package main

import "strconv"

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

// レシーバーは構造体のポインタを受け取る
// そうしないと値渡しになってしまい、メソッド内での変更が反映されない
// どういうことかというと、この関数内でStatusを変更しても、関数を抜けた後に元の値に戻ってしまう
// ポインタを受け取ることで、関数内での変更が元の構造体に反映される
// なぜなら、ポインタを受け取ることで、構造体のアドレスを受け取るから
func (s *Student) Graduate() {
	switch s.Status {
	case StudentStatusHighSchool:
		s.Status = StudentStatusUniversity
	case StudentStatusUniversity:
		s.Status = StatusStatusGraduate
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
	var numString string = "10x0"
	i, err := strconv.Atoi(numString)
	if err != nil {
		println(err)
		return
	}
	println(i)
}
