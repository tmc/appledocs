// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredentialStorage] class.
var (
	URLCredentialStorageClass     _URLCredentialStorageClass
	URLCredentialStorageClassOnce sync.Once
)

func getURLCredentialStorageClass() _URLCredentialStorageClass {
	URLCredentialStorageClassOnce.Do(func() {
		URLCredentialStorageClass = _URLCredentialStorageClass{objc.GetClass("NSURLCredentialStorage")}
	})
	return URLCredentialStorageClass
}

type _URLCredentialStorageClass struct {
	class objc.Class
}

// An interface definition for the [URLCredentialStorage] class.
type IURLCredentialStorage interface {
	objectivec.IObject
	// properties:
	AllCredentials() IDictionary /* already interface */
	// methods:
	CredentialsForProtectionSpace(space IURLProtectionSpace) IDictionary /* already interface */
	DefaultCredentialForProtectionSpace(space IURLProtectionSpace) IURLCredential
	GetCredentialsForProtectionSpaceTaskCompletionHandler(protectionSpace IURLProtectionSpace, task IURLSessionTask, completionHandler IDictionary /* already interface */)
	GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space IURLProtectionSpace, task IURLSessionTask, completionHandler unsafe.Pointer)
	RemoveCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	RemoveCredentialForProtectionSpaceOptions(credential IURLCredential, space IURLProtectionSpace, options IDictionary /* already interface */)
	RemoveCredentialForProtectionSpaceOptionsTask(credential IURLCredential, protectionSpace IURLProtectionSpace, options IDictionary /* already interface */, task IURLSessionTask)
	SetCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	SetCredentialForProtectionSpaceTask(credential IURLCredential, protectionSpace IURLProtectionSpace, task IURLSessionTask)
	SetDefaultCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace)
	SetDefaultCredentialForProtectionSpaceTask(credential IURLCredential, protectionSpace IURLProtectionSpace, task IURLSessionTask)
}

// The manager of a shared credentials cache.
//
// The shared cache stores and retrieves instances of . You can store password-based credentials permanently, based on the they were created with. Certificate-based credentials are never stored permanently.


// The manager of a shared credentials cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage
type URLCredentialStorage struct {
	objectivec.Object
}

// URLCredentialStorageFrom constructs a [URLCredentialStorage] from an unsafe.Pointer.
//
// The manager of a shared credentials cache.
func URLCredentialStorageFrom(ptr unsafe.Pointer) URLCredentialStorage {
	return URLCredentialStorage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLCredentialStorageClass) Alloc() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCredentialStorageClass) New() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCredentialStorage) Init() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCredentialStorage) Autorelease() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredentialStorage creates a new URLCredentialStorage instance.
func NewURLCredentialStorage() URLCredentialStorage {
	return getURLCredentialStorageClass().New()
}



// The shared URL credential storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/shared
func (uc _URLCredentialStorageClass) SharedCredentialStorage() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("sharedCredentialStorage"))
	return rv
}

// Returns a dictionary containing the credentials for the specified protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/credentials(for:)
func (u_ URLCredentialStorage) CredentialsForProtectionSpace(space IURLProtectionSpace) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("credentialsForProtectionSpace:"), space)
	return rv
}


// Returns the default credential for the specified protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/defaultCredential(for:)
func (u_ URLCredentialStorage) DefaultCredentialForProtectionSpace(space IURLProtectionSpace) IURLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("defaultCredentialForProtectionSpace:"), space)
	return rv
}


// Gets a dictionary containing the credentials for the specified protection space, on behalf of the given task, and passes the dictionary to the provided completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/getCredentials(for:task:completionHandler:)
func (u_ URLCredentialStorage) GetCredentialsForProtectionSpaceTaskCompletionHandler(protectionSpace IURLProtectionSpace, task IURLSessionTask, completionHandler IDictionary /* already interface */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getCredentialsForProtectionSpace:task:completionHandler:"), protectionSpace, task, completionHandler)
}


// Gets the default credential for the specified protection space, which is being accessed by the given task, and passes it to the provided completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/getDefaultCredential(for:task:completionHandler:)
func (u_ URLCredentialStorage) GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space IURLProtectionSpace, task IURLSessionTask, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getDefaultCredentialForProtectionSpace:task:completionHandler:"), space, task, completionHandler)
}


// Removes the specified credential from the credential storage for the specified protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:)
func (u_ URLCredentialStorage) RemoveCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCredential:forProtectionSpace:"), credential, space)
}


// Removes the specified credential from the credential storage for the specified protection space using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:options:)
func (u_ URLCredentialStorage) RemoveCredentialForProtectionSpaceOptions(credential IURLCredential, space IURLProtectionSpace, options IDictionary /* already interface */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCredential:forProtectionSpace:options:"), credential, space, options)
}


// Removes the specified credential from the credential storage for the specified protection space, on behalf of the given task and using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:options:task:)
func (u_ URLCredentialStorage) RemoveCredentialForProtectionSpaceOptionsTask(credential IURLCredential, protectionSpace IURLProtectionSpace, options IDictionary /* already interface */, task IURLSessionTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCredential:forProtectionSpace:options:task:"), credential, protectionSpace, options, task)
}


// Adds a credential to the credential storage for the specified protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/set(_:for:)
func (u_ URLCredentialStorage) SetCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCredential:forProtectionSpace:"), credential, space)
}


// Adds a credential to the credential storage for the specified protection space, on behalf of the specified task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/set(_:for:task:)
func (u_ URLCredentialStorage) SetCredentialForProtectionSpaceTask(credential IURLCredential, protectionSpace IURLProtectionSpace, task IURLSessionTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCredential:forProtectionSpace:task:"), credential, protectionSpace, task)
}


// Sets the default credential for a specified protection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/setDefaultCredential(_:for:)
func (u_ URLCredentialStorage) SetDefaultCredentialForProtectionSpace(credential IURLCredential, space IURLProtectionSpace) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDefaultCredential:forProtectionSpace:"), credential, space)
}


// Sets the default credential for a given protection space, which is being accessed by the given task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/setDefaultCredential(_:for:task:)
func (u_ URLCredentialStorage) SetDefaultCredentialForProtectionSpaceTask(credential IURLCredential, protectionSpace IURLProtectionSpace, task IURLSessionTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDefaultCredential:forProtectionSpace:task:"), credential, protectionSpace, task)
}


// The credentials for all available protection spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/allCredentials
func (u_ URLCredentialStorage) AllCredentials() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("allCredentials"))
	return rv
}


// The shared URL credential storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/shared
func (u_ URLCredentialStorage) SharedCredentialStorage() IURLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("sharedCredentialStorage"))
	return rv
}



