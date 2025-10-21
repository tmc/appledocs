// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileVersion] class.
var (
	FileVersionClass     _FileVersionClass
	FileVersionClassOnce sync.Once
)

func getFileVersionClass() _FileVersionClass {
	FileVersionClassOnce.Do(func() {
		FileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}
	})
	return FileVersionClass
}

type _FileVersionClass struct {
	class objc.Class
}

// An interface definition for the [FileVersion] class.
type IFileVersion interface {
	objectivec.IObject
}

// A snapshot of a file at a specific point in time.
//
// Use the methods of this class to access, create, and manage file revisions in your app. Each file version instance contains metadata about a single revision, including the location of the associated file, the modification date of the revision, and whether the revision is discardable. In Mac apps, you can use file version objects to track changes to a local file over time and to prevent the loss of data during editing. When managing local versions, the document architecture creates versions at specific points in the lifetime of your application. Your application can also create versions explicitly at times that your application designates as appropriate. In addition to managing local files, the system also uses this class to manage cloud-based files. For files in the cloud, there is usually only one version of the file at any given time. However, additional file versions may be created in cases where two different computers attempt to save the file to the cloud at the same time. In that case, one file is chosen as the current version and any other versions are tagged as being in conflict with the original. Conflict versions are reported to the appropriate file presenter objects and should be resolved as soon as possible so that the corresponding files can be removed from the cloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion
type FileVersion struct {
	objectivec.Object
}

// FileVersionFrom constructs a [FileVersion] from an unsafe.Pointer.
//
// A snapshot of a file at a specific point in time.
func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileVersionClass) Alloc() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileVersionClass) New() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileVersion) Init() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileVersion) Autorelease() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileVersion creates a new FileVersion instance.
func NewFileVersion() FileVersion {
	return getFileVersionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasLocalContents
func (f_ FileVersion) HasLocalContents() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasLocalContents"))
	return rv
}

// The identifier for this version of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/persistentIdentifier
func (f_ FileVersion) PersistentIdentifier() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("persistentIdentifier"))
	return rv
}

// The URL identifying the location of the file associated with the file version object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/url
func (f_ FileVersion) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URL"))
	return rv
}



