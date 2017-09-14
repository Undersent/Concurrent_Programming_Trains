
package graph

import (
	"bufio"
	"configs"
	"fifo"
	"myswitch"
	"os"
	"strconv"
	"track"
)

// EDGE - edge in graph
const EDGE = 0

// VERTEX - vertex in graph
const VERTEX = 1

// Node - graph node
type Node struct {
	typee int
	id    int
}

// Path - graph path
type Path struct {
	Array []*Node
}

// NewPath - create empty path with @size
func NewPath(size int) *Path {
	p := new(Path)
	p.Array = make([]*Node, size)

	return p
}

// NewNode - create new Node
func NewNode(t int, id int) *Node {
	n := new(Node)
	n.typee = t
	n.id = id

	return n
}

// Type - get node type
func (n *Node) Type() int {
	return n.typee
}

// ID - get ID
func (n *Node) ID() int {
	return n.id
}

// Load - load graph from file
func Load() {
	file, _ := os.Open("../configs/graph.txt")
	scanner := bufio.NewScanner(file)

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
		c = c + 2
		oc = c

		t := track.Tracks.GetTrackByID(id)

		/* Parse Vers */
		for line[c-1] != ']' {

			/* get single switch id */
			for line[c] != ';' && line[c] != ']' {
				c++
			}

			s, _ := strconv.Atoi(line[oc:c])
			c++
			oc = c

			t.InsertVer(s)

		}
		scanner.Scan()
	}

	/* for each line ( switch ) DO */
	for i := 0; i < configs.Conf.NumSwitches(); i++ {
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
		c = c + 2
		oc = c

		s := myswitch.Switches.GetSwitchByID(id)

		/* Parse Vers */
		for line[c-1] != ']' {

			/* get single track id */
			for line[c] != ';' && line[c] != ']' {
				c++
			}

			t, _ := strconv.Atoi(line[oc:c])
			c++
			oc = c

			s.InsertEdge(t)

		}
		scanner.Scan()
	}
}

// FindPath - find path between two nodes
func FindPath(n1 *Node, n2 *Node) *Path {
	var rp *Path
	if n1.Type() == VERTEX && n2.Type() == VERTEX {
		rp = findPathVer(n1.ID(), n2.ID())

		/* delete 1st vertex */
		rp.Array[0] = NewNode(VERTEX, 0)

		/* add FAKE NODE */
		rp.Array[len(rp.Array)-1] = NewNode(VERTEX, 0)
	} else if n1.Type() == VERTEX && n2.Type() == EDGE {
		p1 := findPathVer(n1.ID(), track.Tracks.GetTrackByID(n2.ID()).Vers()[0])
		if track.Tracks.GetTrackByID(n2.ID()).Vers()[1] != 0 {
			p2 := findPathVer(n1.ID(), track.Tracks.GetTrackByID(n2.ID()).Vers()[1])
			if len(p1.Array) <= len(p2.Array) {
				rp = p1
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, track.Tracks.GetTrackByID(n2.ID()).Vers()[0])
			} else {
				rp = p2
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, track.Tracks.GetTrackByID(n2.ID()).Vers()[1])
			}
		} else {
			rp = p1
		}
		/* delete 1st vertex */
		rp.Array[0] = NewNode(VERTEX, 0)
	} else if n1.Type() == EDGE && n2.Type() == VERTEX {
		p1 := findPathVer(track.Tracks.GetTrackByID(n1.ID()).Vers()[0], n2.ID())
		if track.Tracks.GetTrackByID(n1.ID()).Vers()[1] != 0 {
			p2 := findPathVer(track.Tracks.GetTrackByID(n1.ID()).Vers()[1], n2.ID())
			if len(p1.Array) <= len(p2.Array) {
				rp = p1
			} else {
				rp = p2
			}
		} else {
			rp = p1
		}

		/* add fake node */
		rp.Array[len(rp.Array)-1] = NewNode(VERTEX, 0)
	} else {
		vers1 := track.Tracks.GetTrackByID(n1.ID()).Vers()
		vers2 := track.Tracks.GetTrackByID(n2.ID()).Vers()

		if vers1[1] != 0 && vers2[1] != 0 {
			p1 := findPathVer(vers1[0], vers2[0])
			p2 := findPathVer(vers1[0], vers2[1])
			p3 := findPathVer(vers1[1], vers2[0])
			p4 := findPathVer(vers1[1], vers2[1])

			/* find minimum */
			if len(p1.Array) <= len(p2.Array) {
				rp = p1
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
			} else {
				rp = p2
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[1])
			}

			if len(rp.Array) > len(p3.Array) {
				rp = p3
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
			}
			if len(rp.Array) > len(p4.Array) {
				rp = p4
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[1])
			}
		} else if vers1[1] != 0 {
			p1 := findPathVer(vers1[0], vers2[0])
			p2 := findPathVer(vers1[1], vers2[0])

			/* find minimum */
			if len(p1.Array) <= len(p2.Array) {
				rp = p1
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
			} else {
				rp = p2
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
			}
		} else if vers2[1] != 0 {
			p1 := findPathVer(vers1[0], vers2[0])
			p2 := findPathVer(vers1[0], vers2[1])

			/* find minimum */
			if len(p1.Array) <= len(p2.Array) {
				rp = p1
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
			} else {
				rp = p2
				rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[1])
			}
		} else {
			rp = findPathVer(vers1[0], vers2[0])
			rp.Array[len(rp.Array)-1] = NewNode(VERTEX, vers2[0])
		}
	}

	return rp
}

func findPathVer(v1 int, v2 int) *Path {
	if v1 == v2 {
		return NewPath(1)
	}

	visited := make([]bool, configs.Conf.NumSwitches()+1)
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}

	tpath := make([]*track.Track, configs.Conf.NumSwitches()+1)

	queue := fifo.NewFifo()

	/* Visit Start Point */
	visited[v1] = true
	queue.Enqueue(v1)

	for !queue.IsEmpty() {
		var ver int
		ver = queue.Dequeue().(int)

		if ver == v2 {

			/* get real path from tpath */
			qp := fifo.NewFifo()
			var c int
			c = 0

			var vers []int
			vers = tpath[ver].Vers()
			for vers[0] != v1 && vers[1] != v1 {
				qp.Enqueue(NewNode(EDGE, tpath[ver].ID()))
				c++
				if vers[0] != ver {
					qp.Enqueue(NewNode(VERTEX, vers[0]))
					c++
					ver = vers[0]
				} else {
					qp.Enqueue(NewNode(VERTEX, vers[1]))
					c++
					ver = vers[1]
				}
				vers = tpath[ver].Vers()
			}

			/* Insert last */
			qp.Enqueue(NewNode(EDGE, tpath[ver].ID()))
			c++
			if vers[0] != ver {
				qp.Enqueue(NewNode(VERTEX, vers[0]))
				c++
				ver = vers[0]
			} else {
				qp.Enqueue(NewNode(VERTEX, vers[1]))
				c++
				ver = vers[1]
			}

			/* Create Path with reverse order */
			path := NewPath(c + 1)
			for !qp.IsEmpty() {
				c--
				path.Array[c] = qp.Dequeue().(*Node)
			}
			return path
		}

		/* Get edges */
		edges := myswitch.Switches.GetSwitchByID(ver).Edges()
		var i int
		i = 0
		for edges[i] != 0 {
			vers := track.Tracks.GetTrackByID(edges[i]).Vers()
			if vers[0] != ver {
				if vers[0] != 0 && !visited[vers[0]] {
					/* visit new ver */
					queue.Enqueue(vers[0])
					visited[vers[0]] = true

					/* save edge */
					tpath[vers[0]] = track.Tracks.GetTrackByID(edges[i])
				}
			} else if vers[1] != ver {
				if vers[1] != 0 && !visited[vers[1]] {
					/* visit new ver */
					queue.Enqueue(vers[1])
					visited[vers[1]] = true

					/* save edge */
					tpath[vers[1]] = track.Tracks.GetTrackByID(edges[i])
				}
			}
			i++
		}
	}
	return nil
}
