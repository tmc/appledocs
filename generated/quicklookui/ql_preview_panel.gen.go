// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PreviewPanel] class.
var (
	PreviewPanelClass     _PreviewPanelClass
	PreviewPanelClassOnce sync.Once
)

func getPreviewPanelClass() _PreviewPanelClass {
	PreviewPanelClassOnce.Do(func() {
		PreviewPanelClass = _PreviewPanelClass{objc.GetClass("QLPreviewPanel")}
	})
	return PreviewPanelClass
}

type _PreviewPanelClass struct {
	class objc.Class
}

// An interface definition for the [PreviewPanel] class.
type IPreviewPanel interface {
	appkit.IPanel
}

// A class that implements the Quick Look preview panel to display a preview of a list of items.
//
// Every application has a single shared instance of accessible through . The preview panel follows the responder chain and adapts to the first responder willing to control it. A preview panel controller provides the content through methods defined in the protocol. You can’t subclass ; you can, however, customize its behavior using a . See the protocol for the methods to customize a preview panel’s behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel
type PreviewPanel struct {
	appkit.Panel
}

// PreviewPanelFrom constructs a [PreviewPanel] from an unsafe.Pointer.
//
// A class that implements the Quick Look preview panel to display a preview of a list of items.
func PreviewPanelFrom(ptr unsafe.Pointer) PreviewPanel {
	return PreviewPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewPanelClass) Alloc() PreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewPanelClass) New() PreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewPanel) Init() PreviewPanel {
	rv := objc.Send[PreviewPanel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewPanel) Autorelease() PreviewPanel {
	rv := objc.Send[PreviewPanel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewPanel creates a new PreviewPanel instance.
func NewPreviewPanel() PreviewPanel {
	return getPreviewPanelClass().New()
}


// Returns the shared Quick Look preview panel instance.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/shared()
func (pc _PreviewPanelClass) SharedPreviewPanel() PreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("sharedPreviewPanel"))
	return rv
}

// The current first responder accepting to control the preview panel.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/currentController
func (p_ PreviewPanel) CurrentController() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("currentController"))
	return rv
}

// The property that indicates whether the panel is in full screen mode.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/isInFullScreenMode
func (p_ PreviewPanel) InFullScreenMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("inFullScreenMode"))
	return rv
}

// The currently previewed item.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/currentpreviewitem
func (p_ PreviewPanel) CurrentPreviewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentPreviewItem"))
	return rv
}


// SetCurrentPreviewItem sets the value of the currentPreviewItem property.
// The currently previewed item.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/currentpreviewitem
func (p_ PreviewPanel) SetCurrentPreviewItem(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItem:"), value)
}

// The index of the current preview item.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/currentpreviewitemindex
func (p_ PreviewPanel) CurrentPreviewItemIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPreviewItemIndex"))
	return rv
}


// SetCurrentPreviewItemIndex sets the value of the currentPreviewItemIndex property.
// The index of the current preview item.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/currentpreviewitemindex
func (p_ PreviewPanel) SetCurrentPreviewItemIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItemIndex:"), value)
}

// The preview panel data source.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/datasource
func (p_ PreviewPanel) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The preview panel data source.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/datasource
func (p_ PreviewPanel) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}

// The delegate object that controls the preview panel’s behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/delegate
func (p_ PreviewPanel) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that controls the preview panel’s behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/delegate
func (p_ PreviewPanel) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// The preview panel’s display state.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/displaystate
func (p_ PreviewPanel) DisplayState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("displayState"))
	return rv
}


// SetDisplayState sets the value of the displayState property.
// The preview panel’s display state.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/displaystate
func (p_ PreviewPanel) SetDisplayState(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayState:"), value)
}

// The property that indicates whether the panel is in full screen mode.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/isinfullscreenmode
func (p_ PreviewPanel) IsInFullScreenMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInFullScreenMode"))
	return rv
}


// SetIsInFullScreenMode sets the value of the isInFullScreenMode property.
// The property that indicates whether the panel is in full screen mode.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/isinfullscreenmode
func (p_ PreviewPanel) SetIsInFullScreenMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInFullScreenMode:"), value)
}



