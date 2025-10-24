// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIAztecCodeDescriptor */


/* debug [class_header]: Header for CIAztecCodeDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AztecCodeDescriptor */
// An interface definition for the [AztecCodeDescriptor] class.
type IAztecCodeDescriptor interface {
	IBarcodeDescriptor
	
/* debug [class_interface_properties]: Properties for AztecCodeDescriptor */
	// properties:
	DataCodewordCount() int
	ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */
	IsCompact() bool
	LayerCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AztecCodeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AztecCodeDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AztecCodeDescriptorClass) Alloc() AztecCodeDescriptor {
	rv := objc.Send[AztecCodeDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AztecCodeDescriptor */
// A concrete subclass the Core Image Barcode Descriptor that represents an Aztec code symbol.
//
// An Aztec code symbol is a 2D barcode format defined by the ISO/IEC 24778:2008 standard. It encodes data in concentric square rings around a central bullseye pattern.


// A concrete subclass the Core Image Barcode Descriptor that represents an Aztec code symbol.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AztecCodeDescriptor */

// Initializes an Aztec code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/init(payload:isCompact:layerCount:dataCodewordCount:)
func NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	instance := getAztecCodeDescriptorClass().Alloc()
	rv := objc.Send[AztecCodeDescriptor](instance.ID, objc.Sel("initWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AztecCodeDescriptor */

// Creates an Aztec code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/descriptorWithPayload:isCompact:layerCount:dataCodewordCount:
func (ac _AztecCodeDescriptorClass) DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, isCompact bool, layerCount int, dataCodewordCount int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptorWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPayloadIsCompactLayerCountDataCodewordCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AztecCodeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AztecCodeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AztecCodeDescriptor */

// The number of non-error-correction codewords carried by the Aztec code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-swift.property
func (a_ AztecCodeDescriptor) DataCodewordCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("dataCodewordCount"))
	return rv
}/* debug [instance_properties/getter]: dataCodewordCount */


// The error-corrected payload that comprises the the Aztec code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-swift.property
func (a_ AztecCodeDescriptor) ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}/* debug [instance_properties/getter]: errorCorrectedPayload */


// A Boolean value telling if the Aztec code is compact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-swift.property
func (a_ AztecCodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompact"))
	return rv
}/* debug [instance_properties/getter]: isCompact */


// The number of data layers in the Aztec code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-swift.property
func (a_ AztecCodeDescriptor) LayerCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("layerCount"))
	return rv
}/* debug [instance_properties/getter]: layerCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIAztecCodeDescriptor */


