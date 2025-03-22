package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sgl26/influscan-api/internal/middleware"
	"github.com/sgl26/influscan-api/internal/repository"
)

// ScanHandler handles scan-related HTTP requests
type ScanHandler struct {
	repo *repository.ScanRepository
}

// NewScanHandler creates a new scan handler
func NewScanHandler(repo *repository.ScanRepository) *ScanHandler {
	return &ScanHandler{repo: repo}
}

// GetScans handles the GET /scans endpoint
func (h *ScanHandler) GetScans(w http.ResponseWriter, r *http.Request) {
	// Ensure user is authenticated
	user, ok := middleware.RequireAuth(w, r)
	if !ok {
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Get scans from repository
	scans, err := h.repo.GetScans(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encode and send response
	if err := json.NewEncoder(w).Encode(scans); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
