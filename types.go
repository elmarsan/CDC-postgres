package listen

type SqlEvent string

const (
	Insert SqlEvent = "Insert"
	Update SqlEvent = "Update"
	Delete SqlEvent = "Delete"
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
