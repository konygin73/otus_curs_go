package main

import (
	"fmt"
	"os"
	"strconv"
)

func desk(row int, col int) error {
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
	if len(os.Args) != 3 {
		fmt.Printf("запуск с параметрами: количество строк и столбцов: %d %s",
			len(os.Args), os.Args[0])
		return
	}
	row := os.Args[1]
	col := os.Args[2]
	rowNum, err := strconv.Atoi(row)
	if err != nil {
		fmt.Printf("Error converting first string to int: %s", err)
		return
	}
	colNum, err := strconv.Atoi(col)
	if err != nil {
		fmt.Printf("Error converting second string to int: %s", err)
		return
	}
	fmt.Printf("row: %d col: %d\n", rowNum, colNum)
	err = desk(rowNum, colNum)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
