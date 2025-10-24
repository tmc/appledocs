// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterDoorLock */


/* debug [class_header]: Header for MTRBaseClusterDoorLock */
// The class instance for the [MTRBaseClusterDoorLock] class.
var (
	MTRBaseClusterDoorLockClass     _MTRBaseClusterDoorLockClass
	MTRBaseClusterDoorLockClassOnce sync.Once
)

func getMTRBaseClusterDoorLockClass() _MTRBaseClusterDoorLockClass {
	MTRBaseClusterDoorLockClassOnce.Do(func() {
		MTRBaseClusterDoorLockClass = _MTRBaseClusterDoorLockClass{objc.GetClass("MTRBaseClusterDoorLock")}
	})
	return MTRBaseClusterDoorLockClass
}

type _MTRBaseClusterDoorLockClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterDoorLock */
// An interface definition for the [MTRBaseClusterDoorLock] class.
type IMTRBaseClusterDoorLock interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterDoorLock */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterDoorLock */
	// methods:
	ClearAliroReaderConfigWithCompletion(completion unsafe.Pointer)
	ClearAliroReaderConfigWithParamsCompletion(params IMTRDoorLockClusterClearAliroReaderConfigParams, completion unsafe.Pointer)
	ClearCredentialWithParamsCompletion(params IMTRDoorLockClusterClearCredentialParams, completion unsafe.Pointer)
	ClearHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearHolidayScheduleParams, completion unsafe.Pointer)
	ClearUserWithParamsCompletion(params IMTRDoorLockClusterClearUserParams, completion unsafe.Pointer)
	ClearWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearWeekDayScheduleParams, completion unsafe.Pointer)
	ClearYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearYearDayScheduleParams, completion unsafe.Pointer)
	GetCredentialStatusWithParamsCompletion(params IMTRDoorLockClusterGetCredentialStatusParams, completion unsafe.Pointer)
	GetHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetHolidayScheduleParams, completion unsafe.Pointer)
	GetUserWithParamsCompletion(params IMTRDoorLockClusterGetUserParams, completion unsafe.Pointer)
	GetWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetWeekDayScheduleParams, completion unsafe.Pointer)
	GetYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetYearDayScheduleParams, completion unsafe.Pointer)
	LockDoorWithCompletion(completion unsafe.Pointer)
	LockDoorWithParamsCompletion(params IMTRDoorLockClusterLockDoorParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeActuatorEnabledWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroBLEAdvertisingVersionWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroGroupResolvingKeyWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroReaderGroupIdentifierWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroReaderGroupSubIdentifierWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroReaderVerificationKeyWithCompletion(completion unsafe.Pointer)
	ReadAttributeAliroSupportedBLEUWBProtocolVersionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAutoRelockTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCredentialRulesSupportWithCompletion(completion unsafe.Pointer)
	ReadAttributeDefaultConfigurationRegisterWithCompletion(completion unsafe.Pointer)
	ReadAttributeDoorClosedEventsWithCompletion(completion unsafe.Pointer)
	ReadAttributeDoorOpenEventsWithCompletion(completion unsafe.Pointer)
	ReadAttributeDoorStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnableInsideStatusLEDWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnableLocalProgrammingWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnableOneTouchLockingWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnablePrivacyModeButtonWithCompletion(completion unsafe.Pointer)
	ReadAttributeExpiringUserTimeoutWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeLanguageWithCompletion(completion unsafe.Pointer)
	ReadAttributeLEDSettingsWithCompletion(completion unsafe.Pointer)
	ReadAttributeLocalProgrammingFeaturesWithCompletion(completion unsafe.Pointer)
	ReadAttributeLockStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeLockTypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxPINCodeLengthWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxRFIDCodeLengthWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinPINCodeLengthWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinRFIDCodeLengthWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfAliroCredentialIssuerKeysSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfAliroEndpointKeysSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfCredentialsSupportedPerUserWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfHolidaySchedulesSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfPINUsersSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfRFIDUsersSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfTotalUsersSupportedWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithCompletion(completion unsafe.Pointer)
	ReadAttributeOpenPeriodWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperatingModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeRequirePINforRemoteOperationWithCompletion(completion unsafe.Pointer)
	ReadAttributeSendPINOverTheAirWithCompletion(completion unsafe.Pointer)
	ReadAttributeSoundVolumeWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedOperatingModesWithCompletion(completion unsafe.Pointer)
	ReadAttributeUserCodeTemporaryDisableTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeWrongCodeEntryLimitWithCompletion(completion unsafe.Pointer)
	SetAliroReaderConfigWithParamsCompletion(params IMTRDoorLockClusterSetAliroReaderConfigParams, completion unsafe.Pointer)
	SetCredentialWithParamsCompletion(params IMTRDoorLockClusterSetCredentialParams, completion unsafe.Pointer)
	SetHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetHolidayScheduleParams, completion unsafe.Pointer)
	SetUserWithParamsCompletion(params IMTRDoorLockClusterSetUserParams, completion unsafe.Pointer)
	SetWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetWeekDayScheduleParams, completion unsafe.Pointer)
	SetYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetYearDayScheduleParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActuatorEnabledWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroBLEAdvertisingVersionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroGroupResolvingKeyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroReaderGroupIdentifierWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroReaderGroupSubIdentifierWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroReaderVerificationKeyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAliroSupportedBLEUWBProtocolVersionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAutoRelockTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCredentialRulesSupportWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDefaultConfigurationRegisterWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDoorClosedEventsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDoorOpenEventsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDoorStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnableInsideStatusLEDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnableLocalProgrammingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnableOneTouchLockingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnablePrivacyModeButtonWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeExpiringUserTimeoutWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLanguageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLEDSettingsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLocalProgrammingFeaturesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLockStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLockTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxPINCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinPINCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfAliroCredentialIssuerKeysSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfAliroEndpointKeysSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfCredentialsSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfHolidaySchedulesSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfPINUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfRFIDUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfTotalUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfWeekDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfYearDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOpenPeriodWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperatingModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRequirePINforRemoteOperationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSendPINOverTheAirWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSoundVolumeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedOperatingModesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUserCodeTemporaryDisableTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeWrongCodeEntryLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	UnboltDoorWithCompletion(completion unsafe.Pointer)
	UnboltDoorWithParamsCompletion(params IMTRDoorLockClusterUnboltDoorParams, completion unsafe.Pointer)
	UnlockDoorWithCompletion(completion unsafe.Pointer)
	UnlockDoorWithParamsCompletion(params IMTRDoorLockClusterUnlockDoorParams, completion unsafe.Pointer)
	UnlockWithTimeoutWithParamsCompletion(params IMTRDoorLockClusterUnlockWithTimeoutParams, completion unsafe.Pointer)
	WriteAttributeAutoRelockTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeAutoRelockTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeDoorClosedEventsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeDoorClosedEventsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeDoorOpenEventsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeDoorOpenEventsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeEnableInsideStatusLEDWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeEnableInsideStatusLEDWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeEnableLocalProgrammingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeEnableLocalProgrammingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeEnableOneTouchLockingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeEnableOneTouchLockingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeEnablePrivacyModeButtonWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeEnablePrivacyModeButtonWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeExpiringUserTimeoutWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeExpiringUserTimeoutWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeLanguageWithValueCompletion(value objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
	WriteAttributeLanguageWithValueParamsCompletion(value objc.IObject /* cross-framework: NSString */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeLEDSettingsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeLEDSettingsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeLocalProgrammingFeaturesWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeLocalProgrammingFeaturesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOpenPeriodWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOpenPeriodWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOperatingModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOperatingModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeRequirePINforRemoteOperationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeRequirePINforRemoteOperationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeSendPINOverTheAirWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSendPINOverTheAirWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeSoundVolumeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSoundVolumeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeUserCodeTemporaryDisableTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeUserCodeTemporaryDisableTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeWrongCodeEntryLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeWrongCodeEntryLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterDoorLock */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDoorLockClass) Alloc() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterDoorLockClass) New() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDoorLock) Init() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDoorLock) Autorelease() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDoorLock creates a new MTRBaseClusterDoorLock instance.
func NewMTRBaseClusterDoorLock() MTRBaseClusterDoorLock {
	return getMTRBaseClusterDoorLockClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterDoorLock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock
type MTRBaseClusterDoorLock struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDoorLockFrom constructs a [MTRBaseClusterDoorLock] from an unsafe.Pointer.
func MTRBaseClusterDoorLockFrom(ptr unsafe.Pointer) MTRBaseClusterDoorLock {
	return MTRBaseClusterDoorLock{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterDoorLock */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/init(device:endpointID:queue:)
func NewMTRBaseClusterDoorLockWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterDoorLock {
	instance := getMTRBaseClusterDoorLockClass().Alloc()
	rv := objc.Send[MTRBaseClusterDoorLock](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterDoorLockWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/init(device:endpoint:queue:)
func NewMTRBaseClusterDoorLockWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterDoorLock {
	instance := getMTRBaseClusterDoorLockClass().Alloc()
	rv := objc.Send[MTRBaseClusterDoorLock](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterDoorLockWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterDoorLock */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeActuatorEnabled(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeActuatorEnabledWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActuatorEnabledWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeActuatorEnabledWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeActuatorEnabled(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeActuatorEnabledWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActuatorEnabledWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeActuatorEnabledWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroBLEAdvertisingVersion(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroBLEAdvertisingVersionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroBLEAdvertisingVersionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroBLEAdvertisingVersionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroExpeditedTransactionSupportedProtocolVersions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroGroupResolvingKey(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroGroupResolvingKeyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroGroupResolvingKeyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroGroupResolvingKeyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderGroupIdentifier(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroReaderGroupIdentifierWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroReaderGroupIdentifierWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroReaderGroupIdentifierWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderGroupSubIdentifier(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroReaderGroupSubIdentifierWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroReaderGroupSubIdentifierWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroReaderGroupSubIdentifierWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderVerificationKey(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroReaderVerificationKeyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroReaderVerificationKeyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroReaderVerificationKeyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroSupportedBLEUWBProtocolVersions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAliroSupportedBLEUWBProtocolVersionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAliroSupportedBLEUWBProtocolVersionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAliroSupportedBLEUWBProtocolVersionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAutoRelockTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAutoRelockTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAutoRelockTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAutoRelockTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAutoRelockTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeAutoRelockTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAutoRelockTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAutoRelockTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeCredentialRulesSupport(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeCredentialRulesSupportWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCredentialRulesSupportWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCredentialRulesSupportWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeCredentialRulesSupport(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeCredentialRulesSupportWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCredentialRulesSupportWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCredentialRulesSupportWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDefaultConfigurationRegister(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDefaultConfigurationRegisterWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultConfigurationRegisterWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDefaultConfigurationRegisterWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDefaultConfigurationRegister(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDefaultConfigurationRegisterWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultConfigurationRegisterWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDefaultConfigurationRegisterWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorClosedEvents(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorClosedEventsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorClosedEventsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorClosedEventsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorClosedEvents(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorClosedEventsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorClosedEventsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorClosedEventsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorOpenEvents(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorOpenEventsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorOpenEventsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorOpenEventsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorOpenEvents(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorOpenEventsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorOpenEventsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorOpenEventsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorState(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorStateWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorStateWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorStateWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeDoorStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDoorStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDoorStateWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableInsideStatusLED(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableInsideStatusLEDWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableInsideStatusLEDWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableInsideStatusLEDWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableInsideStatusLED(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableInsideStatusLEDWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableInsideStatusLEDWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableInsideStatusLEDWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableLocalProgramming(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableLocalProgrammingWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableLocalProgrammingWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableLocalProgrammingWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableLocalProgramming(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableLocalProgrammingWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableLocalProgrammingWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableLocalProgrammingWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableOneTouchLocking(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableOneTouchLockingWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableOneTouchLockingWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableOneTouchLockingWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableOneTouchLocking(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnableOneTouchLockingWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnableOneTouchLockingWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnableOneTouchLockingWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnablePrivacyModeButton(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnablePrivacyModeButtonWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnablePrivacyModeButtonWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnablePrivacyModeButtonWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnablePrivacyModeButton(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeEnablePrivacyModeButtonWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnablePrivacyModeButtonWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnablePrivacyModeButtonWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeExpiringUserTimeout(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeExpiringUserTimeoutWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeExpiringUserTimeoutWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeExpiringUserTimeoutWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeExpiringUserTimeout(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeExpiringUserTimeoutWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeExpiringUserTimeoutWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeExpiringUserTimeoutWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLanguage(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLanguageWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLanguageWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLanguageWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLanguage(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLanguageWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLanguageWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLanguageWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLEDSettings(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLEDSettingsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLEDSettingsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLEDSettingsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLEDSettings(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLEDSettingsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLEDSettingsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLEDSettingsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLocalProgrammingFeatures(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLocalProgrammingFeaturesWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalProgrammingFeaturesWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalProgrammingFeaturesWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLocalProgrammingFeatures(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLocalProgrammingFeaturesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalProgrammingFeaturesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalProgrammingFeaturesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockState(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLockStateWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLockStateWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLockStateWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLockStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLockStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLockStateWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLockTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLockTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLockTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeLockTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLockTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLockTypeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxPINCodeLength(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMaxPINCodeLengthWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxPINCodeLengthWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxPINCodeLengthWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxPINCodeLength(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMaxPINCodeLengthWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxPINCodeLengthWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxPINCodeLengthWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxRFIDCodeLength(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMaxRFIDCodeLengthWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxRFIDCodeLengthWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxRFIDCodeLengthWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxRFIDCodeLength(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMaxRFIDCodeLengthWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxRFIDCodeLengthWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxRFIDCodeLengthWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinPINCodeLength(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMinPINCodeLengthWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinPINCodeLengthWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinPINCodeLengthWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinPINCodeLength(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMinPINCodeLengthWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinPINCodeLengthWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinPINCodeLengthWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinRFIDCodeLength(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMinRFIDCodeLengthWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinRFIDCodeLengthWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinRFIDCodeLengthWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinRFIDCodeLength(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeMinRFIDCodeLengthWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinRFIDCodeLengthWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinRFIDCodeLengthWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfAliroCredentialIssuerKeysSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfAliroCredentialIssuerKeysSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfAliroCredentialIssuerKeysSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfAliroCredentialIssuerKeysSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfAliroEndpointKeysSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfAliroEndpointKeysSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfAliroEndpointKeysSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfAliroEndpointKeysSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfCredentialsSupportedPerUser(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfCredentialsSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfCredentialsSupportedPerUserWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfCredentialsSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfCredentialsSupportedPerUser(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfCredentialsSupportedPerUserWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfCredentialsSupportedPerUserWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfCredentialsSupportedPerUserWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfHolidaySchedulesSupported(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfHolidaySchedulesSupportedWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfHolidaySchedulesSupportedWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfHolidaySchedulesSupportedWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfHolidaySchedulesSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfHolidaySchedulesSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfHolidaySchedulesSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfHolidaySchedulesSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfPINUsersSupported(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfPINUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfPINUsersSupportedWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfPINUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfPINUsersSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfPINUsersSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfPINUsersSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfPINUsersSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfRFIDUsersSupported(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfRFIDUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfRFIDUsersSupportedWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfRFIDUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfRFIDUsersSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfRFIDUsersSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfRFIDUsersSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfRFIDUsersSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfTotalUsersSupported(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfTotalUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfTotalUsersSupportedWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfTotalUsersSupportedWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfTotalUsersSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfTotalUsersSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfTotalUsersSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfTotalUsersSupportedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfWeekDaySchedulesSupportedPerUser(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfWeekDaySchedulesSupportedPerUserWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfWeekDaySchedulesSupportedPerUser(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfWeekDaySchedulesSupportedPerUserWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfYearDaySchedulesSupportedPerUser(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfYearDaySchedulesSupportedPerUserWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfYearDaySchedulesSupportedPerUser(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfYearDaySchedulesSupportedPerUserWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOpenPeriod(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeOpenPeriodWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOpenPeriodWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOpenPeriodWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOpenPeriod(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeOpenPeriodWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOpenPeriodWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOpenPeriodWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOperatingMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeOperatingModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperatingModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOperatingModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOperatingMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeOperatingModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperatingModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOperatingModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeRequirePINforRemoteOperation(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeRequirePINforRemoteOperationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRequirePINforRemoteOperationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRequirePINforRemoteOperationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeRequirePINforRemoteOperation(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeRequirePINforRemoteOperationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRequirePINforRemoteOperationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRequirePINforRemoteOperationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSendPINOverTheAir(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSendPINOverTheAirWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSendPINOverTheAirWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSendPINOverTheAirWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSendPINOverTheAir(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSendPINOverTheAirWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSendPINOverTheAirWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSendPINOverTheAirWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSoundVolume(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSoundVolumeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSoundVolumeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSoundVolumeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSoundVolume(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSoundVolumeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSoundVolumeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSoundVolumeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSupportedOperatingModes(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSupportedOperatingModesWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedOperatingModesWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSupportedOperatingModesWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSupportedOperatingModes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeSupportedOperatingModesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedOperatingModesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSupportedOperatingModesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeUserCodeTemporaryDisableTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeUserCodeTemporaryDisableTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUserCodeTemporaryDisableTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUserCodeTemporaryDisableTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeUserCodeTemporaryDisableTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeUserCodeTemporaryDisableTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUserCodeTemporaryDisableTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUserCodeTemporaryDisableTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeWrongCodeEntryLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeWrongCodeEntryLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWrongCodeEntryLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWrongCodeEntryLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeWrongCodeEntryLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDoorLockClass) ReadAttributeWrongCodeEntryLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWrongCodeEntryLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWrongCodeEntryLimitWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterDoorLock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterDoorLock */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearAliroReaderConfig(completion:)
func (m_ MTRBaseClusterDoorLock) ClearAliroReaderConfigWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearAliroReaderConfigWithCompletion:"), completion)
}/* debug [instance_methods/method]: ClearAliroReaderConfigWithCompletion */


// Command ClearAliroReaderConfig
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearAliroReaderConfig(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearAliroReaderConfigWithParamsCompletion(params IMTRDoorLockClusterClearAliroReaderConfigParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearAliroReaderConfigWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearAliroReaderConfigWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearCredential(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearCredentialWithParamsCompletion(params IMTRDoorLockClusterClearCredentialParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearCredentialWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearCredentialWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearHolidaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearHolidayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearHolidayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearHolidayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearUser(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearUserWithParamsCompletion(params IMTRDoorLockClusterClearUserParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearUserWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearUserWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearWeekDaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearWeekDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearWeekDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearWeekDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/clearYearDaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) ClearYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterClearYearDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearYearDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearYearDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/getCredentialStatus(with:completion:)
func (m_ MTRBaseClusterDoorLock) GetCredentialStatusWithParamsCompletion(params IMTRDoorLockClusterGetCredentialStatusParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getCredentialStatusWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetCredentialStatusWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/getHolidaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) GetHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetHolidayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getHolidayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetHolidayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/getUserWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) GetUserWithParamsCompletion(params IMTRDoorLockClusterGetUserParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getUserWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetUserWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/getWeekDaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) GetWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetWeekDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getWeekDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetWeekDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/getYearDaySchedule(with:completion:)
func (m_ MTRBaseClusterDoorLock) GetYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterGetYearDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getYearDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetYearDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/lockDoor(completion:)
func (m_ MTRBaseClusterDoorLock) LockDoorWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("lockDoorWithCompletion:"), completion)
}/* debug [instance_methods/method]: LockDoorWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/lockDoor(with:completion:)
func (m_ MTRBaseClusterDoorLock) LockDoorWithParamsCompletion(params IMTRDoorLockClusterLockDoorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("lockDoorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: LockDoorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeActuatorEnabled(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeActuatorEnabledWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActuatorEnabledWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeActuatorEnabledWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroBLEAdvertisingVersion(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroBLEAdvertisingVersionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroBLEAdvertisingVersionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroBLEAdvertisingVersionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroExpeditedTransactionSupportedProtocolVersions(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroGroupResolvingKey(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroGroupResolvingKeyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroGroupResolvingKeyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroGroupResolvingKeyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderGroupIdentifier(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroReaderGroupIdentifierWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroReaderGroupIdentifierWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroReaderGroupIdentifierWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderGroupSubIdentifier(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroReaderGroupSubIdentifierWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroReaderGroupSubIdentifierWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroReaderGroupSubIdentifierWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroReaderVerificationKey(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroReaderVerificationKeyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroReaderVerificationKeyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroReaderVerificationKeyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAliroSupportedBLEUWBProtocolVersions(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAliroSupportedBLEUWBProtocolVersionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAliroSupportedBLEUWBProtocolVersionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAliroSupportedBLEUWBProtocolVersionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeAutoRelockTime(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeAutoRelockTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAutoRelockTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAutoRelockTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeCredentialRulesSupport(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeCredentialRulesSupportWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCredentialRulesSupportWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCredentialRulesSupportWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDefaultConfigurationRegister(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeDefaultConfigurationRegisterWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDefaultConfigurationRegisterWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDefaultConfigurationRegisterWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorClosedEvents(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeDoorClosedEventsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDoorClosedEventsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDoorClosedEventsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorOpenEvents(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeDoorOpenEventsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDoorOpenEventsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDoorOpenEventsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeDoorState(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeDoorStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDoorStateWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDoorStateWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableInsideStatusLED(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeEnableInsideStatusLEDWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnableInsideStatusLEDWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnableInsideStatusLEDWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableLocalProgramming(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeEnableLocalProgrammingWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnableLocalProgrammingWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnableLocalProgrammingWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnableOneTouchLocking(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeEnableOneTouchLockingWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnableOneTouchLockingWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnableOneTouchLockingWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeEnablePrivacyModeButton(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeEnablePrivacyModeButtonWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnablePrivacyModeButtonWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnablePrivacyModeButtonWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeExpiringUserTimeout(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeExpiringUserTimeoutWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeExpiringUserTimeoutWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeExpiringUserTimeoutWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLanguage(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeLanguageWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLanguageWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLanguageWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLEDSettings(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeLEDSettingsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLEDSettingsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLEDSettingsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLocalProgrammingFeatures(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeLocalProgrammingFeaturesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLocalProgrammingFeaturesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLocalProgrammingFeaturesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockState(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeLockStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLockStateWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLockStateWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeLockType(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeLockTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLockTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLockTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxPINCodeLength(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeMaxPINCodeLengthWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxPINCodeLengthWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxPINCodeLengthWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMaxRFIDCodeLength(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeMaxRFIDCodeLengthWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxRFIDCodeLengthWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxRFIDCodeLengthWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinPINCodeLength(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeMinPINCodeLengthWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinPINCodeLengthWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinPINCodeLengthWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeMinRFIDCodeLength(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeMinRFIDCodeLengthWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinRFIDCodeLengthWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinRFIDCodeLengthWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfAliroCredentialIssuerKeysSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfAliroCredentialIssuerKeysSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfAliroCredentialIssuerKeysSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfAliroCredentialIssuerKeysSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfAliroEndpointKeysSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfAliroEndpointKeysSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfAliroEndpointKeysSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfAliroEndpointKeysSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfCredentialsSupportedPerUser(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfCredentialsSupportedPerUserWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfCredentialsSupportedPerUserWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfCredentialsSupportedPerUserWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfHolidaySchedulesSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfHolidaySchedulesSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfHolidaySchedulesSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfHolidaySchedulesSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfPINUsersSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfPINUsersSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfPINUsersSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfPINUsersSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfRFIDUsersSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfRFIDUsersSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfRFIDUsersSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfRFIDUsersSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfTotalUsersSupported(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfTotalUsersSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfTotalUsersSupportedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfTotalUsersSupportedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfWeekDaySchedulesSupportedPerUser(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfWeekDaySchedulesSupportedPerUserWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfWeekDaySchedulesSupportedPerUserWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeNumberOfYearDaySchedulesSupportedPerUser(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfYearDaySchedulesSupportedPerUserWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfYearDaySchedulesSupportedPerUserWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOpenPeriod(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeOpenPeriodWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOpenPeriodWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOpenPeriodWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeOperatingMode(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeOperatingModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperatingModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOperatingModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeRequirePINforRemoteOperation(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeRequirePINforRemoteOperationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRequirePINforRemoteOperationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeRequirePINforRemoteOperationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSendPINOverTheAir(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeSendPINOverTheAirWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSendPINOverTheAirWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSendPINOverTheAirWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSoundVolume(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeSoundVolumeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSoundVolumeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSoundVolumeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeSupportedOperatingModes(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeSupportedOperatingModesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedOperatingModesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSupportedOperatingModesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeUserCodeTemporaryDisableTime(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeUserCodeTemporaryDisableTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUserCodeTemporaryDisableTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUserCodeTemporaryDisableTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/readAttributeWrongCodeEntryLimit(completion:)
func (m_ MTRBaseClusterDoorLock) ReadAttributeWrongCodeEntryLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeWrongCodeEntryLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeWrongCodeEntryLimitWithCompletion */


// Command SetAliroReaderConfig
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setAliroReaderConfigWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetAliroReaderConfigWithParamsCompletion(params IMTRDoorLockClusterSetAliroReaderConfigParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAliroReaderConfigWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetAliroReaderConfigWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setCredentialWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetCredentialWithParamsCompletion(params IMTRDoorLockClusterSetCredentialParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetCredentialWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setHolidayScheduleWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetHolidayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetHolidayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetHolidayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setUserWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetUserWithParamsCompletion(params IMTRDoorLockClusterSetUserParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetUserWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setWeekDayScheduleWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetWeekDayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetWeekDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetWeekDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/setYearDayScheduleWith(_:completion:)
func (m_ MTRBaseClusterDoorLock) SetYearDayScheduleWithParamsCompletion(params IMTRDoorLockClusterSetYearDayScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetYearDayScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeActuatorEnabled(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeActuatorEnabledWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActuatorEnabledWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeActuatorEnabledWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroBLEAdvertisingVersion(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroBLEAdvertisingVersionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroBLEAdvertisingVersionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroBLEAdvertisingVersionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroExpeditedTransactionSupportedProtocolVersions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroExpeditedTransactionSupportedProtocolVersionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroGroupResolvingKey(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroGroupResolvingKeyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroGroupResolvingKeyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroGroupResolvingKeyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroReaderGroupIdentifier(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroReaderGroupIdentifierWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroReaderGroupIdentifierWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroReaderGroupIdentifierWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroReaderGroupSubIdentifier(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroReaderGroupSubIdentifierWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroReaderGroupSubIdentifierWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroReaderGroupSubIdentifierWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroReaderVerificationKey(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroReaderVerificationKeyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroReaderVerificationKeyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroReaderVerificationKeyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAliroSupportedBLEUWBProtocolVersions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAliroSupportedBLEUWBProtocolVersionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAliroSupportedBLEUWBProtocolVersionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAliroSupportedBLEUWBProtocolVersionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeAutoRelockTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeAutoRelockTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAutoRelockTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAutoRelockTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeCredentialRulesSupport(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeCredentialRulesSupportWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCredentialRulesSupportWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCredentialRulesSupportWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeDefaultConfigurationRegister(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeDefaultConfigurationRegisterWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDefaultConfigurationRegisterWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDefaultConfigurationRegisterWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeDoorClosedEvents(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeDoorClosedEventsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDoorClosedEventsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDoorClosedEventsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeDoorOpenEvents(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeDoorOpenEventsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDoorOpenEventsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDoorOpenEventsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeDoorState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeDoorStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDoorStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDoorStateWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeEnableInsideStatusLED(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeEnableInsideStatusLEDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnableInsideStatusLEDWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnableInsideStatusLEDWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeEnableLocalProgramming(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeEnableLocalProgrammingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnableLocalProgrammingWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnableLocalProgrammingWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeEnableOneTouchLocking(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeEnableOneTouchLockingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnableOneTouchLockingWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnableOneTouchLockingWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeEnablePrivacyModeButton(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeEnablePrivacyModeButtonWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnablePrivacyModeButtonWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnablePrivacyModeButtonWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeExpiringUserTimeout(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeExpiringUserTimeoutWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeExpiringUserTimeoutWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeExpiringUserTimeoutWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeLanguage(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeLanguageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLanguageWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLanguageWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeLEDSettings(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeLEDSettingsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLEDSettingsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLEDSettingsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeLocalProgrammingFeatures(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeLocalProgrammingFeaturesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLocalProgrammingFeaturesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLocalProgrammingFeaturesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeLockState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeLockStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLockStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLockStateWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeLockType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeLockTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLockTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLockTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeMaxPINCodeLength(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeMaxPINCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxPINCodeLengthWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxPINCodeLengthWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeMaxRFIDCodeLength(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeMaxRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxRFIDCodeLengthWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeMinPINCodeLength(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeMinPINCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinPINCodeLengthWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinPINCodeLengthWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeMinRFIDCodeLength(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeMinRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinRFIDCodeLengthWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinRFIDCodeLengthWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfAliroCredentialIssuerKeysSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfAliroCredentialIssuerKeysSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfAliroCredentialIssuerKeysSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfAliroCredentialIssuerKeysSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfAliroEndpointKeysSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfAliroEndpointKeysSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfAliroEndpointKeysSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfAliroEndpointKeysSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfCredentialsSupportedPerUser(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfCredentialsSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfCredentialsSupportedPerUserWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfCredentialsSupportedPerUserWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfHolidaySchedulesSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfHolidaySchedulesSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfHolidaySchedulesSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfHolidaySchedulesSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfPINUsersSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfPINUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfPINUsersSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfPINUsersSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfRFIDUsersSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfRFIDUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfRFIDUsersSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfRFIDUsersSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfTotalUsersSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfTotalUsersSupportedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfTotalUsersSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfTotalUsersSupportedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfWeekDaySchedulesSupportedPerUser(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfWeekDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfWeekDaySchedulesSupportedPerUserWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfWeekDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeNumberOfYearDaySchedulesSupportedPerUser(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeNumberOfYearDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfYearDaySchedulesSupportedPerUserWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfYearDaySchedulesSupportedPerUserWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeOpenPeriod(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeOpenPeriodWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOpenPeriodWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOpenPeriodWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeOperatingMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeOperatingModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperatingModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOperatingModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeRequirePINforRemoteOperation(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeRequirePINforRemoteOperationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRequirePINforRemoteOperationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeRequirePINforRemoteOperationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeSendPINOverTheAir(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeSendPINOverTheAirWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSendPINOverTheAirWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSendPINOverTheAirWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeSoundVolume(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeSoundVolumeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSoundVolumeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSoundVolumeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeSupportedOperatingModes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeSupportedOperatingModesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedOperatingModesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSupportedOperatingModesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeUserCodeTemporaryDisableTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeUserCodeTemporaryDisableTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUserCodeTemporaryDisableTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUserCodeTemporaryDisableTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/subscribeAttributeWrongCodeEntryLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDoorLock) SubscribeAttributeWrongCodeEntryLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeWrongCodeEntryLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeWrongCodeEntryLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/unboltDoor(completion:)
func (m_ MTRBaseClusterDoorLock) UnboltDoorWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unboltDoorWithCompletion:"), completion)
}/* debug [instance_methods/method]: UnboltDoorWithCompletion */


// Command UnboltDoor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/unboltDoor(with:completion:)
func (m_ MTRBaseClusterDoorLock) UnboltDoorWithParamsCompletion(params IMTRDoorLockClusterUnboltDoorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unboltDoorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: UnboltDoorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/unlockDoor(completion:)
func (m_ MTRBaseClusterDoorLock) UnlockDoorWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unlockDoorWithCompletion:"), completion)
}/* debug [instance_methods/method]: UnlockDoorWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/unlockDoor(with:completion:)
func (m_ MTRBaseClusterDoorLock) UnlockDoorWithParamsCompletion(params IMTRDoorLockClusterUnlockDoorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unlockDoorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: UnlockDoorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/unlockWithTimeout(with:completion:)
func (m_ MTRBaseClusterDoorLock) UnlockWithTimeoutWithParamsCompletion(params IMTRDoorLockClusterUnlockWithTimeoutParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unlockWithTimeoutWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: UnlockWithTimeoutWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeAutoRelockTime(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeAutoRelockTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeAutoRelockTimeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeAutoRelockTimeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeAutoRelockTime(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeAutoRelockTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeAutoRelockTimeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeAutoRelockTimeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeDoorClosedEvents(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeDoorClosedEventsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDoorClosedEventsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeDoorClosedEventsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeDoorClosedEvents(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeDoorClosedEventsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDoorClosedEventsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeDoorClosedEventsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeDoorOpenEvents(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeDoorOpenEventsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDoorOpenEventsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeDoorOpenEventsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeDoorOpenEvents(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeDoorOpenEventsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDoorOpenEventsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeDoorOpenEventsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableInsideStatusLED(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableInsideStatusLEDWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableInsideStatusLEDWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableInsideStatusLEDWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableInsideStatusLED(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableInsideStatusLEDWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableInsideStatusLEDWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableInsideStatusLEDWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableLocalProgramming(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableLocalProgrammingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableLocalProgrammingWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableLocalProgrammingWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableLocalProgramming(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableLocalProgrammingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableLocalProgrammingWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableLocalProgrammingWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableOneTouchLocking(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableOneTouchLockingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableOneTouchLockingWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableOneTouchLockingWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnableOneTouchLocking(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnableOneTouchLockingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnableOneTouchLockingWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeEnableOneTouchLockingWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnablePrivacyModeButton(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnablePrivacyModeButtonWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnablePrivacyModeButtonWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeEnablePrivacyModeButtonWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeEnablePrivacyModeButton(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeEnablePrivacyModeButtonWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEnablePrivacyModeButtonWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeEnablePrivacyModeButtonWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeExpiringUserTimeout(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeExpiringUserTimeoutWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeExpiringUserTimeoutWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeExpiringUserTimeoutWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeExpiringUserTimeout(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeExpiringUserTimeoutWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeExpiringUserTimeoutWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeExpiringUserTimeoutWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLanguage(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLanguageWithValueCompletion(value objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLanguageWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeLanguageWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLanguage(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLanguageWithValueParamsCompletion(value objc.IObject /* cross-framework: NSString */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLanguageWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeLanguageWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLEDSettings(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLEDSettingsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLEDSettingsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeLEDSettingsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLEDSettings(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLEDSettingsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLEDSettingsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeLEDSettingsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLocalProgrammingFeatures(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLocalProgrammingFeaturesWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLocalProgrammingFeaturesWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeLocalProgrammingFeaturesWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeLocalProgrammingFeatures(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeLocalProgrammingFeaturesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLocalProgrammingFeaturesWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeLocalProgrammingFeaturesWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeOpenPeriod(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeOpenPeriodWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOpenPeriodWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOpenPeriodWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeOpenPeriod(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeOpenPeriodWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOpenPeriodWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOpenPeriodWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeOperatingMode(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeOperatingModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOperatingModeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOperatingModeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeOperatingMode(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeOperatingModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOperatingModeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOperatingModeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeRequirePINforRemoteOperation(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeRequirePINforRemoteOperationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRequirePINforRemoteOperationWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeRequirePINforRemoteOperationWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeRequirePINforRemoteOperation(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeRequirePINforRemoteOperationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRequirePINforRemoteOperationWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeRequirePINforRemoteOperationWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeSendPINOverTheAir(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeSendPINOverTheAirWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSendPINOverTheAirWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeSendPINOverTheAirWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeSendPINOverTheAir(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeSendPINOverTheAirWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSendPINOverTheAirWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeSendPINOverTheAirWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeSoundVolume(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeSoundVolumeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSoundVolumeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeSoundVolumeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeSoundVolume(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeSoundVolumeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSoundVolumeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeSoundVolumeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeUserCodeTemporaryDisableTime(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeUserCodeTemporaryDisableTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserCodeTemporaryDisableTimeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeUserCodeTemporaryDisableTimeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeUserCodeTemporaryDisableTime(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeUserCodeTemporaryDisableTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserCodeTemporaryDisableTimeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeUserCodeTemporaryDisableTimeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeWrongCodeEntryLimit(withValue:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeWrongCodeEntryLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeWrongCodeEntryLimitWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeWrongCodeEntryLimitWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock/writeAttributeWrongCodeEntryLimit(withValue:params:completion:)
func (m_ MTRBaseClusterDoorLock) WriteAttributeWrongCodeEntryLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeWrongCodeEntryLimitWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeWrongCodeEntryLimitWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterDoorLock */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterDoorLock */


