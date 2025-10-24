// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNBatchNormalizationNode] class.
var (
	CNNBatchNormalizationNodeClass     _CNNBatchNormalizationNodeClass
	CNNBatchNormalizationNodeClassOnce sync.Once
)

func getCNNBatchNormalizationNodeClass() _CNNBatchNormalizationNodeClass {
	CNNBatchNormalizationNodeClassOnce.Do(func() {
		CNNBatchNormalizationNodeClass = _CNNBatchNormalizationNodeClass{objc.GetClass("MPSCNNBatchNormalizationNode")}
	})
	return CNNBatchNormalizationNodeClass
}

type _CNNBatchNormalizationNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNBatchNormalizationNode] class.
type ICNNBatchNormalizationNode interface {
	IFilterNode
	

	// properties:
	Flags() CNNBatchNormalizationFlags get set /* not a class type */
	SetFlags(value CNNBatchNormalizationFlags get set /* not a class type */)
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationNodeClass) Alloc() CNNBatchNormalizationNode {
	rv := objc.Send[CNNBatchNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationNodeClass) New() CNNBatchNormalizationNode {
	rv := objc.Send[CNNBatchNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationNode) Init() CNNBatchNormalizationNode {
	rv := objc.Send[CNNBatchNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationNode) Autorelease() CNNBatchNormalizationNode {
	rv := objc.Send[CNNBatchNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationNode creates a new CNNBatchNormalizationNode instance.
func NewCNNBatchNormalizationNode() CNNBatchNormalizationNode {
	return getCNNBatchNormalizationNodeClass().New()
}





// A representation of a batch normalization kernel.


// A representation of a batch normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode
type CNNBatchNormalizationNode struct {
	FilterNode
}

// CNNBatchNormalizationNodeFrom constructs a [CNNBatchNormalizationNode] from an unsafe.Pointer.
//
// A representation of a batch normalization kernel.
func CNNBatchNormalizationNodeFrom(ptr unsafe.Pointer) CNNBatchNormalizationNode {
	return CNNBatchNormalizationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948004-initwithsource
func NewCNNBatchNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNBatchNormalizationNode {
	instance := getCNNBatchNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948033-nodewithsource
func (cc _CNNBatchNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags
func (c_ CNNBatchNormalizationNode) Flags() CNNBatchNormalizationFlags get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("flags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags
func (c_ CNNBatchNormalizationNode) SetFlags(value CNNBatchNormalizationFlags get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlags:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle
func (c_ CNNBatchNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle
func (c_ CNNBatchNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}







