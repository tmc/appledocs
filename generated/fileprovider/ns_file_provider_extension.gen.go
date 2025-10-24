// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderExtension] class.
var (
	FileProviderExtensionClass     _FileProviderExtensionClass
	FileProviderExtensionClassOnce sync.Once
)

func getFileProviderExtensionClass() _FileProviderExtensionClass {
	FileProviderExtensionClassOnce.Do(func() {
		FileProviderExtensionClass = _FileProviderExtensionClass{objc.GetClass("NSFileProviderExtension")}
	})
	return FileProviderExtensionClass
}

type _FileProviderExtensionClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderExtension] class.
type IFileProviderExtension interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The principal class for the nonreplicated File Provider extension.
//
// To create a nonreplicated File Provider extension, start by creating a subclass of the class. When implementing your subclass, remember: Override all of the extension’s methods (except the deprecated methods), even if your implementation is only an empty method. Use your method implementations to provide access to the documents and folders managed by your file provider. Don’t call in your method implementations. Don’t use the class in macOS. Instead, create an subclass that adopts the and protocols. For more information, see .


// The principal class for the nonreplicated File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension
type FileProviderExtension struct {
	objectivec.Object
}

// FileProviderExtensionFrom constructs a [FileProviderExtension] from an unsafe.Pointer.
//
// The principal class for the nonreplicated File Provider extension.
func FileProviderExtensionFrom(ptr unsafe.Pointer) FileProviderExtension {
	return FileProviderExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderExtensionClass) Alloc() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderExtensionClass) New() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderExtension) Init() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderExtension) Autorelease() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderExtension creates a new FileProviderExtension instance.
func NewFileProviderExtension() FileProviderExtension {
	return getFileProviderExtensionClass().New()
}



// Returns a placeholder URL for a given document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/placeholderURL(for:)
func (fc _FileProviderExtensionClass) PlaceholderURLForURL(url objc.IObject /* cross-framework: NSURL */) objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}


// Writes a document placeholder with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/writePlaceholder(at:withMetadata:)
func (fc _FileProviderExtensionClass) WritePlaceholderAtURLWithMetadataError(placeholderURL objc.IObject /* cross-framework: NSURL */, metadata foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}


