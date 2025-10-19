// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UIResponder] class.
var (
	uIResponderClass     _UIResponderClass
	uIResponderClassOnce sync.Once
)

func getUIResponderClass() _UIResponderClass {
	uIResponderClassOnce.Do(func() {
		uIResponderClass = _UIResponderClass{objc.GetClass("UIResponder")}
	})
	return uIResponderClass
}

type _UIResponderClass struct {
	class objc.Class
}

// An interface definition for the [UIResponder] class.
type IUIResponder interface {
	objectivec.IObject
}

// A parent class referenced by other SpriteKit classes. [Full Topic]
type UIResponder struct {
	objectivec.Object
}

// UIResponderFrom constructs a [UIResponder] from an unsafe.Pointer.
//
// A parent class referenced by other SpriteKit classes.
func UIResponderFrom(ptr unsafe.Pointer) UIResponder {
	return UIResponder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UIResponderClass) Alloc() UIResponder {
	rv := objc.Send[UIResponder](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UIResponderClass) New() UIResponder {
	rv := objc.Send[UIResponder](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UIResponder) Init() UIResponder {
	rv := objc.Send[UIResponder](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UIResponder) Autorelease() UIResponder {
	rv := objc.Send[UIResponder](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUIResponder creates a new UIResponder instance.
func NewUIResponder() UIResponder {
	return getUIResponderClass().New()
}




