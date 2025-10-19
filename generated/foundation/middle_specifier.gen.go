// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MiddleSpecifier] class.
var middleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}

type _MiddleSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [MiddleSpecifier] class.
type IMiddleSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMiddleSpecifier

type MiddleSpecifier struct {
	ScriptObjectSpecifier
}

// MiddleSpecifierFrom constructs a [MiddleSpecifier] from an unsafe.Pointer.
//
// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object.
func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MiddleSpecifierClass) Alloc() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MiddleSpecifierClass) New() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MiddleSpecifier) Init() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MiddleSpecifier) Autorelease() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMiddleSpecifier creates a new MiddleSpecifier instance.
func NewMiddleSpecifier() MiddleSpecifier {
	return middleSpecifierClass.New()
}




