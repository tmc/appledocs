// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCard] class.
var (
	TKSmartCardClass     _TKSmartCardClass
	TKSmartCardClassOnce sync.Once
)

func getTKSmartCardClass() _TKSmartCardClass {
	TKSmartCardClassOnce.Do(func() {
		TKSmartCardClass = _TKSmartCardClass{objc.GetClass("TKSmartCard")}
	})
	return TKSmartCardClass
}

type _TKSmartCardClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCard] class.
type ITKSmartCard interface {
	objectivec.IObject
	BeginSessionWithReply(reply unsafe.Pointer)
	EndSession()
	Context() objc.ID
	SetContext(value objc.ID)
	AllowedProtocols() unsafe.Pointer
	SetAllowedProtocols(value unsafe.Pointer)
	Cla() unsafe.Pointer
	SetCla(value unsafe.Pointer)
	CurrentProtocol() unsafe.Pointer
	SetCurrentProtocol(value unsafe.Pointer)
	IsSensitive() bool
	SetIsSensitive(value bool)
	IsValid() bool
	SetIsValid(value bool)
	Slot() TKSmartCardSlot
	SetSlot(value ITKSmartCardSlot)
	UseCommandChaining() bool
	SetUseCommandChaining(value bool)
	UseExtendedLength() bool
	SetUseExtendedLength(value bool)
}

// A representation of a smart card.
//
// This class provides an interface for managing sessions with a smart card, transmitting requests, and facilitating user interaction. You can create a object when a smart card is inserted into a slot, by calling the method on the corresponding object. To start communicating with the smart card, call the method on the object. Once an exclusive session has been established, you transmit data using the method. After you’ve finished communicating with a smart card, you call the method. If the smart card is physically removed from its slot, the session object becomes invalid, and any further calls to will return an error. You can use Key-Value Observing on the property to be notified when a smart card is invalidated, due to being removed from the slot or another reason.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard
type TKSmartCard struct {
	objectivec.Object
}

// TKSmartCardFrom constructs a [TKSmartCard] from an unsafe.Pointer.
//
// A representation of a smart card.
func TKSmartCardFrom(ptr unsafe.Pointer) TKSmartCard {
	return TKSmartCard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardClass) Alloc() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardClass) New() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCard) Init() TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCard) Autorelease() TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCard creates a new TKSmartCard instance.
func NewTKSmartCard() TKSmartCard {
	return getTKSmartCardClass().New()
}


// Begins a session with the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/beginSession(reply:)
func (t_ TKSmartCard) BeginSessionWithReply(reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginSessionWithReply:"), reply)
}

// Completes any pending transmissions and ends the session to the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/endSession()
func (t_ TKSmartCard) EndSession() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endSession"))
}

// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) Context() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("context"))
	return rv
}


// SetContext sets the value of the context property.
// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) SetContext(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContext:"), value)
}

// The protocols allowed for communication with the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/allowedprotocols
func (t_ TKSmartCard) AllowedProtocols() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("allowedProtocols"))
	return rv
}


// SetAllowedProtocols sets the value of the allowedProtocols property.
// The protocols allowed for communication with the Smart Card.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/allowedprotocols
func (t_ TKSmartCard) SetAllowedProtocols(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedProtocols:"), value)
}

// The CLA byte used for APDU transmission.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/cla
func (t_ TKSmartCard) Cla() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("cla"))
	return rv
}


// SetCla sets the value of the cla property.
// The CLA byte used for APDU transmission.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/cla
func (t_ TKSmartCard) SetCla(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCla:"), value)
}

// The protocol used for communication with the Smart Card. Returns
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/currentprotocol
func (t_ TKSmartCard) CurrentProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("currentProtocol"))
	return rv
}


// SetCurrentProtocol sets the value of the currentProtocol property.
// The protocol used for communication with the Smart Card. Returns

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/currentprotocol
func (t_ TKSmartCard) SetCurrentProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentProtocol:"), value)
}

// Whether sessions established for the Smart Card should be considered sensitive.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/issensitive
func (t_ TKSmartCard) IsSensitive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSensitive"))
	return rv
}


// SetIsSensitive sets the value of the isSensitive property.
// Whether sessions established for the Smart Card should be considered sensitive.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/issensitive
func (t_ TKSmartCard) SetIsSensitive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSensitive:"), value)
}

// Whether the Smart Card is valid and accessible from its slot.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/isvalid
func (t_ TKSmartCard) IsValid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isValid"))
	return rv
}


// SetIsValid sets the value of the isValid property.
// Whether the Smart Card is valid and accessible from its slot.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/isvalid
func (t_ TKSmartCard) SetIsValid(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsValid:"), value)
}

// The slot in which the Smart Card is inserted.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/slot
func (t_ TKSmartCard) Slot() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("slot"))
	return rv
}


// SetSlot sets the value of the slot property.
// The slot in which the Smart Card is inserted.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/slot
func (t_ TKSmartCard) SetSlot(value ITKSmartCardSlot) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSlot:"), value)
}

// Whether to use command chaining of APDU with a data field longer than 255 bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/usecommandchaining
func (t_ TKSmartCard) UseCommandChaining() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("useCommandChaining"))
	return rv
}


// SetUseCommandChaining sets the value of the useCommandChaining property.
// Whether to use command chaining of APDU with a data field longer than 255 bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/usecommandchaining
func (t_ TKSmartCard) SetUseCommandChaining(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUseCommandChaining:"), value)
}

// Whether to use extended length APDU.
//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/useextendedlength
func (t_ TKSmartCard) UseExtendedLength() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("useExtendedLength"))
	return rv
}


// SetUseExtendedLength sets the value of the useExtendedLength property.
// Whether to use extended length APDU.

//
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/useextendedlength
func (t_ TKSmartCard) SetUseExtendedLength(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUseExtendedLength:"), value)
}



