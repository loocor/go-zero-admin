package pmsmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	pmsFreightTemplateFieldNames          = builderx.RawFieldNames(&PmsFreightTemplate{})
	pmsFreightTemplateRows                = strings.Join(pmsFreightTemplateFieldNames, ",")
	pmsFreightTemplateRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsFreightTemplateFieldNames, "id", "create_time", "update_time"), ",")
	pmsFreightTemplateRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsFreightTemplateFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsFreightTemplateModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsFreightTemplate struct {
		Id             int64   `db:"id"`
		Name           string  `db:"name"`
		ChargeType     int64   `db:"charge_type"`  // 计费类型:0->按重量; 1->按件数
		FirstWeight    float64 `db:"first_weight"` // 首重kg
		FirstFee       float64 `db:"first_fee"`    // 首费 (元)
		ContinueWeight float64 `db:"continue_weight"`
		ContinueFee    float64 `db:"continue_fee"`
		Dest           string  `db:"dest"` // 目的地 (省、市)
	}
)

func NewPmsFreightTemplateModel(conn sqlx.SqlConn) *PmsFreightTemplateModel {
	return &PmsFreightTemplateModel{
		conn:  conn,
		table: "pms_freight_template",
	}
}

func (m *PmsFreightTemplateModel) Insert(data PmsFreightTemplate) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, pmsFreightTemplateRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.ChargeType, data.FirstWeight, data.FirstFee, data.ContinueWeight, data.ContinueFee, data.Dest)
	return ret, err
}

func (m *PmsFreightTemplateModel) FindOne(id int64) (*PmsFreightTemplate, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsFreightTemplateRows, m.table)
	var resp PmsFreightTemplate
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *PmsFreightTemplateModel) FindAll(Current int64, PageSize int64) (*[]PmsFreightTemplate, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", pmsFreightTemplateRows, m.table)
	var resp []PmsFreightTemplate
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *PmsFreightTemplateModel) Update(data PmsFreightTemplate) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsFreightTemplateRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.ChargeType, data.FirstWeight, data.FirstFee, data.ContinueWeight, data.ContinueFee, data.Dest, data.Id)
	return err
}

func (m *PmsFreightTemplateModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
