// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RandomSpecifier] class.
var (
	RandomSpecifierClass     _RandomSpecifierClass
	RandomSpecifierClassOnce sync.Once
)

func getRandomSpecifierClass() _RandomSpecifierClass {
	RandomSpecifierClassOnce.Do(func() {
		RandomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}
	})
	return RandomSpecifierClass
}

type _RandomSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [RandomSpecifier] class.
type IRandomSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRandomSpecifier
type RandomSpecifier struct {
	ScriptObjectSpecifier
}

// RandomSpecifierFrom constructs a [RandomSpecifier] from an unsafe.Pointer.
//
// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.
func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RandomSpecifierClass) Alloc() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RandomSpecifierClass) New() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RandomSpecifier) Init() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RandomSpecifier) Autorelease() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRandomSpecifier creates a new RandomSpecifier instance.
func NewRandomSpecifier() RandomSpecifier {
	return getRandomSpecifierClass().New()
}




