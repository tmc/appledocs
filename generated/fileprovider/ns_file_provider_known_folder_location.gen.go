// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FileProviderKnownFolderLocation] class.
type IFileProviderKnownFolderLocation interface {
	objectivec.IObject
	DesktopLocation() NSFileProviderKnownFolderLocation
	SetDesktopLocation(value IFileProviderKnownFolderLocation)
	DocumentsLocation() NSFileProviderKnownFolderLocation
	SetDocumentsLocation(value IFileProviderKnownFolderLocation)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location
type FileProviderKnownFolderLocation struct {
	objectivec.Object
}

// FileProviderKnownFolderLocationFrom constructs a [FileProviderKnownFolderLocation] from an unsafe.Pointer.
func FileProviderKnownFolderLocationFrom(ptr unsafe.Pointer) FileProviderKnownFolderLocation {
	return FileProviderKnownFolderLocation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderKnownFolderLocationClass) Alloc() FileProviderKnownFolderLocation {
	rv := objc.Send[FileProviderKnownFolderLocation](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocation) DesktopLocation() NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f_.ID, objc.Sel("desktopLocation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/desktoplocation
func (f_ FileProviderKnownFolderLocation) SetDesktopLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDesktopLocation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocation) DocumentsLocation() NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f_.ID, objc.Sel("documentsLocation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderknownfolderlocations/documentslocation
func (f_ FileProviderKnownFolderLocation) SetDocumentsLocation(value IFileProviderKnownFolderLocation) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentsLocation:"), value)
}



