# Full eatured Golang App

### End points:
---
>> User Authenticity: 
1. POST http://localhost:8080/api/v1/auth/register 

input:
```json
{
	"name":"nase",
	"email":"nas@gmail.com",
	"password":"nas12345"
}
```
Response:
```json
{
    "id": 1,
    "name": "nase",
    "email": "nas@gmail.com",
    "Password": "$2a$04$wJcYp.0ogKNJqniCKfKMRuxsJ4vPfnV275goiHfYvBYUBkZk8QH4K"
}
```
2. POST http://localhost:8080/api/v1/auth/login

Input:
```json
{
	"name":"nase",
	"email":"nas@gmail.com",
	"password":"nas12345"
}
```
Response:
```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTExNDc2NTEsInVzZXJfaWQiOjF9.WP8lJkZJFXD9fE4UFPcC68nXdCKAqPe6P83khvaTEhM",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTExNDQwNTEsInVzZXJfaWQiOjF9.Ec45_KMIN5HU9sgUWedQgRGCNz5QjJDLHMwmH3aijMg"
}
```
3. GET http://localhost:8080/api/v1/auth/user

Response:
```json
{
    "data": {
        "id": 1,
        "name": "nase",
        "email": "nas@gmail.com"
    },
    "message": "success"
}
```

>> Category
1. GET http://localhost:8080/api/v1/categories

Response:
```json
[
    {
        "id": 1,
        "name": "Drama"
    },
    {
        "id": 2,
        "name": "Advanture"
    }
]
```
2. POST http://localhost:8080/api/v1/categories/

Input:
```json
{
	"name":"Romantic"
}
```
Response:
```json
[
    {
        "id": 1,
        "name": "Romantic"
    },
    {
        "id": 2,
        "name": "Advanture"
    }
]
```
3. PUT http://localhost:8080/api/v1/categories/:categorId

input:
```json
{
	"name":"Drama"
}
```
Response:
```
{
    "id": 1,
    "name": "Drama"
}
```
4. Delete http://localhost:8080/api/v1/categories/:categorId

>> Posts
1. GET http://localhost:8080/api/v1/posts
2. POST http://localhost:8080/api/v1/posts

Input:
```json
{
    "id": 1,
    "title": " this is fourth title--",
    "description": "this is description---",
    "image": "media/images/Screenshot from 2022-04-27 13-50-20.png",
    "category": 1
}
```
Response:
```json
{
    "id": 1,
    "title": " this is fourth title--",
    "description": "this is description---",
    "image": "media/images/Screenshot from 2022-04-27 13-50-20.png",
    "created_at": "2022-04-28T16:25:12.45206321+06:00",
    "updated_at": "2022-04-28T16:25:12.45206321+06:00",
    "category": {
        "id": 2,
        "name": "Advanture"
    }
}
```
3. GET http://localhost:8080/api/v1/posts/:postId

Response:
```json
{
    "id": 1,
    "title": " this is fourth title--",
    "description": "this is description---",
    "image": "media/images/Screenshot from 2022-04-27 13-50-20.png",
    "created_at": "2022-04-28T16:25:12.45206321+06:00",
    "updated_at": "2022-04-28T16:25:12.45206321+06:00",
    "category": {
        "id": 2,
        "name": "Advanture"
    }
}
```
4. PUT http://localhost:8080/api/v1/posts/:postId

Input:
```json
{
    "title": " this is fourth title",
    "description": "this is description---",
    "image": "media/images/Screenshot from 2022-04-27 13-50-20.png",
    "category": 1
}
```
Response:
```json
{
    "id": 1,
    "title": " this is fourth title--",
    "description": "this is description---",
    "image": "media/images/Screenshot from 2022-04-27 13-50-20.png",
    "created_at": "2022-04-28T16:25:12.45206321+06:00",
    "updated_at": "2022-04-28T16:25:12.45206321+06:00",
    "category": {
        "id": 2,
        "name": "Advanture"
    }
}
```
5. DELETE http://localhost:8080/api/v1/posts/:postId