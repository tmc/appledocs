// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSData */


/* debug [class_header]: Header for NSData */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Data */
// An interface definition for the [Data] class.
type IData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Data */
	// properties:
	Bytes() objectivec.IObject
	Description() IString
	Length() uint
	NSCompressionErrorMaximum() int
	SetNSCompressionErrorMaximum(value int)
	NSCompressionErrorMinimum() int
	SetNSCompressionErrorMinimum(value int)
	NSCompressionFailedError() int
	SetNSCompressionFailedError(value int)
	NSDecompressionFailedError() int
	SetNSDecompressionFailedError(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Data */
	// methods:
	Base64EncodedDataWithOptions(options DataBase64EncodingOptions) IData
	Base64EncodedStringWithOptions(options DataBase64EncodingOptions) IString
	CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) objectivec.IObject
	DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) objectivec.IObject
	EnumerateByteRangesUsingBlock(block unsafe.Pointer)
	GetBytesLength(buffer objectivec.IObject, length uint)
	GetBytesRange(buffer objectivec.IObject, range_ objc.IObject /* cross-framework: Range */)
	IsEqualToData(other IData) bool
	RangeOfDataOptionsRange(dataToFind IData, mask DataSearchOptions, searchRange objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Range */
	SubdataWithRange(range_ objc.IObject /* cross-framework: Range */) IData
	WriteToURLAtomically(url IURL, atomically bool) bool
	WriteToURLOptionsError(url IURL, writeOptionsMask DataWritingOptions, errorPtr IError) bool
	WriteToFileAtomically(path IString, useAuxiliaryFile bool) bool
	WriteToFileOptionsError(path IString, writeOptionsMask DataWritingOptions, errorPtr IError) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Data */
// Alloc allocates a new instance without initialization.
func (dc _DataClass) Alloc() Data {
	rv := objc.Send[Data](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Data */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Data */

// Initializes a data object with the given Base64 encoded data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64EncodedData:options:)
func NewDataWithBase64EncodedDataOptions(base64Data IData, options DataBase64DecodingOptions) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBase64EncodedDataOptions */


// Initializes a data object with the given Base64 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64EncodedString:options:)
func NewDataWithBase64EncodedStringOptions(base64String IString, options DataBase64DecodingOptions) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64EncodedString:options:"), base64String, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBase64EncodedStringOptions */


// Initializes a data object initialized with the given Base64 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoding:)
func NewDataWithBase64Encoding(base64String IString) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64Encoding:"), base64String)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBase64Encoding */


// Initializes a data object filled with a given number of bytes copied from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytes:length:)
func NewDataWithBytesLength(bytes objectivec.IObject, length uint) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytes:length:"), bytes, length)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBytesLength */


// Initializes a data object filled with a given number of bytes of data from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:)
func NewDataWithBytesNoCopyLength(bytes objectivec.IObject, length uint) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBytesNoCopyLength */


// Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:deallocator:)
func NewDataWithBytesNoCopyLengthDeallocator(bytes objectivec.IObject, length uint, deallocator unsafe.Pointer) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBytesNoCopyLengthDeallocator */


// Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:freeWhenDone:)
func NewDataWithBytesNoCopyLengthFreeWhenDone(bytes objectivec.IObject, length uint, b bool) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithBytesNoCopyLengthFreeWhenDone */


// Initializes a data object with the content of the file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:)
func NewDataWithContentsOfFile(path IString) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithContentsOfFile */


// Initializes a data object with the content of the file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:options:)
func NewDataWithContentsOfFileOptionsError(path IString, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfFile:options:error:"), path, readOptionsMask, errorPtr)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithContentsOfFileOptionsError */


// Initializes a data object with the contents of the mapped file specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfMappedFile:)
func NewDataWithContentsOfMappedFile(path IString) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfMappedFile:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithContentsOfMappedFile */


// Creates a data object from the data at the specified file URL, or returns if the system can’t create one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:)-6rrnr
func NewDataWithContentsOfURL(url IURL) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithContentsOfURL */


// Creates a data object from the data at the provided file URL using specific reading options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:options:)-5abi3
func NewDataWithContentsOfURLOptionsError(url IURL, readOptionsMask DataReadingOptions, errorPtr IError) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, errorPtr)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithContentsOfURLOptionsError */


// Initializes a data object with the contents of another data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(data:)
func NewDataWithData(data IData) Data {
	instance := getDataClass().Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Data */

// Creates an empty data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/data
func (dc _DataClass) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("data"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Data) */


// Creates a data object containing a given number of bytes copied from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytes:length:
func (dc _DataClass) DataWithBytesLength(bytes objectivec.IObject, length uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithBytes:length:"), bytes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithBytesLength) */


// Creates a data object that holds a given number of bytes from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:
func (dc _DataClass) DataWithBytesNoCopyLength(bytes objectivec.IObject, length uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithBytesNoCopy:length:"), bytes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithBytesNoCopyLength) */


// Creates a data object that holds a given number of bytes from a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:freeWhenDone:
func (dc _DataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes objectivec.IObject, length uint, b bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithBytesNoCopyLengthFreeWhenDone) */


// Creates a data object by reading every byte from the file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:
func (dc _DataClass) DataWithContentsOfFile(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithContentsOfFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContentsOfFile) */


// Creates a data object by reading every byte from the file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:options:error:
func (dc _DataClass) DataWithContentsOfFileOptionsError(path IString, readOptionsMask DataReadingOptions, errorPtr IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithContentsOfFile:options:error:"), path, readOptionsMask, errorPtr)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContentsOfFileOptionsError) */


// Creates a data object from the mapped file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfMappedFile(_:)
func (dc _DataClass) DataWithContentsOfMappedFile(path IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("dataWithContentsOfMappedFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContentsOfMappedFile) */


// Creates a data object containing the contents of another data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithData:
func (dc _DataClass) DataWithData(data IData) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithData) */


// Creates a data object from the data at the specified file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:)-6foqd
func (dc _DataClass) DataWithContentsOfURL(url IURL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithContentsOfURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContentsOfURL) */


// Creates a data object from the data at the provided file URL using specific reading options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:options:)-95rht
func (dc _DataClass) DataWithContentsOfURLOptionsError(url IURL, readOptionsMask DataReadingOptions, errorPtr IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dataWithContentsOfURL:options:error:"), url, readOptionsMask, errorPtr)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContentsOfURLOptionsError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Data */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Data */

// Creates a Base64, UTF-8 encoded data object from the string using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedData(options:)
func (d_ Data) Base64EncodedDataWithOptions(options DataBase64EncodingOptions) IData {
	rv := objc.Send[Data](d_.ID, objc.Sel("base64EncodedDataWithOptions:"), options)
	return rv
}/* debug [instance_methods/method]: Base64EncodedDataWithOptions */


// Creates a Base64 encoded string from the string using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedString(options:)
func (d_ Data) Base64EncodedStringWithOptions(options DataBase64EncodingOptions) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("base64EncodedStringWithOptions:"), options)
	return rv
}/* debug [instance_methods/method]: Base64EncodedStringWithOptions */


// Returns a new data object by compressing the data object’s bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/compressed(using:)
func (d_ Data) CompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, error_)
	return rv
}/* debug [instance_methods/method]: CompressedDataUsingAlgorithmError */


// Returns a new data object by decompressing data object’s bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/decompressed(using:)
func (d_ Data) DecompressedDataUsingAlgorithmError(algorithm DataCompressionAlgorithm, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, error_)
	return rv
}/* debug [instance_methods/method]: DecompressedDataUsingAlgorithmError */


// Enumerates each range of bytes in the data object using a block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/enumerateBytes(_:)
func (d_ Data) EnumerateByteRangesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateByteRangesUsingBlock:"), block)
}/* debug [instance_methods/method]: EnumerateByteRangesUsingBlock */


// Copies a number of bytes from the start of the data object into a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:length:)
func (d_ Data) GetBytesLength(buffer objectivec.IObject, length uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getBytes:length:"), buffer, length)
}/* debug [instance_methods/method]: GetBytesLength */


// Copies a range of bytes from the data object into a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:range:)
func (d_ Data) GetBytesRange(buffer objectivec.IObject, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getBytes:range:"), buffer, range_)
}/* debug [instance_methods/method]: GetBytesRange */


// Returns a Boolean value indicating whether this data object is the same as another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/isEqual(to:)
func (d_ Data) IsEqualToData(other IData) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToData:"), other)
	return rv
}/* debug [instance_methods/method]: IsEqualToData */


// Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/range(of:options:in:)
func (d_ Data) RangeOfDataOptionsRange(dataToFind IData, mask DataSearchOptions, searchRange objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("rangeOfData:options:range:"), dataToFind, mask, searchRange)
	return rv
}/* debug [instance_methods/method]: RangeOfDataOptionsRange */


// Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/subdata(with:)
func (d_ Data) SubdataWithRange(range_ objc.IObject /* cross-framework: Range */) IData {
	rv := objc.Send[Data](d_.ID, objc.Sel("subdataWithRange:"), range_)
	return rv
}/* debug [instance_methods/method]: SubdataWithRange */


// Writes the data object’s bytes to the location specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(to:atomically:)
func (d_ Data) WriteToURLAtomically(url IURL, atomically bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}/* debug [instance_methods/method]: WriteToURLAtomically */


// Writes the data object’s bytes to the location specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(to:options:)
func (d_ Data) WriteToURLOptionsError(url IURL, writeOptionsMask DataWritingOptions, errorPtr IError) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:options:error:"), url, writeOptionsMask, errorPtr)
	return rv
}/* debug [instance_methods/method]: WriteToURLOptionsError */


// Writes the data object’s bytes to the file specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:atomically:)
func (d_ Data) WriteToFileAtomically(path IString, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToFile:atomically:"), path, useAuxiliaryFile)
	return rv
}/* debug [instance_methods/method]: WriteToFileAtomically */


// Writes the data object’s bytes to the file specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:options:)
func (d_ Data) WriteToFileOptionsError(path IString, writeOptionsMask DataWritingOptions, errorPtr IError) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToFile:options:error:"), path, writeOptionsMask, errorPtr)
	return rv
}/* debug [instance_methods/method]: WriteToFileOptionsError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Data */

// A pointer to the data object’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/bytes
func (d_ Data) Bytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("bytes"))
	return rv
}/* debug [instance_properties/getter]: bytes */


// A string that contains a hexadecimal representation of the data object’s contents in a property list format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/description
func (d_ Data) Description() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("description"))
	return rv
}/* debug [instance_properties/getter]: description */


// The number of bytes contained by the data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/length
func (d_ Data) Length() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (d_ Data) NSCompressionErrorMaximum() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionErrorMaximum */


// The end of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrormaximum-swift.var
func (d_ Data) SetNSCompressionErrorMaximum(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSCompressionErrorMaximum */


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (d_ Data) NSCompressionErrorMinimum() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionErrorMinimum */


// The start of the range of error codes reserved for compression errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionerrorminimum-swift.var
func (d_ Data) SetNSCompressionErrorMinimum(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSCompressionErrorMinimum */


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (d_ Data) NSCompressionFailedError() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSCompressionFailedError"))
	return rv
}/* debug [instance_properties/getter]: NSCompressionFailedError */


// An error code value that indicates a failure to compress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompressionfailederror-swift.var
func (d_ Data) SetNSCompressionFailedError(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSCompressionFailedError:"), value)
}/* debug [instance_properties/setter]: NSCompressionFailedError */


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (d_ Data) NSDecompressionFailedError() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSDecompressionFailedError"))
	return rv
}/* debug [instance_properties/getter]: NSDecompressionFailedError */


// An error code value that indicates a failure to decompress data using the provided algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecompressionfailederror-swift.var
func (d_ Data) SetNSDecompressionFailedError(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSDecompressionFailedError:"), value)
}/* debug [instance_properties/setter]: NSDecompressionFailedError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSData */


