// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class MEByteSource */


/* debug [class_header]: Header for MEByteSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEByteSource */
// An interface definition for the [MEByteSource] class.
type IMEByteSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEByteSource */
	// properties:
	ContentType() uniformtypeidentifiers.UTType
	FileLength() int64
	FileName() objc.IObject /* cross-framework: NSString */
	RelatedFileNamesInSameDirectory() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEByteSource */
	// methods:
	AvailableLengthAtOffset(offset int64) int64
	ByteSourceForRelatedFileNameError(fileName objc.IObject /* cross-framework: NSString */, errorOut unsafe.Pointer) IMEByteSource
	ReadDataOfLengthFromOffsetCompletionHandler(length uintptr /* not a class type */, offset int64, completionHandler unsafe.Pointer)
	ReadDataOfLengthFromOffsetToDestinationBytesReadError(length uintptr /* not a class type */, offset int64, dest unsafe.Pointer, bytesReadOut uintptr /* not a class type */, error_ unsafe.Pointer) bool
	ReadDataOfLengthFromOffsetToDestinationCompletionHandler(length uintptr /* not a class type */, offset int64, dest unsafe.Pointer, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEByteSource */
// Alloc allocates a new instance without initialization.
func (mc _MEByteSourceClass) Alloc() MEByteSource {
	rv := objc.Send[MEByteSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEByteSource */
// Provides read access to the data in a media asset file.
//
// Media Toolbox passes an instance for the media asset’s primary file when it initializes an object. The format reader may call to request additional byte sources for related files in the same directory as the primary file.


// Provides read access to the data in a media asset file.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEByteSource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEByteSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEByteSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEByteSource */

// Gets the number of available bytes from the offset within the byte source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/availableLength(at:)
func (m_ MEByteSource) AvailableLengthAtOffset(offset int64) int64 {
	rv := objc.Send[int64](m_.ID, objc.Sel("availableLengthAtOffset:"), offset)
	return rv
}/* debug [instance_methods/method]: AvailableLengthAtOffset */


// Creates a new byte source for a related file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/byteSourceForRelatedFileName(_:)
func (m_ MEByteSource) ByteSourceForRelatedFileNameError(fileName objc.IObject /* cross-framework: NSString */, errorOut unsafe.Pointer) MEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSourceForRelatedFileName:error:"), fileName, errorOut)
	return rv
}/* debug [instance_methods/method]: ByteSourceForRelatedFileNameError */


// Reads bytes from a byte source into a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/read(length:from:completionHandler:)
func (m_ MEByteSource) ReadDataOfLengthFromOffsetCompletionHandler(length uintptr /* not a class type */, offset int64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readDataOfLength:fromOffset:completionHandler:"), length, offset, completionHandler)
}/* debug [instance_methods/method]: ReadDataOfLengthFromOffsetCompletionHandler */


// Reads bytes from a byte source into a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/readDataOfLength:fromOffset:toDestination:bytesRead:error:
func (m_ MEByteSource) ReadDataOfLengthFromOffsetToDestinationBytesReadError(length uintptr /* not a class type */, offset int64, dest unsafe.Pointer, bytesReadOut uintptr /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readDataOfLength:fromOffset:toDestination:bytesRead:error:"), length, offset, dest, bytesReadOut, error_)
	return rv
}/* debug [instance_methods/method]: ReadDataOfLengthFromOffsetToDestinationBytesReadError */


// Reads bytes from a byte source into a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/readDataOfLength:fromOffset:toDestination:completionHandler:
func (m_ MEByteSource) ReadDataOfLengthFromOffsetToDestinationCompletionHandler(length uintptr /* not a class type */, offset int64, dest unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readDataOfLength:fromOffset:toDestination:completionHandler:"), length, offset, dest, completionHandler)
}/* debug [instance_methods/method]: ReadDataOfLengthFromOffsetToDestinationCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEByteSource */

// The format of the byte source file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/contentType
func (m_ MEByteSource) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](m_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The length of the byte source file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/fileLength
func (m_ MEByteSource) FileLength() int64 {
	rv := objc.Send[int64](m_.ID, objc.Sel("fileLength"))
	return rv
}/* debug [instance_properties/getter]: fileLength */


// The name of the file for the byte source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/fileName
func (m_ MEByteSource) FileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("fileName"))
	return rv
}/* debug [instance_properties/getter]: fileName */


// An array of related file names in the parent directory of the byte source file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEByteSource/relatedFileNamesInSameDirectory
func (m_ MEByteSource) RelatedFileNamesInSameDirectory() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("relatedFileNamesInSameDirectory"))
	return rv
}/* debug [instance_properties/getter]: relatedFileNamesInSameDirectory */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEByteSource */



