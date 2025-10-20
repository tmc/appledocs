// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [manInTheMiddleProtection] class.
var (
	ManInTheMiddleProtectionClass     _manInTheMiddleProtectionClass
	ManInTheMiddleProtectionClassOnce sync.Once
)

func getmanInTheMiddleProtectionClass() _manInTheMiddleProtectionClass {
	ManInTheMiddleProtectionClassOnce.Do(func() {
		ManInTheMiddleProtectionClass = _manInTheMiddleProtectionClass{objc.GetClass("manInTheMiddleProtection")}
	})
	return ManInTheMiddleProtectionClass
}

type _manInTheMiddleProtectionClass struct {
	class objc.Class
}

// An interface definition for the [manInTheMiddleProtection] class.
type ImanInTheMiddleProtection interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-c.ivar
type manInTheMiddleProtection struct {
	objectivec.Object
}

// manInTheMiddleProtectionFrom constructs a [manInTheMiddleProtection] from an unsafe.Pointer.
func manInTheMiddleProtectionFrom(ptr unsafe.Pointer) manInTheMiddleProtection {
	return manInTheMiddleProtection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _manInTheMiddleProtectionClass) Alloc() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _manInTheMiddleProtectionClass) New() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ manInTheMiddleProtection) Init() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ manInTheMiddleProtection) Autorelease() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmanInTheMiddleProtection creates a new manInTheMiddleProtection instance.
func NewmanInTheMiddleProtection() manInTheMiddleProtection {
	return getmanInTheMiddleProtectionClass().New()
}




