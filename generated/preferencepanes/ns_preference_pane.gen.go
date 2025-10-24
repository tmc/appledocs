// Code generated from Apple documentation for PreferencePanes. DO NOT EDIT.

package preferencepanes

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPreferencePane */


/* debug [class_header]: Header for NSPreferencePane */
// The class instance for the [PreferencePane] class.
var (
	PreferencePaneClass     _PreferencePaneClass
	PreferencePaneClassOnce sync.Once
)

func getPreferencePaneClass() _PreferencePaneClass {
	PreferencePaneClassOnce.Do(func() {
		PreferencePaneClass = _PreferencePaneClass{objc.GetClass("NSPreferencePane")}
	})
	return PreferencePaneClass
}

type _PreferencePaneClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreferencePane */
// An interface definition for the [PreferencePane] class.
type IPreferencePane interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreferencePane */
	// properties:
	AutoSaveTextFields() bool
	Bundle() foundation.Bundle
	FirstKeyView() appkit.View
	SetFirstKeyView(value appkit.View)
	InitialKeyView() appkit.View
	SetInitialKeyView(value appkit.View)
	Selected() bool
	LastKeyView() appkit.View
	SetLastKeyView(value appkit.View)
	MainNibName() objc.IObject /* cross-framework: NSString */
	MainView() appkit.View
	SetMainView(value appkit.View)
	ShouldUnselect() PreferencePaneUnselectReply
	IsSelected() bool
	SetIsSelected(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreferencePane */
	// methods:
	AssignMainView()
	DidSelect()
	DidUnselect()
	LoadMainView() appkit.View
	MainViewDidLoad()
	ReplyToShouldUnselect(shouldUnselect bool)
	UpdateHelpMenuWithArray(inArrayOfMenuItems foundation.IDictionary)
	WillSelect()
	WillUnselect()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreferencePane */
// Alloc allocates a new instance without initialization.
func (pc _PreferencePaneClass) Alloc() PreferencePane {
	rv := objc.Send[PreferencePane](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreferencePaneClass) New() PreferencePane {
	rv := objc.Send[PreferencePane](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreferencePane) Init() PreferencePane {
	rv := objc.Send[PreferencePane](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreferencePane) Autorelease() PreferencePane {
	rv := objc.Send[PreferencePane](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreferencePane creates a new PreferencePane instance.
func NewPreferencePane() PreferencePane {
	return getPreferencePaneClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreferencePane */
// The interface for providing preference panes to System Preferences or other apps.
//
// Preference panes are subclasses of , packaged up in bundles and loaded by a preference application, such as System Preferences. These bundles have a suffix of . Bundles intended for use by System Preferences are located in the directories of the various file system domains. See the chapter in for information about domains. The preference pane bundle normally contains a nib file with the user interface for modifying user preferences. The nib file contains a window assigned to the _window outlet of the preference pane instance (the nib’s File’s Owner). The implementation of , invoked by the preference application, loads the nib file and uses the content view of _window as the preference pane’s main view. Override this method if you need a different technique for creating the user interface. The subclass is responsible for initializing the user interface with the current preference settings and recording any modifications the user makes. Through a series of , , and methods, the preference application notifies the preference pane when the pane is selected (displayed) and deselected, allowing the pane to perform the necessary actions at the appropriate times. Implement these methods (and any additional target-action methods connected to the interface) as needed to produce the desired behavior for your preference pane. Preference panes support Help menu items. Specify global help menu items in your bundle’s file under the key. To add dynamic help items, implement the method.


// The interface for providing preference panes to System Preferences or other apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane
type PreferencePane struct {
	objectivec.Object
}

// PreferencePaneFrom constructs a [PreferencePane] from an unsafe.Pointer.
//
// The interface for providing preference panes to System Preferences or other apps.
func PreferencePaneFrom(ptr unsafe.Pointer) PreferencePane {
	return PreferencePane{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreferencePane */

// Initializes a preference pane with the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/init(bundle:)
func NewPreferencePaneWithBundle(bundle foundation.Bundle) PreferencePane {
	instance := getPreferencePaneClass().Alloc()
	rv := objc.Send[PreferencePane](instance.ID, objc.Sel("initWithBundle:"), bundle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreferencePaneWithBundle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreferencePane */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreferencePane */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreferencePane */

// Locates and assigns the preference pane’s main view from the nib file loaded by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/assignMainView()
func (p_ PreferencePane) AssignMainView() {
	objc.Send[objc.ID](p_.ID, objc.Sel("assignMainView"))
}/* debug [instance_methods/method]: AssignMainView */


// Notifies the preference pane that the main app has just displayed the preference pane’s main view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/didSelect()
func (p_ PreferencePane) DidSelect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("didSelect"))
}/* debug [instance_methods/method]: DidSelect */


// Notifies the preference pane that the main app has just stopped displaying the preference pane’s main view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/didUnselect()
func (p_ PreferencePane) DidUnselect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("didUnselect"))
}/* debug [instance_methods/method]: DidUnselect */


// Loads the preference pane’s user interface into its main view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/loadMainView()
func (p_ PreferencePane) LoadMainView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("loadMainView"))
	return rv
}/* debug [instance_methods/method]: LoadMainView */


// Notifies the preference pane that the main view is set up and prepared to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainViewDidLoad()
func (p_ PreferencePane) MainViewDidLoad() {
	objc.Send[objc.ID](p_.ID, objc.Sel("mainViewDidLoad"))
}/* debug [instance_methods/method]: MainViewDidLoad */


// Notifies the main application of the preference pane’s ability to be deselected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/reply(toShouldUnselect:)
func (p_ PreferencePane) ReplyToShouldUnselect(shouldUnselect bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("replyToShouldUnselect:"), shouldUnselect)
}/* debug [instance_methods/method]: ReplyToShouldUnselect */


// Updates the help menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/updateHelpMenu(with:)
func (p_ PreferencePane) UpdateHelpMenuWithArray(inArrayOfMenuItems foundation.IDictionary) {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateHelpMenuWithArray:"), inArrayOfMenuItems)
}/* debug [instance_methods/method]: UpdateHelpMenuWithArray */


// Notifies the preference pane that the main app is about to display the preference pane’s main view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/willSelect()
func (p_ PreferencePane) WillSelect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("willSelect"))
}/* debug [instance_methods/method]: WillSelect */


// Notifies the preference pane that the main app is about to stop displaying the preference pane’s main view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/willUnselect()
func (p_ PreferencePane) WillUnselect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("willUnselect"))
}/* debug [instance_methods/method]: WillUnselect */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreferencePane */

// A Boolean value that indicates whether text fields save their values before changing preference panes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/autoSaveTextFields
func (p_ PreferencePane) AutoSaveTextFields() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoSaveTextFields"))
	return rv
}/* debug [instance_properties/getter]: autoSaveTextFields */


// The preference pane’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/bundle
func (p_ PreferencePane) Bundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](p_.ID, objc.Sel("bundle"))
	return rv
}/* debug [instance_properties/getter]: bundle */


// The first view in the keyboard focus chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/firstKeyView
func (p_ PreferencePane) FirstKeyView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("firstKeyView"))
	return rv
}/* debug [instance_properties/getter]: firstKeyView */


// The first view in the keyboard focus chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/firstKeyView
func (p_ PreferencePane) SetFirstKeyView(value appkit.View) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFirstKeyView:"), value)
}/* debug [instance_properties/setter]: firstKeyView */


// The view that should have keyboard focus when the pane is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/initialKeyView
func (p_ PreferencePane) InitialKeyView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("initialKeyView"))
	return rv
}/* debug [instance_properties/getter]: initialKeyView */


// The view that should have keyboard focus when the pane is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/initialKeyView
func (p_ PreferencePane) SetInitialKeyView(value appkit.View) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInitialKeyView:"), value)
}/* debug [instance_properties/setter]: initialKeyView */


// A Boolean value that indicates whether the preference pane is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/isSelected
func (p_ PreferencePane) Selected() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selected"))
	return rv
}/* debug [instance_properties/getter]: selected */


// The last view in the keyboard focus chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/lastKeyView
func (p_ PreferencePane) LastKeyView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("lastKeyView"))
	return rv
}/* debug [instance_properties/getter]: lastKeyView */


// The last view in the keyboard focus chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/lastKeyView
func (p_ PreferencePane) SetLastKeyView(value appkit.View) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastKeyView:"), value)
}/* debug [instance_properties/setter]: lastKeyView */


// The name of the preference pane’s nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainNibName
func (p_ PreferencePane) MainNibName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("mainNibName"))
	return rv
}/* debug [instance_properties/getter]: mainNibName */


// The main view of the preference pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainView
func (p_ PreferencePane) MainView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("mainView"))
	return rv
}/* debug [instance_properties/getter]: mainView */


// The main view of the preference pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainView
func (p_ PreferencePane) SetMainView(value appkit.View) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMainView:"), value)
}/* debug [instance_properties/setter]: mainView */


// A Boolean value that indicates whether the preference pane is able to be deselected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/shouldUnselect
func (p_ PreferencePane) ShouldUnselect() PreferencePaneUnselectReply {
	rv := objc.Send[PreferencePaneUnselectReply](p_.ID, objc.Sel("shouldUnselect"))
	return rv
}/* debug [instance_properties/getter]: shouldUnselect */


// A Boolean value that indicates whether the preference pane is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/preferencepanes/nspreferencepane/isselected
func (p_ PreferencePane) IsSelected() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSelected"))
	return rv
}/* debug [instance_properties/getter]: isSelected */


// A Boolean value that indicates whether the preference pane is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/preferencepanes/nspreferencepane/isselected
func (p_ PreferencePane) SetIsSelected(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSelected:"), value)
}/* debug [instance_properties/setter]: isSelected */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPreferencePane */


