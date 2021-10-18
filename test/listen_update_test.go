package test

import (
	listen "cdc"
	"testing"
)

func TestSetupUpdateListenerForExistingTable(t *testing.T) {
	table := "users"

	update := listen.Update{}

	listener, err := update.Listener(listen.Event{
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

func TestSetupUpdateListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	update := listen.Update{}

	listener, err := update.Listener(listen.Event{
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
