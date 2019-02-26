package driver

import (
	"strings"

	"github.com/QHasaki/Server/module/v1"
)

// GetQuerySQL returns query sql
// the input arg `document` : cannot be nil
// the input arg `column` : at least include one string, "*" means query all
// the input arg `where` : can be nil
// check above before use func `Get` & `GetOne`
func GetQuerySQL(document string, where module.Data) (string, []interface{}, error) {
	// checkout the input args
	// if document == "" {
	// 	return "", nil, ErrMissDocument
	// }

	// if len(column) == 0 {
	// 	return "", nil, ErrMissColumn
	// }

	var (
		columns string
		wheres  []string
		values  []interface{}
	)
	sql := "SELECT * "

	for k, v := range where {
		switch v.(type) {
		case module.Where:
			wheres = append(wheres, k+" "+v.(module.Where).Operator+" ?")
			values = append(values, v.(module.Where).Value)
		default:
			wheres = append(wheres, k+" = ?")
			values = append(values, v)
		}
	}

	sql += columns + " FROM `" + document + "`"
	if len(wheres) != 0 {
		sql += " WHERE " + strings.Join(wheres, " AND ")
	}

	return sql, values, nil
}

// GetExecSQL returns exec sql
// the input arg `document` : cannot be nil
// the input arg `data` : cannot be nil
// the input arg `where` : cannot be nil
// check above before use func `Get` & `GetOne`
func GetExecSQL(document string, data module.Data, where module.Data) (string, []interface{}, error) {
	// checkout the input args
	// if document == "" {
	// 	return "", nil, ErrMissDocument
	// }

	// if len(data) == 0 {
	// 	return "", nil, ErrMissData
	// }

	// where == nil : create a new record
	if where == nil {
		var (
			column []string
			needed []string
			values []interface{}
		)

		sql := "INSERT INTO `" + document + "` "

		for k, v := range data {
			column = append(column, k)
			needed = append(needed, "?")
			values = append(values, v)
		}

		return sql + "(" + strings.Join(column, ",") + ") VALUES (" + strings.Join(needed, ", ") + ")", values, nil
	}

	// where != nil : update old record
	var (
		datas  []string
		wheres []string
		values []interface{}
	)

	sql := "UPDATE `" + document + "` SET "

	for k, v := range data {
		datas = append(datas, k+" = ?")
		values = append(values, v)
	}

	for k, v := range where {
		switch v.(type) {
		case module.Where:
			wheres = append(wheres, k+" "+v.(module.Where).Operator+" ?")
			values = append(values, v.(module.Where).Value)
		default:
			wheres = append(wheres, k+" = ?")
			values = append(values, v)
		}
	}

	sql += strings.Join(datas, ", ")

	if len(wheres) != 0 {
		sql += " WHERE " + strings.Join(wheres, " AND ")
	}

	return sql, values, nil

}
