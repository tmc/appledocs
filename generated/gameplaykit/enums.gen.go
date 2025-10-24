// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

/* debug [enums.gen.go]: Generating 2 enums for GameplayKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum GKMeshGraphTriangulationMode (3 cases) */
// GKMeshGraphTriangulationMode - Options for how to place graph nodes when generating the graph, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraphTriangulationMode
type GKMeshGraphTriangulationMode uint

const (
	// GKMeshGraphTriangulationModeCenters - An option to place graph nodes at the center of each polygon in the generated mesh.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraphTriangulationMode/centers
	GKMeshGraphTriangulationModeCenters GKMeshGraphTriangulationMode = 0
	// GKMeshGraphTriangulationModeEdgeMidpoints - An option to place graph nodes at the midpoint of each in the generated mesh.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraphTriangulationMode/edgeMidpoints
	GKMeshGraphTriangulationModeEdgeMidpoints GKMeshGraphTriangulationMode = 0
	// GKMeshGraphTriangulationModeVertices - An option to place graph nodes at each vertex in the generated mesh.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKMeshGraphTriangulationMode/vertices
	GKMeshGraphTriangulationModeVertices GKMeshGraphTriangulationMode = 0
)

/* debug [enums.gen.go]: Processing enum GKRTreeSplitStrategy (4 cases) */
// GKRTreeSplitStrategy - Options that control how a tree balances its internal structure when adding elements, used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy
type GKRTreeSplitStrategy uint

const (
	// GKRTreeSplitStrategyHalve - An option to split groups of elements in half based on the order they were added to the tree in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/halve
	GKRTreeSplitStrategyHalve GKRTreeSplitStrategy = 0
	// GKRTreeSplitStrategyLinear - An option to split groups of elements by finding a line that divides space so that half of the elements are on either side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/linear
	GKRTreeSplitStrategyLinear GKRTreeSplitStrategy = 0
	// GKRTreeSplitStrategyQuadratic - An option to split groups of elements by finding the subgroups that occupy the least area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/quadratic
	GKRTreeSplitStrategyQuadratic GKRTreeSplitStrategy = 0
	// GKRTreeSplitStrategyReduceOverlap - An option to split groups of elements by finding the subgroups whose areas overlap the least.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTreeSplitStrategy/reduceOverlap
	GKRTreeSplitStrategyReduceOverlap GKRTreeSplitStrategy = 0
)


