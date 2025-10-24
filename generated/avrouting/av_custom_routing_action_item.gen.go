// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class AVCustomRoutingActionItem */


/* debug [class_header]: Header for AVCustomRoutingActionItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomRoutingActionItem */
// An interface definition for the [CustomRoutingActionItem] class.
type ICustomRoutingActionItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CustomRoutingActionItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomRoutingActionItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomRoutingActionItem */
// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingActionItemClass) Alloc() CustomRoutingActionItem {
	rv := objc.Send[CustomRoutingActionItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomRoutingActionItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomRoutingActionItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomRoutingActionItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomRoutingActionItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomRoutingActionItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomRoutingActionItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCustomRoutingActionItem */


