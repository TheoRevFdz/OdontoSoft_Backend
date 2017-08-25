package models

import "time"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	ID         int64  `gorm:"type: serial; primary_key", json:"id"`
	Work       string `gorm:"type: varchar(100), json:"work"`
	Quantity   int
	Cost       float64 `gorm:"type:decimal(36,2)"`
	CodPatient int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Treatments slice de tratamientos
type Treatments []Treatment
