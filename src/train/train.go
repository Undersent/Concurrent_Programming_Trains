
package train

import (
	"bufio"
	"configs"
	"fmt"
	"os"
	"strconv"
)

// POS_NONE - start position (train is not in use)
const POS_NONE = 0

// POS_STATION - station
const POS_STATION = 1

// POS_TRACK - track
const POS_TRACK = 2

// POS_SWITCH - switch
const POS_SWITCH = 3

// Train - train structure only private fields.
type Train struct {
	id         int
	maxSpeed   int
	capacity   int
	startPoint int
	posType    int
	posID      int
	route      []int
	curRoute   int
}

// ATrains - Array of Train
type ATrains struct {
	trains   []*Train
}

// Trains Global Array of Trains
var Trains ATrains

// NewTrain - Create new Train
/*
	PARAMS
	@IN id - ID
	@IN maxSpeed - MaxSpeed
	@IN Capacity - Capacity
	@IN startPoint - startPoint
	@IN n - size of Route
*/
func NewTrain(id int, maxSpeed int, capacity int, startPoint int, n int) *Train {
	train := new(Train)

	train.id = id
	train.maxSpeed = maxSpeed
	train.capacity = capacity
	train.startPoint = startPoint
	train.posType = POS_NONE
	train.posID = 0

	train.route = make([]int, n)
	for i := 0; i < len(train.route); i++ {
		train.route[i] = 0
	}

	train.curRoute = 0

	return train
}

// ID - get ID.
func (t *Train) ID() int {
	return t.id
}

// MaxSpeed - get maxSpeed.
func (t *Train) MaxSpeed() int {
	return t.maxSpeed
}

// Capacity - get capcity.
func (t *Train) Capacity() int {
	return t.capacity
}

// StartPoint - get StartPoint.
func (t *Train) StartPoint() int {
	return t.startPoint
}

// PosType - get posType.
func (t *Train) PosType() int {
	return t.posType
}

// PosID - get PosID.
func (t *Train) PosID() int {
	return t.posID
}

// Route - get Route
func (t *Train) Route() []int {
	return t.route
}

// CurRoute - get curRoute
func (t *Train) CurRoute() int {
	return t.curRoute
}

// SetID - set ID
func (t *Train) SetID(id int) {
	t.id = id
}

// SetMaxSpeed - set
func (t *Train) SetMaxSpeed(s int) {
	t.maxSpeed = s
}

// SetCapacity - set Capacity
func (t *Train) SetCapacity(c int) {
	t.capacity = c
}

// SetStartPoint - set StartPoint
func (t *Train) SetStartPoint(sp int) {
	t.startPoint = sp
}

// ChangePos - change posision
func (t *Train) ChangePos(pos int, id int) {
	t.posType = pos
	t.posID = id
}

// AddRoute - add track id to route
func (t *Train) AddRoute(id int) {
	if t.curRoute < len(t.route) {
		t.route[t.curRoute] = id
		t.curRoute++
	}
}

// ShowPos - show only posision
func (t *Train) ShowPos() {
	fmt.Println("pociag:")
	fmt.Printf("ID: %d\n", t.id)

	fmt.Printf("aktualna pozycja: ")
	switch t.posType {
	case POS_STATION:
		fmt.Printf("stacja %d\n", t.posID)
	case POS_TRACK:
		fmt.Printf("tor %d\n", t.posID)
	case POS_SWITCH:
		fmt.Printf("zwrotnica %d\n", t.posID)
	default:
		fmt.Printf("nie uzywany\n")
	}
}

// Show - show train info
func (t *Train) Show() {
	fmt.Println("Pociag:")
	fmt.Printf("ID: %d\n", t.id)
	fmt.Printf("max predkosc: %d\n", t.maxSpeed)
	fmt.Printf("pojemnosc: %d\n", t.capacity)
	fmt.Printf("poczatek trasy: %d\n", t.startPoint)
	fmt.Printf("droga: [")
	for i := 0; i < len(t.route); i++ {
		if t.route[i] != 0 {
			fmt.Printf("%d ", t.route[i])
		}
	}
	fmt.Println("]")

	fmt.Printf("aktualna pozycja: ")
	switch t.posType {
	case POS_STATION:
		fmt.Printf("stacja %d\n", t.posID)
	case POS_TRACK:
		fmt.Printf("tor %d\n", t.posID)
	case POS_SWITCH:
		fmt.Printf("zwrotnica %d\n", t.posID)
	default:
		fmt.Printf("nie uzywany\n")
	}

	fmt.Println("")
}

// NewTrains - create array with Trains
func (t *ATrains) NewTrains(n int) {
	t.trains = make([]*Train, n)
}

// Insert - insert new train
func (t *ATrains) Insert(train *Train) {
	if train.ID() >= 1 && train.ID() <= len(t.trains) {
		t.trains[train.ID() - 1] = train
	}
}

// Trains - get trains
func (t *ATrains) Trains() []*Train {
	return t.trains
}

// GetTrainByID - get Train by id
func (t *ATrains) GetTrainByID(id int) *Train {
	return t.trains[id-1]
}

// Show - show all info about Trains
func (t *ATrains) Show() {
	fmt.Println("wszystkie pociagi")

	for i := 0; i < len(t.trains); i++ {
		t.trains[i].Show()
	}

	fmt.Println("")
}

// ShowPos show only posision
func (t *ATrains) ShowPos() {
	fmt.Println("wszystkie pociagi - pozycje")

	for i := 0; i < len(t.trains); i++ {
		t.trains[i].ShowPos()
	}

	fmt.Println("")
}

// Load - Load Trains from file
func (t *ATrains) Load() {
	file, _ := os.Open("../configs/trains.txt")
	scanner := bufio.NewScanner(file)

	/* SKIP ALL COMMENTS */
	for scanner.Scan() && scanner.Text()[0] == '#' {
	}

	/* for each line ( train ) DO */
	for i := 0; i < configs.Conf.NumTrains(); i++ {
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

		/* find MaxSpeed value */
		for line[c] != ';' {
			c++
		}

		/* Get MaxSpeed value */
		speed, _ := strconv.Atoi(line[oc:c])
		c++
		oc = c

		/* find Capacity value */
		for line[c] != ';' {
			c++
		}

		/* Get Capacity value */
		capacity, _ := strconv.Atoi(line[oc:c])
		c++
		oc = c

		/* find StartPoint value */
		for line[c] != ';' {
			c++
		}

		/* Get StartPoint value */
		startPoint, _ := strconv.Atoi(line[oc:c])
		c = c + 2
		oc = c

		/* Create Train and add to global array */
		train := NewTrain(id, speed, capacity, startPoint, configs.Conf.NumTracks())
		Trains.Insert(train)

		/* Parse Route */
		for line[c-1] != ']' {
			for line[c] != ';' && line[c] != ']' {
				c++
			}

			route, _ := strconv.Atoi(line[oc:c])
			c++
			oc = c

			train.AddRoute(route)
		}
		scanner.Scan()
	}
	file.Close()
}
