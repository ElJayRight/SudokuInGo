package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"


	"image/color"
	"log"
	"strconv"
)
var (
	mplusNormalFont font.Face
	fontSize        int = 36
	grid = [81]int{0, 7, 2, 0, 3, 1, 9, 0, 0, 8, 0, 0, 0, 0, 9, 0, 5, 7, 5, 0, 0, 8, 2, 0, 0, 1, 0, 2, 0, 4, 0, 0, 3, 0, 9, 0, 3, 9, 6, 2, 1, 0, 0, 4, 5, 1, 0, 0, 0, 0, 6, 0, 3, 2, 0, 0, 3, 0, 0, 0, 0, 2, 0, 4, 0, 5, 9, 6, 0, 0, 0, 8, 9, 0, 0, 0, 0, 4, 0, 0, 1}

)
type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	rect := ebiten.NewImage(70, 70)
	rect.Fill(color.Black)

	rect2 := ebiten.NewImage(68, 68)
	rect2.Fill(color.White)
	for i := 0; i != 81; i++ {
		op := &ebiten.DrawImageOptions{}
		op2 := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i%9*70), float64(i/9*70))
		screen.DrawImage(rect, op)
		op2.GeoM.Translate(float64(1+i%9*70), float64(1+i/9*70))
		screen.DrawImage(rect2, op2)
		if grid[i]!=0{
			row := i%9
			col := i/9
			text.Draw(screen, fmt.Sprintf(strconv.Itoa(grid[i])), mplusNormalFont,25+row*70,45+col*70,color.Black)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 630, 630
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Sudoku")
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	//grid := [81]int{0, 7, 2, 0, 3, 1, 9, 0, 0, 8, 0, 0, 0, 0, 9, 0, 5, 7, 5, 0, 0, 8, 2, 0, 0, 1, 0, 2, 0, 4, 0, 0, 3, 0, 9, 0, 3, 9, 6, 2, 1, 0, 0, 4, 5, 1, 0, 0, 0, 0, 6, 0, 3, 2, 0, 0, 3, 0, 0, 0, 0, 2, 0, 4, 0, 5, 9, 6, 0, 0, 0, 8, 9, 0, 0, 0, 0, 4, 0, 0, 1}
	//var move [3]int

	for 2 > 1 {
		render_board(grid)
		fmt.Println("Would you like to\n1. Check the board for errors.\n2. Enter a number.\n3. Exit")
		var out int
		fmt.Scanln(&out)
		if out == 3 {
			break
		} else if out == 1 {
			if generate_and_check_chunks(grid) {
				fmt.Println("There is an error :(")
			} else {
				fmt.Println("There are no errors.")
			}
			fmt.Print("Press Enter to continue.")
			fmt.Scanln()
		} else if out == 2 {
			move = user_input()
			grid = update_board(grid, move)
		}
	}
}
*/

func check(row [9]int) bool {
	set := map[int]bool{}
	c := 0
	v := 0
	for i := 0; i != 9; i++ {
		value := row[i]
		set[value] = true
		if value == 0 {
			c += 1
			v = 1
		}
	}
	if len(set)+c-v == len(row) {
		return false
	} else {
		return true
	}
}

func generate_and_check_chunks(grid [81]int) bool {
	for i := 0; i != 81; i += 9 {
		var a [9]int
		for j := 0; j != 9; j++ {
			a[j] = grid[i+j]
		}
		if check(a) == true {
			return true
		}
	}
	for i := 0; i != 9; i++ {
		var a [9]int
		for j := 0; j != 81; j += 9 {
			a[j/9] = grid[i+j]
		}
		if check(a) == true {
			return true
		}
	}
	for i := 0; i != 3; i++ {
		for j := 0; j != 3; j++ {
			var a [9]int
			for k := 0; k != 27; k += 9 {
				a[k/3] = i*27 + j*3 + k
				a[k/3+1] = i*27 + j*3 + k + 1
				a[k/3+2] = i*27 + j*3 + k + 2
			}
			if check(a) == true {
				return true
			}

		}
	}
	return false
}

func select_location(x float32, y float32) int {
	width_of_cell := 7
	x /= float32(width_of_cell)
	y /= float32(width_of_cell)
	v := int(x)*9 + int(y)
	return v
}

func write_number(grid [81]int, position int, number int) bool {
	if grid[position] == 0 {
		grid[position] = number
		return true
	} else {
		return false
	}
}

/*
func render_board(grid [81]int) {
	for i := 0; i != 9; i++ {
		out := "| "
		for j := 0; j != 9; j++ {
			out += strconv.Itoa(grid[i*9+j]) + " "
			if j%3 == 2 {
				out += "| "
			}
		}
		if i%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
		fmt.Println(out)
	}
	fmt.Println("+-------+-------+-------+")
}
*/
/*
func user_input() [3]int {
	fmt.Println("Enter x y and number. Example 1,1,9 ")
	var out string
	fmt.Scanln(&out)
	if len(out) != 5 {
		fmt.Println("Why couldn't you just enter three numbers?")
		return [3]int{0, 0, 0}
	}
	var x int
	var y int
	var number int
	x, _ = strconv.Atoi(string(out[0]))
	y, _ = strconv.Atoi(string(out[2]))
	number, _ = strconv.Atoi(string(out[4]))
	if x > 0 && x < 10 && y > 0 && y < 10 && 0 < number && number < 10 {
		return [3]int{x - 1, y - 1, number}
	} else {
		fmt.Println("Numbers have to be between 1 and 9 (inclusively).")
	}
	return [3]int{0, 0, 0}
}
*/
func update_board(grid [81]int, move [3]int) [81]int {
	location := move[0] + move[1]*9
	grid[location] = move[2]
	return grid
}

func convert_input(pos [2]int) [2]int {
	x := pos[0] / 600 / 9
	y := pos[1] / 600 / 9
	return [2]int{x, y}
}
