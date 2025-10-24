// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKObstacleGraph */


/* debug [class_header]: Header for GKObstacleGraph */
// The class instance for the [ObstacleGraph] class.
var (
	ObstacleGraphClass     _ObstacleGraphClass
	ObstacleGraphClassOnce sync.Once
)

func getObstacleGraphClass() _ObstacleGraphClass {
	ObstacleGraphClassOnce.Do(func() {
		ObstacleGraphClass = _ObstacleGraphClass{objc.GetClass("GKObstacleGraph")}
	})
	return ObstacleGraphClass
}

type _ObstacleGraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ObstacleGraph */
// An interface definition for the [ObstacleGraph] class.
type IObstacleGraph interface {
	IGraph
	
/* debug [class_interface_properties]: Properties for ObstacleGraph */
	// properties:
	BufferRadius() float32
	Obstacles() []PolygonObstacle
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ObstacleGraph */
	// methods:
	AddObstacles(obstacles []PolygonObstacle)
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeUsingObstacles(node objectivec.IObject)
	ConnectNodeUsingObstaclesIgnoringObstacles(node objectivec.IObject, obstaclesToIgnore []PolygonObstacle)
	ConnectNodeUsingObstaclesIgnoringBufferRadiusOfObstacles(node objectivec.IObject, obstaclesBufferRadiusToIgnore []PolygonObstacle)
	IsConnectionLockedFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject) bool
	LockConnectionFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject)
	NodesForObstacle(obstacle IGKPolygonObstacle) []GraphNode2D
	RemoveAllObstacles()
	RemoveObstacles(obstacles []PolygonObstacle)
	UnlockConnectionFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ObstacleGraph */
// Alloc allocates a new instance without initialization.
func (oc _ObstacleGraphClass) Alloc() ObstacleGraph {
	rv := objc.Send[ObstacleGraph](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ObstacleGraphClass) New() ObstacleGraph {
	rv := objc.Send[ObstacleGraph](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ObstacleGraph) Init() ObstacleGraph {
	rv := objc.Send[ObstacleGraph](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ObstacleGraph) Autorelease() ObstacleGraph {
	rv := objc.Send[ObstacleGraph](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObstacleGraph creates a new ObstacleGraph instance.
func NewObstacleGraph() ObstacleGraph {
	return getObstacleGraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ObstacleGraph */
// A navigation graph for 2D game worlds that creates a minimal network for precise pathfinding around obstacles.
//
// You create an obstacle graph with a collection of objects. To use the graph for pathfinding, you add objects representing points of interest (such as the current position of a game character and the location it needs to find a route to). Then use methods of the superclass to find routes through the graph. Unlike the related class, an obstacle graph creates a minimal network of graph nodes, resulting in paths that are efficient but not smooth. To learn more about graphs and pathfinding, see in .


// A navigation graph for 2D game worlds that creates a minimal network for precise pathfinding around obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph
type ObstacleGraph struct {
	Graph
}

// ObstacleGraphFrom constructs a [ObstacleGraph] from an unsafe.Pointer.
//
// A navigation graph for 2D game worlds that creates a minimal network for precise pathfinding around obstacles.
func ObstacleGraphFrom(ptr unsafe.Pointer) ObstacleGraph {
	return ObstacleGraph{
		Graph: GraphFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ObstacleGraph */

// Initializes a graph with the specified list of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/init(obstacles:bufferRadius:)
func NewObstacleGraphWithObstaclesBufferRadius(obstacles []PolygonObstacle, bufferRadius float32) ObstacleGraph {
	instance := getObstacleGraphClass().Alloc()
	rv := objc.Send[ObstacleGraph](instance.ID, objc.Sel("initWithObstacles:bufferRadius:"), obstacles, bufferRadius)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewObstacleGraphWithObstaclesBufferRadius */


// Initializes a graph with the specified list of obstacles, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/init(obstacles:bufferRadius:nodeClass:)
func NewObstacleGraphWithObstaclesBufferRadiusNodeClass(obstacles []PolygonObstacle, bufferRadius float32, nodeClass objc.Class) ObstacleGraph {
	instance := getObstacleGraphClass().Alloc()
	rv := objc.Send[ObstacleGraph](instance.ID, objc.Sel("initWithObstacles:bufferRadius:nodeClass:"), obstacles, bufferRadius, nodeClass)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewObstacleGraphWithObstaclesBufferRadiusNodeClass */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ObstacleGraph */

// Creates a graph with the specified list of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/graphWithObstacles:bufferRadius:
func (oc _ObstacleGraphClass) GraphWithObstaclesBufferRadius(obstacles []PolygonObstacle, bufferRadius float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("graphWithObstacles:bufferRadius:"), obstacles, bufferRadius)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithObstaclesBufferRadius) */


// Creates a graph with the specified list of obstacles, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/graphWithObstacles:bufferRadius:nodeClass:
func (oc _ObstacleGraphClass) GraphWithObstaclesBufferRadiusNodeClass(obstacles []PolygonObstacle, bufferRadius float32, nodeClass objc.Class) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("graphWithObstacles:bufferRadius:nodeClass:"), obstacles, bufferRadius, nodeClass)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithObstaclesBufferRadiusNodeClass) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ObstacleGraph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ObstacleGraph */

// Adds new obstacles to the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/addObstacles(_:)
func (o_ ObstacleGraph) AddObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObstacles:"), obstacles)
}/* debug [instance_methods/method]: AddObstacles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/classForGenericArgument(at:)
func (o_ ObstacleGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ClassForGenericArgumentAtIndex */


// Adds the specified node to the graph, connecting it to its nearest neighbors without creating connections that pass through obstacles or their buffer regions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:)
func (o_ ObstacleGraph) ConnectNodeUsingObstacles(node objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:"), node)
}/* debug [instance_methods/method]: ConnectNodeUsingObstacles */


// Adds the specified node to the graph, connecting it to its nearest neighbors while ignoring the area occupied by the specified obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:ignoring:)
func (o_ ObstacleGraph) ConnectNodeUsingObstaclesIgnoringObstacles(node objectivec.IObject, obstaclesToIgnore []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:ignoringObstacles:"), node, obstaclesToIgnore)
}/* debug [instance_methods/method]: ConnectNodeUsingObstaclesIgnoringObstacles */


// Adds the specified node to the graph, connecting it to its nearest neighbors while ignoring the buffer regions around the specified obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:ignoringBufferRadiusOf:)
func (o_ ObstacleGraph) ConnectNodeUsingObstaclesIgnoringBufferRadiusOfObstacles(node objectivec.IObject, obstaclesBufferRadiusToIgnore []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:ignoringBufferRadiusOfObstacles:"), node, obstaclesBufferRadiusToIgnore)
}/* debug [instance_methods/method]: ConnectNodeUsingObstaclesIgnoringBufferRadiusOfObstacles */


// Returns a Boolean value indicating whether the specified nodes are protected from disconnection due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/isConnectionLocked(from:to:)
func (o_ ObstacleGraph) IsConnectionLockedFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isConnectionLockedFromNode:toNode:"), startNode, endNode)
	return rv
}/* debug [instance_methods/method]: IsConnectionLockedFromNodeToNode */


// Prevents the specified nodes from being disconnected due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/lockConnection(from:to:)
func (o_ ObstacleGraph) LockConnectionFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("lockConnectionFromNode:toNode:"), startNode, endNode)
}/* debug [instance_methods/method]: LockConnectionFromNodeToNode */


// Returns the group of nodes corresponding to an obstacle in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/nodes(for:)
func (o_ ObstacleGraph) NodesForObstacle(obstacle IGKPolygonObstacle) []GraphNode2D {
	rv := objc.Send[[]GraphNode2D](o_.ID, objc.Sel("nodesForObstacle:"), obstacle)
	return rv
}/* debug [instance_methods/method]: NodesForObstacle */


// Removes all obstacles from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/removeAllObstacles()
func (o_ ObstacleGraph) RemoveAllObstacles() {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeAllObstacles"))
}/* debug [instance_methods/method]: RemoveAllObstacles */


// Removes the specified obstacle from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/removeObstacles(_:)
func (o_ ObstacleGraph) RemoveObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObstacles:"), obstacles)
}/* debug [instance_methods/method]: RemoveObstacles */


// Allows the specified nodes to be disconnected due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/unlockConnection(from:to:)
func (o_ ObstacleGraph) UnlockConnectionFromNodeToNode(startNode objectivec.IObject, endNode objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("unlockConnectionFromNode:toNode:"), startNode, endNode)
}/* debug [instance_methods/method]: UnlockConnectionFromNodeToNode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ObstacleGraph */

// The distance from obstacle edges that should also be considered impassable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/bufferRadius
func (o_ ObstacleGraph) BufferRadius() float32 {
	rv := objc.Send[float32](o_.ID, objc.Sel("bufferRadius"))
	return rv
}/* debug [instance_properties/getter]: bufferRadius */


// The list of obstacle objects in the graph, each of which describes a polygon-shaped impassable area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/obstacles
func (o_ ObstacleGraph) Obstacles() []PolygonObstacle {
	rv := objc.Send[[]PolygonObstacle](o_.ID, objc.Sel("obstacles"))
	return rv
}/* debug [instance_properties/getter]: obstacles */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKObstacleGraph */


