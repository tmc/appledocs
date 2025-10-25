// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMutableData */


/* debug [class_header]: Header for NSMutableData */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableData */
// An interface definition for the [MutableData] class.
type IMutableData interface {
	IData
	
/* debug [class_interface_properties]: Properties for MutableData */
	// properties:
	Length() uint
	SetLength(value uint)
	MutableBytes() objectivec.IObject
	NSCompressionErrorMaximum() int
	SetNSCompressionErrorMaximum(value int)
	NSCompressionErrorMinimum() int
	SetNSCompressionErrorMinimum(value int)
	NSCompressionFailedError() int
	SetNSCompressionFailedError(value int)
	NSDecompressionFailedError() int
	SetNSDecompressionFailedError(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableData */
	// methods:
	AppendData(other IData)
	AppendBytesLength(bytes objectivec.IObject, length uint)
	CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) bool
	DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) bool
	IncreaseLengthBy(extraLength uint)
	ReplaceBytesInRangeWithBytes(range_ objc.IObject /* cross-framework: Range */, bytes objectivec.IObject)
	ReplaceBytesInRangeWithBytesLength(range_ objc.IObject /* cross-framework: Range */, replacementBytes objectivec.IObject, replacementLength uint)
	ResetBytesInRange(range_ objc.IObject /* cross-framework: Range */)
	SetData(data IData)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableData */
// Alloc allocates a new instance without initialization.
func (mc _MutableDataClass) Alloc() MutableData {
	rv := objc.Send[MutableData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableData */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableData */

// Returns an initialized mutable data object capable of holding the specified number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/init(capacity:)
func NewMutableDataWithCapacity(capacity uint) MutableData {
	instance := getMutableDataClass().Alloc()
	rv := objc.Send[MutableData](instance.ID, objc.Sel("initWithCapacity:"), capacity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableDataWithCapacity */


// Initializes and returns a mutable data object containing a given number of zeroed bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/init(length:)
func NewMutableDataWithLength(length uint) MutableData {
	instance := getMutableDataClass().Alloc()
	rv := objc.Send[MutableData](instance.ID, objc.Sel("initWithLength:"), length)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableDataWithLength */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableData */

// Creates and returns a mutable data object capable of holding the specified number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/dataWithCapacity:
func (mc _MutableDataClass) DataWithCapacity(aNumItems uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("dataWithCapacity:"), aNumItems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithCapacity) */


// Creates and returns an mutable data object containing a given number of zeroed bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/dataWithLength:
func (mc _MutableDataClass) DataWithLength(length uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("dataWithLength:"), length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithLength) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableData */

// Appends the content of another data object to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:)
func (m_ MutableData) AppendData(other IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendData:"), other)
}/* debug [instance_methods/method]: AppendData */


// Appends to the receiver a given number of bytes from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/append(_:length:)
func (m_ MutableData) AppendBytesLength(bytes objectivec.IObject, length uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendBytes:length:"), bytes, length)
}/* debug [instance_methods/method]: AppendBytesLength */


// Compresses the data object’s bytes using an algorithm that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/compress(using:)
func (m_ MutableData) CompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("compressUsingAlgorithm:error:"), algorithm, error_)
	return rv
}/* debug [instance_methods/method]: CompressUsingAlgorithmError */


// Decompresses the data object’s bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/decompress(using:)
func (m_ MutableData) DecompressUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("decompressUsingAlgorithm:error:"), algorithm, error_)
	return rv
}/* debug [instance_methods/method]: DecompressUsingAlgorithmError */


// Increases the length of the receiver by a given number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/increaseLength(by:)
func (m_ MutableData) IncreaseLengthBy(extraLength uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("increaseLengthBy:"), extraLength)
}/* debug [instance_methods/method]: IncreaseLengthBy */


// Replaces with a given set of bytes a given range within the contents of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:)
func (m_ MutableData) ReplaceBytesInRangeWithBytes(range_ objc.IObject /* cross-framework: Range */, bytes objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceBytesInRange:withBytes:"), range_, bytes)
}/* debug [instance_methods/method]: ReplaceBytesInRangeWithBytes */


// Replaces with a given set of bytes a given range within the contents of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/replaceBytes(in:withBytes:length:)
func (m_ MutableData) ReplaceBytesInRangeWithBytesLength(range_ objc.IObject /* cross-framework: Range */, replacementBytes objectivec.IObject, replacementLength uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceBytesInRange:withBytes:length:"), range_, replacementBytes, replacementLength)
}/* debug [instance_methods/method]: ReplaceBytesInRangeWithBytesLength */


// Replaces with zeroes the contents of the receiver in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/resetBytes(in:)
func (m_ MutableData) ResetBytesInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetBytesInRange:"), range_)
}/* debug [instance_methods/method]: ResetBytesInRange */


// Replaces the entire contents of the receiver with the contents of another data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/setData(_:)
func (m_ MutableData) SetData(data IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), data)
}/* debug [instance_methods/method]: SetData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableData */

// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/length
func (m_ MutableData) Length() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// The number of bytes contained in the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/length
func (m_ MutableData) SetLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLength:"), value)
}/* debug [instance_properties/setter]: length */


// A pointer to the data contained by the mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData/mutableBytes
func (m_ MutableData) MutableBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("mutableBytes"))
	return rv
}/* debug [instance_properties/getter]: mutableBytes */


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) NSCompressionErrorMaximum() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionErrorMaximum */


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (m_ MutableData) SetNSCompressionErrorMaximum(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSCompressionErrorMaximum */


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) NSCompressionErrorMinimum() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionErrorMinimum */


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (m_ MutableData) SetNSCompressionErrorMinimum(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSCompressionErrorMinimum */


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) NSCompressionFailedError() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionFailedError */


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (m_ MutableData) SetNSCompressionFailedError(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSCompressionFailedError:"), value)
}/* debug [instance_properties/setter]: NSCompressionFailedError */


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) NSDecompressionFailedError() int {
	rv := objc.Send[int](m_.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}/* debug [instance_properties/getter]: NSDecompressionFailedError */


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (m_ MutableData) SetNSDecompressionFailedError(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}/* debug [instance_properties/setter]: NSDecompressionFailedError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableData */


