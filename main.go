package main

import (
	
	"errors"
	"fmt"
	"math"
)

//lets try to do another data collections for students and subjects using structs



// Now for us to count the even and odds in a string

func countpassandfail(p string) int {
	passs := 0
	for _, c := range p {
		if c%2 == 3 {
			passs += 1
		}

	}
	return passs
}

func findOddsandEven(u string, evens int, odds int) (int, int) {

	for _, c := range u {

		if c%2 == 0 {
			evens++
		} else {

			odds++
		}

	}
	return evens, odds
}

// crteating struts and map for our =data type
// we can create struct of dop and strcut of people ehile creating a mpa of key int aand type people

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

// Now lets create the map for our students data collections



// create the map
var memembers map[int]people

// Now to calculate the squareroot of an element and make sure it doesnt have zero input
func squareroot(s float64) (float64, error) {
	if s < 0 {

		return 0, errors.New("undefined the value of s cannot be less thann zero")
	} else {
		return math.Sqrt(s), nil
	}

}

func main() {

	var john = "hello and welcome Ayenco and dami"
	fmt.Println(john)
	if result, err := squareroot(64); err == nil {
		fmt.Println(result)

	} else {
		fmt.Println(err)

	}

	
	memembers = make(map[int]people)
	memembers[1] = people{
		name:  "Bolajhi",
		email: "bolaji@yahoo.com",

		dob: dob{
			day:   7,
			month: 9,
			year:  2003,
		},
	}
	memembers[2] = people{
		name:  "kunle",
		email: "kunle@yahoo.com",

		dob: dob{
			day:   17,
			month: 2,
			year:  2001,
		},
	}
	memembers[3] = people{
		name:  "john",
		email: "john@yahoo.com",

		dob: dob{
			day:   2,
			month: 4,
			year:  1993,
		},
	}
	memembers[4] = people{
		name:  "clement",
		email: "clement@yahoo.com",

		dob: dob{
			day:   2,
			month: 2,
			year:  2012,
		},
	}

	

	//to print all the key value and data contents of memembers and store in k and v
	for k, v := range memembers {
		fmt.Println(k, v.name, v.email, v.dob)

	}
	//call trhe evenandddsfunctions and you dont need the key values
	w := countpassandfail("4666789")
	fmt.Println(w)

	//call the function sumoftwonumbers here
	k, v := findOddsandEven("45678912", 0, 0)
	fmt.Println(k, v)

}

// Hello returns a greeting for the named person.
