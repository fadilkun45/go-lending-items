package helper

import (
	"errors"
	"regexp"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var fkColumnRegex = regexp.MustCompile(`FOREIGN KEY \(` + "`" + `(\w+)` + "`" + `\)`)

func HandleDBError(err error, notFoundMsg string) {
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(NotFound(notFoundMsg))
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg := "referenced id does not exist"
		if m := fkColumnRegex.FindStringSubmatch(mysqlErr.Message); len(m) == 2 {
			msg = m[1] + " does not exist"
		}
		panic(BadRequest(msg))
	}
	panic(err)
}
