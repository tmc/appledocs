// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationNodeClass) Alloc() CNNGroupNormalizationNode {
	rv := objc.Send[CNNGroupNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152572-initwithsource
func NewCNNGroupNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNGroupNormalizationNode {
	instance := getCNNGroupNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNGroupNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152573-nodewithsource
func (cc _CNNGroupNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle
func (c_ CNNGroupNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle
func (c_ CNNGroupNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}







