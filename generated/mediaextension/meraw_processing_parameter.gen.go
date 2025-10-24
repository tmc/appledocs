// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Enabled() bool
	SetEnabled(value bool)
	Key() objc.IObject /* cross-framework: NSString */
	SetKey(value objc.IObject /* cross-framework: NSString */)
	LongDescription() objc.IObject /* cross-framework: NSString */
	SetLongDescription(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	ProcessingParameters() IMERAWProcessingParameter
	SetProcessingParameters(value IMERAWProcessingParameter)
	// methods:
}

// An object for the RAW processor to describe each processing parameter the processor exposes.
//
// This protocol provides an interface for Video Toolbox to query descriptions of the different parameters that can be used to influence RAW processor operation. A distinct is created for each parameter supported by the RAW processor, and the set of supported parameters is returned by the interface.


// An object for the RAW processor to describe each processing parameter the processor exposes.
//
// [Full Topic]
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



// A Boolean value that indicates whether the extension enables the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/enabled
func (m_ MERAWProcessingParameter) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the extension enables the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/enabled
func (m_ MERAWProcessingParameter) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}


// A unique key string identifying the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/key
func (m_ MERAWProcessingParameter) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("key"))
	return rv
}


// A unique key string identifying the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/key
func (m_ MERAWProcessingParameter) SetKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/longdescription
func (m_ MERAWProcessingParameter) LongDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("longDescription"))
	return rv
}


// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/longdescription
func (m_ MERAWProcessingParameter) SetLongDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLongDescription:"), value)
}


// A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/name
func (m_ MERAWProcessingParameter) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/name
func (m_ MERAWProcessingParameter) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) ProcessingParameters() IMERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("processingParameters"))
	return rv
}


// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) SetProcessingParameters(value IMERAWProcessingParameter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProcessingParameters:"), value)
}



