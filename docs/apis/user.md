### Introduction

#### `GET` /api/v1/lorem

Find lorem Ipsum text api.
Lorem Ipsum is simply dummy text of the printing and typesetting industry.
<br />

### Sequence Diagram

![](./seq-img/test.png)

#### Request Header

| Key           | <div style="width: 735px;">Value</div> |
| ------------- | -------------------------------------- |
| Content-Type  | application/json                       |
| Authorization | Bearer {your-access-token}             |
| Connection    | keep-alive                             |

#### Request Body

```json
{
  "id": "{id}"
}
```

#### Response

1. Success case

```json
{
  "message": "success",
  "code": 1000
}
```

2. Error case

```json
{
  "message": "error bra bra bra",
  "code": 5000
}
```

#### Change logs

| Date       | <div style="width: 600px;">Description</div> | Developer     |
| ---------- | -------------------------------------------- | ------------- |
| 20/01/2025 | Change API #1                                | Kan Wongsawan |
| 14/02/2025 | Change API #2                                | Kan Wongsawan |
| 05/03/2025 | Change API #3                                | Kan Wongsawan |
