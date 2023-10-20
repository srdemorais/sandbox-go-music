package musicalnotes

import "fmt"

var notes = [7]string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}

type MusicalNote struct {
	Note  string
	Sound string
}

func (n *MusicalNote) TestUser() {
	(*n).CheckNext()
	(*n).CheckPrevious()
}

func (n *MusicalNote) GetNext() string {
	var next string
	for i, v := range notes {
		if v == (*n).Note {
			if i == 6 {
				next = notes[0]
			} else {
				next = notes[i+1]
			}
			break
		}
	}
	return next
}

func (n *MusicalNote) GetPrevious() string {
	var next string
	for i, v := range notes {
		if v == (*n).Note {
			if i == 0 {
				next = notes[6]
			} else {
				next = notes[i-1]
			}
			break
		}
	}
	return next
}

func (n *MusicalNote) CheckNext() {
	next := (*n).GetNext()
	fmt.Printf("What is the note after \"%v\" ? ", (*n).Note)

	var iNext string
	fmt.Scanln(&iNext)

	fmt.Println("Result: ", next == iNext)
}

func (n *MusicalNote) CheckPrevious() {
	previous := (*n).GetPrevious()
	fmt.Printf("What is the note before \"%v\" ? ", (*n).Note)

	var iPrevious string
	fmt.Scanln(&iPrevious)

	fmt.Println("Result: ", previous == iPrevious)
}
