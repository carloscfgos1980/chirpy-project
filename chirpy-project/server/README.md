

# chirpy API

* Serve the Client
* Endpoint to check the health of the API
* Endpoint to post a rewuest from a webhook
* Endpoint for login the user
* Endpoint for refreshing the token
* Endpoint to revoke the refresh token
* Endpoint to create users
* Enpoint to update users. This is protect so the header requesto should have a Bearer token
* Endpoint to create a chirp
* Endpoint to get all the chirps
* Endpoint to get a chirp passing the chirpId
* Endpoint to delete chirp passing the chirpId.
* Endpoint to check metrics (how many times the web has been visited)
* Endpoint to reset the metrics

## database

This app save the data in postgres database and use goose to run the migrations from SQL files.
To upddate database a new migration has to be created at <sql/schema> and then run the cmd in the terminal with server as root

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

* Run migrations

```bash
goose postgres postgres://<username>:@localhost:5432/<databasename> up
```

Then create a querie in <sql/queries> and run sqlc command in server as root, this will generate the code to run the queries:

```bash
sqlc generate
```

## TEST

in <internal/auth/auth_test> there are 3 test set:

* Check password hash function
* Check Validate JWT function
* Check Get Token

## Login

This will check if the username (email) and password exist in databse. And response with the User data except password, of course and a token and refresh token

## Create a Chirp

This endpoint is protected, it requires a token in the header request. The chirp (message) should not be longer that 140 characters and 3 words ("kerfuffle", "sharbert, "fornax") will substitute for "*****", they are supposed to be bad words

## Get chirps

There is a option to pass a parameter to get the al chirps from a single user and also to sort them by date

## Delete a chirp

This endpoint is protected so it requires a token. Users can only delete their own chirp