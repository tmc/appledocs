
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Panel] class.
var PanelClass _PanelClass

func init() {
	PanelClass = _PanelClass{objc.GetClass("NSPanel")}
}

type _PanelClass struct {
	objc.Class
}

// An interface definition for the [Panel] class.
type IPanel interface {
	ID() objc.ID
}

type Panel struct {
	id objc.ID
}

func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ Panel) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PanelClass) Alloc() Panel {
	rv := objc.Send[Panel](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PanelClass) New() Panel {
	rv := objc.Send[Panel](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPanel creates and returns a new initialized instance.
func NewPanel() Panel {
	return PanelClass.New()
}

// Init initializes the instance.
func (p_ Panel) Init() Panel {
	rv := objc.Send[Panel](p_.ID(), selInit)
	return rv
}
// A Boolean value that indicates whether the receiver becomes the key window only when needed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("becomesKeyOnlyIfNeeded"))
	return rv
}
// SetBecomesKeyOnlyIfNeeded sets the value of the becomesKeyOnlyIfNeeded property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setBecomesKeyOnlyIfNeeded:"), value)
}
// A Boolean value that indicates whether the receiver is a floating panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) FloatingPanel() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("floatingPanel"))
	return rv
}
// SetFloatingPanel sets the value of the floatingPanel property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) SetFloatingPanel(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setFloatingPanel:"), value)
}
// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) WorksWhenModal() bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("worksWhenModal"))
	return rv
}
// SetWorksWhenModal sets the value of the worksWhenModal property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setWorksWhenModal:"), value)
}
