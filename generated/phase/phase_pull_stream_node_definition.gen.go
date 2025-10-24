// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEPullStreamNodeDefinition */


/* debug [class_header]: Header for PHASEPullStreamNodeDefinition */
// The class instance for the [PHASEPullStreamNodeDefinition] class.
var (
	PHASEPullStreamNodeDefinitionClass     _PHASEPullStreamNodeDefinitionClass
	PHASEPullStreamNodeDefinitionClassOnce sync.Once
)

func getPHASEPullStreamNodeDefinitionClass() _PHASEPullStreamNodeDefinitionClass {
	PHASEPullStreamNodeDefinitionClassOnce.Do(func() {
		PHASEPullStreamNodeDefinitionClass = _PHASEPullStreamNodeDefinitionClass{objc.GetClass("PHASEPullStreamNodeDefinition")}
	})
	return PHASEPullStreamNodeDefinitionClass
}

type _PHASEPullStreamNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEPullStreamNodeDefinition */
// An interface definition for the [PHASEPullStreamNodeDefinition] class.
type IPHASEPullStreamNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASEPullStreamNodeDefinition */
	// properties:
	Format() avfaudio.AudioFormat
	Normalize() bool
	SetNormalize(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEPullStreamNodeDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEPullStreamNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEPullStreamNodeDefinitionClass) Alloc() PHASEPullStreamNodeDefinition {
	rv := objc.Send[PHASEPullStreamNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEPullStreamNodeDefinitionClass) New() PHASEPullStreamNodeDefinition {
	rv := objc.Send[PHASEPullStreamNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPullStreamNodeDefinition) Init() PHASEPullStreamNodeDefinition {
	rv := objc.Send[PHASEPullStreamNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPullStreamNodeDefinition) Autorelease() PHASEPullStreamNodeDefinition {
	rv := objc.Send[PHASEPullStreamNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPullStreamNodeDefinition creates a new PHASEPullStreamNodeDefinition instance.
func NewPHASEPullStreamNodeDefinition() PHASEPullStreamNodeDefinition {
	return getPHASEPullStreamNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEPullStreamNodeDefinition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition
type PHASEPullStreamNodeDefinition struct {
	PHASEGeneratorNodeDefinition
}

// PHASEPullStreamNodeDefinitionFrom constructs a [PHASEPullStreamNodeDefinition] from an unsafe.Pointer.
func PHASEPullStreamNodeDefinitionFrom(ptr unsafe.Pointer) PHASEPullStreamNodeDefinition {
	return PHASEPullStreamNodeDefinition{
		PHASEGeneratorNodeDefinition: PHASEGeneratorNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEPullStreamNodeDefinition */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/init(mixerDefinition:format:)
func NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormat(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat) PHASEPullStreamNodeDefinition {
	instance := getPHASEPullStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPullStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:"), mixerDefinition, format)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/init(mixerDefinition:format:identifier:)
func NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormatIdentifier(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat, identifier objc.IObject /* cross-framework: NSString */) PHASEPullStreamNodeDefinition {
	instance := getPHASEPullStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPullStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:identifier:"), mixerDefinition, format, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormatIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEPullStreamNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEPullStreamNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEPullStreamNodeDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEPullStreamNodeDefinition */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/format
func (p_ PHASEPullStreamNodeDefinition) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/normalize
func (p_ PHASEPullStreamNodeDefinition) Normalize() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("normalize"))
	return rv
}/* debug [instance_properties/getter]: normalize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/normalize
func (p_ PHASEPullStreamNodeDefinition) SetNormalize(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalize:"), value)
}/* debug [instance_properties/setter]: normalize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEPullStreamNodeDefinition */


