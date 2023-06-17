# api-gateway-kong-golang


## Architecture

![Architecture Image](/images/arch.png)

- Users sends requests to the service which are handled by kong
- All requests involving logging in the user are directed to the golang service
- The golang service provides a jwt token
- All other requests are redirected first to golang service for authentication
- Once authenticated it is redirected to the backend service
- Golang service acts as authentication mechanism and python service acts as backend 

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