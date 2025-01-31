package musicalnotes

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

var notes = [...]string{"C2", "Db2", "D2", "Eb2", "E2", "F2", "Gb2", "G2", "Ab2", "A2", "Bb2", "B2", "C3", "Db3", "D3", "Eb3", "E3", "F3", "Gb3", "G3", "Ab3", "A3", "Bb3", "B3", "C4", "Db4", "D4", "Eb4", "E4", "F4", "Gb4", "G4", "Ab4", "A4", "Bb4", "B4", "C5", "Db5", "D5", "Eb5", "E5", "F5", "Gb5", "G5", "Ab5", "A5", "Bb5", "B5", "C6"}
var notesPos = [...]int{1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7, 8, 9, 9, 10, 10, 11, 12, 12, 13, 13, 14, 14, 15, 16, 16, 17, 17, 18, 19, 19, 20, 20, 21, 21, 22, 23, 23, 24, 24, 25, 26, 26, 27, 27, 28, 28, 29}

type MusicalNote struct {
	Idx       int
	Note      string
	AudioPath string
}

func Init() MusicalNote {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	idx := rand.Intn(len(notes))

	var oMusicalNote MusicalNote
	oMusicalNote.Idx = idx
	oMusicalNote.Note = notes[idx]
	oMusicalNote.AudioPath = "mp3/" + notes[idx] + ".mp3"

	return oMusicalNote
}

func (n *MusicalNote) TestUser() bool {
	return n.CheckNext() && n.CheckPrevious() && n.CheckPosition() && n.CheckSound()
}

func (n *MusicalNote) CheckNext() bool {
	next := n.GetNext()
	fmt.Printf("What is the note after \"%v\" ? ", n.Note)

	var iNext string
	fmt.Scanln(&iNext)

	return strings.ToUpper(next) == strings.ToUpper(iNext)
}

func (n *MusicalNote) GetNext() string {
	var next string
	if n.Note == "C6" {
		next = "Db6"
	} else {
		next = notes[n.Idx+1]
	}
	return next
}

func (n *MusicalNote) CheckPrevious() bool {
	previous := n.GetPrevious()
	fmt.Printf("What is the note before \"%v\" ? ", n.Note)

	var iPrevious string
	fmt.Scanln(&iPrevious)

	return strings.ToUpper(previous) == strings.ToUpper(iPrevious)
}

func (n *MusicalNote) GetPrevious() string {
	var previous string
	if n.Note == "C2" {
		previous = "B1"
	} else {
		previous = notes[n.Idx-1]
	}
	return previous
}

func (n *MusicalNote) CheckPosition() bool {
	DisplayStaff()

	fmt.Printf("What is the note position ? ")

	var iPosition int
	fmt.Scanln(&iPosition)

	return notesPos[n.Idx] == iPosition
}

func (n *MusicalNote) CheckSound() bool {

	guessNotes, pos := n.getGuessNotes()

	for _, v := range guessNotes {
		if err := RunNote(v.AudioPath); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("What is the note position ? ")

	var iPos int
	fmt.Scanln(&iPos)

	for _, v := range guessNotes {
		fmt.Printf("%v - ", v.Note)
	}
	fmt.Println("")

	for i := 0; i < 4; i++ {
		fmt.Printf("Again... ")

		for _, v := range guessNotes {
			if err := RunNote(v.AudioPath); err != nil {
				log.Fatal(err)
			}
		}
	}

	return pos+1 == iPos
}

func (n *MusicalNote) getGuessNotes() ([6]MusicalNote, int) {

	var guessNotes [6]MusicalNote
	var pos int

	var randPos [6]int
	randPos[0] = n.Idx

	rand.Seed(time.Now().UnixNano())

	p := 1
	for {
		tmpIdx := rand.Intn(len(notes))
		if tmpIdx != n.Idx {
			randPos[p] = tmpIdx
			p++
		}
		if p > len(randPos)-1 {
			break
		}
	}

	// shuffle randpos
	rand.Shuffle(len(randPos), func(i, j int) { randPos[i], randPos[j] = randPos[j], randPos[i] })

	// set guessNotes
	for i := 0; i < 6; i++ {
		guessNotes[i].Idx = randPos[i]
		guessNotes[i].Note = notes[randPos[i]]
		guessNotes[i].AudioPath = "mp3/" + notes[randPos[i]] + ".mp3"

		if n.Note == guessNotes[i].Note {
			pos = i
		}
	}

	return guessNotes, pos
}

func RunNote(audioPath string) error {
	f, err := os.Open(audioPath)
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

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

/*
	  In Western musical notation, the staff (UK also stave; plural: staffs or staves),
		also occasionally referred to as a pentagram, is a set of five horizontal lines and
		four spaces that each represent a different musical pitch or in the case of a
		ercussion staff, different percussion instruments.
*/
func DisplayStaff() {
	fmt.Println("")
	fmt.Println("C6                      ---                   29")
	fmt.Println("                                              28")
	fmt.Println("                        ---                   27")
	fmt.Println("                                              26")
	fmt.Println("           ----------------------------       25")
	fmt.Println("                                              24")
	fmt.Println("           ----------------------------       23")
	fmt.Println("C5                                            22")
	fmt.Println("           ----------------------------       21")
	fmt.Println("                                              20")
	fmt.Println("G4         ----------------------------       19")
	fmt.Println("                                              18")
	fmt.Println("           ----------------------------       17")
	fmt.Println("                                              16")
	fmt.Println("C4                      ---                   15")
	fmt.Println("                                              14")
	fmt.Println("           ----------------------------       13")
	fmt.Println("                                              12")
	fmt.Println("F3         ----------------------------       11")
	fmt.Println("                                              10")
	fmt.Println("           ----------------------------       9")
	fmt.Println("C3                                            8")
	fmt.Println("           ----------------------------       7")
	fmt.Println("                                              6")
	fmt.Println("           ----------------------------       5")
	fmt.Println("                                              4")
	fmt.Println("                        ---                   3")
	fmt.Println("                                              2")
	fmt.Println("C2                      ---                   1")
	fmt.Println("")
}
