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

## Run With Docker
Please pull image from [registry](https://hub.docker.com/r/rezairfanwijaya/go-todo)

## API Reference
API documentation can view [here](https://documenter.getpostman.com/view/11940636/2s8Z6vZEic)