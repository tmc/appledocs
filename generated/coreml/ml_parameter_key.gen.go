// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLParameterKey */


/* debug [class_header]: Header for MLParameterKey */
// The class instance for the [ParameterKey] class.
var (
	ParameterKeyClass     _ParameterKeyClass
	ParameterKeyClassOnce sync.Once
)

func getParameterKeyClass() _ParameterKeyClass {
	ParameterKeyClassOnce.Do(func() {
		ParameterKeyClass = _ParameterKeyClass{objc.GetClass("MLParameterKey")}
	})
	return ParameterKeyClass
}

type _ParameterKeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ParameterKey */
// An interface definition for the [ParameterKey] class.
type IParameterKey interface {
	IKey
	
/* debug [class_interface_properties]: Properties for ParameterKey */
	// properties:
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	Parameters() IMLParameterKey
	SetParameters(value IMLParameterKey)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ParameterKey */
	// methods:
	ScopedTo(scope objc.IObject /* cross-framework: NSString */) IParameterKey
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ParameterKey */
// Alloc allocates a new instance without initialization.
func (pc _ParameterKeyClass) Alloc() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterKeyClass) New() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterKey) Init() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterKey) Autorelease() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterKey creates a new ParameterKey instance.
func NewParameterKey() ParameterKey {
	return getParameterKeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ParameterKey */
// The keys for the parameter dictionary in a model configuration or a model update context.
//
// Use an to retrieve a model’s parameter value using: The model’s method The dictionary of an The dictionary of an


// The keys for the parameter dictionary in a model configuration or a model update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey
type ParameterKey struct {
	Key
}

// ParameterKeyFrom constructs a [ParameterKey] from an unsafe.Pointer.
//
// The keys for the parameter dictionary in a model configuration or a model update context.
func ParameterKeyFrom(ptr unsafe.Pointer) ParameterKey {
	return ParameterKey{
		Key: KeyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ParameterKey *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ParameterKey */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ParameterKey */

// The key you use to access the Adam optimizer’s first beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta1
func (pc _ParameterKeyClass) Beta1() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("beta1"))
	return rv
}/* debug [class_properties_class/property]: beta1 */

// The key you use to access the Adam optimizer’s second beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta2
func (pc _ParameterKeyClass) Beta2() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("beta2"))
	return rv
}/* debug [class_properties_class/property]: beta2 */

// The key you use to access the biases of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/biases
func (pc _ParameterKeyClass) Biases() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("biases"))
	return rv
}/* debug [class_properties_class/property]: biases */

// The key you use to access the optimizer’s epochs parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/epochs
func (pc _ParameterKeyClass) Epochs() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("epochs"))
	return rv
}/* debug [class_properties_class/property]: epochs */

// The key you use to access the Adam optimizer’s epsilon parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/eps
func (pc _ParameterKeyClass) Eps() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("eps"))
	return rv
}/* debug [class_properties_class/property]: eps */

// The key you use to access the optimizer’s learning rate parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/learningRate
func (pc _ParameterKeyClass) LearningRate() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("learningRate"))
	return rv
}/* debug [class_properties_class/property]: learningRate */

// The key you use to access the linked model’s filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelFileName
func (pc _ParameterKeyClass) LinkedModelFileName() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("linkedModelFileName"))
	return rv
}/* debug [class_properties_class/property]: linkedModelFileName */

// The key you use to access the linked model’s search path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelSearchPath
func (pc _ParameterKeyClass) LinkedModelSearchPath() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("linkedModelSearchPath"))
	return rv
}/* debug [class_properties_class/property]: linkedModelSearchPath */

// The key you use to access the optimizer’s mini batch-size parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/miniBatchSize
func (pc _ParameterKeyClass) MiniBatchSize() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("miniBatchSize"))
	return rv
}/* debug [class_properties_class/property]: miniBatchSize */

// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/momentum
func (pc _ParameterKeyClass) Momentum() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("momentum"))
	return rv
}/* debug [class_properties_class/property]: momentum */

// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/numberOfNeighbors
func (pc _ParameterKeyClass) NumberOfNeighbors() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("numberOfNeighbors"))
	return rv
}/* debug [class_properties_class/property]: numberOfNeighbors */

// The key you use to access the seed parameter that initializes the random number generator for the shuffle option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/seed
func (pc _ParameterKeyClass) Seed() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("seed"))
	return rv
}/* debug [class_properties_class/property]: seed */

// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/shuffle
func (pc _ParameterKeyClass) Shuffle() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("shuffle"))
	return rv
}/* debug [class_properties_class/property]: shuffle */

// The key you use to access the weights of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/weights
func (pc _ParameterKeyClass) Weights() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("weights"))
	return rv
}/* debug [class_properties_class/property]: weights */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ParameterKey */

// Creates a copy of a parameter key and adds the scope to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/scoped(to:)
func (p_ ParameterKey) ScopedTo(scope objc.IObject /* cross-framework: NSString */) IParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("scopedTo:"), scope)
	return rv
}/* debug [instance_methods/method]: ScopedTo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ParameterKey */

// The key you use to access the Adam optimizer’s first beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta1
func (p_ ParameterKey) Beta1() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("beta1"))
	return rv
}/* debug [instance_properties/getter]: beta1 */


// The key you use to access the Adam optimizer’s second beta parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/beta2
func (p_ ParameterKey) Beta2() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("beta2"))
	return rv
}/* debug [instance_properties/getter]: beta2 */


// The key you use to access the biases of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/biases
func (p_ ParameterKey) Biases() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("biases"))
	return rv
}/* debug [instance_properties/getter]: biases */


// The key you use to access the optimizer’s epochs parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/epochs
func (p_ ParameterKey) Epochs() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("epochs"))
	return rv
}/* debug [instance_properties/getter]: epochs */


// The key you use to access the Adam optimizer’s epsilon parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/eps
func (p_ ParameterKey) Eps() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("eps"))
	return rv
}/* debug [instance_properties/getter]: eps */


// The key you use to access the optimizer’s learning rate parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/learningRate
func (p_ ParameterKey) LearningRate() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("learningRate"))
	return rv
}/* debug [instance_properties/getter]: learningRate */


// The key you use to access the linked model’s filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelFileName
func (p_ ParameterKey) LinkedModelFileName() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("linkedModelFileName"))
	return rv
}/* debug [instance_properties/getter]: linkedModelFileName */


// The key you use to access the linked model’s search path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/linkedModelSearchPath
func (p_ ParameterKey) LinkedModelSearchPath() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("linkedModelSearchPath"))
	return rv
}/* debug [instance_properties/getter]: linkedModelSearchPath */


// The key you use to access the optimizer’s mini batch-size parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/miniBatchSize
func (p_ ParameterKey) MiniBatchSize() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("miniBatchSize"))
	return rv
}/* debug [instance_properties/getter]: miniBatchSize */


// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/momentum
func (p_ ParameterKey) Momentum() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("momentum"))
	return rv
}/* debug [instance_properties/getter]: momentum */


// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/numberOfNeighbors
func (p_ ParameterKey) NumberOfNeighbors() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("numberOfNeighbors"))
	return rv
}/* debug [instance_properties/getter]: numberOfNeighbors */


// The key you use to access the seed parameter that initializes the random number generator for the shuffle option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/seed
func (p_ ParameterKey) Seed() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/shuffle
func (p_ ParameterKey) Shuffle() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("shuffle"))
	return rv
}/* debug [instance_properties/getter]: shuffle */


// The key you use to access the weights of a layer in a neural network model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey/weights
func (p_ ParameterKey) Weights() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (p_ ParameterKey) Configuration() IMLModelConfiguration {
	rv := objc.Send[ModelConfiguration](p_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The configuration of the model set during initialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/configuration
func (p_ ParameterKey) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (p_ ParameterKey) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](p_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (p_ ParameterKey) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (p_ ParameterKey) Parameters() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("parameters"))
	return rv
}/* debug [instance_properties/getter]: parameters */


// A dictionary of configuration settings your app can override when loading a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/parameters
func (p_ ParameterKey) SetParameters(value IMLParameterKey) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameters:"), value)
}/* debug [instance_properties/setter]: parameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLParameterKey */



