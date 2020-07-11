package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	// Here 'name' is the result of the access
	// Also, as "Un" doesn't exist in the map, 'name' receives
	// the 0 value of the strings that is the empty string
	// 'ok' is the boolean value indicating if the operation was successful
	name, ok := elements["Un"]
	fmt.Println("'"+name+"'", ok)

	// You can put statements before evaluate the if expression
	if name, ok := elements["Un"]; !ok {
		fmt.Println("'"+name+"'", ok)
	}
}
