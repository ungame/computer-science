package subjects

import "fmt"

type Subject struct {
	Name  string
	Items []Subject
}

func (s Subject) Print() {
	fmt.Println(s.Name)
	for _, subject := range s.Items {
		subject.printChild(s.Name)
	}
}

func (s Subject) printChild(parent string) {
	fmt.Printf("%s/%s\n", parent, s.Name)
	for _, subject := range s.Items {
		subject.printChild(s.Name)
	}
}
