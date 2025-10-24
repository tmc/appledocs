// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASCredentialServiceIdentifier */


/* debug [class_header]: Header for ASCredentialServiceIdentifier */
// The class instance for the [CredentialServiceIdentifier] class.
var (
	CredentialServiceIdentifierClass     _CredentialServiceIdentifierClass
	CredentialServiceIdentifierClassOnce sync.Once
)

func getCredentialServiceIdentifierClass() _CredentialServiceIdentifierClass {
	CredentialServiceIdentifierClassOnce.Do(func() {
		CredentialServiceIdentifierClass = _CredentialServiceIdentifierClass{objc.GetClass("ASCredentialServiceIdentifier")}
	})
	return CredentialServiceIdentifierClass
}

type _CredentialServiceIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CredentialServiceIdentifier */
// An interface definition for the [CredentialServiceIdentifier] class.
type ICredentialServiceIdentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CredentialServiceIdentifier */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Type() CredentialServiceIdentifierType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CredentialServiceIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CredentialServiceIdentifier */
// Alloc allocates a new instance without initialization.
func (cc _CredentialServiceIdentifierClass) Alloc() CredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CredentialServiceIdentifierClass) New() CredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CredentialServiceIdentifier) Init() CredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CredentialServiceIdentifier) Autorelease() CredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCredentialServiceIdentifier creates a new CredentialServiceIdentifier instance.
func NewCredentialServiceIdentifier() CredentialServiceIdentifier {
	return getCredentialServiceIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CredentialServiceIdentifier */
// An identifier representing a particular service for which the user needs a credential, like a web site.


// An identifier representing a particular service for which the user needs a credential, like a web site.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier
type CredentialServiceIdentifier struct {
	objectivec.Object
}

// CredentialServiceIdentifierFrom constructs a [CredentialServiceIdentifier] from an unsafe.Pointer.
//
// An identifier representing a particular service for which the user needs a credential, like a web site.
func CredentialServiceIdentifierFrom(ptr unsafe.Pointer) CredentialServiceIdentifier {
	return CredentialServiceIdentifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CredentialServiceIdentifier */

// Initializes a credential service identifier instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/init(identifier:type:)
func NewCredentialServiceIdentifierWithIdentifierType(identifier objc.IObject /* cross-framework: NSString */, type_ CredentialServiceIdentifierType) CredentialServiceIdentifier {
	instance := getCredentialServiceIdentifierClass().Alloc()
	rv := objc.Send[CredentialServiceIdentifier](instance.ID, objc.Sel("initWithIdentifier:type:"), identifier, type_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCredentialServiceIdentifierWithIdentifierType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CredentialServiceIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CredentialServiceIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CredentialServiceIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CredentialServiceIdentifier */

// A string that names the identified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/identifier
func (c_ CredentialServiceIdentifier) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The kind of services that the identifier represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/type
func (c_ CredentialServiceIdentifier) Type() CredentialServiceIdentifierType {
	rv := objc.Send[CredentialServiceIdentifierType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASCredentialServiceIdentifier */


