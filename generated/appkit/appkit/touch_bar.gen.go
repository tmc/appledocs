// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TouchBar] class.
var TouchBarClass objc.Class

func init() {
	TouchBarClass = objc.GetClass("NSTouchBar")
}

type TouchBar struct {
	objc.ID
}

func TouchBarFrom(ptr unsafe.Pointer) TouchBar {
	return TouchBar{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc TouchBar) Alloc() TouchBar {
	ret := objc.ID(TouchBarClass).Send(objc.RegisterName("alloc"))
	return TouchBar{ret}
}

// Init initializes the instance.
func (t_ TouchBar) Init() TouchBar {
	ret := t_.ID.Send(objc.RegisterName("init"))
	return TouchBar{ret}
}
// Creates a Touch Bar object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/init()
func NewTouchBar() TouchBar {
	instance := TouchBar{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/init(coder:)
func NewTouchBarWithCoder(coder unsafe.Pointer) TouchBar {
	instance := TouchBar{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = TouchBar{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns the Touch Bar item that corresponds to a given identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/item(forIdentifier:)
func (t_ TouchBar) ItemForIdentifier(identifier unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("itemForIdentifier:")
	ret := t_.ID.Send(sel, identifier)
	return unsafe.Pointer(ret)
}

