# Database schema
Eventually, we will be making a proper database diagram using some tool. But, this will suffice during the planning phase. 

## User 

| Field| Type | Relation
|----|---------- |-
| uuid | string
| username | string 
| email_address | string 
| hashed_password | string 
| biography | string 
| profile_picture | string (or maybe BLOB?)

## user_follows

| Field | Type | Relation
| -|-|-
| follower | string | FK user.uuid
| followee | string | FK user.uuid

**follower** is the person following **followee**.
