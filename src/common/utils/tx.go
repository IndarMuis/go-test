package utils

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		fmt.Println("ROLLBACK")
		tx.Rollback()
		//if errorRollback != nil {
		//panic(errorRollback)
		//}
	} else {
		fmt.Println("COMMIT")
		tx.Commit()
		//if errorCommit != nil {
		//panic(errorCommit)
	}
}
