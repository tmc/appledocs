// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GraphNode] class.
var (
	GraphNodeClass     _GraphNodeClass
	GraphNodeClassOnce sync.Once
)

func getGraphNodeClass() _GraphNodeClass {
	GraphNodeClassOnce.Do(func() {
		GraphNodeClass = _GraphNodeClass{objc.GetClass("GKGraphNode")}
	})
	return GraphNodeClass
}

type _GraphNodeClass struct {
	class objc.Class
}

// An interface definition for the [GraphNode] class.
type IGraphNode interface {
	objectivec.IObject
	AddConnectionsToNodesBidirectional(nodes []GraphNode, bidirectional bool)
	CostToNode(node IGKGraphNode) float32
	EstimatedCostToNode(node IGKGraphNode) float32
	FindPathFromNode(startNode IGKGraphNode) []GraphNode
	FindPathToNode(goalNode IGKGraphNode) []GraphNode
	RemoveConnectionsToNodesBidirectional(nodes []GraphNode, bidirectional bool)
	ConnectedNodes() []GraphNode
}

// A single node in a navigation graph for use in pathfinding.
//
// A set of connected nodes form a graph that describes the navigability of a game world. Use graph nodes together with a object (or one of its subclasses) to perform actions that relate to the network of nodes as a whole, such as pathfinding to determine routes through the network. This class describes the general features of graph nodes, but does not contain geometry information that relates the graph to a game world. You can construct a graph with this class or any of its subclasses: On its own, the class is useful for worlds such as board games, where the connections between nodes are important but their spatial position has no effect on gameplay. Create objects (for use with the class) to model worlds where movement is constrained to a two-dimensional integer grid. Create objects to model worlds that allow full freedom of movement in a two-dimensional plane. Use these nodes together with the or class to create graphs that route around impassable obstacles. Create objects to model worlds that allow full freedom of movement in three-dimensional space. To learn more about graphs and pathfinding, see in .


// A single node in a navigation graph for use in pathfinding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode
type GraphNode struct {
	objectivec.Object
}

// GraphNodeFrom constructs a [GraphNode] from an unsafe.Pointer.
//
// A single node in a navigation graph for use in pathfinding.
func GraphNodeFrom(ptr unsafe.Pointer) GraphNode {
	return GraphNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphNodeClass) Alloc() GraphNode {
	rv := objc.Send[GraphNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphNodeClass) New() GraphNode {
	rv := objc.Send[GraphNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphNode) Init() GraphNode {
	rv := objc.Send[GraphNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphNode) Autorelease() GraphNode {
	rv := objc.Send[GraphNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphNode creates a new GraphNode instance.
func NewGraphNode() GraphNode {
	return getGraphNodeClass().New()
}



// Connects this node to all nodes in the specified list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/addConnections(to:bidirectional:)
func (g_ GraphNode) AddConnectionsToNodesBidirectional(nodes []GraphNode, bidirectional bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("addConnectionsToNodes:bidirectional:"), nodes, bidirectional)
}


// Returns the cost to travel from this node to the specified, directly connected, node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/cost(to:)
func (g_ GraphNode) CostToNode(node IGKGraphNode) float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("costToNode:"), node)
	return rv
}


// Returns an underestimate of the cost of travel from this node to the specified node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/estimatedCost(to:)
func (g_ GraphNode) EstimatedCostToNode(node IGKGraphNode) float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("estimatedCostToNode:"), node)
	return rv
}


// Computes and returns a sequence of nodes that represents the lowest-cost graph traversal from the specified node to this node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/findPath(from:)
func (g_ GraphNode) FindPathFromNode(startNode IGKGraphNode) []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("findPathFromNode:"), startNode)
	return rv
}


// Computes and returns a sequence of nodes that represents the lowest-cost graph traversal from this node to the specified node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/findPath(to:)
func (g_ GraphNode) FindPathToNode(goalNode IGKGraphNode) []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("findPathToNode:"), goalNode)
	return rv
}


// Removes the connections from this node to the specified nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/removeConnections(to:bidirectional:)
func (g_ GraphNode) RemoveConnectionsToNodesBidirectional(nodes []GraphNode, bidirectional bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeConnectionsToNodes:bidirectional:"), nodes, bidirectional)
}


// The list of other nodes connected to this node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKGraphNode/connectedNodes
func (g_ GraphNode) ConnectedNodes() []GraphNode {
	rv := objc.Send[[]GraphNode](g_.ID, objc.Sel("connectedNodes"))
	return rv
}



