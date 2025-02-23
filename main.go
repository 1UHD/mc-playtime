package main

import (
	"fmt"
	"mc-playtime/tools"
)

func main() {
	fmt.Println(tools.Get_time_for_directory(tools.GetPath(tools.Vanilla_mac)))
}
