// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLDictionaryFeatureProvider */


/* debug [class_header]: Header for MLDictionaryFeatureProvider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DictionaryFeatureProvider */
// An interface definition for the [DictionaryFeatureProvider] class.
type IDictionaryFeatureProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DictionaryFeatureProvider */
	// properties:
	Dictionary() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DictionaryFeatureProvider */
	// methods:
	ObjectForKeyedSubscript(featureName objc.IObject /* cross-framework: NSString */) IFeatureValue
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DictionaryFeatureProvider */
// Alloc allocates a new instance without initialization.
func (dc _DictionaryFeatureProviderClass) Alloc() DictionaryFeatureProvider {
	rv := objc.Send[DictionaryFeatureProvider](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DictionaryFeatureProvider */
// A convenience wrapper for the given dictionary of data.
//
// If your input data is stored in a dictionary, consider this type of that is backed by a dictionary. It is a convenience interface, saving you the trouble of iterating through the dictionary to assign all of its values.


// A convenience wrapper for the given dictionary of data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DictionaryFeatureProvider */

// Creates the feature provider based on a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/init(dictionary:)
func NewDictionaryFeatureProviderWithDictionaryError(dictionary foundation.IDictionary, error_ objectivec.IObject) DictionaryFeatureProvider {
	instance := getDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[DictionaryFeatureProvider](instance.ID, objc.Sel("initWithDictionary:error:"), dictionary, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryFeatureProviderWithDictionaryError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DictionaryFeatureProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DictionaryFeatureProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DictionaryFeatureProvider */

// Subscript interface for the feature provider to pass through to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/subscript(_:)
func (d_ DictionaryFeatureProvider) ObjectForKeyedSubscript(featureName objc.IObject /* cross-framework: NSString */) IFeatureValue {
	rv := objc.Send[FeatureValue](d_.ID, objc.Sel("objectForKeyedSubscript:"), featureName)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DictionaryFeatureProvider */

// The backing dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/dictionary
func (d_ DictionaryFeatureProvider) Dictionary() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](d_.ID, objc.Sel("dictionary"))
	return rv
}/* debug [instance_properties/getter]: dictionary */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLDictionaryFeatureProvider */


