package musicalnotes

import (
	"fmt"
	"math/rand"
	"time"
	"io"
	"os"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

var notes = [7]string{"Do", "Re", "Mi", "Fa", "Sol", "La", "Si"}
var positions = [7]int{0, 1, 2, 3, 4, 5, 6}

type MusicalNote struct {
	Note     string
	Position int
	Sound    string
}

func Init() MusicalNote {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	position := rand.Intn(len(notes))

	var oMusicalNote MusicalNote
	oMusicalNote.Position = position
	oMusicalNote.Note = notes[position]

	return oMusicalNote
}

func (n *MusicalNote) TestUser() bool {
	return (*n).CheckNext() && (*n).CheckPrevious() && (*n).CheckPosition()
}

func (n *MusicalNote) GetNext() string {
	var next string
	if (*n).Position == 6 {
		next = notes[0]
	} else {
		next = notes[(*n).Position+1]
	}
	return next
}

func (n *MusicalNote) GetPrevious() string {
	var next string
	if (*n).Position == 0 {
		next = notes[6]
	} else {
		next = notes[(*n).Position-1]
	}
	return next
}

func (n *MusicalNote) CheckNext() bool {
	next := (*n).GetNext()
	fmt.Printf("What is the note after \"%v\" ? ", (*n).Note)

	var iNext string
	fmt.Scanln(&iNext)

	return next == iNext
}

func (n *MusicalNote) CheckPrevious() bool {
	previous := (*n).GetPrevious()
	fmt.Printf("What is the note before \"%v\" ? ", (*n).Note)

	var iPrevious string
	fmt.Scanln(&iPrevious)

	return previous == iPrevious
}

func (n *MusicalNote) CheckPosition() bool {
	fmt.Println("")
	fmt.Println("11---------------------0------")
	fmt.Println("10                   0        ")
	fmt.Println("9------------------0----------")
	fmt.Println("8                0            ")
	fmt.Println("7--------------0--------------")
	fmt.Println("6            0                ")
	fmt.Println("5----------0------------------")
	fmt.Println("4        0                    ")
	fmt.Println("3------0----------------------")
	fmt.Println("2    0                        ")
	fmt.Println("1 -0-                         ")
	fmt.Println("")

	fmt.Printf("What is the note position ? ")

	var iPosition int
	fmt.Scanln(&iPosition)

	return (*n).Position+1 == iPosition
}

func RunNote(note string) error {
	f, err := os.Open("mp3/" + note + ".mp3")
	fmt.Println("mp3/" + note + ".mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

  fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
