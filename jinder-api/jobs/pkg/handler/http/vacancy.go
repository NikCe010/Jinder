package http

import (
	vacancy2 "Jinder/jinder-api/jobs/pkg/service/dto/vacancy"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateVacancyHandler(w http.ResponseWriter, r *http.Request) {
	var vacancy vacancy2.Vacancy
	err := json.NewDecoder(r.Body).Decode(&vacancy)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Services.CreateVacancy(vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *Handler) UpdateVacancyHandler(w http.ResponseWriter, r *http.Request) {
	var vacancy vacancy2.Vacancy
	err := json.NewDecoder(r.Body).Decode(&vacancy)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Services.UpdateVacancy(vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *Handler) GetVacancyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	vacancyId, err := uuid.Parse(vars["vacancy_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Services.GetVacancy(vacancyId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetVacanciesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["user_id"])
	log.Debug(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count := r.URL.Query().Get("count")
	if count == "" {
		count = "10"
	}

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}

	user, err := h.Services.GetVacancies(userId, count, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func (h *Handler) DeleteVacancyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	vacancyId, err := uuid.Parse(vars["vacancy_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Services.DeleteVacancy(vacancyId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
