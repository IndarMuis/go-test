package exception

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
		panic(err)
	}
}
