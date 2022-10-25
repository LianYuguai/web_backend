// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OerDao is the data access object for table oer.
type OerDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns OerColumns // columns contains all the column names of Table for convenient usage.
}

// OerColumns defines and stores column names for table oer.
type OerColumns struct {
	ToolChamberDate string //
	StableSlot      string //
	StableWafer     string //
	Count           string //
}

// oerColumns holds the columns for table oer.
var oerColumns = OerColumns{
	ToolChamberDate: "toolChamberDate",
	StableSlot:      "stable_slot",
	StableWafer:     "stable_wafer",
	Count:           "count",
}

// NewOerDao creates and returns a new DAO object for table data access.
func NewOerDao() *OerDao {
	return &OerDao{
		group:   "default",
		table:   "oer",
		columns: oerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OerDao) Columns() OerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}