# API

#### GET /api/v1/view

Request: none

Response:
```json
{
    "data": [
        {
            "id": "uuid",
            "username": "name",
            "avatar": "http://image_url",
            "badges": [
              {
                "id": "uuid",
                "type": "cloud",
                "created_at": "date time ISO 8601",
                "point": {
                  "x": 0.0,
                  "y": 0.0
                }
              },
              {
                "id": "uuid",
                "type": "sun",
                "created_at": "date time ISO 8601",
                "point": {
                  "x": 0.0,
                  "y": 0.0
                }
              }
            ]
        }
    ],
    "errors": [
        {
            "code": "13313312",
            "detail": "some error text"
        }
    ]
}
```

// auth from http baseauth username


#### POST /api/v1/user/create

Request:

```json
{
    "username": "some name",
    "avatar": "img.jpg"
}

```

Response:

```json
{
    "id":   "uuid",
    "username": "some name",
    "avatar": "img.jpg"
}

```

#### POST /api/v1/badge/add 

Request:

```json
{
  "user_id": "uuid",
  "type": "cloud|sun", // enum
  "coords": {
    "x": 0.0,
    "y": 0.0
  }
}
```

Response:
```json
{
  "id": "uuid",
  "username": "name",
  "type": "cloud|sun", // enum
  "created_at": "date time ISO 8601",
  "coords": {
    "x": 0.0,
    "y": 0.0
  }
}
```


#### POST /api/v1/badge/update 

Request:

```json
{
  "id": "uuid",
  "coords": {
    "x": 0.0,
    "y": 0.0
  }
}
```

#### POST /api/v1/image/upload

Request:

binary data

Response:
```json
{
    "data": "image/path.jpg"
}
```
