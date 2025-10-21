// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MeshGraph] class.
var (
	MeshGraphClass     _MeshGraphClass
	MeshGraphClassOnce sync.Once
)

func getMeshGraphClass() _MeshGraphClass {
	MeshGraphClassOnce.Do(func() {
		MeshGraphClass = _MeshGraphClass{objc.GetClass("GKMeshGraph")}
	})
	return MeshGraphClass
}

type _MeshGraphClass struct {
	class objc.Class
}

// An interface definition for the [MeshGraph] class.
type IMeshGraph interface {
	IGraph
	AddObstacles(obstacles []PolygonObstacle)
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeUsingObstacles(node unsafe.Pointer)
	RemoveObstacles(obstacles []PolygonObstacle)
	TriangleAtIndex(index uint) unsafe.Pointer
	Triangulate()
}

// A navigation graph for 2D game worlds that creates a space-filling network for smooth pathfinding around obstacles.
//
// To use a mesh graph for pathfinding, add a collection of objects representing impassable areas and objects representing points of interest (such as the current position of a game character and the location it needs to find a route to). Then use methods of the superclass to find routes through the graph. Unlike the related class, a mesh graph creates a space-filling network of graph nodes, resulting in paths that are smooth but not the most efficient. To learn more about graphs and pathfinding, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph
type MeshGraph struct {
	Graph
}

// MeshGraphFrom constructs a [MeshGraph] from an unsafe.Pointer.
//
// A navigation graph for 2D game worlds that creates a space-filling network for smooth pathfinding around obstacles.
func MeshGraphFrom(ptr unsafe.Pointer) MeshGraph {
	return MeshGraph{
		Graph: GraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MeshGraphClass) Alloc() MeshGraph {
	rv := objc.Send[MeshGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MeshGraphClass) New() MeshGraph {
	rv := objc.Send[MeshGraph](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MeshGraph) Init() MeshGraph {
	rv := objc.Send[MeshGraph](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MeshGraph) Autorelease() MeshGraph {
	rv := objc.Send[MeshGraph](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeshGraph creates a new MeshGraph instance.
func NewMeshGraph() MeshGraph {
	return getMeshGraphClass().New()
}




// Initializes a graph to cover the specified area.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/init(bufferRadius:minCoordinate:maxCoordinate:)
func NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinate(bufferRadius unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer) MeshGraph {
	instance := getMeshGraphClass().Alloc()
	rv := objc.Send[MeshGraph](instance.ID, objc.Sel("initWithBufferRadius:minCoordinate:maxCoordinate:"), bufferRadius, min, max)
	rv.Autorelease()
	return rv
}



// Initializes a graph to cover the specified area, using the specified node class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/init(bufferRadius:minCoordinate:maxCoordinate:nodeClass:)
func NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass(bufferRadius unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer, nodeClass objc.Class) MeshGraph {
	instance := getMeshGraphClass().Alloc()
	rv := objc.Send[MeshGraph](instance.ID, objc.Sel("initWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:"), bufferRadius, min, max, nodeClass)
	rv.Autorelease()
	return rv
}


// Creates a graph to cover the specified area.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/graphWithBufferRadius:minCoordinate:maxCoordinate:
func (mc _MeshGraphClass) GraphWithBufferRadiusMinCoordinateMaxCoordinate(bufferRadius unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("graphWithBufferRadius:minCoordinate:maxCoordinate:"), bufferRadius, min, max)
	return rv
}

// Creates a graph to cover the specified area, using the specified node class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/graphWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:
func (mc _MeshGraphClass) GraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass(bufferRadius unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer, nodeClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("graphWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:"), bufferRadius, min, max, nodeClass)
	return rv
}

// Adds new obstacles to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/addObstacles(_:)
func (m_ MeshGraph) AddObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObstacles:"), obstacles)
}

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/classForGenericArgument(at:)
func (m_ MeshGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](m_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}

// Adds the specified node to the graph, connecting it to its nearest neighbors without creating connections that pass through obstacles or their buffer regions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/connectUsingObstacles(node:)
func (m_ MeshGraph) ConnectNodeUsingObstacles(node unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("connectNodeUsingObstacles:"), node)
}

// Removes the specified obstacle from the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/removeObstacles(_:)
func (m_ MeshGraph) RemoveObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObstacles:"), obstacles)
}

// The triangle definition at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangle(at:)
func (m_ MeshGraph) TriangleAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("triangleAtIndex:"), index)
	return rv
}

// Creates or updates the graph with a network of nodes that describes the open space around its obstacles.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulate()
func (m_ MeshGraph) Triangulate() {
	objc.Send[objc.ID](m_.ID, objc.Sel("triangulate"))
}

// The distance from obstacle edges that should also be considered impassable.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/bufferRadius
func (m_ MeshGraph) BufferRadius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bufferRadius"))
	return rv
}

// The list of obstacle objects in the graph, each of which describes a polygon-shaped impassable area.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/obstacles
func (m_ MeshGraph) Obstacles() []PolygonObstacle {
	rv := objc.Send[[]PolygonObstacle](m_.ID, objc.Sel("obstacles"))
	return rv
}

// The number of triangles in the mesh.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangleCount
func (m_ MeshGraph) TriangleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("triangleCount"))
	return rv
}

// A set of options for how to place graph nodes when triangulating the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulationMode
func (m_ MeshGraph) TriangulationMode() MeshGraphTriangulationMode {
	rv := objc.Send[MeshGraphTriangulationMode](m_.ID, objc.Sel("triangulationMode"))
	return rv
}


// SetTriangulationMode sets the value of the triangulationMode property.
// A set of options for how to place graph nodes when triangulating the graph.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulationMode
func (m_ MeshGraph) SetTriangulationMode(value MeshGraphTriangulationMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangulationMode:"), value)
}


