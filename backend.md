# Backend Todo List

## GraphQL Service (frontend queries it)

  - [x] Go server with MySQL database transactions serving graphql
  - [ ] Elasticsearch support for stream/user search

## CI/CD 
  
   1. [x] Dockerize frontend 
   2. [ ] Integrate test automation and deploy automation scripts
   3. [x] Implement localhost automation scripts
   4. [ ] Implement container orchestration

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
  
## Stream Chat

  1. [ ] Implement stream chat with socketio 
  2. [ ] Implement unicode/utf-8 support (for emotes)
  3. [ ] Implement emote support
  4. [ ] Implement message deletion by moderator
  5. [ ] Implement special icons for admins in chat (backend involvement will be necessary for this)
  6. [ ] Somewhere down the line, implement hype train stuff + integrated GIFs/chat animations?!

## Auth system

  1. [x] Implement auth
  2. [x] Implement login system on backend with redis sessions 
  3. [x] Implement cookie assignment with redis  

## Object transactions

  1. [ ] Implement MinIO client for profile picture implementation
  2. [ ] Implement link storage in profile_picture column for easy image attainment
