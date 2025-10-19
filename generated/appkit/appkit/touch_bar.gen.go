// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TouchBar] class.
var (
	touchBarClass     _TouchBarClass
	touchBarClassOnce sync.Once
)

func getTouchBarClass() _TouchBarClass {
	touchBarClassOnce.Do(func() {
		touchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}
	})
	return touchBarClass
}

type _TouchBarClass struct {
	class objc.Class
}

// An interface definition for the [TouchBar] class.
type ITouchBar interface {
	objectivec.IObject
	ItemForIdentifier(identifier unsafe.Pointer) unsafe.Pointer
}

// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar

type TouchBar struct {
	objectivec.Object
}

// TouchBarFrom constructs a [TouchBar] from an unsafe.Pointer.
//
// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro.
func TouchBarFrom(ptr unsafe.Pointer) TouchBar {
	return TouchBar{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TouchBarClass) New() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TouchBar) Init() TouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TouchBar) Autorelease() TouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouchBar creates a new TouchBar instance.
func NewTouchBar() TouchBar {
	return getTouchBarClass().New()
}


// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/init(coder:)
func NewTouchBarWithCoder(coder unsafe.Pointer) TouchBar {
	instance := getTouchBarClass().Alloc()
	rv := objc.Send[TouchBar](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns the Touch Bar item that corresponds to a given identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/item(forIdentifier:)
func (t_ TouchBar) ItemForIdentifier(identifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemForIdentifier:"), identifier)
	return rv
}

