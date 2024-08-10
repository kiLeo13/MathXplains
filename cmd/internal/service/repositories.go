package service

import "MathXplains/internal/domain/sqlite/repository"

var (
	apptmRepo     repository.AppointmentRepository
	configRepo    repository.ConfigRepository
	professorRepo repository.ProfessorRepository
	subjectRepo   repository.SubjectRepository
	userRepo      repository.UserRepository
)

func SetAppointmentRepository(repo *repository.AppointmentRepository) {
	apptmRepo = *repo
}

func SetConfigRepository(repo *repository.ConfigRepository) {
	configRepo = *repo
}

func SetProfessorRepository(repo *repository.ProfessorRepository) {
	professorRepo = *repo
}

func SetSubjectRepository(repo *repository.SubjectRepository) {
	subjectRepo = *repo
}

func SetUserRepository(repo *repository.UserRepository) {
	userRepo = *repo
}
