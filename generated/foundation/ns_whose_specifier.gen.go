// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WhoseSpecifier] class.
var (
	WhoseSpecifierClass     _WhoseSpecifierClass
	WhoseSpecifierClassOnce sync.Once
)

func getWhoseSpecifierClass() _WhoseSpecifierClass {
	WhoseSpecifierClassOnce.Do(func() {
		WhoseSpecifierClass = _WhoseSpecifierClass{objc.GetClass("NSWhoseSpecifier")}
	})
	return WhoseSpecifierClass
}

type _WhoseSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [WhoseSpecifier] class.
type IWhoseSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier that indicates every object in a collection matching a condition.
//
// specifies every object in a collection (or every element in a container) that matches the condition defined by a single Boolean expression or multiple Boolean expressions connected by logical operators. is unique among object specifiers in that its top-level container is typically not the application object but an evaluated object specifier involved in the tested-for condition. An object encapsulates a “test” object for defining this condition. A test object is instantiated from a subclass of the abstract class, whose one declared method is . See “Boolean Expressions and Logical Operations” in and the descriptions in NSComparisonMethods and NSScriptingComparisonMethods for more information. The set of elements specified by an object can be a subset of those that pass the object’s test. This subset is specified by the various sub-element properties of the object . Consider as an example the specifier . This would be represented by an object that uses a test specifier and another object specifier to identify a subset of the objects with the specified property. That is, the specifier’s property is ; the test specifier is an index specifier with property and ; and the qualifier is a key value qualifier for key and value . The test object specifier ( ) is evaluated for each object (paragraph) using that object as the container; the resulting objects (if any) are tested with the qualifier ( ). is part of Cocoa’s built-in script handling. You don’t normally subclass it.
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
