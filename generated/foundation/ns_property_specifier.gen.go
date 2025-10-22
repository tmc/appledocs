// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PropertySpecifier] class.
var (
	PropertySpecifierClass     _PropertySpecifierClass
	PropertySpecifierClassOnce sync.Once
)

func getPropertySpecifierClass() _PropertySpecifierClass {
	PropertySpecifierClassOnce.Do(func() {
		PropertySpecifierClass = _PropertySpecifierClass{objc.GetClass("NSPropertySpecifier")}
	})
	return PropertySpecifierClass
}

type _PropertySpecifierClass struct {
	class objc.Class
}

// An interface definition for the [PropertySpecifier] class.
type IPropertySpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
//
// You don’t typically subclass .


// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPropertySpecifier
type PropertySpecifier struct {
	ScriptObjectSpecifier
}

// PropertySpecifierFrom constructs a [PropertySpecifier] from an unsafe.Pointer.
//
// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PropertySpecifierClass) Alloc() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertySpecifierClass) New() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertySpecifier) Init() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertySpecifier) Autorelease() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertySpecifier creates a new PropertySpecifier instance.
func NewPropertySpecifier() PropertySpecifier {
	return getPropertySpecifierClass().New()
}




