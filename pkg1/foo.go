package pkg1

import (
	"fmt"

	"github.com/bep/bugrepogoplsgenericfuncref/pkg2"
)

func Foo() {
	b := pkg2.Bar[string]{
		BarFunc: func() string {
			return "Hello from BarFunc"
		},
	}

	fmt.Println(b)
}
