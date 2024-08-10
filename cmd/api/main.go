package main

import (
	"MathXplains/internal/domain/sqlite"
	"MathXplains/internal/domain/sqlite/repository"
	cognito "MathXplains/internal/infrastructure/aws/cognito"
	"MathXplains/internal/jobs"
	"MathXplains/internal/routes"
	"MathXplains/internal/service"
	"github.com/robfig/cron"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	cognito.NewCognitoClient(os.Getenv("COGNITO_CLIENT_ID"))
	e := echo.New()
	c := cron.New()

	err := c.AddFunc("@hourly", jobs.UpdateDollarExchange)
	if err != nil {
		panic(err)
	}

	db, err := sqlite.Init()
	if err != nil {
		panic(err)
	}

	e.Static("/", ".")

	userRepo := repository.NewUserRepository(db)
	configRepo := repository.NewConfigRepository(db)
	apptms := repository.NewAppointmentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	professorRepo := repository.NewProfessorRepository(db)

	jobs.SetConfigRepo(configRepo)
	service.SetUserRepository(userRepo)
	service.SetConfigRepository(configRepo)
	service.SetConfigRepository(configRepo)
	service.SetAppointmentRepository(apptms)
	service.SetSubjectRepository(subjectRepo)
	service.SetProfessorRepository(professorRepo)

	e.POST("/api/users", routes.CreateUser)
	e.POST("/api/users/verify", routes.ConfirmAccount)
	e.POST("/api/users/login", routes.LoginUser)
	e.POST("/api/users/refresh", routes.RefreshToken)
	e.GET("/api/users", routes.GetUsers)
	e.GET("/api/users/:id", routes.GetUserByID)

	e.GET("/api/appointments", routes.GetAppointments)
	e.POST("/api/appointments", routes.CreateAppointment)
	e.DELETE("/api/appointments/:id", routes.DeleteAppointment)

	e.GET("/api/subjects", routes.GetSubjects)
	e.GET("/api/professors", routes.GetProfessors)

	e.GET("/api/sales", routes.GetSales)
	e.PATCH("/api/sales", routes.PatchSalesCount)

	e.GET("/api/dollar-exchange-rate", routes.GetDollarExchange)

	c.Start()

	if err := e.Start(":80"); err != nil {
		e.Logger.Fatal(err)
	}
}
