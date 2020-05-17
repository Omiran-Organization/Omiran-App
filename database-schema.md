# Database schema
Eventually we will be making a proper database diagram using some tool. But this will suffice during the planning phase. 
Strings will also need to be given a length limit.

## user
| Field| Type | Relation
|----|---------- |-
| uuid | string
| username | string 
| email_address | string 
| hashed_password | string 
| biography | string 
| profile_picture | string (or maybe BLOB?)

profile_picture holds the URL to the image on AWS s3 bucket.

## user_follows
| Field | Type | Relation
| -|-|-
| follower | string | FK user.uuid
| followee | string | FK user.uuid

**follower** is the person following **followee**.
