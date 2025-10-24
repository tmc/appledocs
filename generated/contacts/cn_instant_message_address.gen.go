// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNInstantMessageAddress */


/* debug [class_header]: Header for CNInstantMessageAddress */
// The class instance for the [CNInstantMessageAddress] class.
var (
	CNInstantMessageAddressClass     _CNInstantMessageAddressClass
	CNInstantMessageAddressClassOnce sync.Once
)

func getCNInstantMessageAddressClass() _CNInstantMessageAddressClass {
	CNInstantMessageAddressClassOnce.Do(func() {
		CNInstantMessageAddressClass = _CNInstantMessageAddressClass{objc.GetClass("CNInstantMessageAddress")}
	})
	return CNInstantMessageAddressClass
}

type _CNInstantMessageAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNInstantMessageAddress */
// An interface definition for the [CNInstantMessageAddress] class.
type ICNInstantMessageAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNInstantMessageAddress */
	// properties:
	Service() objc.IObject /* cross-framework: NSString */
	Username() objc.IObject /* cross-framework: NSString */
	CNInstantMessageAddressServiceKey() objc.IObject /* cross-framework: NSString */
	CNInstantMessageAddressUsernameKey() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceAIM() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceFacebook() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceGaduGadu() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceGoogleTalk() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceICQ() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceJabber() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceMSN() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceQQ() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceSkype() objc.IObject /* cross-framework: NSString */
	CNInstantMessageServiceYahoo() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNInstantMessageAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNInstantMessageAddress */
// Alloc allocates a new instance without initialization.
func (cc _CNInstantMessageAddressClass) Alloc() CNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNInstantMessageAddressClass) New() CNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNInstantMessageAddress) Init() CNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNInstantMessageAddress) Autorelease() CNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNInstantMessageAddress creates a new CNInstantMessageAddress instance.
func NewCNInstantMessageAddress() CNInstantMessageAddress {
	return getCNInstantMessageAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNInstantMessageAddress */
// An immutable object representing an instant message address for the contact.
//
// Use the methods and properties of to identify instant messaging addresses. Some instant message services, such as Facebook and Skype are predefined in this class. You can also specify your own instant message service using the method. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object representing an instant message address for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress
type CNInstantMessageAddress struct {
	objectivec.Object
}

// CNInstantMessageAddressFrom constructs a [CNInstantMessageAddress] from an unsafe.Pointer.
//
// An immutable object representing an instant message address for the contact.
func CNInstantMessageAddressFrom(ptr unsafe.Pointer) CNInstantMessageAddress {
	return CNInstantMessageAddress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNInstantMessageAddress */

// Returns a object initialized with the specified user name and service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/init(username:service:)
func NewCNInstantMessageAddressWithUsernameService(username objc.IObject /* cross-framework: NSString */, service objc.IObject /* cross-framework: NSString */) CNInstantMessageAddress {
	instance := getCNInstantMessageAddressClass().Alloc()
	rv := objc.Send[CNInstantMessageAddress](instance.ID, objc.Sel("initWithUsername:service:"), username, service)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNInstantMessageAddressWithUsernameService */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNInstantMessageAddress */

// Returns a string containing the localized property name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forKey:)
func (cc _CNInstantMessageAddressClass) LocalizedStringForKey(key objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForKey) */


// Returns a string containing the localized name of the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forService:)
func (cc _CNInstantMessageAddressClass) LocalizedStringForService(service objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForService:"), service)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForService) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNInstantMessageAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNInstantMessageAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNInstantMessageAddress */

// The name of the instant message address service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/service
func (c_ CNInstantMessageAddress) Service() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("service"))
	return rv
}/* debug [instance_properties/getter]: service */


// The user name for instant message service address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/username
func (c_ CNInstantMessageAddress) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("username"))
	return rv
}/* debug [instance_properties/getter]: username */


// Instant message address service key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressservicekey
func (c_ CNInstantMessageAddress) CNInstantMessageAddressServiceKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageAddressServiceKey"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageAddressServiceKey */


// Instant message address username key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressusernamekey
func (c_ CNInstantMessageAddress) CNInstantMessageAddressUsernameKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageAddressUsernameKey"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageAddressUsernameKey */


// Instant message service for AIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceaim
func (c_ CNInstantMessageAddress) CNInstantMessageServiceAIM() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceAIM"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceAIM */


// Instant message service for Facebook.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicefacebook
func (c_ CNInstantMessageAddress) CNInstantMessageServiceFacebook() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceFacebook"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceFacebook */


// Instant message service for Gadu Gadu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegadugadu
func (c_ CNInstantMessageAddress) CNInstantMessageServiceGaduGadu() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceGaduGadu"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceGaduGadu */


// Instant message service for Google Talk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegoogletalk
func (c_ CNInstantMessageAddress) CNInstantMessageServiceGoogleTalk() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceGoogleTalk"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceGoogleTalk */


// Instant message service for ICQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceicq
func (c_ CNInstantMessageAddress) CNInstantMessageServiceICQ() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceICQ"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceICQ */


// Instant message service for Jabber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicejabber
func (c_ CNInstantMessageAddress) CNInstantMessageServiceJabber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceJabber"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceJabber */


// Instant message service for MSN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicemsn
func (c_ CNInstantMessageAddress) CNInstantMessageServiceMSN() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceMSN"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceMSN */


// Instant message service for QQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceqq
func (c_ CNInstantMessageAddress) CNInstantMessageServiceQQ() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceQQ"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceQQ */


// Instant message service for Skype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceskype
func (c_ CNInstantMessageAddress) CNInstantMessageServiceSkype() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceSkype"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceSkype */


// Instant message service for Yahoo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceyahoo
func (c_ CNInstantMessageAddress) CNInstantMessageServiceYahoo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNInstantMessageServiceYahoo"))
	return rv
}/* debug [instance_properties/getter]: CNInstantMessageServiceYahoo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNInstantMessageAddress */


