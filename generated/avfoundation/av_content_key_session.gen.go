// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVContentKeySession */


/* debug [class_header]: Header for AVContentKeySession */
// The class instance for the [ContentKeySession] class.
var (
	ContentKeySessionClass     _ContentKeySessionClass
	ContentKeySessionClassOnce sync.Once
)

func getContentKeySessionClass() _ContentKeySessionClass {
	ContentKeySessionClassOnce.Do(func() {
		ContentKeySessionClass = _ContentKeySessionClass{objc.GetClass("AVContentKeySession")}
	})
	return ContentKeySessionClass
}

type _ContentKeySessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentKeySession */
// An interface definition for the [ContentKeySession] class.
type IContentKeySession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentKeySession */
	// properties:
	ContentKeyRecipients() []objc.ID
	ContentProtectionSessionIdentifier() objc.IObject /* cross-framework: NSData */
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
	KeySystem() ContentKeySystem /* typedef */
	StorageURL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentKeySession */
	// methods:
	AddContentKeyRecipient(recipient unsafe.Pointer)
	Expire()
	InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer)
	InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer)
	MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)
	ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objc.IObject, initializationData objc.IObject /* cross-framework: NSData */, options foundation.IDictionary)
	RemoveContentKeyRecipient(recipient unsafe.Pointer)
	RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IAVContentKeyRequest)
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentKeySession */
// Alloc allocates a new instance without initialization.
func (cc _ContentKeySessionClass) Alloc() ContentKeySession {
	rv := objc.Send[ContentKeySession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentKeySessionClass) New() ContentKeySession {
	rv := objc.Send[ContentKeySession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentKeySession) Init() ContentKeySession {
	rv := objc.Send[ContentKeySession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentKeySession) Autorelease() ContentKeySession {
	rv := objc.Send[ContentKeySession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentKeySession creates a new ContentKeySession instance.
func NewContentKeySession() ContentKeySession {
	return getContentKeySessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentKeySession */
// An object that creates and tracks decryption keys for media data.


// An object that creates and tracks decryption keys for media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession
type ContentKeySession struct {
	objectivec.Object
}

// ContentKeySessionFrom constructs a [ContentKeySession] from an unsafe.Pointer.
//
// An object that creates and tracks decryption keys for media data.
func ContentKeySessionFrom(ptr unsafe.Pointer) ContentKeySession {
	return ContentKeySession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentKeySession */

// Creates a content key session to manage a collection of content decryption keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:)
func NewContentKeySessionWithKeySystem(keySystem ContentKeySystem /* typedef */) ContentKeySession {
	rv := objc.Send[ContentKeySession](objc.ID(getContentKeySessionClass().class), objc.Sel("contentKeySessionWithKeySystem:"), keySystem)
	return rv
}/* debug [class_init_methods/constructor]: NewContentKeySessionWithKeySystem */


// Creates a content key session to manage a collection of content decryption keys; points to a directory that stores abnormal session termination reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:storageDirectoryAt:)
func NewContentKeySessionWithKeySystemStorageDirectoryAtURL(keySystem ContentKeySystem /* typedef */, storageURL objc.IObject /* cross-framework: NSURL */) ContentKeySession {
	rv := objc.Send[ContentKeySession](objc.ID(getContentKeySessionClass().class), objc.Sel("contentKeySessionWithKeySystem:storageDirectoryAtURL:"), keySystem, storageURL)
	return rv
}/* debug [class_init_methods/constructor]: NewContentKeySessionWithKeySystemStorageDirectoryAtURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentKeySession */

// Creates a content key session to manage a collection of content decryption keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:)
func (cc _ContentKeySessionClass) ContentKeySessionWithKeySystem(keySystem ContentKeySystem /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeySessionWithKeySystem:"), keySystem)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentKeySessionWithKeySystem) */


// Creates a content key session to manage a collection of content decryption keys; points to a directory that stores abnormal session termination reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:storageDirectoryAt:)
func (cc _ContentKeySessionClass) ContentKeySessionWithKeySystemStorageDirectoryAtURL(keySystem ContentKeySystem /* typedef */, storageURL objc.IObject /* cross-framework: NSURL */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeySessionWithKeySystem:storageDirectoryAtURL:"), keySystem, storageURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentKeySessionWithKeySystemStorageDirectoryAtURL) */


// Returns the expired session reports for content key sessions created with the specified app identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/pendingExpiredSessionReports(withAppIdentifier:storageDirectoryAt:)
func (cc _ContentKeySessionClass) PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(appIdentifier objc.IObject /* cross-framework: NSData */, storageURL objc.IObject /* cross-framework: NSURL */) []foundation.Data {
	rv := objc.Send[[]foundation.Data](objc.ID(cc.class), objc.Sel("pendingExpiredSessionReportsWithAppIdentifier:storageDirectoryAtURL:"), appIdentifier, storageURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL) */


// Removes expired session reports from storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/removePendingExpiredSessionReports(_:withAppIdentifier:storageDirectoryAt:)
func (cc _ContentKeySessionClass) RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(expiredSessionReports []foundation.Data, appIdentifier objc.IObject /* cross-framework: NSData */, storageURL objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("removePendingExpiredSessionReports:withAppIdentifier:storageDirectoryAtURL:"), expiredSessionReports, appIdentifier, storageURL)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentKeySession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentKeySession */

// Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/addContentKeyRecipient(_:)
func (c_ ContentKeySession) AddContentKeyRecipient(recipient unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addContentKeyRecipient:"), recipient)
}/* debug [instance_methods/method]: AddContentKeyRecipient */


// Tells the delegate that the session expired as the result of normal, intentional processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/expire()
func (c_ ContentKeySession) Expire() {
	objc.Send[objc.ID](c_.ID, objc.Sel("expire"))
}/* debug [instance_methods/method]: Expire */


// Invalidates all of an app’s persistable content keys and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/invalidateAllPersistableContentKeys(forApp:options:completionHandler:)
func (c_ ContentKeySession) InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateAllPersistableContentKeysForApp:options:completionHandler:"), appIdentifier, options, handler)
}/* debug [instance_methods/method]: InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler */


// Invalidates the persistable content key and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/invalidatePersistableContentKey(_:options:completionHandler:)
func (c_ ContentKeySession) InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidatePersistableContentKey:options:completionHandler:"), persistableContentKeyData, options, handler)
}/* debug [instance_methods/method]: InvalidatePersistableContentKeyOptionsCompletionHandler */


// Creates a secure server playback context that the client sends to the key server to get an expiration date for the given persistable content key data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/makeSecureTokenForExpirationDate(ofPersistableContentKey:completionHandler:)
func (c_ ContentKeySession) MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("makeSecureTokenForExpirationDateOfPersistableContentKey:completionHandler:"), persistableContentKeyData, handler)
}/* debug [instance_methods/method]: MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler */


// Tells the delegate to start loading the content decryption key with the specified identifier and initialization data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/processContentKeyRequest(withIdentifier:initializationData:options:)
func (c_ ContentKeySession) ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objc.IObject, initializationData objc.IObject /* cross-framework: NSData */, options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("processContentKeyRequestWithIdentifier:initializationData:options:"), identifier, initializationData, options)
}/* debug [instance_methods/method]: ProcessContentKeyRequestWithIdentifierInitializationDataOptions */


// Tells the delegate to remove the specified recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/removeContentKeyRecipient(_:)
func (c_ ContentKeySession) RemoveContentKeyRecipient(recipient unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeContentKeyRecipient:"), recipient)
}/* debug [instance_methods/method]: RemoveContentKeyRecipient */


// Tells the delegate that previously provided response data for a content key request is about to expire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/renewExpiringResponseData(for:)
func (c_ ContentKeySession) RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IAVContentKeyRequest) {
	objc.Send[objc.ID](c_.ID, objc.Sel("renewExpiringResponseDataForContentKeyRequest:"), contentKeyRequest)
}/* debug [instance_methods/method]: RenewExpiringResponseDataForContentKeyRequest */


// Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/setDelegate(_:queue:)
func (c_ ContentKeySession) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentKeySession */

// An array of content key recipients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/contentKeyRecipients
func (c_ ContentKeySession) ContentKeyRecipients() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("contentKeyRecipients"))
	return rv
}/* debug [instance_properties/getter]: contentKeyRecipients */


// The identifier for the current content protection session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/contentProtectionSessionIdentifier
func (c_ ContentKeySession) ContentProtectionSessionIdentifier() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("contentProtectionSessionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: contentProtectionSessionIdentifier */


// The content key session’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/delegate
func (c_ ContentKeySession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue the session uses to invoke delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/delegateQueue
func (c_ ContentKeySession) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */


// The type of key system used to retrieve keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/keySystem
func (c_ ContentKeySession) KeySystem() ContentKeySystem /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keySystem"))
	return rv
}/* debug [instance_properties/getter]: keySystem */


// A URL that points to a writable storage directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/storageURL
func (c_ ContentKeySession) StorageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("storageURL"))
	return rv
}/* debug [instance_properties/getter]: storageURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContentKeySession */


