package main

import "fmt"

/*
Jumping on the Clouds

Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus  or . She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from her starting postion to the last cloud. It is always possible to win the game.

For each game, Emma will get an array of clouds numbered  if they are safe or  if they must be avoided. For example,  indexed from . The number on each cloud is its index in the list so she must avoid the clouds at indexes  and . She could follow the following two paths:  or . The first path takes  jumps while the second takes.
*/

func jumpingOnClouds(c []int32) int32 {
	var j int32 = 0
	pos := 0
	// constraints
	if len(c) < 2 || len(c) > 100{
		return 0
	}
	if  c[0] != 0 || c[pos] != 0 {
		return 0
	}

	for pos < len(c)-1 {
		fmt.Printf("pos: %v\n", pos)
	    if pos+2 == len(c) || c[pos+2] == 1 {
                pos++
	    } else {
	        pos += 2
	    }
	    j++
        }
        return j
}

func main() {
	c := []int32{0,1,0}
	fmt.Println(jumpingOnClouds(c))
	c = []int32{1,0,0}
	fmt.Println(jumpingOnClouds(c))
	c = []int32{0,1,0,0,0,1,0}
	fmt.Println(jumpingOnClouds(c))
	c = []int32{0,0,1,0,0,1,0}
	fmt.Println(jumpingOnClouds(c))
	c = []int32{0,0,1,0,1,0,0}
	fmt.Println(jumpingOnClouds(c))
}

