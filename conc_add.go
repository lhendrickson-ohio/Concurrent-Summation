// Written by Liam Hendrickson
// 1/20/2020
// This programs utilizes golang's concurrency features to sum
// a list of 5000 integers from a file that range from 1 to 5000


package main

import( "io/ioutil"
	"fmt"
	"strings"
	"strconv"
)



func main(){

	input,err := ioutil.ReadFile("sample_in.txt")
	if err != nil{
		panic(err)
	}
	line := strings.Split(string(input), "\n")	//splits lines at newline  char

	data := make([]int, 0, len(line)) 		//makes an empty slice the number of lines in the file


	for _, r:= range line{			//itterates through the entire file
		if len(r) == 0{			//covers empty lines, fixes a parsing error with Atoi
			continue
		}
		n, err := strconv.Atoi(r)		//converts the string into an integer
		if err != nil{
			fmt.Print(err)
		}
		data = append(data, n)			//adds the new integer to the slice
	}


	c:= make(chan int, 5)				//creates a channel with a capacity for 5
	var sum1,sum2,sum3,sum4,sum5 int		//the sums of each fifth of the array

	go add(data, 0, 1000, c)
	go add(data, 1001, 2000, c)
	go add(data, 2001, 3000, c)
	go add(data, 3001, 4000, c)
	go add(data, 4001, 5000, c)

	sum1= <- c
	sum2= <- c
	sum3= <- c
	sum4= <- c
	sum5= <- c

	fmt.Print("\n", sum1+sum2+sum3+sum4+sum5, "\n")



}

func add(data []int, start int, end int, c chan int) {
	var added int
	added = 0

	for i:=start; i<end; i++{
		added = added + data[i]

	}
	c <- added

}

