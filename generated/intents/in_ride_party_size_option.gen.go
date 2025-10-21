// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




