// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNInstanceNormalizationNode] class.
var (
	CNNInstanceNormalizationNodeClass     _CNNInstanceNormalizationNodeClass
	CNNInstanceNormalizationNodeClassOnce sync.Once
)

func getCNNInstanceNormalizationNodeClass() _CNNInstanceNormalizationNodeClass {
	CNNInstanceNormalizationNodeClassOnce.Do(func() {
		CNNInstanceNormalizationNodeClass = _CNNInstanceNormalizationNodeClass{objc.GetClass("MPSCNNInstanceNormalizationNode")}
	})
	return CNNInstanceNormalizationNodeClass
}

type _CNNInstanceNormalizationNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNInstanceNormalizationNode] class.
type ICNNInstanceNormalizationNode interface {
	IFilterNode
	

	// properties:
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationNodeClass) Alloc() CNNInstanceNormalizationNode {
	rv := objc.Send[CNNInstanceNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationNodeClass) New() CNNInstanceNormalizationNode {
	rv := objc.Send[CNNInstanceNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationNode) Init() CNNInstanceNormalizationNode {
	rv := objc.Send[CNNInstanceNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationNode) Autorelease() CNNInstanceNormalizationNode {
	rv := objc.Send[CNNInstanceNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationNode creates a new CNNInstanceNormalizationNode instance.
func NewCNNInstanceNormalizationNode() CNNInstanceNormalizationNode {
	return getCNNInstanceNormalizationNodeClass().New()
}





// A representation of an instance normalization kernel.


// A representation of an instance normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode
type CNNInstanceNormalizationNode struct {
	FilterNode
}

// CNNInstanceNormalizationNodeFrom constructs a [CNNInstanceNormalizationNode] from an unsafe.Pointer.
//
// A representation of an instance normalization kernel.
func CNNInstanceNormalizationNodeFrom(ptr unsafe.Pointer) CNNInstanceNormalizationNode {
	return CNNInstanceNormalizationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951940-initwithsource
func NewCNNInstanceNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNInstanceNormalizationNode {
	instance := getCNNInstanceNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNInstanceNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951941-nodewithsource
func (cc _CNNInstanceNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle
func (c_ CNNInstanceNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle
func (c_ CNNInstanceNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}







