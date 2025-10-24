// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNFilterNode */


/* debug [class_header]: Header for MPSNNFilterNode */
// The class instance for the [FilterNode] class.
var (
	FilterNodeClass     _FilterNodeClass
	FilterNodeClassOnce sync.Once
)

func getFilterNodeClass() _FilterNodeClass {
	FilterNodeClassOnce.Do(func() {
		FilterNodeClass = _FilterNodeClass{objc.GetClass("MPSNNFilterNode")}
	})
	return FilterNodeClass
}

type _FilterNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FilterNode */
// An interface definition for the [FilterNode] class.
type IFilterNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FilterNode */
	// properties:
	Label() objectivec.IObject
	SetLabel(value objectivec.IObject)
	ResultStates() IMPSNNStateNode
	SetResultStates(value IMPSNNStateNode)
	ResultImage() IMPSNNImageNode
	SetResultImage(value IMPSNNImageNode)
	PaddingPolicy() Padding get set /* not a class type */
	SetPaddingPolicy(value Padding get set /* not a class type */)
	ResultState() IMPSNNStateNode
	SetResultState(value IMPSNNStateNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FilterNode */
	// methods:
	GradientFilter()
	GradientFilterWithSources(gradientImages unsafe.Pointer) IGradientFilterNode
	GradientFilters()
	GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer
	GradientFilterWithSource(gradientImage IImageNode) IGradientFilterNode
	GradientFiltersWithSource(gradientImage IImageNode) unsafe.Pointer
	TrainingGraph()
	TrainingGraphWithSourceGradientNodeHandler(gradientImage IImageNode, nodeHandler GradientNodeBlock /* not a class type */) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FilterNode */
// Alloc allocates a new instance without initialization.
func (fc _FilterNodeClass) Alloc() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FilterNodeClass) New() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterNode) Init() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterNode) Autorelease() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterNode creates a new FilterNode instance.
func NewFilterNode() FilterNode {
	return getFilterNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FilterNode */
// A placeholder node denoting a neural network filter stage.


// A placeholder node denoting a neural network filter stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode
type FilterNode struct {
	objectivec.Object
}

// FilterNodeFrom constructs a [FilterNode] from an unsafe.Pointer.
//
// A placeholder node denoting a neural network filter stage.
func FilterNodeFrom(ptr unsafe.Pointer) FilterNode {
	return FilterNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FilterNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FilterNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FilterNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FilterNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2948052-gradientfilter
func (f_ FilterNode) GradientFilter() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilter"))
}/* debug [instance_methods/method]: GradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2948052-gradientfilterwithsources
func (f_ FilterNode) GradientFilterWithSources(gradientImages unsafe.Pointer) IGradientFilterNode {
	rv := objc.Send[GradientFilterNode](f_.ID, objc.Sel("gradientFilterWithSources:"), gradientImages)
	return rv
}/* debug [instance_methods/method]: GradientFilterWithSources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2951955-gradientfilters
func (f_ FilterNode) GradientFilters() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilters"))
}/* debug [instance_methods/method]: GradientFilters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2951955-gradientfilterswithsources
func (f_ FilterNode) GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSources:"), gradientImages)
	return rv
}/* debug [instance_methods/method]: GradientFiltersWithSources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2953941-gradientfilterwithsource
func (f_ FilterNode) GradientFilterWithSource(gradientImage IImageNode) IGradientFilterNode {
	rv := objc.Send[GradientFilterNode](f_.ID, objc.Sel("gradientFilterWithSource:"), gradientImage)
	return rv
}/* debug [instance_methods/method]: GradientFilterWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2953944-gradientfilterswithsource
func (f_ FilterNode) GradientFiltersWithSource(gradientImage IImageNode) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSource:"), gradientImage)
	return rv
}/* debug [instance_methods/method]: GradientFiltersWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/3020688-traininggraph
func (f_ FilterNode) TrainingGraph() {
	objc.Send[objc.ID](f_.ID, objc.Sel("trainingGraph"))
}/* debug [instance_methods/method]: TrainingGraph */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/3020688-traininggraphwithsourcegradient
func (f_ FilterNode) TrainingGraphWithSourceGradientNodeHandler(gradientImage IImageNode, nodeHandler GradientNodeBlock /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("trainingGraphWithSourceGradient:nodeHandler:"), gradientImage, nodeHandler)
	return rv
}/* debug [instance_methods/method]: TrainingGraphWithSourceGradientNodeHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FilterNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866465-label
func (f_ FilterNode) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866465-label
func (f_ FilterNode) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866486-resultstates
func (f_ FilterNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](f_.ID, objc.Sel("resultStates"))
	return rv
}/* debug [instance_properties/getter]: resultStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866486-resultstates
func (f_ FilterNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultStates:"), value)
}/* debug [instance_properties/setter]: resultStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866492-resultimage
func (f_ FilterNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](f_.ID, objc.Sel("resultImage"))
	return rv
}/* debug [instance_properties/getter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866492-resultimage
func (f_ FilterNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultImage:"), value)
}/* debug [instance_properties/setter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866496-paddingpolicy
func (f_ FilterNode) PaddingPolicy() Padding get set /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("paddingPolicy"))
	return rv
}/* debug [instance_properties/getter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866496-paddingpolicy
func (f_ FilterNode) SetPaddingPolicy(value Padding get set /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPaddingPolicy:"), value)
}/* debug [instance_properties/setter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866503-resultstate
func (f_ FilterNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](f_.ID, objc.Sel("resultState"))
	return rv
}/* debug [instance_properties/getter]: resultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866503-resultstate
func (f_ FilterNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultState:"), value)
}/* debug [instance_properties/setter]: resultState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNFilterNode */



