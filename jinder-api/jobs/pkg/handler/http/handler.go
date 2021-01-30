package http

import (
	"Jinder/jinder-api/jobs/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("TestHandler"))
	}).Methods("GET")

	r.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		h.RegisterHandler(writer, request)
	}).Methods("POST")

	r.HandleFunc("/sign_in", func(writer http.ResponseWriter, request *http.Request) {
		h.SignInHandler(writer, request)
	}).Methods("POST")

	r.HandleFunc("/sign_out", func(writer http.ResponseWriter, request *http.Request) {
		h.SignOutHandler(writer, request)
	}).Methods("POST")

	r.HandleFunc("/users/{user_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.UpdateUserHandler(writer, request)
	}).Methods("PUT")

	r.HandleFunc("/users/{user_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.GetUserHandler(writer, request)
	}).Methods("GET")

	u := r.PathPrefix("/users/{user_id}").Subrouter()
	u.Use(h.authorizationMiddleware)
	u.Use(h.loggingMiddleware)

	u.HandleFunc("/resumes", func(writer http.ResponseWriter, request *http.Request) {
		h.CreateResumeHandler(writer, request)
	}).Methods("POST")

	u.HandleFunc("/resumes", func(writer http.ResponseWriter, request *http.Request) {
		h.GetResumesHandler(writer, request)
	}).Methods("GET")

	u.HandleFunc("/resumes/{resume_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.GetResumeHandler(writer, request)
	}).Methods("GET")

	u.HandleFunc("/resumes/{resume_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.UpdateResumeHandler(writer, request)
	}).Methods("PUT")

	u.HandleFunc("/resumes/{resume_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.DeleteResumeHandler(writer, request)
	}).Methods("DELETE")

	u.HandleFunc("/vacancies", func(writer http.ResponseWriter, request *http.Request) {
		h.CreateVacancyHandler(writer, request)
	}).Methods("POST")

	u.HandleFunc("/vacancies", func(writer http.ResponseWriter, request *http.Request) {
		h.GetVacanciesHandler(writer, request)
	}).Methods("GET")

	u.HandleFunc("/vacancies/{vacancy_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.GetVacancyHandler(writer, request)
	}).Methods("GET")

	u.HandleFunc("/vacancies/{vacancy_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.UpdateVacancyHandler(writer, request)
	}).Methods("PUT")

	u.HandleFunc("/vacancies/{vacancy_id}", func(writer http.ResponseWriter, request *http.Request) {
		h.DeleteVacancyHandler(writer, request)
	}).Methods("DELETE")

	return r
}
