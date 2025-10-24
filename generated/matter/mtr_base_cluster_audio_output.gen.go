// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterAudioOutput] class.
var (
	MTRBaseClusterAudioOutputClass     _MTRBaseClusterAudioOutputClass
	MTRBaseClusterAudioOutputClassOnce sync.Once
)

func getMTRBaseClusterAudioOutputClass() _MTRBaseClusterAudioOutputClass {
	MTRBaseClusterAudioOutputClassOnce.Do(func() {
		MTRBaseClusterAudioOutputClass = _MTRBaseClusterAudioOutputClass{objc.GetClass("MTRBaseClusterAudioOutput")}
	})
	return MTRBaseClusterAudioOutputClass
}

type _MTRBaseClusterAudioOutputClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterAudioOutput] class.
type IMTRBaseClusterAudioOutput interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterAudioOutput
type MTRBaseClusterAudioOutput struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterAudioOutputFrom constructs a [MTRBaseClusterAudioOutput] from an unsafe.Pointer.
func MTRBaseClusterAudioOutputFrom(ptr unsafe.Pointer) MTRBaseClusterAudioOutput {
	return MTRBaseClusterAudioOutput{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterAudioOutputClass) Alloc() MTRBaseClusterAudioOutput {
	rv := objc.Send[MTRBaseClusterAudioOutput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterAudioOutputClass) New() MTRBaseClusterAudioOutput {
	rv := objc.Send[MTRBaseClusterAudioOutput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterAudioOutput) Init() MTRBaseClusterAudioOutput {
	rv := objc.Send[MTRBaseClusterAudioOutput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterAudioOutput) Autorelease() MTRBaseClusterAudioOutput {
	rv := objc.Send[MTRBaseClusterAudioOutput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterAudioOutput creates a new MTRBaseClusterAudioOutput instance.
func NewMTRBaseClusterAudioOutput() MTRBaseClusterAudioOutput {
	return getMTRBaseClusterAudioOutputClass().New()
}




