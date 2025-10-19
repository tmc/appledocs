// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QRCodeDescriptor] class.
var (
	qRCodeDescriptorClass     _QRCodeDescriptorClass
	qRCodeDescriptorClassOnce sync.Once
)

func getQRCodeDescriptorClass() _QRCodeDescriptorClass {
	qRCodeDescriptorClassOnce.Do(func() {
		qRCodeDescriptorClass = _QRCodeDescriptorClass{objc.GetClass("CIQRCodeDescriptor")}
	})
	return qRCodeDescriptorClass
}

type _QRCodeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [QRCodeDescriptor] class.
type IQRCodeDescriptor interface {
	IBarcodeDescriptor
}

// A concrete subclass of the Core Image Barcode Descriptor that represents a square QR code symbol.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/init(payload:symbolVersion:maskPattern:errorCorrectionLevel:)
func NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload unsafe.Pointer, symbolVersion int, maskPattern unsafe.Pointer, errorCorrectionLevel unsafe.Pointer) QRCodeDescriptor {
	instance := getQRCodeDescriptorClass().Alloc()
	rv := objc.Send[QRCodeDescriptor](instance.ID, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	rv.Autorelease()
	return rv
}


// Creates a QR code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:
func (qc _QRCodeDescriptorClass) DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload unsafe.Pointer, symbolVersion int, maskPattern unsafe.Pointer, errorCorrectionLevel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(qc.class), objc.Sel("descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}

