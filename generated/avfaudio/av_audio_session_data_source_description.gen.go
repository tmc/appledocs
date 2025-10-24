// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSessionDataSourceDescription */


/* debug [class_header]: Header for AVAudioSessionDataSourceDescription */
// The class instance for the [AudioSessionDataSourceDescription] class.
var (
	AudioSessionDataSourceDescriptionClass     _AudioSessionDataSourceDescriptionClass
	AudioSessionDataSourceDescriptionClassOnce sync.Once
)

func getAudioSessionDataSourceDescriptionClass() _AudioSessionDataSourceDescriptionClass {
	AudioSessionDataSourceDescriptionClassOnce.Do(func() {
		AudioSessionDataSourceDescriptionClass = _AudioSessionDataSourceDescriptionClass{objc.GetClass("AVAudioSessionDataSourceDescription")}
	})
	return AudioSessionDataSourceDescriptionClass
}

type _AudioSessionDataSourceDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSessionDataSourceDescription */
// An interface definition for the [AudioSessionDataSourceDescription] class.
type IAudioSessionDataSourceDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSessionDataSourceDescription */
	// properties:
	OutputDataSource() IAVAudioSessionDataSourceDescription
	SetOutputDataSource(value IAVAudioSessionDataSourceDescription)
	OutputDataSources() IAVAudioSessionDataSourceDescription
	SetOutputDataSources(value IAVAudioSessionDataSourceDescription)
	DataSources() IAVAudioSessionDataSourceDescription
	SetDataSources(value IAVAudioSessionDataSourceDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSessionDataSourceDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSessionDataSourceDescription */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionDataSourceDescriptionClass) Alloc() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionDataSourceDescriptionClass) New() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionDataSourceDescription) Init() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionDataSourceDescription) Autorelease() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionDataSourceDescription creates a new AudioSessionDataSourceDescription instance.
func NewAudioSessionDataSourceDescription() AudioSessionDataSourceDescription {
	return getAudioSessionDataSourceDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSessionDataSourceDescription */
// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
//
// You obtain data source descriptions from the shared object or the objects corresponding to its input and output ports. Only built-in microphone ports on certain devices support the location, orientation, and polar pattern properties. If a port doesn’t support these features, the value of its property is . This class is especially useful for differentiating between microphone configurations on devices having more than one built-in microphone. Such devices may also support signal processing features for spatial filtering, or , in which the system makes the device more sensitive to audio signals from a particular direction. See for more information.


// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription
type AudioSessionDataSourceDescription struct {
	objectivec.Object
}

// AudioSessionDataSourceDescriptionFrom constructs a [AudioSessionDataSourceDescription] from an unsafe.Pointer.
//
// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
func AudioSessionDataSourceDescriptionFrom(ptr unsafe.Pointer) AudioSessionDataSourceDescription {
	return AudioSessionDataSourceDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSessionDataSourceDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSessionDataSourceDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSessionDataSourceDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSessionDataSourceDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSessionDataSourceDescription */

// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasource
func (a_ AudioSessionDataSourceDescription) OutputDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSource"))
	return rv
}/* debug [instance_properties/getter]: outputDataSource */


// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasource
func (a_ AudioSessionDataSourceDescription) SetOutputDataSource(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputDataSource:"), value)
}/* debug [instance_properties/setter]: outputDataSource */


// An array of available output data sources for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasources
func (a_ AudioSessionDataSourceDescription) OutputDataSources() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSources"))
	return rv
}/* debug [instance_properties/getter]: outputDataSources */


// An array of available output data sources for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasources
func (a_ AudioSessionDataSourceDescription) SetOutputDataSources(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputDataSources:"), value)
}/* debug [instance_properties/setter]: outputDataSources */


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) DataSources() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("dataSources"))
	return rv
}/* debug [instance_properties/getter]: dataSources */


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) SetDataSources(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSources:"), value)
}/* debug [instance_properties/setter]: dataSources */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSessionDataSourceDescription */


