// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IKFilterBrowserView */


/* debug [class_header]: Header for IKFilterBrowserView */
// The class instance for the [IKFilterBrowserView] class.
var (
	IKFilterBrowserViewClass     _IKFilterBrowserViewClass
	IKFilterBrowserViewClassOnce sync.Once
)

func getIKFilterBrowserViewClass() _IKFilterBrowserViewClass {
	IKFilterBrowserViewClassOnce.Do(func() {
		IKFilterBrowserViewClass = _IKFilterBrowserViewClass{objc.GetClass("IKFilterBrowserView")}
	})
	return IKFilterBrowserViewClass
}

type _IKFilterBrowserViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKFilterBrowserView */
// An interface definition for the [IKFilterBrowserView] class.
type IIKFilterBrowserView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKFilterBrowserView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKFilterBrowserView */
	// methods:
	FilterName() foundation.String
	SetPreviewState(inState bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKFilterBrowserView */
// Alloc allocates a new instance without initialization.
func (ic _IKFilterBrowserViewClass) Alloc() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKFilterBrowserViewClass) New() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterBrowserView) Init() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterBrowserView) Autorelease() IKFilterBrowserView {
	rv := objc.Send[IKFilterBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterBrowserView creates a new IKFilterBrowserView instance.
func NewIKFilterBrowserView() IKFilterBrowserView {
	return getIKFilterBrowserViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKFilterBrowserView */
// The class is used as a container for the elements of an object.


// The class is used as a container for the elements of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserView
type IKFilterBrowserView struct {
	appkit.View
}

// IKFilterBrowserViewFrom constructs a [IKFilterBrowserView] from an unsafe.Pointer.
//
// The class is used as a container for the elements of an object.
func IKFilterBrowserViewFrom(ptr unsafe.Pointer) IKFilterBrowserView {
	return IKFilterBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKFilterBrowserView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKFilterBrowserView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKFilterBrowserView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKFilterBrowserView */

// Returns the name of the filter that is currently selected in the filter browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserView/filterName()
func (i_ IKFilterBrowserView) FilterName() foundation.String {
	rv := objc.Send[foundation.String](i_.ID, objc.Sel("filterName"))
	return rv
}/* debug [instance_methods/method]: FilterName */


// Sets the preview state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserView/setPreviewState(_:)
func (i_ IKFilterBrowserView) SetPreviewState(inState bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreviewState:"), inState)
}/* debug [instance_methods/method]: SetPreviewState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKFilterBrowserView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKFilterBrowserView */



