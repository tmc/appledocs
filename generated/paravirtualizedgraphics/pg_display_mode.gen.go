// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PGDisplayMode] class.
var (
	PGDisplayModeClass     _PGDisplayModeClass
	PGDisplayModeClassOnce sync.Once
)

func getPGDisplayModeClass() _PGDisplayModeClass {
	PGDisplayModeClassOnce.Do(func() {
		PGDisplayModeClass = _PGDisplayModeClass{objc.GetClass("PGDisplayMode")}
	})
	return PGDisplayModeClass
}

type _PGDisplayModeClass struct {
	class objc.Class
}

// An interface definition for the [PGDisplayMode] class.
type IPGDisplayMode interface {
	objectivec.IObject
}

// A description of a supported display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode
type PGDisplayMode struct {
	objectivec.Object
}

// PGDisplayModeFrom constructs a [PGDisplayMode] from an unsafe.Pointer.
//
// A description of a supported display mode.
func PGDisplayModeFrom(ptr unsafe.Pointer) PGDisplayMode {
	return PGDisplayMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PGDisplayModeClass) Alloc() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PGDisplayModeClass) New() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PGDisplayMode) Init() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PGDisplayMode) Autorelease() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDisplayMode creates a new PGDisplayMode instance.
func NewPGDisplayMode() PGDisplayMode {
	return getPGDisplayModeClass().New()
}


// Creates a new display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/init(sizeInPixels:refreshRateInHz:)
func NewPGDisplayModeWithSizeInPixelsRefreshRateInHz(sizeInPixels unsafe.Pointer, refreshRateInHz unsafe.Pointer) PGDisplayMode {
	instance := getPGDisplayModeClass().Alloc()
	rv := objc.Send[PGDisplayMode](instance.ID, objc.Sel("initWithSizeInPixels:refreshRateInHz:"), sizeInPixels, refreshRateInHz)
	rv.Autorelease()
	return rv
}


// The mode’s refresh rate.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/refreshRate
func (p_ PGDisplayMode) RefreshRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("refreshRate"))
	return rv
}

// The display mode’s dimensions in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/sizeInPixels
func (p_ PGDisplayMode) SizeInPixels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sizeInPixels"))
	return rv
}


