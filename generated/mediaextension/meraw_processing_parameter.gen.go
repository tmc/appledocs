// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MERAWProcessingParameter] class.
var (
	MERAWProcessingParameterClass     _MERAWProcessingParameterClass
	MERAWProcessingParameterClassOnce sync.Once
)

func getMERAWProcessingParameterClass() _MERAWProcessingParameterClass {
	MERAWProcessingParameterClassOnce.Do(func() {
		MERAWProcessingParameterClass = _MERAWProcessingParameterClass{objc.GetClass("MERAWProcessingParameter")}
	})
	return MERAWProcessingParameterClass
}

type _MERAWProcessingParameterClass struct {
	class objc.Class
}

// An interface definition for the [MERAWProcessingParameter] class.
type IMERAWProcessingParameter interface {
	objectivec.IObject
}

// An object for the RAW processor to describe each processing parameter the processor exposes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter
type MERAWProcessingParameter struct {
	objectivec.Object
}

// MERAWProcessingParameterFrom constructs a [MERAWProcessingParameter] from an unsafe.Pointer.
//
// An object for the RAW processor to describe each processing parameter the processor exposes.
func MERAWProcessingParameterFrom(ptr unsafe.Pointer) MERAWProcessingParameter {
	return MERAWProcessingParameter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingParameterClass) Alloc() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MERAWProcessingParameterClass) New() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingParameter) Init() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingParameter) Autorelease() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingParameter creates a new MERAWProcessingParameter instance.
func NewMERAWProcessingParameter() MERAWProcessingParameter {
	return getMERAWProcessingParameterClass().New()
}




