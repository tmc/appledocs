// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramArgument] class.
var (
	ModelStructureProgramArgumentClass     _ModelStructureProgramArgumentClass
	ModelStructureProgramArgumentClassOnce sync.Once
)

func getModelStructureProgramArgumentClass() _ModelStructureProgramArgumentClass {
	ModelStructureProgramArgumentClassOnce.Do(func() {
		ModelStructureProgramArgumentClass = _ModelStructureProgramArgumentClass{objc.GetClass("MLModelStructureProgramArgument")}
	})
	return ModelStructureProgramArgumentClass
}

type _ModelStructureProgramArgumentClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramArgument] class.
type IModelStructureProgramArgument interface {
	objectivec.IObject
}

// A class representing an argument in the Program.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type ModelStructureProgramArgument struct {
	objectivec.Object
}

// ModelStructureProgramArgumentFrom constructs a [ModelStructureProgramArgument] from an unsafe.Pointer.
//
// A class representing an argument in the Program.
func ModelStructureProgramArgumentFrom(ptr unsafe.Pointer) ModelStructureProgramArgument {
	return ModelStructureProgramArgument{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramArgumentClass) Alloc() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramArgumentClass) New() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramArgument) Init() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramArgument) Autorelease() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramArgument creates a new ModelStructureProgramArgument instance.
func NewModelStructureProgramArgument() ModelStructureProgramArgument {
	return getModelStructureProgramArgumentClass().New()
}


// The array of bindings.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/bindings
func (m_ ModelStructureProgramArgument) Bindings() []ModelStructureProgramBinding {
	rv := objc.Send[[]ModelStructureProgramBinding](m_.ID, objc.Sel("bindings"))
	return rv
}



