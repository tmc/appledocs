// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIQRCodeDescriptor */


/* debug [class_header]: Header for CIQRCodeDescriptor */
// The class instance for the [QRCodeDescriptor] class.
var (
	QRCodeDescriptorClass     _QRCodeDescriptorClass
	QRCodeDescriptorClassOnce sync.Once
)

func getQRCodeDescriptorClass() _QRCodeDescriptorClass {
	QRCodeDescriptorClassOnce.Do(func() {
		QRCodeDescriptorClass = _QRCodeDescriptorClass{objc.GetClass("CIQRCodeDescriptor")}
	})
	return QRCodeDescriptorClass
}

type _QRCodeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QRCodeDescriptor */
// An interface definition for the [QRCodeDescriptor] class.
type IQRCodeDescriptor interface {
	IBarcodeDescriptor
	
/* debug [class_interface_properties]: Properties for QRCodeDescriptor */
	// properties:
	ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */
	ErrorCorrectionLevel() QRCodeErrorCorrectionLevel
	MaskPattern() uint8 /* not a class type */
	SymbolVersion() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QRCodeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QRCodeDescriptor */
// Alloc allocates a new instance without initialization.
func (qc _QRCodeDescriptorClass) Alloc() QRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QRCodeDescriptorClass) New() QRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QRCodeDescriptor) Init() QRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QRCodeDescriptor) Autorelease() QRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQRCodeDescriptor creates a new QRCodeDescriptor instance.
func NewQRCodeDescriptor() QRCodeDescriptor {
	return getQRCodeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QRCodeDescriptor */
// A concrete subclass of the Core Image Barcode Descriptor that represents a square QR code symbol.
//
// ISO/IEC 18004 defines versions from 1 to 40, where a higher symbol version indicates a larger data-carrying capacity. QR Codes can encode text, vCard contact information, or Uniform Resource Identifiers (URI).


// A concrete subclass of the Core Image Barcode Descriptor that represents a square QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor
type QRCodeDescriptor struct {
	BarcodeDescriptor
}

// QRCodeDescriptorFrom constructs a [QRCodeDescriptor] from an unsafe.Pointer.
//
// A concrete subclass of the Core Image Barcode Descriptor that represents a square QR code symbol.
func QRCodeDescriptorFrom(ptr unsafe.Pointer) QRCodeDescriptor {
	return QRCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QRCodeDescriptor */

// Initializes a QR code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/init(payload:symbolVersion:maskPattern:errorCorrectionLevel:)
func NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, symbolVersion int, maskPattern uint8 /* not a class type */, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	instance := getQRCodeDescriptorClass().Alloc()
	rv := objc.Send[QRCodeDescriptor](instance.ID, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QRCodeDescriptor */

// Creates a QR code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:
func (qc _QRCodeDescriptorClass) DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, symbolVersion int, maskPattern uint8 /* not a class type */, errorCorrectionLevel QRCodeErrorCorrectionLevel) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(qc.class), objc.Sel("descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QRCodeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QRCodeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QRCodeDescriptor */

// The error-corrected codeword payload that comprises the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectedPayload-swift.property
func (q_ QRCodeDescriptor) ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](q_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}/* debug [instance_properties/getter]: errorCorrectedPayload */


// The error correction level of the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-swift.property
func (q_ QRCodeDescriptor) ErrorCorrectionLevel() QRCodeErrorCorrectionLevel {
	rv := objc.Send[QRCodeErrorCorrectionLevel](q_.ID, objc.Sel("errorCorrectionLevel"))
	return rv
}/* debug [instance_properties/getter]: errorCorrectionLevel */


// The data mask pattern for the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-swift.property
func (q_ QRCodeDescriptor) MaskPattern() uint8 /* not a class type */ {
	rv := objc.Send[uint8](q_.ID, objc.Sel("maskPattern"))
	return rv
}/* debug [instance_properties/getter]: maskPattern */


// The version of the QR code which corresponds to the size of the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-swift.property
func (q_ QRCodeDescriptor) SymbolVersion() int {
	rv := objc.Send[int](q_.ID, objc.Sel("symbolVersion"))
	return rv
}/* debug [instance_properties/getter]: symbolVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIQRCodeDescriptor */


