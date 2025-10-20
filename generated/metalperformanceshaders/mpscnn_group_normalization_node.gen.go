// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNGroupNormalizationNode] class.
var (
	CNNGroupNormalizationNodeClass     _CNNGroupNormalizationNodeClass
	CNNGroupNormalizationNodeClassOnce sync.Once
)

func getCNNGroupNormalizationNodeClass() _CNNGroupNormalizationNodeClass {
	CNNGroupNormalizationNodeClassOnce.Do(func() {
		CNNGroupNormalizationNodeClass = _CNNGroupNormalizationNodeClass{objc.GetClass("MPSCNNGroupNormalizationNode")}
	})
	return CNNGroupNormalizationNodeClass
}

type _CNNGroupNormalizationNodeClass struct {
	class objc.Class
}

// An interface definition for the [CNNGroupNormalizationNode] class.
type ICNNGroupNormalizationNode interface {
	IFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode
type CNNGroupNormalizationNode struct {
	FilterNode
}

// CNNGroupNormalizationNodeFrom constructs a [CNNGroupNormalizationNode] from an unsafe.Pointer.
func CNNGroupNormalizationNodeFrom(ptr unsafe.Pointer) CNNGroupNormalizationNode {
	return CNNGroupNormalizationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationNodeClass) Alloc() CNNGroupNormalizationNode {
	rv := objc.Send[CNNGroupNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNGroupNormalizationNodeClass) New() CNNGroupNormalizationNode {
	rv := objc.Send[CNNGroupNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalizationNode) Init() CNNGroupNormalizationNode {
	rv := objc.Send[CNNGroupNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalizationNode) Autorelease() CNNGroupNormalizationNode {
	rv := objc.Send[CNNGroupNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalizationNode creates a new CNNGroupNormalizationNode instance.
func NewCNNGroupNormalizationNode() CNNGroupNormalizationNode {
	return getCNNGroupNormalizationNodeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/init(source:dataSource:)
func NewCNNGroupNormalizationNodeWithSourceDataSource(source unsafe.Pointer, dataSource objc.ID) CNNGroupNormalizationNode {
	instance := getCNNGroupNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNGroupNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/nodeWithSource:dataSource:
func (cc _CNNGroupNormalizationNodeClass) NodeWithSourceDataSource(source unsafe.Pointer, dataSource objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/trainingStyle
func (c_ CNNGroupNormalizationNode) TrainingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("trainingStyle"))
	return rv
}

// SetTrainingStyle sets the value of the trainingStyle property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/trainingStyle
func (c_ CNNGroupNormalizationNode) SetTrainingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}
