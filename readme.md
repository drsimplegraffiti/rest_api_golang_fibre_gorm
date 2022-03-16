#### Initialize your project

> go mod init "github.com/drsimplegraffit/fibre-api"

#### Install dependencies

> go get "gorm.io/driver/sqlite"
> go get "gorm.io/gorm"
> go get "github.com/gofiber/fiber/v2"

#### Live reload for go apps

[here](https://github.com/cosmtrek/air)
To install:
`curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin`

#### Check if air is installed

> air -v

---

#### Add your air.toml file

---

#### Create a server

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3002"))
}
```

#### Run your server

> go run main.go

---

#### Test your app

> 127.0.0.1:3002/api

#### To run with airflow, type

> air

#### Debugging your golang setup

> If you have issues with:
> `exec: "gcc": executable file not found in %PATH%`

Download this `https://jmeubank.github.io/tdm-gcc/download/`

#### Install sqlite extension on vscode
