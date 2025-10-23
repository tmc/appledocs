// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NameSpecifier] class.
var (
	NameSpecifierClass     _NameSpecifierClass
	NameSpecifierClassOnce sync.Once
)

func getNameSpecifierClass() _NameSpecifierClass {
	NameSpecifierClassOnce.Do(func() {
		NameSpecifierClass = _NameSpecifierClass{objc.GetClass("NSNameSpecifier")}
	})
	return NameSpecifierClass
}

type _NameSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [NameSpecifier] class.
type INameSpecifier interface {
	IScriptObjectSpecifier
	Name() string
	SetName(value string)
}

// A specifier for an object in a collection (or container) by name.
//
// As an example, the following script specifies both an application and a window by name. In this script, the named window’s implicitly specified container is the Finder application’s list of open windows. This specifier works only for objects that have a name property. You don’t normally subclass . The evaluation of an instance of follows these steps until the specified object is found: If the container implements a method whose selector matches the relevant pattern established by scripting key-value coding, the method is invoked. This method can potentially be very fast, and it may be relatively easy to implement. As is the case when evaluating any script object specifier, the container of the specified object is given a chance to evaluate the object specifier. If the container class implements the method, the method is invoked. This method can potentially be very fast, but it is relatively difficult to implement. An instance of that specifies the first object whose relevant attribute matches the name is synthesized and evaluated. The instance of must search through all of the keyed elements in the container, looking for a match. The search is potentially very slow.


// A specifier for an object in a collection (or container) by name.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getNameSpecifierClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/init(coder:)
func NewNameSpecifierWithCoder(inCoder ICoder) NameSpecifier {
	instance := getNameSpecifierClass().Alloc()
	rv := objc.Send[NameSpecifier](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}


// Invokes the super class’s method and then sets the name instance variable to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/init(containerClassDescription:containerSpecifier:key:name:)
func NewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string, name string) NameSpecifier {
	instance := getNameSpecifierClass().Alloc()
	rv := objc.Send[NameSpecifier](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:name:"), classDesc, container, objc.String(property), objc.String(name))
	rv.Autorelease()
	return rv
}



// Sets the name encapsulated with the receiver for the specified object in the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/name
func (n_ NameSpecifier) Name() string {
	rv := objc.Send[string](n_.ID, objc.Sel("name"))
	return rv
}


// Sets the name encapsulated with the receiver for the specified object in the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier/name
func (n_ NameSpecifier) SetName(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setName:"), objc.String(value))
}


