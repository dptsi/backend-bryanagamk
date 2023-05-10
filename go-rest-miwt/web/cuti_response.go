package web

import "github.com/google/uuid"

type CutiResponse struct {
	IdRiwayatCuti   uuid.UUID `json:"id_riwayat_cuti"`
	IdSdm           uuid.UUID `json:"id_sdm"`
	IdJenisCuti     int       `json:"id_jenis_cuti"`
	IsValid         bool      `json:"is_valid"`
	TglSurat        string    `json:"tgl_surat"`
	Lama            int       `json:"lama"`
	Keterangan      string    `json:"keterangan"`
	NoTelp          string    `json:"no_telp"`
	TglAwalCuti     string    `json:"tgl_awal_cuti"`
	TglAkhirCuti    string    `json:"tgl_akhir_cuti"`
	Lokasi          string    `json:"lokasi"`
	SisaCutiTahunan int       `json:"sisa_cuti_tahunan"`
}
