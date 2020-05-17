# Database schema
Eventually, we will be making a proper database diagram using some tool. But, this will suffice during the planning phase. 

## User 

| Field| Type | Relation
|----|---------- |-
| uuid | VARCHAR(36)
| username | string 
| email_address | string 
| password | string 
| description | string 
| profile_picture | (undecided)

(all fields will be hashed with salts for sake of user security)

## user_follows

| Field | Type | Relation
| -|-|-
| follower | string | FK user.uuid
| followee | string | FK user.uuid

**follower** is the person following **followee**.
