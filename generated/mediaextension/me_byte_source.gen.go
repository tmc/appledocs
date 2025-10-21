// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEByteSource] class.
var (
	MEByteSourceClass     _MEByteSourceClass
	MEByteSourceClassOnce sync.Once
)

func getMEByteSourceClass() _MEByteSourceClass {
	MEByteSourceClassOnce.Do(func() {
		MEByteSourceClass = _MEByteSourceClass{objc.GetClass("MEByteSource")}
	})
	return MEByteSourceClass
}

type _MEByteSourceClass struct {
	class objc.Class
}

// An interface definition for the [MEByteSource] class.
type IMEByteSource interface {
	objectivec.IObject
	ByteSourceForRelatedFileNameError(fileName string, errorOut unsafe.Pointer) unsafe.Pointer
	ReadDataOfLengthFromOffsetToDestinationCompletionHandler(length unsafe.Pointer, offset unsafe.Pointer, dest unsafe.Pointer, completionHandler unsafe.Pointer)
}

// Provides read access to the data in a media asset file.
//
// Media Toolbox passes an instance for the media asset’s primary file when it initializes an object. The format reader may call to request additional byte sources for related files in the same directory as the primary file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource
type MEByteSource struct {
	objectivec.Object
}

// MEByteSourceFrom constructs a [MEByteSource] from an unsafe.Pointer.
//
// Provides read access to the data in a media asset file.
func MEByteSourceFrom(ptr unsafe.Pointer) MEByteSource {
	return MEByteSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEByteSourceClass) Alloc() MEByteSource {
	rv := objc.Send[MEByteSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEByteSourceClass) New() MEByteSource {
	rv := objc.Send[MEByteSource](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEByteSource) Init() MEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEByteSource) Autorelease() MEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEByteSource creates a new MEByteSource instance.
func NewMEByteSource() MEByteSource {
	return getMEByteSourceClass().New()
}


// Creates a new byte source for a related file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/byteSourceForRelatedFileName(_:)
func (m_ MEByteSource) ByteSourceForRelatedFileNameError(fileName string, errorOut unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("byteSourceForRelatedFileName:error:"), objc.String(fileName), errorOut)
	return rv
}

// Reads bytes from a byte source into a buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/readDataOfLength:fromOffset:toDestination:completionHandler:
func (m_ MEByteSource) ReadDataOfLengthFromOffsetToDestinationCompletionHandler(length unsafe.Pointer, offset unsafe.Pointer, dest unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readDataOfLength:fromOffset:toDestination:completionHandler:"), length, offset, dest, completionHandler)
}

// The name of the file for the byte source.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/filename
func (m_ MEByteSource) FileName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("fileName"))
	return rv
}


// SetFileName sets the value of the fileName property.
// The name of the file for the byte source.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/filename
func (m_ MEByteSource) SetFileName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFileName:"), objc.String(value))
}

// The format of the byte source file.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/contenttype
func (m_ MEByteSource) ContentType() UTType {
	rv := objc.Send[UTType](m_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
// The format of the byte source file.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/contenttype
func (m_ MEByteSource) SetContentType(value UTType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContentType:"), value)
}

// The length of the byte source file.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/filelength
func (m_ MEByteSource) FileLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fileLength"))
	return rv
}


// SetFileLength sets the value of the fileLength property.
// The length of the byte source file.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/filelength
func (m_ MEByteSource) SetFileLength(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFileLength:"), value)
}

// An array of related file names in the parent directory of the byte source file.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/relatedfilenamesinsamedirectory
func (m_ MEByteSource) RelatedFileNamesInSameDirectory() string {
	rv := objc.Send[string](m_.ID, objc.Sel("relatedFileNamesInSameDirectory"))
	return rv
}


// SetRelatedFileNamesInSameDirectory sets the value of the relatedFileNamesInSameDirectory property.
// An array of related file names in the parent directory of the byte source file.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mebytesource/relatedfilenamesinsamedirectory
func (m_ MEByteSource) SetRelatedFileNamesInSameDirectory(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRelatedFileNamesInSameDirectory:"), objc.String(value))
}



