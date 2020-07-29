# API

GET /api/v1/view

Request: none

Response:
```json
{
    "data": [
        {
            "username": "name",
            "avatar": "http://image_url",
            "badges": [
              {
                "type": "cloud",
                "date": "date time ISO 8601",
                "coords": {
                  "x": 0.0,
                  "y": 0.0,
                }
              },
              {
                "type": "sun",
                "date": "date time ISO 8601",
                "coords": {
                  "x": 0.0,
                  "y": 0.0,
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


/api/v1/user/create

Request:

```json
{
    "username": "some name",
    "avatar": "img.jpg"
}

```

/api/v1/badge/add 
```json
{
  "username": "name",
  "type": "cloud|sun", // enum
  "date": "date time ISO 8601",
  "coords": {
    "x": 0.0,
    "y": 0.0,
  }
}

```

/api/v1/image/upload

Request:

binary data

Response:
```json
{
    "data": "image/path.jpg"
}
```