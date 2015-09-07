package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
				line += bigDigits[digit][row]
				if column < len(stringOfDigits)-1 { //we only add that extra space to every number but the last
					line += " "
				}
			} else {
				log.Fatal("invalid whole numnber")
			}
		}

		if row == 0 {
			fmt.Println(strings.Repeat("*", len(line)))
		}

		fmt.Println(line)

		if row == len(bigDigits[0])-1 { //after the very last row of the number is built we build the line, so doenst matter which index we pick to get that length since all of them are the same
			fmt.Println(strings.Repeat("*", len(line)))
		}
	}
	fmt.Printf("chris %d", len(bigDigits[4]))
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
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

// This solution does not use a flag and has some god awful nesting - to run this remove "UglySolution" from the main function only
func mainUglySolution() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0])) /* filepath.Base prints off the program name */
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]                /* get the arg that has the number */
	for row := range bigDigitsUglySolution[1] { /* get the length of each value in array  , which is 9 in this case*/
		line := ""
		for column := range stringOfDigits { /* iterate each byte of the arg i.e3. 123 */
			digit := stringOfDigits[column] - '0' /* So we retrieve the byte value of the command-line string at the given column and subtract the byte value of digit 0 from it to get the number it represents - utf code points where 0 code point is 48 */
			if 0 <= digit && digit <= 9 {
				if row == len(bigDigitsUglySolution[1])-1 || row == 0 { //just use this logic for the first and last row, this is to not print the extra * at the end of the bar
					if column == len(stringOfDigits)-1 {
						if row == len(bigDigitsUglySolution[1])-1 || row == 0 {
							line += bigDigitsUglySolution[digit][row] + " "
						}
					} else {
						line += bigDigitsUglySolution[digit][row] + "*"
					}
				} else {
					line += bigDigitsUglySolution[digit][row] + " "
				}
			} else {
				log.Fatal("invalid whole numnber")
			}
		}

		fmt.Println(line)
	}
}

var bigDigitsUglySolution = [][]string{
	{"*******",
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
		"*******"},
	{"***", " 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111", "***"},
	{"*****", " 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222", "*****"},
	{"*****", " 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 ", "*****"},
	{"******", "   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  ", "******"},
	{"*****", "55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 ", "*****"},
	{"*****", " 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 ", "*****"},
	{"*****", "77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    ", "*****"},
	{"*****", " 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 ", "*****"},
	{"*****", " 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9", "*****"},
}
