// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNInstantMessageAddress] class.
type ICNInstantMessageAddress interface {
	objectivec.IObject
	Service() string
	SetService(value string)
	Username() string
	SetUsername(value string)
	CNInstantMessageAddressServiceKey() string
	CNInstantMessageAddressUsernameKey() string
	CNInstantMessageServiceAIM() string
	CNInstantMessageServiceFacebook() string
	CNInstantMessageServiceGaduGadu() string
	CNInstantMessageServiceGoogleTalk() string
	CNInstantMessageServiceICQ() string
	CNInstantMessageServiceJabber() string
	CNInstantMessageServiceMSN() string
	CNInstantMessageServiceQQ() string
	CNInstantMessageServiceSkype() string
	CNInstantMessageServiceYahoo() string
}

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

// Alloc allocates a new instance without initialization.
func (cc _CNInstantMessageAddressClass) Alloc() CNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns a object initialized with the specified user name and service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/init(username:service:)

func NewCNInstantMessageAddressWithUsernameService(username string, service string) CNInstantMessageAddress {
	instance := getCNInstantMessageAddressClass().Alloc()
	rv := objc.Send[CNInstantMessageAddress](instance.ID, objc.Sel("initWithUsername:service:"), objc.String(username), objc.String(service))
	rv.Autorelease()
	return rv
}



// Returns a string containing the localized property name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forKey:)

func (cc _CNInstantMessageAddressClass) LocalizedStringForKey(key string) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), objc.String(key))
	return rv
}


// Returns a string containing the localized name of the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forService:)

func (cc _CNInstantMessageAddressClass) LocalizedStringForService(service string) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForService:"), objc.String(service))
	return rv
}


// The name of the instant message address service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/service

func (c_ CNInstantMessageAddress) Service() string {
	rv := objc.Send[string](c_.ID, objc.Sel("service"))
	return rv
}


// The name of the instant message address service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/service

func (c_ CNInstantMessageAddress) SetService(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setService:"), objc.String(value))
}


// The user name for instant message service address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/username

func (c_ CNInstantMessageAddress) Username() string {
	rv := objc.Send[string](c_.ID, objc.Sel("username"))
	return rv
}


// The user name for instant message service address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/username

func (c_ CNInstantMessageAddress) SetUsername(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsername:"), objc.String(value))
}


// Instant message address service key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressservicekey

func (c_ CNInstantMessageAddress) CNInstantMessageAddressServiceKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageAddressServiceKey"))
	return rv
}


// Instant message address username key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressusernamekey

func (c_ CNInstantMessageAddress) CNInstantMessageAddressUsernameKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageAddressUsernameKey"))
	return rv
}


// Instant message service for AIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceaim

func (c_ CNInstantMessageAddress) CNInstantMessageServiceAIM() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceAIM"))
	return rv
}


// Instant message service for Facebook.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicefacebook

func (c_ CNInstantMessageAddress) CNInstantMessageServiceFacebook() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceFacebook"))
	return rv
}


// Instant message service for Gadu Gadu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegadugadu

func (c_ CNInstantMessageAddress) CNInstantMessageServiceGaduGadu() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceGaduGadu"))
	return rv
}


// Instant message service for Google Talk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegoogletalk

func (c_ CNInstantMessageAddress) CNInstantMessageServiceGoogleTalk() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceGoogleTalk"))
	return rv
}


// Instant message service for ICQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceicq

func (c_ CNInstantMessageAddress) CNInstantMessageServiceICQ() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceICQ"))
	return rv
}


// Instant message service for Jabber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicejabber

func (c_ CNInstantMessageAddress) CNInstantMessageServiceJabber() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceJabber"))
	return rv
}


// Instant message service for MSN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicemsn

func (c_ CNInstantMessageAddress) CNInstantMessageServiceMSN() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceMSN"))
	return rv
}


// Instant message service for QQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceqq

func (c_ CNInstantMessageAddress) CNInstantMessageServiceQQ() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceQQ"))
	return rv
}


// Instant message service for Skype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceskype

func (c_ CNInstantMessageAddress) CNInstantMessageServiceSkype() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceSkype"))
	return rv
}


// Instant message service for Yahoo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceyahoo

func (c_ CNInstantMessageAddress) CNInstantMessageServiceYahoo() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceYahoo"))
	return rv
}


