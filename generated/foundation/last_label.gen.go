// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class lastLabel */


/* debug [class_header]: Header for lastLabel */
// The class instance for the [lastLabel] class.
var (
	LastLabelClass     _lastLabelClass
	LastLabelClassOnce sync.Once
)

func getlastLabelClass() _lastLabelClass {
	LastLabelClassOnce.Do(func() {
		LastLabelClass = _lastLabelClass{objc.GetClass("lastLabel")}
	})
	return LastLabelClass
}

type _lastLabelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for lastLabel */
// An interface definition for the [lastLabel] class.
type IlastLabel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for lastLabel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for lastLabel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for lastLabel */
// Alloc allocates a new instance without initialization.
func (lc _lastLabelClass) Alloc() lastLabel {
	rv := objc.Send[lastLabel](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _lastLabelClass) New() lastLabel {
	rv := objc.Send[lastLabel](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ lastLabel) Init() lastLabel {
	rv := objc.Send[lastLabel](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ lastLabel) Autorelease() lastLabel {
	rv := objc.Send[lastLabel](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlastLabel creates a new lastLabel instance.
func NewlastLabel() lastLabel {
	return getlastLabelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for lastLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/lastLabel
type lastLabel struct {
	objectivec.Object
}

// lastLabelFrom constructs a [lastLabel] from an unsafe.Pointer.
func lastLabelFrom(ptr unsafe.Pointer) lastLabel {
	return lastLabel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for lastLabel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for lastLabel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for lastLabel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for lastLabel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for lastLabel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class lastLabel */



