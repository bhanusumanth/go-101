package data

import "fmt"

var Countries [10]string

// var VARIABLE_NAME COLLECTION_TYPE
var SliceDemo []int

var mapDemo map[int]string

// var <VARIABLE_NAME> map[KEY_TYPE]VALUE_TYPE

func init() {
	fmt.Println(" first Initiated called in collections")
	Countries[0] = "Earth"
	Countries[1] = "Jupiter"
	fmt.Println("Array Printing")
	fmt.Println(Countries, "Capacity : ", cap(Countries), " Length:", len(Countries))
	fmt.Println("Printing default Values for Slices, Map", SliceDemo, "\n", mapDemo)
}

func init() {
	fmt.Println("second init called in collections file")
	Countries[0] = "Sun"
}
