// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIDataMatrixCodeDescriptor */


/* debug [class_header]: Header for CIDataMatrixCodeDescriptor */
// The class instance for the [DataMatrixCodeDescriptor] class.
var (
	DataMatrixCodeDescriptorClass     _DataMatrixCodeDescriptorClass
	DataMatrixCodeDescriptorClassOnce sync.Once
)

func getDataMatrixCodeDescriptorClass() _DataMatrixCodeDescriptorClass {
	DataMatrixCodeDescriptorClassOnce.Do(func() {
		DataMatrixCodeDescriptorClass = _DataMatrixCodeDescriptorClass{objc.GetClass("CIDataMatrixCodeDescriptor")}
	})
	return DataMatrixCodeDescriptorClass
}

type _DataMatrixCodeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DataMatrixCodeDescriptor */
// An interface definition for the [DataMatrixCodeDescriptor] class.
type IDataMatrixCodeDescriptor interface {
	IBarcodeDescriptor
	
/* debug [class_interface_properties]: Properties for DataMatrixCodeDescriptor */
	// properties:
	ColumnCount() int
	EccVersion() DataMatrixCodeECCVersion
	ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */
	RowCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DataMatrixCodeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DataMatrixCodeDescriptor */
// Alloc allocates a new instance without initialization.
func (dc _DataMatrixCodeDescriptorClass) Alloc() DataMatrixCodeDescriptor {
	rv := objc.Send[DataMatrixCodeDescriptor](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DataMatrixCodeDescriptorClass) New() DataMatrixCodeDescriptor {
	rv := objc.Send[DataMatrixCodeDescriptor](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DataMatrixCodeDescriptor) Init() DataMatrixCodeDescriptor {
	rv := objc.Send[DataMatrixCodeDescriptor](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DataMatrixCodeDescriptor) Autorelease() DataMatrixCodeDescriptor {
	rv := objc.Send[DataMatrixCodeDescriptor](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataMatrixCodeDescriptor creates a new DataMatrixCodeDescriptor instance.
func NewDataMatrixCodeDescriptor() DataMatrixCodeDescriptor {
	return getDataMatrixCodeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DataMatrixCodeDescriptor */
// A concrete subclass the Core Image Barcode Descriptor that represents an Data Matrix code symbol.
//
// A Data Matrix code symbol is a 2D barcode format defined by the ISO/IEC 16022:2006(E) standard. It encodes data in square or rectangular symbol with solid lines on the left and bottom sides


// A concrete subclass the Core Image Barcode Descriptor that represents an Data Matrix code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor
type DataMatrixCodeDescriptor struct {
	BarcodeDescriptor
}

// DataMatrixCodeDescriptorFrom constructs a [DataMatrixCodeDescriptor] from an unsafe.Pointer.
//
// A concrete subclass the Core Image Barcode Descriptor that represents an Data Matrix code symbol.
func DataMatrixCodeDescriptorFrom(ptr unsafe.Pointer) DataMatrixCodeDescriptor {
	return DataMatrixCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DataMatrixCodeDescriptor */

// Initializes a Data Matrix code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/init(payload:rowCount:columnCount:eccVersion:)
func NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) DataMatrixCodeDescriptor {
	instance := getDataMatrixCodeDescriptorClass().Alloc()
	rv := objc.Send[DataMatrixCodeDescriptor](instance.ID, objc.Sel("initWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DataMatrixCodeDescriptor */

// Creates a Data Matrix code descriptor for the given payload and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/descriptorWithPayload:rowCount:columnCount:eccVersion:
func (dc _DataMatrixCodeDescriptorClass) DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload objc.IObject /* cross-framework: NSData */, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("descriptorWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPayloadRowCountColumnCountEccVersion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DataMatrixCodeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DataMatrixCodeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DataMatrixCodeDescriptor */

// The number of columns in the Data Matrix code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-swift.property
func (d_ DataMatrixCodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("columnCount"))
	return rv
}/* debug [instance_properties/getter]: columnCount */


// The error correction version of the Data Matrix code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-swift.property
func (d_ DataMatrixCodeDescriptor) EccVersion() DataMatrixCodeECCVersion {
	rv := objc.Send[DataMatrixCodeECCVersion](d_.ID, objc.Sel("eccVersion"))
	return rv
}/* debug [instance_properties/getter]: eccVersion */


// The error-corrected payload containing the data encoded in the Data Matrix code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/errorCorrectedPayload-swift.property
func (d_ DataMatrixCodeDescriptor) ErrorCorrectedPayload() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}/* debug [instance_properties/getter]: errorCorrectedPayload */


// The number of rows in the Data Matrix code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-swift.property
func (d_ DataMatrixCodeDescriptor) RowCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("rowCount"))
	return rv
}/* debug [instance_properties/getter]: rowCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIDataMatrixCodeDescriptor */


