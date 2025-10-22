// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SQuantityItem] class.
var (
	SQuantityItemClass     _SQuantityItemClass
	SQuantityItemClassOnce sync.Once
)

func getSQuantityItemClass() _SQuantityItemClass {
	SQuantityItemClassOnce.Do(func() {
		SQuantityItemClass = _SQuantityItemClass{objc.GetClass("CLSQuantityItem")}
	})
	return SQuantityItemClass
}

type _SQuantityItemClass struct {
	class objc.Class
}

// An interface definition for the [SQuantityItem] class.
type ISQuantityItem interface {
	ISActivityItem
	Quantity() float64
	SetQuantity(value float64)
}

// Activity information that signifies a quantity.
//
// Use an activity item of this type to associate a discrete value with a task. For example, you might use it to indicate how many times the user requested a hint while taking a quiz.


// Activity information that signifies a quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSQuantityItem

type SQuantityItem struct {
	SActivityItem
}

// SQuantityItemFrom constructs a [SQuantityItem] from an unsafe.Pointer.
//
// Activity information that signifies a quantity.
func SQuantityItemFrom(ptr unsafe.Pointer) SQuantityItem {
	return SQuantityItem{
		SActivityItem: SActivityItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SQuantityItemClass) Alloc() SQuantityItem {
	rv := objc.Send[SQuantityItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SQuantityItemClass) New() SQuantityItem {
	rv := objc.Send[SQuantityItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SQuantityItem) Init() SQuantityItem {
	rv := objc.Send[SQuantityItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SQuantityItem) Autorelease() SQuantityItem {
	rv := objc.Send[SQuantityItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSQuantityItem creates a new SQuantityItem instance.
func NewSQuantityItem() SQuantityItem {
	return getSQuantityItemClass().New()
}




// Initializes an activity item that records a discrete quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSQuantityItem/init(identifier:title:)

func NewSQuantityItemWithIdentifierTitle(identifier string, title string) SQuantityItem {
	instance := getSQuantityItemClass().Alloc()
	rv := objc.Send[SQuantityItem](instance.ID, objc.Sel("initWithIdentifier:title:"), objc.String(identifier), objc.String(title))
	rv.Autorelease()
	return rv
}



// A quantity associated with the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSQuantityItem/quantity

func (s_ SQuantityItem) Quantity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("quantity"))
	return rv
}


// A quantity associated with the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSQuantityItem/quantity

func (s_ SQuantityItem) SetQuantity(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQuantity:"), value)
}


