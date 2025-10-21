// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [DDMatchMoneyAmount] class.
var (
	DDMatchMoneyAmountClass     _DDMatchMoneyAmountClass
	DDMatchMoneyAmountClassOnce sync.Once
)

func getDDMatchMoneyAmountClass() _DDMatchMoneyAmountClass {
	DDMatchMoneyAmountClassOnce.Do(func() {
		DDMatchMoneyAmountClass = _DDMatchMoneyAmountClass{objc.GetClass("DDMatchMoneyAmount")}
	})
	return DDMatchMoneyAmountClass
}

type _DDMatchMoneyAmountClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchMoneyAmount] class.
type IDDMatchMoneyAmount interface {
	IDDMatch
}

// An object that contains an amount of money that the data detection system matches.
//
// The DataDetection framework returns a match for an amount of money in a object, which contains an amount of money and an ISO currency code.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchMoneyAmount
type DDMatchMoneyAmount struct {
	DDMatch
}

// DDMatchMoneyAmountFrom constructs a [DDMatchMoneyAmount] from an unsafe.Pointer.
//
// An object that contains an amount of money that the data detection system matches.
func DDMatchMoneyAmountFrom(ptr unsafe.Pointer) DDMatchMoneyAmount {
	return DDMatchMoneyAmount{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchMoneyAmountClass) Alloc() DDMatchMoneyAmount {
	rv := objc.Send[DDMatchMoneyAmount](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchMoneyAmountClass) New() DDMatchMoneyAmount {
	rv := objc.Send[DDMatchMoneyAmount](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchMoneyAmount) Init() DDMatchMoneyAmount {
	rv := objc.Send[DDMatchMoneyAmount](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchMoneyAmount) Autorelease() DDMatchMoneyAmount {
	rv := objc.Send[DDMatchMoneyAmount](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchMoneyAmount creates a new DDMatchMoneyAmount instance.
func NewDDMatchMoneyAmount() DDMatchMoneyAmount {
	return getDDMatchMoneyAmountClass().New()
}


// A number that represents an amount of money.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchMoneyAmount/amount
func (d_ DDMatchMoneyAmount) Amount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("amount"))
	return rv
}

// A string that contains an ISO currency code, which the data detection system identifies from the matched string and user preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchMoneyAmount/currency
func (d_ DDMatchMoneyAmount) Currency() appkit.string {
	rv := objc.Send[appkit.string](d_.ID, objc.Sel("currency"))
	return rv
}



