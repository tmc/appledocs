// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	CursorGlyphHandler() unsafe.Pointer
	SetCursorGlyphHandler(value unsafe.Pointer)
	CursorMoveHandler() unsafe.Pointer
	SetCursorMoveHandler(value unsafe.Pointer)
	CursorShowHandler() unsafe.Pointer
	SetCursorShowHandler(value unsafe.Pointer)
	ModeChangeHandler() unsafe.Pointer
	SetModeChangeHandler(value unsafe.Pointer)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	NewFrameEventHandler() unsafe.Pointer
	SetNewFrameEventHandler(value unsafe.Pointer)
	Queue() unsafe.Pointer
	SetQueue(value unsafe.Pointer)
	SizeInMillimeters() objc.IObject /* cross-framework: Size */
	SetSizeInMillimeters(value objc.IObject /* cross-framework: Size */)
	// methods:
}

// A descriptor for a virtual display.


// A descriptor for a virtual display.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorGlyphHandler
func (p_ PGDisplayDescriptor) CursorGlyphHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorGlyphHandler"))
	return rv
}


// A handler that the framework calls to change the cursor’s appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorGlyphHandler
func (p_ PGDisplayDescriptor) SetCursorGlyphHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorGlyphHandler:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorMoveHandler
func (p_ PGDisplayDescriptor) CursorMoveHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorMoveHandler"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorMoveHandler
func (p_ PGDisplayDescriptor) SetCursorMoveHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorMoveHandler:"), value)
}


// A handler that the framework calls to change the cursor’s visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorShowHandler
func (p_ PGDisplayDescriptor) CursorShowHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cursorShowHandler"))
	return rv
}


// A handler that the framework calls to change the cursor’s visibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorShowHandler
func (p_ PGDisplayDescriptor) SetCursorShowHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorShowHandler:"), value)
}


// A handler that the framework calls to change the virtual display’s graphics mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/modeChangeHandler
func (p_ PGDisplayDescriptor) ModeChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modeChangeHandler"))
	return rv
}


// A handler that the framework calls to change the virtual display’s graphics mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/modeChangeHandler
func (p_ PGDisplayDescriptor) SetModeChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModeChangeHandler:"), value)
}


// The display’s name as seen in the guest operating environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/name
func (p_ PGDisplayDescriptor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("name"))
	return rv
}


// The display’s name as seen in the guest operating environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/name
func (p_ PGDisplayDescriptor) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}


// A handler that the framework calls when the guest environment has a new frame to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/newFrameEventHandler
func (p_ PGDisplayDescriptor) NewFrameEventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("newFrameEventHandler"))
	return rv
}


// A handler that the framework calls when the guest environment has a new frame to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/newFrameEventHandler
func (p_ PGDisplayDescriptor) SetNewFrameEventHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNewFrameEventHandler:"), value)
}


// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/queue
func (p_ PGDisplayDescriptor) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("queue"))
	return rv
}


// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/queue
func (p_ PGDisplayDescriptor) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setQueue:"), value)
}


// The size in millimeters of the virtual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/sizeInMillimeters
func (p_ PGDisplayDescriptor) SizeInMillimeters() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("sizeInMillimeters"))
	return rv
}


// The size in millimeters of the virtual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/sizeInMillimeters
func (p_ PGDisplayDescriptor) SetSizeInMillimeters(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSizeInMillimeters:"), value)
}



