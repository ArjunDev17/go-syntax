package main

// talentacquisition
import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	// str := "talentacquisition"

	user := AstrikAndAddres("Arjun")
	fmt.Println("Name", user.Name)

}
func AstrikAndAddres(name string) *User {
	return &User{
		Name: name,
	}
}
