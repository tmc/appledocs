// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GraphNode3D] class.
type IGraphNode3D interface {
	IGraphNode
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GraphNode3DClass) Alloc() GraphNode3D {
	rv := objc.Send[GraphNode3D](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/init(point:)
func NewGraphNode3DWithPoint(point unsafe.Pointer) GraphNode3D {
	instance := getGraphNode3DClass().Alloc()
	rv := objc.Send[GraphNode3D](instance.ID, objc.Sel("initWithPoint:"), point)
	rv.Autorelease()
	return rv
}



// Creates a graph node with the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/node(withPoint:)
func (gc _GraphNode3DClass) NodeWithPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("nodeWithPoint:"), point)
	return rv
}


// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/position
func (g_ GraphNode3D) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("position"))
	return rv
}


// The position of the node in continuous 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode3D/position
func (g_ GraphNode3D) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}


