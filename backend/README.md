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

# Nginx server
The nginx server currently connects to the Go server on `localhost:8080/streamauth`.
This should probably be replaced with some CLI argument or config file in the future.
The stream itself runs on `localhost:8008` as to not collide with the Go server on `localhost:8080`.

```
docker build --tag streamserver .
docker run -it --network host --rm streamserver
```

When running on localhost `--network host` is required to authenticate against the 
Go server. Else it can't find the localhost of the host machine.

# Stream example
### OBS Configuration
* Stream Type: `Custom Streaming Server`
* URL: `rtmp://localhost:1935/stream`
* Stream Key: `<username>?psk=<private_stream_key>`

### Watch Stream
**Replace `<username>` with the actual username**  
* In Safari, VLC or any HLS player, open:
```
http://<server ip>:8008/live/<username>.m3u8
```
* Example Playlist: `http://localhost:8008/live/<username>.m3u8`. This can be pasted into the VideoJS Player below.
* [VideoJS Player](https://video-dev.github.io/hls.js/stable/demo/?src=http%3A%2F%2Flocalhost%3A8080%2Flive%2Fhello.m3u8)
* FFplay: `ffplay -fflags nobuffer rtmp://localhost:1935/stream/<username>`
