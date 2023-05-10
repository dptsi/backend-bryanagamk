package domain

import "github.com/google/uuid"

type Cuti struct {
	IdRiwayatCuti   uuid.UUID
	IdSdm           uuid.UUID
	IdJenisCuti     int
	IsValid         bool
	TglSurat        string
	Lama            int
	Keterangan      string
	NoTelp          string
	TglAwalCuti     string
	TglAkhirCuti    string
	Lokasi          string
	SisaCutiTahunan int
}
