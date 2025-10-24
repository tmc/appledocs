// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UIScrollView */


/* debug [class_header]: Header for UIScrollView */
// The class instance for the [ScrollView] class.
var (
	ScrollViewClass     _ScrollViewClass
	ScrollViewClassOnce sync.Once
)

func getScrollViewClass() _ScrollViewClass {
	ScrollViewClassOnce.Do(func() {
		ScrollViewClass = _ScrollViewClass{objc.GetClass("UIScrollView")}
	})
	return ScrollViewClass
}

type _ScrollViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrollView */
// An interface definition for the [ScrollView] class.
type IScrollView interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScrollView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrollView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrollView */
// Alloc allocates a new instance without initialization.
func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollView) Init() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollView) Autorelease() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollView creates a new ScrollView instance.
func NewScrollView() ScrollView {
	return getScrollViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrollView */
// A parent class referenced by other PencilKit classes.


// A parent class referenced by other PencilKit classes. [Full Topic]
type ScrollView struct {
	objectivec.Object
}

// ScrollViewFrom constructs a [ScrollView] from an unsafe.Pointer.
//
// A parent class referenced by other PencilKit classes.
func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrollView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrollView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrollView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrollView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrollView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UIScrollView */



