package main

import (
	"fmt"
    "database/sql"
    
    _ "gopkg.in/goracle.v2"
)

func main() {
    db, err := sql.Open("goracle", connectString)
    
	if err != nil {
		fmt.Println(err)
		return
    }
	defer db.Close()

    rows, err := db.Query("select 2+2 from dual")
    if err != nil {
        fmt.Println("Error fetching addition")
        fmt.Println(err)
        return
    }
    defer rows.Close()
    
    for rows.Next() {
        var sum int
        rows.Scan(&sum)
        fmt.Printf("2 + 2 always equals: %d\n", sum)
    }
}
