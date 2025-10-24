// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLModel */


/* debug [class_header]: Header for NLModel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Model */
// An interface definition for the [Model] class.
type IModel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Model */
	// properties:
	Configuration() INLModelConfiguration
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Model */
	// methods:
	PredictedLabelForString(string_ objc.IObject /* cross-framework: NSString */) foundation.String
	PredictedLabelHypothesesForStringMaximumCount(string_ objc.IObject /* cross-framework: NSString */, maximumCount uint) foundation.IDictionary
	PredictedLabelHypothesesForTokensMaximumCount(tokens []string, maximumCount uint) foundation.IDictionary
	PredictedLabelsForTokens(tokens []string) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Model */
// Alloc allocates a new instance without initialization.
func (mc _ModelClass) Alloc() Model {
	rv := objc.Send[Model](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Model */
// A custom model trained to classify or tag natural language text.
//
// With , you can create text classifier ( ) or word tagger ( ) models. Use to integrate those models into your app. This integration ensures that your tokenization and tagger configurations are identical when you train your model and use it in your app. If you create a text classifier as described in doc:creating-a-text-classifier-model , you can integrate that model into your app and use it to make predictions like this: If you create a custom word tagger as described in , you can integrate that model into your app and generate tags for new text input like this:


// A custom model trained to classify or tag natural language text.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Model */

// Creates a new natural language model based on a compiled Core ML model at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelWithContentsOfURLError */


// Creates a new natural language model based on the given Core ML model instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func NewModelWithMLModelError(mlModel IModel, error_ objectivec.IObject) Model {
	rv := objc.Send[Model](objc.ID(getModelClass().class), objc.Sel("modelWithMLModel:error:"), mlModel, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelWithMLModelError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Model */

// Creates a new natural language model based on a compiled Core ML model at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func (mc _ModelClass) ModelWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelWithContentsOfURLError) */


// Creates a new natural language model based on the given Core ML model instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func (mc _ModelClass) ModelWithMLModelError(mlModel IModel, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelWithMLModel:error:"), mlModel, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelWithMLModelError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Model */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Model */

// Predicts a label for the given input string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabel(for:)
func (m_ Model) PredictedLabelForString(string_ objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("predictedLabelForString:"), string_)
	return rv
}/* debug [instance_methods/method]: PredictedLabelForString */


// Predicts multiple possible labels for the given input string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForString:maximumCount:
func (m_ Model) PredictedLabelHypothesesForStringMaximumCount(string_ objc.IObject /* cross-framework: NSString */, maximumCount uint) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("predictedLabelHypothesesForString:maximumCount:"), string_, maximumCount)
	return rv
}/* debug [instance_methods/method]: PredictedLabelHypothesesForStringMaximumCount */


// Predicts multiple possible labels for each string in the given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForTokens:maximumCount:
func (m_ Model) PredictedLabelHypothesesForTokensMaximumCount(tokens []string, maximumCount uint) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("predictedLabelHypothesesForTokens:maximumCount:"), tokens, maximumCount)
	return rv
}/* debug [instance_methods/method]: PredictedLabelHypothesesForTokensMaximumCount */


// Predicts a label for each string in the given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabels(forTokens:)
func (m_ Model) PredictedLabelsForTokens(tokens []string) []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("predictedLabelsForTokens:"), tokens)
	return rv
}/* debug [instance_methods/method]: PredictedLabelsForTokens */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Model */

// A configuration describing the natural language model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/configuration
func (m_ Model) Configuration() INLModelConfiguration {
	rv := objc.Send[ModelConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLModel */


