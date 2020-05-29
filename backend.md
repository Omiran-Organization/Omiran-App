# Backend Todo List

## GraphQL Service (frontend queries it)

  - [x] Go server with MySQL database transactions serving graphql
  - [ ] Elasticsearch support for stream/user search

## CI/CD 

   1. [ ] Dockerize application monolithically
   2. [ ] Integrate test automation and deploy automation scripts

## Database CRUD
   
   1. [x] Create database and prerequisite tables with `backend/database-boilerplate.sql`
   2. [x] Connect to database via go server (will probably use a `.yml` config for credentials and other necessary data)
   3. [x] Perform Read and Write of data from both tables with successful struct scanning for programmatic use of data
   
## Live Streaming Service

  1. [ ] Implement OBS integration
  2. [ ] Implement connection to backend service with something similar to Twitch's streamkey solution
  3. [ ] Implement media streaming with either RTMP or WebRTC
  4. [ ] Implement auth resolution feature
  5. [ ] Allow users to choose reception resolution
  
