package riwayat_cuti

import (
	"context"
	riwayatcutiRepo "github.com/dptsi-bryanagamk/go-rest-miwt-ddd/domain/riwayat_cuti"
)

type service struct {
	riwayatcutiRepo riwayatcutiRepo.RepositoryRiwayatCuti
}

func NewService(riwayatcutiRepo riwayatcutiRepo.RepositoryRiwayatCuti) *service {
	r := service{
		riwayatcutiRepo: riwayatcutiRepo,
	}

	return &r
}

func (s *service) GetAllRiwayatCuti(ctx context.Context) (out []MessageResponse, err error) {
	data, err := s.riwayatcutiRepo.GetAllRiwayatCuti(ctx)
	if err != nil {
		return out, err
	}

	if len(data) > 0 {
		out = append(out, MessageResponse{})
	}

	return out, nil
}
