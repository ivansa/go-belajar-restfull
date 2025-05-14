package main

import (
	"go-belajar-restfull/helper"
	"go-belajar-restfull/wire"
)

func main() {
	err := wire.InitializeServer().ListenAndServe()
	helper.CheckError(err)
}
