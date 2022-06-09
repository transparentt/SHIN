package shin

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Editor struct{}

func Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	//cursorStyle := tcell.StyleDefault.Background(tcell.ColorGhostWhite).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(defStyle)
	s.Clear()

	note := ReadNo(1)
	/*for i, line := range note.Contents {
		drawText(s, 0, i, 79, i, defStyle, line)
	}*/

	// Event loop
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	var cursorX, cursorY int = 0, 0
	for {
		// Update screen
		s.Clear()

		for i, line := range note.Contents {
			drawText(s, 0, i, 79, i, defStyle, line)
		}

		s.ShowCursor(cursorX, cursorY)

		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				quit()
			}

			if ev.Key() == tcell.KeyCtrlS {
				note.Write()
				quit()
			}
			if ev.Key() == tcell.KeyBackspace2 || ev.Key() == tcell.KeyBackspace {

				if cursorX == 0 && cursorY == 0 {
					continue
				}

				currentLine := note.Contents[cursorY]

				if len(currentLine) != 0 {
					if cursorX != 0 {

						if cursorX-1 < 0 {
							cursorX = 0
						} else {
							cursorX = cursorX - 1
						}

						newLine := currentLine[:cursorX] + currentLine[cursorX+1:]
						note.Update(newLine, cursorY)

					} else {

						if cursorY-1 < 0 {
							cursorY = 0
						} else {
							cursorY = cursorY - 1
							cursorX = len(note.Contents[cursorY])
						}

						upperNewLine := note.Contents[cursorY] + currentLine
						note.Update(upperNewLine, cursorY)

						newContents := append(note.Contents[:cursorY+1], note.Contents[cursorY+1+1:]...)
						note.Contents = newContents

					}
				} else {
					if cursorX != 0 {
						//
					} else {

						if cursorY-1 < 0 {
							cursorY = 0
						} else {
							cursorY = cursorY - 1
							cursorX = len(note.Contents[cursorY])
						}

						newContents := append(note.Contents[:cursorY+1], note.Contents[cursorY+1+1:]...)
						note.Contents = newContents

					}
				}
			}

			if ev.Key() == tcell.KeyLeft {
				if cursorX-1 < 0 {
					cursorX = 0
				} else {
					cursorX = cursorX - 1
				}
			}

			if ev.Key() == tcell.KeyRight {
				if cursorX+1 >= len(note.Contents[cursorY]) {
					cursorX = len(note.Contents[cursorY])
				} else {
					cursorX = cursorX + 1
				}
			}

			if ev.Key() == tcell.KeyUp {
				if cursorY-1 < 0 {
					cursorY = 0
				} else {
					cursorY = cursorY - 1
					if cursorX > len(note.Contents[cursorY]) {
						cursorX = len(note.Contents[cursorY])
					}
				}
			}

			if ev.Key() == tcell.KeyDown {
				if cursorY+1 >= len(note.Contents) {
					cursorY = len(note.Contents) - 1
				} else {
					cursorY = cursorY + 1
					if cursorX > len(note.Contents[cursorY]) {
						cursorX = len(note.Contents[cursorY])
					}
				}
			}

		}
	}

}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
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

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	drawText(s, x1, y1, x2-3, y2-3, style, text)
}
