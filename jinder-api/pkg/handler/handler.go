package handler

import (
	"Jinder/jinder-api/pkg/handler/resume"
	"Jinder/jinder-api/pkg/handler/user"
	"Jinder/jinder-api/pkg/handler/vacancy"
	"Jinder/jinder-api/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/register", user.RegisterHandler).Methods("POST")
	r.HandleFunc("/sign_in", user.SignInHandler).Methods("POST")
	r.HandleFunc("/sign_out", user.SignOutHandler).Methods("POST")

	r.HandleFunc("/users/{user_id}", user.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{user_id}", user.GetUserHandler).Methods("GET")

	u := r.PathPrefix("/users/{user_id}").Subrouter()
	u.HandleFunc("/resumes", resume.CreateResumeHandler).Methods("POST")
	u.HandleFunc("/resumes", resume.GetResumesHandler).Methods("GET")

	u.HandleFunc("/resumes/{resume_id}", resume.GetResumeHandler).Methods("GET")
	u.HandleFunc("/resumes/{resume_id}", resume.UpdateResumeHandler).Methods("PUT")
	u.HandleFunc("/resumes/{resume_id}", resume.DeleteResumeHandler).Methods("DELETE")

	u.HandleFunc("/vacancies", vacancy.CreateVacancyHandler).Methods("POST")
	u.HandleFunc("/vacancies", vacancy.GetVacanciesHandler).Methods("GET")

	u.HandleFunc("/vacancies/{vacancy_id}", vacancy.GetVacancyHandler).Methods("GET")
	u.HandleFunc("/vacancies/{vacancy_id}", vacancy.UpdateVacancyHandler).Methods("PUT")
	u.HandleFunc("/vacancies/{vacancy_id}", vacancy.DeleteVacancyHandler).Methods("DELETE")
}
