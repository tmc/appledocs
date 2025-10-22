// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// SystemConfiguration Functions (178 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CNCopyCurrentNetworkInfo func(unsafe.Pointer) unsafe.Pointer
	_CNCopySupportedInterfaces func() unsafe.Pointer
	_CNMarkPortalOffline func(unsafe.Pointer) unsafe.Pointer
	_CNMarkPortalOnline func(unsafe.Pointer) unsafe.Pointer
	_CNSetSupportedSSIDs func(unsafe.Pointer) unsafe.Pointer
	_DHCPClientPreferencesCopyApplicationOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DHCPClientPreferencesSetApplicationOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DHCPInfoGetLeaseExpirationTime func(unsafe.Pointer) unsafe.Pointer
	_DHCPInfoGetLeaseStartTime func(unsafe.Pointer) unsafe.Pointer
	_DHCPInfoGetOptionData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceCopyAll func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceCopyAvailableMemberInterfaces func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceCopyStatus func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceCreate func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceGetMemberInterfaces func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceGetOptions func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceRemove func(unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceSetLocalizedDisplayName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceSetMemberInterfaces func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCBondInterfaceSetOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCBondStatusGetInterfaceStatus func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCBondStatusGetMemberInterfaces func(unsafe.Pointer) unsafe.Pointer
	_SCBondStatusGetTypeID func() unsafe.Pointer
	_SCCopyLastError func() unsafe.Pointer
	_SCDynamicStoreAddTemporaryValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreAddValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyComputerName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyConsoleUser func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyDHCPInfo func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyKeyList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyLocalHostName func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyLocation func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyMultiple func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyNotifiedKeys func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyProxies func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCopyValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreCreateWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreGetTypeID func() unsafe.Pointer
	_SCDynamicStoreKeyCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateComputerName func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateConsoleUser func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateHostNames func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateLocation func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateNetworkGlobalEntity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateNetworkInterface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateNetworkInterfaceEntity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateNetworkServiceEntity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreKeyCreateProxies func(unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreNotifyValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreSetMultiple func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreSetNotificationKeys func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCError func() int
	_SCErrorString func(int) unsafe.Pointer
	_SCNetworkCheckReachabilityByAddress func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkCheckReachabilityByName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyExtendedStatus func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyServiceID func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyStatistics func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyUserOptions func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyUserPreferences func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCreateWithServiceID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionGetStatus func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionGetTypeID func() unsafe.Pointer
	_SCNetworkConnectionScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionStart func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionStop func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCopyAll func() unsafe.Pointer
	_SCNetworkInterfaceCopyMTU func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCopyMediaOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCopyMediaSubTypeOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCopyMediaSubTypes func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCreateWithInterface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceForceConfigurationRefresh func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetBSDName func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetConfiguration func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetExtendedConfiguration func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetHardwareAddressString func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetInterface func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetInterfaceType func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetLocalizedDisplayName func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetSupportedInterfaceTypes func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetSupportedProtocolTypes func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceGetTypeID func() unsafe.Pointer
	_SCNetworkInterfaceRefreshConfiguration func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceSetConfiguration func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceSetExtendedConfiguration func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceSetMTU func(unsafe.Pointer, int) unsafe.Pointer
	_SCNetworkInterfaceSetMediaOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkProtocolGetConfiguration func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkProtocolGetEnabled func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkProtocolGetProtocolType func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkProtocolGetTypeID func() unsafe.Pointer
	_SCNetworkProtocolSetConfiguration func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkProtocolSetEnabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityCreateWithAddress func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityCreateWithAddressPair func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityCreateWithName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityGetFlags func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityGetTypeID func() unsafe.Pointer
	_SCNetworkReachabilityScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilitySetCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilitySetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceAddProtocolType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceCopyAll func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceCopyProtocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceCopyProtocols func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceEstablishDefaultConfiguration func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceGetEnabled func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceGetInterface func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceGetName func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceGetServiceID func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceGetTypeID func() unsafe.Pointer
	_SCNetworkServiceRemove func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceRemoveProtocolType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceSetEnabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceSetName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetAddService func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetContainsInterface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetCopyAll func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetCopyCurrent func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetCopyServices func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetCreate func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetGetName func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetGetServiceOrder func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetGetSetID func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetGetTypeID func() unsafe.Pointer
	_SCNetworkSetRemove func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetRemoveService func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetSetCurrent func(unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetSetName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkSetSetServiceOrder func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesAddValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesApplyChanges func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesCommitChanges func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesCopyKeyList func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesCreateWithAuthorization func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesGetSignature func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesGetTypeID func() unsafe.Pointer
	_SCPreferencesGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesLock func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathCreateUniqueChild func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathGetLink func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathSetLink func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetComputerName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetLocalHostName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSynchronize func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesUnlock func(unsafe.Pointer) unsafe.Pointer
	_SCPreferencesUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceCopyAll func(unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceCopyAvailablePhysicalInterfaces func() unsafe.Pointer
	_SCVLANInterfaceCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceGetOptions func(unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceGetPhysicalInterface func(unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceGetTag func(unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceRemove func(unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceSetLocalizedDisplayName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceSetOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCVLANInterfaceSetPhysicalInterfaceAndTag func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CNCopyCurrentNetworkInfo, lib, "CNCopyCurrentNetworkInfo")
	tryRegister(&_CNCopySupportedInterfaces, lib, "CNCopySupportedInterfaces")
	tryRegister(&_CNMarkPortalOffline, lib, "CNMarkPortalOffline")
	tryRegister(&_CNMarkPortalOnline, lib, "CNMarkPortalOnline")
	tryRegister(&_CNSetSupportedSSIDs, lib, "CNSetSupportedSSIDs")
	tryRegister(&_DHCPClientPreferencesCopyApplicationOptions, lib, "DHCPClientPreferencesCopyApplicationOptions")
	tryRegister(&_DHCPClientPreferencesSetApplicationOptions, lib, "DHCPClientPreferencesSetApplicationOptions")
	tryRegister(&_DHCPInfoGetLeaseExpirationTime, lib, "DHCPInfoGetLeaseExpirationTime")
	tryRegister(&_DHCPInfoGetLeaseStartTime, lib, "DHCPInfoGetLeaseStartTime")
	tryRegister(&_DHCPInfoGetOptionData, lib, "DHCPInfoGetOptionData")
	tryRegister(&_SCBondInterfaceCopyAll, lib, "SCBondInterfaceCopyAll")
	tryRegister(&_SCBondInterfaceCopyAvailableMemberInterfaces, lib, "SCBondInterfaceCopyAvailableMemberInterfaces")
	tryRegister(&_SCBondInterfaceCopyStatus, lib, "SCBondInterfaceCopyStatus")
	tryRegister(&_SCBondInterfaceCreate, lib, "SCBondInterfaceCreate")
	tryRegister(&_SCBondInterfaceGetMemberInterfaces, lib, "SCBondInterfaceGetMemberInterfaces")
	tryRegister(&_SCBondInterfaceGetOptions, lib, "SCBondInterfaceGetOptions")
	tryRegister(&_SCBondInterfaceRemove, lib, "SCBondInterfaceRemove")
	tryRegister(&_SCBondInterfaceSetLocalizedDisplayName, lib, "SCBondInterfaceSetLocalizedDisplayName")
	tryRegister(&_SCBondInterfaceSetMemberInterfaces, lib, "SCBondInterfaceSetMemberInterfaces")
	tryRegister(&_SCBondInterfaceSetOptions, lib, "SCBondInterfaceSetOptions")
	tryRegister(&_SCBondStatusGetInterfaceStatus, lib, "SCBondStatusGetInterfaceStatus")
	tryRegister(&_SCBondStatusGetMemberInterfaces, lib, "SCBondStatusGetMemberInterfaces")
	tryRegister(&_SCBondStatusGetTypeID, lib, "SCBondStatusGetTypeID")
	tryRegister(&_SCCopyLastError, lib, "SCCopyLastError")
	tryRegister(&_SCDynamicStoreAddTemporaryValue, lib, "SCDynamicStoreAddTemporaryValue")
	tryRegister(&_SCDynamicStoreAddValue, lib, "SCDynamicStoreAddValue")
	tryRegister(&_SCDynamicStoreCopyComputerName, lib, "SCDynamicStoreCopyComputerName")
	tryRegister(&_SCDynamicStoreCopyConsoleUser, lib, "SCDynamicStoreCopyConsoleUser")
	tryRegister(&_SCDynamicStoreCopyDHCPInfo, lib, "SCDynamicStoreCopyDHCPInfo")
	tryRegister(&_SCDynamicStoreCopyKeyList, lib, "SCDynamicStoreCopyKeyList")
	tryRegister(&_SCDynamicStoreCopyLocalHostName, lib, "SCDynamicStoreCopyLocalHostName")
	tryRegister(&_SCDynamicStoreCopyLocation, lib, "SCDynamicStoreCopyLocation")
	tryRegister(&_SCDynamicStoreCopyMultiple, lib, "SCDynamicStoreCopyMultiple")
	tryRegister(&_SCDynamicStoreCopyNotifiedKeys, lib, "SCDynamicStoreCopyNotifiedKeys")
	tryRegister(&_SCDynamicStoreCopyProxies, lib, "SCDynamicStoreCopyProxies")
	tryRegister(&_SCDynamicStoreCopyValue, lib, "SCDynamicStoreCopyValue")
	tryRegister(&_SCDynamicStoreCreate, lib, "SCDynamicStoreCreate")
	tryRegister(&_SCDynamicStoreCreateRunLoopSource, lib, "SCDynamicStoreCreateRunLoopSource")
	tryRegister(&_SCDynamicStoreCreateWithOptions, lib, "SCDynamicStoreCreateWithOptions")
	tryRegister(&_SCDynamicStoreGetTypeID, lib, "SCDynamicStoreGetTypeID")
	tryRegister(&_SCDynamicStoreKeyCreate, lib, "SCDynamicStoreKeyCreate")
	tryRegister(&_SCDynamicStoreKeyCreateComputerName, lib, "SCDynamicStoreKeyCreateComputerName")
	tryRegister(&_SCDynamicStoreKeyCreateConsoleUser, lib, "SCDynamicStoreKeyCreateConsoleUser")
	tryRegister(&_SCDynamicStoreKeyCreateHostNames, lib, "SCDynamicStoreKeyCreateHostNames")
	tryRegister(&_SCDynamicStoreKeyCreateLocation, lib, "SCDynamicStoreKeyCreateLocation")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkGlobalEntity, lib, "SCDynamicStoreKeyCreateNetworkGlobalEntity")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkInterface, lib, "SCDynamicStoreKeyCreateNetworkInterface")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkInterfaceEntity, lib, "SCDynamicStoreKeyCreateNetworkInterfaceEntity")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkServiceEntity, lib, "SCDynamicStoreKeyCreateNetworkServiceEntity")
	tryRegister(&_SCDynamicStoreKeyCreateProxies, lib, "SCDynamicStoreKeyCreateProxies")
	tryRegister(&_SCDynamicStoreNotifyValue, lib, "SCDynamicStoreNotifyValue")
	tryRegister(&_SCDynamicStoreRemoveValue, lib, "SCDynamicStoreRemoveValue")
	tryRegister(&_SCDynamicStoreSetDispatchQueue, lib, "SCDynamicStoreSetDispatchQueue")
	tryRegister(&_SCDynamicStoreSetMultiple, lib, "SCDynamicStoreSetMultiple")
	tryRegister(&_SCDynamicStoreSetNotificationKeys, lib, "SCDynamicStoreSetNotificationKeys")
	tryRegister(&_SCDynamicStoreSetValue, lib, "SCDynamicStoreSetValue")
	tryRegister(&_SCError, lib, "SCError")
	tryRegister(&_SCErrorString, lib, "SCErrorString")
	tryRegister(&_SCNetworkCheckReachabilityByAddress, lib, "SCNetworkCheckReachabilityByAddress")
	tryRegister(&_SCNetworkCheckReachabilityByName, lib, "SCNetworkCheckReachabilityByName")
	tryRegister(&_SCNetworkConnectionCopyExtendedStatus, lib, "SCNetworkConnectionCopyExtendedStatus")
	tryRegister(&_SCNetworkConnectionCopyServiceID, lib, "SCNetworkConnectionCopyServiceID")
	tryRegister(&_SCNetworkConnectionCopyStatistics, lib, "SCNetworkConnectionCopyStatistics")
	tryRegister(&_SCNetworkConnectionCopyUserOptions, lib, "SCNetworkConnectionCopyUserOptions")
	tryRegister(&_SCNetworkConnectionCopyUserPreferences, lib, "SCNetworkConnectionCopyUserPreferences")
	tryRegister(&_SCNetworkConnectionCreateWithServiceID, lib, "SCNetworkConnectionCreateWithServiceID")
	tryRegister(&_SCNetworkConnectionGetStatus, lib, "SCNetworkConnectionGetStatus")
	tryRegister(&_SCNetworkConnectionGetTypeID, lib, "SCNetworkConnectionGetTypeID")
	tryRegister(&_SCNetworkConnectionScheduleWithRunLoop, lib, "SCNetworkConnectionScheduleWithRunLoop")
	tryRegister(&_SCNetworkConnectionSetDispatchQueue, lib, "SCNetworkConnectionSetDispatchQueue")
	tryRegister(&_SCNetworkConnectionStart, lib, "SCNetworkConnectionStart")
	tryRegister(&_SCNetworkConnectionStop, lib, "SCNetworkConnectionStop")
	tryRegister(&_SCNetworkConnectionUnscheduleFromRunLoop, lib, "SCNetworkConnectionUnscheduleFromRunLoop")
	tryRegister(&_SCNetworkInterfaceCopyAll, lib, "SCNetworkInterfaceCopyAll")
	tryRegister(&_SCNetworkInterfaceCopyMTU, lib, "SCNetworkInterfaceCopyMTU")
	tryRegister(&_SCNetworkInterfaceCopyMediaOptions, lib, "SCNetworkInterfaceCopyMediaOptions")
	tryRegister(&_SCNetworkInterfaceCopyMediaSubTypeOptions, lib, "SCNetworkInterfaceCopyMediaSubTypeOptions")
	tryRegister(&_SCNetworkInterfaceCopyMediaSubTypes, lib, "SCNetworkInterfaceCopyMediaSubTypes")
	tryRegister(&_SCNetworkInterfaceCreateWithInterface, lib, "SCNetworkInterfaceCreateWithInterface")
	tryRegister(&_SCNetworkInterfaceForceConfigurationRefresh, lib, "SCNetworkInterfaceForceConfigurationRefresh")
	tryRegister(&_SCNetworkInterfaceGetBSDName, lib, "SCNetworkInterfaceGetBSDName")
	tryRegister(&_SCNetworkInterfaceGetConfiguration, lib, "SCNetworkInterfaceGetConfiguration")
	tryRegister(&_SCNetworkInterfaceGetExtendedConfiguration, lib, "SCNetworkInterfaceGetExtendedConfiguration")
	tryRegister(&_SCNetworkInterfaceGetHardwareAddressString, lib, "SCNetworkInterfaceGetHardwareAddressString")
	tryRegister(&_SCNetworkInterfaceGetInterface, lib, "SCNetworkInterfaceGetInterface")
	tryRegister(&_SCNetworkInterfaceGetInterfaceType, lib, "SCNetworkInterfaceGetInterfaceType")
	tryRegister(&_SCNetworkInterfaceGetLocalizedDisplayName, lib, "SCNetworkInterfaceGetLocalizedDisplayName")
	tryRegister(&_SCNetworkInterfaceGetSupportedInterfaceTypes, lib, "SCNetworkInterfaceGetSupportedInterfaceTypes")
	tryRegister(&_SCNetworkInterfaceGetSupportedProtocolTypes, lib, "SCNetworkInterfaceGetSupportedProtocolTypes")
	tryRegister(&_SCNetworkInterfaceGetTypeID, lib, "SCNetworkInterfaceGetTypeID")
	tryRegister(&_SCNetworkInterfaceRefreshConfiguration, lib, "SCNetworkInterfaceRefreshConfiguration")
	tryRegister(&_SCNetworkInterfaceSetConfiguration, lib, "SCNetworkInterfaceSetConfiguration")
	tryRegister(&_SCNetworkInterfaceSetExtendedConfiguration, lib, "SCNetworkInterfaceSetExtendedConfiguration")
	tryRegister(&_SCNetworkInterfaceSetMTU, lib, "SCNetworkInterfaceSetMTU")
	tryRegister(&_SCNetworkInterfaceSetMediaOptions, lib, "SCNetworkInterfaceSetMediaOptions")
	tryRegister(&_SCNetworkProtocolGetConfiguration, lib, "SCNetworkProtocolGetConfiguration")
	tryRegister(&_SCNetworkProtocolGetEnabled, lib, "SCNetworkProtocolGetEnabled")
	tryRegister(&_SCNetworkProtocolGetProtocolType, lib, "SCNetworkProtocolGetProtocolType")
	tryRegister(&_SCNetworkProtocolGetTypeID, lib, "SCNetworkProtocolGetTypeID")
	tryRegister(&_SCNetworkProtocolSetConfiguration, lib, "SCNetworkProtocolSetConfiguration")
	tryRegister(&_SCNetworkProtocolSetEnabled, lib, "SCNetworkProtocolSetEnabled")
	tryRegister(&_SCNetworkReachabilityCreateWithAddress, lib, "SCNetworkReachabilityCreateWithAddress")
	tryRegister(&_SCNetworkReachabilityCreateWithAddressPair, lib, "SCNetworkReachabilityCreateWithAddressPair")
	tryRegister(&_SCNetworkReachabilityCreateWithName, lib, "SCNetworkReachabilityCreateWithName")
	tryRegister(&_SCNetworkReachabilityGetFlags, lib, "SCNetworkReachabilityGetFlags")
	tryRegister(&_SCNetworkReachabilityGetTypeID, lib, "SCNetworkReachabilityGetTypeID")
	tryRegister(&_SCNetworkReachabilityScheduleWithRunLoop, lib, "SCNetworkReachabilityScheduleWithRunLoop")
	tryRegister(&_SCNetworkReachabilitySetCallback, lib, "SCNetworkReachabilitySetCallback")
	tryRegister(&_SCNetworkReachabilitySetDispatchQueue, lib, "SCNetworkReachabilitySetDispatchQueue")
	tryRegister(&_SCNetworkReachabilityUnscheduleFromRunLoop, lib, "SCNetworkReachabilityUnscheduleFromRunLoop")
	tryRegister(&_SCNetworkServiceAddProtocolType, lib, "SCNetworkServiceAddProtocolType")
	tryRegister(&_SCNetworkServiceCopy, lib, "SCNetworkServiceCopy")
	tryRegister(&_SCNetworkServiceCopyAll, lib, "SCNetworkServiceCopyAll")
	tryRegister(&_SCNetworkServiceCopyProtocol, lib, "SCNetworkServiceCopyProtocol")
	tryRegister(&_SCNetworkServiceCopyProtocols, lib, "SCNetworkServiceCopyProtocols")
	tryRegister(&_SCNetworkServiceCreate, lib, "SCNetworkServiceCreate")
	tryRegister(&_SCNetworkServiceEstablishDefaultConfiguration, lib, "SCNetworkServiceEstablishDefaultConfiguration")
	tryRegister(&_SCNetworkServiceGetEnabled, lib, "SCNetworkServiceGetEnabled")
	tryRegister(&_SCNetworkServiceGetInterface, lib, "SCNetworkServiceGetInterface")
	tryRegister(&_SCNetworkServiceGetName, lib, "SCNetworkServiceGetName")
	tryRegister(&_SCNetworkServiceGetServiceID, lib, "SCNetworkServiceGetServiceID")
	tryRegister(&_SCNetworkServiceGetTypeID, lib, "SCNetworkServiceGetTypeID")
	tryRegister(&_SCNetworkServiceRemove, lib, "SCNetworkServiceRemove")
	tryRegister(&_SCNetworkServiceRemoveProtocolType, lib, "SCNetworkServiceRemoveProtocolType")
	tryRegister(&_SCNetworkServiceSetEnabled, lib, "SCNetworkServiceSetEnabled")
	tryRegister(&_SCNetworkServiceSetName, lib, "SCNetworkServiceSetName")
	tryRegister(&_SCNetworkSetAddService, lib, "SCNetworkSetAddService")
	tryRegister(&_SCNetworkSetContainsInterface, lib, "SCNetworkSetContainsInterface")
	tryRegister(&_SCNetworkSetCopy, lib, "SCNetworkSetCopy")
	tryRegister(&_SCNetworkSetCopyAll, lib, "SCNetworkSetCopyAll")
	tryRegister(&_SCNetworkSetCopyCurrent, lib, "SCNetworkSetCopyCurrent")
	tryRegister(&_SCNetworkSetCopyServices, lib, "SCNetworkSetCopyServices")
	tryRegister(&_SCNetworkSetCreate, lib, "SCNetworkSetCreate")
	tryRegister(&_SCNetworkSetGetName, lib, "SCNetworkSetGetName")
	tryRegister(&_SCNetworkSetGetServiceOrder, lib, "SCNetworkSetGetServiceOrder")
	tryRegister(&_SCNetworkSetGetSetID, lib, "SCNetworkSetGetSetID")
	tryRegister(&_SCNetworkSetGetTypeID, lib, "SCNetworkSetGetTypeID")
	tryRegister(&_SCNetworkSetRemove, lib, "SCNetworkSetRemove")
	tryRegister(&_SCNetworkSetRemoveService, lib, "SCNetworkSetRemoveService")
	tryRegister(&_SCNetworkSetSetCurrent, lib, "SCNetworkSetSetCurrent")
	tryRegister(&_SCNetworkSetSetName, lib, "SCNetworkSetSetName")
	tryRegister(&_SCNetworkSetSetServiceOrder, lib, "SCNetworkSetSetServiceOrder")
	tryRegister(&_SCPreferencesAddValue, lib, "SCPreferencesAddValue")
	tryRegister(&_SCPreferencesApplyChanges, lib, "SCPreferencesApplyChanges")
	tryRegister(&_SCPreferencesCommitChanges, lib, "SCPreferencesCommitChanges")
	tryRegister(&_SCPreferencesCopyKeyList, lib, "SCPreferencesCopyKeyList")
	tryRegister(&_SCPreferencesCreate, lib, "SCPreferencesCreate")
	tryRegister(&_SCPreferencesCreateWithAuthorization, lib, "SCPreferencesCreateWithAuthorization")
	tryRegister(&_SCPreferencesGetSignature, lib, "SCPreferencesGetSignature")
	tryRegister(&_SCPreferencesGetTypeID, lib, "SCPreferencesGetTypeID")
	tryRegister(&_SCPreferencesGetValue, lib, "SCPreferencesGetValue")
	tryRegister(&_SCPreferencesLock, lib, "SCPreferencesLock")
	tryRegister(&_SCPreferencesPathCreateUniqueChild, lib, "SCPreferencesPathCreateUniqueChild")
	tryRegister(&_SCPreferencesPathGetLink, lib, "SCPreferencesPathGetLink")
	tryRegister(&_SCPreferencesPathGetValue, lib, "SCPreferencesPathGetValue")
	tryRegister(&_SCPreferencesPathRemoveValue, lib, "SCPreferencesPathRemoveValue")
	tryRegister(&_SCPreferencesPathSetLink, lib, "SCPreferencesPathSetLink")
	tryRegister(&_SCPreferencesPathSetValue, lib, "SCPreferencesPathSetValue")
	tryRegister(&_SCPreferencesRemoveValue, lib, "SCPreferencesRemoveValue")
	tryRegister(&_SCPreferencesScheduleWithRunLoop, lib, "SCPreferencesScheduleWithRunLoop")
	tryRegister(&_SCPreferencesSetCallback, lib, "SCPreferencesSetCallback")
	tryRegister(&_SCPreferencesSetComputerName, lib, "SCPreferencesSetComputerName")
	tryRegister(&_SCPreferencesSetDispatchQueue, lib, "SCPreferencesSetDispatchQueue")
	tryRegister(&_SCPreferencesSetLocalHostName, lib, "SCPreferencesSetLocalHostName")
	tryRegister(&_SCPreferencesSetValue, lib, "SCPreferencesSetValue")
	tryRegister(&_SCPreferencesSynchronize, lib, "SCPreferencesSynchronize")
	tryRegister(&_SCPreferencesUnlock, lib, "SCPreferencesUnlock")
	tryRegister(&_SCPreferencesUnscheduleFromRunLoop, lib, "SCPreferencesUnscheduleFromRunLoop")
	tryRegister(&_SCVLANInterfaceCopyAll, lib, "SCVLANInterfaceCopyAll")
	tryRegister(&_SCVLANInterfaceCopyAvailablePhysicalInterfaces, lib, "SCVLANInterfaceCopyAvailablePhysicalInterfaces")
	tryRegister(&_SCVLANInterfaceCreate, lib, "SCVLANInterfaceCreate")
	tryRegister(&_SCVLANInterfaceGetOptions, lib, "SCVLANInterfaceGetOptions")
	tryRegister(&_SCVLANInterfaceGetPhysicalInterface, lib, "SCVLANInterfaceGetPhysicalInterface")
	tryRegister(&_SCVLANInterfaceGetTag, lib, "SCVLANInterfaceGetTag")
	tryRegister(&_SCVLANInterfaceRemove, lib, "SCVLANInterfaceRemove")
	tryRegister(&_SCVLANInterfaceSetLocalizedDisplayName, lib, "SCVLANInterfaceSetLocalizedDisplayName")
	tryRegister(&_SCVLANInterfaceSetOptions, lib, "SCVLANInterfaceSetOptions")
	tryRegister(&_SCVLANInterfaceSetPhysicalInterfaceAndTag, lib, "SCVLANInterfaceSetPhysicalInterfaceAndTag")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the current network information for a given network interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/1614126-cncopycurrentnetworkinfo
func CNCopyCurrentNetworkInfo(p0 unsafe.Pointer) unsafe.Pointer {
	return _CNCopyCurrentNetworkInfo(p0)
	}


// Returns the names of all network interfaces Captive Network Support is monitoring. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNCopySupportedInterfaces
func CNCopySupportedInterfaces() unsafe.Pointer {
	return _CNCopySupportedInterfaces()
	}


// Informs Captive Network Support that the device is not authenticated on a captive network. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOffline
func CNMarkPortalOffline(interfaceName unsafe.Pointer) unsafe.Pointer {
	return _CNMarkPortalOffline(interfaceName)
	}


// Informs Captive Network Support that the application has successfully authenticated the device to a captive network. Captive Network Support notifies the rest of the system that WiFi is a viable interface. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOnline
func CNMarkPortalOnline(interfaceName unsafe.Pointer) unsafe.Pointer {
	return _CNMarkPortalOnline(interfaceName)
	}


// Specifies an updated list of captive network SSIDs that the application performs authentication on. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNSetSupportedSSIDs
func CNSetSupportedSSIDs(ssidArray unsafe.Pointer) unsafe.Pointer {
	return _CNSetSupportedSSIDs(ssidArray)
	}


// Returns the list of options for the specified application ID. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesCopyApplicationOptions
func DHCPClientPreferencesCopyApplicationOptions(applicationID unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	return _DHCPClientPreferencesCopyApplicationOptions(applicationID, count)
	}


// Updates the DHCP client preferences to include the specified list of options for the specified application ID. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesSetApplicationOptions
func DHCPClientPreferencesSetApplicationOptions(applicationID unsafe.Pointer, options unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	return _DHCPClientPreferencesSetApplicationOptions(applicationID, options, count)
	}


// Returns the lease expiration time data. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseExpirationTime
func DHCPInfoGetLeaseExpirationTime(info unsafe.Pointer) unsafe.Pointer {
	return _DHCPInfoGetLeaseExpirationTime(info)
	}


// Returns the lease start time data. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseStartTime
func DHCPInfoGetLeaseStartTime(info unsafe.Pointer) unsafe.Pointer {
	return _DHCPInfoGetLeaseStartTime(info)
	}


// Returns DHCP option data, if present. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetOptionData
func DHCPInfoGetOptionData(info unsafe.Pointer, code unsafe.Pointer) unsafe.Pointer {
	return _DHCPInfoGetOptionData(info, code)
	}


// Returns all Ethernet bond interfaces on the system. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAll(_:)
func SCBondInterfaceCopyAll(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceCopyAll(prefs)
	}


// Returns all network capable devices on the system that can be added to an Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAvailableMemberInterfaces(_:)
func SCBondInterfaceCopyAvailableMemberInterfaces(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceCopyAvailableMemberInterfaces(prefs)
	}


// Returns the status of the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyStatus(_:)
func SCBondInterfaceCopyStatus(bond unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceCopyStatus(bond)
	}


// Creates a new Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCreate(_:)
func SCBondInterfaceCreate(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceCreate(prefs)
	}


// Returns the member interfaces for the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetMemberInterfaces(_:)
func SCBondInterfaceGetMemberInterfaces(bond unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceGetMemberInterfaces(bond)
	}


// Returns the configuration settings associated with the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetOptions(_:)
func SCBondInterfaceGetOptions(bond unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceGetOptions(bond)
	}


// Removes the Ethernet bond interface from the configuration. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceRemove(_:)
func SCBondInterfaceRemove(bond unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceRemove(bond)
	}


// Sets the localized display name for the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetLocalizedDisplayName(_:_:)
func SCBondInterfaceSetLocalizedDisplayName(bond unsafe.Pointer, newName unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceSetLocalizedDisplayName(bond, newName)
	}


// Sets the member interfaces for the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetMemberInterfaces(_:_:)
func SCBondInterfaceSetMemberInterfaces(bond unsafe.Pointer, members unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceSetMemberInterfaces(bond, members)
	}


// Sets the configuration settings for the specified Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetOptions(_:_:)
func SCBondInterfaceSetOptions(bond unsafe.Pointer, newOptions unsafe.Pointer) unsafe.Pointer {
	return _SCBondInterfaceSetOptions(bond, newOptions)
	}


// Returns the status of the specified member interface of an Ethernet bond or the status of the bond as a whole. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetInterfaceStatus(_:_:)
func SCBondStatusGetInterfaceStatus(bondStatus unsafe.Pointer, interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCBondStatusGetInterfaceStatus(bondStatus, interface_)
	}


// Returns the member interfaces that are represented with the Ethernet bond interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetMemberInterfaces(_:)
func SCBondStatusGetMemberInterfaces(bondStatus unsafe.Pointer) unsafe.Pointer {
	return _SCBondStatusGetMemberInterfaces(bondStatus)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetTypeID()
func SCBondStatusGetTypeID() unsafe.Pointer {
	return _SCBondStatusGetTypeID()
	}


// Returns an error or status code associated with the most recent function call. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCCopyLastError()
func SCCopyLastError() unsafe.Pointer {
	return _SCCopyLastError()
	}


// Temporarily adds the specified key-value pair to the dynamic store, if no such key already exists. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddTemporaryValue(_:_:_:)
func SCDynamicStoreAddTemporaryValue(store unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreAddTemporaryValue(store, key, value)
	}


// Adds the specified key-value pair to the dynamic store, if no such key already exists. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddValue(_:_:_:)
func SCDynamicStoreAddValue(store unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreAddValue(store, key, value)
	}


// Returns the current computer name. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyComputerName(_:_:)
func SCDynamicStoreCopyComputerName(store unsafe.Pointer, nameEncoding unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyComputerName(store, nameEncoding)
	}


// Returns information about the user currently logged into the system. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyConsoleUser(_:_:_:)
func SCDynamicStoreCopyConsoleUser(store unsafe.Pointer, uid unsafe.Pointer, gid unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyConsoleUser(store, uid, gid)
	}


// Returns the DHCP information for the specified service. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyDHCPInfo
func SCDynamicStoreCopyDHCPInfo(store unsafe.Pointer, serviceID unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyDHCPInfo(store, serviceID)
	}


// Returns the keys that represent the current dynamic store entries that match the specified pattern. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyKeyList(_:_:)
func SCDynamicStoreCopyKeyList(store unsafe.Pointer, pattern unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyKeyList(store, pattern)
	}


// Returns the current local host name. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocalHostName(_:)
func SCDynamicStoreCopyLocalHostName(store unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyLocalHostName(store)
	}


// Returns the current location identifier. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocation(_:)
func SCDynamicStoreCopyLocation(store unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyLocation(store)
	}


// Returns the key-value pairs that match the specified keys and key patterns. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyMultiple(_:_:_:)
func SCDynamicStoreCopyMultiple(store unsafe.Pointer, keys unsafe.Pointer, patterns unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyMultiple(store, keys, patterns)
	}


// Returns the keys that have changed since the last call to this function. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyNotifiedKeys(_:)
func SCDynamicStoreCopyNotifiedKeys(store unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyNotifiedKeys(store)
	}


// Returns the key-value pairs that represent the current internet proxy settings. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyProxies(_:)
func SCDynamicStoreCopyProxies(store unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyProxies(store)
	}


// Returns the value associated with the specified key. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyValue(_:_:)
func SCDynamicStoreCopyValue(store unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCopyValue(store, key)
	}


// Creates a new session used to interact with the dynamic store maintained by the System Configuration server. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreate(_:_:_:_:)
func SCDynamicStoreCreate(allocator unsafe.Pointer, name unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCreate(allocator, name, callout, context)
	}


// Creates a run loop source object that can be added to the application’s run loop. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateRunLoopSource(_:_:_:)
func SCDynamicStoreCreateRunLoopSource(allocator unsafe.Pointer, store unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCreateRunLoopSource(allocator, store, order)
	}


// Creates a new session used to interact with the dynamic store maintained by the System Configuration server. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateWithOptions(_:_:_:_:_:)
func SCDynamicStoreCreateWithOptions(allocator unsafe.Pointer, name unsafe.Pointer, storeOptions unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreCreateWithOptions(allocator, name, storeOptions, callout, context)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreGetTypeID()
func SCDynamicStoreGetTypeID() unsafe.Pointer {
	return _SCDynamicStoreGetTypeID()
	}


// Creates a dynamic store key using the specified format. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreate
func SCDynamicStoreKeyCreate(allocator unsafe.Pointer, fmt unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreate(allocator, fmt)
	}


// Creates a key that can be used to receive notifications when the current computer name changes. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateComputerName(_:)
func SCDynamicStoreKeyCreateComputerName(allocator unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateComputerName(allocator)
	}


// Creates a key that can be used to receive notifications when the current console user changes. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateConsoleUser(_:)
func SCDynamicStoreKeyCreateConsoleUser(allocator unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateConsoleUser(allocator)
	}


// Creates a key that can be used to receive notifications when the entity changes. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateHostNames(_:)
func SCDynamicStoreKeyCreateHostNames(allocator unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateHostNames(allocator)
	}


// Creates a key that can be used to receive notifications when the location identifier changes. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateLocation(_:)
func SCDynamicStoreKeyCreateLocation(allocator unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateLocation(allocator)
	}


// Creates a dynamic store key that can be used to access a specific global (as opposed to a per-service or per-interface) network configuration entity. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkGlobalEntity(_:_:_:)
func SCDynamicStoreKeyCreateNetworkGlobalEntity(allocator unsafe.Pointer, domain unsafe.Pointer, entity unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateNetworkGlobalEntity(allocator, domain, entity)
	}


// Creates a dynamic store key that can be used to access the network interface configuration information in the dynamic store. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterface(_:_:)
func SCDynamicStoreKeyCreateNetworkInterface(allocator unsafe.Pointer, domain unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateNetworkInterface(allocator, domain)
	}


// Creates a dynamic store key that can be used to access the per-interface network configuration information in the dynamic store. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterfaceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator unsafe.Pointer, domain unsafe.Pointer, ifname unsafe.Pointer, entity unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator, domain, ifname, entity)
	}


// Creates a dynamic store key that can be used to access the per-service network configuration information. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkServiceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkServiceEntity(allocator unsafe.Pointer, domain unsafe.Pointer, serviceID unsafe.Pointer, entity unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateNetworkServiceEntity(allocator, domain, serviceID, entity)
	}


// Creates a key that can be used to receive notifications when the current network proxy settings are changed. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateProxies(_:)
func SCDynamicStoreKeyCreateProxies(allocator unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreKeyCreateProxies(allocator)
	}


// Causes a notification to be delivered for the specified key in the dynamic store. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreNotifyValue(_:_:)
func SCDynamicStoreNotifyValue(store unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreNotifyValue(store, key)
	}


// Removes the value of the specified key from the dynamic store. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreRemoveValue(_:_:)
func SCDynamicStoreRemoveValue(store unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreRemoveValue(store, key)
	}


// Initiates notifications for the notification keys, using the specified dispatch queue for the callback. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetDispatchQueue(_:_:)
func SCDynamicStoreSetDispatchQueue(store unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreSetDispatchQueue(store, queue)
	}


// Updates multiple values in the dynamic store. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetMultiple(_:_:_:_:)
func SCDynamicStoreSetMultiple(store unsafe.Pointer, keysToSet unsafe.Pointer, keysToRemove unsafe.Pointer, keysToNotify unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreSetMultiple(store, keysToSet, keysToRemove, keysToNotify)
	}


// Specifies a set of keys and key patterns that should be monitored for changes. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetNotificationKeys(_:_:_:)
func SCDynamicStoreSetNotificationKeys(store unsafe.Pointer, keys unsafe.Pointer, patterns unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreSetNotificationKeys(store, keys, patterns)
	}


// Adds or replaces a value in the dynamic store for the specified key. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetValue(_:_:_:)
func SCDynamicStoreSetValue(store unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreSetValue(store, key, value)
	}


// Returns an error or status code associated with the most recent function call. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCError()
func SCError() int {
	return _SCError()
	}


// Returns a string describing the specified status code or error code. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCErrorString(_:)
func SCErrorString(status int) unsafe.Pointer {
	return _SCErrorString(status)
	}


// Determines whether the specified network address is reachable using the current network configuration. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkCheckReachabilityByAddress
func SCNetworkCheckReachabilityByAddress(address unsafe.Pointer, addrlen unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkCheckReachabilityByAddress(address, addrlen, flags)
	}


// Determines whether the specified network host or node name is reachable using the current network configuration. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkCheckReachabilityByName
func SCNetworkCheckReachabilityByName(nodename unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkCheckReachabilityByName(nodename, flags)
	}


// Returns the extended status of the connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyExtendedStatus(_:)
func SCNetworkConnectionCopyExtendedStatus(connection unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyExtendedStatus(connection)
	}


// Returns the service ID associated with the specified network connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyServiceID(_:)
func SCNetworkConnectionCopyServiceID(connection unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyServiceID(connection)
	}


// Returns the statistics of the specified connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyStatistics(_:)
func SCNetworkConnectionCopyStatistics(connection unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyStatistics(connection)
	}


// Gets the user options used to start the specified connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserOptions(_:)
func SCNetworkConnectionCopyUserOptions(connection unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyUserOptions(connection)
	}


// Provides the default service ID and a dictionary of user options for the specified connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserPreferences(_:_:_:)
func SCNetworkConnectionCopyUserPreferences(selectionOptions unsafe.Pointer, serviceID unsafe.Pointer, userOptions unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyUserPreferences(selectionOptions, serviceID, userOptions)
	}


// Creates a new connection reference to use for getting the status or for connecting or disconnecting the associated service. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCreateWithServiceID(_:_:_:_:)
func SCNetworkConnectionCreateWithServiceID(allocator unsafe.Pointer, serviceID unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCreateWithServiceID(allocator, serviceID, callout, context)
	}


// Returns the status of the specified network connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetStatus(_:)
func SCNetworkConnectionGetStatus(connection unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionGetStatus(connection)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetTypeID()
func SCNetworkConnectionGetTypeID() unsafe.Pointer {
	return _SCNetworkConnectionGetTypeID()
	}


// Schedules the specified connection with the specified run loop. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionScheduleWithRunLoop(_:_:_:)
func SCNetworkConnectionScheduleWithRunLoop(connection unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionScheduleWithRunLoop(connection, runLoop, runLoopMode)
	}


// Specifies a dispatch queue to use for the connection’s callback function and enables notifications. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionSetDispatchQueue(_:_:)
func SCNetworkConnectionSetDispatchQueue(connection unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionSetDispatchQueue(connection, queue)
	}


// Starts the connection process for the specified network connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStart(_:_:_:)
func SCNetworkConnectionStart(connection unsafe.Pointer, userOptions unsafe.Pointer, linger unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionStart(connection, userOptions, linger)
	}


// Stops the connection process for the specified network connection. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStop(_:_:)
func SCNetworkConnectionStop(connection unsafe.Pointer, forceDisconnect unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionStop(connection, forceDisconnect)
	}


// Unschedules the specified connection from the specified run loop. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionUnscheduleFromRunLoop(_:_:_:)
func SCNetworkConnectionUnscheduleFromRunLoop(connection unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionUnscheduleFromRunLoop(connection, runLoop, runLoopMode)
	}


// Returns all network-capable interfaces on the system. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyAll()
func SCNetworkInterfaceCopyAll() unsafe.Pointer {
	return _SCNetworkInterfaceCopyAll()
	}


// Returns the current MTU setting and the range of allowable values for the specified network interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMTU(_:_:_:_:)
func SCNetworkInterfaceCopyMTU(interface_ unsafe.Pointer, mtu_cur unsafe.Pointer, mtu_min unsafe.Pointer, mtu_max unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMTU(interface_, mtu_cur, mtu_min, mtu_max)
	}


// Returns information media options for the specified network interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaOptions(_:_:_:_:_:)
func SCNetworkInterfaceCopyMediaOptions(interface_ unsafe.Pointer, current unsafe.Pointer, active unsafe.Pointer, available unsafe.Pointer, filter unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMediaOptions(interface_, current, active, available, filter)
	}


// Returns a list of available media options for the specified interface configuration options and subtype. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypeOptions(_:_:)
func SCNetworkInterfaceCopyMediaSubTypeOptions(available unsafe.Pointer, subType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMediaSubTypeOptions(available, subType)
	}


// Returns a list of available media subtypes for the specified interface configuration options. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypes(_:)
func SCNetworkInterfaceCopyMediaSubTypes(available unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMediaSubTypes(available)
	}


// Creates a new network interface layered on top of the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCreateWithInterface(_:_:)
func SCNetworkInterfaceCreateWithInterface(interface_ unsafe.Pointer, interfaceType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCreateWithInterface(interface_, interfaceType)
	}


// Sends a notification to interested network configuration agents to immediately retry their configuration. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceForceConfigurationRefresh(_:)
func SCNetworkInterfaceForceConfigurationRefresh(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceForceConfigurationRefresh(interface_)
	}


// Returns the BSD interface or device name for the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetBSDName(_:)
func SCNetworkInterfaceGetBSDName(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetBSDName(interface_)
	}


// Returns the configuration settings associated with the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetConfiguration(_:)
func SCNetworkInterfaceGetConfiguration(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetConfiguration(interface_)
	}


// Returns the extended configuration settings associated with the specified interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetExtendedConfiguration(_:_:)
func SCNetworkInterfaceGetExtendedConfiguration(interface_ unsafe.Pointer, extendedType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetExtendedConfiguration(interface_, extendedType)
	}


// Returns a displayable link layer address for the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetHardwareAddressString(_:)
func SCNetworkInterfaceGetHardwareAddressString(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetHardwareAddressString(interface_)
	}


// Returns the underlying interface, for layered network interfaces. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterface(_:)
func SCNetworkInterfaceGetInterface(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetInterface(interface_)
	}


// Returns the network interface type of the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterfaceType(_:)
func SCNetworkInterfaceGetInterfaceType(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetInterfaceType(interface_)
	}


// Returns the localized display name, such as “Ethernet” or “FireWire”, for the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetLocalizedDisplayName(_:)
func SCNetworkInterfaceGetLocalizedDisplayName(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetLocalizedDisplayName(interface_)
	}


// Identifies all of the network interface types, such as PPP, that can be layered on top of the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedInterfaceTypes(_:)
func SCNetworkInterfaceGetSupportedInterfaceTypes(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetSupportedInterfaceTypes(interface_)
	}


// Identifies all of the network protocol types, such as IPv4 and IPv6, that can be layered on top of the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedProtocolTypes(_:)
func SCNetworkInterfaceGetSupportedProtocolTypes(interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceGetSupportedProtocolTypes(interface_)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetTypeID()
func SCNetworkInterfaceGetTypeID() unsafe.Pointer {
	return _SCNetworkInterfaceGetTypeID()
	}


// Sends a notification to interested configuration agents to have them immediately retry their configuration over a particular network interface. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceRefreshConfiguration
func SCNetworkInterfaceRefreshConfiguration(ifName unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceRefreshConfiguration(ifName)
	}


// Stores the configuration settings for the specified interface. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetConfiguration(_:_:)
func SCNetworkInterfaceSetConfiguration(interface_ unsafe.Pointer, config unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceSetConfiguration(interface_, config)
	}


// Stores the extended configuration settings for the specified interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetExtendedConfiguration(_:_:_:)
func SCNetworkInterfaceSetExtendedConfiguration(interface_ unsafe.Pointer, extendedType unsafe.Pointer, config unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceSetExtendedConfiguration(interface_, extendedType, config)
	}


// Sets the requested MTU setting for the specified network interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMTU(_:_:)
func SCNetworkInterfaceSetMTU(interface_ unsafe.Pointer, mtu int) unsafe.Pointer {
	return _SCNetworkInterfaceSetMTU(interface_, mtu)
	}


// Sets the requested media subtype and options for the specified network interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMediaOptions(_:_:_:)
func SCNetworkInterfaceSetMediaOptions(interface_ unsafe.Pointer, subtype unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceSetMediaOptions(interface_, subtype, options)
	}


// Returns the configuration settings associated with the specified protocol. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetConfiguration(_:)
func SCNetworkProtocolGetConfiguration(protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolGetConfiguration(protocol_)
	}


// Returns a Boolean value indicating whether the specified protocol is enabled. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetEnabled(_:)
func SCNetworkProtocolGetEnabled(protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolGetEnabled(protocol_)
	}


// Returns the type of the specified network protocol. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetProtocolType(_:)
func SCNetworkProtocolGetProtocolType(protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolGetProtocolType(protocol_)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetTypeID()
func SCNetworkProtocolGetTypeID() unsafe.Pointer {
	return _SCNetworkProtocolGetTypeID()
	}


// Stores the configuration settings for the specified network protocol. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetConfiguration(_:_:)
func SCNetworkProtocolSetConfiguration(protocol_ unsafe.Pointer, config unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolSetConfiguration(protocol_, config)
	}


// Enables or disables the specified protocol. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetEnabled(_:_:)
func SCNetworkProtocolSetEnabled(protocol_ unsafe.Pointer, enabled unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolSetEnabled(protocol_, enabled)
	}


// Creates a reachability reference to the specified network address. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddress(_:_:)
func SCNetworkReachabilityCreateWithAddress(allocator unsafe.Pointer, address unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityCreateWithAddress(allocator, address)
	}


// Creates a reachability reference to the specified network address. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddressPair(_:_:_:)
func SCNetworkReachabilityCreateWithAddressPair(allocator unsafe.Pointer, localAddress unsafe.Pointer, remoteAddress unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityCreateWithAddressPair(allocator, localAddress, remoteAddress)
	}


// Creates a reachability reference to the specified network host or node name. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithName(_:_:)
func SCNetworkReachabilityCreateWithName(allocator unsafe.Pointer, nodename unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityCreateWithName(allocator, nodename)
	}


// Determines if the specified network target is reachable using the current network configuration. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetFlags(_:_:)
func SCNetworkReachabilityGetFlags(target unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityGetFlags(target, flags)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetTypeID()
func SCNetworkReachabilityGetTypeID() unsafe.Pointer {
	return _SCNetworkReachabilityGetTypeID()
	}


// Schedules the specified network target with the specified run loop and mode. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityScheduleWithRunLoop(_:_:_:)
func SCNetworkReachabilityScheduleWithRunLoop(target unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityScheduleWithRunLoop(target, runLoop, runLoopMode)
	}


// Assigns a client to the specified target, which receives callbacks when the reachability of the target changes. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetCallback(_:_:_:)
func SCNetworkReachabilitySetCallback(target unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilitySetCallback(target, callout, context)
	}


// Schedules callbacks for the specified target on the specified dispatch queue. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetDispatchQueue(_:_:)
func SCNetworkReachabilitySetDispatchQueue(target unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilitySetDispatchQueue(target, queue)
	}


// Unschedules the specified target from the specified run loop and mode. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityUnscheduleFromRunLoop(_:_:_:)
func SCNetworkReachabilityUnscheduleFromRunLoop(target unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityUnscheduleFromRunLoop(target, runLoop, runLoopMode)
	}


// Adds the network protocol of the specified type to the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceAddProtocolType(_:_:)
func SCNetworkServiceAddProtocolType(service unsafe.Pointer, protocolType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceAddProtocolType(service, protocolType)
	}


// Returns the network service with the specified identifier. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopy(_:_:)
func SCNetworkServiceCopy(prefs unsafe.Pointer, serviceID unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceCopy(prefs, serviceID)
	}


// Returns all available network services for the specified preferences. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyAll(_:)
func SCNetworkServiceCopyAll(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceCopyAll(prefs)
	}


// Returns the network protocol of the specified type for the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocol(_:_:)
func SCNetworkServiceCopyProtocol(service unsafe.Pointer, protocolType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceCopyProtocol(service, protocolType)
	}


// Returns all network protocols associated with the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocols(_:)
func SCNetworkServiceCopyProtocols(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceCopyProtocols(service)
	}


// Creates a new network service for the specified interface in the configuration. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCreate(_:_:)
func SCNetworkServiceCreate(prefs unsafe.Pointer, interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceCreate(prefs, interface_)
	}


// Establishes the default configuration for the specified network service. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceEstablishDefaultConfiguration(_:)
func SCNetworkServiceEstablishDefaultConfiguration(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceEstablishDefaultConfiguration(service)
	}


// Returns a Boolean value indicating whether the specified service is enabled. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetEnabled(_:)
func SCNetworkServiceGetEnabled(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceGetEnabled(service)
	}


// Returns the network interface associated with the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetInterface(_:)
func SCNetworkServiceGetInterface(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceGetInterface(service)
	}


// Returns the user-specified name associated with the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetName(_:)
func SCNetworkServiceGetName(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceGetName(service)
	}


// Returns the identifier for the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetServiceID(_:)
func SCNetworkServiceGetServiceID(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceGetServiceID(service)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetTypeID()
func SCNetworkServiceGetTypeID() unsafe.Pointer {
	return _SCNetworkServiceGetTypeID()
	}


// Removes the specified network service from the configuration. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemove(_:)
func SCNetworkServiceRemove(service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceRemove(service)
	}


// Removes the network protocol of the specified type from the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemoveProtocolType(_:_:)
func SCNetworkServiceRemoveProtocolType(service unsafe.Pointer, protocolType unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceRemoveProtocolType(service, protocolType)
	}


// Enables or disables the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetEnabled(_:_:)
func SCNetworkServiceSetEnabled(service unsafe.Pointer, enabled unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceSetEnabled(service, enabled)
	}


// Stores the user-specified name for the specified service. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetName(_:_:)
func SCNetworkServiceSetName(service unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceSetName(service, name)
	}


// Adds the specified network service to the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetAddService(_:_:)
func SCNetworkSetAddService(set unsafe.Pointer, service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetAddService(set, service)
	}


// Returns a Boolean value indicating whether the specified interface is represented by at least one network service in the specified set. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetContainsInterface(_:_:)
func SCNetworkSetContainsInterface(set unsafe.Pointer, interface_ unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetContainsInterface(set, interface_)
	}


// Returns the set with the specified identifier. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopy(_:_:)
func SCNetworkSetCopy(prefs unsafe.Pointer, setID unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetCopy(prefs, setID)
	}


// Returns all available sets for the specified preferences session. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyAll(_:)
func SCNetworkSetCopyAll(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetCopyAll(prefs)
	}


// Returns the current set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyCurrent(_:)
func SCNetworkSetCopyCurrent(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetCopyCurrent(prefs)
	}


// Returns all network services associated with the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyServices(_:)
func SCNetworkSetCopyServices(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetCopyServices(set)
	}


// Creates a new set in the configuration. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCreate(_:)
func SCNetworkSetCreate(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetCreate(prefs)
	}


// Returns the user-specified name associated with the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetName(_:)
func SCNetworkSetGetName(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetGetName(set)
	}


// Returns the user-specified ordering of network services within the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetServiceOrder(_:)
func SCNetworkSetGetServiceOrder(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetGetServiceOrder(set)
	}


// Returns the identifier for the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetSetID(_:)
func SCNetworkSetGetSetID(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetGetSetID(set)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetTypeID()
func SCNetworkSetGetTypeID() unsafe.Pointer {
	return _SCNetworkSetGetTypeID()
	}


// Removes the specified set from the configuration. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemove(_:)
func SCNetworkSetRemove(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetRemove(set)
	}


// Removes the specified network service from the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemoveService(_:_:)
func SCNetworkSetRemoveService(set unsafe.Pointer, service unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetRemoveService(set, service)
	}


// Specifies the set that should be the current set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetCurrent(_:)
func SCNetworkSetSetCurrent(set unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetSetCurrent(set)
	}


// Stores the user-specified name for the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetName(_:_:)
func SCNetworkSetSetName(set unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetSetName(set, name)
	}


// Stores the user-specified ordering of network services for the specified set. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetServiceOrder(_:_:)
func SCNetworkSetSetServiceOrder(set unsafe.Pointer, newOrder unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkSetSetServiceOrder(set, newOrder)
	}


// Associates the specified value with the specified preference key. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesAddValue(_:_:_:)
func SCPreferencesAddValue(prefs unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesAddValue(prefs, key, value)
	}


// Requests that the currently stored configuration preferences be applied to the active configuration. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesApplyChanges(_:)
func SCPreferencesApplyChanges(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesApplyChanges(prefs)
	}


// Commits changes made to the configuration preferences to persistent storage. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCommitChanges(_:)
func SCPreferencesCommitChanges(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesCommitChanges(prefs)
	}


// Returns the currently defined preference keys. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCopyKeyList(_:)
func SCPreferencesCopyKeyList(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesCopyKeyList(prefs)
	}


// Initiates access to the per-system set of configuration preferences. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreate(_:_:_:)
func SCPreferencesCreate(allocator unsafe.Pointer, name unsafe.Pointer, prefsID unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesCreate(allocator, name, prefsID)
	}


// Initiates access to the per-system set of configuration preferences with the specified authorization. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreateWithAuthorization(_:_:_:_:)
func SCPreferencesCreateWithAuthorization(allocator unsafe.Pointer, name unsafe.Pointer, prefsID unsafe.Pointer, authorization unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesCreateWithAuthorization(allocator, name, prefsID, authorization)
	}


// Returns a value that can be used to determine if the saved configuration preferences have changed. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetSignature(_:)
func SCPreferencesGetSignature(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesGetSignature(prefs)
	}


// Returns the type identifier of all instances. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetTypeID()
func SCPreferencesGetTypeID() unsafe.Pointer {
	return _SCPreferencesGetTypeID()
	}


// Retrieves the value associated with the specified preference key. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetValue(_:_:)
func SCPreferencesGetValue(prefs unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesGetValue(prefs, key)
	}


// Locks access to the configuration preferences. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesLock(_:_:)
func SCPreferencesLock(prefs unsafe.Pointer, wait unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesLock(prefs, wait)
	}


// Creates a new path component rooted at the specified path in the dictionary hierarchy. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathCreateUniqueChild(_:_:)
func SCPreferencesPathCreateUniqueChild(prefs unsafe.Pointer, prefix unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathCreateUniqueChild(prefs, prefix)
	}


// Returns the link associated with the specified path. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetLink(_:_:)
func SCPreferencesPathGetLink(prefs unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathGetLink(prefs, path)
	}


// Returns the dictionary associated with the specified path. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetValue(_:_:)
func SCPreferencesPathGetValue(prefs unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathGetValue(prefs, path)
	}


// Removes the data associated with the specified path. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathRemoveValue(_:_:)
func SCPreferencesPathRemoveValue(prefs unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathRemoveValue(prefs, path)
	}


// Associates a link to a second dictionary at the specified path. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetLink(_:_:_:)
func SCPreferencesPathSetLink(prefs unsafe.Pointer, path unsafe.Pointer, link unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathSetLink(prefs, path, link)
	}


// Associates the specified dictionary with the specified path. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetValue(_:_:_:)
func SCPreferencesPathSetValue(prefs unsafe.Pointer, path unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesPathSetValue(prefs, path, value)
	}


// Removes the data associated with the specified preference key. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesRemoveValue(_:_:)
func SCPreferencesRemoveValue(prefs unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesRemoveValue(prefs, key)
	}


// Schedules commit and apply notifications for the specified preferences session using the specified run loop and mode. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesScheduleWithRunLoop(_:_:_:)
func SCPreferencesScheduleWithRunLoop(prefs unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesScheduleWithRunLoop(prefs, runLoop, runLoopMode)
	}


// Assigns the specified callback to the specified preferences session. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetCallback(_:_:_:)
func SCPreferencesSetCallback(prefs unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetCallback(prefs, callout, context)
	}


// Sets the computer name preference to the specified name. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetComputerName(_:_:_:)
func SCPreferencesSetComputerName(prefs unsafe.Pointer, name unsafe.Pointer, nameEncoding unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetComputerName(prefs, name, nameEncoding)
	}


// Schedules commit and apply notifications for the specified preferences session using the specified dispatch queue. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetDispatchQueue(_:_:)
func SCPreferencesSetDispatchQueue(prefs unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetDispatchQueue(prefs, queue)
	}


// Sets the local host name to the specified name. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetLocalHostName(_:_:)
func SCPreferencesSetLocalHostName(prefs unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetLocalHostName(prefs, name)
	}


// Updates the data associated with the specified preference key with the specified value. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetValue(_:_:_:)
func SCPreferencesSetValue(prefs unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetValue(prefs, key, value)
	}


// Synchronizes accessed preferences with committed changes. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSynchronize(_:)
func SCPreferencesSynchronize(prefs unsafe.Pointer) {
	_SCPreferencesSynchronize(prefs)
	}


// Releases exclusive access to the configuration preferences. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnlock(_:)
func SCPreferencesUnlock(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesUnlock(prefs)
	}


// Unschedules commit and apply notifications for the specified preferences session from the specified run loop and mode. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnscheduleFromRunLoop(_:_:_:)
func SCPreferencesUnscheduleFromRunLoop(prefs unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesUnscheduleFromRunLoop(prefs, runLoop, runLoopMode)
	}


// Returns all virtual LAN (VLAN) interfaces on the system. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAll(_:)
func SCVLANInterfaceCopyAll(prefs unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceCopyAll(prefs)
	}


// Returns the network capable devices on the system that can be associated with a virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAvailablePhysicalInterfaces()
func SCVLANInterfaceCopyAvailablePhysicalInterfaces() unsafe.Pointer {
	return _SCVLANInterfaceCopyAvailablePhysicalInterfaces()
	}


// Creates a new virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCreate(_:_:_:)
func SCVLANInterfaceCreate(prefs unsafe.Pointer, physical unsafe.Pointer, tag unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceCreate(prefs, physical, tag)
	}


// Returns the configuration settings associated with the virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetOptions(_:)
func SCVLANInterfaceGetOptions(vlan unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceGetOptions(vlan)
	}


// Returns the physical interface for the specified virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetPhysicalInterface(_:)
func SCVLANInterfaceGetPhysicalInterface(vlan unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceGetPhysicalInterface(vlan)
	}


// Returns the tag for the specified virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetTag(_:)
func SCVLANInterfaceGetTag(vlan unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceGetTag(vlan)
	}


// Removes the virtual LAN (VLAN) interface from the configuration. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceRemove(_:)
func SCVLANInterfaceRemove(vlan unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceRemove(vlan)
	}


// Sets the localized display name for the specified virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetLocalizedDisplayName(_:_:)
func SCVLANInterfaceSetLocalizedDisplayName(vlan unsafe.Pointer, newName unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceSetLocalizedDisplayName(vlan, newName)
	}


// Sets the specified configuration settings for the specified virtual LAN (VLAN) interface. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetOptions(_:_:)
func SCVLANInterfaceSetOptions(vlan unsafe.Pointer, newOptions unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceSetOptions(vlan, newOptions)
	}


// Updates the specified virtual LAN (VLAN) interface with the specified information. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetPhysicalInterfaceAndTag(_:_:_:)
func SCVLANInterfaceSetPhysicalInterfaceAndTag(vlan unsafe.Pointer, physical unsafe.Pointer, tag unsafe.Pointer) unsafe.Pointer {
	return _SCVLANInterfaceSetPhysicalInterfaceAndTag(vlan, physical, tag)
	}




