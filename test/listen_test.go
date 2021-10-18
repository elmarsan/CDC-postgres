package test

import (
	listen "cdc"
	"testing"
)

var connParams = listen.DBConnParams{
	Host: "localhost",
	Port: 5432,
	User: "postgres",
	Pass: "password",
	Name: "test",
}

func TestSetupListenerForExistingTable(t *testing.T) {
	table := "users"

	insert := listen.Insert{}

	_, err := insert.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestSetupListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	insert := listen.Insert{}

	_, err := insert.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err == nil {
		t.Errorf("Unexpected listener established for table %s", table)
	}
}
