package main

import (
	"fmt"

	"github.com/magdyamr542/point-quad-tree/quadtree"
)

func main() {
	topLeft := quadtree.Point{X: 0, Y: 0}
	bottomRight := quadtree.Point{X: 8, Y: 8}
	tree := quadtree.NewQuadtree(topLeft, bottomRight)
	if inserted := tree.Insert(quadtree.Node{Point: quadtree.Point{X: 1, Y: 1}, Data: "first"}); !inserted {
		panic("couldn't insert node")
	}

	if inserted := tree.Insert(quadtree.Node{Point: quadtree.Point{X: 2, Y: 5}, Data: "second"}); !inserted {
		panic("couldn't insert node")
	}

	if inserted := tree.Insert(quadtree.Node{Point: quadtree.Point{X: 7, Y: 6}, Data: "third"}); !inserted {
		panic("couldn't insert node")
	}

	searchFor := func(point quadtree.Point) {
		data, exists := tree.Search(point)
		if !exists {
			fmt.Printf("Data at %+v doesn't exist\n", point)
		} else {
			fmt.Printf("Data at %+v: %v\n", point, data)
		}
	}

	searchFor(quadtree.Point{X: 1, Y: 1})
	searchFor(quadtree.Point{X: 1, Y: 2})
	searchFor(quadtree.Point{X: 7, Y: 6})
}
