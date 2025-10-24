// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ArrayBatchProvider] class.
var (
	ArrayBatchProviderClass     _ArrayBatchProviderClass
	ArrayBatchProviderClassOnce sync.Once
)

func getArrayBatchProviderClass() _ArrayBatchProviderClass {
	ArrayBatchProviderClassOnce.Do(func() {
		ArrayBatchProviderClass = _ArrayBatchProviderClass{objc.GetClass("MLArrayBatchProvider")}
	})
	return ArrayBatchProviderClass
}

type _ArrayBatchProviderClass struct {
	class objc.Class
}

// An interface definition for the [ArrayBatchProvider] class.
type IArrayBatchProvider interface {
	objectivec.IObject
	// properties:
	Array() []objc.ID
	// methods:
}

// A convenience wrapper for batches of feature providers.
//
// This batch provider supports an array of feature providers or a dictionary of arrays of feature values.


// A convenience wrapper for batches of feature providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider
type ArrayBatchProvider struct {
	objectivec.Object
}

// ArrayBatchProviderFrom constructs a [ArrayBatchProvider] from an unsafe.Pointer.
//
// A convenience wrapper for batches of feature providers.
func ArrayBatchProviderFrom(ptr unsafe.Pointer) ArrayBatchProvider {
	return ArrayBatchProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ArrayBatchProviderClass) Alloc() ArrayBatchProvider {
	rv := objc.Send[ArrayBatchProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArrayBatchProviderClass) New() ArrayBatchProvider {
	rv := objc.Send[ArrayBatchProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArrayBatchProvider) Init() ArrayBatchProvider {
	rv := objc.Send[ArrayBatchProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArrayBatchProvider) Autorelease() ArrayBatchProvider {
	rv := objc.Send[ArrayBatchProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArrayBatchProvider creates a new ArrayBatchProvider instance.
func NewArrayBatchProvider() ArrayBatchProvider {
	return getArrayBatchProviderClass().New()
}



// Creates a batch provider based on feature names and their associated arrays of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(dictionary:)
func NewArrayBatchProviderWithDictionaryError(dictionary foundation.IDictionary, error_ unsafe.Pointer) ArrayBatchProvider {
	instance := getArrayBatchProviderClass().Alloc()
	rv := objc.Send[ArrayBatchProvider](instance.ID, objc.Sel("initWithDictionary:error:"), dictionary, error_)
	rv.Autorelease()
	return rv
}


// Creates the batch provider based on the array of feature providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(array:)
func NewArrayBatchProviderWithFeatureProviderArray(array []objc.ID) ArrayBatchProvider {
	instance := getArrayBatchProviderClass().Alloc()
	rv := objc.Send[ArrayBatchProvider](instance.ID, objc.Sel("initWithFeatureProviderArray:"), array)
	rv.Autorelease()
	return rv
}



// The array of feature providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/array
func (a_ ArrayBatchProvider) Array() []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("array"))
	return rv
}


