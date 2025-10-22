// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DecimalNumberHandler] class.
var (
	DecimalNumberHandlerClass     _DecimalNumberHandlerClass
	DecimalNumberHandlerClassOnce sync.Once
)

func getDecimalNumberHandlerClass() _DecimalNumberHandlerClass {
	DecimalNumberHandlerClassOnce.Do(func() {
		DecimalNumberHandlerClass = _DecimalNumberHandlerClass{objc.GetClass("NSDecimalNumberHandler")}
	})
	return DecimalNumberHandlerClass
}

type _DecimalNumberHandlerClass struct {
	class objc.Class
}

// An interface definition for the [DecimalNumberHandler] class.
type IDecimalNumberHandler interface {
	objectivec.IObject
	RoundingBehavior() NSDecimalNumberHandler
	SetRoundingBehavior(value IDecimalNumberHandler)
	RoundingIncrement() Number
	SetRoundingIncrement(value INumber)
	RoundingMode() unsafe.Pointer
	SetRoundingMode(value unsafe.Pointer)
}

// A class that adopts the decimal number behaviors protocol.
//
// This class allows you to set the way an object rounds off and handles errors, without having to create a custom class. You can use an instance of this class as an argument to any of the methods that end with . If you don’t think you need special behavior, you probably don’t need this class—it is likely that ’s default behavior will suit your needs. For more information, see the protocol specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumberHandler
type DecimalNumberHandler struct {
	objectivec.Object
}

// DecimalNumberHandlerFrom constructs a [DecimalNumberHandler] from an unsafe.Pointer.
//
// A class that adopts the decimal number behaviors protocol.
func DecimalNumberHandlerFrom(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DecimalNumberHandlerClass) Alloc() DecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DecimalNumberHandlerClass) New() DecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DecimalNumberHandler) Init() DecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DecimalNumberHandler) Autorelease() DecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDecimalNumberHandler creates a new DecimalNumberHandler instance.
func NewDecimalNumberHandler() DecimalNumberHandler {
	return getDecimalNumberHandlerClass().New()
}


// The rounding behavior used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingbehavior
func (d_ DecimalNumberHandler) RoundingBehavior() NSDecimalNumberHandler {
	rv := objc.Send[NSDecimalNumberHandler](d_.ID, objc.Sel("roundingBehavior"))
	return rv
}


// SetRoundingBehavior sets the value of the roundingBehavior property.
// The rounding behavior used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingbehavior
func (d_ DecimalNumberHandler) SetRoundingBehavior(value IDecimalNumberHandler) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRoundingBehavior:"), value)
}

// The rounding increment used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingincrement
func (d_ DecimalNumberHandler) RoundingIncrement() Number {
	rv := objc.Send[Number](d_.ID, objc.Sel("roundingIncrement"))
	return rv
}


// SetRoundingIncrement sets the value of the roundingIncrement property.
// The rounding increment used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingincrement
func (d_ DecimalNumberHandler) SetRoundingIncrement(value INumber) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRoundingIncrement:"), value)
}

// The rounding mode used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingmode-swift.property
func (d_ DecimalNumberHandler) RoundingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("roundingMode"))
	return rv
}


// SetRoundingMode sets the value of the roundingMode property.
// The rounding mode used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/roundingmode-swift.property
func (d_ DecimalNumberHandler) SetRoundingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRoundingMode:"), value)
}



