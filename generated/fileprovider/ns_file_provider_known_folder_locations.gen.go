// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FileProviderKnownFolderLocations] class.
type IFileProviderKnownFolderLocations interface {
	objectivec.IObject
	// properties:
	DesktopLocation() objc.IObject /* cross-framework: FileProviderKnownFolderLocation */
	SetDesktopLocation(value objc.IObject /* cross-framework: FileProviderKnownFolderLocation */)
	DocumentsLocation() objc.IObject /* cross-framework: FileProviderKnownFolderLocation */
	SetDocumentsLocation(value objc.IObject /* cross-framework: FileProviderKnownFolderLocation */)
	ShouldCreateBinaryCompatibilitySymlink() bool
	SetShouldCreateBinaryCompatibilitySymlink(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (fc _FileProviderKnownFolderLocationsClass) Alloc() FileProviderKnownFolderLocations {
	rv := objc.Send[FileProviderKnownFolderLocations](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocations) DesktopLocation() objc.IObject /* cross-framework: FileProviderKnownFolderLocation */ {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("desktopLocation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocations) SetDesktopLocation(value objc.IObject /* cross-framework: FileProviderKnownFolderLocation */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDesktopLocation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocations) DocumentsLocation() objc.IObject /* cross-framework: FileProviderKnownFolderLocation */ {
	rv := objc.Send[FileProviderKnownFolderLocation](f_.ID, objc.Sel("documentsLocation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocations) SetDocumentsLocation(value objc.IObject /* cross-framework: FileProviderKnownFolderLocation */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentsLocation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/shouldcreatebinarycompatibilitysymlink
func (f_ FileProviderKnownFolderLocations) ShouldCreateBinaryCompatibilitySymlink() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("shouldCreateBinaryCompatibilitySymlink"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/shouldcreatebinarycompatibilitysymlink
func (f_ FileProviderKnownFolderLocations) SetShouldCreateBinaryCompatibilitySymlink(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setShouldCreateBinaryCompatibilitySymlink:"), value)
}



