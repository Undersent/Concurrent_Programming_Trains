
package track

import (
	"bufio"
	"configs"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// STATION - track is a station
const STATION = 0

// NORMAL - normal track
const NORMAL = 1

// Track - "protected" track
type Track struct {
	// common
	id     int
	free   bool
	broken bool
	typee  int
	vers   []int
	curVer int
	mutex  sync.Mutex

	// station
	hTime float64

	// normal
	len   int
	speed int
}

// ATracks array with Track
type ATracks struct {
	tracks []*Track
}

// Tracks - global Array with Tracks
var Tracks ATracks

// newTrack - private constructor
func newTrack(n int) *Track {
	t := new(Track)
	t.id = 0
	t.free = true
	t.broken = false
	t.typee = NORMAL
	t.hTime = 0.0
	t.len = 0
	t.speed = 0

	t.vers = make([]int, n)
	for i := 0; i < len(t.vers); i++ {
		t.vers[i] = 0
	}

	t.curVer = 0

	return t
}

// NewNormalTrack - Create Normal Track
/*
   PARAMS
   @IN id - id
   @IN len - len
   @IN speed - speed
   @IN n - size of vers array
*/
func NewNormalTrack(id int, len int, speed int, n int) *Track {
	t := newTrack(n)
	t.typee = NORMAL
	t.id = id
	t.len = len
	t.speed = speed

	return t
}

// NewStationTrack - Create Station Track
/*
   PARAMS
   @IN id - id
   @IN hTime- halt time
   @IN n - size of vers array
*/
func NewStationTrack(id int, hTime float64, n int) *Track {
	t := newTrack(n)
	t.typee = STATION
	t.id = id
	t.hTime = hTime

	return t
}

// ID - get id
func (t *Track) ID() int {
	return t.id
}

// Type - get Track type
func (t *Track) Type() int {
	return t.typee
}

// Vers - Get vers
func (t *Track) Vers() []int {
	return t.vers
}

// HTime - get halt time
func (t *Track) HTime() float64 {
	return t.hTime
}

// Len - get len
func (t *Track) Len() int {
	return t.len
}

// Speed get Available speed
func (t *Track) Speed() int {
	return t.speed
}

// SetID - set id
func (t *Track) SetID(id int) {
	t.id = id
}

// SetType - set type
func (t *Track) SetType(typee int) {
	t.typee = typee
}

// SetHTime - set halt time
func (t *Track) SetHTime(time float64) {
	t.hTime = time
}

// SetLen - set len
func (t *Track) SetLen(len int) {
	t.len = len
}

// SetSpeed - set speed
func (t *Track) SetSpeed(speed int) {
	t.speed = speed
}

// Free - set track as free
func (t *Track) Free() {
	t.mutex.Unlock()
	t.free = true
}

// Busy - set track  as busy
func (t *Track) Busy() {
	t.mutex.Lock()
	t.free = false
}

// Fix - fix Track
func (t *Track) Fix() {
	if configs.Conf.Mode() == configs.NOISY {
		fmt.Printf("tor nr %d jest naprawiony \n", t.id)
	}

	t.mutex.Unlock()
	t.broken = false
}

// Breaking - breaking Track
func (t *Track) Breaking() {
	if configs.Conf.Mode() == configs.NOISY {
		fmt.Printf("tor nr %d jest zepsuty\n", t.id)
	}

	t.mutex.Lock()
	t.broken = true
}

// IsFree - is track free ?
func (t *Track) IsFree() bool {
	return t.free
}

// isBroken - is Track Broken ?
func (t *Track) isBroken() bool {
	return t.broken
}

// InsertVer - insert switch id as ver
func (t *Track) InsertVer(ver int) {
	if t.curVer < len(t.vers) {
		t.vers[t.curVer] = ver
		t.curVer++
	}
}

// Show - show track info
func (t *Track) Show() {
	if t.typee == NORMAL {
		fmt.Println("tor:")
	} else {
		fmt.Println("stacja:")
	}

	fmt.Printf("ID: %d\n", t.id)

	if t.broken {
		fmt.Println("zepsuty!")
	}

	if t.free {
		fmt.Println("stan: wolny")
	} else {
		fmt.Println("stan: zajety")
	}

	if t.typee == NORMAL {
		fmt.Printf("dlugosc: %d\n", t.len)
		fmt.Printf("aktualna predkosc: %d\n", t.speed)
	} else {
		fmt.Printf("czas postoju: %f\n", t.hTime)
	}

	fmt.Printf("porusza sie po wierzcholkach: [")
	for i := 0; i < len(t.vers); i++ {
		if t.vers[i] != 0 {
			fmt.Printf("%d ", t.vers[i])
		}
	}

	fmt.Printf("]\n\n")
}

// NewTracks - create array with Tracks
func (t *ATracks) NewTracks(n int) {
	t.tracks = make([]*Track, n)
}

// Insert - insert track to array
func (t *ATracks) Insert(tr *Track) {
	if tr.ID() >= 1 && tr.ID() <= len(t.tracks) {
		t.tracks[tr.ID()-1] = tr
	}
}

// Tracks - get Array with tracks
func (t *ATracks) Tracks() []*Track {
	return t.tracks
}

// GetTrackByID - get Track by id
func (t *ATracks) GetTrackByID(id int) *Track {
	return t.tracks[id-1]
}

// Show - show all Tracks
func (t *ATracks) Show() {
	fmt.Println("wszystkie tory")

	for i := 0; i < len(t.tracks); i++ {
		t.tracks[i].Show()
	}

	fmt.Println("")
}

// Load - load Tracks from file
func (t *ATracks) Load() {
	file, _ := os.Open("../configs/tracks.txt")
	scanner := bufio.NewScanner(file)
	var typee int
	/* SKIP ALL COMMENTS */
	for scanner.Scan() && scanner.Text()[0] == '#' {
	}

	/* for each line ( track ) DO */
	for i := 0; i < configs.Conf.NumTracks(); i++ {
		line := scanner.Text()

		/* Start from begining */
		oc := 0
		c := 0

		/* find ID value */
		for line[c] != ';' {
			c++
		}

		/* Get ID value */
		id, _ := strconv.Atoi(line[oc:c])
		c++
		oc = c

		/* find Type value */
		for line[c] != ';' {
			c++
		}

		/* Get Type */
		if line[oc:c] == "NORMAL" {
			typee = NORMAL
		} else {
			typee = STATION
		}

		c++
		oc = c

		/* Load normal */
		if typee == NORMAL {
			/* find len value */
			for line[c] != ';' {
				c++
			}

			/* Get len value */
			l, _ := strconv.Atoi(line[oc:c])
			c++
			oc = c

			/* Get Available speed */
			speed, _ := strconv.Atoi(line[oc:len(line)])

			/* Create and Add new Track */
			track := NewNormalTrack(id, l, speed, configs.Conf.NumSwitches())
			t.Insert(track)
		} else { /* Load station */

			/* get halt time */
			hTime, _ := strconv.ParseFloat(line[oc:len(line)], 64)

			/* Create and Add new Track */
			track := NewStationTrack(id, hTime, configs.Conf.NumSwitches())
			t.Insert(track)
		}

		scanner.Scan()
	}

}
