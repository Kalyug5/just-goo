package router

import (
	"github.com/Kalyug5/just-goo/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/",controllers.Home)
	app.Post("/api/todos",controllers.GetTodos)
	app.Post("/api/todo/:id",controllers.GetOneTodo)
	app.Post("/api/todo",controllers.CreateTodo)
	app.Put("/api/todo/:id",controllers.UpdateTodo)
	app.Delete("/api/todo/:id",controllers.DeleteOneTodo)
	app.Delete("/api/todo",controllers.DeleteTodos)
	app.Post("/api/generate-itinerary",controllers.CreateTravelIternery)
	app.Post("/api/trip",controllers.GetTrip)
	app.Post("/api/trips",controllers.GetAllTrip)
	app.Delete("/api/trip/:id",controllers.DeleteOneTrip)
	app.Post("/api/sign-up",controllers.Register)
	app.Post("/api/sign-in",controllers.Login)
	app.Get("/api/user",controllers.User)
	app.Post("/api/logout",controllers.Logout)

	return app

}