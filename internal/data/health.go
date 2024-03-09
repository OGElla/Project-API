package data 

import (
	"time"
	"database/sql"
	"github.com/OGElla/Project-API/internal/validator"
	"errors"
)

type Health struct {
	ID int64 `json:"id"` //unique integer ID
	CreatedAt time.Time `json:"-"` //timestamp 
	Walking Walking `json:"walking,omitempty"` //steps
	Hydrate Hydrate `json:"hydrate,omitempty"` //water
	Sleep Sleep `json:"sleep,omitempty"`//time
	Version int32 `json:"version"`//the version number
}
//reusable
func ValidateDaily(v *validator.Validator, movie *Health) {
	v.Check(movie.Walking > 0, "walking", "must be a positive integer")
}

type HealthModel struct{
	DB *sql.DB
}

func (m HealthModel) Insert(health *Health) error{
	query := `INSERT INTO healthtracker (walking, hydrate, sleep) VALUES($1, $2, $3) RETURNING id, created_at, version` 

	args := []interface{}{health.Walking, health.Hydrate, health.Sleep}

	return m.DB.QueryRow(query, args...).Scan(&health.ID, &health.CreatedAt, &health.Version)
}

func (m HealthModel) Get(id int64) (*Health, error) {
	if id < 1{
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, created_at, walking, hydrate, sleep, version FROM healthtracker WHERE id = $1`

	var health Health

	err := m.DB.QueryRow(query, id).Scan(
		&health.ID,
		&health.CreatedAt,
		&health.Walking,
		&health.Hydrate,
		&health.Sleep,
		&health.Version,
	)

	if err != nil{
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &health, nil
}

func (m HealthModel) Update(health *Health) error {
	query := `UPDATE healthtracker SET walking = $1, hydrate = $2, sleep =$3, version = version + 1 WHERE id = $4 RETURNING version`

	args := []interface{}{
		health.Walking,
		health.Hydrate, 
		health.Sleep,
		health.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&health.Version)
}

func (m HealthModel) Delete(id int64) error {
	if id < 1{
		return ErrRecordNotFound
	}

	query := `DELETE FROM healthtracker WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil{
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}