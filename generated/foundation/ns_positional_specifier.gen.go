// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PositionalSpecifier] class.
var (
	PositionalSpecifierClass     _PositionalSpecifierClass
	PositionalSpecifierClassOnce sync.Once
)

func getPositionalSpecifierClass() _PositionalSpecifierClass {
	PositionalSpecifierClassOnce.Do(func() {
		PositionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}
	})
	return PositionalSpecifierClass
}

type _PositionalSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objectivec.IObject
	SetInsertionClassDescription(classDescription unsafe.Pointer)
}

// A specifier for an insertion point in a container relative to another object in the container.
//
// Instances of specify an insertion point in a container relative to another object in the container, for example, or . The container is specified by an instance of . objects commonly encapsulate object specifiers used as arguments to the ( ) and commands and indicate where the created or moved object is to be inserted relative to the object represented by an object specifier. Invoking an accessor method to obtain information about an instance of causes the object to be evaluated if it hasn’t been already. You don’t normally subclass .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier
type PositionalSpecifier struct {
	objectivec.Object
}

// PositionalSpecifierFrom constructs a [PositionalSpecifier] from an unsafe.Pointer.
//
// A specifier for an insertion point in a container relative to another object in the container.
func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PositionalSpecifierClass) Alloc() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PositionalSpecifierClass) New() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PositionalSpecifier) Init() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PositionalSpecifier) Autorelease() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPositionalSpecifier creates a new PositionalSpecifier instance.
func NewPositionalSpecifier() PositionalSpecifier {
	return getPositionalSpecifierClass().New()
}

// Sets the class description for the object or objects to be inserted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/setInsertionClassDescription(_:)
func (p_ PositionalSpecifier) SetInsertionClassDescription(classDescription unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertionClassDescription:"), classDescription)
}

// Returns the object specifier specified at initialization time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/objectSpecifier
func (p_ PositionalSpecifier) ObjectSpecifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectSpecifier"))
	return rv
}

// Returns the insertion position specified at initialization time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/position
func (p_ PositionalSpecifier) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("position"))
	return rv
}
