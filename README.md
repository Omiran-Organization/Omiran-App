# Omiran-App

An open source live streaming application that is developer-oriented and removes the political overhead and censorship that other streaming platforms exhibit. The application's name, "Omiran", is Yoruba for, "alternative". This application, from a user's perspective can be interpreted as an alternative; an alternative to the harshly controlled and revenue focused platforms. We will also implement a 0% deduction of streamers' subscription money; funded by either github sponsors, or ads.

## Features (not developed yet)

- FOSS
- Live websocket stream chat
- Live streaming and future quirk integration
- 0% deduction in subscription revenue to streamer
- User and livestream search yielding optimized responses using elasticsearch

## Tech Stack

- MinIO (potentially)
- NextJS
- React
- React Native 
- Go 
- Go gin router
- MySQL
- GraphQL
- Docker
- RTMP
- TailwindCSS
- Redis (caching)
- SocketIO
- Bash
- Typescript
- Docker-Compose
- Redux
- Apollo-Client

## Installation and Setup

1. Clone the repository
2. Go into `docker-compose.yml` (in the root directory of the project)
3. Locate these environment variables:
```yml
      - MYSQL_ROOT_HOST=
      - MYSQL_USER=
      - MYSQL_PASSWORD=
      - MYSQL_ROOT_PASSWORD=
```

```yml
      - ALGOLIA_ADMIN_KEY=
      - ALGOLIA_SEARCH_KEY=
      - ALGOLIA_APP_ID=
```

4. Fill them with your preferred values (database name, for the time being, must be, `Omiran`, which is already taken care of for you) 

5. Run: `./automate-build.sh`

6. After running `./automate-build.sh` for the first time, run
`docker ps`

7. Locate the PID for mysql, and run this:
`sudo docker exec -it [PID] bash`


7. Happy Hacking!

*note, running `automate-build.sh` will take a long time as it builds everything from scratch*

*directions regarding creating database tables have not been included yet as there is ambiguity in development currently about whether this process of configuration should be abstracted away from the user; algolia configuration directions are soon to come as well*

## Architecture And Design

### [Database Schema](database-schema.md)

### [Chat Architecture](architecture-prototypes/chat.png)

### [GraphQL service](architecture-prototypes/view_data_querying_architecture.png)

### [Streaming Architecture](architecture-prototypes/streaming.png)

### [Container Orchestration Conception](architecture-prototypes/container-orchestration-prototype.png)
