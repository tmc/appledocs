// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardPINFormat */


/* debug [class_header]: Header for TKSmartCardPINFormat */
// The class instance for the [TKSmartCardPINFormat] class.
var (
	TKSmartCardPINFormatClass     _TKSmartCardPINFormatClass
	TKSmartCardPINFormatClassOnce sync.Once
)

func getTKSmartCardPINFormatClass() _TKSmartCardPINFormatClass {
	TKSmartCardPINFormatClassOnce.Do(func() {
		TKSmartCardPINFormatClass = _TKSmartCardPINFormatClass{objc.GetClass("TKSmartCardPINFormat")}
	})
	return TKSmartCardPINFormatClass
}

type _TKSmartCardPINFormatClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardPINFormat */
// An interface definition for the [TKSmartCardPINFormat] class.
type ITKSmartCardPINFormat interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardPINFormat */
	// properties:
	Charset() TKSmartCardPINCharset
	SetCharset(value TKSmartCardPINCharset)
	Encoding() TKSmartCardPINEncoding
	SetEncoding(value TKSmartCardPINEncoding)
	MaxPINLength() int
	SetMaxPINLength(value int)
	MinPINLength() int
	SetMinPINLength(value int)
	PINBitOffset() int
	SetPINBitOffset(value int)
	PINBlockByteLength() int
	SetPINBlockByteLength(value int)
	PINJustification() TKSmartCardPINJustification
	SetPINJustification(value TKSmartCardPINJustification)
	PINLengthBitOffset() int
	SetPINLengthBitOffset(value int)
	PINLengthBitSize() int
	SetPINLengthBitSize(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardPINFormat */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardPINFormat */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardPINFormatClass) Alloc() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardPINFormatClass) New() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardPINFormat) Init() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardPINFormat) Autorelease() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardPINFormat creates a new TKSmartCardPINFormat instance.
func NewTKSmartCardPINFormat() TKSmartCardPINFormat {
	return getTKSmartCardPINFormatClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardPINFormat */
// The formatting properties for a PIN, such as character encoding and length constraints.
//
// You typically interact with objects when calling the and methods on an instance of .


// The formatting properties for a PIN, such as character encoding and length constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat
type TKSmartCardPINFormat struct {
	objectivec.Object
}

// TKSmartCardPINFormatFrom constructs a [TKSmartCardPINFormat] from an unsafe.Pointer.
//
// The formatting properties for a PIN, such as character encoding and length constraints.
func TKSmartCardPINFormatFrom(ptr unsafe.Pointer) TKSmartCardPINFormat {
	return TKSmartCardPINFormat{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardPINFormat *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardPINFormat */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardPINFormat */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardPINFormat */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardPINFormat */

// The format of PIN characters. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/charset-swift.property
func (t_ TKSmartCardPINFormat) Charset() TKSmartCardPINCharset {
	rv := objc.Send[TKSmartCardPINCharset](t_.ID, objc.Sel("charset"))
	return rv
}/* debug [instance_properties/getter]: charset */


// The format of PIN characters. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/charset-swift.property
func (t_ TKSmartCardPINFormat) SetCharset(value TKSmartCardPINCharset) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCharset:"), value)
}/* debug [instance_properties/setter]: charset */


// The encoding of PIN characters. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/encoding-swift.property
func (t_ TKSmartCardPINFormat) Encoding() TKSmartCardPINEncoding {
	rv := objc.Send[TKSmartCardPINEncoding](t_.ID, objc.Sel("encoding"))
	return rv
}/* debug [instance_properties/getter]: encoding */


// The encoding of PIN characters. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/encoding-swift.property
func (t_ TKSmartCardPINFormat) SetEncoding(value TKSmartCardPINEncoding) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEncoding:"), value)
}/* debug [instance_properties/setter]: encoding */


// The maximum number of characters to form a valid PIN. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/maxPINLength
func (t_ TKSmartCardPINFormat) MaxPINLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maxPINLength"))
	return rv
}/* debug [instance_properties/getter]: maxPINLength */


// The maximum number of characters to form a valid PIN. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/maxPINLength
func (t_ TKSmartCardPINFormat) SetMaxPINLength(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxPINLength:"), value)
}/* debug [instance_properties/setter]: maxPINLength */


// The minimum number of characters to form a valid PIN. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/minPINLength
func (t_ TKSmartCardPINFormat) MinPINLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("minPINLength"))
	return rv
}/* debug [instance_properties/getter]: minPINLength */


// The minimum number of characters to form a valid PIN. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/minPINLength
func (t_ TKSmartCardPINFormat) SetMinPINLength(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinPINLength:"), value)
}/* debug [instance_properties/setter]: minPINLength */


// The offset, in bits, within the PIN block to mark a location for filling in the formatted PIN, which is justified with respect to the property value. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBitOffset
func (t_ TKSmartCardPINFormat) PINBitOffset() int {
	rv := objc.Send[int](t_.ID, objc.Sel("PINBitOffset"))
	return rv
}/* debug [instance_properties/getter]: PINBitOffset */


// The offset, in bits, within the PIN block to mark a location for filling in the formatted PIN, which is justified with respect to the property value. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBitOffset
func (t_ TKSmartCardPINFormat) SetPINBitOffset(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINBitOffset:"), value)
}/* debug [instance_properties/setter]: PINBitOffset */


// The total length of the PIN block in bytes. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBlockByteLength
func (t_ TKSmartCardPINFormat) PINBlockByteLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("PINBlockByteLength"))
	return rv
}/* debug [instance_properties/getter]: PINBlockByteLength */


// The total length of the PIN block in bytes. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBlockByteLength
func (t_ TKSmartCardPINFormat) SetPINBlockByteLength(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINBlockByteLength:"), value)
}/* debug [instance_properties/setter]: PINBlockByteLength */


// The justification within the PIN block. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification
func (t_ TKSmartCardPINFormat) PINJustification() TKSmartCardPINJustification {
	rv := objc.Send[TKSmartCardPINJustification](t_.ID, objc.Sel("PINJustification"))
	return rv
}/* debug [instance_properties/getter]: PINJustification */


// The justification within the PIN block. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification
func (t_ TKSmartCardPINFormat) SetPINJustification(value TKSmartCardPINJustification) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINJustification:"), value)
}/* debug [instance_properties/setter]: PINJustification */


// The offset, in bits, within the PIN block to mark a location for filling in the PIN length, which is always left justified. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitOffset
func (t_ TKSmartCardPINFormat) PINLengthBitOffset() int {
	rv := objc.Send[int](t_.ID, objc.Sel("PINLengthBitOffset"))
	return rv
}/* debug [instance_properties/getter]: PINLengthBitOffset */


// The offset, in bits, within the PIN block to mark a location for filling in the PIN length, which is always left justified. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitOffset
func (t_ TKSmartCardPINFormat) SetPINLengthBitOffset(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINLengthBitOffset:"), value)
}/* debug [instance_properties/setter]: PINLengthBitOffset */


// The size, in bits, of the PIN length field. If set to , PIN length is not written. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitSize
func (t_ TKSmartCardPINFormat) PINLengthBitSize() int {
	rv := objc.Send[int](t_.ID, objc.Sel("PINLengthBitSize"))
	return rv
}/* debug [instance_properties/getter]: PINLengthBitSize */


// The size, in bits, of the PIN length field. If set to , PIN length is not written. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitSize
func (t_ TKSmartCardPINFormat) SetPINLengthBitSize(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINLengthBitSize:"), value)
}/* debug [instance_properties/setter]: PINLengthBitSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardPINFormat */



