// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLossDescriptor */


/* debug [class_header]: Header for MPSCNNLossDescriptor */
// The class instance for the [CNNLossDescriptor] class.
var (
	CNNLossDescriptorClass     _CNNLossDescriptorClass
	CNNLossDescriptorClassOnce sync.Once
)

func getCNNLossDescriptorClass() _CNNLossDescriptorClass {
	CNNLossDescriptorClassOnce.Do(func() {
		CNNLossDescriptorClass = _CNNLossDescriptorClass{objc.GetClass("MPSCNNLossDescriptor")}
	})
	return CNNLossDescriptorClass
}

type _CNNLossDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLossDescriptor */
// An interface definition for the [CNNLossDescriptor] class.
type ICNNLossDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNNLossDescriptor */
	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	LabelSmoothing() objectivec.IObject
	SetLabelSmoothing(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	LossType() CNNLossType get set /* not a class type */
	SetLossType(value CNNLossType get set /* not a class type */)
	NumberOfClasses() objectivec.IObject
	SetNumberOfClasses(value objectivec.IObject)
	ReductionType() CNNReductionType get set /* not a class type */
	SetReductionType(value CNNReductionType get set /* not a class type */)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLossDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLossDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CNNLossDescriptorClass) Alloc() CNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLossDescriptorClass) New() CNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLossDescriptor) Init() CNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLossDescriptor) Autorelease() CNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLossDescriptor creates a new CNNLossDescriptor instance.
func NewCNNLossDescriptor() CNNLossDescriptor {
	return getCNNLossDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLossDescriptor */
// An object that specifies properties used by a loss kernel.


// An object that specifies properties used by a loss kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor
type CNNLossDescriptor struct {
	objectivec.Object
}

// CNNLossDescriptorFrom constructs a [CNNLossDescriptor] from an unsafe.Pointer.
//
// An object that specifies properties used by a loss kernel.
func CNNLossDescriptorFrom(ptr unsafe.Pointer) CNNLossDescriptor {
	return CNNLossDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLossDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLossDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942373-cnnlossdescriptorwithtype
func (cc _CNNLossDescriptorClass) CnnLossDescriptorWithTypeReductionType(lossType CNNLossType, reductionType CNNReductionType) ICNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](objc.ID(cc.class), objc.Sel("cnnLossDescriptorWithType:reductionType:"), lossType, reductionType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CnnLossDescriptorWithTypeReductionType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLossDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLossDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLossDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942362-epsilon
func (c_ CNNLossDescriptor) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942362-epsilon
func (c_ CNNLossDescriptor) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942367-weight
func (c_ CNNLossDescriptor) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942367-weight
func (c_ CNNLossDescriptor) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942369-labelsmoothing
func (c_ CNNLossDescriptor) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942369-labelsmoothing
func (c_ CNNLossDescriptor) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942378-delta
func (c_ CNNLossDescriptor) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942378-delta
func (c_ CNNLossDescriptor) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942381-losstype
func (c_ CNNLossDescriptor) LossType() CNNLossType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942381-losstype
func (c_ CNNLossDescriptor) SetLossType(value CNNLossType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942382-numberofclasses
func (c_ CNNLossDescriptor) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942382-numberofclasses
func (c_ CNNLossDescriptor) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942388-reductiontype
func (c_ CNNLossDescriptor) ReductionType() CNNReductionType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942388-reductiontype
func (c_ CNNLossDescriptor) SetReductionType(value CNNReductionType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/3547982-reduceacrossbatch
func (c_ CNNLossDescriptor) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/3547982-reduceacrossbatch
func (c_ CNNLossDescriptor) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLossDescriptor */



