// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCYOLOLossDescriptor */


/* debug [class_header]: Header for MLCYOLOLossDescriptor */
// The class instance for the [CYOLOLossDescriptor] class.
var (
	CYOLOLossDescriptorClass     _CYOLOLossDescriptorClass
	CYOLOLossDescriptorClassOnce sync.Once
)

func getCYOLOLossDescriptorClass() _CYOLOLossDescriptorClass {
	CYOLOLossDescriptorClassOnce.Do(func() {
		CYOLOLossDescriptorClass = _CYOLOLossDescriptorClass{objc.GetClass("MLCYOLOLossDescriptor")}
	})
	return CYOLOLossDescriptorClass
}

type _CYOLOLossDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CYOLOLossDescriptor */
// An interface definition for the [CYOLOLossDescriptor] class.
type ICYOLOLossDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CYOLOLossDescriptor */
	// properties:
	AnchorBoxCount() uint
	AnchorBoxes() objc.IObject /* cross-framework: NSData */
	MaximumIOUForObjectAbsence() float32
	SetMaximumIOUForObjectAbsence(value float32)
	MinimumIOUForObjectPresence() float32
	SetMinimumIOUForObjectPresence(value float32)
	ScaleClassLoss() float32
	SetScaleClassLoss(value float32)
	ScaleNoObjectConfidenceLoss() float32
	SetScaleNoObjectConfidenceLoss(value float32)
	ScaleObjectConfidenceLoss() float32
	SetScaleObjectConfidenceLoss(value float32)
	ScaleSpatialPositionLoss() float32
	SetScaleSpatialPositionLoss(value float32)
	ScaleSpatialSizeLoss() float32
	SetScaleSpatialSizeLoss(value float32)
	ShouldRescore() bool
	SetShouldRescore(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CYOLOLossDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CYOLOLossDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CYOLOLossDescriptorClass) Alloc() CYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CYOLOLossDescriptorClass) New() CYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CYOLOLossDescriptor) Init() CYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CYOLOLossDescriptor) Autorelease() CYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCYOLOLossDescriptor creates a new CYOLOLossDescriptor instance.
func NewCYOLOLossDescriptor() CYOLOLossDescriptor {
	return getCYOLOLossDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CYOLOLossDescriptor */
// The configuration object you use to create the YOLO loss layer.


// The configuration object you use to create the YOLO loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor
type CYOLOLossDescriptor struct {
	objectivec.Object
}

// CYOLOLossDescriptorFrom constructs a [CYOLOLossDescriptor] from an unsafe.Pointer.
//
// The configuration object you use to create the YOLO loss layer.
func CYOLOLossDescriptorFrom(ptr unsafe.Pointer) CYOLOLossDescriptor {
	return CYOLOLossDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CYOLOLossDescriptor */

// Creates a YOLO loss filter descriptor with the anchor box data and number of anchor boxes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/init(anchorBoxes:anchorBoxCount:)
func NewCYOLOLossDescriptorWithAnchorBoxesAnchorBoxCount(anchorBoxes objc.IObject /* cross-framework: NSData */, anchorBoxCount uint) CYOLOLossDescriptor {
	rv := objc.Send[CYOLOLossDescriptor](objc.ID(getCYOLOLossDescriptorClass().class), objc.Sel("descriptorWithAnchorBoxes:anchorBoxCount:"), anchorBoxes, anchorBoxCount)
	return rv
}/* debug [class_init_methods/constructor]: NewCYOLOLossDescriptorWithAnchorBoxesAnchorBoxCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CYOLOLossDescriptor */

// Creates a YOLO loss filter descriptor with the anchor box data and number of anchor boxes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/init(anchorBoxes:anchorBoxCount:)
func (cc _CYOLOLossDescriptorClass) DescriptorWithAnchorBoxesAnchorBoxCount(anchorBoxes objc.IObject /* cross-framework: NSData */, anchorBoxCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithAnchorBoxes:anchorBoxCount:"), anchorBoxes, anchorBoxCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithAnchorBoxesAnchorBoxCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CYOLOLossDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CYOLOLossDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CYOLOLossDescriptor */

// The number of anchor boxes you use to detect an object in each grid cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/anchorBoxCount
func (c_ CYOLOLossDescriptor) AnchorBoxCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("anchorBoxCount"))
	return rv
}/* debug [instance_properties/getter]: anchorBoxCount */


// Data that contains the width and height for all the anchor boxes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/anchorBoxes
func (c_ CYOLOLossDescriptor) AnchorBoxes() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("anchorBoxes"))
	return rv
}/* debug [instance_properties/getter]: anchorBoxes */


// The negative intersection over union (IOU).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/maximumIOUForObjectAbsence
func (c_ CYOLOLossDescriptor) MaximumIOUForObjectAbsence() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maximumIOUForObjectAbsence"))
	return rv
}/* debug [instance_properties/getter]: maximumIOUForObjectAbsence */


// The negative intersection over union (IOU).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/maximumIOUForObjectAbsence
func (c_ CYOLOLossDescriptor) SetMaximumIOUForObjectAbsence(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumIOUForObjectAbsence:"), value)
}/* debug [instance_properties/setter]: maximumIOUForObjectAbsence */


// The positive intersection over union (IOU).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/minimumIOUForObjectPresence
func (c_ CYOLOLossDescriptor) MinimumIOUForObjectPresence() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("minimumIOUForObjectPresence"))
	return rv
}/* debug [instance_properties/getter]: minimumIOUForObjectPresence */


// The positive intersection over union (IOU).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/minimumIOUForObjectPresence
func (c_ CYOLOLossDescriptor) SetMinimumIOUForObjectPresence(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumIOUForObjectPresence:"), value)
}/* debug [instance_properties/setter]: minimumIOUForObjectPresence */


// The scale factor you use for loss when there are no object classes, and for loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleClassLoss
func (c_ CYOLOLossDescriptor) ScaleClassLoss() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scaleClassLoss"))
	return rv
}/* debug [instance_properties/getter]: scaleClassLoss */


// The scale factor you use for loss when there are no object classes, and for loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleClassLoss
func (c_ CYOLOLossDescriptor) SetScaleClassLoss(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleClassLoss:"), value)
}/* debug [instance_properties/setter]: scaleClassLoss */


// The scale factor you use for no object confidence loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleNoObjectConfidenceLoss
func (c_ CYOLOLossDescriptor) ScaleNoObjectConfidenceLoss() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scaleNoObjectConfidenceLoss"))
	return rv
}/* debug [instance_properties/getter]: scaleNoObjectConfidenceLoss */


// The scale factor you use for no object confidence loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleNoObjectConfidenceLoss
func (c_ CYOLOLossDescriptor) SetScaleNoObjectConfidenceLoss(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleNoObjectConfidenceLoss:"), value)
}/* debug [instance_properties/setter]: scaleNoObjectConfidenceLoss */


// The scale factor you use for object confidence loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleObjectConfidenceLoss
func (c_ CYOLOLossDescriptor) ScaleObjectConfidenceLoss() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scaleObjectConfidenceLoss"))
	return rv
}/* debug [instance_properties/getter]: scaleObjectConfidenceLoss */


// The scale factor you use for object confidence loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleObjectConfidenceLoss
func (c_ CYOLOLossDescriptor) SetScaleObjectConfidenceLoss(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleObjectConfidenceLoss:"), value)
}/* debug [instance_properties/setter]: scaleObjectConfidenceLoss */


// The scale factor you use for spatial position loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleSpatialPositionLoss
func (c_ CYOLOLossDescriptor) ScaleSpatialPositionLoss() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scaleSpatialPositionLoss"))
	return rv
}/* debug [instance_properties/getter]: scaleSpatialPositionLoss */


// The scale factor you use for spatial position loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleSpatialPositionLoss
func (c_ CYOLOLossDescriptor) SetScaleSpatialPositionLoss(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleSpatialPositionLoss:"), value)
}/* debug [instance_properties/setter]: scaleSpatialPositionLoss */


// The scale factor you use for spatial size loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleSpatialSizeLoss
func (c_ CYOLOLossDescriptor) ScaleSpatialSizeLoss() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scaleSpatialSizeLoss"))
	return rv
}/* debug [instance_properties/getter]: scaleSpatialSizeLoss */


// The scale factor you use for spatial size loss and loss gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/scaleSpatialSizeLoss
func (c_ CYOLOLossDescriptor) SetScaleSpatialSizeLoss(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleSpatialSizeLoss:"), value)
}/* debug [instance_properties/setter]: scaleSpatialSizeLoss */


// A Boolean that indicates whether the layer scales the object confidence loss by the intersection over union (IOU) overlap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/shouldRescore
func (c_ CYOLOLossDescriptor) ShouldRescore() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldRescore"))
	return rv
}/* debug [instance_properties/getter]: shouldRescore */


// A Boolean that indicates whether the layer scales the object confidence loss by the intersection over union (IOU) overlap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossDescriptor/shouldRescore
func (c_ CYOLOLossDescriptor) SetShouldRescore(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldRescore:"), value)
}/* debug [instance_properties/setter]: shouldRescore */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCYOLOLossDescriptor */


