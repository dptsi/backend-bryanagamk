package grpc

import (
	"github.com/dptsi-bryanagamk/go-rest-miwt-ddd/container"
	riwayatcuti "github.com/dptsi-bryanagamk/go-rest-miwt-ddd/domain/riwayat_cuti"
)

type Handler struct {
	riwayatCutiHandler riwayatCutiPb.RiwayatCutiServiceServer
}

func SetupHandler(co *container.Container) *Handler {
	riwayatCutiHandler := riwayatcuti.SetupRiwayatCutiHandler(co.RiwayatCutiSvc)

	return &Handler{
		riwayatCutiHandler: riwayatCutiHandler,
	}
}
