package main

import (
	"MathXplains/internal/domain/sqlite"
	"MathXplains/internal/domain/sqlite/repository"
	cognito "MathXplains/internal/infrastructure/aws/cognito"
	"MathXplains/internal/routes"
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORS())

	db, err := sqlite.Init()
	if err != nil {
		panic(err)
	}

	e.Static("/", ".")

	apptms := repository.NewAppointmentRepository(db)
	noteRepo := repository.NewNoteRepository(db)
	professorRepo := repository.NewProfessorRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	userRepo := repository.NewUserRepository(db)

	service.SetAppointmentRepository(apptms)
	service.SetNoteRepository(noteRepo)
	service.SetProfessorRepository(professorRepo)
	service.SetSubjectRepository(subjectRepo)
	service.SetUserRepository(userRepo)

	e.POST("/api/users", routes.CreateUser)
	e.POST("/api/users/verify", routes.ConfirmAccount)
	e.POST("/api/users/verify/resend", routes.ResendConfirmation)
	e.POST("/api/users/global-logout", routes.GlobalLogOut)
	e.POST("/api/users/login", routes.LogInUser)
	e.POST("/api/users/refresh", routes.RefreshToken)
	e.GET("/api/users", routes.GetUsers)
	e.GET("/api/users/:id", routes.GetUserByID)
	e.DELETE("/api/users", routes.DeleteSelfUser)

	e.GET("/api/appointments", routes.GetAppointments)
	e.POST("/api/appointments", routes.CreateAppointment)
	e.DELETE("/api/appointments/:id", routes.DeleteAppointment)

	e.GET("/api/notes", routes.GetNotes)
	e.POST("/api/notes", routes.CreateNote)
	e.PUT("/api/notes/:id", routes.PutNote)
	e.DELETE("/api/notes/:id", routes.DeleteNote)

	e.GET("/api/subjects", routes.GetSubjects)
	e.GET("/api/professors", routes.GetProfessors)

	if err := e.Start(":80"); err != nil {
		e.Logger.Fatal(err)
	}
}
