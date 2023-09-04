package main

import "fmt"

type Bootcamp struct {
	Organization string
	Language     string
	Students     []Student
}

type Student struct {
	Name      string
	EntryDate string
}

func NewBootcamp(org, lan string) Bootcamp {
	return Bootcamp{
		Organization: org,
		Language:     lan,
	}
}

func NewStudent(name, entry string) Student {
	return Student{
		Name:      name,
		EntryDate: entry,
	}
}

func (b *Bootcamp) AddStudent(s Student) {
	b.Students = append(b.Students, s)
}

func main() {
	bootcampGo := NewBootcamp("MELI", "GO")

	fmt.Println(bootcampGo)

	student1 := NewStudent("Juan", "05/02/2023")

	student2 := NewStudent("Martin", "05/02/2023")

	student3 := NewStudent("Paula", "01/02/2023")

	bootcampGo.AddStudent(student1)
	bootcampGo.AddStudent(student2)
	bootcampGo.AddStudent(student3)

	fmt.Println(bootcampGo)
}
