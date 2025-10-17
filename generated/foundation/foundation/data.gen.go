// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Data] class.
var DataClass objc.Class

func init() {
	DataClass = objc.GetClass("NSData")
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
func (dc Data) Alloc() Data {
	ret := objc.ID(DataClass).Send(objc.RegisterName("alloc"))
	return Data{ret}
}

// Init initializes the instance.
func (d_ Data) Init() Data {
	ret := d_.ID.Send(objc.RegisterName("init"))
	return Data{ret}
}
// Initializes a data object with the given Base64 encoded data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(base64EncodedData:options:)
func NewDataWithBase64EncodedDataOptions(base64Data unsafe.Pointer, options unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBase64EncodedData:options:")
	ret := instance.ID.Send(sel, base64Data, options)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object with the given Base64 encoded string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(base64EncodedString:options:)
func NewDataWithBase64EncodedStringOptions(base64String string, options unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBase64EncodedString:options:")
	ret := instance.ID.Send(sel, base64String, options)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object initialized with the given Base64 encoded string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(base64Encoding:)
func NewDataWithBase64Encoding(base64String string) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBase64Encoding:")
	ret := instance.ID.Send(sel, base64String)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object filled with a given number of bytes copied from a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(bytes:length:)
func NewDataWithBytesLength(bytes unsafe.Pointer, length uint) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBytes:length:")
	ret := instance.ID.Send(sel, bytes, length)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object filled with a given number of bytes of data from a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(bytesNoCopy:length:)
func NewDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBytesNoCopy:length:")
	ret := instance.ID.Send(sel, bytes, length)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(bytesNoCopy:length:deallocator:)
func NewDataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBytesNoCopy:length:deallocator:")
	ret := instance.ID.Send(sel, bytes, length, deallocator)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated data object by adding the given number of bytes from the given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(bytesNoCopy:length:freeWhenDone:)
func NewDataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithBytesNoCopy:length:freeWhenDone:")
	ret := instance.ID.Send(sel, bytes, length, b)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object with the content of the file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfFile:)
func NewDataWithContentsOfFile(path string) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:")
	ret := instance.ID.Send(sel, path)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object with the content of the file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfFile:options:)
func NewDataWithContentsOfFileOptionsError(path string, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:options:error:")
	ret := instance.ID.Send(sel, path, readOptionsMask, errorPtr)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object with the contents of the mapped file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfMappedFile:)
func NewDataWithContentsOfMappedFile(path string) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfMappedFile:")
	ret := instance.ID.Send(sel, path)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a data object from the data at the specified file URL, or returns   if the system can’t create one. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfURL:)-6rrnr
func NewDataWithContentsOfURL(url unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:")
	ret := instance.ID.Send(sel, url)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a data object from the data at the provided file URL using specific reading options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfURL:options:)-5abi3
func NewDataWithContentsOfURLOptionsError(url unsafe.Pointer, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:options:error:")
	ret := instance.ID.Send(sel, url, readOptionsMask, errorPtr)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a data object with the contents of another data object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(data:)
func NewDataWithData(data unsafe.Pointer) Data {
	instance := Data{}.Alloc()
	sel := objc.RegisterName("initWithData:")
	ret := instance.ID.Send(sel, data)
	instance = Data{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates an empty data object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/data
func (dc Data) Data() unsafe.Pointer {
	sel := objc.RegisterName("data")
	ret := objc.ID(DataClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Creates a data object containing a given number of bytes copied from a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithBytes:length:
func (dc Data) DataWithBytesLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	sel := objc.RegisterName("dataWithBytes:length:")
	ret := objc.ID(DataClass).Send(sel, bytes, length)
	return unsafe.Pointer(ret)
}
// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithBytesNoCopy:length:
func (dc Data) DataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	sel := objc.RegisterName("dataWithBytesNoCopy:length:")
	ret := objc.ID(DataClass).Send(sel, bytes, length)
	return unsafe.Pointer(ret)
}
// Creates a data object that holds a given number of bytes from a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithBytesNoCopy:length:freeWhenDone:
func (dc Data) DataWithBytesNoCopyLengthFreeWhenDone(bytes unsafe.Pointer, length uint, b bool) unsafe.Pointer {
	sel := objc.RegisterName("dataWithBytesNoCopy:length:freeWhenDone:")
	ret := objc.ID(DataClass).Send(sel, bytes, length, b)
	return unsafe.Pointer(ret)
}
// Creates a data object by reading every byte from the file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithContentsOfFile:
func (dc Data) DataWithContentsOfFile(path string) unsafe.Pointer {
	sel := objc.RegisterName("dataWithContentsOfFile:")
	ret := objc.ID(DataClass).Send(sel, path)
	return unsafe.Pointer(ret)
}
// Creates a data object by reading every byte from the file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithContentsOfFile:options:error:
func (dc Data) DataWithContentsOfFileOptionsError(path string, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithContentsOfFile:options:error:")
	ret := objc.ID(DataClass).Send(sel, path, readOptionsMask, errorPtr)
	return unsafe.Pointer(ret)
}
// Creates a data object from the mapped file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithContentsOfMappedFile(_:)
func (dc Data) DataWithContentsOfMappedFile(path string) objc.ID {
	sel := objc.RegisterName("dataWithContentsOfMappedFile:")
	ret := objc.ID(DataClass).Send(sel, path)
	return ret
}
// Creates a data object containing the contents of another data object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/dataWithData:
func (dc Data) DataWithData(data unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithData:")
	ret := objc.ID(DataClass).Send(sel, data)
	return unsafe.Pointer(ret)
}
// Creates a data object from the data at the specified file URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfURL:)-6foqd
func (dc Data) DataWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithContentsOfURL:")
	ret := objc.ID(DataClass).Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates a data object from the data at the provided file URL using specific reading options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/init(contentsOfURL:options:)-95rht
func (dc Data) DataWithContentsOfURLOptionsError(url unsafe.Pointer, readOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithContentsOfURL:options:error:")
	ret := objc.ID(DataClass).Send(sel, url, readOptionsMask, errorPtr)
	return unsafe.Pointer(ret)
}
// Creates a Base64, UTF-8 encoded data object from the string using the given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/base64EncodedData(options:)
func (d_ Data) Base64EncodedDataWithOptions(options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("base64EncodedDataWithOptions:")
	ret := d_.ID.Send(sel, options)
	return unsafe.Pointer(ret)
}
// Creates a Base64 encoded string from the string using the given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/base64EncodedString(options:)
func (d_ Data) Base64EncodedStringWithOptions(options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("base64EncodedStringWithOptions:")
	ret := d_.ID.Send(sel, options)
	return unsafe.Pointer(ret)
}
// Initializes a Base64 encoded string from the string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/base64Encoding()
func (d_ Data) Base64Encoding() unsafe.Pointer {
	sel := objc.RegisterName("base64Encoding")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a new data object by compressing the data object’s bytes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/compressed(using:)
func (d_ Data) CompressedDataUsingAlgorithmError(algorithm unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("compressedDataUsingAlgorithm:error:")
	ret := d_.ID.Send(sel, algorithm, error)
	return unsafe.Pointer(ret)
}
// Returns a new data object by decompressing data object’s bytes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/decompressed(using:)
func (d_ Data) DecompressedDataUsingAlgorithmError(algorithm unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("decompressedDataUsingAlgorithm:error:")
	ret := d_.ID.Send(sel, algorithm, error)
	return unsafe.Pointer(ret)
}
// Enumerates each range of bytes in the data object using a block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/enumerateBytes(_:)
func (d_ Data) EnumerateByteRangesUsingBlock(block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateByteRangesUsingBlock:")
	d_.ID.Send(sel, block)
}
// Copies a data object’s contents into a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/getBytes(_:)
func (d_ Data) GetBytes(buffer unsafe.Pointer) {
	sel := objc.RegisterName("getBytes:")
	d_.ID.Send(sel, buffer)
}
// Copies a number of bytes from the start of the data object into a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/getBytes(_:length:)
func (d_ Data) GetBytesLength(buffer unsafe.Pointer, length uint) {
	sel := objc.RegisterName("getBytes:length:")
	d_.ID.Send(sel, buffer, length)
}
// Copies a range of bytes from the data object into a given buffer. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/getBytes(_:range:)
func (d_ Data) GetBytesRange(buffer unsafe.Pointer, range_ unsafe.Pointer) {
	sel := objc.RegisterName("getBytes:range:")
	d_.ID.Send(sel, buffer, range_)
}
// Returns a Boolean value indicating whether this data object is the same as another. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/isEqual(to:)
func (d_ Data) IsEqualToData(other unsafe.Pointer) bool {
	sel := objc.RegisterName("isEqualToData:")
	ret := d_.ID.Send(sel, other)
	return ret != 0
}
// Finds and returns the range of the first occurrence of the given data, within the given range, subject to given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/range(of:options:in:)
func (d_ Data) RangeOfDataOptionsRange(dataToFind unsafe.Pointer, mask unsafe.Pointer, searchRange unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("rangeOfData:options:range:")
	ret := d_.ID.Send(sel, dataToFind, mask, searchRange)
	return unsafe.Pointer(ret)
}
// Returns a new data object containing the data object’s bytes that fall within the limits specified by a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/subdata(with:)
func (d_ Data) SubdataWithRange(range_ unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("subdataWithRange:")
	ret := d_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Writes the data object’s bytes to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/write(to:atomically:)
func (d_ Data) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	sel := objc.RegisterName("writeToURL:atomically:")
	ret := d_.ID.Send(sel, url, atomically)
	return ret != 0
}
// Writes the data object’s bytes to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/write(to:options:)
func (d_ Data) WriteToURLOptionsError(url unsafe.Pointer, writeOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToURL:options:error:")
	ret := d_.ID.Send(sel, url, writeOptionsMask, errorPtr)
	return ret != 0
}
// Writes the data object’s bytes to the file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/write(toFile:atomically:)
func (d_ Data) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	sel := objc.RegisterName("writeToFile:atomically:")
	ret := d_.ID.Send(sel, path, useAuxiliaryFile)
	return ret != 0
}
// Writes the data object’s bytes to the file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSData/write(toFile:options:)
func (d_ Data) WriteToFileOptionsError(path string, writeOptionsMask unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToFile:options:error:")
	ret := d_.ID.Send(sel, path, writeOptionsMask, errorPtr)
	return ret != 0
}

