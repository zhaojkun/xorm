package xorm

import (
	"reflect"

	core "github.com/zhaojkun/xorm-core"
)

var (
	ptrPkType = reflect.TypeOf(&core.PK{})
	pkType    = reflect.TypeOf(core.PK{})
)
