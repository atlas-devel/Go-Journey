package main

import "fmt"

type Student struct {
	name  string
	score int
}

func (s *Student) updateScore() {
	s.score += 20
	fmt.Println(s.score)
}

func main() {
	student := Student{
		name:  "Leon",
		score: 90,
	}
	student.updateScore()
	fmt.Println(student.score)
}