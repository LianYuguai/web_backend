// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Log is the golang structure for table log.
type Log struct {
	Id         uint        `json:"id"         description:"Log ID"`
	Passport   string      `json:"passport"   description:"Log Passport"`
	PassportId int         `json:"passportId" description:"Log Passport ID"`
	Time       *gtime.Time `json:"time"       description:"Log Time"`
}
