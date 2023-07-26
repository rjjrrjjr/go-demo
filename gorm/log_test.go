package gorm

import (
	"context"
	"database/sql"
	"encoding/json"
	"git.code.oa.com/phoenixs/trpc-log-meet/log"
	"git.code.oa.com/trpc-go/trpc-database/gorm"
	"testing"
)

func TestPrintGormLog(t *testing.T) {
	rpcLogger := log.Logger(context.Background()).WithFields("Type", "Rpc", "Callee", "asdfg", "Caller",
		"asdfgh", "Remote", "asdfghj",
		"Cost", "12345")
	res := &gorm.Request{
		Op: gorm.OpQueryContext,
	}
	req := &gorm.Response{
		Rows: &sql.Rows{
			//closed: true,
		},
	}
	rpcJson, _ := json.Marshal(&RpcInfo{ResInfo: res, ReqInfo: req})
	rpcLogger.Infof("%s", rpcJson)
}

// RpcInfo RPC 请求信息
type RpcInfo struct {
	ReqInfo interface{}
	ResInfo interface{}
}