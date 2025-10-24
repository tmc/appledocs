// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNYOLOLoss */


/* debug [class_header]: Header for MPSCNNYOLOLoss */
// The class instance for the [CNNYOLOLoss] class.
var (
	CNNYOLOLossClass     _CNNYOLOLossClass
	CNNYOLOLossClassOnce sync.Once
)

func getCNNYOLOLossClass() _CNNYOLOLossClass {
	CNNYOLOLossClassOnce.Do(func() {
		CNNYOLOLossClass = _CNNYOLOLossClass{objc.GetClass("MPSCNNYOLOLoss")}
	})
	return CNNYOLOLossClass
}

type _CNNYOLOLossClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNYOLOLoss */
// An interface definition for the [CNNYOLOLoss] class.
type ICNNYOLOLoss interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNYOLOLoss */
	// properties:
	AnchorBoxes() objectivec.IObject
	SetAnchorBoxes(value objectivec.IObject)
	LossClasses() IMPSCNNLoss
	SetLossClasses(value IMPSCNNLoss)
	LossConfidence() IMPSCNNLoss
	SetLossConfidence(value IMPSCNNLoss)
	LossWH() IMPSCNNLoss
	SetLossWH(value IMPSCNNLoss)
	LossXY() IMPSCNNLoss
	SetLossXY(value IMPSCNNLoss)
	MaxIOUForObjectAbsence() objectivec.IObject
	SetMaxIOUForObjectAbsence(value objectivec.IObject)
	MinIOUForObjectPresence() objectivec.IObject
	SetMinIOUForObjectPresence(value objectivec.IObject)
	NumberOfAnchorBoxes() objectivec.IObject
	SetNumberOfAnchorBoxes(value objectivec.IObject)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
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

	
/* debug [class_interface_methods]: Methods for CNNYOLOLoss */
	// methods:
	Encode()
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */)
	EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage
	EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNYOLOLoss */
// Alloc allocates a new instance without initialization.
func (cc _CNNYOLOLossClass) Alloc() CNNYOLOLoss {
	rv := objc.Send[CNNYOLOLoss](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNYOLOLossClass) New() CNNYOLOLoss {
	rv := objc.Send[CNNYOLOLoss](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNYOLOLoss) Init() CNNYOLOLoss {
	rv := objc.Send[CNNYOLOLoss](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNYOLOLoss) Autorelease() CNNYOLOLoss {
	rv := objc.Send[CNNYOLOLoss](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNYOLOLoss creates a new CNNYOLOLoss instance.
func NewCNNYOLOLoss() CNNYOLOLoss {
	return getCNNYOLOLossClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNYOLOLoss */
// A kernel that computes the YOLO loss and loss gradient between specified predictions and labels.


// A kernel that computes the YOLO loss and loss gradient between specified predictions and labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLoss
type CNNYOLOLoss struct {
	CNNKernel
}

// CNNYOLOLossFrom constructs a [CNNYOLOLoss] from an unsafe.Pointer.
//
// A kernel that computes the YOLO loss and loss gradient between specified predictions and labels.
func CNNYOLOLossFrom(ptr unsafe.Pointer) CNNYOLOLoss {
	return CNNYOLOLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNYOLOLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976480-initwithcoder
func NewCNNYOLOLossWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNYOLOLoss {
	instance := getCNNYOLOLossClass().Alloc()
	rv := objc.Send[CNNYOLOLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNYOLOLossWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976481-initwithdevice
func NewCNNYOLOLossWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNYOLOLossDescriptor) CNNYOLOLoss {
	instance := getCNNYOLOLossClass().Alloc()
	rv := objc.Send[CNNYOLOLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNYOLOLossWithDeviceLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNYOLOLoss */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNYOLOLoss */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNYOLOLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encode
func (c_ CNNYOLOLoss) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encodebatchtocommandbuffer
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), commandBuffer, sourceImage, labels)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976477-encodebatchtocommandbuffer
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:destinationImages:"), commandBuffer, sourceImage, labels, destinationImage)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976478-encodetocommandbuffer
func (c_ CNNYOLOLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), commandBuffer, sourceImage, labels)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976479-encodetocommandbuffer
func (c_ CNNYOLOLoss) EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:destinationImage:"), commandBuffer, sourceImage, labels, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageLabelsDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNYOLOLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976475-anchorboxes
func (c_ CNNYOLOLoss) AnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("anchorBoxes"))
	return rv
}/* debug [instance_properties/getter]: anchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976475-anchorboxes
func (c_ CNNYOLOLoss) SetAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAnchorBoxes:"), value)
}/* debug [instance_properties/setter]: anchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976482-lossclasses
func (c_ CNNYOLOLoss) LossClasses() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossClasses"))
	return rv
}/* debug [instance_properties/getter]: lossClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976482-lossclasses
func (c_ CNNYOLOLoss) SetLossClasses(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossClasses:"), value)
}/* debug [instance_properties/setter]: lossClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976483-lossconfidence
func (c_ CNNYOLOLoss) LossConfidence() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossConfidence"))
	return rv
}/* debug [instance_properties/getter]: lossConfidence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976483-lossconfidence
func (c_ CNNYOLOLoss) SetLossConfidence(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossConfidence:"), value)
}/* debug [instance_properties/setter]: lossConfidence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976484-losswh
func (c_ CNNYOLOLoss) LossWH() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossWH"))
	return rv
}/* debug [instance_properties/getter]: lossWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976484-losswh
func (c_ CNNYOLOLoss) SetLossWH(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossWH:"), value)
}/* debug [instance_properties/setter]: lossWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976485-lossxy
func (c_ CNNYOLOLoss) LossXY() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossXY"))
	return rv
}/* debug [instance_properties/getter]: lossXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976485-lossxy
func (c_ CNNYOLOLoss) SetLossXY(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossXY:"), value)
}/* debug [instance_properties/setter]: lossXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976486-maxiouforobjectabsence
func (c_ CNNYOLOLoss) MaxIOUForObjectAbsence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}/* debug [instance_properties/getter]: maxIOUForObjectAbsence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976486-maxiouforobjectabsence
func (c_ CNNYOLOLoss) SetMaxIOUForObjectAbsence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxIOUForObjectAbsence:"), value)
}/* debug [instance_properties/setter]: maxIOUForObjectAbsence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976487-miniouforobjectpresence
func (c_ CNNYOLOLoss) MinIOUForObjectPresence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minIOUForObjectPresence"))
	return rv
}/* debug [instance_properties/getter]: minIOUForObjectPresence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976487-miniouforobjectpresence
func (c_ CNNYOLOLoss) SetMinIOUForObjectPresence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinIOUForObjectPresence:"), value)
}/* debug [instance_properties/setter]: minIOUForObjectPresence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976488-numberofanchorboxes
func (c_ CNNYOLOLoss) NumberOfAnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfAnchorBoxes"))
	return rv
}/* debug [instance_properties/getter]: numberOfAnchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976488-numberofanchorboxes
func (c_ CNNYOLOLoss) SetNumberOfAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfAnchorBoxes:"), value)
}/* debug [instance_properties/setter]: numberOfAnchorBoxes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976489-reductiontype
func (c_ CNNYOLOLoss) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976489-reductiontype
func (c_ CNNYOLOLoss) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976490-scaleclass
func (c_ CNNYOLOLoss) ScaleClass() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleClass"))
	return rv
}/* debug [instance_properties/getter]: scaleClass */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976490-scaleclass
func (c_ CNNYOLOLoss) SetScaleClass(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleClass:"), value)
}/* debug [instance_properties/setter]: scaleClass */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976491-scalenoobject
func (c_ CNNYOLOLoss) ScaleNoObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleNoObject"))
	return rv
}/* debug [instance_properties/getter]: scaleNoObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976491-scalenoobject
func (c_ CNNYOLOLoss) SetScaleNoObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleNoObject:"), value)
}/* debug [instance_properties/setter]: scaleNoObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976492-scaleobject
func (c_ CNNYOLOLoss) ScaleObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleObject"))
	return rv
}/* debug [instance_properties/getter]: scaleObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976492-scaleobject
func (c_ CNNYOLOLoss) SetScaleObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleObject:"), value)
}/* debug [instance_properties/setter]: scaleObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976493-scalewh
func (c_ CNNYOLOLoss) ScaleWH() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleWH"))
	return rv
}/* debug [instance_properties/getter]: scaleWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976493-scalewh
func (c_ CNNYOLOLoss) SetScaleWH(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleWH:"), value)
}/* debug [instance_properties/setter]: scaleWH */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976494-scalexy
func (c_ CNNYOLOLoss) ScaleXY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleXY"))
	return rv
}/* debug [instance_properties/getter]: scaleXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976494-scalexy
func (c_ CNNYOLOLoss) SetScaleXY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleXY:"), value)
}/* debug [instance_properties/setter]: scaleXY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/3547983-reduceacrossbatch
func (c_ CNNYOLOLoss) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/3547983-reduceacrossbatch
func (c_ CNNYOLOLoss) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNYOLOLoss */


