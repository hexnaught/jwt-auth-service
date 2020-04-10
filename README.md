# jwt-auth-service

### Run Local:

- Fill in details in `.env` file.
- `docker run -p 27017:27017 --name mongoc -v /ProgramData/DockerStore:/DockerResources/data/mongodb mongo`
- `go run .`
- `npm run dev`

_Tool for viewing MongoDB data/Schema_

- `docker run -it --rm -p 8081:8081 --link mongoc:mongo mongo-express`

_Building the docker image_

`docker build -t jokerdan/jwt-auth-service:v1.0 .`

_Running the docker image_

`docker run -p 8080:8080 --detach --name auth-service jokerdan/jwt-auth-service`

### Docker/Compose:

`docker-compose buil`
`docker-compose up`

### Notes

#### Invoke API via Powershell

```ps1
Invoke-WebRequest http://localhost:8080/api/v1/user/register `
	-ContentType "application/json" `
	-Method POST `
	-Body '{"username":"test_0001","password":"qwerty","email":"test@example.com"}' | Select-Object -expand RawContent

Invoke-WebRequest http://localhost:8080/api/v1/user/login `
	-ContentType "application/json" `
	-Method POST `
	-Body '{"username":"test_0001","password":"qwerty"}' | Select-Object -expand RawContent
```

```ps1
$token = ""
Invoke-WebRequest http://localhost:8080/api/v1/token/validate `
	-Method GET `
	-Header @{Authorization = "Bearer $token"}
```
