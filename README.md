# jwt-auth-service

## Running the App

**Note: Currently, this has not been optimized for security or tailored for a production environment yet.**

### Environment Variables

There are some default environment variables set in the `docker-compose.yml` file, some of the optional and/or secret variables are not currently in there and would have to be manually added to the file and/or to the machine. See below for a full list of supported environment variables.

| Environment Variable | Example Value | Notes |
| -------------------- | ------------- | ----- |
| `MONGO_PORT` | `27017` | The port the mongo container will run on. If changed, you also need to change the port exposed in the `docker-compose` file. |
| `MONGO_HOST` | `mongodb | 172.16.X.X` | The address of the mongo instance, the service uses this to connect to mongo. If connecting to an instance of mongo not in the docker-compose stack, i.e. on the local network, you also need to set the following env vars; `MONGO_USERNAME`, `MONGO_PASSWORD`, `MONGO_AUTH_SOURCE`. |
| `MONGO_DB_NAME` | `authservice` | The name of the mongo database the service will use |
| `MONGO_USERNAME` | `auth_service_user` | Username used to authenticate with mongo |
| `MONGO_PASSWORD` | `auth_5erv1ce_p4ssw0rd` | Password used to authenticate with mongo |
| `MONGO_AUTH_SOURCE` | `admin` | Mongo auth source is the mongo database holding credentials to allow connections, likely defaulted to `admin` |
| `BCRYPT_COST_FACTOR` | `12` | The Default value, and minimum the service will use is 10, Max value is 31. Prefer: 12 \| Cost of 12 = 200ms Response time, every cost increase of 1 roughly doubles response time. |
| `JWT_SECRET` | `jwt_5ecr3t` | Secret used for generating and verifying JWT signatures. See **https://jwt.io** for more information. |
| `JWT_TTL` | `15` | Time to live in minutes for the token, tokens will expire after this time has passed. |
| `CERT_DOMAINS` | `example.com,test.example.com,example.org` | A comma separated list of domains for `autotls` to attempt to create certificates for. |
| `GIN_MODE` | `release | debug | test` | Set the mode for GIN to use. |
| `DEBUG` | `true | false` | Not currently used. |

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
- mongoex _(mongo-express to allow viewing mongo via browser frontend)_

The mongo-express container is not required for the service to function, I have included it in the stack to make my life a bit easier while testing/developing.

## Notes

### Invoke API via PowerShell

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

See the `TODO` file in the root of the repository.

