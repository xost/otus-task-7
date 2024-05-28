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

