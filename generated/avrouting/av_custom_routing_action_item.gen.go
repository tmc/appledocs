// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CustomRoutingActionItem] class.
var (
	CustomRoutingActionItemClass     _CustomRoutingActionItemClass
	CustomRoutingActionItemClassOnce sync.Once
)

func getCustomRoutingActionItemClass() _CustomRoutingActionItemClass {
	CustomRoutingActionItemClassOnce.Do(func() {
		CustomRoutingActionItemClass = _CustomRoutingActionItemClass{objc.GetClass("AVCustomRoutingActionItem")}
	})
	return CustomRoutingActionItemClass
}

type _CustomRoutingActionItemClass struct {
	class objc.Class
}

// An interface definition for the [CustomRoutingActionItem] class.
type ICustomRoutingActionItem interface {
	objectivec.IObject
	OverrideTitle() string
	SetOverrideTitle(value string)
	Type() objectivec.IObject
	SetType(value objectivec.IObject)
}

// An object that represents a custom action item to display in a device route picker.
//
// Use this class to specify supplemental action items to display in the list of discovered routes. Tapping a custom item dismisses the picker and calls the method of .


// An object that represents a custom action item to display in a device route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem
type CustomRoutingActionItem struct {
	objectivec.Object
}

// CustomRoutingActionItemFrom constructs a [CustomRoutingActionItem] from an unsafe.Pointer.
//
// An object that represents a custom action item to display in a device route picker.
func CustomRoutingActionItemFrom(ptr unsafe.Pointer) CustomRoutingActionItem {
	return CustomRoutingActionItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingActionItemClass) Alloc() CustomRoutingActionItem {
	rv := objc.Send[CustomRoutingActionItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomRoutingActionItemClass) New() CustomRoutingActionItem {
	rv := objc.Send[CustomRoutingActionItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingActionItem) Init() CustomRoutingActionItem {
	rv := objc.Send[CustomRoutingActionItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingActionItem) Autorelease() CustomRoutingActionItem {
	rv := objc.Send[CustomRoutingActionItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingActionItem creates a new CustomRoutingActionItem instance.
func NewCustomRoutingActionItem() CustomRoutingActionItem {
	return getCustomRoutingActionItemClass().New()
}



// A string to use to override the title of the item’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/overrideTitle
func (c_ CustomRoutingActionItem) OverrideTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("overrideTitle"))
	return rv
}


// A string to use to override the title of the item’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/overrideTitle
func (c_ CustomRoutingActionItem) SetOverrideTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOverrideTitle:"), objc.String(value))
}


// A type with an identifier that matches a value in the app’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/type
func (c_ CustomRoutingActionItem) Type() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("type"))
	return rv
}


// A type with an identifier that matches a value in the app’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/type
func (c_ CustomRoutingActionItem) SetType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}



