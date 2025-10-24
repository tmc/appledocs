// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INCurrencyAmount] class.
var (
	INCurrencyAmountClass     _INCurrencyAmountClass
	INCurrencyAmountClassOnce sync.Once
)

func getINCurrencyAmountClass() _INCurrencyAmountClass {
	INCurrencyAmountClassOnce.Do(func() {
		INCurrencyAmountClass = _INCurrencyAmountClass{objc.GetClass("INCurrencyAmount")}
	})
	return INCurrencyAmountClass
}

type _INCurrencyAmountClass struct {
	class objc.Class
}

// An interface definition for the [INCurrencyAmount] class.
type IINCurrencyAmount interface {
	objectivec.IObject
	Amount() foundation.DecimalNumber
	SetAmount(value foundation.IDecimalNumber)
	CurrencyCode() string
	SetCurrencyCode(value string)
}

// An amount of money to transfer during a financial transaction.
//
// An object encapsulates a monetary value and the currency used to express that value. You use these objects to specify payment amounts when handling intents that involve the transfer of money.

// An amount of money to transfer during a financial transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCurrencyAmount
type INCurrencyAmount struct {
	objectivec.Object
}

// INCurrencyAmountFrom constructs a [INCurrencyAmount] from an unsafe.Pointer.
//
// An amount of money to transfer during a financial transaction.
func INCurrencyAmountFrom(ptr unsafe.Pointer) INCurrencyAmount {
	return INCurrencyAmount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INCurrencyAmountClass) Alloc() INCurrencyAmount {
	rv := objc.Send[INCurrencyAmount](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCurrencyAmountClass) New() INCurrencyAmount {
	rv := objc.Send[INCurrencyAmount](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCurrencyAmount) Init() INCurrencyAmount {
	rv := objc.Send[INCurrencyAmount](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCurrencyAmount) Autorelease() INCurrencyAmount {
	rv := objc.Send[INCurrencyAmount](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCurrencyAmount creates a new INCurrencyAmount instance.
func NewINCurrencyAmount() INCurrencyAmount {
	return getINCurrencyAmountClass().New()
}

// The monetary amount associated with the currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incurrencyamount/amount
func (i_ INCurrencyAmount) Amount() foundation.DecimalNumber {
	rv := objc.Send[foundation.DecimalNumber](i_.ID, objc.Sel("amount"))
	return rv
}

// The monetary amount associated with the currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incurrencyamount/amount
func (i_ INCurrencyAmount) SetAmount(value foundation.IDecimalNumber) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAmount:"), value)
}

// The ISO 4217 currency code that applies to the monetary amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incurrencyamount/currencycode
func (i_ INCurrencyAmount) CurrencyCode() string {
	rv := objc.Send[string](i_.ID, objc.Sel("currencyCode"))
	return rv
}

// The ISO 4217 currency code that applies to the monetary amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incurrencyamount/currencycode
func (i_ INCurrencyAmount) SetCurrencyCode(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrencyCode:"), objc.String(value))
}
