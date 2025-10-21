// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INDateComponentsRange] class.
var (
	INDateComponentsRangeClass     _INDateComponentsRangeClass
	INDateComponentsRangeClassOnce sync.Once
)

func getINDateComponentsRangeClass() _INDateComponentsRangeClass {
	INDateComponentsRangeClassOnce.Do(func() {
		INDateComponentsRangeClass = _INDateComponentsRangeClass{objc.GetClass("INDateComponentsRange")}
	})
	return INDateComponentsRangeClass
}

type _INDateComponentsRangeClass struct {
	class objc.Class
}

// An interface definition for the [INDateComponentsRange] class.
type IINDateComponentsRange interface {
	objectivec.IObject
}

// A span of time.
//
// Use an object to specify date- or time-related information when responding to an intent. For example, a ride service might use this object to specify possible pickup times for the user. You create date components range objects when providing a response that includes a time span. Use this object to specify a range of times to include in the corresponding response. When displaying the date range information to the user, Siri or Maps formats the information appropriately.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDateComponentsRange
type INDateComponentsRange struct {
	objectivec.Object
}

// INDateComponentsRangeFrom constructs a [INDateComponentsRange] from an unsafe.Pointer.
//
// A span of time.
func INDateComponentsRangeFrom(ptr unsafe.Pointer) INDateComponentsRange {
	return INDateComponentsRange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INDateComponentsRangeClass) Alloc() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INDateComponentsRangeClass) New() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INDateComponentsRange) Init() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INDateComponentsRange) Autorelease() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINDateComponentsRange creates a new INDateComponentsRange instance.
func NewINDateComponentsRange() INDateComponentsRange {
	return getINDateComponentsRangeClass().New()
}




