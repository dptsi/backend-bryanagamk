package helper

import (
	"bryanagamk/go-rest-miwt/domain"
	"bryanagamk/go-rest-miwt/web"
)

func ToCutiResponse(cuti domain.Cuti) web.CutiResponse {
	return web.CutiResponse{
		IdRiwayatCuti:   cuti.IdRiwayatCuti,
		IdSdm:           cuti.IdSdm,
		IdJenisCuti:     cuti.IdJenisCuti,
		TglSurat:        cuti.TglSurat,
		Lama:            cuti.Lama,
		Keterangan:      cuti.Keterangan,
		NoTelp:          cuti.NoTelp,
		TglAwalCuti:     cuti.TglAwalCuti,
		TglAkhirCuti:    cuti.TglAkhirCuti,
		Lokasi:          cuti.Lokasi,
		SisaCutiTahunan: cuti.SisaCutiTahunan,
	}
}

func ToCutiResponses(cutiList []domain.Cuti) []web.CutiResponse {
	var cutiResponses []web.CutiResponse
	for _, cuti := range cutiList {
		cutiResponses = append(cutiResponses, ToCutiResponse(cuti))
	}

	return cutiResponses
}
