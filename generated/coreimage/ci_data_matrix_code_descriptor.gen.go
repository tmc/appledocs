// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [DataMatrixCodeDescriptor] class.
type IDataMatrixCodeDescriptor interface {
	IBarcodeDescriptor
}

// A concrete subclass the Core Image Barcode Descriptor that represents an Data Matrix code symbol.
//
// A Data Matrix code symbol is a 2D barcode format defined by the ISO/IEC 16022:2006(E) standard. It encodes data in square or rectangular symbol with solid lines on the left and bottom sides
//
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

// Alloc allocates a new instance without initialization.
func (dc _DataMatrixCodeDescriptorClass) Alloc() DataMatrixCodeDescriptor {
	rv := objc.Send[DataMatrixCodeDescriptor](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a Data Matrix code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/init(payload:rowCount:columnCount:eccVersion:)
func NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload unsafe.Pointer, rowCount int, columnCount int, eccVersion unsafe.Pointer) DataMatrixCodeDescriptor {
	instance := getDataMatrixCodeDescriptorClass().Alloc()
	rv := objc.Send[DataMatrixCodeDescriptor](instance.ID, objc.Sel("initWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	rv.Autorelease()
	return rv
}


// Creates a Data Matrix code descriptor for the given payload and parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/descriptorWithPayload:rowCount:columnCount:eccVersion:
func (dc _DataMatrixCodeDescriptorClass) DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload unsafe.Pointer, rowCount int, columnCount int, eccVersion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("descriptorWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return rv
}

// The number of columns in the Data Matrix code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-swift.property
func (d_ DataMatrixCodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("columnCount"))
	return rv
}

// The error correction version of the Data Matrix code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-swift.property
func (d_ DataMatrixCodeDescriptor) EccVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("eccVersion"))
	return rv
}

// The error-corrected payload containing the data encoded in the Data Matrix code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/errorCorrectedPayload-swift.property
func (d_ DataMatrixCodeDescriptor) ErrorCorrectedPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("errorCorrectedPayload"))
	return rv
}

// The number of rows in the Data Matrix code symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-swift.property
func (d_ DataMatrixCodeDescriptor) RowCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("rowCount"))
	return rv
}


