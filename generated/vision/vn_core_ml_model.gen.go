// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CoreMLModel] class.
var (
	CoreMLModelClass     _CoreMLModelClass
	CoreMLModelClassOnce sync.Once
)

func getCoreMLModelClass() _CoreMLModelClass {
	CoreMLModelClassOnce.Do(func() {
		CoreMLModelClass = _CoreMLModelClass{objc.GetClass("VNCoreMLModel")}
	})
	return CoreMLModelClass
}

type _CoreMLModelClass struct {
	class objc.Class
}

// An interface definition for the [CoreMLModel] class.
type ICoreMLModel interface {
	objectivec.IObject
	FeatureProvider() objc.ID
	SetFeatureProvider(value objc.ID)
	InputImageFeatureName() string
	SetInputImageFeatureName(value string)
	Model() VNCoreMLModel
	SetModel(value IVNCoreMLModel)
}

// A container for the model to use with Vision requests.
//
// A model encapsulates the information trained from a data set used to drive Vision recognition requests. See for instructions on training your own model. Once you train the model, use this class to initialize a for identification.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel
type CoreMLModel struct {
	objectivec.Object
}

// CoreMLModelFrom constructs a [CoreMLModel] from an unsafe.Pointer.
//
// A container for the model to use with Vision requests.
func CoreMLModelFrom(ptr unsafe.Pointer) CoreMLModel {
	return CoreMLModel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoreMLModelClass) Alloc() CoreMLModel {
	rv := objc.Send[CoreMLModel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoreMLModelClass) New() CoreMLModel {
	rv := objc.Send[CoreMLModel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoreMLModel) Init() CoreMLModel {
	rv := objc.Send[CoreMLModel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoreMLModel) Autorelease() CoreMLModel {
	rv := objc.Send[CoreMLModel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLModel creates a new CoreMLModel instance.
func NewCoreMLModel() CoreMLModel {
	return getCoreMLModelClass().New()
}




// Creates a model container to use with a Core ML request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/init(for:)
func NewCoreMLModelForMLModelError(model coreml.IModel, error_ unsafe.Pointer) CoreMLModel {
	rv := objc.Send[CoreMLModel](objc.ID(getCoreMLModelClass().class), objc.Sel("modelForMLModel:error:"), model, error_)
	return rv
}


// Creates a model container to use with a Core ML request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/init(for:)
func (cc _CoreMLModelClass) ModelForMLModelError(model coreml.IModel, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("modelForMLModel:error:"), model, error_)
	return rv
}

// An optional object to support inputs outside Vision.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/featureProvider
func (c_ CoreMLModel) FeatureProvider() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("featureProvider"))
	return rv
}


// SetFeatureProvider sets the value of the featureProvider property.
// An optional object to support inputs outside Vision.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/featureProvider
func (c_ CoreMLModel) SetFeatureProvider(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureProvider:"), value)
}

// The name of the feature value that Vision sets from the request handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/inputImageFeatureName
func (c_ CoreMLModel) InputImageFeatureName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("inputImageFeatureName"))
	return rv
}


// SetInputImageFeatureName sets the value of the inputImageFeatureName property.
// The name of the feature value that Vision sets from the request handler.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLModel/inputImageFeatureName
func (c_ CoreMLModel) SetInputImageFeatureName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputImageFeatureName:"), objc.String(value))
}

// The model to base the image analysis request on.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/model
func (c_ CoreMLModel) Model() VNCoreMLModel {
	rv := objc.Send[VNCoreMLModel](c_.ID, objc.Sel("model"))
	return rv
}


// SetModel sets the value of the model property.
// The model to base the image analysis request on.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/model
func (c_ CoreMLModel) SetModel(value IVNCoreMLModel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModel:"), value)
}


