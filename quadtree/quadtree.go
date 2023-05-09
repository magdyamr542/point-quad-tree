package quadtree

// Point is a point in the 2d space
type Point struct {
	X int
	Y int
}

// Node is a quad tree node that has data and a location represented by a point
type Node struct {
	Point Point
	Data  any
}

// Quadtree is a point quad tree
type Quadtree struct {
	Node *Node

	// The borders of this quad tree.
	TopLeftP     Point
	BottomRightP Point

	// The four regions that are split with the current quad tree are represented with theses pointer.
	// If a pointer is nil, it means the region is empty.
	TopLeftTree     *Quadtree
	TopRightTree    *Quadtree
	BottomLeftTree  *Quadtree
	BottomRightTree *Quadtree
}

func NewQuadtree(topLeft, bottomRight Point) Quadtree {
	q := Quadtree{TopLeftP: topLeft, BottomRightP: bottomRight}
	return q
}

func (q *Quadtree) inbound(node Node) bool {
	return node.Point.X >= q.TopLeftP.X && node.Point.X <= q.BottomRightP.X &&
		node.Point.Y >= q.TopLeftP.Y && node.Point.Y <= q.BottomRightP.Y
}

// Search returns the date of the node that is associated with the point. It returns false if such node can't be found.
func (q *Quadtree) Search(point Point) (any, bool) {

	// Node cannot be inserted to this tree.
	if !q.inbound(Node{Point: point}) {
		return nil, false
	}

	topLeftX := q.TopLeftP.X
	topLeftY := q.TopLeftP.Y
	bottomRightX := q.BottomRightP.X
	bottomRightY := q.BottomRightP.Y

	inLeft := (topLeftX+bottomRightX)/2 >= point.X
	inTop := (topLeftY+bottomRightY)/2 >= point.Y

	// We are at a quad of unit area
	// We cannot subdivide this quad further.
	if (topLeftX-bottomRightX) <= 1 && (bottomRightY-topLeftY) <= 1 {
		if q.Node == nil {
			return nil, false
		}
		return q.Node.Data, true
	}

	// The node is in the left part of the tree
	if inLeft {

		// The node is in the top left part of the tree
		if inTop {
			if q.TopLeftTree != nil {
				return q.TopLeftTree.Search(point)
			}

		} else {
			// The node is in the bottom left part of the tree
			if q.BottomLeftTree != nil {
				return q.BottomLeftTree.Search(point)
			}

		}

	} else {
		// The node is in the right part of the tree

		// The node is in the top right part of the tree
		if inTop {
			if q.TopRightTree != nil {
				return q.TopRightTree.Search(point)
			}

		} else {
			// The node is in the bottom right part of the tree
			if q.BottomRightTree != nil {
				return q.BottomRightTree.Search(point)
			}

		}

	}

	return nil, false
}

// Insert inserts node to the tree. It returns true if the node was inserted
func (q *Quadtree) Insert(node Node) bool {

	// Node cannot be inserted to this tree.
	if !q.inbound(node) {
		return false
	}

	topLeftX := q.TopLeftP.X
	topLeftY := q.TopLeftP.Y
	bottomRightX := q.BottomRightP.X
	bottomRightY := q.BottomRightP.Y

	middleX := (topLeftX + bottomRightX) / 2
	middleY := (topLeftY + bottomRightY) / 2

	inLeft := (topLeftX+bottomRightX)/2 >= node.Point.X
	inTop := (topLeftY+bottomRightY)/2 >= node.Point.Y

	// We are at a quad of unit area
	// We cannot subdivide this quad further.
	if (topLeftX-bottomRightX) <= 1 && (bottomRightY-topLeftY) <= 1 {
		if q.Node == nil {
			q.Node = &node
			return true
		}
		return false
	}

	// The node is in the left part of the tree
	if inLeft {

		// The node is in the top left part of the tree
		if inTop {
			if q.TopLeftTree == nil {
				q.TopLeftTree = &Quadtree{TopLeftP: q.TopLeftP, BottomRightP: Point{X: middleX, Y: middleY}}
			}

			return q.TopLeftTree.Insert(node)

		} else {
			// The node is in the bottom left part of the tree
			if q.BottomLeftTree == nil {
				q.BottomLeftTree = &Quadtree{TopLeftP: Point{X: topLeftX, Y: middleY}, BottomRightP: Point{X: middleX, Y: bottomRightY}}
			}

			return q.BottomLeftTree.Insert(node)
		}

	} else {
		// The node is in the right part of the tree

		// The node is in the top right part of the tree
		if inTop {
			if q.TopRightTree == nil {
				q.TopRightTree = &Quadtree{TopLeftP: Point{X: middleX, Y: topLeftY}, BottomRightP: Point{X: bottomRightX, Y: middleY}}
			}

			return q.TopRightTree.Insert(node)

		} else {
			// The node is in the bottom right part of the tree
			if q.BottomRightTree == nil {
				q.BottomRightTree = &Quadtree{TopLeftP: Point{X: middleX, Y: middleY}, BottomRightP: q.BottomRightP}
			}

			return q.BottomRightTree.Insert(node)
		}

	}
}
