// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AztecCodeDescriptor] class.
var (
	AztecCodeDescriptorClass     _AztecCodeDescriptorClass
	AztecCodeDescriptorClassOnce sync.Once
)

func getAztecCodeDescriptorClass() _AztecCodeDescriptorClass {
	AztecCodeDescriptorClassOnce.Do(func() {
		AztecCodeDescriptorClass = _AztecCodeDescriptorClass{objc.GetClass("CIAztecCodeDescriptor")}
	})
	return AztecCodeDescriptorClass
}

type _AztecCodeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AztecCodeDescriptor] class.
type IAztecCodeDescriptor interface {
	IBarcodeDescriptor
}

// A concrete subclass the Core Image Barcode Descriptor that represents an Aztec code symbol.
//
// An Aztec code symbol is a 2D barcode format defined by the ISO/IEC 24778:2008 standard. It encodes data in concentric square rings around a central bullseye pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor
type AztecCodeDescriptor struct {
	BarcodeDescriptor
}

// AztecCodeDescriptorFrom constructs a [AztecCodeDescriptor] from an unsafe.Pointer.
//
// A concrete subclass the Core Image Barcode Descriptor that represents an Aztec code symbol.
func AztecCodeDescriptorFrom(ptr unsafe.Pointer) AztecCodeDescriptor {
	return AztecCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AztecCodeDescriptorClass) Alloc() AztecCodeDescriptor {
	rv := objc.Send[AztecCodeDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AztecCodeDescriptorClass) New() AztecCodeDescriptor {
	rv := objc.Send[AztecCodeDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AztecCodeDescriptor) Init() AztecCodeDescriptor {
	rv := objc.Send[AztecCodeDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AztecCodeDescriptor) Autorelease() AztecCodeDescriptor {
	rv := objc.Send[AztecCodeDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAztecCodeDescriptor creates a new AztecCodeDescriptor instance.
func NewAztecCodeDescriptor() AztecCodeDescriptor {
	return getAztecCodeDescriptorClass().New()
}


// Initializes an Aztec code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/init(payload:isCompact:layerCount:dataCodewordCount:)
func NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload unsafe.Pointer, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	instance := getAztecCodeDescriptorClass().Alloc()
	rv := objc.Send[AztecCodeDescriptor](instance.ID, objc.Sel("initWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	rv.Autorelease()
	return rv
}


// Creates an Aztec code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/descriptorWithPayload:isCompact:layerCount:dataCodewordCount:
func (ac _AztecCodeDescriptorClass) DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload unsafe.Pointer, isCompact bool, layerCount int, dataCodewordCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("descriptorWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return rv
}

// The number of non-error-correction codewords carried by the Aztec code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-swift.property
func (a_ AztecCodeDescriptor) DataCodewordCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("dataCodewordCount"))
	return rv
}

// The error-corrected payload that comprises the the Aztec code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-swift.property
func (a_ AztecCodeDescriptor) ErrorCorrectedPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}

// A Boolean value telling if the Aztec code is compact.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-swift.property
func (a_ AztecCodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompact"))
	return rv
}

// The number of data layers in the Aztec code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-swift.property
func (a_ AztecCodeDescriptor) LayerCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("layerCount"))
	return rv
}


