// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SActivityItem] class.
var (
	SActivityItemClass     _SActivityItemClass
	SActivityItemClassOnce sync.Once
)

func getSActivityItemClass() _SActivityItemClass {
	SActivityItemClassOnce.Do(func() {
		SActivityItemClass = _SActivityItemClass{objc.GetClass("CLSActivityItem")}
	})
	return SActivityItemClass
}

type _SActivityItemClass struct {
	class objc.Class
}

// An interface definition for the [SActivityItem] class.
type ISActivityItem interface {
	ISObject
}

// An abstract base class for gathering information about an activity.
//
// You don’t typically use an instance of this class directly. Instead, use one of its subclasses to represent a particular activity metric. For example, use a to add a score to a activity.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivityItem
type SActivityItem struct {
	SObject
}

// SActivityItemFrom constructs a [SActivityItem] from an unsafe.Pointer.
//
// An abstract base class for gathering information about an activity.
func SActivityItemFrom(ptr unsafe.Pointer) SActivityItem {
	return SActivityItem{
		SObject: SObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SActivityItemClass) Alloc() SActivityItem {
	rv := objc.Send[SActivityItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SActivityItemClass) New() SActivityItem {
	rv := objc.Send[SActivityItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SActivityItem) Init() SActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SActivityItem) Autorelease() SActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSActivityItem creates a new SActivityItem instance.
func NewSActivityItem() SActivityItem {
	return getSActivityItemClass().New()
}


// An identifier for the activity item.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivityItem/identifier
func (s_ SActivityItem) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}

// A human readable name for the activity item.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivityItem/title
func (s_ SActivityItem) Title() string {
	rv := objc.Send[string](s_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// A human readable name for the activity item.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivityItem/title
func (s_ SActivityItem) SetTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), objc.String(value))
}



