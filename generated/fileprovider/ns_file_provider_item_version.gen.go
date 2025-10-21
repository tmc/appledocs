// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [FileProviderItemVersion] class.
type IFileProviderItemVersion interface {
	objectivec.IObject
}

// The version of the item’s content and its metadata.
//
// Each item has a separate version object for its metadata and its content. As a result, the file provider can update an item’s metadata without uploading or downloading a new copy of its content.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FileProviderItemVersionClass) Alloc() FileProviderItemVersion {
	rv := objc.Send[FileProviderItemVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new version object.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/init(contentVersion:metadataVersion:)
func NewFileProviderItemVersionWithContentVersionMetadataVersion(contentVersion unsafe.Pointer, metadataVersion unsafe.Pointer) FileProviderItemVersion {
	instance := getFileProviderItemVersionClass().Alloc()
	rv := objc.Send[FileProviderItemVersion](instance.ID, objc.Sel("initWithContentVersion:metadataVersion:"), contentVersion, metadataVersion)
	rv.Autorelease()
	return rv
}


// A Boolean value indicating that this version predates the version returned by the file provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/beforeFirstSyncComponent
func (fc _FileProviderItemVersionClass) BeforeFirstSyncComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("beforeFirstSyncComponent"))
	return rv
}
// A Boolean value indicating that this version predates the version returned by the file provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/beforeFirstSyncComponent
func (f_ FileProviderItemVersion) BeforeFirstSyncComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("beforeFirstSyncComponent"))
	return rv
}

// An opaque object used to track versions of the item’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/contentVersion
func (f_ FileProviderItemVersion) ContentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contentVersion"))
	return rv
}

// An opaque object used to track versions of the item’s metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemVersion/metadataVersion
func (f_ FileProviderItemVersion) MetadataVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("metadataVersion"))
	return rv
}


