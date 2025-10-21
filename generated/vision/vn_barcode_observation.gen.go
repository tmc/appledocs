// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [BarcodeObservation] class.
type IBarcodeObservation interface {
	IRectangleObservation
}

// An object that represents barcode information that an image analysis request detects.
//
// This type of observation results from a . It contains information about the detected barcode, including parsed payload data for supported symbologies.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BarcodeObservationClass) Alloc() BarcodeObservation {
	rv := objc.Send[BarcodeObservation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An object that describes the low-level details about the barcode and its data.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/barcodeDescriptor
func (b_ BarcodeObservation) BarcodeDescriptor() coreimage.BarcodeDescriptor {
	rv := objc.Send[coreimage.BarcodeDescriptor](b_.ID, objc.Sel("barcodeDescriptor"))
	return rv
}

// A Boolean value that indicates whether the barcode is color inverted.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isColorInverted
func (b_ BarcodeObservation) IsColorInverted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isColorInverted"))
	return rv
}

// A Boolean value that indicates whether the barcode carries any global standards data.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isGS1DataCarrier
func (b_ BarcodeObservation) IsGS1DataCarrier() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isGS1DataCarrier"))
	return rv
}

// The raw data representation of the barcode’s payload.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadData
func (b_ BarcodeObservation) PayloadData() foundation.NSData {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("payloadData"))
	return rv
}

// A string value that represents the barcode payload.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadStringValue
func (b_ BarcodeObservation) PayloadStringValue() appkit.string {
	rv := objc.Send[appkit.string](b_.ID, objc.Sel("payloadStringValue"))
	return rv
}

// The supplemental composite type.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalCompositeType
func (b_ BarcodeObservation) SupplementalCompositeType() BarcodeCompositeType {
	rv := objc.Send[BarcodeCompositeType](b_.ID, objc.Sel("supplementalCompositeType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadData
func (b_ BarcodeObservation) SupplementalPayloadData() foundation.NSData {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("supplementalPayloadData"))
	return rv
}

// The supplemental code decoded as a string value.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadString
func (b_ BarcodeObservation) SupplementalPayloadString() appkit.string {
	rv := objc.Send[appkit.string](b_.ID, objc.Sel("supplementalPayloadString"))
	return rv
}

// The symbology of the observed barcode.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/symbology
func (b_ BarcodeObservation) Symbology() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("symbology"))
	return rv
}

// The results of a barcode detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (b_ BarcodeObservation) Results() VNBarcodeObservation {
	rv := objc.Send[VNBarcodeObservation](b_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of a barcode detection request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (b_ BarcodeObservation) SetResults(value IVNBarcodeObservation) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResults:"), value)
}



