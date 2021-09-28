package test

type User struct {
	Name       string
	NameLength int
}
type Users []User

func Lesson1(names []string) Users {
	res := make(Users, len(names))
	for i, name := range names {
		res[i] = User{
			Name:       name,
			NameLength: len(name),
		}
	}
	return res
}
func Lesson2(n int) {

}

type Coin struct {
	Thousand    int
	FiveHundred int
	hundred     int
	Fifty       int
	Ten         int
	Five        int
	One         int
}

func Lesson3(n int) Coin {

}
