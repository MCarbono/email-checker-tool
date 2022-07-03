package entity

import "fmt"

type Email struct {
	Local  string
	Domain Domain
}

func (e Email) String() string {
	return fmt.Sprintf("Local: %v%v", e.Local, e.Domain)
}
