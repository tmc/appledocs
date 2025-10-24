// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKTokenSmartCardPINAuthOperation */


/* debug [class_header]: Header for TKTokenSmartCardPINAuthOperation */
// The class instance for the [TKTokenSmartCardPINAuthOperation] class.
var (
	TKTokenSmartCardPINAuthOperationClass     _TKTokenSmartCardPINAuthOperationClass
	TKTokenSmartCardPINAuthOperationClassOnce sync.Once
)

func getTKTokenSmartCardPINAuthOperationClass() _TKTokenSmartCardPINAuthOperationClass {
	TKTokenSmartCardPINAuthOperationClassOnce.Do(func() {
		TKTokenSmartCardPINAuthOperationClass = _TKTokenSmartCardPINAuthOperationClass{objc.GetClass("TKTokenSmartCardPINAuthOperation")}
	})
	return TKTokenSmartCardPINAuthOperationClass
}

type _TKTokenSmartCardPINAuthOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenSmartCardPINAuthOperation */
// An interface definition for the [TKTokenSmartCardPINAuthOperation] class.
type ITKTokenSmartCardPINAuthOperation interface {
	ITKTokenAuthOperation
	
/* debug [class_interface_properties]: Properties for TKTokenSmartCardPINAuthOperation */
	// properties:
	APDUTemplate() objc.IObject /* cross-framework: NSData */
	SetAPDUTemplate(value objc.IObject /* cross-framework: NSData */)
	PIN() objc.IObject /* cross-framework: NSString */
	SetPIN(value objc.IObject /* cross-framework: NSString */)
	PINByteOffset() int
	SetPINByteOffset(value int)
	PINFormat() ITKSmartCardPINFormat
	SetPINFormat(value ITKSmartCardPINFormat)
	SmartCard() ITKSmartCard
	SetSmartCard(value ITKSmartCard)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenSmartCardPINAuthOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenSmartCardPINAuthOperation */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenSmartCardPINAuthOperationClass) Alloc() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenSmartCardPINAuthOperationClass) New() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenSmartCardPINAuthOperation) Init() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenSmartCardPINAuthOperation) Autorelease() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenSmartCardPINAuthOperation creates a new TKTokenSmartCardPINAuthOperation instance.
func NewTKTokenSmartCardPINAuthOperation() TKTokenSmartCardPINAuthOperation {
	return getTKTokenSmartCardPINAuthOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenSmartCardPINAuthOperation */
// A Smart Card PIN authentication operation.


// A Smart Card PIN authentication operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation
type TKTokenSmartCardPINAuthOperation struct {
	TKTokenAuthOperation
}

// TKTokenSmartCardPINAuthOperationFrom constructs a [TKTokenSmartCardPINAuthOperation] from an unsafe.Pointer.
//
// A Smart Card PIN authentication operation.
func TKTokenSmartCardPINAuthOperationFrom(ptr unsafe.Pointer) TKTokenSmartCardPINAuthOperation {
	return TKTokenSmartCardPINAuthOperation{
		TKTokenAuthOperation: TKTokenAuthOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenSmartCardPINAuthOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenSmartCardPINAuthOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenSmartCardPINAuthOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenSmartCardPINAuthOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenSmartCardPINAuthOperation */

// The template into which the PIN is filled in. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/apduTemplate
func (t_ TKTokenSmartCardPINAuthOperation) APDUTemplate() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("APDUTemplate"))
	return rv
}/* debug [instance_properties/getter]: APDUTemplate */


// The template into which the PIN is filled in. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/apduTemplate
func (t_ TKTokenSmartCardPINAuthOperation) SetAPDUTemplate(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAPDUTemplate:"), value)
}/* debug [instance_properties/setter]: APDUTemplate */


// The PIN value resulting from performing the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pin
func (t_ TKTokenSmartCardPINAuthOperation) PIN() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("PIN"))
	return rv
}/* debug [instance_properties/getter]: PIN */


// The PIN value resulting from performing the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pin
func (t_ TKTokenSmartCardPINAuthOperation) SetPIN(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPIN:"), value)
}/* debug [instance_properties/setter]: PIN */


// The offset, in bytes, within the APDU template to mark the location for filling in the PIN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinByteOffset
func (t_ TKTokenSmartCardPINAuthOperation) PINByteOffset() int {
	rv := objc.Send[int](t_.ID, objc.Sel("PINByteOffset"))
	return rv
}/* debug [instance_properties/getter]: PINByteOffset */


// The offset, in bytes, within the APDU template to mark the location for filling in the PIN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinByteOffset
func (t_ TKTokenSmartCardPINAuthOperation) SetPINByteOffset(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINByteOffset:"), value)
}/* debug [instance_properties/setter]: PINByteOffset */


// The PIN format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinFormat
func (t_ TKTokenSmartCardPINAuthOperation) PINFormat() ITKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](t_.ID, objc.Sel("PINFormat"))
	return rv
}/* debug [instance_properties/getter]: PINFormat */


// The PIN format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinFormat
func (t_ TKTokenSmartCardPINAuthOperation) SetPINFormat(value ITKSmartCardPINFormat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINFormat:"), value)
}/* debug [instance_properties/setter]: PINFormat */


// A Smart Card to which the formatted APDU is sent in order to authenticate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/smartCard
func (t_ TKTokenSmartCardPINAuthOperation) SmartCard() ITKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("smartCard"))
	return rv
}/* debug [instance_properties/getter]: smartCard */


// A Smart Card to which the formatted APDU is sent in order to authenticate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/smartCard
func (t_ TKTokenSmartCardPINAuthOperation) SetSmartCard(value ITKSmartCard) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartCard:"), value)
}/* debug [instance_properties/setter]: smartCard */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenSmartCardPINAuthOperation */





