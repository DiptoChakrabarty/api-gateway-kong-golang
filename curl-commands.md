## List of CURL Commands

- Login and get Jwt token 

```
curl -X POST -H "Content-Type: application/json" -d '{"username": "random"}' http://localhost:8000/login

```
This redirects all /login requests to the /login route in golang service based on the config/kong.yaml login-route.

- Send post request to backend service with jwt token 
```
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"email": "random@gmail.com"  ,"name": "random"}' http://localhost:8000/api/entry
```
This redirects all / requests to the / route in golang service based on the config/kong.yaml all-routes.