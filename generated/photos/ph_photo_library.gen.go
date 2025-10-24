// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPhotoLibrary] class.
var (
	PHPhotoLibraryClass     _PHPhotoLibraryClass
	PHPhotoLibraryClassOnce sync.Once
)

func getPHPhotoLibraryClass() _PHPhotoLibraryClass {
	PHPhotoLibraryClassOnce.Do(func() {
		PHPhotoLibraryClass = _PHPhotoLibraryClass{objc.GetClass("PHPhotoLibrary")}
	})
	return PHPhotoLibraryClass
}

type _PHPhotoLibraryClass struct {
	class objc.Class
}

// An interface definition for the [PHPhotoLibrary] class.
type IPHPhotoLibrary interface {
	objectivec.IObject
	// properties:
	CurrentChangeToken() IPHPersistentChangeToken
	UnavailabilityReason() objc.IObject /* cross-framework: Error */
	PHLocalIdentifierNotFound() objc.IObject /* cross-framework: NSString */
	// methods:
	CloudIdentifierMappingsForLocalIdentifiers(localIdentifiers []string) foundation.IDictionary
	FetchPersistentChangesSinceTokenError(token IPHPersistentChangeToken, error_ unsafe.Pointer) IPHPersistentChangeFetchResult
	LocalIdentifierMappingsForCloudIdentifiers(cloudIdentifiers []IPHCloudIdentifier) foundation.IDictionary
	PerformChangesCompletionHandler(changeBlock unsafe.Pointer, completionHandler unsafe.Pointer)
	PerformChangesAndWaitError(changeBlock unsafe.Pointer, error_ unsafe.Pointer) bool
	RegisterChangeObserver(observer objectivec.IObject)
	RegisterAvailabilityObserver(observer objectivec.IObject)
	UnregisterAvailabilityObserver(observer objectivec.IObject)
	UnregisterChangeObserver(observer objectivec.IObject)
}

// An object that manages access and changes to the user’s photo library.
//
// The object represents the entire set of assets and collections that the Photos app manages, including assets stored on the local device and those stored in iCloud Photos. Use this object for the following tasks: Retrieving or verifying the user’s permission for your app to access Photos content Making changes to assets and collections; for example, editing asset metadata or content, inserting new assets, or rearranging the members of a collection Determining which records change since a previous state of the Photos library Registering for update messages the system sends when the library changes


// An object that manages access and changes to the user’s photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary
type PHPhotoLibrary struct {
	objectivec.Object
}

// PHPhotoLibraryFrom constructs a [PHPhotoLibrary] from an unsafe.Pointer.
//
// An object that manages access and changes to the user’s photo library.
func PHPhotoLibraryFrom(ptr unsafe.Pointer) PHPhotoLibrary {
	return PHPhotoLibrary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPhotoLibraryClass) Alloc() PHPhotoLibrary {
	rv := objc.Send[PHPhotoLibrary](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPhotoLibraryClass) New() PHPhotoLibrary {
	rv := objc.Send[PHPhotoLibrary](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPhotoLibrary) Init() PHPhotoLibrary {
	rv := objc.Send[PHPhotoLibrary](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPhotoLibrary) Autorelease() PHPhotoLibrary {
	rv := objc.Send[PHPhotoLibrary](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPhotoLibrary creates a new PHPhotoLibrary instance.
func NewPHPhotoLibrary() PHPhotoLibrary {
	return getPHPhotoLibraryClass().New()
}



// Returns information about your app’s authorization to access the user’s photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/authorizationStatus()
func (pc _PHPhotoLibraryClass) AuthorizationStatus() PHAuthorizationStatus {
	rv := objc.Send[PHAuthorizationStatus](objc.ID(pc.class), objc.Sel("authorizationStatus"))
	return rv
}


// Returns the app’s authorization to access the user’s photo library for the specified access level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/authorizationStatus(for:)
func (pc _PHPhotoLibraryClass) AuthorizationStatusForAccessLevel(accessLevel PHAccessLevel) PHAuthorizationStatus {
	rv := objc.Send[PHAuthorizationStatus](objc.ID(pc.class), objc.Sel("authorizationStatusForAccessLevel:"), accessLevel)
	return rv
}


// Requests the user’s permission, if needed, to access the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/requestAuthorization(_:)
func (pc _PHPhotoLibraryClass) RequestAuthorization(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("requestAuthorization:"), handler)
}


// Prompts the user to grant the app permission to access the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/requestAuthorization(for:handler:)
func (pc _PHPhotoLibraryClass) RequestAuthorizationForAccessLevelHandler(accessLevel PHAccessLevel, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("requestAuthorizationForAccessLevel:handler:"), accessLevel, handler)
}


// Retrieves the shared photo library object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/shared()
func (pc _PHPhotoLibraryClass) SharedPhotoLibrary() PHPhotoLibrary {
	rv := objc.Send[PHPhotoLibrary](objc.ID(pc.class), objc.Sel("sharedPhotoLibrary"))
	return rv
}


// Retrieves the cloud identifier mappings for the list of local identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/cloudIdentifierMappingsForLocalIdentifiers:
func (p_ PHPhotoLibrary) CloudIdentifierMappingsForLocalIdentifiers(localIdentifiers []string) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("cloudIdentifierMappingsForLocalIdentifiers:"), localIdentifiers)
	return rv
}


// Retrieves the Photos library changes since the token you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/fetchPersistentChanges(since:)
func (p_ PHPhotoLibrary) FetchPersistentChangesSinceTokenError(token IPHPersistentChangeToken, error_ unsafe.Pointer) IPHPersistentChangeFetchResult {
	rv := objc.Send[PHPersistentChangeFetchResult](p_.ID, objc.Sel("fetchPersistentChangesSinceToken:error:"), token, error_)
	return rv
}


// Retrieves the local identifier mappings for the list of cloud identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/localIdentifierMappingsForCloudIdentifiers:
func (p_ PHPhotoLibrary) LocalIdentifierMappingsForCloudIdentifiers(cloudIdentifiers []IPHCloudIdentifier) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("localIdentifierMappingsForCloudIdentifiers:"), cloudIdentifiers)
	return rv
}


// Asynchronously runs a block that requests changes to the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/performChanges(_:completionHandler:)
func (p_ PHPhotoLibrary) PerformChangesCompletionHandler(changeBlock unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performChanges:completionHandler:"), changeBlock, completionHandler)
}


// Synchronously runs a block that requests changes to be performed in the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/performChangesAndWait(_:)
func (p_ PHPhotoLibrary) PerformChangesAndWaitError(changeBlock unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("performChangesAndWait:error:"), changeBlock, error_)
	return rv
}


// Registers an object to receive messages when objects in the photo library change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/register(_:)-6y3b9
func (p_ PHPhotoLibrary) RegisterChangeObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("registerChangeObserver:"), observer)
}


// Registers an object to observe changes to the photo library’s availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/register(_:)-gm0a
func (p_ PHPhotoLibrary) RegisterAvailabilityObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("registerAvailabilityObserver:"), observer)
}


// Unregisters an object from observing changes to the photo library’s availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/unregisterAvailabilityObserver(_:)
func (p_ PHPhotoLibrary) UnregisterAvailabilityObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterAvailabilityObserver:"), observer)
}


// Unregisters an object so that it no longer receives change messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/unregisterChangeObserver(_:)
func (p_ PHPhotoLibrary) UnregisterChangeObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterChangeObserver:"), observer)
}


// The opaque token that represents the current state of the Photos library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/currentChangeToken
func (p_ PHPhotoLibrary) CurrentChangeToken() IPHPersistentChangeToken {
	rv := objc.Send[PHPersistentChangeToken](p_.ID, objc.Sel("currentChangeToken"))
	return rv
}


// An error that describes the reason the photo library isn’t available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/unavailabilityReason
func (p_ PHPhotoLibrary) UnavailabilityReason() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](p_.ID, objc.Sel("unavailabilityReason"))
	return rv
}


// A constant value that indicates that the system can’t resolve a local object from a global identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlocalidentifiernotfound
func (p_ PHPhotoLibrary) PHLocalIdentifierNotFound() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("PHLocalIdentifierNotFound"))
	return rv
}


