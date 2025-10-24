// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNYOLOLoss] class.
type ICNNYOLOLoss interface {
	ICNNKernel
	

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


	

	// methods:
	Encode()
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */)
	EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage
	EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976480-initwithcoder
func NewCNNYOLOLossWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNYOLOLoss {
	instance := getCNNYOLOLossClass().Alloc()
	rv := objc.Send[CNNYOLOLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976481-initwithdevice
func NewCNNYOLOLossWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNYOLOLossDescriptor) CNNYOLOLoss {
	instance := getCNNYOLOLossClass().Alloc()
	rv := objc.Send[CNNYOLOLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encode
func (c_ CNNYOLOLoss) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976476-encodebatchtocommandbuffer
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), commandBuffer, sourceImage, labels)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976477-encodebatchtocommandbuffer
func (c_ CNNYOLOLoss) EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:destinationImages:"), commandBuffer, sourceImage, labels, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976478-encodetocommandbuffer
func (c_ CNNYOLOLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), commandBuffer, sourceImage, labels)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976479-encodetocommandbuffer
func (c_ CNNYOLOLoss) EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:destinationImage:"), commandBuffer, sourceImage, labels, destinationImage)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976475-anchorboxes
func (c_ CNNYOLOLoss) AnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("anchorBoxes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976475-anchorboxes
func (c_ CNNYOLOLoss) SetAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAnchorBoxes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976482-lossclasses
func (c_ CNNYOLOLoss) LossClasses() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976482-lossclasses
func (c_ CNNYOLOLoss) SetLossClasses(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976483-lossconfidence
func (c_ CNNYOLOLoss) LossConfidence() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossConfidence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976483-lossconfidence
func (c_ CNNYOLOLoss) SetLossConfidence(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossConfidence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976484-losswh
func (c_ CNNYOLOLoss) LossWH() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossWH"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976484-losswh
func (c_ CNNYOLOLoss) SetLossWH(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossWH:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976485-lossxy
func (c_ CNNYOLOLoss) LossXY() IMPSCNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("lossXY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976485-lossxy
func (c_ CNNYOLOLoss) SetLossXY(value IMPSCNNLoss) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossXY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976486-maxiouforobjectabsence
func (c_ CNNYOLOLoss) MaxIOUForObjectAbsence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976486-maxiouforobjectabsence
func (c_ CNNYOLOLoss) SetMaxIOUForObjectAbsence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxIOUForObjectAbsence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976487-miniouforobjectpresence
func (c_ CNNYOLOLoss) MinIOUForObjectPresence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minIOUForObjectPresence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976487-miniouforobjectpresence
func (c_ CNNYOLOLoss) SetMinIOUForObjectPresence(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinIOUForObjectPresence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976488-numberofanchorboxes
func (c_ CNNYOLOLoss) NumberOfAnchorBoxes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfAnchorBoxes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976488-numberofanchorboxes
func (c_ CNNYOLOLoss) SetNumberOfAnchorBoxes(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfAnchorBoxes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976489-reductiontype
func (c_ CNNYOLOLoss) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976489-reductiontype
func (c_ CNNYOLOLoss) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976490-scaleclass
func (c_ CNNYOLOLoss) ScaleClass() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleClass"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976490-scaleclass
func (c_ CNNYOLOLoss) SetScaleClass(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleClass:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976491-scalenoobject
func (c_ CNNYOLOLoss) ScaleNoObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleNoObject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976491-scalenoobject
func (c_ CNNYOLOLoss) SetScaleNoObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleNoObject:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976492-scaleobject
func (c_ CNNYOLOLoss) ScaleObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleObject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976492-scaleobject
func (c_ CNNYOLOLoss) SetScaleObject(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleObject:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976493-scalewh
func (c_ CNNYOLOLoss) ScaleWH() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleWH"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976493-scalewh
func (c_ CNNYOLOLoss) SetScaleWH(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleWH:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976494-scalexy
func (c_ CNNYOLOLoss) ScaleXY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleXY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/2976494-scalexy
func (c_ CNNYOLOLoss) SetScaleXY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleXY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/3547983-reduceacrossbatch
func (c_ CNNYOLOLoss) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololoss/3547983-reduceacrossbatch
func (c_ CNNYOLOLoss) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}







