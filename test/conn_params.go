package test

import listen "cdc"

var connParams = listen.DBConnParams{
	Host: "localhost",
	Port: 5432,
	User: "postgres",
	Pass: "password",
	Name: "test",
}
