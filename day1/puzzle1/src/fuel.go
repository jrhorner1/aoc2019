package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
	// change to request the input file from user
    file, err := os.Open("/Users/jrhorner1/Documents/git_projects/adventofcode2019/day1/puzzle1/input")
    check(err)
    defer file.Close()

    var fuel int64

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	mass, err := strconv.ParseFloat(scanner.Text(), 64)
    	check(err)
    	mass_rd := int64(mass / 3)
    	fuel += mass_rd - 2
    }

    fmt.Println(fuel)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}