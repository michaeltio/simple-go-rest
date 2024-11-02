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
