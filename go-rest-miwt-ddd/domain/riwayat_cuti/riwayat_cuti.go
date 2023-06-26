package riwayatcuti

import (
	"context"
	"strings"
)

type RiwayatCutiRepository struct {
	masterDB safesql.MasterDB
}

const tblRiwayatCuti = "riwayat_cuti"

func NewRiwayatCutiRepository(masterDB safesql.MasterDB) *RiwayatCutiRepository {
	if masterDB == nil {
		panic("Database is nil")
	}

	return &RiwayatCutiRepository{
		masterDB: masterDB,
	}
}

func (cuti *RiwayatCutiRepository) GetAllRiwayatCuti(ctx context.Context) (out []RiwayatCutiEntity, err error) {
	query := strings.Builder{}
	query.WriteString("SELECT * FROM ")
	query.WriteString(tblRiwayatCuti)

	stmt, err := cuti.masterDB.PreparexContext(ctx, query.String())
	if err != nil {
		return out, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		d := RiwayatCutiEntity{}
		err = rows.Scan(&d.Id)

		if err != nil {
			return out, err
		}

		out = append(out, d)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return

}
