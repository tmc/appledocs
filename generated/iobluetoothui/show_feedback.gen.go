// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class showFeedback */


/* debug [class_header]: Header for showFeedback */
// The class instance for the [showFeedback] class.
var (
	ShowFeedbackClass     _showFeedbackClass
	ShowFeedbackClassOnce sync.Once
)

func getshowFeedbackClass() _showFeedbackClass {
	ShowFeedbackClassOnce.Do(func() {
		ShowFeedbackClass = _showFeedbackClass{objc.GetClass("showFeedback")}
	})
	return ShowFeedbackClass
}

type _showFeedbackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for showFeedback */
// An interface definition for the [showFeedback] class.
type IshowFeedback interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for showFeedback */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for showFeedback */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for showFeedback */
// Alloc allocates a new instance without initialization.
func (sc _showFeedbackClass) Alloc() showFeedback {
	rv := objc.Send[showFeedback](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _showFeedbackClass) New() showFeedback {
	rv := objc.Send[showFeedback](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ showFeedback) Init() showFeedback {
	rv := objc.Send[showFeedback](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ showFeedback) Autorelease() showFeedback {
	rv := objc.Send[showFeedback](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewshowFeedback creates a new showFeedback instance.
func NewshowFeedback() showFeedback {
	return getshowFeedbackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for showFeedback */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/showFeedback
type showFeedback struct {
	objectivec.Object
}

// showFeedbackFrom constructs a [showFeedback] from an unsafe.Pointer.
func showFeedbackFrom(ptr unsafe.Pointer) showFeedback {
	return showFeedback{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for showFeedback *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for showFeedback */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for showFeedback */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for showFeedback */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for showFeedback */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class showFeedback */



