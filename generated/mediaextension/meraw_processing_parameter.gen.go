// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object for the RAW processor to describe each processing parameter the processor exposes.
//
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


// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) ProcessingParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("processingParameters"))
	return rv
}


// SetProcessingParameters sets the value of the processingParameters property.
// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) SetProcessingParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProcessingParameters:"), value)
}

// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/longdescription
func (m_ MERAWProcessingParameter) LongDescription() string {
	rv := objc.Send[string](m_.ID, objc.Sel("longDescription"))
	return rv
}


// SetLongDescription sets the value of the longDescription property.
// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/longdescription
func (m_ MERAWProcessingParameter) SetLongDescription(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLongDescription:"), objc.String(value))
}

// A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/name
func (m_ MERAWProcessingParameter) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// A localized human-readable name for the parameter, suitable for displaying in application UI.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/name
func (m_ MERAWProcessingParameter) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// A unique key string identifying the parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/key
func (m_ MERAWProcessingParameter) Key() string {
	rv := objc.Send[string](m_.ID, objc.Sel("key"))
	return rv
}


// SetKey sets the value of the key property.
// A unique key string identifying the parameter.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/key
func (m_ MERAWProcessingParameter) SetKey(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), objc.String(value))
}

// A Boolean value that indicates whether the extension enables the parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/enabled
func (m_ MERAWProcessingParameter) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether the extension enables the parameter.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/enabled
func (m_ MERAWProcessingParameter) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}



