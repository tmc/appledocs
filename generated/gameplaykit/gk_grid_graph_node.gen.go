// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGridGraphNode */


/* debug [class_header]: Header for GKGridGraphNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridGraphNode */
// An interface definition for the [GridGraphNode] class.
type IGridGraphNode interface {
	IGraphNode
	
/* debug [class_interface_properties]: Properties for GridGraphNode */
	// properties:
	GridPosition() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridGraphNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridGraphNode */
// Alloc allocates a new instance without initialization.
func (gc _GridGraphNodeClass) Alloc() GridGraphNode {
	rv := objc.Send[GridGraphNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridGraphNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridGraphNode */

// Initializes a graph node with the specified position on a grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/init(gridPosition:)
func NewGridGraphNodeWithGridPosition(gridPosition objectivec.IObject) GridGraphNode {
	instance := getGridGraphNodeClass().Alloc()
	rv := objc.Send[GridGraphNode](instance.ID, objc.Sel("initWithGridPosition:"), gridPosition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridGraphNodeWithGridPosition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridGraphNode */

// Creates a graph node with the specified position on a grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/nodeWithGridPosition:
func (gc _GridGraphNodeClass) NodeWithGridPosition(gridPosition objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithGridPosition:"), gridPosition)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithGridPosition) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridGraphNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridGraphNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridGraphNode */

// The position of the node on a discrete integer grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGridGraphNode/gridPosition
func (g_ GridGraphNode) GridPosition() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("gridPosition"))
	return rv
}/* debug [instance_properties/getter]: gridPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGridGraphNode */


