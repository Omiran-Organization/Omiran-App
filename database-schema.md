# Database schema
Eventually, we will be making a proper database diagram using some tool. But, this will suffice during the planning phase. 

## User 

| Field| Type | Relation
|----|---------- |-
| uuid | VARCHAR(36) PRIMARY KEY
| username | TEXT 
| email | VARCHAR(255)
| password | TEXT NOT NULL
| description | TEXT 
| profile_picture | VARCHAR(2083)

(all fields will be hashed with salts for sake of user security)

## Follows

| Field | Type | Relation
| -|-|-
| uuid | FOREIGN KEY VARCHAR(36) NOT NULL | FK User.uuid
| user_following | VARCHAR(36) NOT NULL | FK User.uuid

**uuid** is the person following **user_following**.
The relationship in the `Follows` table, follows a one to one schema. 
