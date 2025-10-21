// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DictionaryFeatureProvider] class.
var (
	DictionaryFeatureProviderClass     _DictionaryFeatureProviderClass
	DictionaryFeatureProviderClassOnce sync.Once
)

func getDictionaryFeatureProviderClass() _DictionaryFeatureProviderClass {
	DictionaryFeatureProviderClassOnce.Do(func() {
		DictionaryFeatureProviderClass = _DictionaryFeatureProviderClass{objc.GetClass("MLDictionaryFeatureProvider")}
	})
	return DictionaryFeatureProviderClass
}

type _DictionaryFeatureProviderClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryFeatureProvider] class.
type IDictionaryFeatureProvider interface {
	objectivec.IObject
	ObjectForKeyedSubscript(featureName appkit.string) FeatureValue
}

// A convenience wrapper for the given dictionary of data.
//
// If your input data is stored in a dictionary, consider this type of that is backed by a dictionary. It is a convenience interface, saving you the trouble of iterating through the dictionary to assign all of its values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider
type DictionaryFeatureProvider struct {
	objectivec.Object
}

// DictionaryFeatureProviderFrom constructs a [DictionaryFeatureProvider] from an unsafe.Pointer.
//
// A convenience wrapper for the given dictionary of data.
func DictionaryFeatureProviderFrom(ptr unsafe.Pointer) DictionaryFeatureProvider {
	return DictionaryFeatureProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryFeatureProviderClass) Alloc() DictionaryFeatureProvider {
	rv := objc.Send[DictionaryFeatureProvider](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DictionaryFeatureProviderClass) New() DictionaryFeatureProvider {
	rv := objc.Send[DictionaryFeatureProvider](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DictionaryFeatureProvider) Init() DictionaryFeatureProvider {
	rv := objc.Send[DictionaryFeatureProvider](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DictionaryFeatureProvider) Autorelease() DictionaryFeatureProvider {
	rv := objc.Send[DictionaryFeatureProvider](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionaryFeatureProvider creates a new DictionaryFeatureProvider instance.
func NewDictionaryFeatureProvider() DictionaryFeatureProvider {
	return getDictionaryFeatureProviderClass().New()
}




// Creates the feature provider based on a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/init(dictionary:)
func NewDictionaryFeatureProviderWithDictionaryError(dictionary unsafe.Pointer, error_ unsafe.Pointer) DictionaryFeatureProvider {
	instance := getDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[DictionaryFeatureProvider](instance.ID, objc.Sel("initWithDictionary:error:"), dictionary, error_)
	rv.Autorelease()
	return rv
}


// Subscript interface for the feature provider to pass through to the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/subscript(_:)
func (d_ DictionaryFeatureProvider) ObjectForKeyedSubscript(featureName appkit.string) FeatureValue {
	rv := objc.Send[FeatureValue](d_.ID, objc.Sel("objectForKeyedSubscript:"), featureName)
	return rv
}

// The backing dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/dictionary
func (d_ DictionaryFeatureProvider) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dictionary"))
	return rv
}


