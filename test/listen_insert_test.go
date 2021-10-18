package test

import (
	listen "cdc"
	"testing"
)

func TestSetupInsertListenerForExistingTable(t *testing.T) {
	table := "users"

	insert := listen.Insert{}

	listener, err := insert.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err != nil {
		t.Errorf("%v", err)
	}

	if listener != nil {
		defer listener.Close()
	}
}

func TestSetupInsertListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	insert := listen.Insert{}

	listener, err := insert.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err == nil {
		t.Errorf("Unexpected listener established for table %s", table)
	}

	if listener != nil {
		defer listener.Close()
	}
}
