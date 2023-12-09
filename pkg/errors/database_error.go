package errors

var (
	ErrNoneDatabaseCode              = "11010"
	ErrDatabaseOpenCode              = "11011"
	ErrSQLMapUnmarshalJSONCode       = "11012"
	ErrSQLMapUnmarshalTextCode       = "11013"
	ErrSQLMapMarshalValueCode        = "11014"
	ErrSQLMapUnmarshalScannedCode    = "11015"
	ErrSQLMapInvalidScanCode         = "11016"
	ErrClosingDatabaseConnectionCode = "11017"
	ErrNoneDatabase                  = New(ErrNoneDatabaseCode, Alert, []string{"No Database selected"}, []string{}, []string{"database name is empty"}, []string{"Input a name for the database"})
	ErrSQLMapInvalidScan             = New(ErrSQLMapInvalidScanCode, Alert, []string{"invalid data type: expected []byte"}, []string{}, []string{}, []string{})
)

func ErrDatabaseOpen(err error) error {
	return New(ErrDatabaseOpenCode, Alert, []string{"Unable to open database", err.Error()}, []string{err.Error()}, []string{"Database is unreachable"}, []string{"Make sure your database is reachable"})
}

func ErrSQLMapUnmarshalJSON(err error) error {
	return New(ErrSQLMapUnmarshalJSONCode, Alert, []string{"failed to unmarshal json", err.Error()}, []string{err.Error()}, []string{}, []string{})
}

func ErrSQLMapUnmarshalText(err error) error {
	return New(ErrSQLMapUnmarshalTextCode, Alert, []string{"failed to unmarshal text", err.Error()}, []string{err.Error()}, []string{}, []string{})
}

func ErrSQLMapMarshalValue(err error) error {
	return New(ErrSQLMapMarshalValueCode, Alert, []string{"failed to marshal value", err.Error()}, []string{err.Error()}, []string{}, []string{})
}

func ErrSQLMapUnmarshalScanned(err error) error {
	return New(ErrSQLMapUnmarshalScannedCode, Alert, []string{"failed to unmarshal scanned data", err.Error()}, []string{err.Error()}, []string{}, []string{})
}

func ErrClosingDatabaseConnection(err error) error {
	return New(ErrClosingDatabaseConnectionCode, Alert, []string{"failed to close database connection"}, []string{err.Error()}, []string{"Invalid database instance passed."}, []string{"Make sure the DB handler has a valid database instance."})
}
