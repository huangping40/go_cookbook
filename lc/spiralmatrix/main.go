package main

import (
	"fmt"
	"os"
	"time"
)

type dirction int32

const (
	Left dirction = iota
	Right
	Down
	Up
)

type point struct {
	x int
	y int
}

type sm struct {
	upLeft    *point
	upRight   *point
	downLeft  *point
	downRight *point
	nextGo    dirction
	x         int
	y         int
}

func initSM(rows, columns int) *sm {
	s := &sm{
		x: -1,
		y: -1,
	}
	s.upLeft = &point{0, 0}
	s.upRight = &point{0, columns - 1}
	if rows > 1 {
		s.downRight = &point{rows - 1, columns - 1}
		if columns > 1 {
			s.downLeft = &point{rows - 1, 0}
		}
	}

	return s
}

func (s *sm) next() (int, int, error) {
	if s.x < 0 {
		s.nextGo = Right
		s.x = 0
		s.y = 0

		if s.downLeft != nil && s.upLeft.x+1 < s.downLeft.x {
			s.upLeft.x += 1
		} else {
			s.upLeft = nil
		}
		return 0, 0, nil
	}

	var err error
	switch s.nextGo {
	case Right:
		if s.y < s.upRight.y {
			s.y += 1
		} else {
			if s.downRight == nil {
				err = fmt.Errorf("Game is over.")
			} else {
				s.nextGo = Down
				s.x += 1

				if s.downRight != nil && s.upRight.x+1 < s.downRight.x {
					s.upRight.x += 1
					if s.upLeft != nil && s.upRight.y-1 > s.upLeft.y {
						s.upRight.y -= 1
					} else {
						s.upRight = nil
					}
				} else {
					s.upRight = nil
				}
			}
		}

	case Left:
		if s.y > s.downLeft.y {
			s.y -= 1
		} else {
			if s.upLeft == nil {
				err = fmt.Errorf("Game is over.")
			} else {
				s.nextGo = Up
				s.x -= 1

				if s.upLeft != nil && s.downLeft.x-1 > s.upLeft.x {
					s.downLeft.x -= 1
					if s.downRight != nil && s.downLeft.y+1 < s.downRight.y {
						s.downLeft.y += 1
					} else {
						s.downLeft = nil
					}
				} else {
					s.downLeft = nil
				}
			}
		}

	case Down:
		if s.x < s.downRight.x {
			s.x += 1
		} else {
			if s.downLeft == nil {
				err = fmt.Errorf("Game is over.")
			} else {
				s.nextGo = Left
				s.y -= 1

				if s.upRight != nil && s.downRight.x-1 > s.upRight.x {
					s.downRight.x -= 1
					if s.downLeft != nil && s.downRight.y-1 > s.downLeft.y {
						s.downRight.y -= 1
					} else {
						s.downRight = nil
					}
				} else {
					s.downRight = nil
				}
			}
		}

	case Up:
		if s.x > s.upLeft.x {
			s.x -= 1
		} else {
			if s.upRight == nil {
				err = fmt.Errorf("Game is over.")
			} else {
				s.nextGo = Right
				s.y += 1

				if s.downLeft != nil && s.upLeft.x+1 < s.downLeft.x {
					s.upLeft.x += 1
					if s.upRight != nil && s.upLeft.y+1 < s.upRight.y {
						s.upLeft.y += 1
					} else {
						s.upLeft = nil
					}
				} else {
					s.upLeft = nil
				}
			}
		}

	default:
		fmt.Printf("Unreachable code area, nextGo: %d.", s.nextGo)
		os.Exit(-1)
	}
	return s.x, s.y, err
}

func main() {
	var rows int = 7
	var columns int = 11

	arr := make([][]int, 0)
	for i := 0; i < rows; i++ {
		arr0 := make([]int, 0)
		for j := 0; j < columns; j++ {
			arr0 = append(arr0, int(time.Now().Nanosecond()%1000))
		}
		arr = append(arr, arr0)
	}

	for k, _ := range arr {
		for k1, _ := range arr[k] {
			fmt.Printf("%5d,", arr[k][k1])
		}
		fmt.Println("")
	}

	fmt.Printf("line: %d, height: %d. \n ", len(arr), len(arr[0]))

	var total int
	sm := initSM(len(arr), len(arr[0]))
	for {
		x, y, err := sm.next()
		if err != nil {
			break
		}
		fmt.Printf("[%d,%d]", x, y)
		fmt.Printf("%d,", arr[x][y])
		total++
	}
	fmt.Printf("\n  game is over, %d.\n", total)
}
