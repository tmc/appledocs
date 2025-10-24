// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class VZVirtualMachineView */


/* debug [class_header]: Header for VZVirtualMachineView */
// The class instance for the [VZVirtualMachineView] class.
var (
	VZVirtualMachineViewClass     _VZVirtualMachineViewClass
	VZVirtualMachineViewClassOnce sync.Once
)

func getVZVirtualMachineViewClass() _VZVirtualMachineViewClass {
	VZVirtualMachineViewClassOnce.Do(func() {
		VZVirtualMachineViewClass = _VZVirtualMachineViewClass{objc.GetClass("VZVirtualMachineView")}
	})
	return VZVirtualMachineViewClass
}

type _VZVirtualMachineViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtualMachineView */
// An interface definition for the [VZVirtualMachineView] class.
type IVZVirtualMachineView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for VZVirtualMachineView */
	// properties:
	AutomaticallyReconfiguresDisplay() bool
	SetAutomaticallyReconfiguresDisplay(value bool)
	CapturesSystemKeys() bool
	SetCapturesSystemKeys(value bool)
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtualMachineView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtualMachineView */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineViewClass) Alloc() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtualMachineViewClass) New() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtualMachineView) Init() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtualMachineView) Autorelease() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineView creates a new VZVirtualMachineView instance.
func NewVZVirtualMachineView() VZVirtualMachineView {
	return getVZVirtualMachineViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtualMachineView */
// A view that allows user interaction with a VM.
//
// The is a UI element that shows the contents of the VM frame buffer that you can optionally configure to respond to changes in the host’s display settings. If the VM configuration includes a keyboard and a pointing device, the view forwards keyboard and mouse events to the VM through those devices.


// A view that allows user interaction with a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView
type VZVirtualMachineView struct {
	appkit.View
}

// VZVirtualMachineViewFrom constructs a [VZVirtualMachineView] from an unsafe.Pointer.
//
// A view that allows user interaction with a VM.
func VZVirtualMachineViewFrom(ptr unsafe.Pointer) VZVirtualMachineView {
	return VZVirtualMachineView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtualMachineView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtualMachineView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtualMachineView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtualMachineView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtualMachineView */

// A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/automaticallyReconfiguresDisplay
func (v_ VZVirtualMachineView) AutomaticallyReconfiguresDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("automaticallyReconfiguresDisplay"))
	return rv
}/* debug [instance_properties/getter]: automaticallyReconfiguresDisplay */


// A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/automaticallyReconfiguresDisplay
func (v_ VZVirtualMachineView) SetAutomaticallyReconfiguresDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutomaticallyReconfiguresDisplay:"), value)
}/* debug [instance_properties/setter]: automaticallyReconfiguresDisplay */


// A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/capturesSystemKeys
func (v_ VZVirtualMachineView) CapturesSystemKeys() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("capturesSystemKeys"))
	return rv
}/* debug [instance_properties/getter]: capturesSystemKeys */


// A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/capturesSystemKeys
func (v_ VZVirtualMachineView) SetCapturesSystemKeys(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCapturesSystemKeys:"), value)
}/* debug [instance_properties/setter]: capturesSystemKeys */


// The VM to display in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/virtualMachine
func (v_ VZVirtualMachineView) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v_.ID, objc.Sel("virtualMachine"))
	return rv
}/* debug [instance_properties/getter]: virtualMachine */


// The VM to display in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/virtualMachine
func (v_ VZVirtualMachineView) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVirtualMachine:"), value)
}/* debug [instance_properties/setter]: virtualMachine */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtualMachineView */



