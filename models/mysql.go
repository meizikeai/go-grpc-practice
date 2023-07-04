package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"go-grpc-practice/libs/tool"
	"go-grpc-practice/libs/types"

	log "github.com/sirupsen/logrus"
)

func AddUserForMySQL(param string) (int, error) {
	db := tool.GetMySQLClient("default.master")

	times := time.Now().Unix()
	query := "INSERT INTO `test` (`param`,`create_time`) VALUES (?,?)"

	rows, err := db.Exec(query, param, times)

	if err != nil {
		log.Error(fmt.Sprintf("%s %s %v %s", query, param, times, err.Error()))
		return 0, err
	}

	result := 0
	res, err := rows.RowsAffected()

	if err != nil {
		return 0, err
	}

	if res > 0 {
		result = 1
	}

	return result, err
}

func GetUserForMySQL(uid string, data []string) types.MapStringString {
	db := tool.GetMySQLClient("default.slave")

	result := types.MapStringString{}

	column := strings.Join(data, ",")
	query := fmt.Sprintf("SELECT %s FROM `test` WHERE `uid` = ?", column)

	rows, err := db.Query(query, uid)

	if err != nil {
		log.Error(fmt.Sprintf("%s %s %s", query, uid, err.Error()))
		return result
	}

	back := handleRowsToSlice(rows)

	if len(back) > 0 {
		result = back[0]
	}

	defer rows.Close()

	return result
}

func handleRowsToSlice(rows *sql.Rows) (result []types.MapStringString) {
	columns, _ := rows.Columns()
	count := len(columns)

	values := make([]sql.RawBytes, count)
	pointer := make([]interface{}, count)

	for i := 0; i < count; i++ {
		pointer[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(pointer...)

		row := make(types.MapStringString)

		for i := 0; i < count; i++ {
			row[columns[i]] = string(values[i])
		}

		result = append(result, row)
	}

	return
}
