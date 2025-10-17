// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Data] class.
var DataClass = _DataClass{objc.GetClass("NSData")}

type _DataClass struct {
	class objc.Class
}

type Data struct {
	objc.ID
}

func DataFrom(ptr unsafe.Pointer) Data {
	return Data{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DataClass) Alloc() Data {
	rv := objc.Send[Data](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return DataClass.New()
}
// Initializes a data object with the given Base64 encoded data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64EncodedData:options:)
func NewDataWithBase64EncodedDataOptions(base64Data unsafe.Pointer, options unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64EncodedData:options:"), base64Data, options)
	rv.Autorelease()
	return rv
}
// Initializes a data object with the given Base64 encoded string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64EncodedString:options:)
func NewDataWithBase64EncodedStringOptions(base64String string, options unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64EncodedString:options:"), base64String, options)
	rv.Autorelease()
	return rv
}
// Initializes a data object initialized with the given Base64 encoded string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(base64Encoding:)
func NewDataWithBase64Encoding(base64String string) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBase64Encoding:"), base64String)
	rv.Autorelease()
	return rv
}
// Initializes a data object filled with a given number of bytes copied from a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytes:length:)
func NewDataWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytes:length:"), bytes, length)
	rv.Autorelease()
	return rv
}
// Initializes a data object filled with a given number of bytes of data from a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:)
func NewDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:"), bytes, length)
	rv.Autorelease()
	return rv
}
// Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:deallocator:)
func NewDataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	rv.Autorelease()
	return rv
}
// Initializes a newly allocated data object by adding the given number of bytes from the given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(bytesNoCopy:length:freeWhenDone:)
func NewDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	rv.Autorelease()
	return rv
}
// Initializes a data object with the content of the file at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:)
func NewDataWithContentsOfFile(path string) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}
// Initializes a data object with the content of the file at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfFile:options:)
func NewDataWithContentsOfFileOptionsError(path string, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfFile:options:error:"), path, readOptionsMask, errorPtr)
	rv.Autorelease()
	return rv
}
// Initializes a data object with the contents of the mapped file specified by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfMappedFile:)
func NewDataWithContentsOfMappedFile(path string) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfMappedFile:"), path)
	rv.Autorelease()
	return rv
}
// Creates a data object from the data at the specified file URL, or returns if the system can’t create one. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:)-6rrnr
func NewDataWithContentsOfURL(url unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}
// Creates a data object from the data at the provided file URL using specific reading options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:options:)-5abi3
func NewDataWithContentsOfURLOptionsError(url unsafe.Pointer, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, readOptionsMask, errorPtr)
	rv.Autorelease()
	return rv
}
// Initializes a data object with the contents of another data object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(data:)
func NewDataWithData(data unsafe.Pointer) Data {
	instance := DataClass.Alloc()
	rv := objc.Send[Data](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// Creates an empty data object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/data
func (dc _DataClass) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("data"))
	return rv
}
// Creates a data object containing a given number of bytes copied from a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytes:length:
func (dc _DataClass) DataWithBytesLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithBytes:length:"), bytes, length)
	return rv
}
// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:
func (dc _DataClass) DataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithBytesNoCopy:length:"), bytes, length)
	return rv
}
// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithBytesNoCopy:length:freeWhenDone:
func (dc _DataClass) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithBytesNoCopy:length:freeWhenDone:"), bytes, length, b)
	return rv
}
// Creates a data object by reading every byte from the file at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:
func (dc _DataClass) DataWithContentsOfFile(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithContentsOfFile:"), path)
	return rv
}
// Creates a data object by reading every byte from the file at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfFile:options:error:
func (dc _DataClass) DataWithContentsOfFileOptionsError(path string, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithContentsOfFile:options:error:"), path, readOptionsMask, errorPtr)
	return rv
}
// Creates a data object from the mapped file at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithContentsOfMappedFile(_:)
func (dc _DataClass) DataWithContentsOfMappedFile(path string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("dataWithContentsOfMappedFile:"), path)
	return rv
}
// Creates a data object containing the contents of another data object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/dataWithData:
func (dc _DataClass) DataWithData(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithData:"), data)
	return rv
}
// Creates a data object from the data at the specified file URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:)-6foqd
func (dc _DataClass) DataWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithContentsOfURL:"), url)
	return rv
}
// Creates a data object from the data at the provided file URL using specific reading options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/init(contentsOfURL:options:)-95rht
func (dc _DataClass) DataWithContentsOfURLOptionsError(url unsafe.Pointer, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dataWithContentsOfURL:options:error:"), url, readOptionsMask, errorPtr)
	return rv
}
// Creates a Base64, UTF-8 encoded data object from the string using the given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedData(options:)
func (d_ Data) Base64EncodedDataWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("base64EncodedDataWithOptions:"), options)
	return rv
}
// Creates a Base64 encoded string from the string using the given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/base64EncodedString(options:)
func (d_ Data) Base64EncodedStringWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("base64EncodedStringWithOptions:"), options)
	return rv
}
// Initializes a Base64 encoded string from the string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/base64Encoding()
func (d_ Data) Base64Encoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("base64Encoding"))
	return rv
}
// Returns a new data object by compressing the data object’s bytes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/compressed(using:)
func (d_ Data) CompressedDataUsingAlgorithmError(algorithm unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("compressedDataUsingAlgorithm:error:"), algorithm, error)
	return rv
}
// Returns a new data object by decompressing data object’s bytes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/decompressed(using:)
func (d_ Data) DecompressedDataUsingAlgorithmError(algorithm unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("decompressedDataUsingAlgorithm:error:"), algorithm, error)
	return rv
}
// Enumerates each range of bytes in the data object using a block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/enumerateBytes(_:)
func (d_ Data) EnumerateByteRangesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateByteRangesUsingBlock:"), block)
}
// Copies a data object’s contents into a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:)
func (d_ Data) GetBytes(buffer unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getBytes:"), buffer)
}
// Copies a number of bytes from the start of the data object into a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:length:)
func (d_ Data) GetBytesLength(buffer unsafe.Pointer, length uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getBytes:length:"), buffer, length)
}
// Copies a range of bytes from the data object into a given buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/getBytes(_:range:)
func (d_ Data) GetBytesRange(buffer unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getBytes:range:"), buffer, range_)
}
// Returns a Boolean value indicating whether this data object is the same as another. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/isEqual(to:)
func (d_ Data) IsEqualToData(other unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToData:"), other)
	return rv
}
// Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/range(of:options:in:)
func (d_ Data) RangeOfDataOptionsRange(dataToFind unsafe.Pointer, mask unsafe.Pointer, searchRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("rangeOfData:options:range:"), dataToFind, mask, searchRange)
	return rv
}
// Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/subdata(with:)
func (d_ Data) SubdataWithRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("subdataWithRange:"), range_)
	return rv
}
// Writes the data object’s bytes to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(to:atomically:)
func (d_ Data) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}
// Writes the data object’s bytes to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(to:options:)
func (d_ Data) WriteToURLOptionsError(url unsafe.Pointer, writeOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:options:error:"), url, writeOptionsMask, errorPtr)
	return rv
}
// Writes the data object’s bytes to the file specified by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:atomically:)
func (d_ Data) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToFile:atomically:"), path, useAuxiliaryFile)
	return rv
}
// Writes the data object’s bytes to the file specified by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/write(toFile:options:)
func (d_ Data) WriteToFileOptionsError(path string, writeOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToFile:options:error:"), path, writeOptionsMask, errorPtr)
	return rv
}


