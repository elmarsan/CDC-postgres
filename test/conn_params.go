package test

import listen "cdc"

var connParams = listen.DBConnParams{
	Host: "localhost",
	Port: 45432,
	User: "postgres",
	Pass: "ebitlabs",
	Name: "test",
}
