package pmsmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	pmsProductVerifyRecordFieldNames          = builderx.RawFieldNames(&PmsProductVerifyRecord{})
	pmsProductVerifyRecordRows                = strings.Join(pmsProductVerifyRecordFieldNames, ",")
	pmsProductVerifyRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductVerifyRecordFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductVerifyRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductVerifyRecordFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductVerifyRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductVerifyRecord struct {
		Id         int64     `db:"id"`
		ProductId  int64     `db:"product_id"`
		CreateTime time.Time `db:"create_time"`
		VerifyMan  string    `db:"verify_man"` // 审核人
		Status     int64     `db:"status"`
		Detail     string    `db:"detail"` // 反馈详情
	}
)

func NewPmsProductVerifyRecordModel(conn sqlx.SqlConn) *PmsProductVerifyRecordModel {
	return &PmsProductVerifyRecordModel{
		conn:  conn,
		table: "pms_product_verify_record",
	}
}

func (m *PmsProductVerifyRecordModel) Insert(data PmsProductVerifyRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, pmsProductVerifyRecordRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductId, data.VerifyMan, data.Status, data.Detail)
	return ret, err
}

func (m *PmsProductVerifyRecordModel) FindOne(id int64) (*PmsProductVerifyRecord, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductVerifyRecordRows, m.table)
	var resp PmsProductVerifyRecord
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

func (m *PmsProductVerifyRecordModel) FindAll(Current int64, PageSize int64) (*[]PmsProductVerifyRecord, error) {

	if Current < 1 {
		Current = 1
	}
	if PageSize < 1 {
		PageSize = 20
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductVerifyRecordRows, m.table)
	var resp []PmsProductVerifyRecord
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

func (m *PmsProductVerifyRecordModel) Update(data PmsProductVerifyRecord) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductVerifyRecordRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductId, data.VerifyMan, data.Status, data.Detail, data.Id)
	return err
}

func (m *PmsProductVerifyRecordModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
