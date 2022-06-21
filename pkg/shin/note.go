package shin

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Note struct {
	No       int
	Contents []string
}

func NewNote() Note {

	basePath := GetBasePath()

	paths, _ := filepath.Glob(filepath.Join(basePath, "*.shin"))

	var highest float64 = 0
	for _, path := range paths {
		rep := regexp.MustCompile(`.shin$`)
		filename := filepath.Base(rep.ReplaceAllString(path, ""))

		number, _ := strconv.Atoi(filename)
		highest = math.Max(highest, float64(number))

	}

	return Note{No: int(highest) + 1, Contents: []string{""}}
}

func (n *Note) UpdateLine(line string, row int) {
	if len(n.Contents)-1 < row {
		n.Contents = append(n.Contents, line)
	} else {
		n.Contents[row] = line
	}

}

func (n *Note) UpdateContents(contents []string) {
	n.Contents = contents

}

func (n Note) Write() {

	basePath := GetBasePath()

	number := strconv.Itoa(n.No)
	f, err := os.Create(filepath.Join(basePath, number+".shin"))
	if err != nil {
		fmt.Println(err)
	}

	var textData string = ""
	for _, row := range n.Contents {
		textData += row + "\n"
	}

	output := []byte(textData)
	_, err = f.Write(output)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
}

func ReadNo(no int) Note {

	basePath := GetBasePath()

	f, err := os.Open(filepath.Join(basePath, strconv.Itoa(no)+".shin"))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return Note{No: no, Contents: contents}

}

func DeleteNo(no int) {

	basePath := GetBasePath()

	err := os.Remove(filepath.Join(basePath, strconv.Itoa(no)+".shin"))
	if err != nil {
		fmt.Println(err)
	}
}

func GetBasePath() string {
	var basePath string

	if os.Getenv("SHIN_STORAGE") != "" {
		basePath = os.Getenv("SHIN_STORAGE")
	} else {
		basePath = "~/.shin/"
	}

	return basePath
}
