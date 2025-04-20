package main

import "fmt"

func main() {
	simple_Map := make(map[int]string)
	simple_Map[11] = "cricket team"
	simple_Map[4] = "badminton plalayer"
	fmt.Println("select map data ", simple_Map)
	s_map := map[int]string{
		12: "kabaddi",
		2:  "single baminton",
	}
	fmt.Println("s map :", s_map)
}
