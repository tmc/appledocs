// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [VZVirtualMachineView] class.
type IVZVirtualMachineView interface {
	appkit.IView
}

// A view that allows user interaction with a VM.
//
// The is a UI element that shows the contents of the VM frame buffer that you can optionally configure to respond to changes in the host’s display settings. If the VM configuration includes a keyboard and a pointing device, the view forwards keyboard and mouse events to the VM through those devices.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineViewClass) Alloc() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/automaticallyReconfiguresDisplay
func (v_ VZVirtualMachineView) AutomaticallyReconfiguresDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("automaticallyReconfiguresDisplay"))
	return rv
}


// SetAutomaticallyReconfiguresDisplay sets the value of the automaticallyReconfiguresDisplay property.
// A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/automaticallyReconfiguresDisplay
func (v_ VZVirtualMachineView) SetAutomaticallyReconfiguresDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutomaticallyReconfiguresDisplay:"), value)
}
// A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/capturesSystemKeys
func (v_ VZVirtualMachineView) CapturesSystemKeys() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("capturesSystemKeys"))
	return rv
}


// SetCapturesSystemKeys sets the value of the capturesSystemKeys property.
// A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/capturesSystemKeys
func (v_ VZVirtualMachineView) SetCapturesSystemKeys(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCapturesSystemKeys:"), value)
}
// The VM to display in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/virtualMachine
func (v_ VZVirtualMachineView) VirtualMachine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("virtualMachine"))
	return rv
}


// SetVirtualMachine sets the value of the virtualMachine property.
// The VM to display in the view.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/virtualMachine
func (v_ VZVirtualMachineView) SetVirtualMachine(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVirtualMachine:"), value)
}


