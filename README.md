# bukaresep-go

A simple go package to manage food recipes

## Dependency

bukaresep-go needs a ```SQLlite3``` database to as recipe data store and ```xorm``` dependency as orm database.

## Installation

Add this package to your Golang application. This need using


```bash
go get github.com/imamfzn/bukaresep-go
```

## Scope

bukaresep should be able to:

* add a new recipe,
* update a recipe,
* retrieve detail of a recipe,
* list all recipe,
* delete recipe.

A recipe consists of:

* name (string, required),
* description (string, required),
* ingredients (string, required),
* instructions (string, required),
* id (integer, required),

## Usage

Set up a service

```golang
import (
  "log",
  "github.com/go-xorm/xorm"
  "github.com/imamfzn/bukaresep-go"
  "github.com/imamfzn/bukaresep-go/repository"
  _ "github.com/mattn/go-sqlite3"
)

// create database instance using xorm
// in this example we will use sqlite driver
db, err := xorm.NewEngine("sqlite3", "file::memory:")

if err != nil {
  log.Fatal(err)
}

// syncronize recipe entity to database
err = db.Sync(new(entity.Recipe))

if err != nil {
  return nil, err
}

// create repo for data access
// using db instance before
repo, err := repository.NewXormRepository(db)

if err != nil {
  log.Fatal(err)
}

service := bukaresep.NewService(repo)
```

Then, you can use `bukaresep service` to manage your recipes.

### Add a new recipe

```golang
recipe, err := service.AddRecipe("Chicken Katsu", "Delicious oriental food", "Chicken, egg, salt", "Just merge all ingredients")

if err != nil {
  // Recipe is most likely invalid; a recipe requred all fields
  log.Fatal(err)
}
// recipe will contain a Recipe object from database.
```

### Update a Recipe

```golang
recipe.name = "Chicken Katsu v2.0"
recipe, err = service.UpdateRecipe(recipe)

if err != nil {
  // Recipe is most likely invalid; a recipe requred all fields
  log.Fatal(err)
}
```

### Retrieve detail of a recipe

```golang
recipe, err = service.GetRecipe(1) // the recipe's id

if err != nil {
  // A database error has occured
  log.Fatal(err)
}
```

### List all recipes

```golang
recipes, err = service.GetAllRecipe()

if err != nil {
  // A database error has occured
  log.Fatal(err)
}
```

### Delete a recipe

```golang
err = service.DeleteRecipe(recipe)

if err != nil {
  // A database error has occured
  log.Fatal(err)
}
```

## Onboarding & Development Guide

### Prerequisite

1. Git
2. Golang 1.11 or later
4. SQLite3

### Setup

1. Clone this repository
```bash
git clone -b development https://github.com/imamfzn/bukaresep-go.git
```

### Development Guide

1. To run unit test in all package:
```bash
go test ./...
```
2. To run test-coverage:
```bash
go test ./... -coverprofile cp.out
```
3. To show test-coverage on your browser:
```bash
go tool cover -html=cp.out
```
4. To run linter test:
```bash
golint ./...
```