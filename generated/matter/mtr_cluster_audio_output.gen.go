// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterAudioOutput] class.
var (
	MTRClusterAudioOutputClass     _MTRClusterAudioOutputClass
	MTRClusterAudioOutputClassOnce sync.Once
)

func getMTRClusterAudioOutputClass() _MTRClusterAudioOutputClass {
	MTRClusterAudioOutputClassOnce.Do(func() {
		MTRClusterAudioOutputClass = _MTRClusterAudioOutputClass{objc.GetClass("MTRClusterAudioOutput")}
	})
	return MTRClusterAudioOutputClass
}

type _MTRClusterAudioOutputClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterAudioOutput] class.
type IMTRClusterAudioOutput interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterAudioOutput
type MTRClusterAudioOutput struct {
	MTRGenericCluster
}

// MTRClusterAudioOutputFrom constructs a [MTRClusterAudioOutput] from an unsafe.Pointer.
func MTRClusterAudioOutputFrom(ptr unsafe.Pointer) MTRClusterAudioOutput {
	return MTRClusterAudioOutput{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterAudioOutputClass) Alloc() MTRClusterAudioOutput {
	rv := objc.Send[MTRClusterAudioOutput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterAudioOutputClass) New() MTRClusterAudioOutput {
	rv := objc.Send[MTRClusterAudioOutput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterAudioOutput) Init() MTRClusterAudioOutput {
	rv := objc.Send[MTRClusterAudioOutput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterAudioOutput) Autorelease() MTRClusterAudioOutput {
	rv := objc.Send[MTRClusterAudioOutput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterAudioOutput creates a new MTRClusterAudioOutput instance.
func NewMTRClusterAudioOutput() MTRClusterAudioOutput {
	return getMTRClusterAudioOutputClass().New()
}




