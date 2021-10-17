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

	_, err := listen.Listener(listen.ListenEvent{
		ConnParams: connParams,
		Event:      listen.Insert,
		Table:      table,
	})

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestSetupListenerForNonExistingTable(t *testing.T) {
	table := "posts"

	_, err := listen.Listener(listen.ListenEvent{
		ConnParams: connParams,
		Event:      listen.Insert,
		Table:      table,
	})

	if err == nil {
		t.Errorf("Unexpected listener established for table %s", table)
	}
}
