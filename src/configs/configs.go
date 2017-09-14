
package configs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// SILENT - silent mode, menu is enabled, no default print on stdout
const SILENT = 0

// NOISY - noisy mode, menu is disabled, but a lot prints on stdout
const NOISY = 1

// Config - main config
type Config struct {
	mode        int
	sPerH       int
	numTrains   int
	numTracks   int
	numSwitches int
	probability int
}

// Conf - global Configs
var Conf Config

/* GETTERS */

// Mode - get mode
func (c *Config) Mode() int {
	return c.mode
}

// SPerH - get seconds per hour
func (c *Config) SPerH() int {
	return c.sPerH
}

// NumTrains - get number of trains
func (c *Config) NumTrains() int {
	return c.numTrains
}

// NumTracks - get number of tracks
func (c *Config) NumTracks() int {
	return c.numTracks
}

// NumSwitches - get number of switches
func (c *Config) NumSwitches() int {
	return c.numSwitches
}

// Probability - get break probability
func (c *Config) Probability() int {
	return c.probability
}

/* SETTERS */

// SetMode - set mode
func (c *Config) SetMode(m int) {
	c.mode = m
}

// SetSPerH - set seconds per hour
func (c *Config) SetSPerH(s int) {
	c.sPerH = s
}

// SetNumTrains - set number of trains
func (c *Config) SetNumTrains(n int) {
	c.numTrains = n
}

// SetNumTracks - set number of tracks
func (c *Config) SetNumTracks(n int) {
	c.numTracks = n
}

// SetNumSwitches - set number of switches
func (c *Config) SetNumSwitches(n int) {
	c.numSwitches = n
}

// SetProbability - set break probability
func (c *Config) SetProbability(p int) {
	c.probability = p
}

// Show - show configs
func (c *Config) Show() {
	fmt.Println("konfiguracja:")

	if c.mode == SILENT {
		fmt.Println("tryb cichy")
	} else {
		fmt.Println("tryb glosny")
	}

	fmt.Printf("s/h: %d\n", c.sPerH)
	fmt.Printf("liczba pociagow: %d\n", c.numTrains)
	fmt.Printf("liczba zwrotnic: %d\n", c.numSwitches)
	fmt.Printf("liczba torow: %d\n\n", c.numTracks)
}

// Load - load configs from file
func (c *Config) Load() {
	file, _ := os.Open("../configs/conf.txt")
	scanner := bufio.NewScanner(file)

	/* SKIP ALL COMMENTS */
	for scanner.Scan() && scanner.Text()[0] == '#' {
	}

	/* SET MODE */
	if scanner.Text() == "SILENT" {
		c.mode = SILENT
	} else if scanner.Text() == "NOISY" {
		c.mode = NOISY
	}

	/* SET Seconds per hour */
	scanner.Scan()
	c.sPerH, _ = strconv.Atoi(scanner.Text())

	/* SET Num of Trains */
	scanner.Scan()
	c.numTrains, _ = strconv.Atoi(scanner.Text())

	/* SET Num of Switches */
	scanner.Scan()
	c.numSwitches, _ = strconv.Atoi(scanner.Text())

	/* SET num of Tracks */
	scanner.Scan()
	c.numTracks, _ = strconv.Atoi(scanner.Text())

	/* SET Probability */
	scanner.Scan()
	c.probability, _ = strconv.Atoi(scanner.Text())

	file.Close()
}
