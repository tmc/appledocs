// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBatchNormalizationNode */


/* debug [class_header]: Header for MPSCNNBatchNormalizationNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBatchNormalizationNode */
// An interface definition for the [CNNBatchNormalizationNode] class.
type ICNNBatchNormalizationNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNBatchNormalizationNode */
	// properties:
	Flags() CNNBatchNormalizationFlags get set /* not a class type */
	SetFlags(value CNNBatchNormalizationFlags get set /* not a class type */)
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBatchNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBatchNormalizationNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBatchNormalizationNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBatchNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948004-initwithsource
func NewCNNBatchNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNBatchNormalizationNode {
	instance := getCNNBatchNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationNodeWithSourceDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBatchNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948033-nodewithsource
func (cc _CNNBatchNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceDataSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBatchNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBatchNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBatchNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags
func (c_ CNNBatchNormalizationNode) Flags() CNNBatchNormalizationFlags get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("flags"))
	return rv
}/* debug [instance_properties/getter]: flags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags
func (c_ CNNBatchNormalizationNode) SetFlags(value CNNBatchNormalizationFlags get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlags:"), value)
}/* debug [instance_properties/setter]: flags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle
func (c_ CNNBatchNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}/* debug [instance_properties/getter]: trainingStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle
func (c_ CNNBatchNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}/* debug [instance_properties/setter]: trainingStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBatchNormalizationNode */


