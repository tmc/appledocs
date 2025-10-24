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
	// properties:
	DocumentStorageURL() objc.IObject /* cross-framework: URL */
	SetDocumentStorageURL(value objc.IObject /* cross-framework: URL */)
	ProviderIdentifier() objc.IObject /* cross-framework: NSString */
	SetProviderIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
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



// Adds a domain to the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/add(_:completionHandler:)
func (fc _FileProviderManagerClass) AddDomainCompletionHandler(domain IFileProviderDomain, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("addDomain:completionHandler:"), domain, completionHandler)
}


// Returns all of the File Provider extension’s domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getDomainsWithCompletionHandler(_:)
func (fc _FileProviderManagerClass) GetDomainsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getDomainsWithCompletionHandler:"), completionHandler)
}


// Creates a new domain that takes ownership of on-disk data that your app previously managed without a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/import(_:fromDirectoryAt:completionHandler:)
func (fc _FileProviderManagerClass) ImportDomainFromDirectoryAtURLCompletionHandler(domain IFileProviderDomain, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("importDomain:fromDirectoryAtURL:completionHandler:"), domain, url, completionHandler)
}


// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/documentstorageurl
func (f_ FileProviderManager) DocumentStorageURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}


// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/documentstorageurl
func (f_ FileProviderManager) SetDocumentStorageURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDocumentStorageURL:"), value)
}


// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/provideridentifier
func (f_ FileProviderManager) ProviderIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}


// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidermanager/provideridentifier
func (f_ FileProviderManager) SetProviderIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setProviderIdentifier:"), value)
}



