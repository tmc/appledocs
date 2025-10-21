// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariExtension] class.
var (
	SFSafariExtensionClass     _SFSafariExtensionClass
	SFSafariExtensionClassOnce sync.Once
)

func getSFSafariExtensionClass() _SFSafariExtensionClass {
	SFSafariExtensionClassOnce.Do(func() {
		SFSafariExtensionClass = _SFSafariExtensionClass{objc.GetClass("SFSafariExtension")}
	})
	return SFSafariExtensionClass
}

type _SFSafariExtensionClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariExtension] class.
type ISFSafariExtension interface {
	objectivec.IObject
}

// A proxy for the Safari extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtension
type SFSafariExtension struct {
	objectivec.Object
}

// SFSafariExtensionFrom constructs a [SFSafariExtension] from an unsafe.Pointer.
//
// A proxy for the Safari extension.
func SFSafariExtensionFrom(ptr unsafe.Pointer) SFSafariExtension {
	return SFSafariExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionClass) Alloc() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariExtensionClass) New() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtension) Init() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtension) Autorelease() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtension creates a new SFSafariExtension instance.
func NewSFSafariExtension() SFSafariExtension {
	return getSFSafariExtensionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtension/getBaseURI(completionHandler:)
func (sc _SFSafariExtensionClass) GetBaseURIWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getBaseURIWithCompletionHandler:"), completionHandler)
}



