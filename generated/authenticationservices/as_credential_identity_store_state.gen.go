// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASCredentialIdentityStoreState */


/* debug [class_header]: Header for ASCredentialIdentityStoreState */
// The class instance for the [CredentialIdentityStoreState] class.
var (
	CredentialIdentityStoreStateClass     _CredentialIdentityStoreStateClass
	CredentialIdentityStoreStateClassOnce sync.Once
)

func getCredentialIdentityStoreStateClass() _CredentialIdentityStoreStateClass {
	CredentialIdentityStoreStateClassOnce.Do(func() {
		CredentialIdentityStoreStateClass = _CredentialIdentityStoreStateClass{objc.GetClass("ASCredentialIdentityStoreState")}
	})
	return CredentialIdentityStoreStateClass
}

type _CredentialIdentityStoreStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CredentialIdentityStoreState */
// An interface definition for the [CredentialIdentityStoreState] class.
type ICredentialIdentityStoreState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CredentialIdentityStoreState */
	// properties:
	Enabled() bool
	SupportsIncrementalUpdates() bool
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CredentialIdentityStoreState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CredentialIdentityStoreState */
// Alloc allocates a new instance without initialization.
func (cc _CredentialIdentityStoreStateClass) Alloc() CredentialIdentityStoreState {
	rv := objc.Send[CredentialIdentityStoreState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CredentialIdentityStoreStateClass) New() CredentialIdentityStoreState {
	rv := objc.Send[CredentialIdentityStoreState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CredentialIdentityStoreState) Init() CredentialIdentityStoreState {
	rv := objc.Send[CredentialIdentityStoreState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CredentialIdentityStoreState) Autorelease() CredentialIdentityStoreState {
	rv := objc.Send[CredentialIdentityStoreState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCredentialIdentityStoreState creates a new CredentialIdentityStoreState instance.
func NewCredentialIdentityStoreState() CredentialIdentityStoreState {
	return getCredentialIdentityStoreStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CredentialIdentityStoreState */
// A representation of the state of a credential identity store.


// A representation of the state of a credential identity store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreState
type CredentialIdentityStoreState struct {
	objectivec.Object
}

// CredentialIdentityStoreStateFrom constructs a [CredentialIdentityStoreState] from an unsafe.Pointer.
//
// A representation of the state of a credential identity store.
func CredentialIdentityStoreStateFrom(ptr unsafe.Pointer) CredentialIdentityStoreState {
	return CredentialIdentityStoreState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CredentialIdentityStoreState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CredentialIdentityStoreState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CredentialIdentityStoreState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CredentialIdentityStoreState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CredentialIdentityStoreState */

// A Boolean value indicating whether the credential identity store is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreState/isEnabled
func (c_ CredentialIdentityStoreState) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value indicating whether the credential identity store supports incremental updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreState/supportsIncrementalUpdates
func (c_ CredentialIdentityStoreState) SupportsIncrementalUpdates() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsIncrementalUpdates"))
	return rv
}/* debug [instance_properties/getter]: supportsIncrementalUpdates */


// A Boolean value indicating whether the credential identity store is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/ascredentialidentitystorestate/isenabled
func (c_ CredentialIdentityStoreState) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value indicating whether the credential identity store is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/ascredentialidentitystorestate/isenabled
func (c_ CredentialIdentityStoreState) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASCredentialIdentityStoreState */



