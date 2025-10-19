// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleScript] class.
var (
	appleScriptClass     _AppleScriptClass
	appleScriptClassOnce sync.Once
)

func getAppleScriptClass() _AppleScriptClass {
	appleScriptClassOnce.Do(func() {
		appleScriptClass = _AppleScriptClass{objc.GetClass("NSAppleScript")}
	})
	return appleScriptClass
}

type _AppleScriptClass struct {
	class objc.Class
}

// An interface definition for the [AppleScript] class.
type IAppleScript interface {
	objectivec.IObject
}

// An object that provides the ability to load, compile, and execute scripts.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript
type AppleScript struct {
	objectivec.Object
}

// AppleScriptFrom constructs a [AppleScript] from an unsafe.Pointer.
//
// An object that provides the ability to load, compile, and execute scripts.
func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleScriptClass) Alloc() AppleScript {
	rv := objc.Send[AppleScript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleScriptClass) New() AppleScript {
	rv := objc.Send[AppleScript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleScript) Init() AppleScript {
	rv := objc.Send[AppleScript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleScript) Autorelease() AppleScript {
	rv := objc.Send[AppleScript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleScript creates a new AppleScript instance.
func NewAppleScript() AppleScript {
	return getAppleScriptClass().New()
}




