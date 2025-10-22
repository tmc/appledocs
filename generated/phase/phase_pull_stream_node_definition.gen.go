// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
)

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

// An interface definition for the [PHASEPullStreamNodeDefinition] class.
type IPHASEPullStreamNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
	Format() avfaudio.AudioFormat
	Normalize() bool
	SetNormalize(value bool)
}

//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEPullStreamNodeDefinitionClass) Alloc() PHASEPullStreamNodeDefinition {
	rv := objc.Send[PHASEPullStreamNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/init(mixerDefinition:format:)
func NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormat(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat) PHASEPullStreamNodeDefinition {
	instance := getPHASEPullStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPullStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:"), mixerDefinition, format)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/init(mixerDefinition:format:identifier:)
func NewPHASEPullStreamNodeDefinitionWithMixerDefinitionFormatIdentifier(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat, identifier string) PHASEPullStreamNodeDefinition {
	instance := getPHASEPullStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPullStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:identifier:"), mixerDefinition, format, objc.String(identifier))
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/format
func (p_ PHASEPullStreamNodeDefinition) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/normalize
func (p_ PHASEPullStreamNodeDefinition) Normalize() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("normalize"))
	return rv
}


// SetNormalize sets the value of the normalize property.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNodeDefinition/normalize
func (p_ PHASEPullStreamNodeDefinition) SetNormalize(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalize:"), value)
}


