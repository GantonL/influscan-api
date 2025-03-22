package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sgl26/influscan-api/internal/database"
	"github.com/sgl26/influscan-api/internal/domain"
)

// ScanRepository handles database operations for scans
type ScanRepository struct {
	db *database.SupabaseClient
}

// NewScanRepository creates a new scan repository
func NewScanRepository(db *database.SupabaseClient) *ScanRepository {
	return &ScanRepository{db: db}
}

// GetScans retrieves scans for a specific user
func (r *ScanRepository) GetScans(ctx context.Context, userID string) ([]domain.Scan, error) {
	query := map[string]interface{}{
		"user_id": fmt.Sprintf("eq.%s", userID),
		"order":   "created_at.desc",
	}

	data, err := r.db.Query(ctx, "scans", query)
	if err != nil {
		return nil, fmt.Errorf("error querying scans: %w", err)
	}

	var scans []domain.Scan
	if err := json.Unmarshal(data, &scans); err != nil {
		return nil, fmt.Errorf("error unmarshaling scans: %w", err)
	}

	return scans, nil
}

// CreateScan creates a new scan
func (r *ScanRepository) CreateScan(ctx context.Context, userID string) (*domain.Scan, error) {
	scan := domain.Scan{
		UserID:    userID,
		Status:    "pending",
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(scan)
	if err != nil {
		return nil, fmt.Errorf("error marshaling scan: %w", err)
	}

	result, err := r.db.Insert(ctx, "scans", data)
	if err != nil {
		return nil, fmt.Errorf("error inserting scan: %w", err)
	}

	var createdScan domain.Scan
	if err := json.Unmarshal(result, &createdScan); err != nil {
		return nil, fmt.Errorf("error unmarshaling created scan: %w", err)
	}

	return &createdScan, nil
}
