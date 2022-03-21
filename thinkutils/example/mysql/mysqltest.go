package main

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"sync"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

/*
	sqlStatement := `
	UPDATE users
	SET first_name = $2, last_name = $3
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 5, "NewFirst", "NewLast")
	if err != nil {
	  panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
	  panic(err)
	}
	fmt.Println(count)
*/

type MyType struct {
	Id uint64 `json:"id" field:"id"`
	Name string `json:"name" field:"name"`
}

func basicQueryJSON(wg *sync.WaitGroup) {

	db := thinkutils.ThinkMysql.QuickConn()

	rows, err := db.Query(`
		SELECT 
       		* 
		FROM 
		    sys_user 
		WHERE 
		    user_id > ?`, 0)
	if err != nil {
		return
	}

	defer rows.Close()
	defer wg.Done()

	szRet := thinkutils.ThinkMysql.ToJSON(rows)
	fmt.Println(szRet)
	wg.Done()

}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go basicQueryJSON(&wg)
	}

	wg.Wait()
	//time.Sleep(10 * time.Second)
	//basicQueryJSON()
}
