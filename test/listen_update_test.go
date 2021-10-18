package test

import (
	listen "cdc"
	"testing"
)

func TestSetupUpdateListenerForExistingTable(t *testing.T) {
	table := "users"

	update := listen.Update{}

	_, err := update.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestSetupUpdateListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	update := listen.Update{}

	_, err := update.Listener(listen.Event{
		ConnParams: connParams,
		Event:      listen.InsertSQLEvent,
		Table:      table,
	})

	if err == nil {
		t.Errorf("Unexpected listener established for table %s", table)
	}
}
