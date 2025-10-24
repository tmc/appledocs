// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGraphNode3D */


/* debug [class_header]: Header for GKGraphNode3D */
// The class instance for the [GraphNode3D] class.
var (
	GraphNode3DClass     _GraphNode3DClass
	GraphNode3DClassOnce sync.Once
)

func getGraphNode3DClass() _GraphNode3DClass {
	GraphNode3DClassOnce.Do(func() {
		GraphNode3DClass = _GraphNode3DClass{objc.GetClass("GKGraphNode3D")}
	})
	return GraphNode3DClass
}

type _GraphNode3DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphNode3D */
// An interface definition for the [GraphNode3D] class.
type IGraphNode3D interface {
	IGraphNode
	
/* debug [class_interface_properties]: Properties for GraphNode3D */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphNode3D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphNode3D */
// Alloc allocates a new instance without initialization.
func (gc _GraphNode3DClass) Alloc() GraphNode3D {
	rv := objc.Send[GraphNode3D](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphNode3DClass) New() GraphNode3D {
	rv := objc.Send[GraphNode3D](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphNode3D) Init() GraphNode3D {
	rv := objc.Send[GraphNode3D](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphNode3D) Autorelease() GraphNode3D {
	rv := objc.Send[GraphNode3D](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphNode3D creates a new GraphNode3D instance.
func NewGraphNode3D() GraphNode3D {
	return getGraphNode3DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphNode3D */
// A node in a navigation graph, associated with a point in continuous 3D space.
//
// Together, a network of nodes form a graph that describes the navigability of a game world. Use graph nodes with a object to perform actions that relate to the network of nodes as a whole, such as pathfinding to determine routes through the network. To learn more about graphs and pathfinding, see in .


// A node in a navigation graph, associated with a point in continuous 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D
type GraphNode3D struct {
	GraphNode
}

// GraphNode3DFrom constructs a [GraphNode3D] from an unsafe.Pointer.
//
// A node in a navigation graph, associated with a point in continuous 3D space.
func GraphNode3DFrom(ptr unsafe.Pointer) GraphNode3D {
	return GraphNode3D{
		GraphNode: GraphNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphNode3D */

// Initializes a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/init(point:)
func NewGraphNode3DWithPoint(point objectivec.IObject) GraphNode3D {
	instance := getGraphNode3DClass().Alloc()
	rv := objc.Send[GraphNode3D](instance.ID, objc.Sel("initWithPoint:"), point)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphNode3DWithPoint */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphNode3D */

// Creates a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/node(withPoint:)
func (gc _GraphNode3DClass) NodeWithPoint(point objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithPoint:"), point)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithPoint) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphNode3D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphNode3D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphNode3D */

// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/position
func (g_ GraphNode3D) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/position
func (g_ GraphNode3D) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGraphNode3D */


