// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ObstacleGraph] class.
type IObstacleGraph interface {
	IGraph
	AddObstacles(obstacles []PolygonObstacle)
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	ConnectNodeUsingObstacles(node unsafe.Pointer)
	ConnectNodeUsingObstaclesIgnoringObstacles(node unsafe.Pointer, obstaclesToIgnore []PolygonObstacle)
	ConnectNodeUsingObstaclesIgnoringBufferRadiusOfObstacles(node unsafe.Pointer, obstaclesBufferRadiusToIgnore []PolygonObstacle)
	IsConnectionLockedFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) bool
	LockConnectionFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer)
	NodesForObstacle(obstacle IGKPolygonObstacle) []GraphNode2D
	RemoveAllObstacles()
	RemoveObstacles(obstacles []PolygonObstacle)
	UnlockConnectionFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer)
	BufferRadius() float32
	Obstacles() []PolygonObstacle
}

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

// Alloc allocates a new instance without initialization.
func (oc _ObstacleGraphClass) Alloc() ObstacleGraph {
	rv := objc.Send[ObstacleGraph](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a graph with the specified list of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/init(obstacles:bufferRadius:)
func NewObstacleGraphWithObstaclesBufferRadius(obstacles []PolygonObstacle, bufferRadius float32) ObstacleGraph {
	instance := getObstacleGraphClass().Alloc()
	rv := objc.Send[ObstacleGraph](instance.ID, objc.Sel("initWithObstacles:bufferRadius:"), obstacles, bufferRadius)
	rv.Autorelease()
	return rv
}


// Initializes a graph with the specified list of obstacles, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/init(obstacles:bufferRadius:nodeClass:)
func NewObstacleGraphWithObstaclesBufferRadiusNodeClass(obstacles []PolygonObstacle, bufferRadius float32, nodeClass objc.Class) ObstacleGraph {
	instance := getObstacleGraphClass().Alloc()
	rv := objc.Send[ObstacleGraph](instance.ID, objc.Sel("initWithObstacles:bufferRadius:nodeClass:"), obstacles, bufferRadius, nodeClass)
	rv.Autorelease()
	return rv
}



// Creates a graph with the specified list of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/graphWithObstacles:bufferRadius:
func (oc _ObstacleGraphClass) GraphWithObstaclesBufferRadius(obstacles []PolygonObstacle, bufferRadius float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("graphWithObstacles:bufferRadius:"), obstacles, bufferRadius)
	return rv
}


// Creates a graph with the specified list of obstacles, using the specified node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/graphWithObstacles:bufferRadius:nodeClass:
func (oc _ObstacleGraphClass) GraphWithObstaclesBufferRadiusNodeClass(obstacles []PolygonObstacle, bufferRadius float32, nodeClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("graphWithObstacles:bufferRadius:nodeClass:"), obstacles, bufferRadius, nodeClass)
	return rv
}


// Adds new obstacles to the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/addObstacles(_:)
func (o_ ObstacleGraph) AddObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObstacles:"), obstacles)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/classForGenericArgument(at:)
func (o_ ObstacleGraph) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}


// Adds the specified node to the graph, connecting it to its nearest neighbors without creating connections that pass through obstacles or their buffer regions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:)
func (o_ ObstacleGraph) ConnectNodeUsingObstacles(node unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:"), node)
}


// Adds the specified node to the graph, connecting it to its nearest neighbors while ignoring the area occupied by the specified obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:ignoring:)
func (o_ ObstacleGraph) ConnectNodeUsingObstaclesIgnoringObstacles(node unsafe.Pointer, obstaclesToIgnore []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:ignoringObstacles:"), node, obstaclesToIgnore)
}


// Adds the specified node to the graph, connecting it to its nearest neighbors while ignoring the buffer regions around the specified obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/connectUsingObstacles(node:ignoringBufferRadiusOf:)
func (o_ ObstacleGraph) ConnectNodeUsingObstaclesIgnoringBufferRadiusOfObstacles(node unsafe.Pointer, obstaclesBufferRadiusToIgnore []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("connectNodeUsingObstacles:ignoringBufferRadiusOfObstacles:"), node, obstaclesBufferRadiusToIgnore)
}


// Returns a Boolean value indicating whether the specified nodes are protected from disconnection due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/isConnectionLocked(from:to:)
func (o_ ObstacleGraph) IsConnectionLockedFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isConnectionLockedFromNode:toNode:"), startNode, endNode)
	return rv
}


// Prevents the specified nodes from being disconnected due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/lockConnection(from:to:)
func (o_ ObstacleGraph) LockConnectionFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("lockConnectionFromNode:toNode:"), startNode, endNode)
}


// Returns the group of nodes corresponding to an obstacle in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/nodes(for:)
func (o_ ObstacleGraph) NodesForObstacle(obstacle IGKPolygonObstacle) []GraphNode2D {
	rv := objc.Send[[]GraphNode2D](o_.ID, objc.Sel("nodesForObstacle:"), obstacle)
	return rv
}


// Removes all obstacles from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/removeAllObstacles()
func (o_ ObstacleGraph) RemoveAllObstacles() {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeAllObstacles"))
}


// Removes the specified obstacle from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/removeObstacles(_:)
func (o_ ObstacleGraph) RemoveObstacles(obstacles []PolygonObstacle) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObstacles:"), obstacles)
}


// Allows the specified nodes to be disconnected due to the addition of obstacles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/unlockConnection(from:to:)
func (o_ ObstacleGraph) UnlockConnectionFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("unlockConnectionFromNode:toNode:"), startNode, endNode)
}


// The distance from obstacle edges that should also be considered impassable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/bufferRadius
func (o_ ObstacleGraph) BufferRadius() float32 {
	rv := objc.Send[float32](o_.ID, objc.Sel("bufferRadius"))
	return rv
}


// The list of obstacle objects in the graph, each of which describes a polygon-shaped impassable area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph/obstacles
func (o_ ObstacleGraph) Obstacles() []PolygonObstacle {
	rv := objc.Send[[]PolygonObstacle](o_.ID, objc.Sel("obstacles"))
	return rv
}


