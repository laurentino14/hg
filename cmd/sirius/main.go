package main

import (
	"bitbucket.org/elevatt/sirius/configs"
	"bitbucket.org/elevatt/sirius/internal/infra/data-sources/cockroachgen"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
	db := cockroachgen.GenerateCockrouchDataSource()
	defer db.Close()
	app := configs.SiriusConfig(db)

	log.Fatal(app.Listen(":3000"))
}
