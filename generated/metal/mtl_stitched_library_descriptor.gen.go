// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [StitchedLibraryDescriptor] class.
var (
	StitchedLibraryDescriptorClass     _StitchedLibraryDescriptorClass
	StitchedLibraryDescriptorClassOnce sync.Once
)

func getStitchedLibraryDescriptorClass() _StitchedLibraryDescriptorClass {
	StitchedLibraryDescriptorClassOnce.Do(func() {
		StitchedLibraryDescriptorClass = _StitchedLibraryDescriptorClass{objc.GetClass("MTLStitchedLibraryDescriptor")}
	})
	return StitchedLibraryDescriptorClass
}

type _StitchedLibraryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [StitchedLibraryDescriptor] class.
type IStitchedLibraryDescriptor interface {
	objectivec.IObject
	

	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	FunctionGraphs() []FunctionStitchingGraph
	SetFunctionGraphs(value []FunctionStitchingGraph)
	Functions() []objc.ID
	SetFunctions(value []objc.ID)
	Options() StitchedLibraryOptions
	SetOptions(value StitchedLibraryOptions)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _StitchedLibraryDescriptorClass) Alloc() StitchedLibraryDescriptor {
	rv := objc.Send[StitchedLibraryDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StitchedLibraryDescriptorClass) New() StitchedLibraryDescriptor {
	rv := objc.Send[StitchedLibraryDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StitchedLibraryDescriptor) Init() StitchedLibraryDescriptor {
	rv := objc.Send[StitchedLibraryDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StitchedLibraryDescriptor) Autorelease() StitchedLibraryDescriptor {
	rv := objc.Send[StitchedLibraryDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStitchedLibraryDescriptor creates a new StitchedLibraryDescriptor instance.
func NewStitchedLibraryDescriptor() StitchedLibraryDescriptor {
	return getStitchedLibraryDescriptorClass().New()
}





// A description of a new library of procedurally generated functions.
//
// An describes a library of new stitched functions. A is a visible function you create by composing other Metal shader functions together in a function graph. Configure a stitched library descriptor by assigning an array of one or more instances, each describing a stitched function, to the property. Then assign an array that includes all the functions the graphs depend on to the property. Create a stitched library from the descriptor by passing it to the method of an . You can change the descriptor to create other libraries without affecting any existing ones.


// A description of a new library of procedurally generated functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor
type StitchedLibraryDescriptor struct {
	objectivec.Object
}

// StitchedLibraryDescriptorFrom constructs a [StitchedLibraryDescriptor] from an unsafe.Pointer.
//
// A description of a new library of procedurally generated functions.
func StitchedLibraryDescriptorFrom(ptr unsafe.Pointer) StitchedLibraryDescriptor {
	return StitchedLibraryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/binaryArchives
func (s_ StitchedLibraryDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("binaryArchives"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/binaryArchives
func (s_ StitchedLibraryDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}


// The function graphs that define the new stitched library’s functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functionGraphs
func (s_ StitchedLibraryDescriptor) FunctionGraphs() []FunctionStitchingGraph {
	rv := objc.Send[[]FunctionStitchingGraph](s_.ID, objc.Sel("functionGraphs"))
	return rv
}


// The function graphs that define the new stitched library’s functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functionGraphs
func (s_ StitchedLibraryDescriptor) SetFunctionGraphs(value []FunctionStitchingGraph) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setFunctionGraphs:"), nsArray)
}


// The list of functions for creating the stitched library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functions
func (s_ StitchedLibraryDescriptor) Functions() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("functions"))
	return rv
}


// The list of functions for creating the stitched library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functions
func (s_ StitchedLibraryDescriptor) SetFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setFunctions:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/options
func (s_ StitchedLibraryDescriptor) Options() StitchedLibraryOptions {
	rv := objc.Send[StitchedLibraryOptions](s_.ID, objc.Sel("options"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/options
func (s_ StitchedLibraryDescriptor) SetOptions(value StitchedLibraryOptions) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOptions:"), value)
}








