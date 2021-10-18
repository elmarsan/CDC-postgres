package test

import (
	listen "cdc"
	"testing"
)

func TestSetupInsertListenerForExistingTable(t *testing.T) {
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

func TestSetupInsertListenerForNonExistingTable(t *testing.T) {
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
