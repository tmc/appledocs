// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechLanguageModel */

/* debug [class_header]: Header for SFSpeechLanguageModel */
// The class instance for the [SFSpeechLanguageModel] class.
var (
	SFSpeechLanguageModelClass     _SFSpeechLanguageModelClass
	SFSpeechLanguageModelClassOnce sync.Once
)

func getSFSpeechLanguageModelClass() _SFSpeechLanguageModelClass {
	SFSpeechLanguageModelClassOnce.Do(func() {
		SFSpeechLanguageModelClass = _SFSpeechLanguageModelClass{objc.GetClass("SFSpeechLanguageModel")}
	})
	return SFSpeechLanguageModelClass
}

type _SFSpeechLanguageModelClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechLanguageModel */
// An interface definition for the [SFSpeechLanguageModel] class.
type ISFSpeechLanguageModel interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFSpeechLanguageModel */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechLanguageModel */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechLanguageModel */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechLanguageModelClass) Alloc() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSpeechLanguageModelClass) New() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechLanguageModel) Init() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechLanguageModel) Autorelease() SFSpeechLanguageModel {
	rv := objc.Send[SFSpeechLanguageModel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechLanguageModel creates a new SFSpeechLanguageModel instance.
func NewSFSpeechLanguageModel() SFSpeechLanguageModel {
	return getSFSpeechLanguageModelClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechLanguageModel */
// A language model built from custom training data.
//
// Create this object using or .

// A language model built from custom training data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel
type SFSpeechLanguageModel struct {
	objectivec.Object
}

// SFSpeechLanguageModelFrom constructs a [SFSpeechLanguageModel] from an unsafe.Pointer.
//
// A language model built from custom training data.
func SFSpeechLanguageModelFrom(ptr unsafe.Pointer) SFSpeechLanguageModel {
	return SFSpeechLanguageModel{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechLanguageModel */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechLanguageModel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:clientIdentifier:configuration:completion:)
func (sc _SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlClientIdentifierConfigurationCompletion(asset objc.IObject /* cross-framework: NSURL */, clientIdentifier objc.IObject /* cross-framework: NSString */, configuration ISFSpeechLanguageModelConfiguration, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("prepareCustomLanguageModelForUrl:clientIdentifier:configuration:completion:"), asset, clientIdentifier, configuration, completion)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=PrepareCustomLanguageModelForUrlClientIdentifierConfigurationCompletion) */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:clientIdentifier:configuration:ignoresCache:completion:)
func (sc _SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlClientIdentifierConfigurationIgnoresCacheCompletion(asset objc.IObject /* cross-framework: NSURL */, clientIdentifier objc.IObject /* cross-framework: NSString */, configuration ISFSpeechLanguageModelConfiguration, ignoresCache bool, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("prepareCustomLanguageModelForUrl:clientIdentifier:configuration:ignoresCache:completion:"), asset, clientIdentifier, configuration, ignoresCache, completion)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=PrepareCustomLanguageModelForUrlClientIdentifierConfigurationIgnoresCacheCompletion) */

// Creates a language model from custom training data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:configuration:completion:)
func (sc _SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationCompletion(asset objc.IObject /* cross-framework: NSURL */, configuration ISFSpeechLanguageModelConfiguration, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("prepareCustomLanguageModelForUrl:configuration:completion:"), asset, configuration, completion)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=PrepareCustomLanguageModelForUrlConfigurationCompletion) */

// Creates a language model from custom training data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechLanguageModel/prepareCustomLanguageModel(for:configuration:ignoresCache:completion:)
func (sc _SFSpeechLanguageModelClass) PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion(asset objc.IObject /* cross-framework: NSURL */, configuration ISFSpeechLanguageModelConfiguration, ignoresCache bool, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("prepareCustomLanguageModelForUrl:configuration:ignoresCache:completion:"), asset, configuration, ignoresCache, completion)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=PrepareCustomLanguageModelForUrlConfigurationIgnoresCacheCompletion) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechLanguageModel */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechLanguageModel */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechLanguageModel */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechLanguageModel */
