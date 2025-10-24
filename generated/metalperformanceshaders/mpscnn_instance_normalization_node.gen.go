// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNInstanceNormalizationNode */


/* debug [class_header]: Header for MPSCNNInstanceNormalizationNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNInstanceNormalizationNode */
// An interface definition for the [CNNInstanceNormalizationNode] class.
type ICNNInstanceNormalizationNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNInstanceNormalizationNode */
	// properties:
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNInstanceNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNInstanceNormalizationNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNInstanceNormalizationNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNInstanceNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951940-initwithsource
func NewCNNInstanceNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNInstanceNormalizationNode {
	instance := getCNNInstanceNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNInstanceNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNInstanceNormalizationNodeWithSourceDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNInstanceNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951941-nodewithsource
func (cc _CNNInstanceNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceDataSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNInstanceNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNInstanceNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNInstanceNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle
func (c_ CNNInstanceNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}/* debug [instance_properties/getter]: trainingStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle
func (c_ CNNInstanceNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}/* debug [instance_properties/setter]: trainingStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNInstanceNormalizationNode */


