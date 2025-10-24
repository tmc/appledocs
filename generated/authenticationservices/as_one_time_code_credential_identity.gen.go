// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASOneTimeCodeCredentialIdentity */


/* debug [class_header]: Header for ASOneTimeCodeCredentialIdentity */
// The class instance for the [OneTimeCodeCredentialIdentity] class.
var (
	OneTimeCodeCredentialIdentityClass     _OneTimeCodeCredentialIdentityClass
	OneTimeCodeCredentialIdentityClassOnce sync.Once
)

func getOneTimeCodeCredentialIdentityClass() _OneTimeCodeCredentialIdentityClass {
	OneTimeCodeCredentialIdentityClassOnce.Do(func() {
		OneTimeCodeCredentialIdentityClass = _OneTimeCodeCredentialIdentityClass{objc.GetClass("ASOneTimeCodeCredentialIdentity")}
	})
	return OneTimeCodeCredentialIdentityClass
}

type _OneTimeCodeCredentialIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OneTimeCodeCredentialIdentity */
// An interface definition for the [OneTimeCodeCredentialIdentity] class.
type IOneTimeCodeCredentialIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OneTimeCodeCredentialIdentity */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OneTimeCodeCredentialIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OneTimeCodeCredentialIdentity */
// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialIdentityClass) Alloc() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OneTimeCodeCredentialIdentityClass) New() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredentialIdentity) Init() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredentialIdentity) Autorelease() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredentialIdentity creates a new OneTimeCodeCredentialIdentity instance.
func NewOneTimeCodeCredentialIdentity() OneTimeCodeCredentialIdentity {
	return getOneTimeCodeCredentialIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OneTimeCodeCredentialIdentity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialIdentity
type OneTimeCodeCredentialIdentity struct {
	objectivec.Object
}

// OneTimeCodeCredentialIdentityFrom constructs a [OneTimeCodeCredentialIdentity] from an unsafe.Pointer.
func OneTimeCodeCredentialIdentityFrom(ptr unsafe.Pointer) OneTimeCodeCredentialIdentity {
	return OneTimeCodeCredentialIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OneTimeCodeCredentialIdentity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialIdentity/init(serviceIdentifier:label:recordIdentifier:)
func NewOneTimeCodeCredentialIdentityWithServiceIdentifierLabelRecordIdentifier(serviceIdentifier IASCredentialServiceIdentifier, label objc.IObject /* cross-framework: NSString */, recordIdentifier objc.IObject /* cross-framework: NSString */) OneTimeCodeCredentialIdentity {
	instance := getOneTimeCodeCredentialIdentityClass().Alloc()
	rv := objc.Send[OneTimeCodeCredentialIdentity](instance.ID, objc.Sel("initWithServiceIdentifier:label:recordIdentifier:"), serviceIdentifier, label, recordIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOneTimeCodeCredentialIdentityWithServiceIdentifierLabelRecordIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OneTimeCodeCredentialIdentity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OneTimeCodeCredentialIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OneTimeCodeCredentialIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OneTimeCodeCredentialIdentity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialIdentity/label
func (o_ OneTimeCodeCredentialIdentity) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASOneTimeCodeCredentialIdentity */


