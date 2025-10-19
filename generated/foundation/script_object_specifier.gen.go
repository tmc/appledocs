// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptObjectSpecifier] class.
var (
	scriptObjectSpecifierClass     _ScriptObjectSpecifierClass
	scriptObjectSpecifierClassOnce sync.Once
)

func getScriptObjectSpecifierClass() _ScriptObjectSpecifierClass {
	scriptObjectSpecifierClassOnce.Do(func() {
		scriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}
	})
	return scriptObjectSpecifierClass
}

type _ScriptObjectSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [ScriptObjectSpecifier] class.
type IScriptObjectSpecifier interface {
	objectivec.IObject
}

// An abstract class used to represent natural language expressions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier
type ScriptObjectSpecifier struct {
	objectivec.Object
}

// ScriptObjectSpecifierFrom constructs a [ScriptObjectSpecifier] from an unsafe.Pointer.
//
// An abstract class used to represent natural language expressions.
func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptObjectSpecifierClass) Alloc() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptObjectSpecifierClass) New() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptObjectSpecifier) Autorelease() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptObjectSpecifier creates a new ScriptObjectSpecifier instance.
func NewScriptObjectSpecifier() ScriptObjectSpecifier {
	return getScriptObjectSpecifierClass().New()
}




