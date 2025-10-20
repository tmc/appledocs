// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [symbolVersion] class.
var (
	SymbolVersionClass     _symbolVersionClass
	SymbolVersionClassOnce sync.Once
)

func getsymbolVersionClass() _symbolVersionClass {
	SymbolVersionClassOnce.Do(func() {
		SymbolVersionClass = _symbolVersionClass{objc.GetClass("symbolVersion")}
	})
	return SymbolVersionClass
}

type _symbolVersionClass struct {
	class objc.Class
}

// An interface definition for the [symbolVersion] class.
type IsymbolVersion interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-c.ivar
type symbolVersion struct {
	objectivec.Object
}

// symbolVersionFrom constructs a [symbolVersion] from an unsafe.Pointer.
func symbolVersionFrom(ptr unsafe.Pointer) symbolVersion {
	return symbolVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _symbolVersionClass) Alloc() symbolVersion {
	rv := objc.Send[symbolVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _symbolVersionClass) New() symbolVersion {
	rv := objc.Send[symbolVersion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ symbolVersion) Init() symbolVersion {
	rv := objc.Send[symbolVersion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ symbolVersion) Autorelease() symbolVersion {
	rv := objc.Send[symbolVersion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsymbolVersion creates a new symbolVersion instance.
func NewsymbolVersion() symbolVersion {
	return getsymbolVersionClass().New()
}




