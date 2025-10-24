// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class RPSystemBroadcastPickerView */


/* debug [class_header]: Header for RPSystemBroadcastPickerView */
// The class instance for the [RPSystemBroadcastPickerView] class.
var (
	RPSystemBroadcastPickerViewClass     _RPSystemBroadcastPickerViewClass
	RPSystemBroadcastPickerViewClassOnce sync.Once
)

func getRPSystemBroadcastPickerViewClass() _RPSystemBroadcastPickerViewClass {
	RPSystemBroadcastPickerViewClassOnce.Do(func() {
		RPSystemBroadcastPickerViewClass = _RPSystemBroadcastPickerViewClass{objc.GetClass("RPSystemBroadcastPickerView")}
	})
	return RPSystemBroadcastPickerViewClass
}

type _RPSystemBroadcastPickerViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPSystemBroadcastPickerView */
// An interface definition for the [RPSystemBroadcastPickerView] class.
type IRPSystemBroadcastPickerView interface {
	IView
	
/* debug [class_interface_properties]: Properties for RPSystemBroadcastPickerView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPSystemBroadcastPickerView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPSystemBroadcastPickerView */
// Alloc allocates a new instance without initialization.
func (rc _RPSystemBroadcastPickerViewClass) Alloc() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPSystemBroadcastPickerViewClass) New() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPSystemBroadcastPickerView) Init() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPSystemBroadcastPickerView) Autorelease() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPSystemBroadcastPickerView creates a new RPSystemBroadcastPickerView instance.
func NewRPSystemBroadcastPickerView() RPSystemBroadcastPickerView {
	return getRPSystemBroadcastPickerViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPSystemBroadcastPickerView */
// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
//
// Add this view to your view hierarchy to let users broadcast directly from your app. When a user taps the broadcast button, it displays a broadcast picker, allowing the user to select a broadcast provider. You can limit the picker to a particular broadcast provider by setting to the bundle identifier of a broadcast extension. You can also show or hide the microphone button displayed in the picker by setting the property. Set these properties before presenting , as shown here:


// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPSystemBroadcastPickerView
type RPSystemBroadcastPickerView struct {
	View
}

// RPSystemBroadcastPickerViewFrom constructs a [RPSystemBroadcastPickerView] from an unsafe.Pointer.
//
// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
func RPSystemBroadcastPickerViewFrom(ptr unsafe.Pointer) RPSystemBroadcastPickerView {
	return RPSystemBroadcastPickerView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPSystemBroadcastPickerView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPSystemBroadcastPickerView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPSystemBroadcastPickerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPSystemBroadcastPickerView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPSystemBroadcastPickerView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPSystemBroadcastPickerView */


