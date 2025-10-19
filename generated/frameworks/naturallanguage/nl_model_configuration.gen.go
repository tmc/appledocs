// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLModelConfiguration] class.
var nLModelConfigurationClass = _NLModelConfigurationClass{objc.GetClass("NLModelConfiguration")}

type _NLModelConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [NLModelConfiguration] class.
type INLModelConfiguration interface {
	objectivec.IObject
}

// The configuration parameters of a natural language model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration

type NLModelConfiguration struct {
	objectivec.Object
}

// NLModelConfigurationFrom constructs a [NLModelConfiguration] from an unsafe.Pointer.
//
// The configuration parameters of a natural language model.
func NLModelConfigurationFrom(ptr unsafe.Pointer) NLModelConfiguration {
	return NLModelConfiguration{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLModelConfigurationClass) Alloc() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLModelConfigurationClass) New() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLModelConfiguration) Init() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLModelConfiguration) Autorelease() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLModelConfiguration creates a new NLModelConfiguration instance.
func NewNLModelConfiguration() NLModelConfiguration {
	return nLModelConfigurationClass.New()
}




