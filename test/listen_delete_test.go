package test

import (
	listen "cdc"
	"testing"
)

func TestSetupDeleteListenerForExistingTable(t *testing.T) {
	table := "users"

	deleteL := listen.Delete{}

	listener, err := deleteL.Listener(listen.Event{
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

func TestSetupDeleteListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	deleteL := listen.Delete{}

	listener, err := deleteL.Listener(listen.Event{
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
