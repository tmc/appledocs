// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AddressString() objc.IObject /* cross-framework: NSString */
	RawString() objc.IObject /* cross-framework: NSString */
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/init(rawString:)
func NewMEEmailAddressWithRawString(rawString objc.IObject /* cross-framework: NSString */) MEEmailAddress {
	instance := getMEEmailAddressClass().Alloc()
	rv := objc.Send[MEEmailAddress](instance.ID, objc.Sel("initWithRawString:"), rawString)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/addressString
func (m_ MEEmailAddress) AddressString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("addressString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEEmailAddress/rawString
func (m_ MEEmailAddress) RawString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("rawString"))
	return rv
}


