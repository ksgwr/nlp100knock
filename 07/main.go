package main

import "fmt"

func main() {
	fmt.Println(Template(12, "気温", 22.4))
}

func Template(x int, y string, z float32) string {
	return fmt.Sprintf("%d時の%sは%.01f", x, y, z)
}