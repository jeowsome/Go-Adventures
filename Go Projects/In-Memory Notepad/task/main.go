package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NoteBook struct {
	maxNoteCount    int
	filledNoteCount int
	notes           []Notes
}

type Notes struct {
	no      int
	content string
}

func createNote(n *NoteBook, notes string) {
	if n.filledNoteCount < n.maxNoteCount {
		n.notes = append(n.notes, Notes{
			no:      len(n.notes) + 1,
			content: notes,
		})
		fmt.Println("[OK] the note was successfully created")
		n.filledNoteCount++
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}

func modify(n *NoteBook, index int, action, newNote string) []Notes {
	notes := list(n.notes, false)
	var newNotes []Notes

	if action == "del" {
		for _, note := range notes {
			if note.no != index {
				newNotes = append(newNotes, note)
			}
		}
		n.notes = newNotes
	} else if action == "mod" && newNote != "" {
		if index > len(notes) {
			fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", index, len(notes)+1)
			return nil
		}
		for _, note := range n.notes {
			curNote := Notes{
				no:      note.no,
				content: note.content,
			}

			if index == note.no {
				curNote.content = newNote
			}

			newNotes = append(newNotes, curNote)
		}
	}
	n.notes = newNotes
	return notes
}

func list(n []Notes, printNotes bool) []Notes {
	showNotes := printNotes

	if len(n) == 0 {
		if showNotes {
			fmt.Println("[Info] Notepad is empty")
		}
		return nil
	}

	if showNotes {
		for i, note := range n {
			fmt.Printf("[Info] %d: %s\n", i+1, note.content)
		}
	}
	return n
}

func clear(n *NoteBook) {
	n.notes = make([]Notes, 0)
	fmt.Println("[OK] All notes were successfully deleted")
}

func exit() {
	fmt.Println("[Info] Bye!")
}

func deleteNote(n *NoteBook, ind string) {
	if ind == "" {
		fmt.Println("[Error] Missing position argument")
		return
	}

	i, err := strconv.Atoi(ind)

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", ind)
		return
	}

	er := modify(n, i, "del", "")

	if er != nil {
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", i)
		n.filledNoteCount--
	}
}

func update(n *NoteBook, data string) {
	input := strings.Split(data, " ")
	i, err := strconv.Atoi(input[0])

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", input[0])
		return
	}

	note := strings.Join(input[1:], " ")
	if note == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}

	er := modify(n, i, "mod", note)

	if er != nil {
		fmt.Printf("[OK] The note at position %d was successfully updated\n", i)
	}
}

func canDoCommand(n *NoteBook, data, command string) bool {
	canDo := true

	switch {
	case command == "update" && n.filledNoteCount > 0 && data == "":
		fmt.Println("[Error] Missing position argument")
		canDo = false
	case (command == "create" || command == "update") && data == "":
		fmt.Println("[Error] Missing note argument")
		canDo = false
	case n.filledNoteCount == 0 && (command == "update" || command == "delete"):
		fmt.Printf("[Error] There is nothing to %s\n", command)
		canDo = false
	}
	return canDo
}

func startNoteMaking(noteBook *NoteBook) {
	for {
		fmt.Print("Enter a command and data:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		if !canDoCommand(noteBook, data, command) {
			continue
		}

		switch command {
		case "exit":
			exit()
			return
		case "create":
			createNote(noteBook, data)
		case "list":
			list(noteBook.notes, true)
		case "clear":
			clear(noteBook)
		case "delete":
			deleteNote(noteBook, data)
		case "update":
			update(noteBook, data)
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func main() {
	noteBook := new(NoteBook)
	fmt.Print("Enter the maximum number of notes:")
	_, err := fmt.Scanln(&noteBook.maxNoteCount)

	if err != nil {
		fmt.Println(err)
	}

	if noteBook.maxNoteCount != 0 {
		startNoteMaking(noteBook)
	}
}
