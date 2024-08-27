package saying

import (
	"fmt"
	"testing"
)

func TEstGreet(t *testing.T) {
	// fmt.Println("Enter name for testing")
	value := Greet("Ram")
	if value != "welcome my dear :Ram" {
		t.Error("got", value, "and expected :", "welcome my dear :Ram")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Ram"))
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Ram")
	}
}

//sudo apt install golang-golang-x-tools
