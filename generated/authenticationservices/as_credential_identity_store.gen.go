// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASCredentialIdentityStore */


/* debug [class_header]: Header for ASCredentialIdentityStore */
// The class instance for the [CredentialIdentityStore] class.
var (
	CredentialIdentityStoreClass     _CredentialIdentityStoreClass
	CredentialIdentityStoreClassOnce sync.Once
)

func getCredentialIdentityStoreClass() _CredentialIdentityStoreClass {
	CredentialIdentityStoreClassOnce.Do(func() {
		CredentialIdentityStoreClass = _CredentialIdentityStoreClass{objc.GetClass("ASCredentialIdentityStore")}
	})
	return CredentialIdentityStoreClass
}

type _CredentialIdentityStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CredentialIdentityStore */
// An interface definition for the [CredentialIdentityStore] class.
type ICredentialIdentityStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CredentialIdentityStore */
	// properties:
	ASCredentialIdentityStoreErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CredentialIdentityStore */
	// methods:
	GetCredentialIdentitiesForServiceCredentialIdentityTypesCompletionHandler(serviceIdentifier IASCredentialServiceIdentifier, credentialIdentityTypes CredentialIdentityTypes, completionHandler unsafe.Pointer)
	GetCredentialIdentityStoreStateWithCompletion(completion unsafe.Pointer)
	RemoveAllCredentialIdentitiesWithCompletion(completion unsafe.Pointer)
	RemoveCredentialIdentityEntriesCompletion(credentialIdentities []objc.ID, completion unsafe.Pointer)
	ReplaceCredentialIdentityEntriesCompletion(newCredentialIdentities []objc.ID, completion unsafe.Pointer)
	SaveCredentialIdentityEntriesCompletion(credentialIdentities []objc.ID, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CredentialIdentityStore */
// Alloc allocates a new instance without initialization.
func (cc _CredentialIdentityStoreClass) Alloc() CredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CredentialIdentityStoreClass) New() CredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CredentialIdentityStore) Init() CredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CredentialIdentityStore) Autorelease() CredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCredentialIdentityStore creates a new CredentialIdentityStore instance.
func NewCredentialIdentityStore() CredentialIdentityStore {
	return getCredentialIdentityStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CredentialIdentityStore */
// A container that your extension fills to provide credentials through the QuickType bar.
//
// Make credential identities available to users directly as AutoFill suggestions by adding them to the instance of the identity store. You can add identities during configuration in your extension’s override of the method. You can also update the shared store from within your extension’s host app. Be sure to update the shared store whenever your app’s database changes to avoid showing stale identities as AutoFill suggestions. Take advantage of the incremental change methods and to avoid rewriting the entire store every time you need to make a change. When the user disables your extension, the system clears and disables your shared store. So before making updates, check to see that the store’s enabled to avoid unnecessary activity:


// A container that your extension fills to provide credentials through the QuickType bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore
type CredentialIdentityStore struct {
	objectivec.Object
}

// CredentialIdentityStoreFrom constructs a [CredentialIdentityStore] from an unsafe.Pointer.
//
// A container that your extension fills to provide credentials through the QuickType bar.
func CredentialIdentityStoreFrom(ptr unsafe.Pointer) CredentialIdentityStore {
	return CredentialIdentityStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CredentialIdentityStore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CredentialIdentityStore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CredentialIdentityStore */

// The shared credential identity store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/shared
func (cc _CredentialIdentityStoreClass) SharedStore() CredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](objc.ID(cc.class), objc.Sel("sharedStore"))
	return rv
}/* debug [class_properties_class/property]: sharedStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CredentialIdentityStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/getCredentialIdentitiesForService:credentialIdentityTypes:completionHandler:
func (c_ CredentialIdentityStore) GetCredentialIdentitiesForServiceCredentialIdentityTypesCompletionHandler(serviceIdentifier IASCredentialServiceIdentifier, credentialIdentityTypes CredentialIdentityTypes, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getCredentialIdentitiesForService:credentialIdentityTypes:completionHandler:"), serviceIdentifier, credentialIdentityTypes, completionHandler)
}/* debug [instance_methods/method]: GetCredentialIdentitiesForServiceCredentialIdentityTypesCompletionHandler */


// Gets the state of the credential identity store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/getState(_:)
func (c_ CredentialIdentityStore) GetCredentialIdentityStoreStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getCredentialIdentityStoreStateWithCompletion:"), completion)
}/* debug [instance_methods/method]: GetCredentialIdentityStoreStateWithCompletion */


// Removes all existing credential identities from the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/removeAllCredentialIdentities(_:)
func (c_ CredentialIdentityStore) RemoveAllCredentialIdentitiesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllCredentialIdentitiesWithCompletion:"), completion)
}/* debug [instance_methods/method]: RemoveAllCredentialIdentitiesWithCompletion */


// Remove the given credential identities from the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/removeCredentialIdentities(_:completion:)-67lcw
func (c_ CredentialIdentityStore) RemoveCredentialIdentityEntriesCompletion(credentialIdentities []objc.ID, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeCredentialIdentityEntries:completion:"), credentialIdentities, completion)
}/* debug [instance_methods/method]: RemoveCredentialIdentityEntriesCompletion */


// Replaces existing credential identities with new credential identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/replaceCredentialIdentities(_:completion:)
func (c_ CredentialIdentityStore) ReplaceCredentialIdentityEntriesCompletion(newCredentialIdentities []objc.ID, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("replaceCredentialIdentityEntries:completion:"), newCredentialIdentities, completion)
}/* debug [instance_methods/method]: ReplaceCredentialIdentityEntriesCompletion */


// Save the supplied credential identities to the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/saveCredentialIdentities(_:completion:)-1bbx6
func (c_ CredentialIdentityStore) SaveCredentialIdentityEntriesCompletion(credentialIdentities []objc.ID, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("saveCredentialIdentityEntries:completion:"), credentialIdentities, completion)
}/* debug [instance_methods/method]: SaveCredentialIdentityEntriesCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CredentialIdentityStore */

// The shared credential identity store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/shared
func (c_ CredentialIdentityStore) SharedStore() IASCredentialIdentityStore {
	rv := objc.Send[CredentialIdentityStore](c_.ID, objc.Sel("sharedStore"))
	return rv
}/* debug [instance_properties/getter]: sharedStore */


// The domain for a credential identity store error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/ascredentialidentitystoreerrordomain
func (c_ CredentialIdentityStore) ASCredentialIdentityStoreErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ASCredentialIdentityStoreErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: ASCredentialIdentityStoreErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASCredentialIdentityStore */



