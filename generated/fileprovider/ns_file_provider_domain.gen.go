// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderDomain] class.
var (
	FileProviderDomainClass     _FileProviderDomainClass
	FileProviderDomainClassOnce sync.Once
)

func getFileProviderDomainClass() _FileProviderDomainClass {
	FileProviderDomainClassOnce.Do(func() {
		FileProviderDomainClass = _FileProviderDomainClass{objc.GetClass("NSFileProviderDomain")}
	})
	return FileProviderDomainClass
}

type _FileProviderDomainClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderDomain] class.
type IFileProviderDomain interface {
	objectivec.IObject
}

// A File Provider extension’s domain.
//
// You can use domains to partition a file provider’s content. When you use domains, a single file provider can act as if multiple file providers were installed, and users can dynamically switch from one domain to another. You can use domains to represent different accounts or locations. By default, a File Provider extension has no domain. You can register domains by calling the class’s method. A new instance is created for each domain that you register. The object’s property indicates which domain the file provider belongs to. Any items returned by that file provider also belong to the domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain
type FileProviderDomain struct {
	objectivec.Object
}

// FileProviderDomainFrom constructs a [FileProviderDomain] from an unsafe.Pointer.
//
// A File Provider extension’s domain.
func FileProviderDomainFrom(ptr unsafe.Pointer) FileProviderDomain {
	return FileProviderDomain{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderDomainClass) Alloc() FileProviderDomain {
	rv := objc.Send[FileProviderDomain](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderDomainClass) New() FileProviderDomain {
	rv := objc.Send[FileProviderDomain](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderDomain) Init() FileProviderDomain {
	rv := objc.Send[FileProviderDomain](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderDomain) Autorelease() FileProviderDomain {
	rv := objc.Send[FileProviderDomain](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderDomain creates a new FileProviderDomain instance.
func NewFileProviderDomain() FileProviderDomain {
	return getFileProviderDomainClass().New()
}




// Creates a new file provider domain with the specified URL and display name.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(displayName:userInfo:volumeURL:)
func NewFileProviderDomainWithDisplayNameUserInfoVolumeURL(displayName string, userInfo objc.ID, volumeURL foundation.URL) FileProviderDomain {
	instance := getFileProviderDomainClass().Alloc()
	rv := objc.Send[FileProviderDomain](instance.ID, objc.Sel("initWithDisplayName:userInfo:volumeURL:"), objc.String(displayName), userInfo, volumeURL)
	rv.Autorelease()
	return rv
}



// Creates a new file provider domain with the specified identifier and display name.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:)
func NewFileProviderDomainWithIdentifierDisplayName(identifier unsafe.Pointer, displayName string) FileProviderDomain {
	instance := getFileProviderDomainClass().Alloc()
	rv := objc.Send[FileProviderDomain](instance.ID, objc.Sel("initWithIdentifier:displayName:"), identifier, objc.String(displayName))
	rv.Autorelease()
	return rv
}



// Returns a newly instantiated domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:pathRelativeToDocumentStorage:)
func NewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage(identifier unsafe.Pointer, displayName string, pathRelativeToDocumentStorage string) FileProviderDomain {
	instance := getFileProviderDomainClass().Alloc()
	rv := objc.Send[FileProviderDomain](instance.ID, objc.Sel("initWithIdentifier:displayName:pathRelativeToDocumentStorage:"), identifier, objc.String(displayName), objc.String(pathRelativeToDocumentStorage))
	rv.Autorelease()
	return rv
}


// A Boolean value indicating that the domain is present, but disconnected from the file extension.
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/isdisconnected
func (f_ FileProviderDomain) IsDisconnected() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDisconnected"))
	return rv
}


// SetIsDisconnected sets the value of the isDisconnected property.
// A Boolean value indicating that the domain is present, but disconnected from the file extension.

//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/isdisconnected
func (f_ FileProviderDomain) SetIsDisconnected(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDisconnected:"), value)
}

// The domain managed by this file provider object.
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderextension/domain
func (f_ FileProviderDomain) Domain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("domain"))
	return rv
}


// SetDomain sets the value of the domain property.
// The domain managed by this file provider object.

//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderextension/domain
func (f_ FileProviderDomain) SetDomain(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDomain:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/isreplicated
func (f_ FileProviderDomain) IsReplicated() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReplicated"))
	return rv
}


// SetIsReplicated sets the value of the isReplicated property.
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/isreplicated
func (f_ FileProviderDomain) SetIsReplicated(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsReplicated:"), value)
}

// A Boolean value that determines whether the domain is visible to users.
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/ishidden
func (f_ FileProviderDomain) IsHidden() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that determines whether the domain is visible to users.

//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomain/ishidden
func (f_ FileProviderDomain) SetIsHidden(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsHidden:"), value)
}

// A unique identifier for the backing store used by the system.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/backingStoreIdentity
func (f_ FileProviderDomain) BackingStoreIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("backingStoreIdentity"))
	return rv
}

// The name of the domain displayed in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/displayName
func (f_ FileProviderDomain) DisplayName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("displayName"))
	return rv
}

// The domain’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/identifier
func (f_ FileProviderDomain) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value indicating that the domain is present, but disconnected from the file extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isDisconnected
func (f_ FileProviderDomain) Disconnected() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("disconnected"))
	return rv
}

// A Boolean value that determines whether the domain is visible to users.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isHidden
func (f_ FileProviderDomain) Hidden() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hidden"))
	return rv
}


// SetHidden sets the value of the hidden property.
// A Boolean value that determines whether the domain is visible to users.

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isHidden
func (f_ FileProviderDomain) SetHidden(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHidden:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isReplicated
func (f_ FileProviderDomain) Replicated() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("replicated"))
	return rv
}

// The path of the domain’s subdirectory relative to the file provider’s shared container.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/pathRelativeToDocumentStorage
func (f_ FileProviderDomain) PathRelativeToDocumentStorage() string {
	rv := objc.Send[string](f_.ID, objc.Sel("pathRelativeToDocumentStorage"))
	return rv
}

// A list of known folders that the domain currently replicates.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/replicatedKnownFolders
func (f_ FileProviderDomain) ReplicatedKnownFolders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("replicatedKnownFolders"))
	return rv
}

// A list of known folders that the domain can replicate.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportedKnownFolders
func (f_ FileProviderDomain) SupportedKnownFolders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("supportedKnownFolders"))
	return rv
}


// SetSupportedKnownFolders sets the value of the supportedKnownFolders property.
// A list of known folders that the domain can replicate.

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportedKnownFolders
func (f_ FileProviderDomain) SetSupportedKnownFolders(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportedKnownFolders:"), value)
}

// A Boolean value that indicates whether the provider supports search.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsStringSearchRequest
func (f_ FileProviderDomain) SupportsStringSearchRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsStringSearchRequest"))
	return rv
}


// SetSupportsStringSearchRequest sets the value of the supportsStringSearchRequest property.
// A Boolean value that indicates whether the provider supports search.

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsStringSearchRequest
func (f_ FileProviderDomain) SetSupportsStringSearchRequest(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsStringSearchRequest:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsSyncingTrash
func (f_ FileProviderDomain) SupportsSyncingTrash() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsSyncingTrash"))
	return rv
}


// SetSupportsSyncingTrash sets the value of the supportsSyncingTrash property.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsSyncingTrash
func (f_ FileProviderDomain) SetSupportsSyncingTrash(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsSyncingTrash:"), value)
}

// A mode that gives the File Provider extension more control over the system’s behavior during testing.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/testingModes-swift.property
func (f_ FileProviderDomain) TestingModes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("testingModes"))
	return rv
}


// SetTestingModes sets the value of the testingModes property.
// A mode that gives the File Provider extension more control over the system’s behavior during testing.

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/testingModes-swift.property
func (f_ FileProviderDomain) SetTestingModes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTestingModes:"), value)
}

// A Boolean value that indicates whether the user has enabled or disabled the domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/userEnabled
func (f_ FileProviderDomain) UserEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("userEnabled"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/userInfo
func (f_ FileProviderDomain) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/userInfo
func (f_ FileProviderDomain) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUserInfo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/volumeUUID
func (f_ FileProviderDomain) VolumeUUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("volumeUUID"))
	return rv
}


