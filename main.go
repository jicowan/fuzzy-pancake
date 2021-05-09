package fuzzy_pancake

import "fmt"

func main() {
	fmt.Println(favs())
}

func favs() []string {
	xs := []string{"Now is the time", "for all good men", "to come to the aid", "of their country"}
	return xs
}
