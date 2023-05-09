package quadtree

import "testing"

func TestQuadtree(t *testing.T) {
	// Initialize a new quad tree
	topLeft := Point{X: -10, Y: -10}
	bottomRight := Point{X: 10, Y: 10}
	q := NewQuadtree(topLeft, bottomRight)

	// Insert some nodes
	node1 := Node{Point: Point{X: 5, Y: 5}, Data: "data1"}
	node2 := Node{Point: Point{X: 7, Y: 7}, Data: "data2"}
	node3 := Node{Point: Point{X: 9, Y: 9}, Data: "data3"}

	// Test inserting nodes
	if !q.Insert(node1) {
		t.Error("Failed to insert node1")
	}

	if !q.Insert(node2) {
		t.Error("Failed to insert node2")
	}

	if !q.Insert(node3) {
		t.Error("Failed to insert node3")
	}

	// Test searching for inserted nodes
	data1, found1 := q.Search(node1.Point)
	if !found1 {
		t.Error("Failed to find node1")
	}
	if data1 != node1.Data {
		t.Errorf("Unexpected data for node1: %v", data1)
	}

	data2, found2 := q.Search(node2.Point)
	if !found2 {
		t.Error("Failed to find node2")
	}
	if data2 != node2.Data {
		t.Errorf("Unexpected data for node2: %v", data2)
	}

	data3, found3 := q.Search(node3.Point)
	if !found3 {
		t.Error("Failed to find node3")
	}
	if data3 != node3.Data {
		t.Errorf("Unexpected data for node3: %v", data3)
	}

	// Test searching for nodes that were not inserted
	data4, found4 := q.Search(Point{X: -5, Y: -5})
	if found4 {
		t.Error("Found node4, which should not exist")
	}
	if data4 != nil {
		t.Errorf("Unexpected data for node4: %v", data4)
	}

	// Test inserting a node out of bounds
	outOfBoundsNode := Node{Point: Point{X: 20, Y: 20}, Data: "outOfBounds"}
	if q.Insert(outOfBoundsNode) {
		t.Error("Inserted an out of bounds node")
	}
}
