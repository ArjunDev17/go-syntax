package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Ram is God"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Ram is God" {
		t.Error("got", s, "want", "Ram is God")
	}
}

func TestJoin(t *testing.T) {
	s := "Ram is God"
	xs := strings.Split(s, " ")
	s = Join(xs)
	fmt.Println("s :", s)
	if s != "Ram is God" {
		t.Error("got", s, "want", "Ram is God")
	}
}
func ExampleCat() {
	s := "Ram is God"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output: Ram is God
}

func ExampleJoin() {
	s := "Ram is God"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output: Ram is God
}

const s = "Dear Arjun,As we're approaching the final stages of the recruiting process, we need some information from you to work on the offer release process.Request you to fill up this form at the earliest. Below is the form link."

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
