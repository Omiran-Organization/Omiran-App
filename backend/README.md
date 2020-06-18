# Backend

The backend consists of 4 components right now. 
- Go web server
- MySQL database
- Redis database
- Nginx container (handles streaming)

As of now you have to run them all on localhost. The nginx container is not required 
to use the web server, the others are required. It is also not currently integrated 
with the server (work in progress)

# Config file
The config file is for the database credentials and should be called `config.yaml`.
It has the following format 
```
user: 'username'
password: 'password'
port: 3306
```

# Database schema
The database schema is located in the `database-schema.sql` file. Currently you have to manually 
build the database, and rebuild it whenever there has been changes. Eventually we will implement 
migrations. 

# Running the backend
To start the Go server, use `go run main.go`
MySQL and Redis will have to be started first. 