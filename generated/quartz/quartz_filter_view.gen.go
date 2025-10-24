// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class QuartzFilterView */


/* debug [class_header]: Header for QuartzFilterView */
// The class instance for the [QuartzFilterView] class.
var (
	QuartzFilterViewClass     _QuartzFilterViewClass
	QuartzFilterViewClassOnce sync.Once
)

func getQuartzFilterViewClass() _QuartzFilterViewClass {
	QuartzFilterViewClassOnce.Do(func() {
		QuartzFilterViewClass = _QuartzFilterViewClass{objc.GetClass("QuartzFilterView")}
	})
	return QuartzFilterViewClass
}

type _QuartzFilterViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuartzFilterView */
// An interface definition for the [QuartzFilterView] class.
type IQuartzFilterView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for QuartzFilterView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuartzFilterView */
	// methods:
	SizeToFit()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuartzFilterView */
// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterViewClass) Alloc() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QuartzFilterViewClass) New() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilterView) Init() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilterView) Autorelease() QuartzFilterView {
	rv := objc.Send[QuartzFilterView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilterView creates a new QuartzFilterView instance.
func NewQuartzFilterView() QuartzFilterView {
	return getQuartzFilterViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuartzFilterView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterView
type QuartzFilterView struct {
	appkit.View
}

// QuartzFilterViewFrom constructs a [QuartzFilterView] from an unsafe.Pointer.
func QuartzFilterViewFrom(ptr unsafe.Pointer) QuartzFilterView {
	return QuartzFilterView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuartzFilterView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuartzFilterView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuartzFilterView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuartzFilterView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterView/sizeToFit()
func (q_ QuartzFilterView) SizeToFit() {
	objc.Send[objc.ID](q_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuartzFilterView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QuartzFilterView */





