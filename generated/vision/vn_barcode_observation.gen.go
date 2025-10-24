// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class VNBarcodeObservation */


/* debug [class_header]: Header for VNBarcodeObservation */
// The class instance for the [BarcodeObservation] class.
var (
	BarcodeObservationClass     _BarcodeObservationClass
	BarcodeObservationClassOnce sync.Once
)

func getBarcodeObservationClass() _BarcodeObservationClass {
	BarcodeObservationClassOnce.Do(func() {
		BarcodeObservationClass = _BarcodeObservationClass{objc.GetClass("VNBarcodeObservation")}
	})
	return BarcodeObservationClass
}

type _BarcodeObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BarcodeObservation */
// An interface definition for the [BarcodeObservation] class.
type IBarcodeObservation interface {
	IRectangleObservation
	
/* debug [class_interface_properties]: Properties for BarcodeObservation */
	// properties:
	BarcodeDescriptor() coreimage.BarcodeDescriptor
	IsColorInverted() bool
	IsGS1DataCarrier() bool
	PayloadData() objc.IObject /* cross-framework: NSData */
	PayloadStringValue() objc.IObject /* cross-framework: NSString */
	SupplementalCompositeType() BarcodeCompositeType
	SupplementalPayloadData() objc.IObject /* cross-framework: NSData */
	SupplementalPayloadString() objc.IObject /* cross-framework: NSString */
	Symbology() BarcodeSymbology /* typedef */
	Results() IVNBarcodeObservation
	SetResults(value IVNBarcodeObservation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BarcodeObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BarcodeObservation */
// Alloc allocates a new instance without initialization.
func (bc _BarcodeObservationClass) Alloc() BarcodeObservation {
	rv := objc.Send[BarcodeObservation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BarcodeObservationClass) New() BarcodeObservation {
	rv := objc.Send[BarcodeObservation](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BarcodeObservation) Init() BarcodeObservation {
	rv := objc.Send[BarcodeObservation](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BarcodeObservation) Autorelease() BarcodeObservation {
	rv := objc.Send[BarcodeObservation](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBarcodeObservation creates a new BarcodeObservation instance.
func NewBarcodeObservation() BarcodeObservation {
	return getBarcodeObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BarcodeObservation */
// An object that represents barcode information that an image analysis request detects.
//
// This type of observation results from a . It contains information about the detected barcode, including parsed payload data for supported symbologies.


// An object that represents barcode information that an image analysis request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation
type BarcodeObservation struct {
	RectangleObservation
}

// BarcodeObservationFrom constructs a [BarcodeObservation] from an unsafe.Pointer.
//
// An object that represents barcode information that an image analysis request detects.
func BarcodeObservationFrom(ptr unsafe.Pointer) BarcodeObservation {
	return BarcodeObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BarcodeObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BarcodeObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BarcodeObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BarcodeObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BarcodeObservation */

// An object that describes the low-level details about the barcode and its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/barcodeDescriptor
func (b_ BarcodeObservation) BarcodeDescriptor() coreimage.BarcodeDescriptor {
	rv := objc.Send[coreimage.BarcodeDescriptor](b_.ID, objc.Sel("barcodeDescriptor"))
	return rv
}/* debug [instance_properties/getter]: barcodeDescriptor */


// A Boolean value that indicates whether the barcode is color inverted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isColorInverted
func (b_ BarcodeObservation) IsColorInverted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isColorInverted"))
	return rv
}/* debug [instance_properties/getter]: isColorInverted */


// A Boolean value that indicates whether the barcode carries any global standards data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isGS1DataCarrier
func (b_ BarcodeObservation) IsGS1DataCarrier() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isGS1DataCarrier"))
	return rv
}/* debug [instance_properties/getter]: isGS1DataCarrier */


// The raw data representation of the barcode’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadData
func (b_ BarcodeObservation) PayloadData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("payloadData"))
	return rv
}/* debug [instance_properties/getter]: payloadData */


// A string value that represents the barcode payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadStringValue
func (b_ BarcodeObservation) PayloadStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("payloadStringValue"))
	return rv
}/* debug [instance_properties/getter]: payloadStringValue */


// The supplemental composite type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalCompositeType
func (b_ BarcodeObservation) SupplementalCompositeType() BarcodeCompositeType {
	rv := objc.Send[BarcodeCompositeType](b_.ID, objc.Sel("supplementalCompositeType"))
	return rv
}/* debug [instance_properties/getter]: supplementalCompositeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadData
func (b_ BarcodeObservation) SupplementalPayloadData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("supplementalPayloadData"))
	return rv
}/* debug [instance_properties/getter]: supplementalPayloadData */


// The supplemental code decoded as a string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadString
func (b_ BarcodeObservation) SupplementalPayloadString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("supplementalPayloadString"))
	return rv
}/* debug [instance_properties/getter]: supplementalPayloadString */


// The symbology of the observed barcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/symbology
func (b_ BarcodeObservation) Symbology() BarcodeSymbology /* typedef */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("symbology"))
	return rv
}/* debug [instance_properties/getter]: symbology */


// The results of a barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (b_ BarcodeObservation) Results() IVNBarcodeObservation {
	rv := objc.Send[BarcodeObservation](b_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The results of a barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (b_ BarcodeObservation) SetResults(value IVNBarcodeObservation) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResults:"), value)
}/* debug [instance_properties/setter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNBarcodeObservation */



