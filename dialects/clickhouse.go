package dialects

import (
	"fmt"
	"reflect"

	"gopkg.in/gorp.v1"
)

type ClickHouseDialect struct{}

func (d ClickHouseDialect) QuerySuffix() string { return "" }

func (d ClickHouseDialect) ToSqlType(val reflect.Type, maxsize int, isAutoIncr bool) string {
	switch val.Kind() {
	case reflect.Ptr:
		return d.ToSqlType(val.Elem(), maxsize, isAutoIncr)
	case reflect.Bool:
		return "UInt8"
	case reflect.Int:
		return "Int32"
	case reflect.Int8:
		return "Int8"
	case reflect.Int16:
		return "Int16"
	case reflect.Int32:
		return "Int32"
	case reflect.Uint8:
		return "UInt8"
	case reflect.Uint16:
		return "UInt16"
	case reflect.Uint32:
		return "UInt8"
	case reflect.Int64:
		return "Int64"
	case reflect.Uint64:
		return "UInt64"
	case reflect.Float64:
		return "Float64"
	case reflect.Float32:
		return "Float32"
	case reflect.Slice:
		if val.Elem().Kind() == reflect.Uint8 {
			return "Array(UInt8)"
		}
	}

	switch val.Name() {
	case "NullInt64":
		return "Int64"
	case "NullFloat64":
		return "Float64"
	case "NullBool":
		return "UInt8"
	case "NullTime", "Time":
		return "DateTime"
	}

	return "string"
}

// Returns empty string
func (d ClickHouseDialect) AutoIncrStr() string {
	return ""
}

func (d ClickHouseDialect) AutoIncrBindValue() string {
	return ""
}

func (d ClickHouseDialect) AutoIncrInsertSuffix(col *gorp.ColumnMap) string {
	return ""
}

// Returns suffix
func (d ClickHouseDialect) CreateTableSuffix() string {
	return ""
}

func (d ClickHouseDialect) TruncateClause() string {
	return ""
}

// Returns "$(i+1)"
func (d ClickHouseDialect) BindVar(i int) string {
	return fmt.Sprintf(":%d", i+1)
}

func (d ClickHouseDialect) InsertAutoIncr(exec gorp.SqlExecutor, insertSql string, params ...interface{}) (int64, error) {
	return 0, nil
}

func (d ClickHouseDialect) QuoteField(f string) string {
	return `"` + f + `"`
}

func (d ClickHouseDialect) QuotedTableForQuery(schema string, table string) string {
	return table
}

func (d ClickHouseDialect) IfSchemaNotExists(command, schema string) string {
	return fmt.Sprintf("%s if not exists", command)
}

func (d ClickHouseDialect) IfTableExists(command, schema, table string) string {
	return fmt.Sprintf("%s if exists", command)
}

func (d ClickHouseDialect) IfTableNotExists(command, schema, table string) string {
	return fmt.Sprintf("%s if not exists", command)
}
