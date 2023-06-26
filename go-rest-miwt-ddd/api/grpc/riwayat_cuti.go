package grpc

type riwayatCutiGRPC struct {
	riwayatCutiPB.UnimplementedRiwayatCutiServiceServer
	svc riwayatCutiSvc.Service
}

func SetupRiwayatCutiHandler(svc riwayatCuti.Service) *riwayatCutiGRPC {
	if svc == nil {
		panic("please provice riwayat cuti service")
	}

	h := &riwayatCutiGRPC{
		svc: svc,
	}

	return h
}
