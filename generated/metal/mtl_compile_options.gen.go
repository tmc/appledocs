// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLCompileOptions] class.
var (
	mTLCompileOptionsClass     _MTLCompileOptionsClass
	mTLCompileOptionsClassOnce sync.Once
)

func getMTLCompileOptionsClass() _MTLCompileOptionsClass {
	mTLCompileOptionsClassOnce.Do(func() {
		mTLCompileOptionsClass = _MTLCompileOptionsClass{objc.GetClass("MTLCompileOptions")}
	})
	return mTLCompileOptionsClass
}

type _MTLCompileOptionsClass struct {
	class objc.Class
}

// An interface definition for the [MTLCompileOptions] class.
type IMTLCompileOptions interface {
	objectivec.IObject
}

// Compilation settings for a Metal shader library.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions
type MTLCompileOptions struct {
	objectivec.Object
}

// MTLCompileOptionsFrom constructs a [MTLCompileOptions] from an unsafe.Pointer.
//
// Compilation settings for a Metal shader library.
func MTLCompileOptionsFrom(ptr unsafe.Pointer) MTLCompileOptions {
	return MTLCompileOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTLCompileOptionsClass) Alloc() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTLCompileOptionsClass) New() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTLCompileOptions) Init() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTLCompileOptions) Autorelease() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCompileOptions creates a new MTLCompileOptions instance.
func NewMTLCompileOptions() MTLCompileOptions {
	return getMTLCompileOptionsClass().New()
}




