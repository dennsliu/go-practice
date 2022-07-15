package main

type Team struct {
	Name string
	Num  int
}

func TeamRegister(name string, num int) *Team {
	s := new(Team)
	s.Name = name
	s.Num = num
	return s
}

func main() {
	TeamRegister("Warrior", 18)
}
