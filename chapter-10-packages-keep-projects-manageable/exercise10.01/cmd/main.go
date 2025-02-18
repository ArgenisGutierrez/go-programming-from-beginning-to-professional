// Package main creating a package to calculate areas of various shapes
package main

import "gitgub.com/exercise10.01/pkg/shape"

func main() {
	t := shape.Square{Side: 10}
	shape.PrintShapeDetails(t)
}
