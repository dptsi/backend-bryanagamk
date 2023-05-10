package web

type CutiCreateRequest struct {
	IdJenisCuti  int    `validate:"required" json:"id_jenis_cuti"`
	Lama         int    `validate:"required" json:"lama"`
	Keterangan   string `json:"keterangan"`
	NoTelp       string `json:"no_telp"`
	TglAwalCuti  string `validate:"required" json:"tgl_awal_cuti"`
	TglAkhirCuti string `validate:"required" json:"tgl_akhir_cuti"`
	Lokasi       string `validate:"required" json:"lokasi"`
}
