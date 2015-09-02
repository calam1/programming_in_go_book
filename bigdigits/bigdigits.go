package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0])) /* filepath.Base prints off the program name */
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]    /* get the arg that has the number */
	for row := range bigDigits[1] { /* get the length of each value in array  , which is 7 in this case*/
		line := ""
		for column := range stringOfDigits { /* iterate each byte of the arg i.e3. 123 */
			digit := stringOfDigits[column] - '0' /* So we retrieve the byte value of the command-line string at the given column and subtract the byte value of digit 0 from it to get the number it represents - utf code points where 0 code point is 48 */
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole numnber")
			}
		}

		fmt.Println(line)
	}
}

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
