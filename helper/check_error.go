package helper

import "fmt"

func CheckError(err any) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
