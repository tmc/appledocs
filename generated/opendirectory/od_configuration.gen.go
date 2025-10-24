// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/securityfoundation"
)

/* debug [class.gen.go]: Generating class ODConfiguration */


/* debug [class_header]: Header for ODConfiguration */
// The class instance for the [ODConfiguration] class.
var (
	ODConfigurationClass     _ODConfigurationClass
	ODConfigurationClassOnce sync.Once
)

func getODConfigurationClass() _ODConfigurationClass {
	ODConfigurationClassOnce.Do(func() {
		ODConfigurationClass = _ODConfigurationClass{objc.GetClass("ODConfiguration")}
	})
	return ODConfigurationClass
}

type _ODConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODConfiguration */
// An interface definition for the [ODConfiguration] class.
type IODConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODConfiguration */
	// properties:
	AuthenticationModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetAuthenticationModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	Comment() objc.IObject /* cross-framework: NSString */
	SetComment(value objc.IObject /* cross-framework: NSString */)
	ConnectionIdleTimeoutInSeconds() int
	SetConnectionIdleTimeoutInSeconds(value int)
	ConnectionSetupTimeoutInSeconds() int
	SetConnectionSetupTimeoutInSeconds(value int)
	DefaultMappings() IODMappings
	SetDefaultMappings(value IODMappings)
	DefaultModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetDefaultModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	DiscoveryModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetDiscoveryModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	GeneralModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetGeneralModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	HideRegistration() bool
	SetHideRegistration(value bool)
	ManInTheMiddleProtection() bool
	SetManInTheMiddleProtection(value bool)
	NodeName() objc.IObject /* cross-framework: NSString */
	SetNodeName(value objc.IObject /* cross-framework: NSString */)
	PacketEncryption() int
	SetPacketEncryption(value int)
	PacketSigning() int
	SetPacketSigning(value int)
	PreferredDestinationHostName() objc.IObject /* cross-framework: NSString */
	SetPreferredDestinationHostName(value objc.IObject /* cross-framework: NSString */)
	PreferredDestinationHostPort() uint16 /* not a class type */
	SetPreferredDestinationHostPort(value uint16 /* not a class type */)
	QueryTimeoutInSeconds() int
	SetQueryTimeoutInSeconds(value int)
	TemplateName() objc.IObject /* cross-framework: NSString */
	SetTemplateName(value objc.IObject /* cross-framework: NSString */)
	TrustAccount() objc.IObject /* cross-framework: NSString */
	TrustKerberosPrincipal() objc.IObject /* cross-framework: NSString */
	TrustMetaAccount() objc.IObject /* cross-framework: NSString */
	TrustType() objc.IObject /* cross-framework: NSString */
	TrustUsesKerberosKeytab() bool
	TrustUsesMutualAuthentication() bool
	TrustUsesSystemKeychain() bool
	VirtualSubnodes() objc.IObject /* cross-framework: NSArray */
	SetVirtualSubnodes(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODConfiguration */
	// methods:
	AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType objc.IObject /* cross-framework: NSString */, account objc.IObject /* cross-framework: NSString */, accountPassword objc.IObject /* cross-framework: NSString */, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, join bool, error_ unsafe.Pointer) bool
	RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, deleteAccount bool, error_ unsafe.Pointer) bool
	SaveUsingAuthorizationError(authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODConfiguration */
// Alloc allocates a new instance without initialization.
func (oc _ODConfigurationClass) Alloc() ODConfiguration {
	rv := objc.Send[ODConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ODConfigurationClass) New() ODConfiguration {
	rv := objc.Send[ODConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODConfiguration) Init() ODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODConfiguration) Autorelease() ODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODConfiguration creates a new ODConfiguration instance.
func NewODConfiguration() ODConfiguration {
	return getODConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration
type ODConfiguration struct {
	objectivec.Object
}

// ODConfigurationFrom constructs a [ODConfiguration] from an unsafe.Pointer.
func ODConfigurationFrom(ptr unsafe.Pointer) ODConfiguration {
	return ODConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/configuration
func (oc _ODConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("configuration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Configuration) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustAccount(_:)
func (oc _ODConfigurationClass) SuggestedTrustAccount(hostname objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(oc.class), objc.Sel("suggestedTrustAccount:"), hostname)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SuggestedTrustAccount) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustPassword(_:)
func (oc _ODConfigurationClass) SuggestedTrustPassword(length uintptr /* not a class type */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(oc.class), objc.Sel("suggestedTrustPassword:"), length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SuggestedTrustPassword) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/addTrustType(_:trustAccount:trustPassword:username:password:joinExisting:)
func (o_ ODConfiguration) AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType objc.IObject /* cross-framework: NSString */, account objc.IObject /* cross-framework: NSString */, accountPassword objc.IObject /* cross-framework: NSString */, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, join bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addTrustType:trustAccount:trustPassword:username:password:joinExisting:error:"), trustType, account, accountPassword, username, password, join, error_)
	return rv
}/* debug [instance_methods/method]: AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/removeTrust(usingUsername:password:deleteTrustAccount:)
func (o_ ODConfiguration) RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, deleteAccount bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeTrustUsingUsername:password:deleteTrustAccount:error:"), username, password, deleteAccount, error_)
	return rv
}/* debug [instance_methods/method]: RemoveTrustUsingUsernamePasswordDeleteTrustAccountError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/save(using:)
func (o_ ODConfiguration) SaveUsingAuthorizationError(authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveUsingAuthorization:error:"), authorization, error_)
	return rv
}/* debug [instance_methods/method]: SaveUsingAuthorizationError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) AuthenticationModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("authenticationModuleEntries"))
	return rv
}/* debug [instance_properties/getter]: authenticationModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) SetAuthenticationModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAuthenticationModuleEntries:"), value)
}/* debug [instance_properties/setter]: authenticationModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) Comment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("comment"))
	return rv
}/* debug [instance_properties/getter]: comment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) SetComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), value)
}/* debug [instance_properties/setter]: comment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionIdleTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionIdleTimeoutInSeconds"))
	return rv
}/* debug [instance_properties/getter]: connectionIdleTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionIdleTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionIdleTimeoutInSeconds:"), value)
}/* debug [instance_properties/setter]: connectionIdleTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionSetupTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionSetupTimeoutInSeconds"))
	return rv
}/* debug [instance_properties/getter]: connectionSetupTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionSetupTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionSetupTimeoutInSeconds:"), value)
}/* debug [instance_properties/setter]: connectionSetupTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) DefaultMappings() IODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("defaultMappings"))
	return rv
}/* debug [instance_properties/getter]: defaultMappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) SetDefaultMappings(value IODMappings) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultMappings:"), value)
}/* debug [instance_properties/setter]: defaultMappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) DefaultModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("defaultModuleEntries"))
	return rv
}/* debug [instance_properties/getter]: defaultModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) SetDefaultModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultModuleEntries:"), value)
}/* debug [instance_properties/setter]: defaultModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) DiscoveryModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("discoveryModuleEntries"))
	return rv
}/* debug [instance_properties/getter]: discoveryModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) SetDiscoveryModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDiscoveryModuleEntries:"), value)
}/* debug [instance_properties/setter]: discoveryModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) GeneralModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("generalModuleEntries"))
	return rv
}/* debug [instance_properties/getter]: generalModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) SetGeneralModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGeneralModuleEntries:"), value)
}/* debug [instance_properties/setter]: generalModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) HideRegistration() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hideRegistration"))
	return rv
}/* debug [instance_properties/getter]: hideRegistration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) SetHideRegistration(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setHideRegistration:"), value)
}/* debug [instance_properties/setter]: hideRegistration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) ManInTheMiddleProtection() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("manInTheMiddleProtection"))
	return rv
}/* debug [instance_properties/getter]: manInTheMiddleProtection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) SetManInTheMiddleProtection(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManInTheMiddleProtection:"), value)
}/* debug [instance_properties/setter]: manInTheMiddleProtection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) NodeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("nodeName"))
	return rv
}/* debug [instance_properties/getter]: nodeName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) SetNodeName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNodeName:"), value)
}/* debug [instance_properties/setter]: nodeName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) PacketEncryption() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetEncryption"))
	return rv
}/* debug [instance_properties/getter]: packetEncryption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) SetPacketEncryption(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketEncryption:"), value)
}/* debug [instance_properties/setter]: packetEncryption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) PacketSigning() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetSigning"))
	return rv
}/* debug [instance_properties/getter]: packetSigning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) SetPacketSigning(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketSigning:"), value)
}/* debug [instance_properties/setter]: packetSigning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) PreferredDestinationHostName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("preferredDestinationHostName"))
	return rv
}/* debug [instance_properties/getter]: preferredDestinationHostName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostName:"), value)
}/* debug [instance_properties/setter]: preferredDestinationHostName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) PreferredDestinationHostPort() uint16 /* not a class type */ {
	rv := objc.Send[uint16](o_.ID, objc.Sel("preferredDestinationHostPort"))
	return rv
}/* debug [instance_properties/getter]: preferredDestinationHostPort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostPort(value uint16 /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostPort:"), value)
}/* debug [instance_properties/setter]: preferredDestinationHostPort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) QueryTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("queryTimeoutInSeconds"))
	return rv
}/* debug [instance_properties/getter]: queryTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetQueryTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueryTimeoutInSeconds:"), value)
}/* debug [instance_properties/setter]: queryTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) TemplateName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("templateName"))
	return rv
}/* debug [instance_properties/getter]: templateName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) SetTemplateName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), value)
}/* debug [instance_properties/setter]: templateName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-swift.property
func (o_ ODConfiguration) TrustAccount() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustAccount"))
	return rv
}/* debug [instance_properties/getter]: trustAccount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-swift.property
func (o_ ODConfiguration) TrustKerberosPrincipal() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustKerberosPrincipal"))
	return rv
}/* debug [instance_properties/getter]: trustKerberosPrincipal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-swift.property
func (o_ ODConfiguration) TrustMetaAccount() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustMetaAccount"))
	return rv
}/* debug [instance_properties/getter]: trustMetaAccount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-swift.property
func (o_ ODConfiguration) TrustType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustType"))
	return rv
}/* debug [instance_properties/getter]: trustType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-swift.property
func (o_ ODConfiguration) TrustUsesKerberosKeytab() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesKerberosKeytab"))
	return rv
}/* debug [instance_properties/getter]: trustUsesKerberosKeytab */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-swift.property
func (o_ ODConfiguration) TrustUsesMutualAuthentication() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesMutualAuthentication"))
	return rv
}/* debug [instance_properties/getter]: trustUsesMutualAuthentication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-swift.property
func (o_ ODConfiguration) TrustUsesSystemKeychain() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesSystemKeychain"))
	return rv
}/* debug [instance_properties/getter]: trustUsesSystemKeychain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) VirtualSubnodes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("virtualSubnodes"))
	return rv
}/* debug [instance_properties/getter]: virtualSubnodes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) SetVirtualSubnodes(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVirtualSubnodes:"), value)
}/* debug [instance_properties/setter]: virtualSubnodes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODConfiguration */



