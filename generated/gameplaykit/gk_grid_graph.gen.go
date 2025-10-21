// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GridGraph] class.
var (
	GridGraphClass     _GridGraphClass
	GridGraphClassOnce sync.Once
)

func getGridGraphClass() _GridGraphClass {
	GridGraphClassOnce.Do(func() {
		GridGraphClass = _GridGraphClass{objc.GetClass("GKGridGraph")}
	})
	return GridGraphClass
}

type _GridGraphClass struct {
	class objc.Class
}

// An interface definition for the [GridGraph] class.
type IGridGraph interface {
	IGraph
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeToAdjacentNodes(node unsafe.Pointer)
	NodeAtGridPosition(position unsafe.Pointer) unsafe.Pointer
}

// A navigation graph for 2D game worlds where movement is constrained to an integer grid.
//
// Use this class to generate a graph containing objects representing a specified grid. Then use methods of the superclass to find routes through the graph. To learn more about graphs and pathfinding, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph
type GridGraph struct {
	Graph
}

// GridGraphFrom constructs a [GridGraph] from an unsafe.Pointer.
//
// A navigation graph for 2D game worlds where movement is constrained to an integer grid.
func GridGraphFrom(ptr unsafe.Pointer) GridGraph {
	return GridGraph{
		Graph: GraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GridGraphClass) Alloc() GridGraph {
	rv := objc.Send[GridGraph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GridGraphClass) New() GridGraph {
	rv := objc.Send[GridGraph](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridGraph) Init() GridGraph {
	rv := objc.Send[GridGraph](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridGraph) Autorelease() GridGraph {
	rv := objc.Send[GridGraph](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridGraph creates a new GridGraph instance.
func NewGridGraph() GridGraph {
	return getGridGraphClass().New()
}




// Initializes a graph that describes an integer grid with the specified dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/init(fromGridStartingAt:width:height:diagonalsAllowed:)
func NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowed(position unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, diagonalsAllowed bool) GridGraph {
	instance := getGridGraphClass().Alloc()
	rv := objc.Send[GridGraph](instance.ID, objc.Sel("initFromGridStartingAt:width:height:diagonalsAllowed:"), position, width, height, diagonalsAllowed)
	rv.Autorelease()
	return rv
}



// Initializes a graph that describes an integer grid with the specified dimensions, using the specified node class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/init(fromGridStartingAt:width:height:diagonalsAllowed:nodeClass:)
func NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass(position unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, diagonalsAllowed bool, nodeClass objc.Class) GridGraph {
	instance := getGridGraphClass().Alloc()
	rv := objc.Send[GridGraph](instance.ID, objc.Sel("initFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:"), position, width, height, diagonalsAllowed, nodeClass)
	rv.Autorelease()
	return rv
}


// Creates a graph that describes an integer grid with the specified dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/graphFromGridStartingAt:width:height:diagonalsAllowed:
func (gc _GridGraphClass) GraphFromGridStartingAtWidthHeightDiagonalsAllowed(position unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, diagonalsAllowed bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("graphFromGridStartingAt:width:height:diagonalsAllowed:"), position, width, height, diagonalsAllowed)
	return rv
}

// Creates a graph that describes an integer grid with the specified dimensions, using the specified node class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/graphFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:
func (gc _GridGraphClass) GraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass(position unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, diagonalsAllowed bool, nodeClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("graphFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:"), position, width, height, diagonalsAllowed, nodeClass)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/classForGenericArgument(at:)
func (g_ GridGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](g_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}

// Adds the specified node to the graph, connecting it to its nearest neighbors in the grid.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/connectToAdjacentNodes(node:)
func (g_ GridGraph) ConnectNodeToAdjacentNodes(node unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectNodeToAdjacentNodes:"), node)
}

// Returns the node in the graph at the specified grid coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/node(atGridPosition:)
func (g_ GridGraph) NodeAtGridPosition(position unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("nodeAtGridPosition:"), position)
	return rv
}

// A Boolean value that indicates whether nodes in the grid are connected to their diagonal neighbors.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/diagonalsAllowed
func (g_ GridGraph) DiagonalsAllowed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("diagonalsAllowed"))
	return rv
}

// The number of possible y-coordinates in the grid.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridHeight
func (g_ GridGraph) GridHeight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("gridHeight"))
	return rv
}

// The lowest x- and y-coordinates that appear in the grid.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridOrigin
func (g_ GridGraph) GridOrigin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gridOrigin"))
	return rv
}

// The number of possible x-coordinates in the grid.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridWidth
func (g_ GridGraph) GridWidth() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("gridWidth"))
	return rv
}


