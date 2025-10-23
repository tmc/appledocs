// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderManager] class.
var (
	FileProviderManagerClass     _FileProviderManagerClass
	FileProviderManagerClassOnce sync.Once
)

func getFileProviderManagerClass() _FileProviderManagerClass {
	FileProviderManagerClassOnce.Do(func() {
		FileProviderManagerClass = _FileProviderManagerClass{objc.GetClass("NSFileProviderManager")}
	})
	return FileProviderManagerClass
}

type _FileProviderManagerClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderManager] class.
type IFileProviderManager interface {
	objectivec.IObject
	DocumentStorageURL() foundation.URL
	SetDocumentStorageURL(value foundation.URL)
	ProviderIdentifier() string
	SetProviderIdentifier(value string)
}

// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.


// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager
type FileProviderManager struct {
	objectivec.Object
}

// FileProviderManagerFrom constructs a [FileProviderManager] from an unsafe.Pointer.
//
// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
func FileProviderManagerFrom(ptr unsafe.Pointer) FileProviderManager {
	return FileProviderManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderManagerClass) Alloc() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderManagerClass) New() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderManager) Init() FileProviderManager {
	rv := objc.Send[FileProviderManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderManager) Autorelease() FileProviderManager {
	rv := objc.Send[FileProviderManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderManager creates a new FileProviderManager instance.
func NewFileProviderManager() FileProviderManager {
	return getFileProviderManagerClass().New()
}



// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/documentstorageurl
func (f_ FileProviderManager) DocumentStorageURL() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}


// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/documentstorageurl
func (f_ FileProviderManager) SetDocumentStorageURL(value foundation.URL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentStorageURL:"), value)
}


// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/provideridentifier
func (f_ FileProviderManager) ProviderIdentifier() string {
	rv := objc.Send[string](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}


// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/provideridentifier
func (f_ FileProviderManager) SetProviderIdentifier(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setProviderIdentifier:"), objc.String(value))
}



