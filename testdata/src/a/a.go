package a

func sum1(a, b int) int {
	return a + b
}

func sum2(a int, b int) int {
	return a + b
}

func sum3(a ...int) int { // want "lines are recommended"
	val := 0

	for _, v := range a {
		val += v
	}

	return val
}

func sum4() {}

func sum5(x, y int) int {
	return x + y
}

type User struct {
	ID   int
	Name string
}

func sum6(a, b, c, d, e int) int { // want "args are recommended"
	return a + b + c + d + e
}

func (u *User) SetName(name string) {
	u.Name = name
}
