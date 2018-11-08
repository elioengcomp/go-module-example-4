package example

import (
	"fmt"
	"github.com/elioengcomp/go-module-example-2/anotherpackage"
)

const version = "no-version"

func Exec() string {

	return fmt.Sprintf("This is go-module-example-4 %s Using go-module-example-2 so we have a circular dependency: %s", version, anotherpackage.Exec())
}
