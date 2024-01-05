package musicalnotes

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

var notes = [7]string{"Do", "Re", "Mi", "Fa", "Sol", "La", "Si"}
var positions = [7]int{0, 1, 2, 3, 4, 5, 6}
var outnotes [3]string
var outsounds [3]string

func getOut(note string, outnotes *[3]string, outsounds *[3]string) {
	rand.Seed(time.Now().UnixNano())

	var notes = [7]string{"Do", "Re", "Mi", "Fa", "Sol", "La", "Si"}
	var sounds = [7]string{"C", "D", "E", "F", "G", "A", "B"}

	var idx int
	var randpos [3]int

	// find the index of the input note
	for i, v := range notes {
		if v == note {
			idx = i
			break
		}
	}

	// initialize randpos
	randpos[0] = idx
	p := 1
	for i := 0; i < 7; i++ {
		i = rand.Intn(7)
		if i != idx {
			randpos[p] = i
			p++
		}
		if p > 2 {
			break
		}
	}

	// shuffle randpos
	rand.Shuffle(len(randpos), func(i, j int) { randpos[i], randpos[j] = randpos[j], randpos[i] })

	// set outnotes
	for i := 0; i < 3; i++ {
		outnotes[i] = notes[randpos[i]]
	}

	// set outsounds
	for i := 0; i < 3; i++ {
		outsounds[i] = sounds[randpos[i]]
	}
}

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
	return (*n).CheckNext() && (*n).CheckPrevious() && (*n).CheckPosition() && (*n).CheckSound()
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

func (n *MusicalNote) CheckSound() bool {
	getOut((*n).Note, &outnotes, &outsounds)

	p := -1
	for i := 0; i < 3; i++ {
		if outnotes[i] == (*n).Note {
			p = i + 1
			break
		}
	}

	for _, v := range outsounds {
		if err := RunNote(v); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("What is the note position ? ")

	var iPosition int
	fmt.Scanln(&iPosition)

	fmt.Println(outnotes)
	fmt.Println(outsounds)
	fmt.Println(p == iPosition)

	return p == iPosition
}

func RunNote(note string) error {
	f, err := os.Open("mp3/" + note + ".mp3")
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
