package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/alexflint/go-arg"
	"github.com/transparentt/SHIN/pkg/shin"
)

func main() {

	var args struct {
		NoteNumber string `arg:"positional" help:"read the numbered note"`
		L          bool   `arg:"-l" help:"lists all notes and their first non-empty line of text"`
		C          string `arg:"-c" help:"prints the numbered note to standard output"`
		D          string `arg:"-d" help:"deletes the numbered note"`
	}

	arg.MustParse(&args)


	if args.NoteNumber != "" {
		no, err := strconv.Atoi(args.NoteNumber)
		if err != nil {
			fmt.Println("error: please input number!")
			os.Exit(1)
		}
		note := shin.ReadNo(no)
		editor := shin.NewEditor(note, 100)
		editor.Run()

	} else {
		if args.L {

			basePath := shin.GetBasePath()
			paths, _ := filepath.Glob(basePath + "*.shin")
			sort.Strings(paths)

			for _, path := range paths {

				rep := regexp.MustCompile(`.shin$`)
				filename := filepath.Base(rep.ReplaceAllString(path, ""))
				no, _ := strconv.Atoi(filename)
				note := shin.ReadNo(no)

				for i, line := range note.Contents {
					if line != "" {
						fmt.Println(filename+".shin:",  note.Contents[i])
						break
					}
				}

			}
			//
		} else if args.C != "" {

			no, err := strconv.Atoi(args.C)
			if err != nil {
				fmt.Println("error: please input number!")
				os.Exit(1)
			}

			note := shin.ReadNo(no)
			for _, content := range note.Contents {
				fmt.Println(content)
			}

		} else if args.D != "" {

			no, err := strconv.Atoi(args.D)
			if err != nil {
				fmt.Println("error: please input number!")
				os.Exit(1)
			}

			shin.DeleteNo(no)

		} else {

			note := shin.NewNote()
			editor := shin.NewEditor(note, 100)
			editor.Run()

		}

	}

}
