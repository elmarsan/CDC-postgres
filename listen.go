package listen

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Listen interface {
	Listener(event Event) (*pq.Listener, error)
}

func connect(connParams DBConnParams) *sql.DB {
	connInfo := connInfo(connParams)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func createNotifyEvent(db *sql.DB) error {
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

	return err
}

func connInfo(connParams DBConnParams) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connParams.Host, connParams.Port, connParams.User, connParams.Pass, connParams.Name)
}
