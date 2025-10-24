// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Data] class.
var (
	DataClass     _DataClass
	DataClassOnce sync.Once
)

func getDataClass() _DataClass {
	DataClassOnce.Do(func() {
		DataClass = _DataClass{objc.GetClass("NSData")}
	})
	return DataClass
}

type _DataClass struct {
	class objc.Class
}

// An interface definition for the [Data] class.
type IData interface {
	objectivec.IObject
	// properties:
	NSCompressionErrorMaximum() int
	SetNSCompressionErrorMaximum(value int)
	NSCompressionErrorMinimum() int
	SetNSCompressionErrorMinimum(value int)
	NSCompressionFailedError() int
	SetNSCompressionFailedError(value int)
	Bytes() unsafe.Pointer
	SetBytes(value unsafe.Pointer)
	Description() IString
	SetDescription(value IString)
	Length() int
	SetLength(value int)
	NSDecompressionFailedError() int
	SetNSDecompressionFailedError(value int)
	// methods:
}

// A static byte buffer in memory.
//
// In Swift, the buffer bridges to ; use when you need reference semantics or other Foundation-specific behavior. and its mutable subclass provide data objects, or object-oriented wrappers for byte buffers. Data objects let simple allocated buffers (that is, data with no embedded pointers) take on the behavior of Foundation objects. The size of the data is subject to a theoretical limit of about 8 exabytes (1 EB = 10¹⁸ bytes; in practice, the limit should not be a factor). is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A static byte buffer in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData
type Data struct {
	objectivec.Object
}

// DataFrom constructs a [Data] from an unsafe.Pointer.
//
// A static byte buffer in memory.
func DataFrom(ptr unsafe.Pointer) Data {
	return Data{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DataClass) Alloc() Data {
	rv := objc.Send[Data](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DataClass) New() Data {
	rv := objc.Send[Data](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Data) Init() Data {
	rv := objc.Send[Data](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Data) Autorelease() Data {
	rv := objc.Send[Data](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewData creates a new Data instance.
func NewData() Data {
	return getDataClass().New()
}



// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (d_ Data) NSCompressionErrorMaximum() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (d_ Data) SetNSCompressionErrorMaximum(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (d_ Data) NSCompressionErrorMinimum() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (d_ Data) SetNSCompressionErrorMinimum(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (d_ Data) NSCompressionFailedError() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (d_ Data) SetNSCompressionFailedError(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionFailedError:"), value)
}


// A pointer to the data object’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/bytes
func (d_ Data) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("bytes"))
	return rv
}


// A pointer to the data object’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/bytes
func (d_ Data) SetBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBytes:"), value)
}


// A string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/description
func (d_ Data) Description() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("description"))
	return rv
}


// A string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/description
func (d_ Data) SetDescription(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescription:"), value)
}


// The number of bytes contained by the data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/length
func (d_ Data) Length() int {
	rv := objc.Send[int](d_.ID, objc.Sel("length"))
	return rv
}


// The number of bytes contained by the data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdata/length
func (d_ Data) SetLength(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLength:"), value)
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (d_ Data) NSDecompressionFailedError() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (d_ Data) SetNSDecompressionFailedError(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}



