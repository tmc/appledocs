// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An immutable object representing an instant message address for the contact.
//
// Use the methods and properties of to identify instant messaging addresses. Some instant message services, such as Facebook and Skype are predefined in this class. You can also specify your own instant message service using the method. objects are thread-safe, and you may access their properties from any thread of your app.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/init(username:service:)
func NewCNInstantMessageAddressWithUsernameService(username string, service string) CNInstantMessageAddress {
	instance := getCNInstantMessageAddressClass().Alloc()
	rv := objc.Send[CNInstantMessageAddress](instance.ID, objc.Sel("initWithUsername:service:"), objc.String(username), objc.String(service))
	rv.Autorelease()
	return rv
}


// Returns a string containing the localized property name.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forKey:)
func (cc _CNInstantMessageAddressClass) LocalizedStringForKey(key string) string {
	rv := objc.Send[string](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), objc.String(key))
	return rv
}

// Returns a string containing the localized name of the specified service.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNInstantMessageAddress/localizedString(forService:)
func (cc _CNInstantMessageAddressClass) LocalizedStringForService(service string) string {
	rv := objc.Send[string](objc.ID(cc.class), objc.Sel("localizedStringForService:"), objc.String(service))
	return rv
}


