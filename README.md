## Save My Pass application

### WOP

In this application user must save theirs passwords (with encryption)

****
Backend: Go 1.21, Gin-gonic, GORM, air and delve

Database: Postgresql, Redis

Logs, metrics: Prometheus, Grafana

Server, proxy: nginx

Frontend: idk
*****

Dev-mode:
```shell
mkdir grafana
```

```shell
sudo chmod 777 grafana/ -R
```

```shell
cp .env.example .env
```


```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

Production:
```shell
mkdir grafana
```

```shell
sudo chmod 777 grafana/ -R
```

```shell
cp .env.example .env
```


```shell
docker-compose up -d
```