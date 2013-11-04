package database

import (
	"errors"
	"expvar"
	"github.com/ziutek/mymysql/autorc"
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"reflect"
	"runtime"
	"strings"
)

type T struct{}

// prepared statements go here
var (
	Statements = make(map[string]*autorc.Stmt, 0)
)

func GetFunctionName(i interface{}) string {
	arr := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	if strings.Contains(arr, "database.") {
		strArr := strings.Split(arr, "database.")
		return strArr[1]
	}
	return arr
}

func PrepareAll() error {

	return nil
}

func GetStatement(key string) (stmt *autorc.Stmt, err error) {
	stmt, ok := Statements[key]
	if !ok {
		qry := expvar.Get(key)
		if qry == nil {
			err = errors.New("Invalid query reference")
		}
	}
	return
}
