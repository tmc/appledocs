// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNYOLOLossDescriptor */


/* debug [class_header]: Header for MPSCNNYOLOLossDescriptor */
// The class instance for the [CNNYOLOLossDescriptor] class.
var (
	CNNYOLOLossDescriptorClass     _CNNYOLOLossDescriptorClass
	CNNYOLOLossDescriptorClassOnce sync.Once
)

func getCNNYOLOLossDescriptorClass() _CNNYOLOLossDescriptorClass {
	CNNYOLOLossDescriptorClassOnce.Do(func() {
		CNNYOLOLossDescriptorClass = _CNNYOLOLossDescriptorClass{objc.GetClass("MPSCNNYOLOLossDescriptor")}
	})
	return CNNYOLOLossDescriptorClass
}

type _CNNYOLOLossDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNYOLOLossDescriptor */
// An interface definition for the [CNNYOLOLossDescriptor] class.
type ICNNYOLOLossDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNNYOLOLossDescriptor */
	// properties:
	WhLossDescriptor() IMPSCNNLossDescriptor
	SetWhLossDescriptor(value IMPSCNNLossDescriptor)
	XyLossDescriptor() IMPSCNNLossDescriptor
	SetXyLossDescriptor(value IMPSCNNLossDescriptor)
	AnchorBoxes() objectivec.IObject
	SetAnchorBoxes(value objectivec.IObject)
	ClassesLossDescriptor() IMPSCNNLossDescriptor
	SetClassesLossDescriptor(value IMPSCNNLossDescriptor)
	ConfidenceLossDescriptor() IMPSCNNLossDescriptor
	SetConfidenceLossDescriptor(value IMPSCNNLossDescriptor)
	MaxIOUForObjectAbsence() objectivec.IObject
	SetMaxIOUForObjectAbsence(value objectivec.IObject)
	MinIOUForObjectPresence() objectivec.IObject
	SetMinIOUForObjectPresence(value objectivec.IObject)
	NumberOfAnchorBoxes() objectivec.IObject
	SetNumberOfAnchorBoxes(value objectivec.IObject)
	ReductionType() CNNReductionType get set /* not a class type */
	SetReductionType(value CNNReductionType get set /* not a class type */)
	Rescore() objectivec.IObject
	SetRescore(value objectivec.IObject)
	ScaleClass() objectivec.IObject
	SetScaleClass(value objectivec.IObject)
	ScaleNoObject() objectivec.IObject
	SetScaleNoObject(value objectivec.IObject)
	ScaleObject() objectivec.IObject
	SetScaleObject(value objectivec.IObject)
	ScaleWH() objectivec.IObject
	SetScaleWH(value objectivec.IObject)
	ScaleXY() objectivec.IObject
	SetScaleXY(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNYOLOLossDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNYOLOLossDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CNNYOLOLossDescriptorClass) Alloc() CNNYOLOLossDescriptor {
	rv := objc.Send[CNNYOLOLossDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNYOLOLossDescriptorClass) New() CNNYOLOLossDescriptor {
	rv := objc.Send[CNNYOLOLossDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNYOLOLossDescriptor) Init() CNNYOLOLossDescriptor {
	rv := objc.Send[CNNYOLOLossDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNYOLOLossDescriptor) Autorelease() CNNYOLOLossDescriptor {
	rv := objc.Send[CNNYOLOLossDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNYOLOLossDescriptor creates a new CNNYOLOLossDescriptor instance.
func NewCNNYOLOLossDescriptor() CNNYOLOLossDescriptor {
	return getCNNYOLOLossDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNYOLOLossDescriptor */
// An object that specifies properties used by a YOLO loss kernel.


// An object that specifies properties used by a YOLO loss kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor
type CNNYOLOLossDescriptor struct {
	objectivec.Object
}

// CNNYOLOLossDescriptorFrom constructs a [CNNYOLOLossDescriptor] from an unsafe.Pointer.
//
// An object that specifies properties used by a YOLO loss kernel.
func CNNYOLOLossDescriptorFrom(ptr unsafe.Pointer) CNNYOLOLossDescriptor {
	return CNNYOLOLossDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNYOLOLossDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNYOLOLossDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976500-cnnlossdescriptor
func (cc _CNNYOLOLossDescriptorClass) CnnLossDescriptor() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("cnnLossDescriptor"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CnnLossDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976500-cnnlossdescriptorwithxylosstype
func (cc _CNNYOLOLossDescriptorClass) CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes(XYLossType CNNLossType, WHLossType CNNLossType, confidenceLossType CNNLossType, classesLossType CNNLossType, reductionType CNNReductionType, anchorBoxes objc.IObject /* cross-framework: Data */, numberOfAnchorBoxes uint) ICNNYOLOLossDescriptor {
	rv := objc.Send[CNNYOLOLossDescriptor](objc.ID(cc.class), objc.Sel("cnnLossDescriptorWithXYLossType:WHLossType:confidenceLossType:classesLossType:reductionType:anchorBoxes:numberOfAnchorBoxes:"), XYLossType, WHLossType, confidenceLossType, classesLossType, reductionType, anchorBoxes, numberOfAnchorBoxes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNYOLOLossDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNYOLOLossDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNYOLOLossDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976496-whlossdescriptor
func (c_ CNNYOLOLossDescriptor) WhLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("whLossDescriptor"))
	return rv
}/* debug [instance_properties/getter]: whLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976496-whlossdescriptor
func (c_ CNNYOLOLossDescriptor) SetWhLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhLossDescriptor:"), value)
}/* debug [instance_properties/setter]: whLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976497-xylossdescriptor
func (c_ CNNYOLOLossDescriptor) XyLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("xyLossDescriptor"))
	return rv
}/* debug [instance_properties/getter]: xyLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976497-xylossdescriptor
func (c_ CNNYOLOLossDescriptor) SetXyLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setXyLossDescriptor:"), value)
}/* debug [instance_properties/setter]: xyLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976498-anchorboxes
func (c_ CNNYOLOLossDescriptor) AnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("anchorBoxes"))
	return rv
}/* debug [instance_properties/getter]: anchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976498-anchorboxes
func (c_ CNNYOLOLossDescriptor) SetAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAnchorBoxes:"), value)
}/* debug [instance_properties/setter]: anchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976499-classeslossdescriptor
func (c_ CNNYOLOLossDescriptor) ClassesLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("classesLossDescriptor"))
	return rv
}/* debug [instance_properties/getter]: classesLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976499-classeslossdescriptor
func (c_ CNNYOLOLossDescriptor) SetClassesLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClassesLossDescriptor:"), value)
}/* debug [instance_properties/setter]: classesLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976501-confidencelossdescriptor
func (c_ CNNYOLOLossDescriptor) ConfidenceLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[CNNLossDescriptor](c_.ID, objc.Sel("confidenceLossDescriptor"))
	return rv
}/* debug [instance_properties/getter]: confidenceLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976501-confidencelossdescriptor
func (c_ CNNYOLOLossDescriptor) SetConfidenceLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfidenceLossDescriptor:"), value)
}/* debug [instance_properties/setter]: confidenceLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976502-maxiouforobjectabsence
func (c_ CNNYOLOLossDescriptor) MaxIOUForObjectAbsence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}/* debug [instance_properties/getter]: maxIOUForObjectAbsence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976502-maxiouforobjectabsence
func (c_ CNNYOLOLossDescriptor) SetMaxIOUForObjectAbsence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxIOUForObjectAbsence:"), value)
}/* debug [instance_properties/setter]: maxIOUForObjectAbsence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976503-miniouforobjectpresence
func (c_ CNNYOLOLossDescriptor) MinIOUForObjectPresence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minIOUForObjectPresence"))
	return rv
}/* debug [instance_properties/getter]: minIOUForObjectPresence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976503-miniouforobjectpresence
func (c_ CNNYOLOLossDescriptor) SetMinIOUForObjectPresence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinIOUForObjectPresence:"), value)
}/* debug [instance_properties/setter]: minIOUForObjectPresence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976504-numberofanchorboxes
func (c_ CNNYOLOLossDescriptor) NumberOfAnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfAnchorBoxes"))
	return rv
}/* debug [instance_properties/getter]: numberOfAnchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976504-numberofanchorboxes
func (c_ CNNYOLOLossDescriptor) SetNumberOfAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfAnchorBoxes:"), value)
}/* debug [instance_properties/setter]: numberOfAnchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976505-reductiontype
func (c_ CNNYOLOLossDescriptor) ReductionType() CNNReductionType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976505-reductiontype
func (c_ CNNYOLOLossDescriptor) SetReductionType(value CNNReductionType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976506-rescore
func (c_ CNNYOLOLossDescriptor) Rescore() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("rescore"))
	return rv
}/* debug [instance_properties/getter]: rescore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976506-rescore
func (c_ CNNYOLOLossDescriptor) SetRescore(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRescore:"), value)
}/* debug [instance_properties/setter]: rescore */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976507-scaleclass
func (c_ CNNYOLOLossDescriptor) ScaleClass() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleClass"))
	return rv
}/* debug [instance_properties/getter]: scaleClass */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976507-scaleclass
func (c_ CNNYOLOLossDescriptor) SetScaleClass(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleClass:"), value)
}/* debug [instance_properties/setter]: scaleClass */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976508-scalenoobject
func (c_ CNNYOLOLossDescriptor) ScaleNoObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleNoObject"))
	return rv
}/* debug [instance_properties/getter]: scaleNoObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976508-scalenoobject
func (c_ CNNYOLOLossDescriptor) SetScaleNoObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleNoObject:"), value)
}/* debug [instance_properties/setter]: scaleNoObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976509-scaleobject
func (c_ CNNYOLOLossDescriptor) ScaleObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleObject"))
	return rv
}/* debug [instance_properties/getter]: scaleObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976509-scaleobject
func (c_ CNNYOLOLossDescriptor) SetScaleObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleObject:"), value)
}/* debug [instance_properties/setter]: scaleObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976510-scalewh
func (c_ CNNYOLOLossDescriptor) ScaleWH() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleWH"))
	return rv
}/* debug [instance_properties/getter]: scaleWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976510-scalewh
func (c_ CNNYOLOLossDescriptor) SetScaleWH(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleWH:"), value)
}/* debug [instance_properties/setter]: scaleWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976511-scalexy
func (c_ CNNYOLOLossDescriptor) ScaleXY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleXY"))
	return rv
}/* debug [instance_properties/getter]: scaleXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/2976511-scalexy
func (c_ CNNYOLOLossDescriptor) SetScaleXY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleXY:"), value)
}/* debug [instance_properties/setter]: scaleXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/3547984-reduceacrossbatch
func (c_ CNNYOLOLossDescriptor) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossdescriptor/3547984-reduceacrossbatch
func (c_ CNNYOLOLossDescriptor) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNYOLOLossDescriptor */



