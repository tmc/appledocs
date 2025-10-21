// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

// Enum types and constants
// GKMeshGraphTriangulationMode - Options for how to place graph nodes when generating the graph, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraphTriangulationMode
type MeshGraphTriangulationMode uint

// GKRTreeSplitStrategy - Options that control how a tree balances its internal structure when adding elements, used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy
type RTreeSplitStrategy uint

const (
	// RTreeSplitStrategyHalve - An option to split groups of elements in half based on the order they were added to the tree in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/halve
	RTreeSplitStrategyHalve RTreeSplitStrategy = 0
	// RTreeSplitStrategyLinear - An option to split groups of elements by finding a line that divides space so that half of the elements are on either side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/linear
	RTreeSplitStrategyLinear RTreeSplitStrategy = 0
	// RTreeSplitStrategyQuadratic - An option to split groups of elements by finding the subgroups that occupy the least area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/quadratic
	RTreeSplitStrategyQuadratic RTreeSplitStrategy = 0
	// RTreeSplitStrategyReduceOverlap - An option to split groups of elements by finding the subgroups whose areas overlap the least.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/reduceOverlap
	RTreeSplitStrategyReduceOverlap RTreeSplitStrategy = 0
)


