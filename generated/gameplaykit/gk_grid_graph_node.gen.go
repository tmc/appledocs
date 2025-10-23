// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GridGraphNode] class.
var (
	GridGraphNodeClass     _GridGraphNodeClass
	GridGraphNodeClassOnce sync.Once
)

func getGridGraphNodeClass() _GridGraphNodeClass {
	GridGraphNodeClassOnce.Do(func() {
		GridGraphNodeClass = _GridGraphNodeClass{objc.GetClass("GKGridGraphNode")}
	})
	return GridGraphNodeClass
}

type _GridGraphNodeClass struct {
	class objc.Class
}

// An interface definition for the [GridGraphNode] class.
type IGridGraphNode interface {
	IGraphNode
	GridPosition() unsafe.Pointer
}

// A node in a navigation graph, associated with a position on a discrete two-dimensional grid.
//
// Together, a network of nodes form a graph that describes the navigability of a game world. Use graph nodes with a object (and methods of its superclass ) to perform actions that relate to the network of nodes as a whole, such as pathfinding to determine routes through the network. To learn more about graphs and pathfinding, see in .


// A node in a navigation graph, associated with a position on a discrete two-dimensional grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode
type GridGraphNode struct {
	GraphNode
}

// GridGraphNodeFrom constructs a [GridGraphNode] from an unsafe.Pointer.
//
// A node in a navigation graph, associated with a position on a discrete two-dimensional grid.
func GridGraphNodeFrom(ptr unsafe.Pointer) GridGraphNode {
	return GridGraphNode{
		GraphNode: GraphNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GridGraphNodeClass) Alloc() GridGraphNode {
	rv := objc.Send[GridGraphNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GridGraphNodeClass) New() GridGraphNode {
	rv := objc.Send[GridGraphNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridGraphNode) Init() GridGraphNode {
	rv := objc.Send[GridGraphNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridGraphNode) Autorelease() GridGraphNode {
	rv := objc.Send[GridGraphNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridGraphNode creates a new GridGraphNode instance.
func NewGridGraphNode() GridGraphNode {
	return getGridGraphNodeClass().New()
}



// Initializes a graph node with the specified position on a grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/init(gridPosition:)
func NewGridGraphNodeWithGridPosition(gridPosition unsafe.Pointer) GridGraphNode {
	instance := getGridGraphNodeClass().Alloc()
	rv := objc.Send[GridGraphNode](instance.ID, objc.Sel("initWithGridPosition:"), gridPosition)
	rv.Autorelease()
	return rv
}



// Creates a graph node with the specified position on a grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/nodeWithGridPosition:
func (gc _GridGraphNodeClass) NodeWithGridPosition(gridPosition unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("nodeWithGridPosition:"), gridPosition)
	return rv
}


// The position of the node on a discrete integer grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/gridPosition
func (g_ GridGraphNode) GridPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gridPosition"))
	return rv
}


