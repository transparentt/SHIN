package shin

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	note             Note
	characterPerLine int
	cursorX          int
	cursorY          int
}

func NewEditor(note Note, characterPerLine int) Editor {
	return Editor{
		note:             note,
		characterPerLine: characterPerLine,
		cursorX:          0,
		cursorY:          0,
	}
}

func (e *Editor) keyUp() {
	if e.cursorY-1 < 0 {
		e.cursorY = 0
	} else {
		e.cursorY = e.cursorY - 1
		if e.cursorX > len(e.note.Contents[e.cursorY]) {
			e.cursorX = len(e.note.Contents[e.cursorY])
		}
	}
}

func (e *Editor) keyDown() {
	if e.cursorY+1 >= len(e.note.Contents) {
		e.cursorY = len(e.note.Contents) - 1
	} else {
		e.cursorY = e.cursorY + 1
		if e.cursorX > len(e.note.Contents[e.cursorY]) {
			e.cursorX = len(e.note.Contents[e.cursorY])
		}
	}
}

func (e *Editor) keyLeft() {
	if e.cursorX-1 < 0 {
		e.cursorX = 0
	} else {
		e.cursorX = e.cursorX - 1
	}
}

func (e *Editor) keyRight() {
	if e.cursorX+1 >= len(e.note.Contents[e.cursorY]) {
		e.cursorX = len(e.note.Contents[e.cursorY])
	} else {
		e.cursorX = e.cursorX + 1
	}
}

func (e *Editor) keyEnter() {

	currentRightAll := e.note.Contents[e.cursorY][e.cursorX:]
	currentLeftAll := e.note.Contents[e.cursorY][:e.cursorX]
	e.note.UpdateLine(currentLeftAll, e.cursorY)

	newContents := make([]string, len(e.note.Contents[:e.cursorY+1]))
	copy(newContents, e.note.Contents[:e.cursorY+1])
	newContents = append(newContents, currentRightAll)
	newContents = append(newContents, e.note.Contents[e.cursorY+1:]...)
	e.note.UpdateContents(newContents)

	e.cursorY = e.cursorY + 1
	e.cursorX = 0
}

func (e *Editor) keyBackspace() {
	if e.cursorX == 0 && e.cursorY == 0 {
		return
	}

	currentLine := e.note.Contents[e.cursorY]

	if len(currentLine) != 0 {
		if e.cursorX != 0 {

			if e.cursorX-1 < 0 {
				e.cursorX = 0
			} else {
				e.cursorX = e.cursorX - 1
			}

			newLine := currentLine[:e.cursorX] + currentLine[e.cursorX+1:]
			e.note.UpdateLine(newLine, e.cursorY)

		} else {

			if e.cursorY-1 < 0 {
				e.cursorY = 0
			} else {
				e.cursorY = e.cursorY - 1
				e.cursorX = len(e.note.Contents[e.cursorY])
			}

			upperNewLine := e.note.Contents[e.cursorY] + currentLine
			e.note.UpdateLine(upperNewLine, e.cursorY)

			newContents := append(e.note.Contents[:e.cursorY+1], e.note.Contents[e.cursorY+1+1:]...)
			e.note.UpdateContents(newContents)

		}
	} else {
		if e.cursorX == 0 {

			if e.cursorY-1 < 0 {
				e.cursorY = 0
			} else {
				e.cursorY = e.cursorY - 1
				e.cursorX = len(e.note.Contents[e.cursorY])
			}

			newContents := append(e.note.Contents[:e.cursorY+1], e.note.Contents[e.cursorY+1+1:]...)
			e.note.UpdateContents(newContents)

		}
	}
}

func (e *Editor) keyDel() {
	currentLine := e.note.Contents[e.cursorY]
	if len(currentLine) > e.cursorX+1 {
		newLine := currentLine[:e.cursorX+1] + currentLine[e.cursorX+2:]
		e.note.UpdateLine(newLine, e.cursorY)
	}
}

func (e *Editor) keyCtrlS(s tcell.Screen) {
	e.note.Write()
	s.Fini()
	os.Exit(0)
}

func (e Editor) keyCtrlQ(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}

func (e *Editor) keyRune(event *tcell.EventKey) {
	currentLine := e.note.Contents[e.cursorY]
	newLine := currentLine[:e.cursorX] + string(rune(event.Rune())) + currentLine[e.cursorX:]
	e.note.UpdateLine(newLine, e.cursorY)

	if e.cursorX+1 >= len(e.note.Contents[e.cursorY]) {
		e.cursorX = len(e.note.Contents[e.cursorY])
	} else {
		e.cursorX = e.cursorX + 1
	}
}

func (e Editor) drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (e *Editor) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(defStyle)
	s.Clear()

	// Event loop
	for {
		// Update screen
		s.Clear()

		for i, line := range e.note.Contents {
			e.drawText(s, 0, i, e.characterPerLine, i, defStyle, line)
		}

		s.ShowCursor(e.cursorX, e.cursorY)
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyUp {
				e.keyUp()
			} else if ev.Key() == tcell.KeyDown {
				e.keyDown()
			} else if ev.Key() == tcell.KeyLeft {
				e.keyLeft()
			} else if ev.Key() == tcell.KeyRight {
				e.keyRight()
			} else if ev.Key() == tcell.KeyEnter {
				e.keyEnter()
			} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
				e.keyBackspace()
			} else if ev.Key() == tcell.KeyDEL || ev.Key() == tcell.KeyCtrlD {
				e.keyDel()
			} else if ev.Key() == tcell.KeyCtrlS {
				e.keyCtrlS(s)
			} else if ev.Key() == tcell.KeyCtrlQ || ev.Key() == tcell.KeyEscape {
				e.keyCtrlQ(s)
			} else if ev.Key() == tcell.KeyRune {
				e.keyRune(ev)
			}

		}
	}

}
