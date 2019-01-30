// magic
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

/**
 * Prints the license of this library
 */
func license() string {
	return `MagicSquare v1.0.0b (20170316)
MagicSquare Copyright(C) 2014, 2015, 2016, 2017 Anoop Manakkalath
<anoop.manakkalath@gmail.com>

This program is free software: you can redistribute it and/or modify it
under the terms of the GNU Lesser General Public License version 3 as
published by the Free Software Foundation.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
See the GNU Lesser General Public License for more details.
You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.`
}

/**
 * Gets new line character of OS (Windows or GNU/ Linux)
 */
func getNewLineChar() string {
	osname := runtime.GOOS
	if osname == "windows" {
		return "\r\n"
	}
	return "\n"
}

/**
 * This generates the odd magic square
 * @return int[][] magic square
 */
func createOddMagicSquare(order int) [][]int {
	magicSquare := initializeMagicSquare(order)
	square := order * order
	HALF := order / 2
	number := 1
	row := 0
	col := HALF
	magicSquare[row][col] = number
	for number < square {

		number++
		row--
		col++
		if row < 0 && col >= order {
			row += 2
			col--
		}
		if row < 0 {
			row = order - 1
		}
		if order <= col {
			col = 0
		}
		if magicSquare[row][col] != 0 {
			row += 2
			col--
		}
		if magicSquare[row][col] == 0 {
			magicSquare[row][col] = number
		}
	}
	return magicSquare
}

/**
 * This generates the doubly even magic square
 * @return int[][] magic square
 */
func createDoublyEvenMagicSquare(order int) [][]int {
	magicSquare := initializeMagicSquare(order)
	for i := 0; i < order; i++ {
		if i%4 == 0 || i%4 == 3 {
			for j := 0; j < order; j++ {
				if j%4 == 0 || j%4 == 3 {
					magicSquare[i][j] = i*order + j + 1
				} else {
					magicSquare[i][j] = (order-i)*order - j
				}
			}
		} else {
			for j := 0; j < order; j++ {
				if j%4 == 1 || j%4 == 2 {
					magicSquare[i][j] = i*order + j + 1
				} else {
					magicSquare[i][j] = (order-i)*order - j
				}
			}
		}
	}
	return magicSquare
}

/**
 * This generates the singly even magic square
 * @return int[][] magic square
 */
func createSinglyEvenMagicSquare(order int) [][]int {
	magicSquare := initializeMagicSquare(order)
	fourthsQuotient := (order - 2) / 4
	half := order / 2
	quarter := half / 2
	quarterSquare := half * half
	number := 1
	row := 0
	col := quarter
	stored := 0
	tmpRow := 0
	magicSquare[row][col] = number
	for number != quarterSquare {
		number = number + 1
		row--
		col++
		if row < 0 && col >= half {
			row += 2
			col--
		}
		if row < 0 {
			row = half - 1
		}
		if col >= half {
			col = 0
		}
		if magicSquare[row][col] != 0 {
			row += 2
			col--
		}
		if magicSquare[row][col] == 0 {
			magicSquare[row][col] = number
		}
	}
	number = quarterSquare + 1
	row = half
	col = half + quarter
	magicSquare[row][col] = number
	for number != quarterSquare*2 {
		number = number + 1
		row--
		col++
		if row < half && col >= half*2 {
			row += 2
			col--
		}
		if row < half {
			row = 2*half - 1
		}
		if col >= half*2 {
			col = half
		}
		if magicSquare[row][col] != 0 {
			row += 2
			col--
		}
		if magicSquare[row][col] == 0 {
			magicSquare[row][col] = number
		}
	}
	number = quarterSquare*2 + 1
	row = 0
	col = half + quarter
	magicSquare[row][col] = number
	for number != quarterSquare*3 {
		number = number + 1
		row--
		col++
		if row < 0 && col >= half*2 {
			row += 2
			col--
		}
		if row < 0 {
			row = half - 1
		}
		if col >= half*2 {
			col = half
		}
		if magicSquare[row][col] != 0 {
			row += 2
			col--
		}
		if magicSquare[row][col] == 0 {
			magicSquare[row][col] = number
		}
	}
	number = quarterSquare*3 + 1
	row = half
	col = quarter
	magicSquare[row][col] = number
	for number != quarterSquare*4 {
		number = number + 1
		row--
		col++
		if row < half && col >= half {
			row += 2
			col--
		}
		if row < half {
			row = 2*half - 1
		}
		if col >= half {
			col = 0
		}
		if magicSquare[row][col] != 0 {
			row += 2
			col--
		}
		if magicSquare[row][col] == 0 {
			magicSquare[row][col] = number
		}
	}
	for i := 0; i < fourthsQuotient; i++ {
		tmpRow = 0
		for tmpRow < half {
			if i == 0 && tmpRow == fourthsQuotient {
				tmpRow++
				continue
			}
			stored = magicSquare[tmpRow][i]
			magicSquare[tmpRow][i] = magicSquare[tmpRow+half][i]
			magicSquare[tmpRow+half][i] = stored
			tmpRow++
		}
	}
	stored = magicSquare[fourthsQuotient][fourthsQuotient]
	magicSquare[fourthsQuotient][fourthsQuotient] = magicSquare[fourthsQuotient+half][fourthsQuotient]
	magicSquare[fourthsQuotient+half][fourthsQuotient] = stored
	for i := order; i > (order - fourthsQuotient + 1); i-- {
		stored = 0
		tmpRow = 0
		for tmpRow < half {
			stored = magicSquare[tmpRow][i-1]
			magicSquare[tmpRow][i-1] = magicSquare[tmpRow+half][i-1]
			magicSquare[tmpRow+half][i-1] = stored
			tmpRow++
		}
	}
	return magicSquare
}

/**
* Initailizes a dummy magic square with the given order
 */
func initializeMagicSquare(order int) [][]int {
	magicSquare := make([][]int, order)
	for i := range magicSquare {
		magicSquare[i] = make([]int, order)
	}
	return magicSquare
}

/**
 * Error Message
 */
func getErrMessage() string {
	errMsg := "Please give an order which is a number between 3 and 101."
	return errMsg
}

/**
 * This gives the magic sum of the magic square.
 * @return int magicSum
 */
func getMagicSum(order int) int {
	if order < 3 || order > 101 {
		fmt.Println(getErrMessage())
		return 0
	}
	magicSum := order * (order*order + 1) / 2
	return magicSum
}

/**
 * This method characterizes the magic square.
 * @param int order of the magic square
 */
func discriminateIt(order int) int {
	/**
	 * The type of the magic square
	 */
	var kind int
	residue := order % 4
	if order < 3 || order > 101 {
		fmt.Println(getErrMessage())
		return 0
	}
	switch residue {
	case 0:
		kind = 4
		break
	case 2:
		kind = 2
		break
	default:
		kind = 1
		break
	}
	return kind
}

/**
 * This calls the appropriate methods to generate the magic square.
 * @return magic square as an integer array
 */
func generateMagicSquare(order int) [][]int {
	if order < 3 || order > 101 {
		fmt.Println(getErrMessage())
		return [][]int{}
	}
	magicSquare := make([][]int, order)
	for i := range magicSquare {
		magicSquare[i] = make([]int, order)
	}
	kind := discriminateIt(order)
	switch kind {
	case 1:
		copy(magicSquare, createOddMagicSquare(order))
		break
	case 4:
		copy(magicSquare, createDoublyEvenMagicSquare(order))
		break
	case 2:
		copy(magicSquare, createSinglyEvenMagicSquare(order))
		break
	}
	return magicSquare
}

/**
 * This prints the magic square.
 */
func printMagicSquare(magicSquare [][]int) {
	for row := range magicSquare {
		for col := range magicSquare {
			fmt.Print(magicSquare[row][col], "\t")
		}
		fmt.Println()
	}
}

/**
 * This prints the magic sum.
 */
func printMagicSum(magicSum int) {
	fmt.Println("The magic sum is: ", magicSum)
}

/**
 * This writes the magic square to a file.
 */
func writeMagicSquare(magicSquare [][]int) bool {
	order := len(magicSquare)
	orderStr := strconv.Itoa(order)
	newLine := getNewLineChar()
	pwd, err := os.Getwd()
	if err != nil {
		return false
	}
	file, err := os.Create(pwd + "/" + orderStr + ".txt")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	writeStr := "The " + orderStr + "*" + orderStr + " magic square is:" +
		newLine + newLine
	w.WriteString(writeStr)
	w.Flush()
	for row := range magicSquare {
		writeStr = ""
		for col := range magicSquare {
			writeStr += strconv.Itoa(magicSquare[row][col]) + "\t"
		}
		writeStr += newLine
		w.WriteString(writeStr)
		w.Flush()
	}
	writeStr = newLine + "The magic sum is: " + strconv.Itoa(getMagicSum(order)) +
		newLine
	w.WriteString(writeStr)
	w.Flush()
	return true
}

// Main function
func main() {
	var orderStr string
	var save string
	var bypass bool
	orderStrPtr := flag.String("o", "0", "Enter a number between 3 and 101")
	savePtr := flag.String("s", "n", "Give 'y' option to save the magic square")
	flag.Parse()
	orderStr = *orderStrPtr
	save = *savePtr
	fmt.Println(license())
	fmt.Println()
	fmt.Println("Usage:", "magic -o <order> -s <y/n>")
	//The order of the magic square (int).
	order, err := strconv.Atoi(orderStr)
	if err != nil {
		fmt.Println(getErrMessage())
		bypass = true
	}
	if !bypass && (order < 3 || order > 101) {
		fmt.Println(getErrMessage())
		bypass = true
	}
	if !bypass {
		magicSquare := generateMagicSquare(order)
		if len(magicSquare) > 0 {
			printMagicSquare(magicSquare)
			magicSum := getMagicSum(order)
			printMagicSum(magicSum)
			if save == "y" || save == "Y" {
				write := writeMagicSquare(magicSquare)
				if write {
					fmt.Println("The magicsquare has been written to disk")
				}
			}
		}
	}
	var input string
	fmt.Print(getNewLineChar(), "Hit 'Enter' to continue...")
	fmt.Scanf("%s\n", &input)
}
