// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAttributeCacheContainer] class.
var (
	MTRAttributeCacheContainerClass     _MTRAttributeCacheContainerClass
	MTRAttributeCacheContainerClassOnce sync.Once
)

func getMTRAttributeCacheContainerClass() _MTRAttributeCacheContainerClass {
	MTRAttributeCacheContainerClassOnce.Do(func() {
		MTRAttributeCacheContainerClass = _MTRAttributeCacheContainerClass{objc.GetClass("MTRAttributeCacheContainer")}
	})
	return MTRAttributeCacheContainerClass
}

type _MTRAttributeCacheContainerClass struct {
	class objc.Class
}

// An interface definition for the [MTRAttributeCacheContainer] class.
type IMTRAttributeCacheContainer interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeCacheContainer
type MTRAttributeCacheContainer struct {
	objectivec.Object
}

// MTRAttributeCacheContainerFrom constructs a [MTRAttributeCacheContainer] from an unsafe.Pointer.
func MTRAttributeCacheContainerFrom(ptr unsafe.Pointer) MTRAttributeCacheContainer {
	return MTRAttributeCacheContainer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeCacheContainerClass) Alloc() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAttributeCacheContainerClass) New() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeCacheContainer) Init() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeCacheContainer) Autorelease() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeCacheContainer creates a new MTRAttributeCacheContainer instance.
func NewMTRAttributeCacheContainer() MTRAttributeCacheContainer {
	return getMTRAttributeCacheContainerClass().New()
}




