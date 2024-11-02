# Simple GO-REST APIs

- This API is only for study purposes.

## ENDPOINTS

### Get All Todo List

- **Method**: GET
- **Endpoint**: `/todos`
- **Request**:

```json
{}
```

- **Response**

```json
[
    {
        "id" : int,
        "item" : string,
        "isComplete" : bool,
    }
]
```

### Get Todo List by id

- **Method**: GET
- **Endpoint**: `/todos/{id}`
- **Request**:

```json
{}
```

- **Response**

```json
{
    "id" : int,
    "item" : string,
    "isComplete" : bool,
}
```

### Add New List

- **Method**: POST
- **Endpoint**: `/todos`
- **Request**:

```json
{
    "item" : string,
    "isComplete" boolean,
}
```

- **Response**

```json
{
    "message" : string,
}
```

### update list

- **Method**: PATCH
- **Endpoint**: `/todos/{id}`
- **Request**:

```json
{
    "item" : string,
    "isComplete" boolean,
}
```

- **Response**

```json
{
    "message" : string,
}
```

### Add New List

- **Method**: DELETE
- **Endpoint**: `/todos/{id}`
- **Request**:

```json
{}
```

- **Response**

```json
{
    "message" : string,
}
```
