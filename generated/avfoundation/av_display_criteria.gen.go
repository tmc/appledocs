// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDisplayCriteria */


/* debug [class_header]: Header for AVDisplayCriteria */
// The class instance for the [DisplayCriteria] class.
var (
	DisplayCriteriaClass     _DisplayCriteriaClass
	DisplayCriteriaClassOnce sync.Once
)

func getDisplayCriteriaClass() _DisplayCriteriaClass {
	DisplayCriteriaClassOnce.Do(func() {
		DisplayCriteriaClass = _DisplayCriteriaClass{objc.GetClass("AVDisplayCriteria")}
	})
	return DisplayCriteriaClass
}

type _DisplayCriteriaClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DisplayCriteria */
// An interface definition for the [DisplayCriteria] class.
type IDisplayCriteria interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DisplayCriteria */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DisplayCriteria */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DisplayCriteria */
// Alloc allocates a new instance without initialization.
func (dc _DisplayCriteriaClass) Alloc() DisplayCriteria {
	rv := objc.Send[DisplayCriteria](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DisplayCriteriaClass) New() DisplayCriteria {
	rv := objc.Send[DisplayCriteria](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DisplayCriteria) Init() DisplayCriteria {
	rv := objc.Send[DisplayCriteria](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DisplayCriteria) Autorelease() DisplayCriteria {
	rv := objc.Send[DisplayCriteria](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplayCriteria creates a new DisplayCriteria instance.
func NewDisplayCriteria() DisplayCriteria {
	return getDisplayCriteriaClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DisplayCriteria */
// An object the system uses to guide the selection of a display mode in tvOS.
//
// In tvOS, this object provides the display criteria that an uses to set an appropriate display mode, such as switching to HDR, when presenting a video asset. If your app uses for its player user interface, the system automatically applies the display critera when it presents the asset. If you use a custom player interface, load the value of an asset’s property and set it on the window’s object.


// An object the system uses to guide the selection of a display mode in tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDisplayCriteria
type DisplayCriteria struct {
	objectivec.Object
}

// DisplayCriteriaFrom constructs a [DisplayCriteria] from an unsafe.Pointer.
//
// An object the system uses to guide the selection of a display mode in tvOS.
func DisplayCriteriaFrom(ptr unsafe.Pointer) DisplayCriteria {
	return DisplayCriteria{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DisplayCriteria */

// Creates a display criteria object with the specified refresh rate and format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDisplayCriteria/init(refreshRate:formatDescription:)
func NewDisplayCriteriaWithRefreshRateFormatDescription(refreshRate float32, formatDescription FormatDescriptionRef /* not a class type */) DisplayCriteria {
	instance := getDisplayCriteriaClass().Alloc()
	rv := objc.Send[DisplayCriteria](instance.ID, objc.Sel("initWithRefreshRate:formatDescription:"), refreshRate, formatDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDisplayCriteriaWithRefreshRateFormatDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DisplayCriteria */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DisplayCriteria */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DisplayCriteria */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DisplayCriteria */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDisplayCriteria */


