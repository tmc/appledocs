// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardATR */


/* debug [class_header]: Header for TKSmartCardATR */
// The class instance for the [TKSmartCardATR] class.
var (
	TKSmartCardATRClass     _TKSmartCardATRClass
	TKSmartCardATRClassOnce sync.Once
)

func getTKSmartCardATRClass() _TKSmartCardATRClass {
	TKSmartCardATRClassOnce.Do(func() {
		TKSmartCardATRClass = _TKSmartCardATRClass{objc.GetClass("TKSmartCardATR")}
	})
	return TKSmartCardATRClass
}

type _TKSmartCardATRClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardATR */
// An interface definition for the [TKSmartCardATR] class.
type ITKSmartCardATR interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardATR */
	// properties:
	Bytes() objc.IObject /* cross-framework: NSData */
	HistoricalBytes() objc.IObject /* cross-framework: NSData */
	HistoricalRecords() []TKCompactTLVRecord
	Protocols() []foundation.Number
	Atr() ITKSmartCardATR
	SetAtr(value ITKSmartCardATR)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardATR */
	// methods:
	InterfaceGroupAtIndex(index int) ITKSmartCardATRInterfaceGroup
	InterfaceGroupForProtocol(protocol_ TKSmartCardProtocol) ITKSmartCardATRInterfaceGroup
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardATR */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardATRClass) Alloc() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardATRClass) New() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardATR) Init() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardATR) Autorelease() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardATR creates a new TKSmartCardATR instance.
func NewTKSmartCardATR() TKSmartCardATR {
	return getTKSmartCardATRClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardATR */
// A parsed ATR (Answer To Reset) message from a Smart Card.
//
// This class declares a programmatic interface to parsing an ATR from data or a byte stream, and accessing the individual parts.


// A parsed ATR (Answer To Reset) message from a Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR
type TKSmartCardATR struct {
	objectivec.Object
}

// TKSmartCardATRFrom constructs a [TKSmartCardATR] from an unsafe.Pointer.
//
// A parsed ATR (Answer To Reset) message from a Smart Card.
func TKSmartCardATRFrom(ptr unsafe.Pointer) TKSmartCardATR {
	return TKSmartCardATR{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardATR */

// Initializes a object from a provided data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/init(bytes:)
func NewTKSmartCardATRWithBytes(bytes objc.IObject /* cross-framework: NSData */) TKSmartCardATR {
	instance := getTKSmartCardATRClass().Alloc()
	rv := objc.Send[TKSmartCardATR](instance.ID, objc.Sel("initWithBytes:"), bytes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKSmartCardATRWithBytes */


// Initializes a object from a provided data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/init(source:)
func NewTKSmartCardATRWithSource(source unsafe.Pointer) TKSmartCardATR {
	instance := getTKSmartCardATRClass().Alloc()
	rv := objc.Send[TKSmartCardATR](instance.ID, objc.Sel("initWithSource:"), source)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKSmartCardATRWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardATR */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardATR */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardATR */

// Returns the interface group at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/interfaceGroup(at:)
func (t_ TKSmartCardATR) InterfaceGroupAtIndex(index int) ITKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t_.ID, objc.Sel("interfaceGroupAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: InterfaceGroupAtIndex */


// Returns the interface group with the specified protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/interfaceGroup(for:)
func (t_ TKSmartCardATR) InterfaceGroupForProtocol(protocol_ TKSmartCardProtocol) ITKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t_.ID, objc.Sel("interfaceGroupForProtocol:"), protocol_)
	return rv
}/* debug [instance_methods/method]: InterfaceGroupForProtocol */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardATR */

// The ATR message data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/bytes
func (t_ TKSmartCardATR) Bytes() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("bytes"))
	return rv
}/* debug [instance_properties/getter]: bytes */


// The ATR historical bytes, not including interface bytes or the TCK (check byte).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/historicalBytes
func (t_ TKSmartCardATR) HistoricalBytes() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("historicalBytes"))
	return rv
}/* debug [instance_properties/getter]: historicalBytes */


// A list of compact TLV records parsed from historical bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/historicalRecords
func (t_ TKSmartCardATR) HistoricalRecords() []TKCompactTLVRecord {
	rv := objc.Send[[]TKCompactTLVRecord](t_.ID, objc.Sel("historicalRecords"))
	return rv
}/* debug [instance_properties/getter]: historicalRecords */


// An array of protocols indicated in the ATR
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/protocols
func (t_ TKSmartCardATR) Protocols() []foundation.Number {
	rv := objc.Send[[]foundation.Number](t_.ID, objc.Sel("protocols"))
	return rv
}/* debug [instance_properties/getter]: protocols */


// The ATR (Answer to Reset) of the inserted Smart Card, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardslot/atr
func (t_ TKSmartCardATR) Atr() ITKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t_.ID, objc.Sel("atr"))
	return rv
}/* debug [instance_properties/getter]: atr */


// The ATR (Answer to Reset) of the inserted Smart Card, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardslot/atr
func (t_ TKSmartCardATR) SetAtr(value ITKSmartCardATR) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAtr:"), value)
}/* debug [instance_properties/setter]: atr */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardATR */


