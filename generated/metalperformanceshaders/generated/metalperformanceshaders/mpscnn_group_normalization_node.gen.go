// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNGroupNormalizationNode */


/* debug [class_header]: Header for MPSCNNGroupNormalizationNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNGroupNormalizationNode */
// An interface definition for the [CNNGroupNormalizationNode] class.
type ICNNGroupNormalizationNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNGroupNormalizationNode */
	// properties:
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNGroupNormalizationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNGroupNormalizationNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNGroupNormalizationNode */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNGroupNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152572-initwithsource
func NewCNNGroupNormalizationNodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) CNNGroupNormalizationNode {
	instance := getCNNGroupNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNGroupNormalizationNode](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNGroupNormalizationNodeWithSourceDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNGroupNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152573-nodewithsource
func (cc _CNNGroupNormalizationNodeClass) NodeWithSourceDataSource(source IImageNode, dataSource unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceDataSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNGroupNormalizationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNGroupNormalizationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNGroupNormalizationNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle
func (c_ CNNGroupNormalizationNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}/* debug [instance_properties/getter]: trainingStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle
func (c_ CNNGroupNormalizationNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}/* debug [instance_properties/setter]: trainingStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNGroupNormalizationNode */


