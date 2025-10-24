// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterThermostat */


/* debug [class_header]: Header for MTRBaseClusterThermostat */
// The class instance for the [MTRBaseClusterThermostat] class.
var (
	MTRBaseClusterThermostatClass     _MTRBaseClusterThermostatClass
	MTRBaseClusterThermostatClassOnce sync.Once
)

func getMTRBaseClusterThermostatClass() _MTRBaseClusterThermostatClass {
	MTRBaseClusterThermostatClassOnce.Do(func() {
		MTRBaseClusterThermostatClass = _MTRBaseClusterThermostatClass{objc.GetClass("MTRBaseClusterThermostat")}
	})
	return MTRBaseClusterThermostatClass
}

type _MTRBaseClusterThermostatClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterThermostat */
// An interface definition for the [MTRBaseClusterThermostat] class.
type IMTRBaseClusterThermostat interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterThermostat */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterThermostat */
	// methods:
	AtomicRequestWithParamsCompletion(params IMTRThermostatClusterAtomicRequestParams, completion unsafe.Pointer)
	ClearWeeklyScheduleWithCompletion(completion unsafe.Pointer)
	ClearWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterClearWeeklyScheduleParams, completion unsafe.Pointer)
	GetWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterGetWeeklyScheduleParams, completion unsafe.Pointer)
	ReadAttributeAbsMaxCoolSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeAbsMaxHeatSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeAbsMinCoolSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeAbsMinHeatSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeACCapacityWithCompletion(completion unsafe.Pointer)
	ReadAttributeACCapacityformatWithCompletion(completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeACCoilTemperatureWithCompletion(completion unsafe.Pointer)
	ReadAttributeACCompressorTypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeACErrorCodeWithCompletion(completion unsafe.Pointer)
	ReadAttributeACLouverPositionWithCompletion(completion unsafe.Pointer)
	ReadAttributeACRefrigerantTypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeActivePresetHandleWithCompletion(completion unsafe.Pointer)
	ReadAttributeActiveScheduleHandleWithCompletion(completion unsafe.Pointer)
	ReadAttributeACTypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeControlSequenceOfOperationWithCompletion(completion unsafe.Pointer)
	ReadAttributeEmergencyHeatDeltaWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeHVACSystemTypeConfigurationWithCompletion(completion unsafe.Pointer)
	ReadAttributeLocalTemperatureWithCompletion(completion unsafe.Pointer)
	ReadAttributeLocalTemperatureCalibrationWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxCoolSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxHeatSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinCoolSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinHeatSetpointLimitWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinSetpointDeadBandWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfDailyTransitionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfPresetsWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfSchedulesWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfScheduleTransitionPerDayWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfScheduleTransitionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfWeeklyTransitionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupancyWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupiedCoolingSetpointWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupiedHeatingSetpointWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupiedSetbackWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupiedSetbackMaxWithCompletion(completion unsafe.Pointer)
	ReadAttributeOccupiedSetbackMinWithCompletion(completion unsafe.Pointer)
	ReadAttributeOutdoorTemperatureWithCompletion(completion unsafe.Pointer)
	ReadAttributePICoolingDemandWithCompletion(completion unsafe.Pointer)
	ReadAttributePIHeatingDemandWithCompletion(completion unsafe.Pointer)
	ReadAttributePresetsWithCompletion(completion unsafe.Pointer)
	ReadAttributePresetTypesWithCompletion(completion unsafe.Pointer)
	ReadAttributeRemoteSensingWithCompletion(completion unsafe.Pointer)
	ReadAttributeSchedulesWithCompletion(completion unsafe.Pointer)
	ReadAttributeScheduleTypesWithCompletion(completion unsafe.Pointer)
	ReadAttributeSetpointChangeAmountWithCompletion(completion unsafe.Pointer)
	ReadAttributeSetpointChangeSourceWithCompletion(completion unsafe.Pointer)
	ReadAttributeSetpointChangeSourceTimestampWithCompletion(completion unsafe.Pointer)
	ReadAttributeSetpointHoldExpiryTimestampWithCompletion(completion unsafe.Pointer)
	ReadAttributeStartOfWeekWithCompletion(completion unsafe.Pointer)
	ReadAttributeSystemModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeTemperatureSetpointHoldWithCompletion(completion unsafe.Pointer)
	ReadAttributeTemperatureSetpointHoldDurationWithCompletion(completion unsafe.Pointer)
	ReadAttributeThermostatProgrammingOperationModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeThermostatRunningModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeThermostatRunningStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeUnoccupiedCoolingSetpointWithCompletion(completion unsafe.Pointer)
	ReadAttributeUnoccupiedHeatingSetpointWithCompletion(completion unsafe.Pointer)
	ReadAttributeUnoccupiedSetbackWithCompletion(completion unsafe.Pointer)
	ReadAttributeUnoccupiedSetbackMaxWithCompletion(completion unsafe.Pointer)
	ReadAttributeUnoccupiedSetbackMinWithCompletion(completion unsafe.Pointer)
	SetActivePresetRequestWithParamsCompletion(params IMTRThermostatClusterSetActivePresetRequestParams, completion unsafe.Pointer)
	SetActiveScheduleRequestWithParamsCompletion(params IMTRThermostatClusterSetActiveScheduleRequestParams, completion unsafe.Pointer)
	SetpointRaiseLowerWithParamsCompletion(params IMTRThermostatClusterSetpointRaiseLowerParams, completion unsafe.Pointer)
	SetWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterSetWeeklyScheduleParams, completion unsafe.Pointer)
	SubscribeAttributeAbsMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAbsMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAbsMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAbsMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACCapacityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACCapacityformatWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACCoilTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACCompressorTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACErrorCodeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACLouverPositionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACRefrigerantTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActivePresetHandleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActiveScheduleHandleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeACTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeControlSequenceOfOperationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEmergencyHeatDeltaWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeHVACSystemTypeConfigurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLocalTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLocalTemperatureCalibrationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinSetpointDeadBandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfDailyTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfPresetsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfSchedulesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfScheduleTransitionPerDayWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfScheduleTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfWeeklyTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupancyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupiedSetbackWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOutdoorTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePICoolingDemandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePIHeatingDemandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePresetsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePresetTypesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRemoteSensingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSchedulesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeScheduleTypesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSetpointChangeAmountWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSetpointChangeSourceWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSetpointChangeSourceTimestampWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSetpointHoldExpiryTimestampWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStartOfWeekWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSystemModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTemperatureSetpointHoldWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTemperatureSetpointHoldDurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThermostatProgrammingOperationModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThermostatRunningModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThermostatRunningStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUnoccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUnoccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUnoccupiedSetbackWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUnoccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUnoccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeACCapacityWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACCapacityWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACCapacityformatWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACCapacityformatWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACCompressorTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACCompressorTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACErrorCodeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACErrorCodeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACLouverPositionWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACLouverPositionWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACRefrigerantTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACRefrigerantTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeACTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeACTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeControlSequenceOfOperationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeControlSequenceOfOperationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeEmergencyHeatDeltaWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeEmergencyHeatDeltaWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeHVACSystemTypeConfigurationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeHVACSystemTypeConfigurationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeLocalTemperatureCalibrationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeLocalTemperatureCalibrationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeMaxCoolSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeMaxCoolSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeMaxHeatSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeMaxHeatSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeMinCoolSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeMinCoolSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeMinHeatSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeMinHeatSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeMinSetpointDeadBandWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeMinSetpointDeadBandWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOccupiedCoolingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOccupiedCoolingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOccupiedHeatingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOccupiedHeatingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOccupiedSetbackWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOccupiedSetbackWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributePresetsWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer)
	WriteAttributePresetsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeRemoteSensingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeRemoteSensingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeSchedulesWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer)
	WriteAttributeSchedulesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeSystemModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSystemModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeTemperatureSetpointHoldWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeTemperatureSetpointHoldWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeTemperatureSetpointHoldDurationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeTemperatureSetpointHoldDurationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeThermostatProgrammingOperationModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeThermostatProgrammingOperationModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeUnoccupiedCoolingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeUnoccupiedCoolingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeUnoccupiedHeatingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeUnoccupiedHeatingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeUnoccupiedSetbackWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeUnoccupiedSetbackWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterThermostat */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThermostatClass) Alloc() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterThermostatClass) New() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThermostat) Init() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThermostat) Autorelease() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThermostat creates a new MTRBaseClusterThermostat instance.
func NewMTRBaseClusterThermostat() MTRBaseClusterThermostat {
	return getMTRBaseClusterThermostatClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterThermostat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat
type MTRBaseClusterThermostat struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThermostatFrom constructs a [MTRBaseClusterThermostat] from an unsafe.Pointer.
func MTRBaseClusterThermostatFrom(ptr unsafe.Pointer) MTRBaseClusterThermostat {
	return MTRBaseClusterThermostat{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterThermostat */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/init(device:endpointID:queue:)
func NewMTRBaseClusterThermostatWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterThermostat {
	instance := getMTRBaseClusterThermostatClass().Alloc()
	rv := objc.Send[MTRBaseClusterThermostat](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterThermostatWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/init(device:endpoint:queue:)
func NewMTRBaseClusterThermostatWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterThermostat {
	instance := getMTRBaseClusterThermostatClass().Alloc()
	rv := objc.Send[MTRBaseClusterThermostat](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterThermostatWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterThermostat */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxCoolSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMaxCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMaxCoolSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMaxCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxCoolSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMaxCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMaxCoolSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMaxCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxHeatSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMaxHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMaxHeatSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMaxHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxHeatSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMaxHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMaxHeatSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMaxHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinCoolSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMinCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMinCoolSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMinCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinCoolSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMinCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMinCoolSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMinCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinHeatSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMinHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMinHeatSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMinHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinHeatSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAbsMinHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMinHeatSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAbsMinHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCapacityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCapacityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCapacityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCapacityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCapacityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCapacityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacityformat(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCapacityformatWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCapacityformatWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCapacityformatWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacityformat(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCapacityformatWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCapacityformatWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCapacityformatWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCoilTemperature(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCoilTemperatureWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCoilTemperatureWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCoilTemperatureWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCoilTemperature(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCoilTemperatureWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCoilTemperatureWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCoilTemperatureWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCompressorType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCompressorTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCompressorTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCompressorTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCompressorType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACCompressorTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACCompressorTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACCompressorTypeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACErrorCode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACErrorCodeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACErrorCodeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACErrorCodeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACErrorCode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACErrorCodeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACErrorCodeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACErrorCodeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACLouverPosition(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACLouverPositionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACLouverPositionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACLouverPositionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACLouverPosition(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACLouverPositionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACLouverPositionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACLouverPositionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACRefrigerantType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACRefrigerantTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACRefrigerantTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACRefrigerantTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACRefrigerantType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACRefrigerantTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACRefrigerantTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACRefrigerantTypeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeActivePresetHandle(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeActivePresetHandleWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActivePresetHandleWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeActivePresetHandleWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeActiveScheduleHandle(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeActiveScheduleHandleWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActiveScheduleHandleWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeActiveScheduleHandleWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeACTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeACTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeACTypeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeControlSequenceOfOperation(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeControlSequenceOfOperationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeControlSequenceOfOperationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeControlSequenceOfOperationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeControlSequenceOfOperation(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeControlSequenceOfOperationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeControlSequenceOfOperationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeControlSequenceOfOperationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeEmergencyHeatDelta(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeEmergencyHeatDeltaWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEmergencyHeatDeltaWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEmergencyHeatDeltaWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeEmergencyHeatDelta(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeEmergencyHeatDeltaWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEmergencyHeatDeltaWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEmergencyHeatDeltaWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeHVACSystemTypeConfiguration(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeHVACSystemTypeConfigurationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeHVACSystemTypeConfigurationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeHVACSystemTypeConfigurationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeHVACSystemTypeConfiguration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeHVACSystemTypeConfigurationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeHVACSystemTypeConfigurationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeHVACSystemTypeConfigurationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperature(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeLocalTemperatureWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalTemperatureWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalTemperatureWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperature(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeLocalTemperatureWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalTemperatureWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalTemperatureWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperatureCalibration(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeLocalTemperatureCalibrationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalTemperatureCalibrationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalTemperatureCalibrationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperatureCalibration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeLocalTemperatureCalibrationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalTemperatureCalibrationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeLocalTemperatureCalibrationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxCoolSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMaxCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxCoolSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxCoolSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMaxCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxCoolSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxHeatSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMaxHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxHeatSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxHeatSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMaxHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxHeatSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinCoolSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinCoolSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinCoolSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinCoolSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinCoolSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinCoolSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinHeatSetpointLimit(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinHeatSetpointLimitWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinHeatSetpointLimitWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinHeatSetpointLimit(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinHeatSetpointLimitWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinHeatSetpointLimitWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinSetpointDeadBand(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinSetpointDeadBandWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinSetpointDeadBandWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinSetpointDeadBandWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinSetpointDeadBand(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeMinSetpointDeadBandWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinSetpointDeadBandWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinSetpointDeadBandWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfDailyTransitions(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfDailyTransitionsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfDailyTransitionsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfDailyTransitionsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfDailyTransitions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfDailyTransitionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfDailyTransitionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfDailyTransitionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfPresets(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfPresetsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfPresetsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfPresetsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfSchedules(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfSchedulesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfSchedulesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfSchedulesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfScheduleTransitionPerDay(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfScheduleTransitionPerDayWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfScheduleTransitionPerDayWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfScheduleTransitionPerDayWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfScheduleTransitions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfScheduleTransitionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfScheduleTransitionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfScheduleTransitionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfWeeklyTransitions(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfWeeklyTransitionsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfWeeklyTransitionsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfWeeklyTransitionsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfWeeklyTransitions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeNumberOfWeeklyTransitionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfWeeklyTransitionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfWeeklyTransitionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupancy(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupancyWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupancyWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupancyWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupancy(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupancyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupancyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupancyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedCoolingSetpoint(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedCoolingSetpointWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedCoolingSetpointWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedCoolingSetpointWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedCoolingSetpoint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedCoolingSetpointWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedCoolingSetpointWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedCoolingSetpointWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedHeatingSetpoint(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedHeatingSetpointWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedHeatingSetpointWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedHeatingSetpointWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedHeatingSetpoint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedHeatingSetpointWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedHeatingSetpointWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedHeatingSetpointWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetback(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetback(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMax(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackMaxWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackMaxWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackMaxWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMax(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackMaxWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackMaxWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackMaxWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMin(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackMinWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackMinWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackMinWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMin(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOccupiedSetbackMinWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOccupiedSetbackMinWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOccupiedSetbackMinWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOutdoorTemperature(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOutdoorTemperatureWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOutdoorTemperatureWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOutdoorTemperatureWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOutdoorTemperature(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeOutdoorTemperatureWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOutdoorTemperatureWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOutdoorTemperatureWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePICoolingDemand(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePICoolingDemandWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePICoolingDemandWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePICoolingDemandWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePICoolingDemand(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePICoolingDemandWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePICoolingDemandWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePICoolingDemandWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePIHeatingDemand(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePIHeatingDemandWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePIHeatingDemandWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePIHeatingDemandWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePIHeatingDemand(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePIHeatingDemandWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePIHeatingDemandWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePIHeatingDemandWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePresets(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePresetsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePresetsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePresetsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePresetTypes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributePresetTypesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePresetTypesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePresetTypesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeRemoteSensing(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeRemoteSensingWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemoteSensingWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemoteSensingWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeRemoteSensing(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeRemoteSensingWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemoteSensingWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemoteSensingWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSchedules(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSchedulesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSchedulesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSchedulesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeScheduleTypes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeScheduleTypesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeScheduleTypesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeScheduleTypesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeAmount(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeAmountWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeAmountWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeAmountWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeAmount(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeAmountWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeAmountWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeAmountWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSource(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeSourceWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeSourceWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeSourceWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSource(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeSourceWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeSourceWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeSourceWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSourceTimestamp(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeSourceTimestampWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeSourceTimestampWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeSourceTimestampWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSourceTimestamp(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointChangeSourceTimestampWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointChangeSourceTimestampWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointChangeSourceTimestampWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointHoldExpiryTimestamp(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSetpointHoldExpiryTimestampWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSetpointHoldExpiryTimestampWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSetpointHoldExpiryTimestampWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeStartOfWeek(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeStartOfWeekWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartOfWeekWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartOfWeekWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeStartOfWeek(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeStartOfWeekWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartOfWeekWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartOfWeekWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSystemMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSystemModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSystemModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSystemModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSystemMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeSystemModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSystemModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSystemModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHold(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeTemperatureSetpointHoldWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTemperatureSetpointHoldWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTemperatureSetpointHoldWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHold(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeTemperatureSetpointHoldWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTemperatureSetpointHoldWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTemperatureSetpointHoldWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHoldDuration(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeTemperatureSetpointHoldDurationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTemperatureSetpointHoldDurationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTemperatureSetpointHoldDurationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHoldDuration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeTemperatureSetpointHoldDurationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTemperatureSetpointHoldDurationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTemperatureSetpointHoldDurationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatProgrammingOperationMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatProgrammingOperationModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatProgrammingOperationModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatProgrammingOperationModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatProgrammingOperationMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatProgrammingOperationModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatProgrammingOperationModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatProgrammingOperationModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatRunningModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatRunningModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatRunningModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatRunningModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatRunningModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatRunningModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningState(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatRunningStateWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatRunningStateWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatRunningStateWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeThermostatRunningStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThermostatRunningStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThermostatRunningStateWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedCoolingSetpoint(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedCoolingSetpointWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedCoolingSetpointWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedCoolingSetpointWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedCoolingSetpoint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedCoolingSetpointWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedCoolingSetpointWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedCoolingSetpointWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedHeatingSetpoint(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedHeatingSetpointWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedHeatingSetpointWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedHeatingSetpointWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedHeatingSetpoint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedHeatingSetpointWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedHeatingSetpointWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedHeatingSetpointWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetback(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetback(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMax(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackMaxWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackMaxWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackMaxWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMax(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackMaxWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackMaxWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackMaxWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMin(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackMinWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackMinWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackMinWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMin(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThermostatClass) ReadAttributeUnoccupiedSetbackMinWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUnoccupiedSetbackMinWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUnoccupiedSetbackMinWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterThermostat */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterThermostat */

// Command AtomicRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/atomicRequest(with:completion:)
func (m_ MTRBaseClusterThermostat) AtomicRequestWithParamsCompletion(params IMTRThermostatClusterAtomicRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("atomicRequestWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: AtomicRequestWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/clearWeeklySchedule(completion:)
func (m_ MTRBaseClusterThermostat) ClearWeeklyScheduleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearWeeklyScheduleWithCompletion:"), completion)
}/* debug [instance_methods/method]: ClearWeeklyScheduleWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/clearWeeklySchedule(with:completion:)
func (m_ MTRBaseClusterThermostat) ClearWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterClearWeeklyScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearWeeklyScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ClearWeeklyScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/getWeeklySchedule(with:completion:)
func (m_ MTRBaseClusterThermostat) GetWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterGetWeeklyScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getWeeklyScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GetWeeklyScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxCoolSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAbsMaxCoolSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMaxCoolSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAbsMaxCoolSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMaxHeatSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAbsMaxHeatSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMaxHeatSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAbsMaxHeatSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinCoolSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAbsMinCoolSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMinCoolSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAbsMinCoolSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAbsMinHeatSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAbsMinHeatSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMinHeatSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAbsMinHeatSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacity(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACCapacityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACCapacityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACCapacityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCapacityformat(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACCapacityformatWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACCapacityformatWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACCapacityformatWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCoilTemperature(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACCoilTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACCoilTemperatureWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACCoilTemperatureWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACCompressorType(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACCompressorTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACCompressorTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACCompressorTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACErrorCode(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACErrorCodeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACErrorCodeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACErrorCodeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACLouverPosition(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACLouverPositionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACLouverPositionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACLouverPositionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACRefrigerantType(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACRefrigerantTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACRefrigerantTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACRefrigerantTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeActivePresetHandle(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeActivePresetHandleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActivePresetHandleWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeActivePresetHandleWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeActiveScheduleHandle(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeActiveScheduleHandleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActiveScheduleHandleWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeActiveScheduleHandleWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeACType(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeACTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeACTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeACTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeControlSequenceOfOperation(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeControlSequenceOfOperationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeControlSequenceOfOperationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeControlSequenceOfOperationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeEmergencyHeatDelta(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeEmergencyHeatDeltaWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEmergencyHeatDeltaWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEmergencyHeatDeltaWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeHVACSystemTypeConfiguration(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeHVACSystemTypeConfigurationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeHVACSystemTypeConfigurationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeHVACSystemTypeConfigurationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperature(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeLocalTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLocalTemperatureWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLocalTemperatureWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeLocalTemperatureCalibration(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeLocalTemperatureCalibrationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLocalTemperatureCalibrationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeLocalTemperatureCalibrationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxCoolSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeMaxCoolSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxCoolSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxCoolSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMaxHeatSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeMaxHeatSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxHeatSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxHeatSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinCoolSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeMinCoolSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinCoolSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinCoolSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinHeatSetpointLimit(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeMinHeatSetpointLimitWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinHeatSetpointLimitWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinHeatSetpointLimitWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeMinSetpointDeadBand(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeMinSetpointDeadBandWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinSetpointDeadBandWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinSetpointDeadBandWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfDailyTransitions(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfDailyTransitionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfDailyTransitionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfDailyTransitionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfPresets(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfPresetsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfPresetsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfPresetsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfSchedules(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfSchedulesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfSchedulesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfSchedulesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfScheduleTransitionPerDay(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfScheduleTransitionPerDayWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfScheduleTransitionPerDayWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfScheduleTransitionPerDayWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfScheduleTransitions(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfScheduleTransitionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfScheduleTransitionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfScheduleTransitionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeNumberOfWeeklyTransitions(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeNumberOfWeeklyTransitionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfWeeklyTransitionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfWeeklyTransitionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupancy(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupancyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupancyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupancyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedCoolingSetpoint(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupiedCoolingSetpointWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupiedCoolingSetpointWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupiedCoolingSetpointWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedHeatingSetpoint(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupiedHeatingSetpointWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupiedHeatingSetpointWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupiedHeatingSetpointWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetback(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupiedSetbackWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupiedSetbackWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupiedSetbackWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMax(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupiedSetbackMaxWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupiedSetbackMaxWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupiedSetbackMaxWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOccupiedSetbackMin(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOccupiedSetbackMinWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOccupiedSetbackMinWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOccupiedSetbackMinWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeOutdoorTemperature(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeOutdoorTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOutdoorTemperatureWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOutdoorTemperatureWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePICoolingDemand(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributePICoolingDemandWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePICoolingDemandWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePICoolingDemandWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePIHeatingDemand(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributePIHeatingDemandWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePIHeatingDemandWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePIHeatingDemandWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePresets(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributePresetsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePresetsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePresetsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributePresetTypes(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributePresetTypesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePresetTypesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePresetTypesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeRemoteSensing(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeRemoteSensingWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRemoteSensingWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeRemoteSensingWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSchedules(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSchedulesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSchedulesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSchedulesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeScheduleTypes(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeScheduleTypesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeScheduleTypesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeScheduleTypesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeAmount(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSetpointChangeAmountWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSetpointChangeAmountWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSetpointChangeAmountWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSource(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSetpointChangeSourceWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSetpointChangeSourceWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSetpointChangeSourceWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointChangeSourceTimestamp(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSetpointChangeSourceTimestampWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSetpointChangeSourceTimestampWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSetpointChangeSourceTimestampWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSetpointHoldExpiryTimestamp(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSetpointHoldExpiryTimestampWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSetpointHoldExpiryTimestampWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSetpointHoldExpiryTimestampWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeStartOfWeek(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeStartOfWeekWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStartOfWeekWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeStartOfWeekWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeSystemMode(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeSystemModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSystemModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSystemModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHold(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeTemperatureSetpointHoldWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTemperatureSetpointHoldWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeTemperatureSetpointHoldWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeTemperatureSetpointHoldDuration(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeTemperatureSetpointHoldDurationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTemperatureSetpointHoldDurationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeTemperatureSetpointHoldDurationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatProgrammingOperationMode(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeThermostatProgrammingOperationModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThermostatProgrammingOperationModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeThermostatProgrammingOperationModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningMode(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeThermostatRunningModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThermostatRunningModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeThermostatRunningModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeThermostatRunningState(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeThermostatRunningStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThermostatRunningStateWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeThermostatRunningStateWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedCoolingSetpoint(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeUnoccupiedCoolingSetpointWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUnoccupiedCoolingSetpointWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUnoccupiedCoolingSetpointWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedHeatingSetpoint(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeUnoccupiedHeatingSetpointWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUnoccupiedHeatingSetpointWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUnoccupiedHeatingSetpointWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetback(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeUnoccupiedSetbackWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUnoccupiedSetbackWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUnoccupiedSetbackWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMax(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeUnoccupiedSetbackMaxWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUnoccupiedSetbackMaxWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUnoccupiedSetbackMaxWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/readAttributeUnoccupiedSetbackMin(completion:)
func (m_ MTRBaseClusterThermostat) ReadAttributeUnoccupiedSetbackMinWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUnoccupiedSetbackMinWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUnoccupiedSetbackMinWithCompletion */


// Command SetActivePresetRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/setActivePresetRequestWith(_:completion:)
func (m_ MTRBaseClusterThermostat) SetActivePresetRequestWithParamsCompletion(params IMTRThermostatClusterSetActivePresetRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActivePresetRequestWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetActivePresetRequestWithParamsCompletion */


// Command SetActiveScheduleRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/setActiveScheduleRequestWith(_:completion:)
func (m_ MTRBaseClusterThermostat) SetActiveScheduleRequestWithParamsCompletion(params IMTRThermostatClusterSetActiveScheduleRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveScheduleRequestWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetActiveScheduleRequestWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/setpointRaiseLower(with:completion:)
func (m_ MTRBaseClusterThermostat) SetpointRaiseLowerWithParamsCompletion(params IMTRThermostatClusterSetpointRaiseLowerParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setpointRaiseLowerWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetpointRaiseLowerWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/setWeeklyScheduleWith(_:completion:)
func (m_ MTRBaseClusterThermostat) SetWeeklyScheduleWithParamsCompletion(params IMTRThermostatClusterSetWeeklyScheduleParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeeklyScheduleWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: SetWeeklyScheduleWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAbsMaxCoolSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAbsMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMaxCoolSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAbsMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAbsMaxHeatSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAbsMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMaxHeatSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAbsMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAbsMinCoolSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAbsMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMinCoolSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAbsMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAbsMinHeatSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAbsMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMinHeatSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAbsMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACCapacity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACCapacityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACCapacityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACCapacityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACCapacityformat(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACCapacityformatWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACCapacityformatWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACCapacityformatWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACCoilTemperature(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACCoilTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACCoilTemperatureWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACCoilTemperatureWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACCompressorType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACCompressorTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACCompressorTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACCompressorTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACErrorCode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACErrorCodeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACErrorCodeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACErrorCodeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACLouverPosition(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACLouverPositionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACLouverPositionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACLouverPositionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACRefrigerantType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACRefrigerantTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACRefrigerantTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACRefrigerantTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeActivePresetHandle(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeActivePresetHandleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActivePresetHandleWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeActivePresetHandleWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeActiveScheduleHandle(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeActiveScheduleHandleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActiveScheduleHandleWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeActiveScheduleHandleWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeACType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeACTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeACTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeACTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeControlSequenceOfOperation(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeControlSequenceOfOperationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeControlSequenceOfOperationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeControlSequenceOfOperationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeEmergencyHeatDelta(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeEmergencyHeatDeltaWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEmergencyHeatDeltaWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEmergencyHeatDeltaWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeHVACSystemTypeConfiguration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeHVACSystemTypeConfigurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeHVACSystemTypeConfigurationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeHVACSystemTypeConfigurationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeLocalTemperature(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeLocalTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLocalTemperatureWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLocalTemperatureWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeLocalTemperatureCalibration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeLocalTemperatureCalibrationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLocalTemperatureCalibrationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeLocalTemperatureCalibrationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeMaxCoolSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxCoolSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeMaxHeatSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxHeatSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeMinCoolSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinCoolSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinCoolSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeMinHeatSetpointLimit(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinHeatSetpointLimitWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinHeatSetpointLimitWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeMinSetpointDeadBand(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeMinSetpointDeadBandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinSetpointDeadBandWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinSetpointDeadBandWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfDailyTransitions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfDailyTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfDailyTransitionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfDailyTransitionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfPresets(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfPresetsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfPresetsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfPresetsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfSchedules(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfSchedulesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfSchedulesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfSchedulesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfScheduleTransitionPerDay(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfScheduleTransitionPerDayWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfScheduleTransitionPerDayWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfScheduleTransitionPerDayWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfScheduleTransitions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfScheduleTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfScheduleTransitionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfScheduleTransitionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeNumberOfWeeklyTransitions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeNumberOfWeeklyTransitionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfWeeklyTransitionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfWeeklyTransitionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupancy(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupancyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupancyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupancyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupiedCoolingSetpoint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupiedCoolingSetpointWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupiedHeatingSetpoint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupiedHeatingSetpointWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupiedSetback(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupiedSetbackWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupiedSetbackWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupiedSetbackWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupiedSetbackMax(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupiedSetbackMaxWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOccupiedSetbackMin(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOccupiedSetbackMinWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeOutdoorTemperature(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeOutdoorTemperatureWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOutdoorTemperatureWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOutdoorTemperatureWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributePICoolingDemand(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributePICoolingDemandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePICoolingDemandWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePICoolingDemandWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributePIHeatingDemand(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributePIHeatingDemandWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePIHeatingDemandWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePIHeatingDemandWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributePresets(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributePresetsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePresetsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePresetsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributePresetTypes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributePresetTypesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePresetTypesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePresetTypesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeRemoteSensing(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeRemoteSensingWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRemoteSensingWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeRemoteSensingWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSchedules(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSchedulesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSchedulesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSchedulesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeScheduleTypes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeScheduleTypesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeScheduleTypesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeScheduleTypesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSetpointChangeAmount(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSetpointChangeAmountWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSetpointChangeAmountWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSetpointChangeAmountWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSetpointChangeSource(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSetpointChangeSourceWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSetpointChangeSourceWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSetpointChangeSourceWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSetpointChangeSourceTimestamp(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSetpointChangeSourceTimestampWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSetpointChangeSourceTimestampWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSetpointChangeSourceTimestampWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSetpointHoldExpiryTimestamp(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSetpointHoldExpiryTimestampWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSetpointHoldExpiryTimestampWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSetpointHoldExpiryTimestampWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeStartOfWeek(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeStartOfWeekWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStartOfWeekWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeStartOfWeekWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeSystemMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeSystemModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSystemModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSystemModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeTemperatureSetpointHold(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeTemperatureSetpointHoldWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTemperatureSetpointHoldWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeTemperatureSetpointHoldWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeTemperatureSetpointHoldDuration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeTemperatureSetpointHoldDurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTemperatureSetpointHoldDurationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeTemperatureSetpointHoldDurationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeThermostatProgrammingOperationMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeThermostatProgrammingOperationModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThermostatProgrammingOperationModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeThermostatProgrammingOperationModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeThermostatRunningMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeThermostatRunningModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThermostatRunningModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeThermostatRunningModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeThermostatRunningState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeThermostatRunningStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThermostatRunningStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeThermostatRunningStateWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeUnoccupiedCoolingSetpoint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeUnoccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUnoccupiedCoolingSetpointWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUnoccupiedCoolingSetpointWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeUnoccupiedHeatingSetpoint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeUnoccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUnoccupiedHeatingSetpointWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUnoccupiedHeatingSetpointWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeUnoccupiedSetback(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeUnoccupiedSetbackWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUnoccupiedSetbackWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUnoccupiedSetbackWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeUnoccupiedSetbackMax(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeUnoccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUnoccupiedSetbackMaxWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUnoccupiedSetbackMaxWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/subscribeAttributeUnoccupiedSetbackMin(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThermostat) SubscribeAttributeUnoccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUnoccupiedSetbackMinWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUnoccupiedSetbackMinWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCapacity(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCapacityWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCapacityWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACCapacityWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCapacity(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCapacityWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCapacityWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACCapacityWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCapacityformat(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCapacityformatWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCapacityformatWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACCapacityformatWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCapacityformat(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCapacityformatWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCapacityformatWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACCapacityformatWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCompressorType(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCompressorTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCompressorTypeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACCompressorTypeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACCompressorType(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACCompressorTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACCompressorTypeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACCompressorTypeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACErrorCode(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACErrorCodeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACErrorCodeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACErrorCodeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACErrorCode(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACErrorCodeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACErrorCodeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACErrorCodeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACLouverPosition(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACLouverPositionWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACLouverPositionWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACLouverPositionWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACLouverPosition(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACLouverPositionWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACLouverPositionWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACLouverPositionWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACRefrigerantType(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACRefrigerantTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACRefrigerantTypeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACRefrigerantTypeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACRefrigerantType(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACRefrigerantTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACRefrigerantTypeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACRefrigerantTypeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACType(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACTypeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACTypeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeACTypeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeACType(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeACTypeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeACTypeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeACTypeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeControlSequenceOfOperation(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeControlSequenceOfOperationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeControlSequenceOfOperationWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeControlSequenceOfOperationWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeControlSequenceOfOperation(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeControlSequenceOfOperationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeControlSequenceOfOperationWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeControlSequenceOfOperationWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeEmergencyHeatDelta(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeEmergencyHeatDeltaWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEmergencyHeatDeltaWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeEmergencyHeatDeltaWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeEmergencyHeatDelta(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeEmergencyHeatDeltaWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeEmergencyHeatDeltaWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeEmergencyHeatDeltaWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeHVACSystemTypeConfiguration(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeHVACSystemTypeConfigurationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeHVACSystemTypeConfigurationWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeHVACSystemTypeConfigurationWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeHVACSystemTypeConfiguration(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeHVACSystemTypeConfigurationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeHVACSystemTypeConfigurationWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeHVACSystemTypeConfigurationWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeLocalTemperatureCalibration(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeLocalTemperatureCalibrationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLocalTemperatureCalibrationWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeLocalTemperatureCalibrationWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeLocalTemperatureCalibration(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeLocalTemperatureCalibrationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeLocalTemperatureCalibrationWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeLocalTemperatureCalibrationWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMaxCoolSetpointLimit(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMaxCoolSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMaxCoolSetpointLimitWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeMaxCoolSetpointLimitWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMaxCoolSetpointLimit(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMaxCoolSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMaxCoolSetpointLimitWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeMaxCoolSetpointLimitWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMaxHeatSetpointLimit(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMaxHeatSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMaxHeatSetpointLimitWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeMaxHeatSetpointLimitWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMaxHeatSetpointLimit(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMaxHeatSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMaxHeatSetpointLimitWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeMaxHeatSetpointLimitWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinCoolSetpointLimit(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinCoolSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinCoolSetpointLimitWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeMinCoolSetpointLimitWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinCoolSetpointLimit(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinCoolSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinCoolSetpointLimitWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeMinCoolSetpointLimitWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinHeatSetpointLimit(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinHeatSetpointLimitWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinHeatSetpointLimitWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeMinHeatSetpointLimitWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinHeatSetpointLimit(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinHeatSetpointLimitWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinHeatSetpointLimitWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeMinHeatSetpointLimitWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinSetpointDeadBand(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinSetpointDeadBandWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinSetpointDeadBandWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeMinSetpointDeadBandWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeMinSetpointDeadBand(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeMinSetpointDeadBandWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeMinSetpointDeadBandWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeMinSetpointDeadBandWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedCoolingSetpoint(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedCoolingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedCoolingSetpointWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedCoolingSetpointWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedCoolingSetpoint(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedCoolingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedCoolingSetpointWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedCoolingSetpointWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedHeatingSetpoint(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedHeatingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedHeatingSetpointWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedHeatingSetpointWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedHeatingSetpoint(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedHeatingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedHeatingSetpointWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedHeatingSetpointWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedSetback(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedSetbackWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedSetbackWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedSetbackWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeOccupiedSetback(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeOccupiedSetbackWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOccupiedSetbackWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOccupiedSetbackWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributePresets(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributePresetsWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePresetsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributePresetsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributePresets(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributePresetsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePresetsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributePresetsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeRemoteSensing(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeRemoteSensingWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRemoteSensingWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeRemoteSensingWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeRemoteSensing(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeRemoteSensingWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRemoteSensingWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeRemoteSensingWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeSchedules(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeSchedulesWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSchedulesWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeSchedulesWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeSchedules(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeSchedulesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSchedulesWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeSchedulesWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeSystemMode(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeSystemModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSystemModeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeSystemModeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeSystemMode(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeSystemModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSystemModeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeSystemModeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeTemperatureSetpointHold(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeTemperatureSetpointHoldWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeTemperatureSetpointHoldWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeTemperatureSetpointHoldWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeTemperatureSetpointHold(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeTemperatureSetpointHoldWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeTemperatureSetpointHoldWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeTemperatureSetpointHoldWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeTemperatureSetpointHoldDuration(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeTemperatureSetpointHoldDurationWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeTemperatureSetpointHoldDurationWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeTemperatureSetpointHoldDurationWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeTemperatureSetpointHoldDuration(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeTemperatureSetpointHoldDurationWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeTemperatureSetpointHoldDurationWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeTemperatureSetpointHoldDurationWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeThermostatProgrammingOperationMode(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeThermostatProgrammingOperationModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeThermostatProgrammingOperationModeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeThermostatProgrammingOperationModeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeThermostatProgrammingOperationMode(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeThermostatProgrammingOperationModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeThermostatProgrammingOperationModeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeThermostatProgrammingOperationModeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedCoolingSetpoint(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedCoolingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedCoolingSetpointWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedCoolingSetpointWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedCoolingSetpoint(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedCoolingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedCoolingSetpointWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedCoolingSetpointWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedHeatingSetpoint(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedHeatingSetpointWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedHeatingSetpointWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedHeatingSetpointWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedHeatingSetpoint(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedHeatingSetpointWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedHeatingSetpointWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedHeatingSetpointWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedSetback(withValue:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedSetbackWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedSetbackWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedSetbackWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat/writeAttributeUnoccupiedSetback(withValue:params:completion:)
func (m_ MTRBaseClusterThermostat) WriteAttributeUnoccupiedSetbackWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUnoccupiedSetbackWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeUnoccupiedSetbackWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterThermostat */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterThermostat */


