// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LARightStore */


/* debug [class_header]: Header for LARightStore */
// The class instance for the [RightStore] class.
var (
	RightStoreClass     _RightStoreClass
	RightStoreClassOnce sync.Once
)

func getRightStoreClass() _RightStoreClass {
	RightStoreClassOnce.Do(func() {
		RightStoreClass = _RightStoreClass{objc.GetClass("LARightStore")}
	})
	return RightStoreClass
}

type _RightStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RightStore */
// An interface definition for the [RightStore] class.
type IRightStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RightStore */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RightStore */
	// methods:
	RemoveAllRightsWithCompletion(handler unsafe.Pointer)
	RemoveRightCompletion(right ILAPersistedRight, handler unsafe.Pointer)
	RemoveRightForIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	RightForIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	SaveRightIdentifierCompletion(right ILARight, identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	SaveRightIdentifierSecretCompletion(right ILARight, identifier objc.IObject /* cross-framework: NSString */, secret objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RightStore */
// Alloc allocates a new instance without initialization.
func (rc _RightStoreClass) Alloc() RightStore {
	rv := objc.Send[RightStore](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RightStoreClass) New() RightStore {
	rv := objc.Send[RightStore](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RightStore) Init() RightStore {
	rv := objc.Send[RightStore](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RightStore) Autorelease() RightStore {
	rv := objc.Send[RightStore](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRightStore creates a new RightStore instance.
func NewRightStore() RightStore {
	return getRightStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RightStore */
// A container for data protected by a right.
//
// Use an along with an to make secrets accessible only after certain conditions, including authentication, are met. Storing secrets this way lets you tie the availability of sensitive resources to the authorization status of the user. The following stores a named access token behind the default authorization requirements: The system stores your secret in the keychain and protects it with a unique key in the Secure Enclave. The system associates the key with your right and with an access control list that ensures that the data is only accessible after your access requirements are met. You can retrieve stored secrets later using the right’s identifier:


// A container for data protected by a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore
type RightStore struct {
	objectivec.Object
}

// RightStoreFrom constructs a [RightStore] from an unsafe.Pointer.
//
// A container for data protected by a right.
func RightStoreFrom(ptr unsafe.Pointer) RightStore {
	return RightStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RightStore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RightStore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RightStore */

// A shared object that stores rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/shared
func (rc _RightStoreClass) SharedStore() RightStore {
	rv := objc.Send[RightStore](objc.ID(rc.class), objc.Sel("sharedStore"))
	return rv
}/* debug [class_properties_class/property]: sharedStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RightStore */

// Removes all rights associated with this client from the right store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeAllRights(completion:)
func (r_ RightStore) RemoveAllRightsWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeAllRightsWithCompletion:"), handler)
}/* debug [instance_methods/method]: RemoveAllRightsWithCompletion */


// Removes a right from the right store given an instance of that right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeRight(_:completion:)
func (r_ RightStore) RemoveRightCompletion(right ILAPersistedRight, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeRight:completion:"), right, handler)
}/* debug [instance_methods/method]: RemoveRightCompletion */


// Removes a right from the right store given its unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeRight(forIdentifier:completion:)
func (r_ RightStore) RemoveRightForIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeRightForIdentifier:completion:"), identifier, handler)
}/* debug [instance_methods/method]: RemoveRightForIdentifierCompletion */


// Fetches a previously stored right from the shared right store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/right(forIdentifier:completion:)
func (r_ RightStore) RightForIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightForIdentifier:completion:"), identifier, handler)
}/* debug [instance_methods/method]: RightForIdentifierCompletion */


// Saves a right to a persistent right store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/saveRight(_:identifier:completion:)
func (r_ RightStore) SaveRightIdentifierCompletion(right ILARight, identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("saveRight:identifier:completion:"), right, identifier, handler)
}/* debug [instance_methods/method]: SaveRightIdentifierCompletion */


// Saves a right to a persistent store along with secret data you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/saveRight(_:identifier:secret:completion:)
func (r_ RightStore) SaveRightIdentifierSecretCompletion(right ILARight, identifier objc.IObject /* cross-framework: NSString */, secret objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("saveRight:identifier:secret:completion:"), right, identifier, secret, handler)
}/* debug [instance_methods/method]: SaveRightIdentifierSecretCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RightStore */

// A shared object that stores rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/shared
func (r_ RightStore) SharedStore() ILARightStore {
	rv := objc.Send[RightStore](r_.ID, objc.Sel("sharedStore"))
	return rv
}/* debug [instance_properties/getter]: sharedStore */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LARightStore */



