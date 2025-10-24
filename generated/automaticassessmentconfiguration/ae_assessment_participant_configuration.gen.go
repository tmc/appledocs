// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AEAssessmentParticipantConfiguration */


/* debug [class_header]: Header for AEAssessmentParticipantConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AEAssessmentParticipantConfiguration */
// An interface definition for the [AEAssessmentParticipantConfiguration] class.
type IAEAssessmentParticipantConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AEAssessmentParticipantConfiguration */
	// properties:
	AllowsNetworkAccess() bool
	SetAllowsNetworkAccess(value bool)
	ConfigurationInfo() foundation.IDictionary
	SetConfigurationInfo(value foundation.IDictionary)
	Required() bool
	SetRequired(value bool)
	ConfigurationsByApplication() IAEAssessmentParticipantConfiguration
	SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration)
	MainParticipantConfiguration() IAEAssessmentParticipantConfiguration
	SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration)
	IsRequired() bool
	SetIsRequired(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AEAssessmentParticipantConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AEAssessmentParticipantConfiguration */
// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentParticipantConfigurationClass) Alloc() AEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AEAssessmentParticipantConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AEAssessmentParticipantConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AEAssessmentParticipantConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AEAssessmentParticipantConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AEAssessmentParticipantConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AEAssessmentParticipantConfiguration */

// A Boolean that indicates whether an app can access network resources during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/allowsNetworkAccess
func (a_ AEAssessmentParticipantConfiguration) AllowsNetworkAccess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsNetworkAccess */


// A Boolean that indicates whether an app can access network resources during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/allowsNetworkAccess
func (a_ AEAssessmentParticipantConfiguration) SetAllowsNetworkAccess(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsNetworkAccess */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/configurationInfo
func (a_ AEAssessmentParticipantConfiguration) ConfigurationInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("configurationInfo"))
	return rv
}/* debug [instance_properties/getter]: configurationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/configurationInfo
func (a_ AEAssessmentParticipantConfiguration) SetConfigurationInfo(value foundation.IDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationInfo:"), value)
}/* debug [instance_properties/setter]: configurationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/isRequired
func (a_ AEAssessmentParticipantConfiguration) Required() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("required"))
	return rv
}/* debug [instance_properties/getter]: required */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentParticipantConfiguration/isRequired
func (a_ AEAssessmentParticipantConfiguration) SetRequired(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequired:"), value)
}/* debug [instance_properties/setter]: required */


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication
func (a_ AEAssessmentParticipantConfiguration) ConfigurationsByApplication() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}/* debug [instance_properties/getter]: configurationsByApplication */


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication
func (a_ AEAssessmentParticipantConfiguration) SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationsByApplication:"), value)
}/* debug [instance_properties/setter]: configurationsByApplication */


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration
func (a_ AEAssessmentParticipantConfiguration) MainParticipantConfiguration() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}/* debug [instance_properties/getter]: mainParticipantConfiguration */


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration
func (a_ AEAssessmentParticipantConfiguration) SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMainParticipantConfiguration:"), value)
}/* debug [instance_properties/setter]: mainParticipantConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentparticipantconfiguration/isrequired
func (a_ AEAssessmentParticipantConfiguration) IsRequired() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRequired"))
	return rv
}/* debug [instance_properties/getter]: isRequired */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentparticipantconfiguration/isrequired
func (a_ AEAssessmentParticipantConfiguration) SetIsRequired(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRequired:"), value)
}/* debug [instance_properties/setter]: isRequired */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AEAssessmentParticipantConfiguration */


