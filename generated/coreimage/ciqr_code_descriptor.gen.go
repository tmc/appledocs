// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [QRCodeDescriptor] class.
type IQRCodeDescriptor interface {
	IBarcodeDescriptor
	ErrorCorrectedPayload() foundation.NSData
	ErrorCorrectionLevel() QRCodeErrorCorrectionLevel
	MaskPattern() unsafe.Pointer
	SymbolVersion() int
}

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

// Alloc allocates a new instance without initialization.
func (qc _QRCodeDescriptorClass) Alloc() QRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a QR code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/init(payload:symbolVersion:maskPattern:errorCorrectionLevel:)
func NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.IData, symbolVersion int, maskPattern unsafe.Pointer, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	instance := getQRCodeDescriptorClass().Alloc()
	rv := objc.Send[QRCodeDescriptor](instance.ID, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	rv.Autorelease()
	return rv
}



// Creates a QR code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:
func (qc _QRCodeDescriptorClass) DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.IData, symbolVersion int, maskPattern unsafe.Pointer, errorCorrectionLevel QRCodeErrorCorrectionLevel) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(qc.class), objc.Sel("descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}


// The error-corrected codeword payload that comprises the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectedPayload-swift.property
func (q_ QRCodeDescriptor) ErrorCorrectedPayload() foundation.NSData {
	rv := objc.Send[foundation.NSData](q_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}


// The error correction level of the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-swift.property
func (q_ QRCodeDescriptor) ErrorCorrectionLevel() QRCodeErrorCorrectionLevel {
	rv := objc.Send[QRCodeErrorCorrectionLevel](q_.ID, objc.Sel("errorCorrectionLevel"))
	return rv
}


// The data mask pattern for the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-swift.property
func (q_ QRCodeDescriptor) MaskPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("maskPattern"))
	return rv
}


// The version of the QR code which corresponds to the size of the QR code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-swift.property
func (q_ QRCodeDescriptor) SymbolVersion() int {
	rv := objc.Send[int](q_.ID, objc.Sel("symbolVersion"))
	return rv
}


