## List of CURL Commands

- Login and get wt token 

```
curl -X POST -H "Content-Type: application/json" -d '{"username": "random"}' http://localhost:8000/login

```

- Send post request to backend service with jwt token 
```
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"email": "random@gmail.com"  ,"name": "random"}' http://localhost:8000/api/entry
```