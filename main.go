package main

import (
	"email-service/config"
	"fmt"
	"log"
	"net/http"

	v1 "email-service/controller/v1"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// conn, err := amqp.Dial("amqp://setucdwc:8HPqKaOisQhptp7HARM0S1rUaQeAw2LU@cougar.rmq.cloudamqp.com/setucdwc")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	// ch, err := conn.Channel()
	// failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"forgot-password-email", // name
	// 	true,                    // durable
	// 	false,                   // delete when unused
	// 	false,                   // exclusive
	// 	false,                   // no-wait
	// 	nil,                     // arguments
	// )
	// failOnError(err, "Failed to declare a queue")

	// msgs, err := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )
	// failOnError(err, "Failed to register a consumer")

	// forever := make(chan bool)

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 	}
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever

	e := echo.New()
	e.Use(middleware.Logger())

	apiVersion1 := e.Group("/v1")
	// Routes
	apiVersion1.GET("", HealthCheck)
	apiVersion1.POST("/send", v1.SendEmail)
	apiVersion1.POST("/send/:name", v1.SendWithTemplate)
	apiVersion1.GET("/templates", v1.ListTemplates)
	apiVersion1.GET("/template/version/:name", v1.ListTemplateVersions)

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
