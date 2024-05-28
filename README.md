# otus-task-7

приложения деплоятся в пространство orders


Зарегистрируем пользователя:
```
$curl -v -X POST http://arch.homework/register -d '{"login":"stas","password":"password","email":"xost43@gmail.com","first_name":"Stas","last_name":"Khokhlov"}'
```
Ответ сервиса:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> POST /register HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Content-Length: 108
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:33:39 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 10
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact
{"id": 12}
```
Авторизуемся:
```
$curl -v -X POST http://arch.homework/login -d '{"login":"stas","password":"password"}'
```
Ответ сервера:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> POST /login HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Content-Length: 38
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:37:22 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 15
< Connection: keep-alive
< Set-Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee; HttpOnly
<
* Connection #0 to host arch.homework left intact
{"status":"ok"}
```
"Положим" деньги на свят пользователя:
```
$curl -v -X PUT --cookie session_id=435f1190-913f-468b-a694-bb6f0737c6ee http://arch.homework/account/deposit -d '{"delta":100}'
```
Ответ сервера:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> PUT /account/deposit HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee
> Content-Length: 13
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:40:42 GMT
< Content-Length: 0
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact

```
Проверим баланс средств:
```
$curl -v  --cookie session_id=435f1190-913f-468b-a694-bb6f0737c6ee http://arch.homework/account/get
```
Ответ сервера:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> GET /account/get HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:40:51 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 15
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact
{"balance":100}
```
Создадим заказ:
```
$curl -v -X POST --cookie session_id=435f1190-913f-468b-a694-bb6f0737c6ee http://arch.homework/orders/create -d '{"item":"some stuff","amount":30}'
```
Ответ сервера:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> POST /orders/create HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee
> Content-Length: 33
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:44:39 GMT
< Content-Length: 0
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact

```
Проверим что баланс средств на счету изменился:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> GET /account/get HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee
>
< HTTP/1.1 200 OK
< Date: Tue, 28 May 2024 15:46:09 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 14
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact
{"balance":70}
```
Так как сервис нотификаций не "торчит" наружу, проверим создание нотификации в базе сервиса:
```
$kubectl get pods | grep notif |grep post
```
```
notif-postgresql-0         1/1     Running   0              43m
```
```
$kubectl exec pod/notif-postgresql-0 -it -- /bin/bash
```
```
psql -h localhost -U notifuser notifdb
```
```
select * from notif;
```
```
 id | userid |                  message
----+--------+--------------------------------------------
  4 |     12 | Successfully created order with some stuff
(1 row)
```
Попробуем создать заказ с в суммой, превышающей количество средств на счету:
```
$curl -v -X POST --cookie session_id=435f1190-913f-468b-a694-bb6f0737c6ee http://arch.homework/orders/create -d '{"item":"some stuff","amount":3000}'
```
Ответ сервиса:
```
* Host arch.homework:80 was resolved.
* IPv6: (none)
* IPv4: 192.168.49.2
*   Trying 192.168.49.2:80...
* Connected to arch.homework (192.168.49.2) port 80
> POST /orders/create HTTP/1.1
> Host: arch.homework
> User-Agent: curl/8.5.0
> Accept: */*
> Cookie: session_id=435f1190-913f-468b-a694-bb6f0737c6ee
> Content-Length: 35
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 402 Payment Required
< Date: Tue, 28 May 2024 16:48:03 GMT
< Content-Length: 0
< Connection: keep-alive
<
* Connection #0 to host arch.homework left intact
```
Проверим, что создалась нотификация с ошибкой создания заказа:
```
select * from notif;
```
```
 id | userid |                             message
----+--------+-----------------------------------------------------------------
  4 |     12 | Successfully created order with some stuff
  5 |     12 | Failed to create order due to insufficient funds in the account
(2 rows)
```
