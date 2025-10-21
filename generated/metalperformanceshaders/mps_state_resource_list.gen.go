// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/metal"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StateResourceList] class.
var (
	StateResourceListClass     _StateResourceListClass
	StateResourceListClassOnce sync.Once
)

func getStateResourceListClass() _StateResourceListClass {
	StateResourceListClassOnce.Do(func() {
		StateResourceListClass = _StateResourceListClass{objc.GetClass("MPSStateResourceList")}
	})
	return StateResourceListClass
}

type _StateResourceListClass struct {
	class objc.Class
}

// An interface definition for the [StateResourceList] class.
type IStateResourceList interface {
	objectivec.IObject
	AppendBuffer(size uint)
}

// An interface for objects that define resources for Metal Performance Shaders state containers.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList
type StateResourceList struct {
	objectivec.Object
}

// StateResourceListFrom constructs a [StateResourceList] from an unsafe.Pointer.
//
// An interface for objects that define resources for Metal Performance Shaders state containers.
func StateResourceListFrom(ptr unsafe.Pointer) StateResourceList {
	return StateResourceList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StateResourceListClass) Alloc() StateResourceList {
	rv := objc.Send[StateResourceList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StateResourceListClass) New() StateResourceList {
	rv := objc.Send[StateResourceList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateResourceList) Init() StateResourceList {
	rv := objc.Send[StateResourceList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateResourceList) Autorelease() StateResourceList {
	rv := objc.Send[StateResourceList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateResourceList creates a new StateResourceList instance.
func NewStateResourceList() StateResourceList {
	return getStateResourceListClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/resourceListWithTextureDescriptors:
func (sc _StateResourceListClass) ResourceListWithTextureDescriptors(d metal.ITextureDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("resourceListWithTextureDescriptors:"), d)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/appendBuffer(_:)
func (s_ StateResourceList) AppendBuffer(size uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendBuffer:"), size)
}



