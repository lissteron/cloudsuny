# API

#### GET /api/v1/view

Request: none

Response:
```json
{
    "errors": null,
    "data": [
        {
            "id": "77e4d43d-5be3-4b7b-9277-809cf18824d3",
            "username": "username-1",
            "avatar": "avatar-1",
            "badges": [
                {
                    "id": "783459e6-ea41-4882-86e9-eab51d0e9182",
                    "user_id": "77e4d43d-5be3-4b7b-9277-809cf18824d3",
                    "type": "cloud",
                    "point": {
                        "x": 0.55,
                        "y": -100.99
                    },
                    "created_at": "2020-07-30T08:48:12.796654876Z"
                }
            ]
        },
        {
            "id": "89569f66-8be2-4f8c-884b-bbfb33198a7b",
            "username": "username-2",
            "avatar": "avatar-2"
        },
        {
            "id": "a463d37f-5eb0-485a-a0df-12844616847a",
            "username": "username-3",
            "avatar": "images/1c9khj4qb2kvbh.png"
        },
        {
            "id": "5a28730f-34ba-4871-887f-e18196bcd80e",
            "username": "username-4",
            "avatar": "images/1c9khj4qb2kvbh.png"
        }
    ]
}
```

// auth from http baseauth username


#### POST /api/v1/user/create

Request:

```json
{
	"username": "username",
	"avatar": "images/1c9khj4qb2kvbh.png"
}
```

Response:

```json
{
    "errors": null,
    "data": {
        "id": "5a28730f-34ba-4871-887f-e18196bcd80e",
        "username": "username-4",
        "avatar": "images/1c9khj4qb2kvbh.png"
    }
}
```
Response with error:
```json
{
    "errors": [
        {
            "code": "0",
            "detail": "user already exists"
        }
    ],
    "data": null
}
```

#### POST /api/v1/badge/create

Request:

```json
{
    "user_id": "77e4d43d-5be3-4b7b-9277-809cf18824d3",
    "type": "cloud",
    "coords": {
        "x": 0,
        "y": 0
    }
}
```

Response:
```json
{
    "errors": null,
    "data": {
        "id": "47f08476-4da3-4742-abe1-ef36ea10b95c",
        "user_id": "77e4d43d-5be3-4b7b-9277-809cf18824d3",
        "type": "cloud",
        "point": {
            "x": 0,
            "y": 0
        },
        "created_at": "2020-07-30T15:43:12.507165526Z"
    }
}
```


#### POST /api/v1/badge/update 

Request:

```json
{
    "id": "47f08476-4da3-4742-abe1-ef36ea10b95c",
    "point": {
        "x": 0.55,
        "y": -100.19
    }
}
```

Response:
```json
{
    "errors": null,
    "data": {
        "id": "47f08476-4da3-4742-abe1-ef36ea10b95c",
        "user_id": "77e4d43d-5be3-4b7b-9277-809cf18824d3",
        "type": "cloud",
        "point": {
            "x": 0.55,
            "y": -100.19
        },
        "created_at": "2020-07-30T15:43:12.507165526Z"
    }
}
```

#### POST /api/v1/image/upload

Request:

field data = binary data

Response:
```json
{
    "errors": null,
    "data": {
        "name": "images/1c9kiacgaj25ud.png"
    }
}
```

#### Error codes

|Name                     |Code  |
|-------------------------|------|
|DatabaseError            |1000  |
|SystemError              |1001  |
|InvalidJSONError         |1002  |
|UserNotFound             |2000  |
|UserAlreadyExists        |2001  |
|BadgeNotFound            |2002  |
|UnknownImgFormat         |2003  |
|EncodeImgFailed          |2004  |
|ValidUserUsernameRequired|200000|
|ValidUserAvatarRequired  |200001|
|ValidBadgeIDRequired     |200002|
|ValidBadgeUserIDRequired |200003|
|ValidBadgeTypeRequired   |200004|
|ValidBadgeTypeIn         |200005|