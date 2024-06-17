# server-dns-host
### Представляет собой grpc сервер для изменения имени хоста и серверов DNS, и cli-клиент для взаимодествия с ним.


## Конфигурация сервера
``` yaml
#srv/config/config.yaml
grpc:
  port: 44044
dns:
  pathResolve: "etc/resolv.conf"
```
### ENV
```env
GRPC_PORT = 44004
```


## Конфигурация клиента
``` yaml
grpc:
  host: 127.0.0.1:44044
```
```env
GRPC_HOST = 127.0.0.1:44044
```

## Примеры использования

