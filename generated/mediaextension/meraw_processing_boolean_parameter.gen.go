// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MERAWProcessingBooleanParameter] class.
var (
	MERAWProcessingBooleanParameterClass     _MERAWProcessingBooleanParameterClass
	MERAWProcessingBooleanParameterClassOnce sync.Once
)

func getMERAWProcessingBooleanParameterClass() _MERAWProcessingBooleanParameterClass {
	MERAWProcessingBooleanParameterClassOnce.Do(func() {
		MERAWProcessingBooleanParameterClass = _MERAWProcessingBooleanParameterClass{objc.GetClass("MERAWProcessingBooleanParameter")}
	})
	return MERAWProcessingBooleanParameterClass
}

type _MERAWProcessingBooleanParameterClass struct {
	class objc.Class
}

// An interface definition for the [MERAWProcessingBooleanParameter] class.
type IMERAWProcessingBooleanParameter interface {
	IMERAWProcessingParameter
}

// An object that describes a Boolean parameter of a RAW processor.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean
type MERAWProcessingBooleanParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingBooleanParameterFrom constructs a [MERAWProcessingBooleanParameter] from an unsafe.Pointer.
//
// An object that describes a Boolean parameter of a RAW processor.
func MERAWProcessingBooleanParameterFrom(ptr unsafe.Pointer) MERAWProcessingBooleanParameter {
	return MERAWProcessingBooleanParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingBooleanParameterClass) Alloc() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MERAWProcessingBooleanParameterClass) New() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingBooleanParameter) Init() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingBooleanParameter) Autorelease() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingBooleanParameter creates a new MERAWProcessingBooleanParameter instance.
func NewMERAWProcessingBooleanParameter() MERAWProcessingBooleanParameter {
	return getMERAWProcessingBooleanParameterClass().New()
}




