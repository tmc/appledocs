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
// is the abstract superclass for classes that instantiate objects called “object specifiers.” An object specifier represents an AppleScript reference form, which is a natural-language expression such as or or . The scripting system maps these words or phrases to attributes and relationships of scriptable objects. A reference form rarely occurs in isolation; usually a script statement consists of a series of reference forms preceded by a command and typically connected to each other by , such as: The expression specifies a location in the application’s AppleScript object model—the objects the application makes available to scripters. The classes of objects in the object model often closely match the classes of actual objects in the application, but they are not required to. An object specifier locates objects in the running application that correspond to the specified object model objects. Your application typically creates object specifiers when it implements the method for its scriptable classes. That method is defined by the NSScriptObjectSpecifiers protocol. It is unlikely that you would ever need to create your own subclass of ; the set of valid AppleScript reference forms is determined by Apple Computer and object specifier classes are already implemented for this set. If for some reason you do need to create a subclass, you must override the primitive method to return indices to the elements within the container whose values are matched with the child specifier’s key. In addition, you probably need to declare any special instance variables and implement an initializer that invokes super’s designated initializer, , and initializes these variables. For a comprehensive treatment of object specifiers, including sample code, see in .
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




