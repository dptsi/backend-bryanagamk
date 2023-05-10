package service

import (
	"bryanagamk/go-rest-miwt/domain"
	"bryanagamk/go-rest-miwt/exception"
	"bryanagamk/go-rest-miwt/helper"
	"bryanagamk/go-rest-miwt/repository"
	"bryanagamk/go-rest-miwt/web"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type CutiServiceImpl struct {
	CutiRepository repository.CutiRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCutiService(cutiRepository repository.CutiRepository, DB *sql.DB, validate *validator.Validate) CutiService {
	return &CutiServiceImpl{CutiRepository: cutiRepository, DB: DB, Validate: validate}
}

func (service *CutiServiceImpl) Create(ctx context.Context, request web.CutiCreateRequest) web.CutiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	id, err := uuid.Parse("000311DB-A57E-4F2E-9394-8BAF3D38F863")
	helper.PanicIfError(err)
	cuti := domain.Cuti{
		IdRiwayatCuti: uuid.New(),
		IdSdm:         id,
		IdJenisCuti:   request.IdJenisCuti,
		TglSurat:      time.DateOnly,
		Lama:          request.Lama,
		Keterangan:    request.Keterangan,
		NoTelp:        request.NoTelp,
		TglAwalCuti:   request.TglAwalCuti,
		TglAkhirCuti:  request.TglAkhirCuti,
		Lokasi:        request.Lokasi,
	}

	service.CutiRepository.Save(ctx, tx, cuti)

	return helper.ToCutiResponse(cuti)
}

func (service *CutiServiceImpl) Update(ctx context.Context, request web.CutiUpdateRequest) web.CutiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cuti, errorFindData := service.CutiRepository.FindById(ctx, tx, request.IdRiwayatCuti)
	if errorFindData != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	id, err := uuid.Parse("000311DB-A57E-4F2E-9394-8BAF3D38F863")

	cuti.IdSdm = id
	cuti.IdJenisCuti = request.IdJenisCuti
	cuti.Lama = request.Lama
	cuti.Keterangan = request.Keterangan
	cuti.NoTelp = request.NoTelp
	cuti.TglAwalCuti = request.TglAwalCuti
	cuti.TglAkhirCuti = request.TglAkhirCuti
	cuti.Lokasi = request.Lokasi

	service.CutiRepository.Update(ctx, tx, cuti)

	return helper.ToCutiResponse(cuti)
}

func (service *CutiServiceImpl) Delete(ctx context.Context, cutiId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cuti, errorFindData := service.CutiRepository.FindById(ctx, tx, cutiId)
	if errorFindData != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CutiRepository.Delete(ctx, tx, cuti)
}

func (service *CutiServiceImpl) FindById(ctx context.Context, cutiId string) web.CutiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cuti, errorFindData := service.CutiRepository.FindById(ctx, tx, cutiId)
	if errorFindData != nil {
		panic(exception.NewNotFoundError(errorFindData.Error()))
	}

	return helper.ToCutiResponse(cuti)
}

func (service *CutiServiceImpl) FindAll(ctx context.Context) []web.CutiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cutiList := service.CutiRepository.FindAll(ctx, tx)

	return helper.ToCutiResponses(cutiList)
}
