// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMeshGraph */


/* debug [class_header]: Header for GKMeshGraph */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MeshGraph */
// An interface definition for the [MeshGraph] class.
type IMeshGraph interface {
	IGraph
	
/* debug [class_interface_properties]: Properties for MeshGraph */
	// properties:
	BufferRadius() float32
	Obstacles() []PolygonObstacle
	TriangleCount() uint
	TriangulationMode() MeshGraphTriangulationMode
	SetTriangulationMode(value MeshGraphTriangulationMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MeshGraph */
	// methods:
	AddObstacles(obstacles []PolygonObstacle)
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeUsingObstacles(node objectivec.IObject)
	RemoveObstacles(obstacles []PolygonObstacle)
	TriangleAtIndex(index uint) objc.IObject /* cross-framework: GKTriangle */
	Triangulate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MeshGraph */
// Alloc allocates a new instance without initialization.
func (mc _MeshGraphClass) Alloc() MeshGraph {
	rv := objc.Send[MeshGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MeshGraph */
// A navigation graph for 2D game worlds that creates a space-filling network for smooth pathfinding around obstacles.
//
// To use a mesh graph for pathfinding, add a collection of objects representing impassable areas and objects representing points of interest (such as the current position of a game character and the location it needs to find a route to). Then use methods of the superclass to find routes through the graph. Unlike the related class, a mesh graph creates a space-filling network of graph nodes, resulting in paths that are smooth but not the most efficient. To learn more about graphs and pathfinding, see in .


// A navigation graph for 2D game worlds that creates a space-filling network for smooth pathfinding around obstacles.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MeshGraph */

// Initializes a graph to cover the specified area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/init(bufferRadius:minCoordinate:maxCoordinate:)
func NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinate(bufferRadius float32, min objectivec.IObject, max objectivec.IObject) MeshGraph {
	instance := getMeshGraphClass().Alloc()
	rv := objc.Send[MeshGraph](instance.ID, objc.Sel("initWithBufferRadius:minCoordinate:maxCoordinate:"), bufferRadius, min, max)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinate */


// Initializes a graph to cover the specified area, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/init(bufferRadius:minCoordinate:maxCoordinate:nodeClass:)
func NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass(bufferRadius float32, min objectivec.IObject, max objectivec.IObject, nodeClass objc.Class) MeshGraph {
	instance := getMeshGraphClass().Alloc()
	rv := objc.Send[MeshGraph](instance.ID, objc.Sel("initWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:"), bufferRadius, min, max, nodeClass)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMeshGraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MeshGraph */

// Creates a graph to cover the specified area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/graphWithBufferRadius:minCoordinate:maxCoordinate:
func (mc _MeshGraphClass) GraphWithBufferRadiusMinCoordinateMaxCoordinate(bufferRadius float32, min objectivec.IObject, max objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("graphWithBufferRadius:minCoordinate:maxCoordinate:"), bufferRadius, min, max)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithBufferRadiusMinCoordinateMaxCoordinate) */


// Creates a graph to cover the specified area, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/graphWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:
func (mc _MeshGraphClass) GraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass(bufferRadius float32, min objectivec.IObject, max objectivec.IObject, nodeClass objc.Class) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("graphWithBufferRadius:minCoordinate:maxCoordinate:nodeClass:"), bufferRadius, min, max, nodeClass)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithBufferRadiusMinCoordinateMaxCoordinateNodeClass) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MeshGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MeshGraph */

// Adds new obstacles to the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/addObstacles(_:)
func (m_ MeshGraph) AddObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObstacles:"), obstacles)
}/* debug [instance_methods/method]: AddObstacles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/classForGenericArgument(at:)
func (m_ MeshGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](m_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ClassForGenericArgumentAtIndex */


// Adds the specified node to the graph, connecting it to its nearest neighbors without creating connections that pass through obstacles or their buffer regions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/connectUsingObstacles(node:)
func (m_ MeshGraph) ConnectNodeUsingObstacles(node objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("connectNodeUsingObstacles:"), node)
}/* debug [instance_methods/method]: ConnectNodeUsingObstacles */


// Removes the specified obstacle from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/removeObstacles(_:)
func (m_ MeshGraph) RemoveObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObstacles:"), obstacles)
}/* debug [instance_methods/method]: RemoveObstacles */


// The triangle definition at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangle(at:)
func (m_ MeshGraph) TriangleAtIndex(index uint) objc.IObject /* cross-framework: GKTriangle */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("triangleAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TriangleAtIndex */


// Creates or updates the graph with a network of nodes that describes the open space around its obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulate()
func (m_ MeshGraph) Triangulate() {
	objc.Send[objc.ID](m_.ID, objc.Sel("triangulate"))
}/* debug [instance_methods/method]: Triangulate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MeshGraph */

// The distance from obstacle edges that should also be considered impassable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/bufferRadius
func (m_ MeshGraph) BufferRadius() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("bufferRadius"))
	return rv
}/* debug [instance_properties/getter]: bufferRadius */


// The list of obstacle objects in the graph, each of which describes a polygon-shaped impassable area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/obstacles
func (m_ MeshGraph) Obstacles() []PolygonObstacle {
	rv := objc.Send[[]PolygonObstacle](m_.ID, objc.Sel("obstacles"))
	return rv
}/* debug [instance_properties/getter]: obstacles */


// The number of triangles in the mesh.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangleCount
func (m_ MeshGraph) TriangleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("triangleCount"))
	return rv
}/* debug [instance_properties/getter]: triangleCount */


// A set of options for how to place graph nodes when triangulating the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulationMode
func (m_ MeshGraph) TriangulationMode() MeshGraphTriangulationMode {
	rv := objc.Send[MeshGraphTriangulationMode](m_.ID, objc.Sel("triangulationMode"))
	return rv
}/* debug [instance_properties/getter]: triangulationMode */


// A set of options for how to place graph nodes when triangulating the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraph/triangulationMode
func (m_ MeshGraph) SetTriangulationMode(value MeshGraphTriangulationMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangulationMode:"), value)
}/* debug [instance_properties/setter]: triangulationMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMeshGraph */


