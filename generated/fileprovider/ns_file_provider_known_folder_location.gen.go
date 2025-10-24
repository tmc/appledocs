// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderKnownFolderLocation */


/* debug [class_header]: Header for NSFileProviderKnownFolderLocation */
// The class instance for the [FileProviderKnownFolderLocation] class.
var (
	FileProviderKnownFolderLocationClass     _FileProviderKnownFolderLocationClass
	FileProviderKnownFolderLocationClassOnce sync.Once
)

func getFileProviderKnownFolderLocationClass() _FileProviderKnownFolderLocationClass {
	FileProviderKnownFolderLocationClassOnce.Do(func() {
		FileProviderKnownFolderLocationClass = _FileProviderKnownFolderLocationClass{objc.GetClass("NSFileProviderKnownFolderLocation")}
	})
	return FileProviderKnownFolderLocationClass
}

type _FileProviderKnownFolderLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderKnownFolderLocation */
// An interface definition for the [FileProviderKnownFolderLocation] class.
type IFileProviderKnownFolderLocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderKnownFolderLocation */
	// properties:
	DesktopLocation() IFileProviderKnownFolderLocation
	SetDesktopLocation(value IFileProviderKnownFolderLocation)
	DocumentsLocation() IFileProviderKnownFolderLocation
	SetDocumentsLocation(value IFileProviderKnownFolderLocation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderKnownFolderLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderKnownFolderLocation */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderKnownFolderLocationClass) Alloc() FileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderKnownFolderLocationClass) New() FileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderKnownFolderLocation) Init() FileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderKnownFolderLocation) Autorelease() FileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderKnownFolderLocation creates a new FileProviderKnownFolderLocation instance.
func NewFileProviderKnownFolderLocation() FileProviderKnownFolderLocation {
	return getFileProviderKnownFolderLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderKnownFolderLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location
type FileProviderKnownFolderLocation struct {
	objectivec.Object
}

// FileProviderKnownFolderLocationFrom constructs a [FileProviderKnownFolderLocation] from an unsafe.Pointer.
func FileProviderKnownFolderLocationFrom(ptr unsafe.Pointer) FileProviderKnownFolderLocation {
	return FileProviderKnownFolderLocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderKnownFolderLocation */

// Initialize a location with the item identifier of a folder that already exists on the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(existingItemIdentifier:)
func NewFileProviderKnownFolderLocationWithExistingItemIdentifier(existingItemIdentifier FileProviderItemIdentifier /* typedef */) FileProviderKnownFolderLocation {
	instance := getFileProviderKnownFolderLocationClass().Alloc()
	rv := objc.Send[FileProviderKnownFolderLocation](instance.ID, objc.Sel("initWithExistingItemIdentifier:"), existingItemIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileProviderKnownFolderLocationWithExistingItemIdentifier */


// Initialize a location with the filename of the folder in a specified parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(parentItemIdentifier:filename:)
func NewFileProviderKnownFolderLocationWithParentItemIdentifierFilename(parentItemIdentifier FileProviderItemIdentifier /* typedef */, filename objc.IObject /* cross-framework: NSString */) FileProviderKnownFolderLocation {
	instance := getFileProviderKnownFolderLocationClass().Alloc()
	rv := objc.Send[FileProviderKnownFolderLocation](instance.ID, objc.Sel("initWithParentItemIdentifier:filename:"), parentItemIdentifier, filename)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileProviderKnownFolderLocationWithParentItemIdentifierFilename */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderKnownFolderLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderKnownFolderLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderKnownFolderLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderKnownFolderLocation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocation) DesktopLocation() IFileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("desktopLocation"))
	return rv
}/* debug [instance_properties/getter]: desktopLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocation) SetDesktopLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDesktopLocation:"), value)
}/* debug [instance_properties/setter]: desktopLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocation) DocumentsLocation() IFileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("documentsLocation"))
	return rv
}/* debug [instance_properties/getter]: documentsLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocation) SetDocumentsLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentsLocation:"), value)
}/* debug [instance_properties/setter]: documentsLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderKnownFolderLocation */


