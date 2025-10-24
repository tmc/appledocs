// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderKnownFolderLocations */


/* debug [class_header]: Header for NSFileProviderKnownFolderLocations */
// The class instance for the [FileProviderKnownFolderLocations] class.
var (
	FileProviderKnownFolderLocationsClass     _FileProviderKnownFolderLocationsClass
	FileProviderKnownFolderLocationsClassOnce sync.Once
)

func getFileProviderKnownFolderLocationsClass() _FileProviderKnownFolderLocationsClass {
	FileProviderKnownFolderLocationsClassOnce.Do(func() {
		FileProviderKnownFolderLocationsClass = _FileProviderKnownFolderLocationsClass{objc.GetClass("NSFileProviderKnownFolderLocations")}
	})
	return FileProviderKnownFolderLocationsClass
}

type _FileProviderKnownFolderLocationsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderKnownFolderLocations */
// An interface definition for the [FileProviderKnownFolderLocations] class.
type IFileProviderKnownFolderLocations interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderKnownFolderLocations */
	// properties:
	DesktopLocation() IFileProviderKnownFolderLocation
	SetDesktopLocation(value IFileProviderKnownFolderLocation)
	DocumentsLocation() IFileProviderKnownFolderLocation
	SetDocumentsLocation(value IFileProviderKnownFolderLocation)
	ShouldCreateBinaryCompatibilitySymlink() bool
	SetShouldCreateBinaryCompatibilitySymlink(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderKnownFolderLocations */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderKnownFolderLocations */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderKnownFolderLocationsClass) Alloc() FileProviderKnownFolderLocations {
	rv := objc.Send[FileProviderKnownFolderLocations](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderKnownFolderLocationsClass) New() FileProviderKnownFolderLocations {
	rv := objc.Send[FileProviderKnownFolderLocations](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderKnownFolderLocations) Init() FileProviderKnownFolderLocations {
	rv := objc.Send[FileProviderKnownFolderLocations](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderKnownFolderLocations) Autorelease() FileProviderKnownFolderLocations {
	rv := objc.Send[FileProviderKnownFolderLocations](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderKnownFolderLocations creates a new FileProviderKnownFolderLocations instance.
func NewFileProviderKnownFolderLocations() FileProviderKnownFolderLocations {
	return getFileProviderKnownFolderLocationsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderKnownFolderLocations */
// A class for working with known-folder locations.


// A class for working with known-folder locations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations
type FileProviderKnownFolderLocations struct {
	objectivec.Object
}

// FileProviderKnownFolderLocationsFrom constructs a [FileProviderKnownFolderLocations] from an unsafe.Pointer.
//
// A class for working with known-folder locations.
func FileProviderKnownFolderLocationsFrom(ptr unsafe.Pointer) FileProviderKnownFolderLocations {
	return FileProviderKnownFolderLocations{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderKnownFolderLocations */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderKnownFolderLocations */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderKnownFolderLocations */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderKnownFolderLocations */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderKnownFolderLocations */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/desktopLocation
func (f_ FileProviderKnownFolderLocations) DesktopLocation() IFileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("desktopLocation"))
	return rv
}/* debug [instance_properties/getter]: desktopLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/desktopLocation
func (f_ FileProviderKnownFolderLocations) SetDesktopLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDesktopLocation:"), value)
}/* debug [instance_properties/setter]: desktopLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/documentsLocation
func (f_ FileProviderKnownFolderLocations) DocumentsLocation() IFileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("documentsLocation"))
	return rv
}/* debug [instance_properties/getter]: documentsLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/documentsLocation
func (f_ FileProviderKnownFolderLocations) SetDocumentsLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentsLocation:"), value)
}/* debug [instance_properties/setter]: documentsLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/shouldCreateBinaryCompatibilitySymlink
func (f_ FileProviderKnownFolderLocations) ShouldCreateBinaryCompatibilitySymlink() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("shouldCreateBinaryCompatibilitySymlink"))
	return rv
}/* debug [instance_properties/getter]: shouldCreateBinaryCompatibilitySymlink */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/shouldCreateBinaryCompatibilitySymlink
func (f_ FileProviderKnownFolderLocations) SetShouldCreateBinaryCompatibilitySymlink(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setShouldCreateBinaryCompatibilitySymlink:"), value)
}/* debug [instance_properties/setter]: shouldCreateBinaryCompatibilitySymlink */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderKnownFolderLocations */


