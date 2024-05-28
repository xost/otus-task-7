package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type deltaModel struct {
	Delta int `json:"delta"`
}

type configModel struct {
	dbHost string
	dbPort string
	dbName string
	dbUser string
	dbPass string
	host   string
	port   string
}

const (
	getbalanceTpl    = `SELECT balance FROM account WHERE id=$1`
	updatebalanceTpl = `INSERT INTO account VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET balance = excluded.balance`
)

var (
	getbalanceStmt    *sql.Stmt
	updatebalanceStmt *sql.Stmt
)

func readConf() *configModel {
	cfg := &configModel{
		dbHost: "account-postgresql",
		dbPort: "5432",
		dbName: "accountdb",
		dbUser: "accountuser",
		dbPass: "accountpasswd",
		host:   "0.0.0.0",
		port:   "80",
	}
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if dbHost != "" {
		cfg.dbHost = dbHost
	}
	if dbPort != "" {
		cfg.dbPort = dbPort
	}
	if dbName != "" {
		cfg.dbName = dbName
	}
	if dbUser != "" {
		cfg.dbUser = dbUser
	}
	if dbPass != "" {
		cfg.dbPass = dbPass
	}
	if host != "" {
		cfg.host = host
	}
	if port != "" {
		cfg.port = port
	}
	return cfg
}

func makeDBConn(cfg *configModel) (*sql.DB, error) {
	pgConnString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.dbHost, cfg.dbPort, cfg.dbUser, cfg.dbPass, cfg.dbName,
	)
	log.Println("connection string: ", pgConnString)
	db, err := sql.Open("postgres", pgConnString)
	return db, err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := readConf()

	db, err := makeDBConn(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err = db.PingContext(ctx); err != nil {
		log.Fatal("Failed to check db connection:", err)
	}

	mustPrepareStmts(ctx, db)

	r := mux.NewRouter()

	r.HandleFunc("/account/health", isAuthenticatedMiddleware(health))
	r.HandleFunc("/account/get", isAuthenticatedMiddleware(get))
	r.HandleFunc("/account/deposit", isAuthenticatedMiddleware(deposit)).Methods("PUT")

	bindOn := fmt.Sprintf("%s:%s", cfg.host, cfg.port)
	if err := http.ListenAndServe(bindOn, r); err != nil {
		log.Printf("Failed to bind on [%s]: %s", bindOn, err)
	}
}

func mustPrepareStmts(ctx context.Context, db *sql.DB) {
	var err error

	getbalanceStmt, err = db.PrepareContext(ctx, getbalanceTpl)
	if err != nil {
		panic(err)
	}

	updatebalanceStmt, err = db.PrepareContext(ctx, updatebalanceTpl)
	if err != nil {
		panic(err)
	}

}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "OK"}`))
}

func createbalance(id int) error {
	_, err := updatebalanceStmt.Query(id, 0)
	if err != nil {
		return err
	}
	return nil
}

func getbalance(id int) (int, error) {
	balance := 0
	err := getbalanceStmt.QueryRow(id).Scan(&balance)
	return balance, err
}

func updatebalance(id, delta int) error {
	current, _ := getbalance(id)
	amount := current + delta
	_, err := updatebalanceStmt.Query(id, amount)
	return err
}

func get(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	id, err := strconv.Atoi(headers.Get("X-User-Id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Got wrong header [X-User-Id]: %s", err)
		return
	}
	b, err := getbalance(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("There is no account record for user id [%d]. Create it\n", id)
			if err = createbalance(id); err != nil {
				log.Printf("Failed to create account for user id [%d]: %s\n", id, err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to create account for userID [%d]: %s", id, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"balance":0}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to get account balance for userID [%d]", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"balance":%d}`, b)
}

func deposit(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	id, err := strconv.Atoi(headers.Get("X-User-Id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Got wrong header [X-User-Id]: %s", err)
		return
	}
	d := deltaModel{}
	if err = json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to pasrse data:", err)
		return
	}
	if err = updatebalance(id, d.Delta); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to update balance:", err)
		return
	}
}

func isAuthenticatedMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		fmt.Println(headers)
		if _, ok := headers["X-User-Id"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authenticated"))
			return
		}
		h.ServeHTTP(w, r)
	}
}
