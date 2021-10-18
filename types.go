package listen

type SqlEvent string

const (
	InsertSQLEvent SqlEvent = "Insert"
	UpdateSQLEvent SqlEvent = "Update"
	DeleteSQLEvent SqlEvent = "Delete"
)

type DBConnParams struct {
	Host string
	Port uint16
	User string
	Pass string
	Name string
}

type Event struct {
	ConnParams DBConnParams
	Table      string
	Event      SqlEvent
}
