# GO-TODO

Simple todolist application using golang. There are two main domains, namely `Todo` and `Activity`. One activity can have many todos.


## Run Locally

#### Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/Go-ToDo.git
```

#### Go to the project directory

```bash
  cd Go-Todo
```

#### Make sure it's using mysql and create a new database to use in the ENV later
#### Set ENV with termainal or command prompt
#### example
```bash
  set MYSQL_USER=root
  set MYSQL_PASSWORD=
  set MYSQL_HOST=127.0.0.1
  set MYSQL_DBNAME=go_todo
```
#### Run application
```bash
  go run main.go
```




## API Reference

### ACTIVITY DOMAIN
#### Create New Activity

```http
  POST https://localhost:3030/activity-groups
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title` | `string` | **Required** |
| `email` | `string` | **Required** |

##### Response Success
```bash
{
    "status": "Success",
    "message": "Success",
    "data": {
        "id": 9,
        "title": "test",
        "email": "test@gmail.com",
        "updatedAt": "2022-12-23T23:48:43.1320686+07:00",
        "createdAt": "2022-12-23T23:48:43.1320686+07:00"
    }
}
```

##### Response Failed
```bash
{
    "status": "Bad Request",
    "message": [
        "error on filed: Title, condition: required"
    ]
}
```

#### Get All Activity
```http
  GET https://localhost:3030/activity-groups
```
##### Response Success
```bash
{
    "status": "Success",
    "message": "Success",
    "data": [
        {
            "id": 1,
            "title": "test-lagi",
            "email": "test@gmail.com",
            "updatedAt": "2022-12-24T11:41:56.138+07:00",
            "createdAt": "2022-12-24T11:41:56.138+07:00"
        },
        {
            "id": 2,
            "title": "test-lagi",
            "email": "test@gmail.com",
            "updatedAt": "2022-12-24T11:42:07.857+07:00",
            "createdAt": "2022-12-24T11:42:07.857+07:00"
        }
    ]
}
```


#### Get Activity By ID

```http
  GET https://localhost:3030/activity-groups/1
```

| Param | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** ID Activity |

##### Response Success
```bash
{
    "status": "Success",
    "message": "Success",
    "data": {
        "id": 1,
        "title": "test-lagi",
        "email": "test@gmail.com",
        "updatedAt": "2022-12-24T11:41:56.138+07:00",
        "createdAt": "2022-12-24T11:41:56.138+07:00"
    }
}
```

##### Response Failed
```bash
{
    "status": "Not Found",
    "message": "Activity with ID 90 Not Found"
}
```

#### Delete Activity By ID

```http
  DELETE https://localhost:3030/activity-groups/1
```

| Param | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** ID Activity |

##### Response Success
```bash
{
    "status": "Success",
    "message": "Success",
    "data": {}
}
```

##### Response Failed
```bash
{
    "status": "Not Found",
    "message": "Activity with ID 1 Not Found"
}
```
#### Update Activity By ID

```http
  PATCH https://localhost:3030/activity-groups/1
```

| Param | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int` | **Required** ID Activity |

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title` | `int` | **Required** |

##### Response Success
```bash
{
    "status": "Success",
    "message": "Success",
    "data": {
        "id": 1,
        "title": "update",
        "email": "test@gmail.com",
        "updatedAt": "2022-12-24T21:56:42.499+07:00",
        "createdAt": "2022-12-24T22:20:43.729+07:00"
    }
}
```

##### Response Failed ID
```bash
{
    "status": "Not Found",
    "message": "Activity with ID 1 Not Found"
}
```
##### Response Failed Body
```bash
{
    "status": "Bad Request",
    "message": [
        "error on filed: Title, condition: required"
    ]
}
```


