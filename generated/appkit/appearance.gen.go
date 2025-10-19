// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Appearance] class.
var (
	appearanceClass     _AppearanceClass
	appearanceClassOnce sync.Once
)

func getAppearanceClass() _AppearanceClass {
	appearanceClassOnce.Do(func() {
		appearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}
	})
	return appearanceClass
}

type _AppearanceClass struct {
	class objc.Class
}

// An interface definition for the [Appearance] class.
type IAppearance interface {
	objectivec.IObject
}

// An object that manages standard appearance attributes for UI elements in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearance
type Appearance struct {
	objectivec.Object
}

// AppearanceFrom constructs a [Appearance] from an unsafe.Pointer.
//
// An object that manages standard appearance attributes for UI elements in an app.
func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppearanceClass) New() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Appearance) Init() Appearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Appearance) Autorelease() Appearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppearance creates a new Appearance instance.
func NewAppearance() Appearance {
	return getAppearanceClass().New()
}




