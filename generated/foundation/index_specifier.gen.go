// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IndexSpecifier] class.
var (
	indexSpecifierClass     _IndexSpecifierClass
	indexSpecifierClassOnce sync.Once
)

func getIndexSpecifierClass() _IndexSpecifierClass {
	indexSpecifierClassOnce.Do(func() {
		indexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}
	})
	return indexSpecifierClass
}

type _IndexSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [IndexSpecifier] class.
type IIndexSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier representing an object in a collection (or container) with an index number.
//
// The script terms and specify the object with index , while specifies the object with index of . A negative index indicates a location by counting backward from the last object in the collection. You don’t normally subclass .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier
type IndexSpecifier struct {
	ScriptObjectSpecifier
}

// IndexSpecifierFrom constructs a [IndexSpecifier] from an unsafe.Pointer.
//
// A specifier representing an object in a collection (or container) with an index number.
func IndexSpecifierFrom(ptr unsafe.Pointer) IndexSpecifier {
	return IndexSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IndexSpecifierClass) Alloc() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndexSpecifierClass) New() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexSpecifier) Init() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexSpecifier) Autorelease() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexSpecifier creates a new IndexSpecifier instance.
func NewIndexSpecifier() IndexSpecifier {
	return getIndexSpecifierClass().New()
}




