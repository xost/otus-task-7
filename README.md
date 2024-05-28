# otus-task-7

приложения деплоятся в пространство auth


Результат прогона тестов
```
newman

nginx forward auth

→ регистрация 1
  POST http://arch.homework/register [200 OK, 151B, 83ms]
  ✓  [INFO] Request: {
	"login": "Johann52",
	"password": "CRkL88CWjmzwHEN",
	"email": "Eino16@yahoo.com",
	"first_name": "Krystina",
	"last_name": "Marquardt"
}

  ✓  [INFO] Response: {"id": 58}

→ логин 1
  POST http://arch.homework/login [200 OK, 227B, 11ms]
  ✓  [INFO] Request: {"login": "Johann52", "password": "CRkL88CWjmzwHEN"}
  ✓  [INFO] Response: {"status":"ok"}

→ проверить данные о пользователе 1
  GET http://arch.homework/auth [200 OK, 359B, 7ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"login":"Johann52","password":"","email":"Eino16@yahoo.com","first_name":"Krystina","last_name":"Marquardt"}
  ✓  test token data

→ получить данные о пользователе 1
  GET http://arch.homework/users/me [200 OK, 275B, 9ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"login":"Johann52","password":"","email":"Eino16@yahoo.com","first_name":"Krystina","last_name":"Marquardt","avatar_uri":"","age":0}
  ✓  test token data

→ обновить данные о пользователе 1
  PUT http://arch.homework/users/me [200 OK, 220B, 14ms]
  ✓  [INFO] Request: {"avatar_uri": "https://cdn.fakercloud.com/avatars/vinciarts_128.jpg", "age": 306}
  ✓  [INFO] Response: {"avatar_uri":"https://cdn.fakercloud.com/avatars/vinciarts_128.jpg","age":306}

→ получить данные о пользователе 1 после обновления
  GET http://arch.homework/users/me [200 OK, 329B, 10ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"login":"Johann52","password":"","email":"Eino16@yahoo.com","first_name":"Krystina","last_name":"Marquardt","avatar_uri":"https://cdn.fakercloud.com/avatars/vinciarts_128.jpg","age":306}
  ✓  test token data

→ логаут 1
  GET http://arch.homework/logout [200 OK, 99B, 5ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response:

→ получить данные после разлогина 1
  GET http://arch.homework/users/me [200 OK, 201B, 9ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"message": "Please go to login and provide Login/Password"}

→ регистрация 2
  POST http://arch.homework/register [200 OK, 151B, 13ms]
  ✓  [INFO] Request: {
	"login": "Jeremy_Feest14",
	"password": "X6Hst3hP38oUu1m",
	"email": "General.Bashirian25@yahoo.com",
	"first_name": "Rolando",
	"last_name": "Kassulke"
}

  ✓  [INFO] Response: {"id": 59}

→ логин 2
  POST http://arch.homework/login [200 OK, 227B, 7ms]
  ✓  [INFO] Request: {"login": "Jeremy_Feest14", "password": "X6Hst3hP38oUu1m"}
  ✓  [INFO] Response: {"status":"ok"}

→ проверить данные о пользователе 1
  GET http://arch.homework/auth [200 OK, 393B, 5ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"login":"Jeremy_Feest14","password":"","email":"General.Bashirian25@yahoo.com","first_name":"Rolando","last_name":"Kassulke"}
  ✓  test token data

→ логаут 2
  GET http://arch.homework/logout [200 OK, 99B, 5ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response:

→ получить данные после разлогина 2
  GET http://arch.homework/users/me [200 OK, 201B, 9ms]
  ✓  [INFO] Request: [object Object]
  ✓  [INFO] Response: {"message": "Please go to login and provide Login/Password"}

┌─────────────────────────┬──────────────────┬──────────────────┐
│                         │         executed │           failed │
├─────────────────────────┼──────────────────┼──────────────────┤
│              iterations │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│                requests │               13 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│            test-scripts │               21 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│      prerequest-scripts │               15 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│              assertions │               30 │                0 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 599ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 804B (approx)                            │
├───────────────────────────────────────────────────────────────┤
│ average response time: 14ms [min: 5ms, max: 83ms, s.d.: 20ms] │
└───────────────────────────────────────────────────────────────┘

```
