package main

import (
	"fmt"
	"github.com/chenleijava/go-guava/converter-master"
)

func main() {
	t2t := converter.NewTable2Struct()

	err := t2t.
		SavePath("/home/go/project/model/model.go").
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		Run()
	fmt.Println(err)
}
