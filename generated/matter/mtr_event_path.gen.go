// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREventPath] class.
var (
	MTREventPathClass     _MTREventPathClass
	MTREventPathClassOnce sync.Once
)

func getMTREventPathClass() _MTREventPathClass {
	MTREventPathClassOnce.Do(func() {
		MTREventPathClass = _MTREventPathClass{objc.GetClass("MTREventPath")}
	})
	return MTREventPathClass
}

type _MTREventPathClass struct {
	class objc.Class
}

// An interface definition for the [MTREventPath] class.
type IMTREventPath interface {
	IMTRClusterPath
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventPath
type MTREventPath struct {
	MTRClusterPath
}

// MTREventPathFrom constructs a [MTREventPath] from an unsafe.Pointer.
func MTREventPathFrom(ptr unsafe.Pointer) MTREventPath {
	return MTREventPath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREventPathClass) Alloc() MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREventPathClass) New() MTREventPath {
	rv := objc.Send[MTREventPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventPath) Init() MTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventPath) Autorelease() MTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventPath creates a new MTREventPath instance.
func NewMTREventPath() MTREventPath {
	return getMTREventPathClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventpath/event
func (m_ MTREventPath) Event() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("event"))
	return rv
}


// SetEvent sets the value of the event property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventpath/event
func (m_ MTREventPath) SetEvent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEvent:"), value)
}



