# SHIN: Shin's Hyper Interesting Notes

## Specifications

### Notes
* Notes are stored in the $SHIN_STORAGE directory if that is set.
* If SHIN_STORAGE is not set in the environment, notes are stored in the
  $HOME/.shin directory.
* Notes are numbered starting from 1.
* Notes are saved as <number>.shin in the storage directory mentioned above.
* To find out the number of a new note, shin looks at all notes in the storage
  directory and takes the highest existing note nr + 1.
* Notes are plain text files encoded in UTF-8 (strings in Go are normally UTF-8).

### Command line

* `shin` : Starts taking a new note in the terminal editor (see below).
* `shin <number>` : Edits the given numbered note. Makes it new if it doesn't exist.
* `shin -l` : lists all notes and their first non-empty line of text.
* `shin -c <number>` : prints the numbered numbered note to standard output.
* `shin -d <number>` : deletes the numbered note.


### Terminal Editor
* To implement a simple terminal editor, you can use the tcell library found here:
https://github.com/gdamore/tcell

* The terminal editor has a cursor.
* In the terminal editor you can type in text.
* Text will be inserted at the cursor position.
* Enter starts a new line and moves the cursor down one line.
* Backspace deletes the character before and moves the cursor back.
* Del deletes the character after and does not move the cursor.
* You can move around the cursor with the arrow keys.
* You can save and quit with Control-S.
* You can abandon the note with Control-Q.


## Stretch Goals:
* `shin -g /regexp/`: searches all notes for the regexp and returns the matching
  notes with the matching text line.
* Extra commands for the editor.
* Make sure it also works for Japanese or UTF-8 input.
* Support for the EDITOR environment variable.
* `shin <filename>` : Edits the given file name. It is stored as the file is named, and not as a note.


