package main 

import "fmt"

/*
Counting Valleys

An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly  steps, for every step it was noted if it was an uphill, , or a downhill,  step. Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude. We define the following terms:

A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

Example

The hiker first enters a valley  units deep. Then they climb out and up onto a mountain  units high. Finally, the hiker returns to sea level and ends the hike.

Function Description

Complete the countingValleys function in the editor below.

countingValleys has the following parameter(s):

int steps: the number of steps on the hike
string path: a string describing the path
Returns

int: the number of valleys traversed
Input Format

The first line contains an integer , the number of steps in the hike.
The second line contains a single string , of  characters that describe the path.
*/

func countValleys(steps int32, path string) int32 {
	count := 0
	var valleys int32 
	typeH := 's'

	//check constraints
	if len(path) < 2 || len(path) > 1000000{
		return 0
	}

	if steps >= 2 && steps <= 1000000 {
		for _, c := range path {
		    if c == 'D' {
			    count--
		    } else if c == 'U' {
			    count ++
		    } else {
			    return 0
		    }
		    if typeH == 's' {
			    if count < 0 {
				    typeH = 'v'
			    }else if count > 0 {
				    typeH = 'm'
			    }
		    }else if count == 0 {
			    if typeH == 'v' {
				    valleys++
			    }
			    typeH = 's'
		    }
    		}
	}
	return valleys
}

func main() {
	path := "DDUUUUDD"
	v := countValleys(8, path)
	fmt.Println(v)

	//test case 2
	path = "UDDDUDUU"
	v = countValleys(8, path)
	fmt.Println(v)

	//test case 3
	path = "UUUUUUUU"
	v = countValleys(8, path)
	fmt.Println(v)
	
	//test case 4
	path = "DDDDDDDD"
	v = countValleys(8, path)
	fmt.Println(v)

	//test case 5
	path = "UDDUUDDU"
	v = countValleys(8, path)
	fmt.Println(v)

	//test case 6
	path = "U"
	v = countValleys(1, path)
	fmt.Println(v)

}
