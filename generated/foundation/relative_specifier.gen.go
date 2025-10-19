// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelativeSpecifier] class.
var (
	relativeSpecifierClass     _RelativeSpecifierClass
	relativeSpecifierClassOnce sync.Once
)

func getRelativeSpecifierClass() _RelativeSpecifierClass {
	relativeSpecifierClassOnce.Do(func() {
		relativeSpecifierClass = _RelativeSpecifierClass{objc.GetClass("NSRelativeSpecifier")}
	})
	return relativeSpecifierClass
}

type _RelativeSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [RelativeSpecifier] class.
type IRelativeSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier that indicates an object in a collection by its position relative to another object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier
type RelativeSpecifier struct {
	ScriptObjectSpecifier
}

// RelativeSpecifierFrom constructs a [RelativeSpecifier] from an unsafe.Pointer.
//
// A specifier that indicates an object in a collection by its position relative to another object.
func RelativeSpecifierFrom(ptr unsafe.Pointer) RelativeSpecifier {
	return RelativeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RelativeSpecifierClass) Alloc() RelativeSpecifier {
	rv := objc.Send[RelativeSpecifier](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RelativeSpecifierClass) New() RelativeSpecifier {
	rv := objc.Send[RelativeSpecifier](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelativeSpecifier) Init() RelativeSpecifier {
	rv := objc.Send[RelativeSpecifier](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelativeSpecifier) Autorelease() RelativeSpecifier {
	rv := objc.Send[RelativeSpecifier](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelativeSpecifier creates a new RelativeSpecifier instance.
func NewRelativeSpecifier() RelativeSpecifier {
	return getRelativeSpecifierClass().New()
}




