package musicalnotes

var notes = [7]string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}

type MusicalNote struct {
	note  string
	sound string
}

func (n *MusicalNote) getNext() string {
	var next string
	for i, v := range notes {
		if v == (*n).note {
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

func (n *MusicalNote) getPrevious() string {
	var next string
	for i, v := range notes {
		if v == (*n).note {
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
