// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchMoneyAmount */


/* debug [class_header]: Header for DDMatchMoneyAmount */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchMoneyAmount */
// An interface definition for the [DDMatchMoneyAmount] class.
type IDDMatchMoneyAmount interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchMoneyAmount */
	// properties:
	Amount() float64
	Currency() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchMoneyAmount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchMoneyAmount */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchMoneyAmountClass) Alloc() DDMatchMoneyAmount {
	rv := objc.Send[DDMatchMoneyAmount](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchMoneyAmount */
// An object that contains an amount of money that the data detection system matches.
//
// The DataDetection framework returns a match for an amount of money in a object, which contains an amount of money and an ISO currency code.


// An object that contains an amount of money that the data detection system matches.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchMoneyAmount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchMoneyAmount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchMoneyAmount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchMoneyAmount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchMoneyAmount */

// A number that represents an amount of money.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchMoneyAmount/amount
func (d_ DDMatchMoneyAmount) Amount() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("amount"))
	return rv
}/* debug [instance_properties/getter]: amount */


// A string that contains an ISO currency code, which the data detection system identifies from the matched string and user preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchMoneyAmount/currency
func (d_ DDMatchMoneyAmount) Currency() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("currency"))
	return rv
}/* debug [instance_properties/getter]: currency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchMoneyAmount */



