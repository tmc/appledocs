// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Heading] class.
var (
	headingClass     _HeadingClass
	headingClassOnce sync.Once
)

func getHeadingClass() _HeadingClass {
	headingClassOnce.Do(func() {
		headingClass = _HeadingClass{objc.GetClass("CLHeading")}
	})
	return headingClass
}

type _HeadingClass struct {
	class objc.Class
}

// An interface definition for the [Heading] class.
type IHeading interface {
	objectivec.IObject
}

// The orientation of the user’s device, relative to true or magnetic north.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading
type Heading struct {
	objectivec.Object
}

// HeadingFrom constructs a [Heading] from an unsafe.Pointer.
//
// The orientation of the user’s device, relative to true or magnetic north.
func HeadingFrom(ptr unsafe.Pointer) Heading {
	return Heading{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HeadingClass) Alloc() Heading {
	rv := objc.Send[Heading](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HeadingClass) New() Heading {
	rv := objc.Send[Heading](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ Heading) Init() Heading {
	rv := objc.Send[Heading](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ Heading) Autorelease() Heading {
	rv := objc.Send[Heading](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeading creates a new Heading instance.
func NewHeading() Heading {
	return getHeadingClass().New()
}




