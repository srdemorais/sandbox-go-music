package musicalnotes

var notes = [7]string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}

type MusicalNote struct {
	Note  string
	Sound string
}

func (n *MusicalNote) GetNext() string {
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

func (n *MusicalNote) GetPrevious() string {
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
