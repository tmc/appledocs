// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKGraph */


/* debug [class_header]: Header for GKGraph */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Graph */
// An interface definition for the [Graph] class.
type IGraph interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Graph */
	// properties:
	Nodes() []GraphNode
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Graph */
	// methods:
	AddNodes(nodes []GraphNode)
	ConnectNodeToLowestCostNodeBidirectional(node IGKGraphNode, bidirectional bool)
	FindPathFromNodeToNode(startNode IGKGraphNode, endNode IGKGraphNode) []GraphNode
	RemoveNodes(nodes []GraphNode)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Graph */
// Alloc allocates a new instance without initialization.
func (gc _GraphClass) Alloc() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Graph */
// A collection of nodes that describes the navigability of a game world and provides methods to search for routes through that space.
//
// Individual nodes in a graph represent discrete locations that a character or other object in your game can occupy, and the connections between adjacent nodes represent the ability of a game entity to travel from one location to another. Use the class to create a general graph, or the , , or subclass to generate specialized graphs that contain more information about the geometry of your game world. Each set of graph and node classes can generate graphs for different kinds of spaces: The base classes and contain functionality general to all graphs and nodes. You can also use these classes on their own to construct graphs that contain no geometry information. This option is useful for games where the connections between spaces are more important than their physical locations, such as board games. Use the and classes to describe game worlds that constrain movement to an integer grid, such as tactical role-playing games. Use the or class to describe 2D game worlds that allow continuous movement in open spaces that are interrupted by impassable obstacles ( objects). Obstacle graphs automatically generate nodes containing 2D point information ( objects), and you can also add your own such nodes representing locations of interest. The graphs modeled by this class are always —that is, a connection between two nodes describes one direction of travel between them. To enable travel between two nodes in either direction, you must create a connection in each direction. You can choose to connect both directions at once with the method (for graphs) or the addConnection:bidirectional: method (for nodes). Using a graph for pathfinding typically involves three major steps: Create a graph once (for example, when initializing a game level class) with static information about your game world. When you need to find a route between points, connect temporary nodes to the graph at those points. Use the method to connect nodes using their own geometry information, or the or method to use the additional constraints of obstacle and grid graphs. Call the method to find a route between locations in the graph. This method returns an array of graph nodes, starting with the requested start point of the path, and proceeding to adjacent nodes in order until it reaches the requested end point. Use the geometry information contained in each node to make use of the route—for example, in a SpriteKit game you might create a sequence of move actions to move a character from point to point along the path. The temporary nodes you created for finding a path typically have little usefulness after a path has been found. Remove those nodes before reusing the graph for future searches. To learn more about graphs and pathfinding, see in .


// A collection of nodes that describes the navigability of a game world and provides methods to search for routes through that space.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Graph */

// Initializes a graph with the specified list of nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/init(_:)
func NewGraphWithNodes(nodes []GraphNode) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithNodes:"), nodes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphWithNodes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Graph */

// Creates a graph with the specified list of nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/graphWithNodes:
func (gc _GraphClass) GraphWithNodes(nodes []GraphNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphWithNodes:"), nodes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithNodes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Graph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Graph */

// Adds the specified nodes to the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/add(_:)
func (g_ Graph) AddNodes(nodes []GraphNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("addNodes:"), nodes)
}/* debug [instance_methods/method]: AddNodes */


// Adds a node to the graph, connecting it to the node already in the graph for which the connection has the lowest cost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/connectToLowestCostNode(node:bidirectional:)
func (g_ Graph) ConnectNodeToLowestCostNodeBidirectional(node IGKGraphNode, bidirectional bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectNodeToLowestCostNode:bidirectional:"), node, bidirectional)
}/* debug [instance_methods/method]: ConnectNodeToLowestCostNodeBidirectional */


// Computes and returns a sequence of nodes that represents the shortest traversal of the graph between the specified nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/findPath(from:to:)
func (g_ Graph) FindPathFromNodeToNode(startNode IGKGraphNode, endNode IGKGraphNode) []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("findPathFromNode:toNode:"), startNode, endNode)
	return rv
}/* debug [instance_methods/method]: FindPathFromNodeToNode */


// Removes the specified nodes from the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/remove(_:)
func (g_ Graph) RemoveNodes(nodes []GraphNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeNodes:"), nodes)
}/* debug [instance_methods/method]: RemoveNodes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Graph */

// The list of nodes in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraph/nodes
func (g_ Graph) Nodes() []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("nodes"))
	return rv
}/* debug [instance_properties/getter]: nodes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGraph */


