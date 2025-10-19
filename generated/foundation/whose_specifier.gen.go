// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WhoseSpecifier] class.
var (
	whoseSpecifierClass     _WhoseSpecifierClass
	whoseSpecifierClassOnce sync.Once
)

func getWhoseSpecifierClass() _WhoseSpecifierClass {
	whoseSpecifierClassOnce.Do(func() {
		whoseSpecifierClass = _WhoseSpecifierClass{objc.GetClass("NSWhoseSpecifier")}
	})
	return whoseSpecifierClass
}

type _WhoseSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [WhoseSpecifier] class.
type IWhoseSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier that indicates every object in a collection matching a condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier
type WhoseSpecifier struct {
	ScriptObjectSpecifier
}

// WhoseSpecifierFrom constructs a [WhoseSpecifier] from an unsafe.Pointer.
//
// A specifier that indicates every object in a collection matching a condition.
func WhoseSpecifierFrom(ptr unsafe.Pointer) WhoseSpecifier {
	return WhoseSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (wc _WhoseSpecifierClass) Alloc() WhoseSpecifier {
	rv := objc.Send[WhoseSpecifier](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WhoseSpecifierClass) New() WhoseSpecifier {
	rv := objc.Send[WhoseSpecifier](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WhoseSpecifier) Init() WhoseSpecifier {
	rv := objc.Send[WhoseSpecifier](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WhoseSpecifier) Autorelease() WhoseSpecifier {
	rv := objc.Send[WhoseSpecifier](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWhoseSpecifier creates a new WhoseSpecifier instance.
func NewWhoseSpecifier() WhoseSpecifier {
	return getWhoseSpecifierClass().New()
}




