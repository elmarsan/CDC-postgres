package listen

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"time"
)

type UpdateColumn struct{ Listen }

func (update UpdateColumn) Listener(event Event, column string) (*pq.Listener, error) {
	db := connect(event.ConnParams)
	err := createNotifyEvent(db)
	if err != nil {
		return nil, err
	}

	_, err = db.Query(
		fmt.Sprintf(`
			DROP TRIGGER IF EXISTS %s_notify_event ON %s;
			CREATE TRIGGER %s_notify_event AFTER UPDATE OF %s ON %s FOR EACH ROW EXECUTE PROCEDURE cdc_notify_event();
	`, event.Table, event.Table, event.Table, column, event.Table,
		),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: column `%s` does not exist on %s", column, event.Table))
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
