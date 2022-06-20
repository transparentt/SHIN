package main

import "github.com/transparentt/SHIN/pkg/shin"

func main() {

	note := shin.ReadNo(1)
	editor := shin.NewEditor(note, 100)

	editor.Run()

}
