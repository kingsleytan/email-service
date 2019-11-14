package main

import (
	// "database/sql"

	"email-service/config"
	"email-service/model"
	"fmt"
	"log"
	"net/http"

	v1 "email-service/controller/v1"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*model.Mail)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateRecord :
func CreateRecord(db *pg.DB, i interface{}) error {
	err := db.Insert(i)
	if err != nil {
		panic(err)
	}
	return err
}

func main() {

	db := pg.Connect(&pg.Options{
		User: config.DBUserName,
	})
	fmt.Println("db:", db)
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	apiVersion1 := e.Group("/v1")
	// Routes
	apiVersion1.GET("", HealthCheck)
	apiVersion1.POST("/send", v1.SendEmail)
	apiVersion1.POST("/send/:name", v1.SendWithTemplate)
	apiVersion1.GET("/templates", v1.ListTemplates)
	apiVersion1.GET("/template/:name/version", v1.ListTemplateVersions)

	// Start server
	e.Logger.Fatal(e.Start(":2000"))

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	fmt.Println("ENV", config.Env)
	fmt.Println("MailgunDomain", config.MailgunDomain)
	fmt.Println("MailgunKey", config.MailgunKey)
}

// HealthCheck :
func HealthCheck(c echo.Context) error {
	fmt.Println("debug 1")
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("API is working fine, ENV: %s", config.Env),
	})
}
