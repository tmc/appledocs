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

/* debug [class.gen.go]: Generating class NSFileProviderExtension */


/* debug [class_header]: Header for NSFileProviderExtension */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderExtension */
// An interface definition for the [FileProviderExtension] class.
type IFileProviderExtension interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderExtension */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderExtension */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderExtension */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderExtensionClass) Alloc() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderExtension */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderExtension *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderExtension */

// Returns a placeholder URL for a given document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/placeholderURL(for:)
func (fc _FileProviderExtensionClass) PlaceholderURLForURL(url objc.IObject /* cross-framework: NSURL */) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlaceholderURLForURL) */


// Writes a document placeholder with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/writePlaceholder(at:withMetadata:)
func (fc _FileProviderExtensionClass) WritePlaceholderAtURLWithMetadataError(placeholderURL objc.IObject /* cross-framework: NSURL */, metadata foundation.IDictionary, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WritePlaceholderAtURLWithMetadataError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderExtension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderExtension */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderExtension */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderExtension */


