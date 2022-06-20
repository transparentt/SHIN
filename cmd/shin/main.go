package main

import "github.com/transparentt/SHIN/pkg/shin"

func main() {

	note := shin.NewNote()
	editor := shin.NewEditor(note, 100)

	editor.Run()

}
