// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEEmailAddress] class.
var (
	MEEmailAddressClass     _MEEmailAddressClass
	MEEmailAddressClassOnce sync.Once
)

func getMEEmailAddressClass() _MEEmailAddressClass {
	MEEmailAddressClassOnce.Do(func() {
		MEEmailAddressClass = _MEEmailAddressClass{objc.GetClass("MEEmailAddress")}
	})
	return MEEmailAddressClass
}

type _MEEmailAddressClass struct {
	class objc.Class
}

// An interface definition for the [MEEmailAddress] class.
type IMEEmailAddress interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress
type MEEmailAddress struct {
	objectivec.Object
}

// MEEmailAddressFrom constructs a [MEEmailAddress] from an unsafe.Pointer.
func MEEmailAddressFrom(ptr unsafe.Pointer) MEEmailAddress {
	return MEEmailAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEEmailAddressClass) Alloc() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEEmailAddressClass) New() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEEmailAddress) Init() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEEmailAddress) Autorelease() MEEmailAddress {
	rv := objc.Send[MEEmailAddress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEmailAddress creates a new MEEmailAddress instance.
func NewMEEmailAddress() MEEmailAddress {
	return getMEEmailAddressClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/init(rawString:)
func NewMEEmailAddressWithRawString(rawString string) MEEmailAddress {
	instance := getMEEmailAddressClass().Alloc()
	rv := objc.Send[MEEmailAddress](instance.ID, objc.Sel("initWithRawString:"), objc.String(rawString))
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/meemailaddress/rawstring
func (m_ MEEmailAddress) RawString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("rawString"))
	return rv
}


// SetRawString sets the value of the rawString property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/meemailaddress/rawstring
func (m_ MEEmailAddress) SetRawString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRawString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/addressString
func (m_ MEEmailAddress) AddressString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("addressString"))
	return rv
}


