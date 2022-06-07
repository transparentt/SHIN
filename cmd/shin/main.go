package main

import (
	"github.com/transparentt/SHIN/pkg/shin"
)

func main() {

	note := shin.NewNote()

	for i := range note.Contents {
		note.Update("aiueo\n", i)
	}

	note.Write()

	// shin.DeleteNo(1)

}
