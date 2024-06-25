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

![image](https://github.com/MorZLE/dns-host/assets/122459662/bab92db4-5483-4ec6-a8aa-7b41b7b2dca4)
![image](https://github.com/MorZLE/dns-host/assets/122459662/c00c480b-5af7-445e-942e-b38703691b4b)
![image](https://github.com/MorZLE/dns-host/assets/122459662/82c63c47-3b5e-4f7b-9b9d-64340171f3b3)
