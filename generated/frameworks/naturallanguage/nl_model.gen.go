// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLModel] class.
var nLModelClass = _NLModelClass{objc.GetClass("NLModel")}

type _NLModelClass struct {
	class objc.Class
}

// An interface definition for the [NLModel] class.
type INLModel interface {
	objectivec.IObject
	PredictedLabelForString(string string) unsafe.Pointer
	PredictedLabelHypothesesForStringMaximumCount(string string, maximumCount uint) unsafe.Pointer
	PredictedLabelHypothesesForTokensMaximumCount(tokens unsafe.Pointer, maximumCount uint) unsafe.Pointer
	PredictedLabelsForTokens(tokens unsafe.Pointer) unsafe.Pointer
}

// A custom model trained to classify or tag natural language text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel

type NLModel struct {
	objectivec.Object
}

// NLModelFrom constructs a [NLModel] from an unsafe.Pointer.
//
// A custom model trained to classify or tag natural language text.
func NLModelFrom(ptr unsafe.Pointer) NLModel {
	return NLModel{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLModelClass) Alloc() NLModel {
	rv := objc.Send[NLModel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLModelClass) New() NLModel {
	rv := objc.Send[NLModel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLModel) Init() NLModel {
	rv := objc.Send[NLModel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLModel) Autorelease() NLModel {
	rv := objc.Send[NLModel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLModel creates a new NLModel instance.
func NewNLModel() NLModel {
	return nLModelClass.New()
}


// Creates a new natural language model based on a compiled Core ML model at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) NLModel {
	rv := objc.Send[NLModel](objc.ID(nLModelClass.class), objc.Sel("modelWithContentsOfURL:error:"), url, error)
	rv.Autorelease()
	return rv
}
// Creates a new natural language model based on the given Core ML model instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func NewModelWithMLModelError(mlModel unsafe.Pointer, error unsafe.Pointer) NLModel {
	rv := objc.Send[NLModel](objc.ID(nLModelClass.class), objc.Sel("modelWithMLModel:error:"), mlModel, error)
	rv.Autorelease()
	return rv
}


// Creates a new natural language model based on a compiled Core ML model at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func (nc _NLModelClass) ModelWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("modelWithContentsOfURL:error:"), url, error)
	return rv
}
// Creates a new natural language model based on the given Core ML model instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func (nc _NLModelClass) ModelWithMLModelError(mlModel unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("modelWithMLModel:error:"), mlModel, error)
	return rv
}
// Predicts a label for the given input string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabel(for:)
func (n_ NLModel) PredictedLabelForString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("predictedLabelForString:"), string)
	return rv
}
// Predicts multiple possible labels for the given input string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForString:maximumCount:
func (n_ NLModel) PredictedLabelHypothesesForStringMaximumCount(string string, maximumCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("predictedLabelHypothesesForString:maximumCount:"), string, maximumCount)
	return rv
}
// Predicts multiple possible labels for each string in the given array. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForTokens:maximumCount:
func (n_ NLModel) PredictedLabelHypothesesForTokensMaximumCount(tokens unsafe.Pointer, maximumCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("predictedLabelHypothesesForTokens:maximumCount:"), tokens, maximumCount)
	return rv
}
// Predicts a label for each string in the given array. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabels(forTokens:)
func (n_ NLModel) PredictedLabelsForTokens(tokens unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("predictedLabelsForTokens:"), tokens)
	return rv
}

