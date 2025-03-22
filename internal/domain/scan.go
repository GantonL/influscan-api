package domain

// Scan represents a scan record
type Scan struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}
