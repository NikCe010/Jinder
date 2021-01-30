package http

import (
	resume2 "Jinder/jinder-api/jobs/pkg/service/dto/resume"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateResumeHandler(w http.ResponseWriter, r *http.Request) {
	var resume resume2.Resume
	err := json.NewDecoder(r.Body).Decode(&resume)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Services.CreateResume(resume)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *Handler) UpdateResumeHandler(w http.ResponseWriter, r *http.Request) {
	var resume resume2.Resume
	err := json.NewDecoder(r.Body).Decode(&resume)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Services.UpdateResume(resume)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *Handler) GetResumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resumeId, err := uuid.Parse(vars["resume_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Services.GetResume(resumeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetResumesHandler(w http.ResponseWriter, r *http.Request) {
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

	resumes, err := h.Services.GetResumes(userId, count, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resumes)
}

func (h *Handler) DeleteResumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resumeId, err := uuid.Parse(vars["resume_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Services.DeleteResume(resumeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
