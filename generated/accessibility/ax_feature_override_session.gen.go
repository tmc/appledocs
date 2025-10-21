// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXFeatureOverrideSession] class.
var (
	AXFeatureOverrideSessionClass     _AXFeatureOverrideSessionClass
	AXFeatureOverrideSessionClassOnce sync.Once
)

func getAXFeatureOverrideSessionClass() _AXFeatureOverrideSessionClass {
	AXFeatureOverrideSessionClassOnce.Do(func() {
		AXFeatureOverrideSessionClass = _AXFeatureOverrideSessionClass{objc.GetClass("AXFeatureOverrideSession")}
	})
	return AXFeatureOverrideSessionClass
}

type _AXFeatureOverrideSessionClass struct {
	class objc.Class
}

// An interface definition for the [AXFeatureOverrideSession] class.
type IAXFeatureOverrideSession interface {
	objectivec.IObject
}

// A token object that represents an override session held by your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession
type AXFeatureOverrideSession struct {
	objectivec.Object
}

// AXFeatureOverrideSessionFrom constructs a [AXFeatureOverrideSession] from an unsafe.Pointer.
//
// A token object that represents an override session held by your app.
func AXFeatureOverrideSessionFrom(ptr unsafe.Pointer) AXFeatureOverrideSession {
	return AXFeatureOverrideSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXFeatureOverrideSessionClass) Alloc() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXFeatureOverrideSessionClass) New() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXFeatureOverrideSession) Init() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXFeatureOverrideSession) Autorelease() AXFeatureOverrideSession {
	rv := objc.Send[AXFeatureOverrideSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXFeatureOverrideSession creates a new AXFeatureOverrideSession instance.
func NewAXFeatureOverrideSession() AXFeatureOverrideSession {
	return getAXFeatureOverrideSessionClass().New()
}




