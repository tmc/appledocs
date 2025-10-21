// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTAHeader] class.
var (
	MTROTAHeaderClass     _MTROTAHeaderClass
	MTROTAHeaderClassOnce sync.Once
)

func getMTROTAHeaderClass() _MTROTAHeaderClass {
	MTROTAHeaderClassOnce.Do(func() {
		MTROTAHeaderClass = _MTROTAHeaderClass{objc.GetClass("MTROTAHeader")}
	})
	return MTROTAHeaderClass
}

type _MTROTAHeaderClass struct {
	class objc.Class
}

// An interface definition for the [MTROTAHeader] class.
type IMTROTAHeader interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeader
type MTROTAHeader struct {
	objectivec.Object
}

// MTROTAHeaderFrom constructs a [MTROTAHeader] from an unsafe.Pointer.
func MTROTAHeaderFrom(ptr unsafe.Pointer) MTROTAHeader {
	return MTROTAHeader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTAHeaderClass) Alloc() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTAHeaderClass) New() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTAHeader) Init() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTAHeader) Autorelease() MTROTAHeader {
	rv := objc.Send[MTROTAHeader](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTAHeader creates a new MTROTAHeader instance.
func NewMTROTAHeader() MTROTAHeader {
	return getMTROTAHeaderClass().New()
}




