// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NameSpecifier] class.
var nameSpecifierClass = _NameSpecifierClass{objc.GetClass("NSNameSpecifier")}

type _NameSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [NameSpecifier] class.
type INameSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for an object in a collection (or container) by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier

type NameSpecifier struct {
	ScriptObjectSpecifier
}

// NameSpecifierFrom constructs a [NameSpecifier] from an unsafe.Pointer.
//
// A specifier for an object in a collection (or container) by name.
func NameSpecifierFrom(ptr unsafe.Pointer) NameSpecifier {
	return NameSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (nc _NameSpecifierClass) Alloc() NameSpecifier {
	rv := objc.Send[NameSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NameSpecifierClass) New() NameSpecifier {
	rv := objc.Send[NameSpecifier](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NameSpecifier) Init() NameSpecifier {
	rv := objc.Send[NameSpecifier](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NameSpecifier) Autorelease() NameSpecifier {
	rv := objc.Send[NameSpecifier](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNameSpecifier creates a new NameSpecifier instance.
func NewNameSpecifier() NameSpecifier {
	return nameSpecifierClass.New()
}




