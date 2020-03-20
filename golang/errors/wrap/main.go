package main

import (
	"errors"
	"fmt"
)

var e1 = errors.New("식별할 수 없는 값이 있습니다.")

func x1() error {
	return fmt.Errorf("adding more context: %w", e1)
}

type myError struct {
	err  string
	more string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%s: %s", e.more, e.err)
}
func x2() error {
	return fmt.Errorf("adding more context: %w", &myError{"error", "more"})
}

func main() {
	isE := x1()
	if errors.Is(isE, e1) { // Magical it works
		// handle gracefully
		fmt.Println(isE)
		fmt.Println(e1)
	}
	asE := x2()
	var e2 *myError
	if ok := errors.As(asE, &e2); ok {
		fmt.Println(asE)
		fmt.Println(e2)
	}
}
