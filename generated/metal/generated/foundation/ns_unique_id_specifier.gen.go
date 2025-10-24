// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var (
	UniqueIDSpecifierClass     _UniqueIDSpecifierClass
	UniqueIDSpecifierClassOnce sync.Once
)

func getUniqueIDSpecifierClass() _UniqueIDSpecifierClass {
	UniqueIDSpecifierClassOnce.Do(func() {
		UniqueIDSpecifierClass = _UniqueIDSpecifierClass{objc.GetClass("NSUniqueIDSpecifier")}
	})
	return UniqueIDSpecifierClass
}

type _UniqueIDSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [UniqueIDSpecifier] class.
type IUniqueIDSpecifier interface {
	IScriptObjectSpecifier
	// properties:
	UniqueID() unsafe.Pointer
	SetUniqueID(value unsafe.Pointer)
	// methods:
}

// A specifier for an object in a collection (or container) by unique ID.
//
// This specifier works only for objects that have an ID property. The unique ID object passed to an instance of must be either an object or an object. The exact type should match the scripting dictionary declaration of the ID attribute for the relevant scripting class. You can expect that the ID property will be for any object that supports it. Therefore a scripter can obtain the unique ID for an object and refer to the object by the ID, but cannot set the unique ID. You don’t normally subclass . The evaluation of objects follows these steps until the specified object is found: If the container implements a method whose selector matches the relevant pattern established by scripting key-value coding, the method is invoked. This method can potentially be very fast, and it may be relatively easy to implement. As is the case when evaluating any script object specifier, the container of the specified object is given a chance to evaluate the object specifier. If the container class implements the method, the method is invoked. This method can potentially be very fast, but it is relatively difficult to implement. An object that specifies the first object whose relevant attribute matches the ID is synthesized and evaluated. The object must search through all of the keyed elements in the container, looking for a match. The search is potentially very slow.


// A specifier for an object in a collection (or container) by unique ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier
type UniqueIDSpecifier struct {
	ScriptObjectSpecifier
}

// UniqueIDSpecifierFrom constructs a [UniqueIDSpecifier] from an unsafe.Pointer.
//
// A specifier for an object in a collection (or container) by unique ID.
func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UniqueIDSpecifierClass) Alloc() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UniqueIDSpecifierClass) New() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UniqueIDSpecifier) Init() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UniqueIDSpecifier) Autorelease() UniqueIDSpecifier {
	rv := objc.Send[UniqueIDSpecifier](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUniqueIDSpecifier creates a new UniqueIDSpecifier instance.
func NewUniqueIDSpecifier() UniqueIDSpecifier {
	return getUniqueIDSpecifierClass().New()
}



// Returns the ID encapsulated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier/uniqueid
func (u_ UniqueIDSpecifier) UniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uniqueID"))
	return rv
}


// Returns the ID encapsulated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuniqueidspecifier/uniqueid
func (u_ UniqueIDSpecifier) SetUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUniqueID:"), value)
}



