// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreimage"
)

/* debug [class.gen.go]: Generating class IKFilterUIView */


/* debug [class_header]: Header for IKFilterUIView */
// The class instance for the [IKFilterUIView] class.
var (
	IKFilterUIViewClass     _IKFilterUIViewClass
	IKFilterUIViewClassOnce sync.Once
)

func getIKFilterUIViewClass() _IKFilterUIViewClass {
	IKFilterUIViewClassOnce.Do(func() {
		IKFilterUIViewClass = _IKFilterUIViewClass{objc.GetClass("IKFilterUIView")}
	})
	return IKFilterUIViewClass
}

type _IKFilterUIViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKFilterUIView */
// An interface definition for the [IKFilterUIView] class.
type IIKFilterUIView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKFilterUIView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKFilterUIView */
	// methods:
	Filter() coreimage.Filter
	ObjectController() appkit.ObjectController
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKFilterUIView */
// Alloc allocates a new instance without initialization.
func (ic _IKFilterUIViewClass) Alloc() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKFilterUIViewClass) New() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterUIView) Init() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterUIView) Autorelease() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterUIView creates a new IKFilterUIView instance.
func NewIKFilterUIView() IKFilterUIView {
	return getIKFilterUIViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKFilterUIView */
// Input parameters for filtering core image filters.
//
// The class provides a view that contains input parameter controls for a Core Image filter ( ). You need to use this class when providing a user interface for a custom filter. The class creates a view that has an object controller for the given filter. It also retains the filter.


// Input parameters for filtering core image filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView
type IKFilterUIView struct {
	appkit.View
}

// IKFilterUIViewFrom constructs a [IKFilterUIView] from an unsafe.Pointer.
//
// Input parameters for filtering core image filters.
func IKFilterUIViewFrom(ptr unsafe.Pointer) IKFilterUIView {
	return IKFilterUIView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKFilterUIView */

// Initializes a view that contains controls for the input parameters of a filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView/init(frame:filter:)
func NewIKFilterUIViewWithFrameFilter(frameRect Rect /* not a class type */, inFilter coreimage.Filter) IKFilterUIView {
	instance := getIKFilterUIViewClass().Alloc()
	rv := objc.Send[IKFilterUIView](instance.ID, objc.Sel("initWithFrame:filter:"), frameRect, inFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIKFilterUIViewWithFrameFilter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKFilterUIView */

// Creates a view that contains controls for the input parameters of a filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView/view(withFrame:filter:)
func (ic _IKFilterUIViewClass) ViewWithFrameFilter(frameRect Rect /* not a class type */, inFilter coreimage.Filter) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("viewWithFrame:filter:"), frameRect, inFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ViewWithFrameFilter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKFilterUIView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKFilterUIView */

// Returns the Core Image filter associated with the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView/filter()
func (i_ IKFilterUIView) Filter() coreimage.Filter {
	rv := objc.Send[coreimage.Filter](i_.ID, objc.Sel("filter"))
	return rv
}/* debug [instance_methods/method]: Filter */


// Returns the object controller for the bindings between the filter and its view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView/objectController()
func (i_ IKFilterUIView) ObjectController() appkit.ObjectController {
	rv := objc.Send[appkit.ObjectController](i_.ID, objc.Sel("objectController"))
	return rv
}/* debug [instance_methods/method]: ObjectController */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKFilterUIView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKFilterUIView */


