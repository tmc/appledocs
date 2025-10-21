// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INRidePartySizeOption] class.
var (
	INRidePartySizeOptionClass     _INRidePartySizeOptionClass
	INRidePartySizeOptionClassOnce sync.Once
)

func getINRidePartySizeOptionClass() _INRidePartySizeOptionClass {
	INRidePartySizeOptionClassOnce.Do(func() {
		INRidePartySizeOptionClass = _INRidePartySizeOptionClass{objc.GetClass("INRidePartySizeOption")}
	})
	return INRidePartySizeOptionClass
}

type _INRidePartySizeOptionClass struct {
	class objc.Class
}

// An interface definition for the [INRidePartySizeOption] class.
type IINRidePartySizeOption interface {
	objectivec.IObject
}

// The price of a ride involving the specified number of people.
//
// An object describes many passengers and the special pricing that applies to a party of that size. When configuring a object, you create one or more instances of this class to specify pricing for the user’s party. Present this information to the user along with the other ride option information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePartySizeOption
type INRidePartySizeOption struct {
	objectivec.Object
}

// INRidePartySizeOptionFrom constructs a [INRidePartySizeOption] from an unsafe.Pointer.
//
// The price of a ride involving the specified number of people.
func INRidePartySizeOptionFrom(ptr unsafe.Pointer) INRidePartySizeOption {
	return INRidePartySizeOption{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INRidePartySizeOptionClass) Alloc() INRidePartySizeOption {
	rv := objc.Send[INRidePartySizeOption](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRidePartySizeOptionClass) New() INRidePartySizeOption {
	rv := objc.Send[INRidePartySizeOption](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRidePartySizeOption) Init() INRidePartySizeOption {
	rv := objc.Send[INRidePartySizeOption](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRidePartySizeOption) Autorelease() INRidePartySizeOption {
	rv := objc.Send[INRidePartySizeOption](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRidePartySizeOption creates a new INRidePartySizeOption instance.
func NewINRidePartySizeOption() INRidePartySizeOption {
	return getINRidePartySizeOptionClass().New()
}


// The user-visible description of the party size.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/sizedescription
func (i_ INRidePartySizeOption) SizeDescription() string {
	rv := objc.Send[string](i_.ID, objc.Sel("sizeDescription"))
	return rv
}


// SetSizeDescription sets the value of the sizeDescription property.
// The user-visible description of the party size.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/sizedescription
func (i_ INRidePartySizeOption) SetSizeDescription(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSizeDescription:"), objc.String(value))
}

// The number of people in the party, specified as a minimum and maximum value.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/partysizerange
func (i_ INRidePartySizeOption) PartySizeRange() Range {
	rv := objc.Send[Range](i_.ID, objc.Sel("partySizeRange"))
	return rv
}


// SetPartySizeRange sets the value of the partySizeRange property.
// The number of people in the party, specified as a minimum and maximum value.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/partysizerange
func (i_ INRidePartySizeOption) SetPartySizeRange(value Range) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPartySizeRange:"), value)
}

// The pricing information for parties of the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/pricerange
func (i_ INRidePartySizeOption) PriceRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("priceRange"))
	return rv
}


// SetPriceRange sets the value of the priceRange property.
// The pricing information for parties of the specified size.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridepartysizeoption/pricerange
func (i_ INRidePartySizeOption) SetPriceRange(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPriceRange:"), value)
}



