// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var (
	scrubberLayoutAttributesClass     _ScrubberLayoutAttributesClass
	scrubberLayoutAttributesClassOnce sync.Once
)

func getScrubberLayoutAttributesClass() _ScrubberLayoutAttributesClass {
	scrubberLayoutAttributesClassOnce.Do(func() {
		scrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}
	})
	return scrubberLayoutAttributesClass
}

type _ScrubberLayoutAttributesClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberLayoutAttributes] class.
type IScrubberLayoutAttributes interface {
	objectivec.IObject
}

// The layout of a scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes
type ScrubberLayoutAttributes struct {
	objectivec.Object
}

// ScrubberLayoutAttributesFrom constructs a [ScrubberLayoutAttributes] from an unsafe.Pointer.
//
// The layout of a scrubber item.
func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutAttributesClass) Alloc() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberLayoutAttributesClass) New() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberLayoutAttributes) Init() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberLayoutAttributes) Autorelease() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberLayoutAttributes creates a new ScrubberLayoutAttributes instance.
func NewScrubberLayoutAttributes() ScrubberLayoutAttributes {
	return getScrubberLayoutAttributesClass().New()
}




