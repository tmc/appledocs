// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGridGraph */


/* debug [class_header]: Header for GKGridGraph */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridGraph */
// An interface definition for the [GridGraph] class.
type IGridGraph interface {
	IGraph
	
/* debug [class_interface_properties]: Properties for GridGraph */
	// properties:
	DiagonalsAllowed() bool
	GridHeight() uint
	GridOrigin() objectivec.IObject
	GridWidth() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridGraph */
	// methods:
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeToAdjacentNodes(node IGKGridGraphNode)
	NodeAtGridPosition(position objectivec.IObject) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridGraph */
// Alloc allocates a new instance without initialization.
func (gc _GridGraphClass) Alloc() GridGraph {
	rv := objc.Send[GridGraph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridGraph */
// A navigation graph for 2D game worlds where movement is constrained to an integer grid.
//
// Use this class to generate a graph containing objects representing a specified grid. Then use methods of the superclass to find routes through the graph. To learn more about graphs and pathfinding, see in .


// A navigation graph for 2D game worlds where movement is constrained to an integer grid.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridGraph */

// Initializes a graph that describes an integer grid with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/init(fromGridStartingAt:width:height:diagonalsAllowed:)
func NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowed(position objectivec.IObject, width int, height int, diagonalsAllowed bool) GridGraph {
	instance := getGridGraphClass().Alloc()
	rv := objc.Send[GridGraph](instance.ID, objc.Sel("initFromGridStartingAt:width:height:diagonalsAllowed:"), position, width, height, diagonalsAllowed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowed */


// Initializes a graph that describes an integer grid with the specified dimensions, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/init(fromGridStartingAt:width:height:diagonalsAllowed:nodeClass:)
func NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass(position objectivec.IObject, width int, height int, diagonalsAllowed bool, nodeClass objc.Class) GridGraph {
	instance := getGridGraphClass().Alloc()
	rv := objc.Send[GridGraph](instance.ID, objc.Sel("initFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:"), position, width, height, diagonalsAllowed, nodeClass)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridGraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridGraph */

// Creates a graph that describes an integer grid with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/graphFromGridStartingAt:width:height:diagonalsAllowed:
func (gc _GridGraphClass) GraphFromGridStartingAtWidthHeightDiagonalsAllowed(position objectivec.IObject, width int, height int, diagonalsAllowed bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphFromGridStartingAt:width:height:diagonalsAllowed:"), position, width, height, diagonalsAllowed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphFromGridStartingAtWidthHeightDiagonalsAllowed) */


// Creates a graph that describes an integer grid with the specified dimensions, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/graphFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:
func (gc _GridGraphClass) GraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass(position objectivec.IObject, width int, height int, diagonalsAllowed bool, nodeClass objc.Class) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphFromGridStartingAt:width:height:diagonalsAllowed:nodeClass:"), position, width, height, diagonalsAllowed, nodeClass)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphFromGridStartingAtWidthHeightDiagonalsAllowedNodeClass) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridGraph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/classForGenericArgument(at:)
func (g_ GridGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](g_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ClassForGenericArgumentAtIndex */


// Adds the specified node to the graph, connecting it to its nearest neighbors in the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/connectToAdjacentNodes(node:)
func (g_ GridGraph) ConnectNodeToAdjacentNodes(node IGKGridGraphNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectNodeToAdjacentNodes:"), node)
}/* debug [instance_methods/method]: ConnectNodeToAdjacentNodes */


// Returns the node in the graph at the specified grid coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/node(atGridPosition:)
func (g_ GridGraph) NodeAtGridPosition(position objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("nodeAtGridPosition:"), position)
	return rv
}/* debug [instance_methods/method]: NodeAtGridPosition */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridGraph */

// A Boolean value that indicates whether nodes in the grid are connected to their diagonal neighbors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/diagonalsAllowed
func (g_ GridGraph) DiagonalsAllowed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("diagonalsAllowed"))
	return rv
}/* debug [instance_properties/getter]: diagonalsAllowed */


// The number of possible y-coordinates in the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridHeight
func (g_ GridGraph) GridHeight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("gridHeight"))
	return rv
}/* debug [instance_properties/getter]: gridHeight */


// The lowest x- and y-coordinates that appear in the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridOrigin
func (g_ GridGraph) GridOrigin() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("gridOrigin"))
	return rv
}/* debug [instance_properties/getter]: gridOrigin */


// The number of possible x-coordinates in the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraph/gridWidth
func (g_ GridGraph) GridWidth() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("gridWidth"))
	return rv
}/* debug [instance_properties/getter]: gridWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGridGraph */


