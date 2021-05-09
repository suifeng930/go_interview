package main

type Person interface {
	Show()
}
type Student struct {
}

func (stu *Student) Show() {
}

func live() Person {
	var stu *Student
	return stu
}
