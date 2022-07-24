package piscine

import "fmt"

/*
QuadA displays a rectangle for example, the function QuadA(5,3) will output
o---o
|   |
o---o
*/
func QuadA(x, y int) {
	Quad(x, y, "o", "o", "o", "o", "-", "|")
}

/*
QuadB displays a rectangle for example, the QuadB(5,3) function will output
/*** \
*    *
\*** /
*/
func QuadB(x, y int) {
	Quad(x, y, "/", "\\", "\\", "/", "*", "*")
}

/*
QuadC displays a rectangle for example, the function QuadC(5,3) will output
ABBBA
B   B
CBBBC
*/

func QuadC(x, y int) {
	Quad(x, y, "A", "A", "C", "C", "B", "B")
}

/*
QuadD displays a rectangle for example, the QuadD(5,3) function will output
ABBBC
B   B
ABBBC
*/
func QuadD(x, y int) {
	Quad(x, y, "A", "C", "A", "C", "B", "B")
}

/*
QuadE displays a rectangle for example, the QuadE(5,3) function will output
ABBBC
B   B
CBBBA
*/
func QuadE(x, y int) {
	Quad(x, y, "A", "C", "C", "A", "B", "B")
}

func Quad(x, y int, Corner_1, Corner_2, Corner_3, Corner_4, Sim_1, Sim_2 string) {
	if x <= 0 || y <= 0 {
		return
	} else {
		for i := 1; i <= y; i++ {
			if i == 1 {
				FirstLine(x, Corner_1, Corner_2, Sim_1)
			} else if i == y {
				LastLine(x, Corner_3, Corner_4, Sim_1)
			} else {
				MiddleLine(x, Sim_2)
			}
		}
	}
}

// FirstOrLastString prints the first or last lines
func FirstLine(x int, Corner_1, Corner_2, Sim_1 string) {
	for row := 0; row < x; row++ {
		if row == 0 { // check of Сorner #1
			fmt.Print(Corner_1)
		} else if row == x-1 {
			fmt.Printf(Corner_2) // check of Сorner #2
		} else {
			fmt.Printf(Sim_1) // check of another Simbol #1
		}
	}
	fmt.Println("")
}

// MiddleLines prints the sub lines
func MiddleLine(x int, Sim_2 string) {
	for row := 0; row < x; row++ {
		if row == 0 || row == x-1 {
			fmt.Printf(Sim_2)
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Println("")
}

// LastString prints the first or last lines
func LastLine(x int, Corner_3, Corner_4, Sim_1 string) {
	for row := 0; row < x; row++ {
		if row == 0 {
			fmt.Printf(Corner_3) // check of Сorner #3
		} else if row == x-1 {
			fmt.Printf(Corner_4) // check of Сorner #4
		} else {
			fmt.Printf(Sim_1) // check of another Simbol #1
		}
	}
	fmt.Println("")
}
