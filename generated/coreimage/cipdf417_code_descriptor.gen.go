// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIPDF417CodeDescriptor */


/* debug [class_header]: Header for CIPDF417CodeDescriptor */
// The class instance for the [PDF417CodeDescriptor] class.
var (
	PDF417CodeDescriptorClass     _PDF417CodeDescriptorClass
	PDF417CodeDescriptorClassOnce sync.Once
)

func getPDF417CodeDescriptorClass() _PDF417CodeDescriptorClass {
	PDF417CodeDescriptorClassOnce.Do(func() {
		PDF417CodeDescriptorClass = _PDF417CodeDescriptorClass{objc.GetClass("CIPDF417CodeDescriptor")}
	})
	return PDF417CodeDescriptorClass
}

type _PDF417CodeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDF417CodeDescriptor */
// An interface definition for the [PDF417CodeDescriptor] class.
type IPDF417CodeDescriptor interface {
	IBarcodeDescriptor
	
/* debug [class_interface_properties]: Properties for PDF417CodeDescriptor */
	// properties:
	ColumnCount() int
	ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */
	IsCompact() bool
	RowCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDF417CodeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDF417CodeDescriptor */
// Alloc allocates a new instance without initialization.
func (pc _PDF417CodeDescriptorClass) Alloc() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDF417CodeDescriptorClass) New() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDF417CodeDescriptor) Init() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDF417CodeDescriptor) Autorelease() PDF417CodeDescriptor {
	rv := objc.Send[PDF417CodeDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDF417CodeDescriptor creates a new PDF417CodeDescriptor instance.
func NewPDF417CodeDescriptor() PDF417CodeDescriptor {
	return getPDF417CodeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDF417CodeDescriptor */
// A concrete subclass of Core Image Barcode Descriptor that represents a PDF417 symbol.
//
// PDF417 is a stacked linear barcode symbol format used predominantly in transport, ID cards, and inventory management. Each pattern in the code comprises 4 bars and spaces, 17 units long. Refer to the ISO/IEC 15438:2006(E) for the PDF417 symbol specification.


// A concrete subclass of Core Image Barcode Descriptor that represents a PDF417 symbol.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDF417CodeDescriptor */

// Initializes an PDF417 code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/init(payload:isCompact:rowCount:columnCount:)
func NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	instance := getPDF417CodeDescriptorClass().Alloc()
	rv := objc.Send[PDF417CodeDescriptor](instance.ID, objc.Sel("initWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDF417CodeDescriptor */

// Creates an PDF417 code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/descriptorWithPayload:isCompact:rowCount:columnCount:
func (pc _PDF417CodeDescriptorClass) DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, isCompact bool, rowCount int, columnCount int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("descriptorWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPayloadIsCompactRowCountColumnCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDF417CodeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDF417CodeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDF417CodeDescriptor */

// The number of columns in the PDF417 code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/columnCount-swift.property
func (p_ PDF417CodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("columnCount"))
	return rv
}/* debug [instance_properties/getter]: columnCount */


// The error-corrected payload containing the data encoded in the PDF417 code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/errorCorrectedPayload-swift.property
func (p_ PDF417CodeDescriptor) ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}/* debug [instance_properties/getter]: errorCorrectedPayload */


// A boolean value telling if the PDF417 code is compact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/isCompact-swift.property
func (p_ PDF417CodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCompact"))
	return rv
}/* debug [instance_properties/getter]: isCompact */


// The number of rows in the PDF417 code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/rowCount-swift.property
func (p_ PDF417CodeDescriptor) RowCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rowCount"))
	return rv
}/* debug [instance_properties/getter]: rowCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIPDF417CodeDescriptor */


