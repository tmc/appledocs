// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSliderAccessory */


/* debug [class_header]: Header for NSSliderAccessory */
// The class instance for the [SliderAccessory] class.
var (
	SliderAccessoryClass     _SliderAccessoryClass
	SliderAccessoryClassOnce sync.Once
)

func getSliderAccessoryClass() _SliderAccessoryClass {
	SliderAccessoryClassOnce.Do(func() {
		SliderAccessoryClass = _SliderAccessoryClass{objc.GetClass("NSSliderAccessory")}
	})
	return SliderAccessoryClass
}

type _SliderAccessoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SliderAccessory */
// An interface definition for the [SliderAccessory] class.
type ISliderAccessory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SliderAccessory */
	// properties:
	Behavior() ISliderAccessoryBehavior
	SetBehavior(value ISliderAccessoryBehavior)
	Enabled() bool
	SetEnabled(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SliderAccessory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SliderAccessory */
// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryClass) Alloc() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SliderAccessoryClass) New() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderAccessory) Init() SliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderAccessory) Autorelease() SliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderAccessory creates a new SliderAccessory instance.
func NewSliderAccessory() SliderAccessory {
	return getSliderAccessoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SliderAccessory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type SliderAccessory struct {
	objectivec.Object
}

// SliderAccessoryFrom constructs a [SliderAccessory] from an unsafe.Pointer.
func SliderAccessoryFrom(ptr unsafe.Pointer) SliderAccessory {
	return SliderAccessory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SliderAccessory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(image:)
func NewSliderAccessoryWithImage(image IImage) SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(getSliderAccessoryClass().class), objc.Sel("accessoryWithImage:"), image)
	return rv
}/* debug [class_init_methods/constructor]: NewSliderAccessoryWithImage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SliderAccessory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(image:)
func (sc _SliderAccessoryClass) AccessoryWithImage(image IImage) ISliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.class), objc.Sel("accessoryWithImage:"), image)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AccessoryWithImage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SliderAccessory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SliderAccessory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SliderAccessory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/behavior
func (s_ SliderAccessory) Behavior() ISliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/behavior
func (s_ SliderAccessory) SetBehavior(value ISliderAccessoryBehavior) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBehavior:"), value)
}/* debug [instance_properties/setter]: behavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/isEnabled
func (s_ SliderAccessory) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/isEnabled
func (s_ SliderAccessory) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/isenabled
func (s_ SliderAccessory) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/isenabled
func (s_ SliderAccessory) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSliderAccessory */


