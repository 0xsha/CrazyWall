package routes

import (
	"back-end/handlers"
	"back-end/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app fiber.Router) {


	// well
	app.Get("/" , func(ctx *fiber.Ctx) error {
		return ctx.SendString("Nothing To See here :)")
	})

	// health check
	app.Get("/api/health" , func(ctx *fiber.Ctx) error {
		return ctx.SendString("Ok")
	})


	// API related routes
	baseAPI := app.Group("/api/v1")


	// Auth there is a 5 request per M limitation to avoid bots and brute forces
	authAPI := baseAPI.Group("/auth", middlewares.Limited())
	authAPI.Post("/login" , handlers.Login )
	authAPI.Post("/register" , handlers.Register )
	authAPI.Post("/logout" , handlers.Logout)


	// Users
	userAPI := baseAPI.Group("/users" , middlewares.JWTProtected())
	userAPI.Get("/user" , handlers.User)


	// Phone search
	phoneAPI := baseAPI.Group("/phone" , middlewares.JWTProtected())
	phoneAPI.Post("/search" , handlers.SearchPhone)

	// Email search
	emailAPI := baseAPI.Group("/email" , middlewares.JWTProtected())
	emailAPI.Post("/search" , handlers.SearchEmail)


	// Domain search
	domainAPI := baseAPI.Group("/domain" , middlewares.JWTProtected())
	domainAPI.Post("/search" , handlers.SearchDomain)
	domainAPI.Post("/history" , handlers.SearchDomain)
	domainAPI.Post("/subdomain" , handlers.SearchDomain)


	// IP search
	ipAPI := baseAPI.Group("/ip" , middlewares.JWTProtected())
	ipAPI.Post("/search" , handlers.SearchIP)


	// UserName search
	userNameAPI := baseAPI.Group("/username" , middlewares.JWTProtected())
	userNameAPI.Post("/search" , handlers.SearchUserName)


	// Name search
	nameAPI := baseAPI.Group("/name" , middlewares.JWTProtected())
	nameAPI.Post("/search" , handlers.SearchName)

	// Leaked databases
	leakAPI := baseAPI.Group("/leak" , middlewares.JWTProtected())
	leakAPI.Post("/search" , handlers.SearchAssociation)


	// Search history endpoints
	historyAPI := baseAPI.Group("/history" , middlewares.JWTProtected())
	historyAPI.Post("/search" , handlers.GetHistoryByID)
	historyAPI.Post("/mine" , handlers.GetMyHistory)


	// Auxiliary endpoints
	auxiliaryAPI := baseAPI.Group("/auxiliary" , middlewares.JWTProtected())
	auxiliaryAPI.Post("/flag" , handlers.GetFlagURL)



	// Live  endpoints
	liveAPI := baseAPI.Group("/live" , middlewares.JWTProtected())
	liveAPI.Post("/search" , handlers.SearchLive)


}