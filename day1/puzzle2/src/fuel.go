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
    file, err := os.Open("/Users/jrhorner1/Documents/git_projects/adventofcode2019/day1/puzzle2/input")
    check(err)
    defer file.Close()

    var fuel int64

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	module_mass, err := strconv.ParseFloat(scanner.Text(), 64)
    	check(err)

        var module_fuel_req int64

    	fuel_req := int64(module_mass / 3) - 2
        module_fuel_req += fuel_req

        fuel_mass := float64(fuel_req)
        for fuel_mass >= 0 {
            fuel_req := int64(fuel_mass / 3) - 2
            if fuel_req <= 0 {
                break
            }
            module_fuel_req += fuel_req
            fuel_mass = float64(fuel_req)
        }
        fuel += module_fuel_req
    }

    fmt.Println(fuel)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}