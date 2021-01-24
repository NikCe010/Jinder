package handler

import (
	"Jinder/jinder-api/pkg/service/dto"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) CreateResumeHandler(w http.ResponseWriter, r *http.Request) {
	var resume dto.Resume
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
	var resume dto.Resume
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

	count := 10
	offset := 0
	countParam := r.URL.Query().Get("count")
	log.Print(countParam)
	if countParam != "" {
		count, err = strconv.Atoi(countParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	offsetParam := r.URL.Query().Get("offset")
	log.Print(offsetParam)
	if offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
