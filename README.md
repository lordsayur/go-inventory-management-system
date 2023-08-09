# Inventory Management System (IMS)

## Web API

To start IMS web api, change directory to api

```bash
cd api
```

Then run the main.go file

```bash
go main.go

```

By default, the system use sqlite3 database when running the program. It is also possible to use in memory database. In api/main.go file, comment out either in memory repository

```go
itemRepository := repositories.NewMemoryItemRepository()

```

or XORM repository

```go
itemRepository := repositories.NewXORMItemRepository()

```

It is recommeded to use Visual Studio Code as IDE and install all recommended extension. Feel free to import collection and enviroment localed in utils folder in order to test the web api

## UI

To start frontend application, move to ui folder

```bash
cd ui
```

then install all dependencies

```bash
yarn
```

start development server

```bash
yarn dev
```
