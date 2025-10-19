// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PositionalSpecifier] class.
var (
	positionalSpecifierClass     _PositionalSpecifierClass
	positionalSpecifierClassOnce sync.Once
)

func getPositionalSpecifierClass() _PositionalSpecifierClass {
	positionalSpecifierClassOnce.Do(func() {
		positionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}
	})
	return positionalSpecifierClass
}

type _PositionalSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objectivec.IObject
}

// A specifier for an insertion point in a container relative to another object in the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier
type PositionalSpecifier struct {
	objectivec.Object
}

// PositionalSpecifierFrom constructs a [PositionalSpecifier] from an unsafe.Pointer.
//
// A specifier for an insertion point in a container relative to another object in the container.
func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PositionalSpecifierClass) Alloc() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PositionalSpecifierClass) New() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PositionalSpecifier) Init() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PositionalSpecifier) Autorelease() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPositionalSpecifier creates a new PositionalSpecifier instance.
func NewPositionalSpecifier() PositionalSpecifier {
	return getPositionalSpecifierClass().New()
}




