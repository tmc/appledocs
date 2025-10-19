// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var (
	uniqueIDSpecifierClass     _UniqueIDSpecifierClass
	uniqueIDSpecifierClassOnce sync.Once
)

func getUniqueIDSpecifierClass() _UniqueIDSpecifierClass {
	uniqueIDSpecifierClassOnce.Do(func() {
		uniqueIDSpecifierClass = _UniqueIDSpecifierClass{objc.GetClass("NSUniqueIDSpecifier")}
	})
	return uniqueIDSpecifierClass
}

type _UniqueIDSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [UniqueIDSpecifier] class.
type IUniqueIDSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for an object in a collection (or container) by unique ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier
type UniqueIDSpecifier struct {
	ScriptObjectSpecifier
}

// UniqueIDSpecifierFrom constructs a [UniqueIDSpecifier] from an unsafe.Pointer.
//
// A specifier for an object in a collection (or container) by unique ID.
func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UniqueIDSpecifierClass) Alloc() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UniqueIDSpecifierClass) New() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UniqueIDSpecifier) Init() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UniqueIDSpecifier) Autorelease() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUniqueIDSpecifier creates a new UniqueIDSpecifier instance.
func NewUniqueIDSpecifier() UniqueIDSpecifier {
	return getUniqueIDSpecifierClass().New()
}




