// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Graph] class.
var (
	GraphClass     _GraphClass
	GraphClassOnce sync.Once
)

func getGraphClass() _GraphClass {
	GraphClassOnce.Do(func() {
		GraphClass = _GraphClass{objc.GetClass("GKGraph")}
	})
	return GraphClass
}

type _GraphClass struct {
	class objc.Class
}

// An interface definition for the [Graph] class.
type IGraph interface {
	objectivec.IObject
	AddNodes(nodes unsafe.Pointer)
	ConnectNodeToLowestCostNodeBidirectional(node unsafe.Pointer, bidirectional bool)
	FindPathFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) []GraphNode
	RemoveNodes(nodes unsafe.Pointer)
}

// A collection of nodes that describes the navigability of a game world and provides methods to search for routes through that space.
//
// Individual nodes in a graph represent discrete locations that a character or other object in your game can occupy, and the connections between adjacent nodes represent the ability of a game entity to travel from one location to another. Use the class to create a general graph, or the , , or subclass to generate specialized graphs that contain more information about the geometry of your game world. Each set of graph and node classes can generate graphs for different kinds of spaces: The base classes and contain functionality general to all graphs and nodes. You can also use these classes on their own to construct graphs that contain no geometry information. This option is useful for games where the connections between spaces are more important than their physical locations, such as board games. Use the and classes to describe game worlds that constrain movement to an integer grid, such as tactical role-playing games. Use the or class to describe 2D game worlds that allow continuous movement in open spaces that are interrupted by impassable obstacles ( objects). Obstacle graphs automatically generate nodes containing 2D point information ( objects), and you can also add your own such nodes representing locations of interest. The graphs modeled by this class are always —that is, a connection between two nodes describes one direction of travel between them. To enable travel between two nodes in either direction, you must create a connection in each direction. You can choose to connect both directions at once with the method (for graphs) or the addConnection:bidirectional: method (for nodes). Using a graph for pathfinding typically involves three major steps: Create a graph once (for example, when initializing a game level class) with static information about your game world. When you need to find a route between points, connect temporary nodes to the graph at those points. Use the method to connect nodes using their own geometry information, or the or method to use the additional constraints of obstacle and grid graphs. Call the method to find a route between locations in the graph. This method returns an array of graph nodes, starting with the requested start point of the path, and proceeding to adjacent nodes in order until it reaches the requested end point. Use the geometry information contained in each node to make use of the route—for example, in a SpriteKit game you might create a sequence of move actions to move a character from point to point along the path. The temporary nodes you created for finding a path typically have little usefulness after a path has been found. Remove those nodes before reusing the graph for future searches. To learn more about graphs and pathfinding, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph
type Graph struct {
	objectivec.Object
}

// GraphFrom constructs a [Graph] from an unsafe.Pointer.
//
// A collection of nodes that describes the navigability of a game world and provides methods to search for routes through that space.
func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphClass) Alloc() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphClass) New() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Graph) Init() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Graph) Autorelease() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraph creates a new Graph instance.
func NewGraph() Graph {
	return getGraphClass().New()
}




// Initializes a graph with the specified list of nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/init(_:)
func NewGraphWithNodes(nodes unsafe.Pointer) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithNodes:"), nodes)
	rv.Autorelease()
	return rv
}


// Creates a graph with the specified list of nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/graphWithNodes:
func (gc _GraphClass) GraphWithNodes(nodes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("graphWithNodes:"), nodes)
	return rv
}

// Adds the specified nodes to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/add(_:)
func (g_ Graph) AddNodes(nodes unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("addNodes:"), nodes)
}

// Adds a node to the graph, connecting it to the node already in the graph for which the connection has the lowest cost.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/connectToLowestCostNode(node:bidirectional:)
func (g_ Graph) ConnectNodeToLowestCostNodeBidirectional(node unsafe.Pointer, bidirectional bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectNodeToLowestCostNode:bidirectional:"), node, bidirectional)
}

// Computes and returns a sequence of nodes that represents the shortest traversal of the graph between the specified nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/findPath(from:to:)
func (g_ Graph) FindPathFromNodeToNode(startNode unsafe.Pointer, endNode unsafe.Pointer) []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("findPathFromNode:toNode:"), startNode, endNode)
	return rv
}

// Removes the specified nodes from the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/remove(_:)
func (g_ Graph) RemoveNodes(nodes unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeNodes:"), nodes)
}

// The list of nodes in the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/nodes
func (g_ Graph) Nodes() []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("nodes"))
	return rv
}


