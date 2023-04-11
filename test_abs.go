package main

import ("testing"
	"math")

func GPT(){
	
}

func TestGPT(t *testing.T) {
    got := math.Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

func main() {
	x := -25.0
	y := math.Abs(x)
	fmt.Print(y)
}
