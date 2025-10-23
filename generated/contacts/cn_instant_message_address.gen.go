// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Service() string /* primitive/slice/pointer. */
	SetService(value string /* primitive/slice/pointer. */)
	Username() string /* primitive/slice/pointer. */
	SetUsername(value string /* primitive/slice/pointer. */)
	CNInstantMessageAddressServiceKey() string /* primitive/slice/pointer. */
	CNInstantMessageAddressUsernameKey() string /* primitive/slice/pointer. */
	CNInstantMessageServiceAIM() string /* primitive/slice/pointer. */
	CNInstantMessageServiceFacebook() string /* primitive/slice/pointer. */
	CNInstantMessageServiceGaduGadu() string /* primitive/slice/pointer. */
	CNInstantMessageServiceGoogleTalk() string /* primitive/slice/pointer. */
	CNInstantMessageServiceICQ() string /* primitive/slice/pointer. */
	CNInstantMessageServiceJabber() string /* primitive/slice/pointer. */
	CNInstantMessageServiceMSN() string /* primitive/slice/pointer. */
	CNInstantMessageServiceQQ() string /* primitive/slice/pointer. */
	CNInstantMessageServiceSkype() string /* primitive/slice/pointer. */
	CNInstantMessageServiceYahoo() string /* primitive/slice/pointer. */
	// methods:
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



// The name of the instant message address service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/service
func (c_ CNInstantMessageAddress) Service() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("service"))
	return rv
}


// The name of the instant message address service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/service
func (c_ CNInstantMessageAddress) SetService(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setService:"), objc.String(value))
}


// The user name for instant message service address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/username
func (c_ CNInstantMessageAddress) Username() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("username"))
	return rv
}


// The user name for instant message service address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/username
func (c_ CNInstantMessageAddress) SetUsername(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsername:"), objc.String(value))
}


// Instant message address service key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressservicekey
func (c_ CNInstantMessageAddress) CNInstantMessageAddressServiceKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageAddressServiceKey"))
	return rv
}


// Instant message address username key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddressusernamekey
func (c_ CNInstantMessageAddress) CNInstantMessageAddressUsernameKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageAddressUsernameKey"))
	return rv
}


// Instant message service for AIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceaim
func (c_ CNInstantMessageAddress) CNInstantMessageServiceAIM() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceAIM"))
	return rv
}


// Instant message service for Facebook.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicefacebook
func (c_ CNInstantMessageAddress) CNInstantMessageServiceFacebook() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceFacebook"))
	return rv
}


// Instant message service for Gadu Gadu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegadugadu
func (c_ CNInstantMessageAddress) CNInstantMessageServiceGaduGadu() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceGaduGadu"))
	return rv
}


// Instant message service for Google Talk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicegoogletalk
func (c_ CNInstantMessageAddress) CNInstantMessageServiceGoogleTalk() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceGoogleTalk"))
	return rv
}


// Instant message service for ICQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceicq
func (c_ CNInstantMessageAddress) CNInstantMessageServiceICQ() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceICQ"))
	return rv
}


// Instant message service for Jabber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicejabber
func (c_ CNInstantMessageAddress) CNInstantMessageServiceJabber() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceJabber"))
	return rv
}


// Instant message service for MSN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageservicemsn
func (c_ CNInstantMessageAddress) CNInstantMessageServiceMSN() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceMSN"))
	return rv
}


// Instant message service for QQ.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceqq
func (c_ CNInstantMessageAddress) CNInstantMessageServiceQQ() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceQQ"))
	return rv
}


// Instant message service for Skype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceskype
func (c_ CNInstantMessageAddress) CNInstantMessageServiceSkype() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceSkype"))
	return rv
}


// Instant message service for Yahoo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageserviceyahoo
func (c_ CNInstantMessageAddress) CNInstantMessageServiceYahoo() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNInstantMessageServiceYahoo"))
	return rv
}



