# api-gateway-kong-golang

- This is a api gateway microservices proect
- It involves following technologies
      - Kong
      - Golang
      - Python

## Initial Configuration
- Add your Ip Address to gateway in config/kong.yaml

```
extra_hosts:
      - "host.docker.internal:<IP ADDR>"
```

- Set password for postgres in kong-gateway/POSTGRES_PASSWORD file

- Install python dependencies
```
pip3 install -r requirements.txt
```

- Install golang dependecies
```
go mod download
```

## Running application
- Run python backend api
```
python3 backend/app.py
```

- Run golang server
```
go run main.go
```

- Run kong proxy
```
in kong-gateway directory
docker-compose up 
```