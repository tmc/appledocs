// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AEAssessmentParticipantConfiguration] class.
var (
	AEAssessmentParticipantConfigurationClass     _AEAssessmentParticipantConfigurationClass
	AEAssessmentParticipantConfigurationClassOnce sync.Once
)

func getAEAssessmentParticipantConfigurationClass() _AEAssessmentParticipantConfigurationClass {
	AEAssessmentParticipantConfigurationClassOnce.Do(func() {
		AEAssessmentParticipantConfigurationClass = _AEAssessmentParticipantConfigurationClass{objc.GetClass("AEAssessmentParticipantConfiguration")}
	})
	return AEAssessmentParticipantConfigurationClass
}

type _AEAssessmentParticipantConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [AEAssessmentParticipantConfiguration] class.
type IAEAssessmentParticipantConfiguration interface {
	objectivec.IObject
	AllowsNetworkAccess() bool
	SetAllowsNetworkAccess(value bool)
	ConfigurationInfo() unsafe.Pointer
	SetConfigurationInfo(value unsafe.Pointer)
	Required() bool
	SetRequired(value bool)
	ConfigurationsByApplication() AEAssessmentParticipantConfiguration
	SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration)
	MainParticipantConfiguration() AEAssessmentParticipantConfiguration
	SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration)
	IsRequired() bool
	SetIsRequired(value bool)
}

// Configuration information for an app that’s available during an assessment.
//
// Use an instance of this class to configure the properties of an app that you allow to run during an assessment. Associate the participant configuration with an app (an instance) when you call the method of a session configuration.


// Configuration information for an app that’s available during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration

type AEAssessmentParticipantConfiguration struct {
	objectivec.Object
}

// AEAssessmentParticipantConfigurationFrom constructs a [AEAssessmentParticipantConfiguration] from an unsafe.Pointer.
//
// Configuration information for an app that’s available during an assessment.
func AEAssessmentParticipantConfigurationFrom(ptr unsafe.Pointer) AEAssessmentParticipantConfiguration {
	return AEAssessmentParticipantConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentParticipantConfigurationClass) Alloc() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AEAssessmentParticipantConfigurationClass) New() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AEAssessmentParticipantConfiguration) Init() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AEAssessmentParticipantConfiguration) Autorelease() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEAssessmentParticipantConfiguration creates a new AEAssessmentParticipantConfiguration instance.
func NewAEAssessmentParticipantConfiguration() AEAssessmentParticipantConfiguration {
	return getAEAssessmentParticipantConfigurationClass().New()
}




// A Boolean that indicates whether an app can access network resources during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/allowsNetworkAccess

func (a_ AEAssessmentParticipantConfiguration) AllowsNetworkAccess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsNetworkAccess"))
	return rv
}


// A Boolean that indicates whether an app can access network resources during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/allowsNetworkAccess

func (a_ AEAssessmentParticipantConfiguration) SetAllowsNetworkAccess(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/configurationInfo

func (a_ AEAssessmentParticipantConfiguration) ConfigurationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("configurationInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/configurationInfo

func (a_ AEAssessmentParticipantConfiguration) SetConfigurationInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/isRequired

func (a_ AEAssessmentParticipantConfiguration) Required() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("required"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/isRequired

func (a_ AEAssessmentParticipantConfiguration) SetRequired(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequired:"), value)
}


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication

func (a_ AEAssessmentParticipantConfiguration) ConfigurationsByApplication() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication

func (a_ AEAssessmentParticipantConfiguration) SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationsByApplication:"), value)
}


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration

func (a_ AEAssessmentParticipantConfiguration) MainParticipantConfiguration() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration

func (a_ AEAssessmentParticipantConfiguration) SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMainParticipantConfiguration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentparticipantconfiguration/isrequired

func (a_ AEAssessmentParticipantConfiguration) IsRequired() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRequired"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentparticipantconfiguration/isrequired

func (a_ AEAssessmentParticipantConfiguration) SetIsRequired(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRequired:"), value)
}


