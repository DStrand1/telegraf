package sql

import (
	// Blank imports to register the drivers
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/IBM/nzgo/v12"
	_ "github.com/SAP/go-hdb/driver"
	_ "github.com/apache/arrow/go/v13/arrow/flight/flightsql/driver"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/sijms/go-ora/v2"
)
