package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"time"
)

type SqlEvent string

const (
	Insert SqlEvent = "Insert"
	Update SqlEvent = "Update"
	Delete SqlEvent = "Delete"
)

type DBConnParams struct {
	Host string
	Port uint16
	User string
	Pass string
	Name string
}

type ListenEvent struct {
	ConnParams DBConnParams
	Table      string
	Event      SqlEvent
}

func GetListener(event ListenEvent) (*pq.Listener, error) {
	db := connect(event.ConnParams)
	return setupListener(db, event)
}

func connect(connParams DBConnParams) *sql.DB {
	connInfo := connInfo(connParams)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func setupListener(db *sql.DB, event ListenEvent) (*pq.Listener, error) {
	_, err := db.Query(`
		CREATE OR REPLACE FUNCTION cdc_notify_event() RETURNS TRIGGER AS $$

		DECLARE 
			data json;
		    notification json;
		
		BEGIN 
			IF (TG_OP = 'DELETE') THEN
		 		data = row_to_json(OLD);
			ELSE
		 		data = row_to_json(NEW);
			END IF;
		 
			notification = json_build_object(
				'timestamp', NOW(),
				'table',TG_TABLE_NAME, 
				'action', TG_OP,
				'data', data);

			PERFORM pg_notify('events',notification::text);
			RETURN NULL;
		END;

		$$ LANGUAGE plpgsql;
	`)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = db.Query(
		fmt.Sprintf(`
			DROP TRIGGER IF EXISTS %s_notify_event ON %s;
			CREATE TRIGGER %s_notify_event AFTER %s ON %s FOR EACH ROW EXECUTE PROCEDURE cdc_notify_event();
	`, event.Table, event.Table, event.Table, event.Event, event.Table,
		),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: table `%s` does not exist", event.Table))
	}

	_, err = db.Query("LISTEN EVENTS")
	if err != nil {
		return nil, err
	}

	return pq.NewListener(connInfo(event.ConnParams), time.Second, time.Second, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}), nil
}

func connInfo(connParams DBConnParams) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connParams.Host, connParams.Port, connParams.User, connParams.Pass, connParams.Name)
}
