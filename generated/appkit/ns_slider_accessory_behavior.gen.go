// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSliderAccessoryBehavior */


/* debug [class_header]: Header for NSSliderAccessoryBehavior */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SliderAccessoryBehavior */
// An interface definition for the [SliderAccessoryBehavior] class.
type ISliderAccessoryBehavior interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SliderAccessoryBehavior */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SliderAccessoryBehavior */
	// methods:
	HandleAction(sender ISliderAccessory)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SliderAccessoryBehavior */
// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryBehaviorClass) Alloc() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SliderAccessoryBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior
type SliderAccessoryBehavior struct {
	objectivec.Object
}

// SliderAccessoryBehaviorFrom constructs a [SliderAccessoryBehavior] from an unsafe.Pointer.
func SliderAccessoryBehaviorFrom(ptr unsafe.Pointer) SliderAccessoryBehavior {
	return SliderAccessoryBehavior{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SliderAccessoryBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(handler:)
func NewSliderAccessoryBehaviorWithHandler(handler unsafe.Pointer) SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(getSliderAccessoryBehaviorClass().class), objc.Sel("behaviorWithHandler:"), handler)
	return rv
}/* debug [class_init_methods/constructor]: NewSliderAccessoryBehaviorWithHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(target:action:)
func NewSliderAccessoryBehaviorWithTargetAction(target objc.IObject, action objc.SEL) SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(getSliderAccessoryBehaviorClass().class), objc.Sel("behaviorWithTarget:action:"), target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewSliderAccessoryBehaviorWithTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SliderAccessoryBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(handler:)
func (sc _SliderAccessoryBehaviorClass) BehaviorWithHandler(handler unsafe.Pointer) ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("behaviorWithHandler:"), handler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(target:action:)
func (sc _SliderAccessoryBehaviorClass) BehaviorWithTargetAction(target objc.IObject, action objc.SEL) ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("behaviorWithTarget:action:"), target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BehaviorWithTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SliderAccessoryBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/automatic
func (sc _SliderAccessoryBehaviorClass) AutomaticBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("automaticBehavior"))
	return rv
}/* debug [class_properties_class/property]: automaticBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueReset
func (sc _SliderAccessoryBehaviorClass) ValueResetBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("valueResetBehavior"))
	return rv
}/* debug [class_properties_class/property]: valueResetBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueStep
func (sc _SliderAccessoryBehaviorClass) ValueStepBehavior() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("valueStepBehavior"))
	return rv
}/* debug [class_properties_class/property]: valueStepBehavior */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SliderAccessoryBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/handleAction(_:)
func (s_ SliderAccessoryBehavior) HandleAction(sender ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("handleAction:"), sender)
}/* debug [instance_methods/method]: HandleAction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SliderAccessoryBehavior */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/automatic
func (s_ SliderAccessoryBehavior) AutomaticBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("automaticBehavior"))
	return rv
}/* debug [instance_properties/getter]: automaticBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueReset
func (s_ SliderAccessoryBehavior) ValueResetBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("valueResetBehavior"))
	return rv
}/* debug [instance_properties/getter]: valueResetBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueStep
func (s_ SliderAccessoryBehavior) ValueStepBehavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("valueStepBehavior"))
	return rv
}/* debug [instance_properties/getter]: valueStepBehavior */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSliderAccessoryBehavior */


