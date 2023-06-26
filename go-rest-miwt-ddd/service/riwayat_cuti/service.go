package riwayat_cuti

import "context"

type Service interface {
	GetAllRiwayatCuti(ctx context.Context) (out []MessageResponse, err error)
}
