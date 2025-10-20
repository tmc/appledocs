// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKFilterBrowserPanel] class.
var (
	IKFilterBrowserPanelClass     _IKFilterBrowserPanelClass
	IKFilterBrowserPanelClassOnce sync.Once
)

func getIKFilterBrowserPanelClass() _IKFilterBrowserPanelClass {
	IKFilterBrowserPanelClassOnce.Do(func() {
		IKFilterBrowserPanelClass = _IKFilterBrowserPanelClass{objc.GetClass("IKFilterBrowserPanel")}
	})
	return IKFilterBrowserPanelClass
}

type _IKFilterBrowserPanelClass struct {
	class objc.Class
}

// An interface definition for the [IKFilterBrowserPanel] class.
type IIKFilterBrowserPanel interface {
	appkit.IPanel
}

// Presents a user interface for browsing filters.
//
// The class provides a user interface that allows users to browse Core Image filters ( ), to preview a filter, and to get additional information about the filter, such as its description. An object can be displayed as: a separate panel, that is, a utility window that floats on top of document windows a modal dialog a sheet, that is, a dialog that is attached to its parent window and must be dismissed by the user a view that an application can insert into a custom user interface An object can be configured through a style mask to use either the default or brushed metal look for windows. The size and number of visible controls are specified through an options dictionary. An object communicates selection changes through notifications. The class allows the user to create filter collections that are stored with the key in the property list located in .
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel
type IKFilterBrowserPanel struct {
	appkit.Panel
}

// IKFilterBrowserPanelFrom constructs a [IKFilterBrowserPanel] from an unsafe.Pointer.
//
// Presents a user interface for browsing filters.
func IKFilterBrowserPanelFrom(ptr unsafe.Pointer) IKFilterBrowserPanel {
	return IKFilterBrowserPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKFilterBrowserPanelClass) Alloc() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKFilterBrowserPanelClass) New() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterBrowserPanel) Init() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterBrowserPanel) Autorelease() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterBrowserPanel creates a new IKFilterBrowserPanel instance.
func NewIKFilterBrowserPanel() IKFilterBrowserPanel {
	return getIKFilterBrowserPanelClass().New()
}




