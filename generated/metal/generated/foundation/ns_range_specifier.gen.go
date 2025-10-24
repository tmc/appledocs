// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RangeSpecifier] class.
var (
	RangeSpecifierClass     _RangeSpecifierClass
	RangeSpecifierClassOnce sync.Once
)

func getRangeSpecifierClass() _RangeSpecifierClass {
	RangeSpecifierClassOnce.Do(func() {
		RangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}
	})
	return RangeSpecifierClass
}

type _RangeSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [RangeSpecifier] class.
type IRangeSpecifier interface {
	IScriptObjectSpecifier
	// properties:
	EndSpecifier() IScriptObjectSpecifier
	SetEndSpecifier(value IScriptObjectSpecifier)
	StartSpecifier() IScriptObjectSpecifier
	SetStartSpecifier(value IScriptObjectSpecifier)
	// methods:
}

// A specifier for a range of objects in a container.
//
// An object specifies a range (that is, an uninterrupted series) of objects in a container through two delimiting objects. The range is represented by two object specifiers, a start specifier and an end specifier, which can be of any specifier type (such as or object). These specifiers are evaluated in the context of the same container object as the range specifier itself. You don’t normally subclass .


// A specifier for a range of objects in a container.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getRangeSpecifierClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(coder:)
func NewRangeSpecifierWithCoder(inCoder ICoder) RangeSpecifier {
	instance := getRangeSpecifierClass().Alloc()
	rv := objc.Send[RangeSpecifier](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}


// Returns a range specifier initialized with the given properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(containerClassDescription:containerSpecifier:key:start:end:)
func NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property IString, startSpec IScriptObjectSpecifier, endSpec IScriptObjectSpecifier) RangeSpecifier {
	instance := getRangeSpecifierClass().Alloc()
	rv := objc.Send[RangeSpecifier](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:startSpecifier:endSpecifier:"), classDesc, container, property, startSpec, endSpec)
	rv.Autorelease()
	return rv
}



// Sets the object specifier representing the last object of the range to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/endSpecifier
func (r_ RangeSpecifier) EndSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](r_.ID, objc.Sel("endSpecifier"))
	return rv
}


// Sets the object specifier representing the last object of the range to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/endSpecifier
func (r_ RangeSpecifier) SetEndSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndSpecifier:"), value)
}


// Returns the object specifier representing the first object of the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/startSpecifier
func (r_ RangeSpecifier) StartSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](r_.ID, objc.Sel("startSpecifier"))
	return rv
}


// Returns the object specifier representing the first object of the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/startSpecifier
func (r_ RangeSpecifier) SetStartSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartSpecifier:"), value)
}


