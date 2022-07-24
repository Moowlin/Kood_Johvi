INSTRUCTION

Write a function QuadA that prints a valid rectangle with a given width of x and height of y.

The function must draw the rectangles as in the examples.

If x and y are positive numbers, the program should print the rectangles as seen in the examples, otherwise, the function should print nothing.

----------------------------------------------------------------------------
USAGE
Here are possible programs to test your function :

```
package main

import "piscine"

func main() {
        piscine.QuadA(5,3)
}
```
-----------------------------------------------------------------------------


QuadA displays a rectangle with "o" character in the corners, and "-" or "|" characters on the sides of the rectangle.
For example, the function QuadA(5,3) will output
```
o---o
|   |
o---o
```
```
QuadB displays a rectangle with "/" or "\" characters in the corners, and "\*" character on the sides of the rectangle.
For example, the QuadB(5,3) function will output

/***\
*   *
\***/
```

QuadC displays a rectangle with the symbol "A" in the upper corners, the symbol "C" in the lower corners, and the symbol "B" on the sides of the rectangle.
For example, the function QuadC(5,3) will output
```
ABBBA
B   B
CBBBC
```

QuadD displays a rectangle, in the left corners of which there is the symbol "A", in the right corners the symbol "C", and on the sides of the rectangle the symbol "B".
For example, the QuadD(5,3) function will output
```
ABBBC
B   B
ABBBC
```

QuadE displays a rectangle, in the first and fourth corners of which there is the character "A", in the second and third corners there is the character "C", and on the sides of the rectangle the character "B" .
For example, the QuadE(5,3) function will output
```
ABBBC
B   B
CBBBA
```

