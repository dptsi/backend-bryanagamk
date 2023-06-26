package riwayatcuti

import "time"

type RiwayatCutiEntity struct {
	Id        int64     `json:"id" db:"id"`
	Reason    string    `json:"string" db:"reason"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
