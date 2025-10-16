
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HapticFeedbackManager] class.
var HapticFeedbackManagerClass _HapticFeedbackManagerClass

func init() {
	HapticFeedbackManagerClass = _HapticFeedbackManagerClass{objc.GetClass("NSHapticFeedbackManager")}
}

type _HapticFeedbackManagerClass struct {
	objc.Class
}

// An interface definition for the [HapticFeedbackManager] class.
type IHapticFeedbackManager interface {
	ID() objc.ID
}

type HapticFeedbackManager struct {
	id objc.ID
}

func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ HapticFeedbackManager) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _HapticFeedbackManagerClass) Alloc() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _HapticFeedbackManagerClass) New() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewHapticFeedbackManager creates and returns a new initialized instance.
func NewHapticFeedbackManager() HapticFeedbackManager {
	return HapticFeedbackManagerClass.New()
}

// Init initializes the instance.
func (h_ HapticFeedbackManager) Init() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](h_.ID(), selInit)
	return rv
}
