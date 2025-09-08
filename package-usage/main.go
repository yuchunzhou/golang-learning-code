package main

import (
	"package-usage/package1"
	"package-usage/package1/package2"
)

func main() {
	package2.GLOBAL_V = 2
	package1.Foo() // 2
}
