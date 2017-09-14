
package fifo

type node struct {
	next *node
	value interface{}
}

// Fifo - first in first out queue
type Fifo struct {
    head *node
    tail *node
}

// NewFifo - create new Fifo
func NewFifo() *Fifo {
    f := new(Fifo)

    f.head = nil
    f.tail = nil

    return f;
}

// isEmpty - is Fifo Empty ?
func (f *Fifo) IsEmpty() bool {
    return f.head == nil
}

// Enqueue - add val to fifo
func (f *Fifo) Enqueue(val interface{}) {
    n := new(node)
    n.next = nil
    n.value = val

    if f.IsEmpty() {
        f.tail = n
        f.head = n
        f.head.next = f.tail
    } else {
        f.tail.next = n
        f.tail = f.tail.next
    }
}

// Dequeue - delete cal from fifo
func (f *Fifo) Dequeue() interface {} {
    if f.IsEmpty() {
        return nil
    }

    val := f.head.value
    f.head = f.head.next

    if f.head == nil {
        f.tail = nil
    }

    return val
}

// GetHead - get first value
func (f *Fifo) GetHead() interface{} {
    return f.head.value
}
