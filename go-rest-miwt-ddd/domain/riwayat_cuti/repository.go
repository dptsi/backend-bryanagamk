package riwayatcuti

import "context"

type RepositoryRiwayatCuti interface {
	GetAllRiwayatCuti(ctx context.Context) (out []RiwayatCutiEntity, err error)
	GetRiwayatCutiById(ctx context.Context, id int64) (out RiwayatCutiEntity, err error)
}
