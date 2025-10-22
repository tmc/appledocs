// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderDomainVersion] class.
var (
	FileProviderDomainVersionClass     _FileProviderDomainVersionClass
	FileProviderDomainVersionClassOnce sync.Once
)

func getFileProviderDomainVersionClass() _FileProviderDomainVersionClass {
	FileProviderDomainVersionClassOnce.Do(func() {
		FileProviderDomainVersionClass = _FileProviderDomainVersionClass{objc.GetClass("NSFileProviderDomainVersion")}
	})
	return FileProviderDomainVersionClass
}

type _FileProviderDomainVersionClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderDomainVersion] class.
type IFileProviderDomainVersion interface {
	objectivec.IObject
	DomainVersion() NSFileProviderDomainVersion
	SetDomainVersion(value IFileProviderDomainVersion)
}

// An opaque object that identifies a specific version of a domain.
//
// The file provider extension is responsible for assigning and updating the domain version. To specify the domain version, adopt the protocol. The system then calls your extension’s method to read the current version. The system reads the domain version after you call: The completion handler The completion handler The completion handler The completion handler The or method when enumerating the materialized set. The system always reads the domain version on the same dispatch queue as the completion handler. Your extension defines when the domain version changes. When you update the version, call the and passing the constant as the property. This notifies the system of the update. The system ignores any lower versions. When the system discovers a change on disk, it associates that change with the current domain version. It then includes the version in the object passed to the file provider extension. Only file provider extensions based on the use instances of this class. Each version object is immutable. You can use them as keys in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion
type FileProviderDomainVersion struct {
	objectivec.Object
}

// FileProviderDomainVersionFrom constructs a [FileProviderDomainVersion] from an unsafe.Pointer.
//
// An opaque object that identifies a specific version of a domain.
func FileProviderDomainVersionFrom(ptr unsafe.Pointer) FileProviderDomainVersion {
	return FileProviderDomainVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderDomainVersionClass) Alloc() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderDomainVersionClass) New() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderDomainVersion) Init() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderDomainVersion) Autorelease() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderDomainVersion creates a new FileProviderDomainVersion instance.
func NewFileProviderDomainVersion() FileProviderDomainVersion {
	return getFileProviderDomainVersionClass().New()
}


// An opaque object that uniquely identifies the domain’s version.
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/domainversion
func (f_ FileProviderDomainVersion) DomainVersion() NSFileProviderDomainVersion {
	rv := objc.Send[NSFileProviderDomainVersion](f_.ID, objc.Sel("domainVersion"))
	return rv
}


// SetDomainVersion sets the value of the domainVersion property.
// An opaque object that uniquely identifies the domain’s version.

//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/domainversion
func (f_ FileProviderDomainVersion) SetDomainVersion(value IFileProviderDomainVersion) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDomainVersion:"), value)
}



