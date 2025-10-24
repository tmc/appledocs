// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class AMWorkflowView */


/* debug [class_header]: Header for AMWorkflowView */
// The class instance for the [AMWorkflowView] class.
var (
	AMWorkflowViewClass     _AMWorkflowViewClass
	AMWorkflowViewClassOnce sync.Once
)

func getAMWorkflowViewClass() _AMWorkflowViewClass {
	AMWorkflowViewClassOnce.Do(func() {
		AMWorkflowViewClass = _AMWorkflowViewClass{objc.GetClass("AMWorkflowView")}
	})
	return AMWorkflowViewClass
}

type _AMWorkflowViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMWorkflowView */
// An interface definition for the [AMWorkflowView] class.
type IAMWorkflowView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for AMWorkflowView */
	// properties:
	WorkflowView() IAMWorkflowView
	SetWorkflowView(value IAMWorkflowView)
	IsEditable() bool
	SetIsEditable(value bool)
	WorkflowController() IAMWorkflowController
	SetWorkflowController(value IAMWorkflowController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMWorkflowView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMWorkflowView */
// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowViewClass) Alloc() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMWorkflowViewClass) New() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMWorkflowView) Init() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMWorkflowView) Autorelease() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflowView creates a new AMWorkflowView instance.
func NewAMWorkflowView() AMWorkflowView {
	return getAMWorkflowViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMWorkflowView */
// An object that lets you view and edit Automator workflows in your app.
//
// A workflow view displays an instance of . You can use Interface Builder to add an instance of to a window in your app. You can then add an object to the nib window and use the controller’s outlet to connect it to the workflow view. The controller object also has and actions that can be connected to buttons or other user interface elements.


// An object that lets you view and edit Automator workflows in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowView
type AMWorkflowView struct {
	appkit.View
}

// AMWorkflowViewFrom constructs a [AMWorkflowView] from an unsafe.Pointer.
//
// An object that lets you view and edit Automator workflows in your app.
func AMWorkflowViewFrom(ptr unsafe.Pointer) AMWorkflowView {
	return AMWorkflowView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMWorkflowView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMWorkflowView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMWorkflowView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMWorkflowView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMWorkflowView */

// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/workflowview-swift.property
func (a_ AMWorkflowView) WorkflowView() IAMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("workflowView"))
	return rv
}/* debug [instance_properties/getter]: workflowView */


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/workflowview-swift.property
func (a_ AMWorkflowView) SetWorkflowView(value IAMWorkflowView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowView:"), value)
}/* debug [instance_properties/setter]: workflowView */


// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/iseditable
func (a_ AMWorkflowView) IsEditable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/iseditable
func (a_ AMWorkflowView) SetIsEditable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// The view’s workflow controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/workflowcontroller
func (a_ AMWorkflowView) WorkflowController() IAMWorkflowController {
	rv := objc.Send[AMWorkflowController](a_.ID, objc.Sel("workflowController"))
	return rv
}/* debug [instance_properties/getter]: workflowController */


// The view’s workflow controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/workflowcontroller
func (a_ AMWorkflowView) SetWorkflowController(value IAMWorkflowController) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowController:"), value)
}/* debug [instance_properties/setter]: workflowController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMWorkflowView */



