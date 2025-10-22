// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GraphNode2D] class.
type IGraphNode2D interface {
	IGraphNode
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
}

// A node in a navigation graph, associated with a point in continuous 2D space.
//
// Together, a network of nodes form a graph that describes the navigability of a game world. Use graph nodes with a , , or object to perform actions that relate to the network of nodes as a whole, such as pathfinding to determine routes through the network. When you use the or class to describe a game world in terms of open spaces interrupted by obstacles, GameplayKit automatically creates and manages instances that represent positions along possible paths that navigate around those obstacles. To learn more about graphs and pathfinding, see in .
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphNode2DClass) Alloc() GraphNode2D {
	rv := objc.Send[GraphNode2D](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a graph node with the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/init(point:)
func NewGraphNode2DWithPoint(point unsafe.Pointer) GraphNode2D {
	instance := getGraphNode2DClass().Alloc()
	rv := objc.Send[GraphNode2D](instance.ID, objc.Sel("initWithPoint:"), point)
	rv.Autorelease()
	return rv
}


// Creates a graph node with the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/node(withPoint:)
func (gc _GraphNode2DClass) NodeWithPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("nodeWithPoint:"), point)
	return rv
}

// The position of the node in continuous 2D space.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/position
func (g_ GraphNode2D) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the node in continuous 2D space.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode2D/position
func (g_ GraphNode2D) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}


