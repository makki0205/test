package test

type User struct {
	Name       string
	NameLength int
}
type Users []User

func Lesson1(names []string) Users {
	panic("hoge")
}
func Lesson2(n int) {
	panic("hoge")
}

type Coin struct {
	Thousand    int
	FiveHundred int
	hundred     int
	Fifty       int
	Ten         int
	Five        int
}

func Lesson3(n int) Coin {
	panic("hoge")
}
