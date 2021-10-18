package test

import (
	listen "cdc"
	"testing"
)

func TestSetupUpdateColumnListenerForExistingColumn(t *testing.T) {
	table := "users"
	column := "name"

	update := listen.UpdateColumn{}

	listener, err := update.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	}, column)

	if err != nil {
		t.Errorf("%v", err)
	}

	if listener != nil {
		defer listener.Close()
	}
}

func TestSetupUpdateColumnListenerForNonExistingTable(t *testing.T) {
	table := "posts"
	column := "name"

	update := listen.UpdateColumn{}

	listener, err := update.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	}, column)

	if err == nil {
		t.Errorf("Unexpected listener established for column %s of table %s", column, table)
	}

	if listener != nil {
		defer listener.Close()
	}
}

func TestSetupUpdateColumnListenerForNonExistingColumn(t *testing.T) {
	table := "users"
	column := "lastname"

	update := listen.UpdateColumn{}

	listener, err := update.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	}, column)

	if err == nil {
		t.Errorf("Unexpected listener established for column %s of table %s", column, table)
	}

	if listener != nil {
		defer listener.Close()
	}
}
