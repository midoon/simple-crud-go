package main

import (
	"github.com/midoon/simple-crud-go/dependency"
	"github.com/midoon/simple-crud-go/helper"
)



func main() {
	server := dependency.InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
