// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Model] class.
var (
	ModelClass     _ModelClass
	ModelClassOnce sync.Once
)

func getModelClass() _ModelClass {
	ModelClassOnce.Do(func() {
		ModelClass = _ModelClass{objc.GetClass("NLModel")}
	})
	return ModelClass
}

type _ModelClass struct {
	class objc.Class
}

// An interface definition for the [Model] class.
type IModel interface {
	objectivec.IObject
	PredictedLabelForString(string_ string) string
	PredictedLabelHypothesesForStringMaximumCount(string_ string, maximumCount uint) unsafe.Pointer
	PredictedLabelHypothesesForTokensMaximumCount(tokens unsafe.Pointer, maximumCount uint) []foundation.NSDictionary
	PredictedLabelsForTokens(tokens unsafe.Pointer) []string
}

// A custom model trained to classify or tag natural language text.
//
// With , you can create text classifier ( ) or word tagger ( ) models. Use to integrate those models into your app. This integration ensures that your tokenization and tagger configurations are identical when you train your model and use it in your app. If you create a text classifier as described in doc:creating-a-text-classifier-model , you can integrate that model into your app and use it to make predictions like this: If you create a custom word tagger as described in , you can integrate that model into your app and generate tags for new text input like this:
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel
type Model struct {
	objectivec.Object
}

// ModelFrom constructs a [Model] from an unsafe.Pointer.
//
// A custom model trained to classify or tag natural language text.
func ModelFrom(ptr unsafe.Pointer) Model {
	return Model{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelClass) Alloc() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelClass) New() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Model) Init() Model {
	rv := objc.Send[Model](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Model) Autorelease() Model {
	rv := objc.Send[Model](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModel creates a new Model instance.
func NewModel() Model {
	return getModelClass().New()
}




// Creates a new natural language model based on a compiled Core ML model at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}



// Creates a new natural language model based on the given Core ML model instance.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func NewModelWithMLModelError(mlModel unsafe.Pointer, error_ unsafe.Pointer) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithMLModel:error:"), mlModel, error_)
	return rv
}


// Creates a new natural language model based on a compiled Core ML model at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func (mc _ModelClass) ModelWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}

// Creates a new natural language model based on the given Core ML model instance.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func (mc _ModelClass) ModelWithMLModelError(mlModel unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelWithMLModel:error:"), mlModel, error_)
	return rv
}

// Predicts a label for the given input string.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabel(for:)
func (m_ Model) PredictedLabelForString(string_ string) string {
	rv := objc.Send[string](m_.ID, objc.Sel("predictedLabelForString:"), objc.String(string_))
	return rv
}

// Predicts multiple possible labels for the given input string.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForString:maximumCount:
func (m_ Model) PredictedLabelHypothesesForStringMaximumCount(string_ string, maximumCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predictedLabelHypothesesForString:maximumCount:"), objc.String(string_), maximumCount)
	return rv
}

// Predicts multiple possible labels for each string in the given array.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForTokens:maximumCount:
func (m_ Model) PredictedLabelHypothesesForTokensMaximumCount(tokens unsafe.Pointer, maximumCount uint) []foundation.NSDictionary {
	rv := objc.Send[[]foundation.NSDictionary](m_.ID, objc.Sel("predictedLabelHypothesesForTokens:maximumCount:"), tokens, maximumCount)
	return rv
}

// Predicts a label for each string in the given array.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabels(forTokens:)
func (m_ Model) PredictedLabelsForTokens(tokens unsafe.Pointer) []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("predictedLabelsForTokens:"), tokens)
	return rv
}

// A configuration describing the natural language model.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/configuration
func (m_ Model) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("configuration"))
	return rv
}


