// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [StageInputOutputDescriptor] class.
var (
	StageInputOutputDescriptorClass     _StageInputOutputDescriptorClass
	StageInputOutputDescriptorClassOnce sync.Once
)

func getStageInputOutputDescriptorClass() _StageInputOutputDescriptorClass {
	StageInputOutputDescriptorClassOnce.Do(func() {
		StageInputOutputDescriptorClass = _StageInputOutputDescriptorClass{objc.GetClass("MTLStageInputOutputDescriptor")}
	})
	return StageInputOutputDescriptorClass
}

type _StageInputOutputDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [StageInputOutputDescriptor] class.
type IStageInputOutputDescriptor interface {
	objectivec.IObject
}

// A description of the input and output data of a function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor
type StageInputOutputDescriptor struct {
	objectivec.Object
}

// StageInputOutputDescriptorFrom constructs a [StageInputOutputDescriptor] from an unsafe.Pointer.
//
// A description of the input and output data of a function.
func StageInputOutputDescriptorFrom(ptr unsafe.Pointer) StageInputOutputDescriptor {
	return StageInputOutputDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StageInputOutputDescriptorClass) Alloc() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StageInputOutputDescriptorClass) New() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StageInputOutputDescriptor) Init() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StageInputOutputDescriptor) Autorelease() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStageInputOutputDescriptor creates a new StageInputOutputDescriptor instance.
func NewStageInputOutputDescriptor() StageInputOutputDescriptor {
	return getStageInputOutputDescriptorClass().New()
}


// An array that describes where and how to fetch data for the function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/attributes
func (s_ StageInputOutputDescriptor) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("attributes"))
	return rv
}



