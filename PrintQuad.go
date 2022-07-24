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

/*
Quad renders the view rectangle:
Coner_1  Sim_1  Coner_2
Sim_2           Sim_2
Sim_2           Sim_2
Coner_3  Sim_1  Coner_4
*/

func Quad(x, y int, Corner_1, Corner_2, Corner_3, Corner_4, Sim_1, Sim_2 string) {
	if x <= 0 || y <= 0 {
		return
	} else {
		for i := 1; i <= y; i++ {
			if i == 1 {
				Line(x, Corner_1, Sim_1, Corner_2)
			} else if i == y {
				Line(x, Corner_3, Sim_1, Corner_4)
			} else {
				Line(x, Sim_2, " ", Sim_2)
			}
		}
	}
}

// Line prints a lines
func Line(x int, FirstS, MidlS, LastS string) {
	for row := 0; row < x; row++ {
		if row == 0 { // check of Сorner #1
			fmt.Print(FirstS)
		} else if row == x-1 {
			fmt.Printf(LastS) // check of Сorner #2
		} else {
			fmt.Printf(MidlS) // check of another Simbol #1
		}
	}
	fmt.Println("")
}
