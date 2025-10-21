// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PGDisplayDescriptor] class.
var (
	PGDisplayDescriptorClass     _PGDisplayDescriptorClass
	PGDisplayDescriptorClassOnce sync.Once
)

func getPGDisplayDescriptorClass() _PGDisplayDescriptorClass {
	PGDisplayDescriptorClassOnce.Do(func() {
		PGDisplayDescriptorClass = _PGDisplayDescriptorClass{objc.GetClass("PGDisplayDescriptor")}
	})
	return PGDisplayDescriptorClass
}

type _PGDisplayDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PGDisplayDescriptor] class.
type IPGDisplayDescriptor interface {
	objectivec.IObject
}

// A descriptor for a virtual display.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor
type PGDisplayDescriptor struct {
	objectivec.Object
}

// PGDisplayDescriptorFrom constructs a [PGDisplayDescriptor] from an unsafe.Pointer.
//
// A descriptor for a virtual display.
func PGDisplayDescriptorFrom(ptr unsafe.Pointer) PGDisplayDescriptor {
	return PGDisplayDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PGDisplayDescriptorClass) Alloc() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PGDisplayDescriptorClass) New() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PGDisplayDescriptor) Init() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PGDisplayDescriptor) Autorelease() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDisplayDescriptor creates a new PGDisplayDescriptor instance.
func NewPGDisplayDescriptor() PGDisplayDescriptor {
	return getPGDisplayDescriptorClass().New()
}


// A handler that the framework calls to change the cursor’s appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorGlyphHandler
func (p_ PGDisplayDescriptor) CursorGlyphHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorGlyphHandler"))
	return rv
}


// SetCursorGlyphHandler sets the value of the cursorGlyphHandler property.
// A handler that the framework calls to change the cursor’s appearance.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorGlyphHandler
func (p_ PGDisplayDescriptor) SetCursorGlyphHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorGlyphHandler:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorMoveHandler
func (p_ PGDisplayDescriptor) CursorMoveHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorMoveHandler"))
	return rv
}


// SetCursorMoveHandler sets the value of the cursorMoveHandler property.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorMoveHandler
func (p_ PGDisplayDescriptor) SetCursorMoveHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorMoveHandler:"), value)
}

// A handler that the framework calls to change the cursor’s visibility.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorShowHandler
func (p_ PGDisplayDescriptor) CursorShowHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorShowHandler"))
	return rv
}


// SetCursorShowHandler sets the value of the cursorShowHandler property.
// A handler that the framework calls to change the cursor’s visibility.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorShowHandler
func (p_ PGDisplayDescriptor) SetCursorShowHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorShowHandler:"), value)
}

// A handler that the framework calls to change the virtual display’s graphics mode.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/modeChangeHandler
func (p_ PGDisplayDescriptor) ModeChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modeChangeHandler"))
	return rv
}


// SetModeChangeHandler sets the value of the modeChangeHandler property.
// A handler that the framework calls to change the virtual display’s graphics mode.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/modeChangeHandler
func (p_ PGDisplayDescriptor) SetModeChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModeChangeHandler:"), value)
}

// The display’s name as seen in the guest operating environment.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/name
func (p_ PGDisplayDescriptor) Name() string {
	rv := objc.Send[string](p_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The display’s name as seen in the guest operating environment.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/name
func (p_ PGDisplayDescriptor) SetName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), objc.String(value))
}

// A handler that the framework calls when the guest environment has a new frame to display.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/newFrameEventHandler
func (p_ PGDisplayDescriptor) NewFrameEventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("newFrameEventHandler"))
	return rv
}


// SetNewFrameEventHandler sets the value of the newFrameEventHandler property.
// A handler that the framework calls when the guest environment has a new frame to display.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/newFrameEventHandler
func (p_ PGDisplayDescriptor) SetNewFrameEventHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNewFrameEventHandler:"), value)
}

// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/queue
func (p_ PGDisplayDescriptor) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("queue"))
	return rv
}


// SetQueue sets the value of the queue property.
// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/queue
func (p_ PGDisplayDescriptor) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setQueue:"), value)
}

// The size in millimeters of the virtual display.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/sizeInMillimeters
func (p_ PGDisplayDescriptor) SizeInMillimeters() foundation.Size {
	rv := objc.Send[foundation.Size](p_.ID, objc.Sel("sizeInMillimeters"))
	return rv
}


// SetSizeInMillimeters sets the value of the sizeInMillimeters property.
// The size in millimeters of the virtual display.

//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/sizeInMillimeters
func (p_ PGDisplayDescriptor) SetSizeInMillimeters(value foundation.Size) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSizeInMillimeters:"), value)
}



