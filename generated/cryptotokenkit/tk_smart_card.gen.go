// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCard */


/* debug [class_header]: Header for TKSmartCard */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCard */
// An interface definition for the [TKSmartCard] class.
type ITKSmartCard interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCard */
	// properties:
	AllowedProtocols() TKSmartCardProtocol
	SetAllowedProtocols(value TKSmartCardProtocol)
	Cla() unsafe.Pointer
	SetCla(value unsafe.Pointer)
	Context() objc.ID
	SetContext(value objc.ID)
	CurrentProtocol() TKSmartCardProtocol
	Sensitive() bool
	SetSensitive(value bool)
	Valid() bool
	Slot() ITKSmartCardSlot
	UseCommandChaining() bool
	SetUseCommandChaining(value bool)
	UseExtendedLength() bool
	SetUseExtendedLength(value bool)
	IsSensitive() bool
	SetIsSensitive(value bool)
	IsValid() bool
	SetIsValid(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCard */
	// methods:
	BeginSessionWithReply(reply unsafe.Pointer)
	EndSession()
	InSessionWithErrorExecuteBlock(error_ unsafe.Pointer, block unsafe.Pointer) bool
	SendInsP1P2DataLeReply(ins unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, requestData objc.IObject /* cross-framework: NSData */, le objc.IObject /* cross-framework: NSNumber */, reply unsafe.Pointer)
	SendInsP1P2DataLeSwError(ins unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, requestData objc.IObject /* cross-framework: NSData */, le objc.IObject /* cross-framework: NSNumber */, sw unsafe.Pointer, error_ unsafe.Pointer) foundation.Data
	TransmitRequestReply(request objc.IObject /* cross-framework: NSData */, reply unsafe.Pointer)
	UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU objc.IObject /* cross-framework: NSData */, currentPINByteOffset int, newPINByteOffset int) ITKSmartCardUserInteractionForSecurePINChange
	UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU objc.IObject /* cross-framework: NSData */, PINByteOffset int) ITKSmartCardUserInteractionForSecurePINVerification
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCard */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardClass) Alloc() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCard */
// A representation of a smart card.
//
// This class provides an interface for managing sessions with a smart card, transmitting requests, and facilitating user interaction. You can create a object when a smart card is inserted into a slot, by calling the method on the corresponding object. To start communicating with the smart card, call the method on the object. Once an exclusive session has been established, you transmit data using the method. After you’ve finished communicating with a smart card, you call the method. If the smart card is physically removed from its slot, the session object becomes invalid, and any further calls to will return an error. You can use Key-Value Observing on the property to be notified when a smart card is invalidated, due to being removed from the slot or another reason.


// A representation of a smart card.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCard *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCard */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCard */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCard */

// Begins a session with the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/beginSession(reply:)
func (t_ TKSmartCard) BeginSessionWithReply(reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginSessionWithReply:"), reply)
}/* debug [instance_methods/method]: BeginSessionWithReply */


// Completes any pending transmissions and ends the session to the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/endSession()
func (t_ TKSmartCard) EndSession() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endSession"))
}/* debug [instance_methods/method]: EndSession */


// Synchronously begins a session, executes the given block, and ends the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/inSessionWithError:executeBlock:
func (t_ TKSmartCard) InSessionWithErrorExecuteBlock(error_ unsafe.Pointer, block unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("inSessionWithError:executeBlock:"), error_, block)
	return rv
}/* debug [instance_methods/method]: InSessionWithErrorExecuteBlock */


// Asynchronously transmits an APDU command to the card, returning the response in a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/sendIns:p1:p2:data:le:reply:
func (t_ TKSmartCard) SendInsP1P2DataLeReply(ins unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, requestData objc.IObject /* cross-framework: NSData */, le objc.IObject /* cross-framework: NSNumber */, reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sendIns:p1:p2:data:le:reply:"), ins, p1, p2, requestData, le, reply)
}/* debug [instance_methods/method]: SendInsP1P2DataLeReply */


// Synchronously transmits an APDU command to the card and returns the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/sendIns:p1:p2:data:le:sw:error:
func (t_ TKSmartCard) SendInsP1P2DataLeSwError(ins unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, requestData objc.IObject /* cross-framework: NSData */, le objc.IObject /* cross-framework: NSNumber */, sw unsafe.Pointer, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](t_.ID, objc.Sel("sendIns:p1:p2:data:le:sw:error:"), ins, p1, p2, requestData, le, sw, error_)
	return rv
}/* debug [instance_methods/method]: SendInsP1P2DataLeSwError */


// Transmits data in Application Protocol Data Unit (APDU) format to the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/transmit(_:reply:)
func (t_ TKSmartCard) TransmitRequestReply(request objc.IObject /* cross-framework: NSData */, reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("transmitRequest:reply:"), request, reply)
}/* debug [instance_methods/method]: TransmitRequestReply */


// Creates a new user interaction object for secure PIN change using the smart card reader facilities (typically a HW keypad).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/userInteractionForSecurePINChange(_:apdu:currentPINByteOffset:newPINByteOffset:)
func (t_ TKSmartCard) UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU objc.IObject /* cross-framework: NSData */, currentPINByteOffset int, newPINByteOffset int) ITKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](t_.ID, objc.Sel("userInteractionForSecurePINChangeWithPINFormat:APDU:currentPINByteOffset:newPINByteOffset:"), PINFormat, APDU, currentPINByteOffset, newPINByteOffset)
	return rv
}/* debug [instance_methods/method]: UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset */


// Creates and returns a new user interaction object for secure PIN verification using the Smart Card reader facilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/userInteractionForSecurePINVerification(_:apdu:pinByteOffset:)
func (t_ TKSmartCard) UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU objc.IObject /* cross-framework: NSData */, PINByteOffset int) ITKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](t_.ID, objc.Sel("userInteractionForSecurePINVerificationWithPINFormat:APDU:PINByteOffset:"), PINFormat, APDU, PINByteOffset)
	return rv
}/* debug [instance_methods/method]: UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCard */

// The protocols allowed for communication with the Smart Card. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/allowedProtocols
func (t_ TKSmartCard) AllowedProtocols() TKSmartCardProtocol {
	rv := objc.Send[TKSmartCardProtocol](t_.ID, objc.Sel("allowedProtocols"))
	return rv
}/* debug [instance_properties/getter]: allowedProtocols */


// The protocols allowed for communication with the Smart Card. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/allowedProtocols
func (t_ TKSmartCard) SetAllowedProtocols(value TKSmartCardProtocol) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedProtocols:"), value)
}/* debug [instance_properties/setter]: allowedProtocols */


// The CLA byte used for APDU transmission. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/cla
func (t_ TKSmartCard) Cla() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("cla"))
	return rv
}/* debug [instance_properties/getter]: cla */


// The CLA byte used for APDU transmission. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/cla
func (t_ TKSmartCard) SetCla(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCla:"), value)
}/* debug [instance_properties/setter]: cla */


// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) Context() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// User-specified information. This property is automatically set to if the Smart Card is removed or another object begins a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t_ TKSmartCard) SetContext(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContext:"), value)
}/* debug [instance_properties/setter]: context */


// The protocol used for communication with the Smart Card. Returns if no session is currently established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/currentProtocol
func (t_ TKSmartCard) CurrentProtocol() TKSmartCardProtocol {
	rv := objc.Send[TKSmartCardProtocol](t_.ID, objc.Sel("currentProtocol"))
	return rv
}/* debug [instance_properties/getter]: currentProtocol */


// Whether sessions established for the Smart Card should be considered sensitive. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/isSensitive
func (t_ TKSmartCard) Sensitive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("sensitive"))
	return rv
}/* debug [instance_properties/getter]: sensitive */


// Whether sessions established for the Smart Card should be considered sensitive. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/isSensitive
func (t_ TKSmartCard) SetSensitive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSensitive:"), value)
}/* debug [instance_properties/setter]: sensitive */


// Whether the Smart Card is valid and accessible from its slot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/isValid
func (t_ TKSmartCard) Valid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("valid"))
	return rv
}/* debug [instance_properties/getter]: valid */


// The slot in which the Smart Card is inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/slot
func (t_ TKSmartCard) Slot() ITKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("slot"))
	return rv
}/* debug [instance_properties/getter]: slot */


// Whether to use command chaining of APDU with a data field longer than 255 bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useCommandChaining
func (t_ TKSmartCard) UseCommandChaining() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("useCommandChaining"))
	return rv
}/* debug [instance_properties/getter]: useCommandChaining */


// Whether to use command chaining of APDU with a data field longer than 255 bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useCommandChaining
func (t_ TKSmartCard) SetUseCommandChaining(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUseCommandChaining:"), value)
}/* debug [instance_properties/setter]: useCommandChaining */


// Whether to use extended length APDU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useExtendedLength
func (t_ TKSmartCard) UseExtendedLength() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("useExtendedLength"))
	return rv
}/* debug [instance_properties/getter]: useExtendedLength */


// Whether to use extended length APDU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useExtendedLength
func (t_ TKSmartCard) SetUseExtendedLength(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUseExtendedLength:"), value)
}/* debug [instance_properties/setter]: useExtendedLength */


// Whether sessions established for the Smart Card should be considered sensitive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/issensitive
func (t_ TKSmartCard) IsSensitive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSensitive"))
	return rv
}/* debug [instance_properties/getter]: isSensitive */


// Whether sessions established for the Smart Card should be considered sensitive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/issensitive
func (t_ TKSmartCard) SetIsSensitive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSensitive:"), value)
}/* debug [instance_properties/setter]: isSensitive */


// Whether the Smart Card is valid and accessible from its slot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/isvalid
func (t_ TKSmartCard) IsValid() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isValid"))
	return rv
}/* debug [instance_properties/getter]: isValid */


// Whether the Smart Card is valid and accessible from its slot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcard/isvalid
func (t_ TKSmartCard) SetIsValid(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsValid:"), value)
}/* debug [instance_properties/setter]: isValid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCard */



