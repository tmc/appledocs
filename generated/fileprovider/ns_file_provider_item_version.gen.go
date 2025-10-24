// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ContentVersion() objc.IObject /* cross-framework: Data */
	SetContentVersion(value objc.IObject /* cross-framework: Data */)
	MetadataVersion() objc.IObject /* cross-framework: Data */
	SetMetadataVersion(value objc.IObject /* cross-framework: Data */)
	// methods:
}

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



// An opaque object used to track versions of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/contentversion
func (f_ FileProviderItemVersion) ContentVersion() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](f_.ID, objc.Sel("contentVersion"))
	return rv
}


// An opaque object used to track versions of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/contentversion
func (f_ FileProviderItemVersion) SetContentVersion(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setContentVersion:"), value)
}


// An opaque object used to track versions of the item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/metadataversion
func (f_ FileProviderItemVersion) MetadataVersion() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](f_.ID, objc.Sel("metadataVersion"))
	return rv
}


// An opaque object used to track versions of the item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/metadataversion
func (f_ FileProviderItemVersion) SetMetadataVersion(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMetadataVersion:"), value)
}



