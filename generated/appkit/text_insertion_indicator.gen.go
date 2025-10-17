
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextInsertionIndicator] class.
var TextInsertionIndicatorClass _TextInsertionIndicatorClass

func init() {
	TextInsertionIndicatorClass = _TextInsertionIndicatorClass{objc.GetClass("NSTextInsertionIndicator")}
}

type _TextInsertionIndicatorClass struct {
	objc.Class
}

// An interface definition for the [TextInsertionIndicator] class.
type ITextInsertionIndicator interface {
	ID() objc.ID
}

type TextInsertionIndicator struct {
	id objc.ID
}

func TextInsertionIndicatorFrom(ptr unsafe.Pointer) TextInsertionIndicator {
	return TextInsertionIndicator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextInsertionIndicator) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextInsertionIndicatorClass) Alloc() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextInsertionIndicatorClass) New() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextInsertionIndicator creates and returns a new initialized instance.
func NewTextInsertionIndicator() TextInsertionIndicator {
	return TextInsertionIndicatorClass.New()
}

// Init initializes the instance.
func (t_ TextInsertionIndicator) Init() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](t_.ID(), selInit)
	return rv
}
// Options that affect the automatic display mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/automaticModeOptions-swift.property
func (t_ TextInsertionIndicator) AutomaticModeOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("automaticModeOptions"))
	return rv
}
// SetAutomaticModeOptions sets the value of the automaticModeOptions property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/automaticModeOptions-swift.property
func (t_ TextInsertionIndicator) SetAutomaticModeOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAutomaticModeOptions:"), value)
}
// The color of this indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/color
func (t_ TextInsertionIndicator) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("color"))
	return rv
}
// SetColor sets the value of the color property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/color
func (t_ TextInsertionIndicator) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setColor:"), value)
}
// A value that describes the display mode of an indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) DisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("displayMode"))
	return rv
}
// SetDisplayMode sets the value of the displayMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextInsertionIndicator/displayMode-swift.property
func (t_ TextInsertionIndicator) SetDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDisplayMode:"), value)
}
