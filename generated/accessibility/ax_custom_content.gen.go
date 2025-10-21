// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXCustomContent] class.
var (
	AXCustomContentClass     _AXCustomContentClass
	AXCustomContentClassOnce sync.Once
)

func getAXCustomContentClass() _AXCustomContentClass {
	AXCustomContentClassOnce.Do(func() {
		AXCustomContentClass = _AXCustomContentClass{objc.GetClass("AXCustomContent")}
	})
	return AXCustomContentClass
}

type _AXCustomContentClass struct {
	class objc.Class
}

// An interface definition for the [AXCustomContent] class.
type IAXCustomContent interface {
	objectivec.IObject
}

// Objects that define custom content and the timing of its output.
//
// An object contains the accessibility strings for the labels you apply to your accessibility content. Combine them with the protocol to allow your users to experience the content in a more appropriate manner for each assistive technology.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent
type AXCustomContent struct {
	objectivec.Object
}

// AXCustomContentFrom constructs a [AXCustomContent] from an unsafe.Pointer.
//
// Objects that define custom content and the timing of its output.
func AXCustomContentFrom(ptr unsafe.Pointer) AXCustomContent {
	return AXCustomContent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXCustomContentClass) Alloc() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXCustomContentClass) New() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXCustomContent) Init() AXCustomContent {
	rv := objc.Send[AXCustomContent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXCustomContent) Autorelease() AXCustomContent {
	rv := objc.Send[AXCustomContent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXCustomContent creates a new AXCustomContent instance.
func NewAXCustomContent() AXCustomContent {
	return getAXCustomContentClass().New()
}




