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
	Length() uint /* primitive/slice/pointer */
	SetLength(value uint /* primitive/slice/pointer */)
	MutableBytes() unsafe.Pointer
	NSCompressionErrorMaximum() int /* primitive/slice/pointer */
	SetNSCompressionErrorMaximum(value int /* primitive/slice/pointer */)
	NSCompressionErrorMinimum() int /* primitive/slice/pointer */
	SetNSCompressionErrorMinimum(value int /* primitive/slice/pointer */)
	NSCompressionFailedError() int /* primitive/slice/pointer */
	SetNSCompressionFailedError(value int /* primitive/slice/pointer */)
	NSDecompressionFailedError() int /* primitive/slice/pointer */
	SetNSDecompressionFailedError(value int /* primitive/slice/pointer */)
	// methods:
	AppendData(other IData)
	AppendBytesLength(bytes unsafe.Pointer, length uint /* primitive/slice/pointer */)
	CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	IncreaseLengthBy(extraLength uint /* primitive/slice/pointer */)
	ReplaceBytesInRangeWithBytes(range_ Range /* foo */, bytes unsafe.Pointer)
	ReplaceBytesInRangeWithBytesLength(range_ Range /* foo */, replacementBytes unsafe.Pointer, replacementLength uint /* primitive/slice/pointer */)
	ResetBytesInRange(range_ Range /* foo */)
	SetData(data IData)
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



// Returns an initialized mutable data object capable of holding the specified number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/init(capacity:)
func NewMutableDataWithCapacity(capacity uint /* primitive/slice/pointer */) MutableData {
	instance := getMutableDataClass().Alloc()
	rv := objc.Send[MutableData](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	rv.Autorelease()
	return rv
}


// Initializes and returns a mutable data object containing a given number of zeroed bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/init(length:)
func NewMutableDataWithLength(length uint /* primitive/slice/pointer */) MutableData {
	instance := getMutableDataClass().Alloc()
	rv := objc.Send[MutableData](instance.ID, objc.Sel("initWithLength:"), length)
	rv.Autorelease()
	return rv
}



// Appends the content of another data object to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:)
func (m_ MutableData) AppendData(other IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendData:"), other)
}


// Appends to the receiver a given number of bytes from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:length:)
func (m_ MutableData) AppendBytesLength(bytes unsafe.Pointer, length uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendBytes:length:"), bytes, length)
}


// Compresses the data object’s bytes using an algorithm that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/compress(using:)
func (m_ MutableData) CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("compressUsingAlgorithm:error:"), algorithm, error_)
	return rv
}


// Decompresses the data object’s bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/decompress(using:)
func (m_ MutableData) DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("decompressUsingAlgorithm:error:"), algorithm, error_)
	return rv
}


// Increases the length of the receiver by a given number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/increaseLength(by:)
func (m_ MutableData) IncreaseLengthBy(extraLength uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("increaseLengthBy:"), extraLength)
}


// Replaces with a given set of bytes a given range within the contents of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:)
func (m_ MutableData) ReplaceBytesInRangeWithBytes(range_ Range /* foo */, bytes unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceBytesInRange:withBytes:"), range_, bytes)
}


// Replaces with a given set of bytes a given range within the contents of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:length:)
func (m_ MutableData) ReplaceBytesInRangeWithBytesLength(range_ Range /* foo */, replacementBytes unsafe.Pointer, replacementLength uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceBytesInRange:withBytes:length:"), range_, replacementBytes, replacementLength)
}


// Replaces with zeroes the contents of the receiver in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/resetBytes(in:)
func (m_ MutableData) ResetBytesInRange(range_ Range /* foo */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetBytesInRange:"), range_)
}


// Replaces the entire contents of the receiver with the contents of another data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/setData(_:)
func (m_ MutableData) SetData(data IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), data)
}


// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/length
func (m_ MutableData) Length() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](m_.ID, objc.Sel("length"))
	return rv
}


// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/length
func (m_ MutableData) SetLength(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLength:"), value)
}


// A pointer to the data contained by the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/mutableBytes
func (m_ MutableData) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mutableBytes"))
	return rv
}


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) NSCompressionErrorMaximum() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) SetNSCompressionErrorMaximum(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) NSCompressionErrorMinimum() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) SetNSCompressionErrorMinimum(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) NSCompressionFailedError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) SetNSCompressionFailedError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionFailedError:"), value)
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) NSDecompressionFailedError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](m_.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) SetNSDecompressionFailedError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}


