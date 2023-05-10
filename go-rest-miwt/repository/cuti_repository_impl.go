package repository

import (
	"bryanagamk/go-rest-miwt/domain"
	"bryanagamk/go-rest-miwt/exception"
	"bryanagamk/go-rest-miwt/helper"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type CutiRepositoryImpl struct {
}

func NewCutiRepository() CutiRepository {
	return &CutiRepositoryImpl{}
}

func (repository *CutiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cuti domain.Cuti) domain.Cuti {
	SQL := "insert into riwayat_cuti(id_riwayat_cuti, id_sdm, id_jenis_cuti, tgl_surat, lama, keterangan, no_telp, tanggal_awal_cuti, tanggal_akhir_cuti, lokasi, sisa_cuti_tahunan, updater) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL,
		cuti.IdRiwayatCuti,
		cuti.IdSdm,
		cuti.IdJenisCuti,
		cuti.TglSurat,
		cuti.Lama,
		cuti.Keterangan,
		cuti.NoTelp,
		cuti.TglAwalCuti,
		cuti.TglAkhirCuti,
		cuti.Lokasi,
		cuti.SisaCutiTahunan,
		uuid.New(),
	)

	helper.PanicIfError(err)

	return cuti
}

func (repository *CutiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cuti domain.Cuti) domain.Cuti {
	SQL := "update riwayat_cuti set id_sdm = ?, id_jenis_cuti = ?, tgl_surat = ?, lama = ?, keterangan = ?, no_telp = ?, tanggal_awal_cuti = ?, tanggal_akhir_cuti = ?, lokasi = ?, sisa_cuti_tahunan = ?, updater = ? where id_riwayat_cuti = ?"
	_, err := tx.ExecContext(ctx, SQL,
		cuti.IdSdm,
		cuti.IdJenisCuti,
		cuti.TglSurat,
		cuti.Lama,
		cuti.Keterangan,
		cuti.NoTelp,
		cuti.TglAwalCuti,
		cuti.TglAkhirCuti,
		cuti.Lokasi,
		cuti.SisaCutiTahunan,
		uuid.New(),
		cuti.IdRiwayatCuti,
	)
	helper.PanicIfError(err)

	return cuti
}

func (repository *CutiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, cuti domain.Cuti) {
	SQL := "delete riwayat_cuti where id_riwayat_cuti = ?"
	_, err := tx.ExecContext(ctx, SQL, cuti.IdRiwayatCuti)
	helper.PanicIfError(err)
}

func (repository *CutiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cutiId string) (domain.Cuti, error) {
	SQL := "select convert(nvarchar(36), id_riwayat_cuti) as id_riwayat_cuti, convert(nvarchar(36), id_sdm) as id_sdm, id_jenis_cuti, tgl_surat, lama, keterangan, no_telp, tanggal_awal_cuti, tanggal_akhir_cuti, lokasi, ISNULL(sisa_cuti_tahunan, 0) as sisa_cuti_tahunan from riwayat_cuti where id_riwayat_cuti = ?"
	rows, err := tx.QueryContext(ctx, SQL, cutiId)

	cuti := domain.Cuti{}
	if rows.Next() {
		errorRow := rows.Scan(&cuti.IdRiwayatCuti, &cuti.IdSdm, &cuti.IdJenisCuti, &cuti.TglSurat, &cuti.Lama, &cuti.Keterangan, &cuti.NoTelp, &cuti.TglAwalCuti, &cuti.TglAkhirCuti, &cuti.Lokasi, &cuti.SisaCutiTahunan)
		helper.PanicIfError(errorRow)
		return cuti, nil
	} else {
		panic(exception.NewNotFoundError(err.Error()))
	}
}

func (repository *CutiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cuti {
	SQL := "select convert(nvarchar(36), id_riwayat_cuti) as id_riwayat_cuti, convert(nvarchar(36), id_sdm) as id_sdm, id_jenis_cuti, tgl_surat, lama, keterangan, no_telp, tanggal_awal_cuti, tanggal_akhir_cuti, lokasi, ISNULL(sisa_cuti_tahunan, 0) as sisa_cuti_tahunan  from riwayat_cuti"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var cutiList []domain.Cuti
	for rows.Next() {
		cuti := domain.Cuti{}
		errorRow := rows.Scan(&cuti.IdRiwayatCuti, &cuti.IdSdm, &cuti.IdJenisCuti, &cuti.TglSurat, &cuti.Lama, &cuti.Keterangan, &cuti.NoTelp, &cuti.TglAwalCuti, &cuti.TglAkhirCuti, &cuti.Lokasi, &cuti.SisaCutiTahunan)
		helper.PanicIfError(errorRow)

		cutiList = append(cutiList, cuti)
	}

	return cutiList
}
