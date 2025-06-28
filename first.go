package main

import (
	"fmt"
)

func desk(row, col int) error {
	if row <= 1 || col <= 1 {
		return fmt.Errorf("входные переменные должны быть больше единицы")
	}
	for ii := range row {
		ps0 := ' '
		ps1 := '#'
		if ii&1 == 1 {
			ps0 = '#'
			ps1 = ' '
		}
		for jj := range col {
			if jj&1 == 0 {
				fmt.Printf("%c", ps0)
			} else {
				fmt.Printf("%c", ps1)
			}
		}
		fmt.Print("\n")
	}
	return nil
}

func main() {
	fmt.Print("Home task2:\n")
	err := desk(3, 7)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
