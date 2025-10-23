// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [AMWorkflowView] class.
type IAMWorkflowView interface {
	appkit.IView
	Editable() bool
	SetEditable(value bool)
	WorkflowController() AMWorkflowController
	SetWorkflowController(value IAMWorkflowController)
	WorkflowView() AMWorkflowView
	SetWorkflowView(value IAMWorkflowView)
	IsEditable() bool
	SetIsEditable(value bool)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowViewClass) Alloc() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowView/isEditable
func (a_ AMWorkflowView) Editable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowView/isEditable
func (a_ AMWorkflowView) SetEditable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEditable:"), value)
}


// The view’s workflow controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowView/workflowController
func (a_ AMWorkflowView) WorkflowController() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](a_.ID, objc.Sel("workflowController"))
	return rv
}


// The view’s workflow controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowView/workflowController
func (a_ AMWorkflowView) SetWorkflowController(value IAMWorkflowController) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowController:"), value)
}


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/workflowview-swift.property
func (a_ AMWorkflowView) WorkflowView() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("workflowView"))
	return rv
}


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/workflowview-swift.property
func (a_ AMWorkflowView) SetWorkflowView(value IAMWorkflowView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowView:"), value)
}


// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/iseditable
func (a_ AMWorkflowView) IsEditable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value that indicates whether the workflow view is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowview/iseditable
func (a_ AMWorkflowView) SetIsEditable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEditable:"), value)
}



