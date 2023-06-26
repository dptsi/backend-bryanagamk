package container

import (
	"fmt"
	domainRiwayatCuti "github.com/dptsi-bryanagamk/go-rest-miwt-ddd/domain/riwayat_cuti"
	"github.com/dptsi-bryanagamk/go-rest-miwt-ddd/service/riwayat_cuti"
)

type Container struct {
	Config         *config.Config
	MasterDB       safesql.MasterDB
	RiwayatCutiSvc riwayat_cuti.Service
}

func Setup() *Container {
	err := config.Init{
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	}

	if err != nil {
		panic(fmt.Sprintf("failed in initializer config: %v", err))
	}

	cfg := config.Get()

	masterDB, err := safesql.OpenMasterDB("postgres", cfg.Postgres.Master.Address, "postgresql-miwt")

	mapsHttpClient := httpclient.NewClient()

	riwayatCutiRepo := domainRiwayatCuti.NewRiwayatCutiRepository(masterDB)

}
