// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableData] class.
var (
	MutableDataClass     _MutableDataClass
	MutableDataClassOnce sync.Once
)

func getMutableDataClass() _MutableDataClass {
	MutableDataClassOnce.Do(func() {
		MutableDataClass = _MutableDataClass{objc.GetClass("NSMutableData")}
	})
	return MutableDataClass
}

type _MutableDataClass struct {
	class objc.Class
}

// An interface definition for the [MutableData] class.
type IMutableData interface {
	IData
	// properties:
	NSCompressionErrorMaximum() int
	SetNSCompressionErrorMaximum(value int)
	NSCompressionErrorMinimum() int
	SetNSCompressionErrorMinimum(value int)
	NSCompressionFailedError() int
	SetNSCompressionFailedError(value int)
	NSDecompressionFailedError() int
	SetNSDecompressionFailedError(value int)
	Length() int
	SetLength(value int)
	MutableBytes() unsafe.Pointer
	SetMutableBytes(value unsafe.Pointer)
	// methods:
}

// An object representing a dynamic byte buffer in memory.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. and its superclass provide data objects, or object-oriented wrappers for byte buffers. Data objects let simple allocated buffers (that is, data with no embedded pointers) take on the behavior of Foundation objects. They are typically used for data storage and are also useful in Distributed Objects applications, where data contained in data objects can be copied or moved between applications. creates static data objects, and creates dynamic data objects. You can easily convert one type of data object to the other with the initializer that takes an object or an object as an argument. The following methods change when used on a mutable data object: When called, the bytes are immediately copied and then the buffer is freed. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// An object representing a dynamic byte buffer in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData
type MutableData struct {
	Data
}

// MutableDataFrom constructs a [MutableData] from an unsafe.Pointer.
//
// An object representing a dynamic byte buffer in memory.
func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: DataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableDataClass) Alloc() MutableData {
	rv := objc.Send[MutableData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableDataClass) New() MutableData {
	rv := objc.Send[MutableData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableData) Init() MutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableData) Autorelease() MutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableData creates a new MutableData instance.
func NewMutableData() MutableData {
	return getMutableDataClass().New()
}



// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) NSCompressionErrorMaximum() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) SetNSCompressionErrorMaximum(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) NSCompressionErrorMinimum() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) SetNSCompressionErrorMinimum(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) NSCompressionFailedError() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) SetNSCompressionFailedError(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionFailedError:"), value)
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) NSDecompressionFailedError() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) SetNSDecompressionFailedError(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}


// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/length
func (m_ MutableData) Length() int {
	rv := objc.Send[int](m_.ID, objc.Sel("length"))
	return rv
}


// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/length
func (m_ MutableData) SetLength(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLength:"), value)
}


// A pointer to the data contained by the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/mutablebytes
func (m_ MutableData) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mutableBytes"))
	return rv
}


// A pointer to the data contained by the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledata/mutablebytes
func (m_ MutableData) SetMutableBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMutableBytes:"), value)
}



