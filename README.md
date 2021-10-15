## CDC Postgres

Simple library to listen `INSERT`, `UPDATE` and `DELETE` sql statements executed into specific table. <br/>
Only valid with postgres databases.

### How to use it

```go
func main() {
    listenEvent := ListenEvent{
        Table: "products",
        Event: Insert,
        ConnParams: DBConnParams{
            Host: "localhost",
            Port: 45432,
            User: "postgres",
            Pass: "ebitlabs",
            Name: "shop",
        },
    }
    
    listener, err := GetListener(listenEvent)
    if err != nil {
        log.Fatal(err)
    }
    
    err = listener.Listen("events")
    if err != nil {
        log.Fatal(err)
    }
}

func waitForNotification(l *pq.Listener) {
    for {
        select {
            case n := <-l.Notify:
        	fmt.Println("Received data from channel [", n.Channel, "] :")
        	
        	parsedData := bytes.Buffer{}
        	err := json.Indent(&parsedData, []byte(n.Extra), "", "\t")
        	if err != nil {fmt.Println("Error processing JSON: ", err)
        	    return
        	}
        	
        	fmt.Println(string(parsedData.Bytes()))
        	return
        }
    }
}


```