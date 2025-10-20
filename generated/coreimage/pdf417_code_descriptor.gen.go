// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDF417CodeDescriptor] class.
var (
	pDF417CodeDescriptorClass     _PDF417CodeDescriptorClass
	pDF417CodeDescriptorClassOnce sync.Once
)

func getPDF417CodeDescriptorClass() _PDF417CodeDescriptorClass {
	pDF417CodeDescriptorClassOnce.Do(func() {
		pDF417CodeDescriptorClass = _PDF417CodeDescriptorClass{objc.GetClass("CIPDF417CodeDescriptor")}
	})
	return pDF417CodeDescriptorClass
}

type _PDF417CodeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PDF417CodeDescriptor] class.
type IPDF417CodeDescriptor interface {
	IBarcodeDescriptor
}

// A concrete subclass of Core Image Barcode Descriptor that represents a PDF417 symbol.
//
// PDF417 is a stacked linear barcode symbol format used predominantly in transport, ID cards, and inventory management. Each pattern in the code comprises 4 bars and spaces, 17 units long. Refer to the ISO/IEC 15438:2006(E) for the PDF417 symbol specification.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor
type PDF417CodeDescriptor struct {
	BarcodeDescriptor
}

// PDF417CodeDescriptorFrom constructs a [PDF417CodeDescriptor] from an unsafe.Pointer.
//
// A concrete subclass of Core Image Barcode Descriptor that represents a PDF417 symbol.
func PDF417CodeDescriptorFrom(ptr unsafe.Pointer) PDF417CodeDescriptor {
	return PDF417CodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _PDF417CodeDescriptorClass) Alloc() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _PDF417CodeDescriptorClass) New() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ PDF417CodeDescriptor) Init() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ PDF417CodeDescriptor) Autorelease() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDF417CodeDescriptor creates a new PDF417CodeDescriptor instance.
func NewPDF417CodeDescriptor() PDF417CodeDescriptor {
	return getPDF417CodeDescriptorClass().New()
}


// Initializes an PDF417 code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/init(payload:isCompact:rowCount:columnCount:)
func NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload unsafe.Pointer, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	instance := getPDF417CodeDescriptorClass().Alloc()
	rv := objc.Send[PDF417CodeDescriptor](instance.ID, objc.Sel("initWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	rv.Autorelease()
	return rv
}


// Creates an PDF417 code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/descriptorWithPayload:isCompact:rowCount:columnCount:
func (fc _PDF417CodeDescriptorClass) DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload unsafe.Pointer, isCompact bool, rowCount int, columnCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("descriptorWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return rv
}

// The number of columns in the PDF417 code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/columnCount-swift.property
func (f_ PDF417CodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](f_.ID, objc.Sel("columnCount"))
	return rv
}
// The error-corrected payload containing the data encoded in the PDF417 code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/errorCorrectedPayload-swift.property
func (f_ PDF417CodeDescriptor) ErrorCorrectedPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}
// A boolean value telling if the PDF417 code is compact.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/isCompact-swift.property
func (f_ PDF417CodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isCompact"))
	return rv
}
// The number of rows in the PDF417 code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/rowCount-swift.property
func (f_ PDF417CodeDescriptor) RowCount() int {
	rv := objc.Send[int](f_.ID, objc.Sel("rowCount"))
	return rv
}

