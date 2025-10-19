// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RangeSpecifier] class.
var rangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}

type _RangeSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [RangeSpecifier] class.
type IRangeSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for a range of objects in a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier

type RangeSpecifier struct {
	ScriptObjectSpecifier
}

// RangeSpecifierFrom constructs a [RangeSpecifier] from an unsafe.Pointer.
//
// A specifier for a range of objects in a container.
func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (rc _RangeSpecifierClass) Alloc() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RangeSpecifierClass) New() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RangeSpecifier) Init() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RangeSpecifier) Autorelease() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRangeSpecifier creates a new RangeSpecifier instance.
func NewRangeSpecifier() RangeSpecifier {
	return rangeSpecifierClass.New()
}




