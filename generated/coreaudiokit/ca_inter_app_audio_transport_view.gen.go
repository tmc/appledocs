// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

/* debug [class.gen.go]: Generating class CAInterAppAudioTransportView */


/* debug [class_header]: Header for CAInterAppAudioTransportView */
// The class instance for the [InterAppAudioTransportView] class.
var (
	InterAppAudioTransportViewClass     _InterAppAudioTransportViewClass
	InterAppAudioTransportViewClassOnce sync.Once
)

func getInterAppAudioTransportViewClass() _InterAppAudioTransportViewClass {
	InterAppAudioTransportViewClassOnce.Do(func() {
		InterAppAudioTransportViewClass = _InterAppAudioTransportViewClass{objc.GetClass("CAInterAppAudioTransportView")}
	})
	return InterAppAudioTransportViewClass
}

type _InterAppAudioTransportViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InterAppAudioTransportView */
// An interface definition for the [InterAppAudioTransportView] class.
type IInterAppAudioTransportView interface {
	IView
	
/* debug [class_interface_properties]: Properties for InterAppAudioTransportView */
	// properties:
	IsConnected() bool
	SetIsConnected(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsPlaying() bool
	SetIsPlaying(value bool)
	IsRecording() bool
	SetIsRecording(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InterAppAudioTransportView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InterAppAudioTransportView */
// Alloc allocates a new instance without initialization.
func (ic _InterAppAudioTransportViewClass) Alloc() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InterAppAudioTransportViewClass) New() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterAppAudioTransportView) Init() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterAppAudioTransportView) Autorelease() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterAppAudioTransportView creates a new InterAppAudioTransportView instance.
func NewInterAppAudioTransportView() InterAppAudioTransportView {
	return getInterAppAudioTransportViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InterAppAudioTransportView */
// A view that provides an audio transport user interface.


// A view that provides an audio transport user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView
type InterAppAudioTransportView struct {
	View
}

// InterAppAudioTransportViewFrom constructs a [InterAppAudioTransportView] from an unsafe.Pointer.
//
// A view that provides an audio transport user interface.
func InterAppAudioTransportViewFrom(ptr unsafe.Pointer) InterAppAudioTransportView {
	return InterAppAudioTransportView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InterAppAudioTransportView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InterAppAudioTransportView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InterAppAudioTransportView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InterAppAudioTransportView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InterAppAudioTransportView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isconnected
func (i_ InterAppAudioTransportView) IsConnected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isConnected"))
	return rv
}/* debug [instance_properties/getter]: isConnected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isconnected
func (i_ InterAppAudioTransportView) SetIsConnected(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsConnected:"), value)
}/* debug [instance_properties/setter]: isConnected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isenabled
func (i_ InterAppAudioTransportView) IsEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isenabled
func (i_ InterAppAudioTransportView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isplaying
func (i_ InterAppAudioTransportView) IsPlaying() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isPlaying"))
	return rv
}/* debug [instance_properties/getter]: isPlaying */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isplaying
func (i_ InterAppAudioTransportView) SetIsPlaying(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsPlaying:"), value)
}/* debug [instance_properties/setter]: isPlaying */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isrecording
func (i_ InterAppAudioTransportView) IsRecording() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isRecording"))
	return rv
}/* debug [instance_properties/getter]: isRecording */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isrecording
func (i_ InterAppAudioTransportView) SetIsRecording(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRecording:"), value)
}/* debug [instance_properties/setter]: isRecording */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAInterAppAudioTransportView */


