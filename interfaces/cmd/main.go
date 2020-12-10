package main

import (
	"fmt"
	"localrepo/denser/metals"
)

func main() {

	gold := metals.Metal{Mass: 478, Volume: 24}
	silver := metals.Metal{Mass: 100, Volume: 10}

	result := metals.IsDenser(&gold, &silver)

	if result {
		fmt.Println("gold has higer density than silver")
	} else {
		fmt.Println("Silver has higher density than gold")
	}

}
