package shin

type Note struct {
	No       int
	Contents [][]string
}

func (n Note) NewNote() Note {
	var path string
	var highest int
	var contents [][]string

	if env.SHIN_STORAGE != "" {
		path = env.SHIN_STORAGE
	} else {
		path = "~/.ssh/"
	}

	highest = n.getHighest(path) // If there is no note, return 0

	return Note{highest + 1, contents}
}

func (n *Note) Read(no string) {

	n.No = n.getNumber(no)
	n.Contents = n.getContents(no)

}

func (n *Note) Update(c Note.contents) {
	n.Contents = c
}

func (n Note) Write() err {
	if env.SHIN_STORAGE != nil {
		_, err := write(n.Contents, env.SHIN_STORAGE+n.No+".shin")
		return err
	} else {
		_, err := write(n.Contents, "~/.shin/"+n.No+".shin")
		return err
	}
}

func (n Note) Delete(no path) {
	_, err := deleteNote(no)
	return err
}

func (n Note) getNumber(path string) int {
	no := 3
	return no
}

func (n Note) getContents(path string) [][]stringintconv {
	var contents [][]string
	return contents
}

func (n Note) getHighest(path string) int {
	highest := 1
	return highest
}
