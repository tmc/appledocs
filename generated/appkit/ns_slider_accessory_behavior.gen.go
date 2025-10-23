// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SliderAccessoryBehavior] class.
var (
	SliderAccessoryBehaviorClass     _SliderAccessoryBehaviorClass
	SliderAccessoryBehaviorClassOnce sync.Once
)

func getSliderAccessoryBehaviorClass() _SliderAccessoryBehaviorClass {
	SliderAccessoryBehaviorClassOnce.Do(func() {
		SliderAccessoryBehaviorClass = _SliderAccessoryBehaviorClass{objc.GetClass("NSSliderAccessoryBehavior")}
	})
	return SliderAccessoryBehaviorClass
}

type _SliderAccessoryBehaviorClass struct {
	class objc.Class
}

// An interface definition for the [SliderAccessoryBehavior] class.
type ISliderAccessoryBehavior interface {
	objectivec.IObject
	// properties:
	// methods:
	HandleAction(sender ISliderAccessory)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior
type SliderAccessoryBehavior struct {
	objectivec.Object
}

// SliderAccessoryBehaviorFrom constructs a [SliderAccessoryBehavior] from an unsafe.Pointer.
func SliderAccessoryBehaviorFrom(ptr unsafe.Pointer) SliderAccessoryBehavior {
	return SliderAccessoryBehavior{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryBehaviorClass) Alloc() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderAccessoryBehaviorClass) New() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderAccessoryBehavior) Init() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderAccessoryBehavior) Autorelease() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderAccessoryBehavior creates a new SliderAccessoryBehavior instance.
func NewSliderAccessoryBehavior() SliderAccessoryBehavior {
	return getSliderAccessoryBehaviorClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(handler:)
func NewSliderAccessoryBehaviorWithHandler(handler unsafe.Pointer) SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(getSliderAccessoryBehaviorClass().class), objc.Sel("behaviorWithHandler:"), handler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(target:action:)
func NewSliderAccessoryBehaviorWithTargetAction(target objectivec.IObject, action objc.SEL) SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(getSliderAccessoryBehaviorClass().class), objc.Sel("behaviorWithTarget:action:"), target, action)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(handler:)
func (sc _SliderAccessoryBehaviorClass) BehaviorWithHandler(handler unsafe.Pointer) ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("behaviorWithHandler:"), handler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(target:action:)
func (sc _SliderAccessoryBehaviorClass) BehaviorWithTargetAction(target objectivec.IObject, action objc.SEL) ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("behaviorWithTarget:action:"), target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/automatic
func (sc _SliderAccessoryBehaviorClass) AutomaticBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("automaticBehavior"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueReset
func (sc _SliderAccessoryBehaviorClass) ValueResetBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("valueResetBehavior"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueStep
func (sc _SliderAccessoryBehaviorClass) ValueStepBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("valueStepBehavior"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/handleAction(_:)
func (s_ SliderAccessoryBehavior) HandleAction(sender ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("handleAction:"), sender)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/automatic
func (s_ SliderAccessoryBehavior) AutomaticBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("automaticBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueReset
func (s_ SliderAccessoryBehavior) ValueResetBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("valueResetBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueStep
func (s_ SliderAccessoryBehavior) ValueStepBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("valueStepBehavior"))
	return rv
}


