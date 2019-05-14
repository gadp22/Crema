package crema

import (
	"database/sql"
	"strconv"
	"strings"
)

type Query struct {
	QueryString string
}

const get = "SELECT "

const from = " FROM "
const innerJoin = " INNER JOIN "
const leftJoin = " LEFT JOIN "
const on = " ON "

const where = " WHERE "

const insert = "INSERT INTO "
const val = "VALUES"

const update = "UPDATE "
const set = " SET "

const delete = "DELETE "

const all = "*"
const and = " AND "

const equals = " = "
const notEquals = " != "

const returning = " RETURNING "

const greaterEqual = " >= "
const lessEqual = " <= "

const greater = " > "
const less = " < "

const order = " ORDER BY "
const asc = " ASC"
const desc = " DESC"

const limit = " LIMIT "

func SetDB(database *sql.DB) {
	db = database
}

func BeginTransaction() (*sql.Tx, error) {

	return db.Begin()
}

func ExecuteQuery(queryString string) (*sql.Rows, error) {
	rows, err := db.Query(queryString)

	return rows, err
}

func ExecuteNonQuery(queryString string) (sql.Result, error) {
	res, err := db.Exec(queryString)

	return res, err
}

func ExecuteQueryRow(tx *sql.Tx, queryString string) *sql.Row {
	row := tx.QueryRow(queryString)

	return row
}

func ExecuteNonQueryTransaction(tx *sql.Tx, queryString string) (sql.Result, error) {
	res, err := tx.Exec(queryString)

	return res, err
}

func Select(values ...string) *Query {
	var q Query

	q.QueryString = get

	for _, value := range values {
		q.QueryString += value
		q.QueryString += ", "
	}
	q.QueryString = strings.TrimRight(q.QueryString, ", ")

	return &q
}

func Insert(table string) *Query {
	var q Query

	q.QueryString = insert + table

	return &q
}

func Update(table string) *Query {
	var q Query

	q.QueryString = update + table

	return &q
}

func Delete() *Query {
	var q Query

	q.QueryString = delete

	return &q
}

func (q *Query) Set() *Query {
	q.QueryString += set

	return q
}

func (q *Query) Columns(columns []string) *Query {
	q.QueryString += "("

	for _, column := range columns {
		q.QueryString += column
		q.QueryString += ","
	}
	q.QueryString = strings.TrimRight(q.QueryString, ",")

	q.QueryString += ") "

	return q
}

func (q *Query) Values(values []string) *Query {
	q.QueryString += val

	q.QueryString += "("
	for _, value := range values {
		q.QueryString += "'"
		q.QueryString += value
		q.QueryString += "'"
		q.QueryString += ","
	}
	q.QueryString = strings.TrimRight(q.QueryString, ",")

	q.QueryString += ") "

	return q
}

func (q *Query) All() *Query {
	q.QueryString += all

	return q
}

func (q *Query) From(table string) *Query {
	q.QueryString += from
	q.QueryString += table

	return q
}

func (q *Query) InnerJoin(table string) *Query {
	q.QueryString += innerJoin
	q.QueryString += table

	return q
}

func (q *Query) On() *Query {
	q.QueryString += on

	return q
}

func (q *Query) Where() *Query {
	q.QueryString += where

	return q
}

func (q *Query) And() *Query {
	q.QueryString += and

	return q
}

func (q *Query) Equal(key string, val interface{}) *Query {

	q.QueryString += key
	q.QueryString += equals

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) GreaterEqual(key string, val interface{}) *Query {
	q.QueryString += key
	q.QueryString += greaterEqual

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) LessEqual(key string, val interface{}) *Query {
	q.QueryString += key
	q.QueryString += lessEqual

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) Greater(key string, val interface{}) *Query {
	q.QueryString += key
	q.QueryString += greater

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) Less(key string, val interface{}) *Query {
	q.QueryString += key
	q.QueryString += less

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) OrderBy(column string) *Query {
	q.QueryString += order
	q.QueryString += column

	return q
}

func (q *Query) Asc() *Query {
	q.QueryString += asc

	return q
}

func (q *Query) Desc() *Query {
	q.QueryString += desc

	return q
}

func (q *Query) Limit(val string) *Query {
	q.QueryString += limit
	q.QueryString += val

	return q
}

func (q *Query) EqualColumn(key string, val string) *Query {
	q.QueryString += key
	q.QueryString += equals
	q.QueryString += val

	return q
}

func (q *Query) EqualMD5(key string, val string) *Query {
	q.QueryString += key
	q.QueryString += equals
	q.QueryString += "md5('"
	q.QueryString += val
	q.QueryString += "')"

	return q
}

func (q *Query) NotEqual(key string, val interface{}) *Query {
	q.QueryString += key
	q.QueryString += notEquals

	if s := typeof(val); s == "string" {
		q.QueryString += "'"
		q.QueryString += val.(string)
		q.QueryString += "'"
	} else if s := typeof(val); s == "int" {
		q.QueryString += strconv.Itoa(val.(int))
	}

	return q
}

func (q *Query) NotEqualColumn(key string, val string) *Query {
	q.QueryString += key
	q.QueryString += notEquals
	q.QueryString += val

	return q
}

func (q *Query) Returning(columnName string) *Query {
	q.QueryString += returning
	q.QueryString += columnName

	return q
}

func GetGenericSelectQuery(table string, conditions map[string]string) *Query {
	q := Select().All().From(table)

	if conditions["id"] != "" {
		id, err := strconv.Atoi(conditions["id"])

		if err != nil {
			PrintfError(err.Error()) 
			panic(err)
		}

		q.Where().Equal("id", id)	
	}

	return q
}

func GetGenericInsertQuery(table string, values map[string]string) *Query {
	var keyList []string
	var valList []string

	for key, val := range values {
		keyList = append(keyList, key)
		valList = append(valList, val)
	}

	q := Insert(table).Columns(keyList)
	q.Values(valList).Returning("id")

	return q
}

func GetGenericUpdateQuery(table string, values map[string]string) *Query {
	q := Update(table).Set()

	for key, val := range values {
		if key != "id" {
			q.Equal(key, val)
			q.QueryString += " , "
		}
	}
	q.QueryString = strings.TrimRight(q.QueryString, " , ")

	id, err := strconv.Atoi(values["id"])

	if err != nil {
		panic(err)
	}

	q.Where().Equal("id", id)

	return q
}

func GetGenericDeleteQuery(table string, conditions map[string]string) *Query {
	id, err := strconv.Atoi(conditions["id"])

	if err != nil {
		panic(err)
	}

	q := Delete().From(table).Where().Equal("id", id)

	return q
}

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
