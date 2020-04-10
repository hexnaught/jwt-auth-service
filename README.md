# jwt-auth-service

## Running the App

**Note: Currently, this has not been optimized for security or tailored for a production environment yet.**

### Run Local:

Add the following code to the top of `main.go`:

```go
// Load .env file, this is for locally running the app only
err := godotenv.Load()
if err != nil {
	log.Fatal("error loading .env file")
}
```

Do the following actions:

- Fill in details in `.env` file.
- `docker run -p 27017:27017 --name mongoc -v /ProgramData/DockerStore:/DockerResources/data/mongodb mongo`
- `go run .`
- `npm run dev`

_Mongo-Express - Tool for viewing MongoDB data/Schema_

`docker run -it --rm -p 8081:8081 --link mongoc:mongo mongo-express`

_Building the docker image_

`docker build -t jokerdan/jwt-auth-service:v1.0 .`

_Running the docker image_

`docker run -p 8080:8080 --detach --name auth-service jokerdan/jwt-auth-service`

### Docker/Compose:

```ps1
> docker-compose build
> docker-compose up
```

The stack will come up with 3 containers running:
- authservice _(Running the build Go binary)_
- mongodb _(The database container)_
- monoex _(mongo-express to allow viewing mongo via browser frontend)_

The mongo-express container is not required for the service to function, I have included it in the stack to make my life a bit easier while testing/developing.

## Notes

### Invoke API via Powershell

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

### TODO

See the `TODO` file in the root of the repostiroy.

