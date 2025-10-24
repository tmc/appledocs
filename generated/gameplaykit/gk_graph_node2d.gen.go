// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGraphNode2D */


/* debug [class_header]: Header for GKGraphNode2D */
// The class instance for the [GraphNode2D] class.
var (
	GraphNode2DClass     _GraphNode2DClass
	GraphNode2DClassOnce sync.Once
)

func getGraphNode2DClass() _GraphNode2DClass {
	GraphNode2DClassOnce.Do(func() {
		GraphNode2DClass = _GraphNode2DClass{objc.GetClass("GKGraphNode2D")}
	})
	return GraphNode2DClass
}

type _GraphNode2DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphNode2D */
// An interface definition for the [GraphNode2D] class.
type IGraphNode2D interface {
	IGraphNode
	
/* debug [class_interface_properties]: Properties for GraphNode2D */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphNode2D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphNode2D */
// Alloc allocates a new instance without initialization.
func (gc _GraphNode2DClass) Alloc() GraphNode2D {
	rv := objc.Send[GraphNode2D](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphNode2DClass) New() GraphNode2D {
	rv := objc.Send[GraphNode2D](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphNode2D) Init() GraphNode2D {
	rv := objc.Send[GraphNode2D](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphNode2D) Autorelease() GraphNode2D {
	rv := objc.Send[GraphNode2D](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphNode2D creates a new GraphNode2D instance.
func NewGraphNode2D() GraphNode2D {
	return getGraphNode2DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphNode2D */
// A node in a navigation graph, associated with a point in continuous 2D space.
//
// Together, a network of nodes form a graph that describes the navigability of a game world. Use graph nodes with a , , or object to perform actions that relate to the network of nodes as a whole, such as pathfinding to determine routes through the network. When you use the or class to describe a game world in terms of open spaces interrupted by obstacles, GameplayKit automatically creates and manages instances that represent positions along possible paths that navigate around those obstacles. To learn more about graphs and pathfinding, see in .


// A node in a navigation graph, associated with a point in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D
type GraphNode2D struct {
	GraphNode
}

// GraphNode2DFrom constructs a [GraphNode2D] from an unsafe.Pointer.
//
// A node in a navigation graph, associated with a point in continuous 2D space.
func GraphNode2DFrom(ptr unsafe.Pointer) GraphNode2D {
	return GraphNode2D{
		GraphNode: GraphNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphNode2D */

// Initializes a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/init(point:)
func NewGraphNode2DWithPoint(point objectivec.IObject) GraphNode2D {
	instance := getGraphNode2DClass().Alloc()
	rv := objc.Send[GraphNode2D](instance.ID, objc.Sel("initWithPoint:"), point)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphNode2DWithPoint */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphNode2D */

// Creates a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/node(withPoint:)
func (gc _GraphNode2DClass) NodeWithPoint(point objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithPoint:"), point)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithPoint) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphNode2D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphNode2D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphNode2D */

// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/position
func (g_ GraphNode2D) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/position
func (g_ GraphNode2D) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGraphNode2D */


