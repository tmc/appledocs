// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderItemVersion */


/* debug [class_header]: Header for NSFileProviderItemVersion */
// The class instance for the [FileProviderItemVersion] class.
var (
	FileProviderItemVersionClass     _FileProviderItemVersionClass
	FileProviderItemVersionClassOnce sync.Once
)

func getFileProviderItemVersionClass() _FileProviderItemVersionClass {
	FileProviderItemVersionClassOnce.Do(func() {
		FileProviderItemVersionClass = _FileProviderItemVersionClass{objc.GetClass("NSFileProviderItemVersion")}
	})
	return FileProviderItemVersionClass
}

type _FileProviderItemVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderItemVersion */
// An interface definition for the [FileProviderItemVersion] class.
type IFileProviderItemVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderItemVersion */
	// properties:
	ContentVersion() objc.IObject /* cross-framework: NSData */
	MetadataVersion() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderItemVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderItemVersion */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderItemVersionClass) Alloc() FileProviderItemVersion {
	rv := objc.Send[FileProviderItemVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderItemVersionClass) New() FileProviderItemVersion {
	rv := objc.Send[FileProviderItemVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderItemVersion) Init() FileProviderItemVersion {
	rv := objc.Send[FileProviderItemVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderItemVersion) Autorelease() FileProviderItemVersion {
	rv := objc.Send[FileProviderItemVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderItemVersion creates a new FileProviderItemVersion instance.
func NewFileProviderItemVersion() FileProviderItemVersion {
	return getFileProviderItemVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderItemVersion */
// The version of the item’s content and its metadata.
//
// Each item has a separate version object for its metadata and its content. As a result, the file provider can update an item’s metadata without uploading or downloading a new copy of its content.


// The version of the item’s content and its metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion
type FileProviderItemVersion struct {
	objectivec.Object
}

// FileProviderItemVersionFrom constructs a [FileProviderItemVersion] from an unsafe.Pointer.
//
// The version of the item’s content and its metadata.
func FileProviderItemVersionFrom(ptr unsafe.Pointer) FileProviderItemVersion {
	return FileProviderItemVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderItemVersion */

// Creates a new version object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/init(contentVersion:metadataVersion:)
func NewFileProviderItemVersionWithContentVersionMetadataVersion(contentVersion objc.IObject /* cross-framework: NSData */, metadataVersion objc.IObject /* cross-framework: NSData */) FileProviderItemVersion {
	instance := getFileProviderItemVersionClass().Alloc()
	rv := objc.Send[FileProviderItemVersion](instance.ID, objc.Sel("initWithContentVersion:metadataVersion:"), contentVersion, metadataVersion)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileProviderItemVersionWithContentVersionMetadataVersion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderItemVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderItemVersion */

// A Boolean value indicating that this version predates the version returned by the file provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/beforeFirstSyncComponent
func (fc _FileProviderItemVersionClass) BeforeFirstSyncComponent() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](objc.ID(fc.class), objc.Sel("beforeFirstSyncComponent"))
	return rv
}/* debug [class_properties_class/property]: beforeFirstSyncComponent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderItemVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderItemVersion */

// A Boolean value indicating that this version predates the version returned by the file provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/beforeFirstSyncComponent
func (f_ FileProviderItemVersion) BeforeFirstSyncComponent() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("beforeFirstSyncComponent"))
	return rv
}/* debug [instance_properties/getter]: beforeFirstSyncComponent */


// An opaque object used to track versions of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/contentVersion
func (f_ FileProviderItemVersion) ContentVersion() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("contentVersion"))
	return rv
}/* debug [instance_properties/getter]: contentVersion */


// An opaque object used to track versions of the item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/metadataVersion
func (f_ FileProviderItemVersion) MetadataVersion() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("metadataVersion"))
	return rv
}/* debug [instance_properties/getter]: metadataVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderItemVersion */


