package list

import (
	"fmt"
	"strings"
)

type node struct {
	name string
	next *node
}

type List struct {
	head *node
}

func New() *List {
	return &List{}
}

func (l *List) Push(name string) {
	node := &node{
		name: name,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}

	p := l.head
	for {
		if p.next == nil {
			break
		}
		p = p.next
	}
	p.next = node
}

func (l *List) PushMany(names ...string) {
	for _, name := range names {
		l.Push(name)
	}
}

func (l *List) Print() {
	s := make([]string, 0)
	p := l.head
	for {
		if p == nil {
			break
		}
		s = append(s, p.name)
		p = p.next
	}
	res := strings.Join(s, " -> ")
	fmt.Println(res)

}

func (l *List) Pop() string {
	if l.head == nil {
		return ""
	}
	p := l.head
	l.head = p.next
	p.next = nil
	return p.name
}

func (l *List) Delete(name string) {
	if l.head == nil {
		return
	}
	current := l.head
	var prev *node
	for {
		if current == nil {
			return
		}
		if current.name != name {
			prev = current
			current = current.next
			continue
		}
		if prev == nil {
			l.head = current.next
			current.next = nil
			return
		} else {
			prev.next = current.next
			current.next = nil
			return
		}
	}
}
func (l *List) Exist(name string) bool {
	p := l.head
	for {
		if p == nil {
			break
		}
		if p.name == name {
			return true
		}
		p = p.next
	}
	return false
}

func (l *List) Len() int {
	x := 0
	p := l.head
	for {
		if p == nil {
			break
		}
		x = x + 1
		p = p.next
	}
	return x
}
