// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

// Network Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (392 total):

// nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor nw_advertise_descriptor_t, ) nw_txt_record_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_advertise_descriptor_create_application_service(application_service_name const char  *, ) nw_advertise_descriptor_t
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_advertise_descriptor_create_bonjour_service(name const char  *, type ,  const char  *, domain ,  const char  *, ) nw_advertise_descriptor_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_advertise_descriptor_get_application_service_name(advertise_descriptor nw_advertise_descriptor_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor nw_advertise_descriptor_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor nw_advertise_descriptor_t, no_auto_rename ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_advertise_descriptor_set_txt_record(advertise_descriptor nw_advertise_descriptor_t, txt_record ,  const void  *, txt_length ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_set_txt_record_object(advertise_descriptor nw_advertise_descriptor_t, txt_record ,  nw_txt_record_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_create_application_service(application_service_name const char  *, ) nw_browse_descriptor_t
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+


// nw_browse_descriptor_create_bonjour_service(type const char  *, domain ,  const char  *, ) nw_browse_descriptor_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_get_application_service_name(descriptor nw_browse_descriptor_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_browse_descriptor_get_bonjour_service_domain(descriptor nw_browse_descriptor_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browse_descriptor_get_bonjour_service_type(descriptor nw_browse_descriptor_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_get_include_txt_record(descriptor nw_browse_descriptor_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_set_include_txt_record(descriptor nw_browse_descriptor_t, include_txt_record ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browse_result_copy_endpoint(result nw_browse_result_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_copy_txt_record_object(result nw_browse_result_t, ) nw_txt_record_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_enumerate_interfaces(result nw_browse_result_t, enumerator ,  nw_browse_result_enumerate_interface_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browse_result_get_changes(old_result nw_browse_result_t, new_result ,  nw_browse_result_t, ) nw_browse_result_change_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_get_interfaces_count(result nw_browse_result_t, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_cancel(browser nw_browser_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browser_copy_browse_descriptor(browser nw_browser_t, ) nw_browse_descriptor_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_copy_parameters(browser nw_browser_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_create(descriptor nw_browse_descriptor_t, parameters ,  nw_parameters_t, ) nw_browser_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browser_set_browse_results_changed_handler(browser nw_browser_t, handler ,  nw_browser_browse_results_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_set_queue(browser nw_browser_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_set_state_changed_handler(browser nw_browser_t, state_changed_handler ,  nw_browser_state_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_browser_start(browser nw_browser_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_connection_access_establishment_report(connection nw_connection_t, queue ,  dispatch_queue_t, access_block ,  nw_establishment_report_access_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_connection_batch(connection nw_connection_t, batch_block ,  dispatch_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_cancel(connection nw_connection_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_cancel_current_endpoint(connection nw_connection_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_current_path(connection nw_connection_t, ) nw_path_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_copy_description(connection nw_connection_t, ) char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_endpoint(connection nw_connection_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_parameters(connection nw_connection_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_copy_protocol_metadata(connection nw_connection_t, definition ,  nw_protocol_definition_t, ) nw_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_create(endpoint nw_endpoint_t, parameters ,  nw_parameters_t, ) nw_connection_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_create_new_data_transfer_report(connection nw_connection_t, ) nw_data_transfer_report_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_connection_force_cancel(connection nw_connection_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_get_maximum_datagram_size(connection nw_connection_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_group_cancel(group nw_connection_group_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_connection_group_copy_descriptor(group nw_connection_group_t, ) nw_group_descriptor_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_local_endpoint_for_message(group nw_connection_group_t, context ,  nw_content_context_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_parameters(group nw_connection_group_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_connection_group_copy_path_for_message(group nw_connection_group_t, context ,  nw_content_context_t, ) nw_path_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_protocol_metadata(group nw_connection_group_t, definition ,  nw_protocol_definition_t, ) nw_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_copy_protocol_metadata_for_message(group nw_connection_group_t, context ,  nw_content_context_t, definition ,  nw_protocol_definition_t, ) nw_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_connection_group_copy_remote_endpoint_for_message(group nw_connection_group_t, context ,  nw_content_context_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_create(group_descriptor nw_group_descriptor_t, parameters ,  nw_parameters_t, ) nw_connection_group_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_extract_connection(group nw_connection_group_t, endpoint ,  nw_endpoint_t, protocol_options ,  nw_protocol_options_t, ) nw_connection_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_connection_group_extract_connection_for_message(group nw_connection_group_t, context ,  nw_content_context_t, ) nw_connection_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_reinsert_extracted_connection(group nw_connection_group_t, connection ,  nw_connection_t, ) bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_reply(group nw_connection_group_t, inbound_message ,  nw_content_context_t, outbound_message ,  nw_content_context_t, content ,  dispatch_data_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_connection_group_send_message(group nw_connection_group_t, content ,  dispatch_data_t, endpoint ,  nw_endpoint_t, context ,  nw_content_context_t, completion ,  nw_connection_group_send_completion_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_set_new_connection_handler(group nw_connection_group_t, new_connection_handler ,  nw_connection_group_new_connection_handler_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_set_queue(group nw_connection_group_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_connection_group_set_receive_handler(group nw_connection_group_t, maximum_message_size ,  uint32_t, reject_oversized_messages ,  bool, receive_handler ,  nw_connection_group_receive_handler_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_set_state_changed_handler(group nw_connection_group_t, state_changed_handler ,  nw_connection_group_state_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_start(group nw_connection_group_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_connection_receive(connection nw_connection_t, minimum_incomplete_length ,  uint32_t, maximum_length ,  uint32_t, completion ,  nw_connection_receive_completion_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_receive_message(connection nw_connection_t, completion ,  nw_connection_receive_completion_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_restart(connection nw_connection_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_send(connection nw_connection_t, content ,  dispatch_data_t, context ,  nw_content_context_t, is_complete ,  bool, completion ,  nw_connection_send_completion_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_better_path_available_handler(connection nw_connection_t, handler ,  nw_connection_boolean_event_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_path_changed_handler(connection nw_connection_t, handler ,  nw_connection_path_event_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_set_queue(connection nw_connection_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_state_changed_handler(connection nw_connection_t, handler ,  nw_connection_state_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_viability_changed_handler(connection nw_connection_t, handler ,  nw_connection_boolean_event_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_connection_start(connection nw_connection_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_copy_antecedent(context nw_content_context_t, ) nw_content_context_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_copy_protocol_metadata(context nw_content_context_t, protocol ,  nw_protocol_definition_t, ) nw_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_content_context_create(context_identifier const char  *, ) nw_content_context_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_foreach_protocol_metadata(context nw_content_context_t, foreach_block ,  void  (^, definition )( nw_protocol_definition_t, metadata ,  nw_protocol_metadata_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_expiration_milliseconds(context nw_content_context_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_content_context_get_identifier(context nw_content_context_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_is_final(context nw_content_context_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_relative_priority(context nw_content_context_t, ) double
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_content_context_set_antecedent(context nw_content_context_t, antecedent_context ,  nw_content_context_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_expiration_milliseconds(context nw_content_context_t, expiration_milliseconds ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_is_final(context nw_content_context_t, is_final ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_content_context_set_metadata_for_protocol(context nw_content_context_t, protocol_metadata ,  nw_protocol_metadata_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_relative_priority(context nw_content_context_t, relative_priority ,  double, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_data_transfer_report_collect(report nw_data_transfer_report_t, queue ,  dispatch_queue_t, collect_block ,  nw_data_transfer_report_collect_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_copy_path_interface(report nw_data_transfer_report_t, path_index ,  uint32_t, ) nw_interface_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_duration_milliseconds(report nw_data_transfer_report_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_path_count(report nw_data_transfer_report_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_get_path_radio_type(report nw_data_transfer_report_t, path_index ,  uint32_t, ) nw_interface_radio_type_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_data_transfer_report_get_received_application_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_ip_packet_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_get_received_transport_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_transport_duplicate_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_get_sent_application_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_ip_packet_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_transport_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_state(report nw_data_transfer_report_t, ) nw_data_transfer_report_state_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_data_transfer_report_get_transport_rtt_variance(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report nw_data_transfer_report_t, path_index ,  uint32_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_endpoint_copy_address_string(endpoint nw_endpoint_t, ) char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_endpoint_copy_port_string(endpoint nw_endpoint_t, ) char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_copy_txt_record(endpoint nw_endpoint_t, ) nw_txt_record_t
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_endpoint_create_address(address const struct sockaddr  *, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_endpoint_create_bonjour_service(name const char  *, type ,  const char  *, domain ,  const char  *, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_create_host(hostname const char  *, port ,  const char  *, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_create_url(url const char  *, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_endpoint_get_address(endpoint nw_endpoint_t, ) const struct sockaddr  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_bonjour_service_domain(endpoint nw_endpoint_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_bonjour_service_name(endpoint nw_endpoint_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_endpoint_get_bonjour_service_type(endpoint nw_endpoint_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_hostname(endpoint nw_endpoint_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_port(endpoint nw_endpoint_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_endpoint_get_signature(endpoint nw_endpoint_t, out_signature_length ,  size_t  *, ) const uint8_t  *
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_endpoint_get_type(endpoint nw_endpoint_t, ) nw_endpoint_type_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_url(endpoint nw_endpoint_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_error_copy_cf_error(error nw_error_t, ) CFErrorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_error_get_error_code(error nw_error_t, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_error_get_error_domain(error nw_error_t, ) nw_error_domain_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_establishment_report_copy_proxy_endpoint(report nw_establishment_report_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_enumerate_protocols(report nw_establishment_report_t, enumerate_block ,  nw_report_protocol_enumerator_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_enumerate_resolution_reports(report nw_establishment_report_t, enumerate_block ,  nw_report_resolution_report_enumerator_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_establishment_report_enumerate_resolutions(report nw_establishment_report_t, enumerate_block ,  nw_report_resolution_enumerator_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_attempt_started_after_milliseconds(report nw_establishment_report_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_duration_milliseconds(report nw_establishment_report_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_establishment_report_get_previous_attempt_count(report nw_establishment_report_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_proxy_configured(report nw_establishment_report_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_used_proxy(report nw_establishment_report_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ethernet_channel_cancel(ethernet_channel nw_ethernet_channel_t, )
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_create(ether_type uint16_t, interface ,  nw_interface_t, ) nw_ethernet_channel_t
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_create_with_parameters(ether_type uint16_t, interface ,  nw_interface_t, parameters ,  nw_parameters_t, ) nw_ethernet_channel_t
//
// Availability:
//   - macOS 13.0+


// nw_ethernet_channel_get_maximum_payload_size(ethernet_channel nw_ethernet_channel_t, ) uint32_t
//
// Availability:
//   - macOS 13.0+

// nw_ethernet_channel_send(ethernet_channel nw_ethernet_channel_t, content ,  dispatch_data_t, vlan_tag ,  uint16_t, remote_address ,  nw_ethernet_address_t, completion ,  nw_ethernet_channel_send_completion_t, )
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_set_queue(ethernet_channel nw_ethernet_channel_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - macOS 10.15+


// nw_ethernet_channel_set_receive_handler(ethernet_channel nw_ethernet_channel_t, handler ,  nw_ethernet_channel_receive_handler_t, )
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_set_state_changed_handler(ethernet_channel nw_ethernet_channel_t, handler ,  nw_ethernet_channel_state_changed_handler_t, )
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_start(ethernet_channel nw_ethernet_channel_t, )
//
// Availability:
//   - macOS 10.15+


// nw_framer_async(framer nw_framer_t, async_block ,  nw_framer_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_local_endpoint(framer nw_framer_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_options(framer nw_framer_t, ) nw_protocol_options_t
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+


// nw_framer_copy_parameters(framer nw_framer_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_remote_endpoint(framer nw_framer_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_create_definition(identifier const char  *, flags ,  uint32_t, start_handler ,  nw_framer_start_handler_t, ) nw_protocol_definition_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_create_options(framer_definition nw_protocol_definition_t, ) nw_protocol_options_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_deliver_input(framer nw_framer_t, input_buffer ,  const uint8_t  *, input_length ,  size_t, message ,  nw_framer_message_t, is_complete ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_deliver_input_no_copy(framer nw_framer_t, input_length ,  size_t, message ,  nw_framer_message_t, is_complete ,  bool, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_mark_failed_with_error(framer nw_framer_t, error_code ,  int, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_mark_ready(framer nw_framer_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_access_value(message nw_framer_message_t, key ,  const char  *, access_value ,  bool  (^, value )( const void  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_message_copy_object_value(message nw_framer_message_t, key ,  const char  *, ) id
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_create(framer nw_framer_t, ) nw_framer_message_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_set_object_value(message nw_framer_message_t, key ,  const char  *, value ,  id, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_message_set_value(message nw_framer_message_t, key ,  const char  *, value ,  void  *, dispose_value ,  nw_framer_message_dispose_value_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_options_copy_object_value(options nw_protocol_options_t, key ,  const char  *, ) id
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+

// nw_framer_options_set_object_value(options nw_protocol_options_t, key ,  const char  *, value ,  id, )
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+


// nw_framer_parse_input(framer nw_framer_t, minimum_incomplete_length ,  size_t, maximum_length ,  size_t, temp_buffer ,  uint8_t  *, parse ,  nw_framer_parse_completion_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_parse_output(framer nw_framer_t, minimum_incomplete_length ,  size_t, maximum_length ,  size_t, temp_buffer ,  uint8_t  *, parse ,  nw_framer_parse_completion_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_pass_through_input(framer nw_framer_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_pass_through_output(framer nw_framer_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_prepend_application_protocol(framer nw_framer_t, protocol_options ,  nw_protocol_options_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_protocol_create_message(definition nw_protocol_definition_t, ) nw_framer_message_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_schedule_wakeup(framer nw_framer_t, milliseconds ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_cleanup_handler(framer nw_framer_t, cleanup_handler ,  nw_framer_cleanup_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_input_handler(framer nw_framer_t, input_handler ,  nw_framer_input_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_set_output_handler(framer nw_framer_t, output_handler ,  nw_framer_output_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_stop_handler(framer nw_framer_t, stop_handler ,  nw_framer_stop_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_wakeup_handler(framer nw_framer_t, wakeup_handler ,  nw_framer_wakeup_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_framer_write_output(framer nw_framer_t, output_buffer ,  const uint8_t  *, output_length ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_write_output_data(framer nw_framer_t, output_data ,  dispatch_data_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_write_output_no_copy(framer nw_framer_t, output_length ,  size_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_group_descriptor_add_endpoint(descriptor nw_group_descriptor_t, endpoint ,  nw_endpoint_t, ) bool
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_group_descriptor_create_multicast(multicast_group nw_endpoint_t, ) nw_group_descriptor_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_group_descriptor_create_multiplex(remote_endpoint nw_endpoint_t, ) nw_group_descriptor_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_group_descriptor_enumerate_endpoints(descriptor nw_group_descriptor_t, enumerate_block ,  nw_group_descriptor_enumerate_endpoints_block_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_interface_get_index(interface nw_interface_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_interface_get_name(interface nw_interface_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_interface_get_type(interface nw_interface_t, ) nw_interface_type_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_get_ecn_flag(metadata nw_protocol_metadata_t, ) nw_ip_ecn_flag_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_get_receive_time(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_ip_metadata_get_service_class(metadata nw_protocol_metadata_t, ) nw_service_class_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_set_ecn_flag(metadata nw_protocol_metadata_t, ecn_flag ,  nw_ip_ecn_flag_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_set_service_class(metadata nw_protocol_metadata_t, service_class ,  nw_service_class_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_ip_options_set_calculate_receive_time(options nw_protocol_options_t, calculate_receive_time ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_disable_fragmentation(options nw_protocol_options_t, disable_fragmentation ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_disable_multicast_loopback(options nw_protocol_options_t, disable_multicast_loopback ,  bool, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_ip_options_set_hop_limit(options nw_protocol_options_t, hop_limit ,  uint8_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_local_address_preference(options nw_protocol_options_t, preference ,  nw_ip_local_address_preference_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ip_options_set_use_minimum_mtu(options nw_protocol_options_t, use_minimum_mtu ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_ip_options_set_version(options nw_protocol_options_t, version ,  nw_ip_version_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_cancel(listener nw_listener_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_create(parameters nw_parameters_t, ) nw_listener_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_listener_create_with_connection(connection nw_connection_t, parameters ,  nw_parameters_t, ) nw_listener_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_create_with_launchd_key(parameters nw_parameters_t, launchd_key ,  const char  *, ) nw_listener_t
//
// Availability:
//   - macOS 10.14+

// nw_listener_create_with_port(port const char  *, parameters ,  nw_parameters_t, ) nw_listener_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_listener_get_new_connection_limit(listener nw_listener_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.15+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_get_port(listener nw_listener_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_advertise_descriptor(listener nw_listener_t, advertise_descriptor ,  nw_advertise_descriptor_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_listener_set_advertised_endpoint_changed_handler(listener nw_listener_t, handler ,  nw_listener_advertised_endpoint_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_new_connection_group_handler(listener nw_listener_t, handler ,  nw_listener_new_connection_group_handler_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_listener_set_new_connection_handler(listener nw_listener_t, handler ,  nw_listener_new_connection_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_listener_set_new_connection_limit(listener nw_listener_t, new_connection_limit ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.15+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_queue(listener nw_listener_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_state_changed_handler(listener nw_listener_t, handler ,  nw_listener_state_changed_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_listener_start(listener nw_listener_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor nw_group_descriptor_t, ) bool
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor nw_group_descriptor_t, disable_unicast_traffic ,  bool, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_multicast_group_descriptor_set_specific_source(multicast_descriptor nw_group_descriptor_t, source ,  nw_endpoint_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_parameters_clear_prohibited_interface_types(parameters nw_parameters_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_clear_prohibited_interfaces(parameters nw_parameters_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_copy(parameters nw_parameters_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy_default_protocol_stack(parameters nw_parameters_t, ) nw_protocol_stack_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy_local_endpoint(parameters nw_parameters_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_copy_required_interface(parameters nw_parameters_t, ) nw_interface_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_create_custom_ip(custom_ip_protocol_number uint8_t, configure_ip ,  nw_parameters_configure_protocol_block_t, ) nw_parameters_t
//
// Availability:
//   - macOS 10.15+

// nw_parameters_create_quic(configure_quic nw_parameters_configure_protocol_block_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_parameters_create_secure_tcp(configure_tls nw_parameters_configure_protocol_block_t, configure_tcp ,  nw_parameters_configure_protocol_block_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_create_secure_udp(configure_dtls nw_parameters_configure_protocol_block_t, configure_udp ,  nw_parameters_configure_protocol_block_t, ) nw_parameters_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_allow_ultra_constrained(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// nw_parameters_get_attribution(parameters nw_parameters_t, ) nw_parameters_attribution_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_parameters_get_expired_dns_behavior(parameters nw_parameters_t, ) nw_parameters_expired_dns_behavior_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_fast_open_enabled(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_get_include_peer_to_peer(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_local_only(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_multipath_service(parameters nw_parameters_t, ) nw_multipath_service_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_get_prefer_no_proxy(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_prohibit_constrained(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_parameters_get_prohibit_expensive(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_get_required_interface_type(parameters nw_parameters_t, ) nw_interface_type_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_reuse_local_address(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_service_class(parameters nw_parameters_t, ) nw_service_class_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_iterate_prohibited_interface_types(parameters nw_parameters_t, iterate_block ,  nw_parameters_iterate_interface_types_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_iterate_prohibited_interfaces(parameters nw_parameters_t, iterate_block ,  nw_parameters_iterate_interfaces_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_prohibit_interface(parameters nw_parameters_t, interface ,  nw_interface_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_prohibit_interface_type(parameters nw_parameters_t, interface_type ,  nw_interface_type_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_require_interface(parameters nw_parameters_t, interface ,  nw_interface_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_requires_dnssec_validation(parameters nw_parameters_t, ) bool
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+


// nw_parameters_set_allow_ultra_constrained(parameters nw_parameters_t, allow_ultra_constrained ,  bool, )
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// nw_parameters_set_attribution(parameters nw_parameters_t, attribution ,  nw_parameters_attribution_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_parameters_set_expired_dns_behavior(parameters nw_parameters_t, expired_dns_behavior ,  nw_parameters_expired_dns_behavior_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_set_fast_open_enabled(parameters nw_parameters_t, fast_open_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_include_peer_to_peer(parameters nw_parameters_t, include_peer_to_peer ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_local_endpoint(parameters nw_parameters_t, local_endpoint ,  nw_endpoint_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_set_local_only(parameters nw_parameters_t, local_only ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_multipath_service(parameters nw_parameters_t, multipath_service ,  nw_multipath_service_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_prefer_no_proxy(parameters nw_parameters_t, prefer_no_proxy ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_set_privacy_context(parameters nw_parameters_t, privacy_context ,  nw_privacy_context_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_parameters_set_prohibit_constrained(parameters nw_parameters_t, prohibit_constrained ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_parameters_set_prohibit_expensive(parameters nw_parameters_t, prohibit_expensive ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_set_required_interface_type(parameters nw_parameters_t, interface_type ,  nw_interface_type_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_requires_dnssec_validation(parameters nw_parameters_t, requires_dnssec_validation ,  bool, )
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_parameters_set_reuse_local_address(parameters nw_parameters_t, reuse_local_address ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_parameters_set_service_class(parameters nw_parameters_t, service_class ,  nw_service_class_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_copy_effective_local_endpoint(path nw_path_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_copy_effective_remote_endpoint(path nw_path_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_path_enumerate_gateways(path nw_path_t, enumerate_block ,  nw_path_enumerate_gateways_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_path_enumerate_interfaces(path nw_path_t, enumerate_block ,  nw_path_enumerate_interfaces_block_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_get_link_quality(path nw_path_t, ) nw_link_quality_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// nw_path_get_status(path nw_path_t, ) nw_path_status_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_get_unsatisfied_reason(path nw_path_t, ) nw_path_unsatisfied_reason_t
//
// Availability:
//   - Mac Catalyst 14.2+
//   - iOS 14.2+
//   - iPadOS 14.2+
//   - macOS 11.0+
//   - tvOS 14.2+
//   - visionOS 1.0+
//   - watchOS 7.1+

// nw_path_has_dns(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_path_has_ipv4(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_has_ipv6(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_constrained(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_path_is_equal(path nw_path_t, other_path ,  nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_expensive(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_ultra_constrained(path nw_path_t, ) bool
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// nw_path_monitor_cancel(monitor nw_path_monitor_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_create_with_type(required_interface_type nw_interface_type_t, ) nw_path_monitor_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_prohibit_interface_type(monitor nw_path_monitor_t, interface_type ,  nw_interface_type_t, )
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_path_monitor_set_cancel_handler(monitor nw_path_monitor_t, cancel_handler ,  nw_path_monitor_cancel_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_set_queue(monitor nw_path_monitor_t, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_set_update_handler(monitor nw_path_monitor_t, update_handler ,  nw_path_monitor_update_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_path_monitor_start(monitor nw_path_monitor_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_uses_interface_type(path nw_path_t, interface_type ,  nw_interface_type_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_copy_definition(metadata nw_protocol_metadata_t, ) nw_protocol_definition_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_protocol_metadata_is_framer_message(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_protocol_metadata_is_ip(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_quic(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_protocol_metadata_is_tcp(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_tls(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_udp(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_protocol_metadata_is_ws(metadata nw_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_protocol_options_is_quic(options nw_protocol_options_t, ) bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_proxy_config_add_excluded_domain(config nw_proxy_config_t, excluded_domain ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// nw_proxy_config_add_match_domain(config nw_proxy_config_t, match_domain ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_clear_excluded_domains(config nw_proxy_config_t, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_clear_match_domains(config nw_proxy_config_t, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// nw_proxy_config_enumerate_excluded_domains(config nw_proxy_config_t, enumerator ,  nw_proxy_domain_enumerator_t, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_enumerate_match_domains(config nw_proxy_config_t, enumerator ,  nw_proxy_domain_enumerator_t, )
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_quic_add_tls_application_protocol(options nw_protocol_options_t, application_protocol ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_copy_sec_protocol_metadata(metadata nw_protocol_metadata_t, ) sec_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_copy_sec_protocol_options(options nw_protocol_options_t, ) sec_protocol_options_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_application_error(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_application_error_reason(metadata nw_protocol_metadata_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_idle_timeout(options nw_protocol_options_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_data(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_initial_max_stream_data_bidirectional_local(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_stream_data_bidirectional_remote(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_stream_data_unidirectional(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_initial_max_streams_bidirectional(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_streams_unidirectional(options nw_protocol_options_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_keepalive_interval(metadata nw_protocol_metadata_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_local_max_streams_bidirectional(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_local_max_streams_unidirectional(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_max_datagram_frame_size(options nw_protocol_options_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+


// nw_quic_get_max_udp_payload_size(options nw_protocol_options_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_remote_idle_timeout(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_remote_max_streams_bidirectional(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_remote_max_streams_unidirectional(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_application_error(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_id(metadata nw_protocol_metadata_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_stream_is_datagram(options nw_protocol_options_t, ) bool
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_get_stream_is_unidirectional(options nw_protocol_options_t, ) bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_type(stream_metadata nw_protocol_metadata_t, ) uint8_t
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_get_stream_usable_datagram_frame_size(metadata nw_protocol_metadata_t, ) uint16_t
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_application_error(metadata nw_protocol_metadata_t, application_error ,  uint64_t, reason ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_idle_timeout(options nw_protocol_options_t, idle_timeout ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_set_initial_max_data(options nw_protocol_options_t, initial_max_data ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_stream_data_bidirectional_local(options nw_protocol_options_t, initial_max_stream_data_bidirectional_local ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_stream_data_bidirectional_remote(options nw_protocol_options_t, initial_max_stream_data_bidirectional_remote ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_set_initial_max_stream_data_unidirectional(options nw_protocol_options_t, initial_max_stream_data_unidirectional ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_streams_bidirectional(options nw_protocol_options_t, initial_max_streams_bidirectional ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_streams_unidirectional(options nw_protocol_options_t, initial_max_streams_unidirectional ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_set_keepalive_interval(metadata nw_protocol_metadata_t, keepalive_interval ,  uint16_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_local_max_streams_bidirectional(metadata nw_protocol_metadata_t, max_streams_bidirectional ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_local_max_streams_unidirectional(metadata nw_protocol_metadata_t, max_streams_unidirectional ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_set_max_datagram_frame_size(options nw_protocol_options_t, max_datagram_frame_size ,  uint16_t, )
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_max_udp_payload_size(options nw_protocol_options_t, max_udp_payload_size ,  uint16_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_stream_application_error(metadata nw_protocol_metadata_t, application_error ,  uint64_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// nw_quic_set_stream_is_datagram(options nw_protocol_options_t, is_datagram ,  bool, )
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_stream_is_unidirectional(options nw_protocol_options_t, is_unidirectional ,  bool, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_release(obj void  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_resolution_report_copy_preferred_endpoint(resolution_report nw_resolution_report_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_copy_successful_endpoint(resolution_report nw_resolution_report_t, ) nw_endpoint_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_endpoint_count(resolution_report nw_resolution_report_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_resolution_report_get_milliseconds(resolution_report nw_resolution_report_t, ) uint64_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_protocol(resolution_report nw_resolution_report_t, ) nw_report_resolution_protocol_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_source(resolution_report nw_resolution_report_t, ) nw_report_resolution_source_t
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+


// nw_retain(obj void  *, ) void  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_get_available_receive_buffer(metadata nw_protocol_metadata_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_get_available_send_buffer(metadata nw_protocol_metadata_t, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_connection_timeout(options nw_protocol_options_t, connection_timeout ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_disable_ack_stretching(options nw_protocol_options_t, disable_ack_stretching ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_disable_ecn(options nw_protocol_options_t, disable_ecn ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_enable_fast_open(options nw_protocol_options_t, enable_fast_open ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_enable_keepalive(options nw_protocol_options_t, enable_keepalive ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_keepalive_count(options nw_protocol_options_t, keepalive_count ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_keepalive_idle_time(options nw_protocol_options_t, keepalive_idle_time ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_keepalive_interval(options nw_protocol_options_t, keepalive_interval ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_maximum_segment_size(options nw_protocol_options_t, maximum_segment_size ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_multipath_force_version(options nw_protocol_options_t, multipath_force_version ,  nw_multipath_version_t, )
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_tcp_options_set_no_delay(options nw_protocol_options_t, no_delay ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_no_options(options nw_protocol_options_t, no_options ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_no_push(options nw_protocol_options_t, no_push ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_persist_timeout(options nw_protocol_options_t, persist_timeout ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_retransmit_connection_drop_time(options nw_protocol_options_t, retransmit_connection_drop_time ,  uint32_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_tcp_options_set_retransmit_fin_drop(options nw_protocol_options_t, retransmit_fin_drop ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tls_copy_sec_protocol_metadata(metadata nw_protocol_metadata_t, ) sec_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tls_copy_sec_protocol_options(options nw_protocol_options_t, ) sec_protocol_options_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_txt_record_access_bytes(txt_record nw_txt_record_t, access_bytes ,  nw_txt_record_access_bytes_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_access_key(txt_record nw_txt_record_t, key ,  const char  *, access_value ,  nw_txt_record_access_key_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_apply(txt_record nw_txt_record_t, applier ,  nw_txt_record_applier_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_txt_record_copy(txt_record nw_txt_record_t, ) nw_txt_record_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_create_with_bytes(txt_bytes const uint8_t  *, txt_len ,  size_t, ) nw_txt_record_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_find_key(txt_record nw_txt_record_t, key ,  const char  *, ) nw_txt_record_find_key_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_txt_record_get_key_count(txt_record nw_txt_record_t, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_is_dictionary(txt_record nw_txt_record_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_is_equal(left nw_txt_record_t, right ,  nw_txt_record_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_txt_record_remove_key(txt_record nw_txt_record_t, key ,  const char  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_set_key(txt_record nw_txt_record_t, key ,  const char  *, value ,  const uint8_t  *, value_len ,  size_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_udp_options_set_prefer_no_checksum(options nw_protocol_options_t, prefer_no_checksum ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// nw_ws_create_metadata(opcode nw_ws_opcode_t, ) nw_protocol_metadata_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_create_options(version nw_ws_version_t, ) nw_protocol_options_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_copy_server_response(metadata nw_protocol_metadata_t, ) nw_ws_response_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_metadata_get_close_code(metadata nw_protocol_metadata_t, ) nw_ws_close_code_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_get_opcode(metadata nw_protocol_metadata_t, ) nw_ws_opcode_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_set_close_code(metadata nw_protocol_metadata_t, close_code ,  nw_ws_close_code_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_metadata_set_pong_handler(metadata nw_protocol_metadata_t, client_queue ,  dispatch_queue_t, pong_handler ,  nw_ws_pong_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_add_additional_header(options nw_protocol_options_t, name ,  const char  *, value ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_add_subprotocol(options nw_protocol_options_t, subprotocol ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_options_set_auto_reply_ping(options nw_protocol_options_t, auto_reply_ping ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_client_request_handler(options nw_protocol_options_t, client_queue ,  dispatch_queue_t, handler ,  nw_ws_client_request_handler_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_maximum_message_size(options nw_protocol_options_t, maximum_message_size ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_options_set_skip_handshake(options nw_protocol_options_t, skip_handshake ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_request_enumerate_additional_headers(request nw_ws_request_t, enumerator ,  nw_ws_additional_header_enumerator_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_request_enumerate_subprotocols(request nw_ws_request_t, enumerator ,  nw_ws_subprotocol_enumerator_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_response_add_additional_header(response nw_ws_response_t, name ,  const char  *, value ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_create(status nw_ws_response_status_t, selected_subprotocol ,  const char  *, ) nw_ws_response_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_enumerate_additional_headers(response nw_ws_response_t, enumerator ,  nw_ws_additional_header_enumerator_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// nw_ws_response_get_selected_subprotocol(response nw_ws_response_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_get_status(response nw_ws_response_t, ) nw_ws_response_status_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

