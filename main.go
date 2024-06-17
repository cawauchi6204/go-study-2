package main

type Status string

type Body struct {
	Height int
	Weight int
}

const (
	StudentStatusHighSchool Status = Status("HighSchool")
	StudentStatusUniversity Status = Status("University")
	StatusStatusGraduate    Status = Status("Graduated")
)

type Student struct {
	ID     string
	Name   string
	Age    int
	Body   Body
	Status Status
}

func (s Student) Graduate() {
	switch s.Status {
	case StudentStatusHighSchool:
		println("University")
	case StudentStatusUniversity:
		println("Graduated")
	case StatusStatusGraduate:
		println("Already Graduated")
	}
}

// 関数名カッコはメソッドであるという証明
// クラスみたいな感じ
// このメソッドはUser構造体のAgeフィールドを出力する
func (s *Student) OutputAge() {
	println(s.Age)
}

func main() {
	student := Student{ID: "1", Name: "John", Age: 30, Body: Body{Height: 170, Weight: 60}, Status: StudentStatusHighSchool}
	student.Graduate()
}
