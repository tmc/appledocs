// Code generated from Apple documentation for Network. DO NOT EDIT.

package network


import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
)


// Network Functions (417 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_nw_advertise_descriptor_copy_txt_record_object func(Nw_advertise_descriptor_t) Nw_txt_record_t
	_nw_advertise_descriptor_create_application_service func(unsafe.Pointer) Nw_advertise_descriptor_t
	_nw_advertise_descriptor_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Nw_advertise_descriptor_t
	_nw_advertise_descriptor_get_application_service_name func(Nw_advertise_descriptor_t) unsafe.Pointer
	_nw_advertise_descriptor_get_no_auto_rename func(Nw_advertise_descriptor_t) bool
	_nw_advertise_descriptor_set_no_auto_rename func(Nw_advertise_descriptor_t, bool)
	_nw_advertise_descriptor_set_txt_record func(Nw_advertise_descriptor_t, unsafe.Pointer, uintptr)
	_nw_advertise_descriptor_set_txt_record_object func(Nw_advertise_descriptor_t, Nw_txt_record_t)
	_nw_browse_descriptor_create_application_service func(unsafe.Pointer) Nw_browse_descriptor_t
	_nw_browse_descriptor_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer) Nw_browse_descriptor_t
	_nw_browse_descriptor_get_application_service_name func(Nw_browse_descriptor_t) unsafe.Pointer
	_nw_browse_descriptor_get_bonjour_service_domain func(Nw_browse_descriptor_t) unsafe.Pointer
	_nw_browse_descriptor_get_bonjour_service_type func(Nw_browse_descriptor_t) unsafe.Pointer
	_nw_browse_descriptor_get_include_txt_record func(Nw_browse_descriptor_t) bool
	_nw_browse_descriptor_set_include_txt_record func(Nw_browse_descriptor_t, bool)
	_nw_browse_result_copy_endpoint func(Nw_browse_result_t) Nw_endpoint_t
	_nw_browse_result_copy_txt_record_object func(Nw_browse_result_t) Nw_txt_record_t
	_nw_browse_result_enumerate_interfaces func(Nw_browse_result_t, unsafe.Pointer)
	_nw_browse_result_get_changes func(Nw_browse_result_t, Nw_browse_result_t) Nw_browse_result_change_t
	_nw_browse_result_get_interfaces_count func(Nw_browse_result_t) uintptr
	_nw_browser_cancel func(Nw_browser_t)
	_nw_browser_copy_browse_descriptor func(Nw_browser_t) Nw_browse_descriptor_t
	_nw_browser_copy_parameters func(Nw_browser_t) Nw_parameters_t
	_nw_browser_create func(Nw_browse_descriptor_t, Nw_parameters_t) Nw_browser_t
	_nw_browser_set_browse_results_changed_handler func(Nw_browser_t, unsafe.Pointer)
	_nw_browser_set_queue func(Nw_browser_t, unsafe.Pointer)
	_nw_browser_set_state_changed_handler func(Nw_browser_t, unsafe.Pointer)
	_nw_browser_start func(Nw_browser_t)
	_nw_connection_access_establishment_report func(Nw_connection_t, unsafe.Pointer, unsafe.Pointer)
	_nw_connection_batch func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_cancel func(Nw_connection_t)
	_nw_connection_cancel_current_endpoint func(Nw_connection_t)
	_nw_connection_copy_current_path func(Nw_connection_t) Nw_path_t
	_nw_connection_copy_description func(Nw_connection_t) unsafe.Pointer
	_nw_connection_copy_endpoint func(Nw_connection_t) Nw_endpoint_t
	_nw_connection_copy_parameters func(Nw_connection_t) Nw_parameters_t
	_nw_connection_copy_protocol_metadata func(Nw_connection_t, Nw_protocol_definition_t) Nw_protocol_metadata_t
	_nw_connection_create func(Nw_endpoint_t, Nw_parameters_t) Nw_connection_t
	_nw_connection_create_new_data_transfer_report func(Nw_connection_t) Nw_data_transfer_report_t
	_nw_connection_force_cancel func(Nw_connection_t)
	_nw_connection_get_maximum_datagram_size func(Nw_connection_t) uint32
	_nw_connection_group_cancel func(Nw_connection_group_t)
	_nw_connection_group_copy_descriptor func(Nw_connection_group_t) Nw_group_descriptor_t
	_nw_connection_group_copy_parameters func(Nw_connection_group_t) Nw_parameters_t
	_nw_connection_group_copy_path_for_message func(Nw_connection_group_t, Nw_content_context_t) Nw_path_t
	_nw_connection_group_copy_protocol_metadata func(Nw_connection_group_t, Nw_protocol_definition_t) Nw_protocol_metadata_t
	_nw_connection_group_copy_protocol_metadata_for_message func(Nw_connection_group_t, Nw_content_context_t, Nw_protocol_definition_t) Nw_protocol_metadata_t
	_nw_connection_group_copy_remote_endpoint_for_message func(Nw_connection_group_t, Nw_content_context_t) Nw_endpoint_t
	_nw_connection_group_create func(Nw_group_descriptor_t, Nw_parameters_t) Nw_connection_group_t
	_nw_connection_group_extract_connection func(Nw_connection_group_t, Nw_endpoint_t, Nw_protocol_options_t) Nw_connection_t
	_nw_connection_group_extract_connection_for_message func(Nw_connection_group_t, Nw_content_context_t) Nw_connection_t
	_nw_connection_group_reinsert_extracted_connection func(Nw_connection_group_t, Nw_connection_t) bool
	_nw_connection_group_reply func(Nw_connection_group_t, Nw_content_context_t, Nw_content_context_t, unsafe.Pointer)
	_nw_connection_group_send_message func(Nw_connection_group_t, unsafe.Pointer, Nw_endpoint_t, Nw_content_context_t, unsafe.Pointer)
	_nw_connection_group_set_new_connection_handler func(Nw_connection_group_t, unsafe.Pointer)
	_nw_connection_group_set_queue func(Nw_connection_group_t, unsafe.Pointer)
	_nw_connection_group_set_receive_handler func(Nw_connection_group_t, uint32, bool, unsafe.Pointer)
	_nw_connection_group_set_state_changed_handler func(Nw_connection_group_t, unsafe.Pointer)
	_nw_connection_group_start func(Nw_connection_group_t)
	_nw_connection_receive func(Nw_connection_t, uint32, uint32, unsafe.Pointer)
	_nw_connection_receive_message func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_restart func(Nw_connection_t)
	_nw_connection_send func(Nw_connection_t, unsafe.Pointer, Nw_content_context_t, bool, unsafe.Pointer)
	_nw_connection_set_better_path_available_handler func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_set_path_changed_handler func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_set_queue func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_set_state_changed_handler func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_set_viability_changed_handler func(Nw_connection_t, unsafe.Pointer)
	_nw_connection_start func(Nw_connection_t)
	_nw_content_context_copy_antecedent func(Nw_content_context_t) Nw_content_context_t
	_nw_content_context_copy_protocol_metadata func(Nw_content_context_t, Nw_protocol_definition_t) Nw_protocol_metadata_t
	_nw_content_context_create func(unsafe.Pointer) Nw_content_context_t
	_nw_content_context_foreach_protocol_metadata func(Nw_content_context_t)
	_nw_content_context_get_expiration_milliseconds func(Nw_content_context_t) uint64
	_nw_content_context_get_identifier func(Nw_content_context_t) unsafe.Pointer
	_nw_content_context_get_is_final func(Nw_content_context_t) bool
	_nw_content_context_get_relative_priority func(Nw_content_context_t) float64
	_nw_content_context_set_antecedent func(Nw_content_context_t, Nw_content_context_t)
	_nw_content_context_set_expiration_milliseconds func(Nw_content_context_t, uint64)
	_nw_content_context_set_is_final func(Nw_content_context_t, bool)
	_nw_content_context_set_metadata_for_protocol func(Nw_content_context_t, Nw_protocol_metadata_t)
	_nw_content_context_set_relative_priority func(Nw_content_context_t, float64)
	_nw_data_transfer_report_collect func(Nw_data_transfer_report_t, unsafe.Pointer, unsafe.Pointer)
	_nw_data_transfer_report_copy_path_interface func(Nw_data_transfer_report_t, uint32) Nw_interface_t
	_nw_data_transfer_report_get_duration_milliseconds func(Nw_data_transfer_report_t) uint64
	_nw_data_transfer_report_get_path_count func(Nw_data_transfer_report_t) uint32
	_nw_data_transfer_report_get_path_radio_type func(Nw_data_transfer_report_t, uint32) unsafe.Pointer
	_nw_data_transfer_report_get_received_application_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_received_ip_packet_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_received_transport_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_received_transport_duplicate_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_received_transport_out_of_order_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_sent_application_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_sent_ip_packet_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_sent_transport_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_state func(Nw_data_transfer_report_t) unsafe.Pointer
	_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_transport_rtt_variance func(Nw_data_transfer_report_t, uint32) uint64
	_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds func(Nw_data_transfer_report_t, uint32) uint64
	_nw_endpoint_copy_address_string func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_copy_port_string func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_copy_txt_record func(Nw_endpoint_t) Nw_txt_record_t
	_nw_endpoint_create_address func(unsafe.Pointer) Nw_endpoint_t
	_nw_endpoint_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Nw_endpoint_t
	_nw_endpoint_create_host func(unsafe.Pointer, unsafe.Pointer) Nw_endpoint_t
	_nw_endpoint_create_url func(unsafe.Pointer) Nw_endpoint_t
	_nw_endpoint_get_address func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_domain func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_name func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_type func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_hostname func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_port func(Nw_endpoint_t) uint16
	_nw_endpoint_get_signature func(Nw_endpoint_t, unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_type func(Nw_endpoint_t) unsafe.Pointer
	_nw_endpoint_get_url func(Nw_endpoint_t) unsafe.Pointer
	_nw_error_copy_cf_error func(Nw_error_t) ErrorRef
	_nw_error_get_error_code func(Nw_error_t) int
	_nw_error_get_error_domain func(Nw_error_t) unsafe.Pointer
	_nw_establishment_report_copy_proxy_endpoint func(Nw_establishment_report_t) Nw_endpoint_t
	_nw_establishment_report_enumerate_protocols func(Nw_establishment_report_t, unsafe.Pointer)
	_nw_establishment_report_enumerate_resolution_reports func(Nw_establishment_report_t, unsafe.Pointer)
	_nw_establishment_report_enumerate_resolutions func(Nw_establishment_report_t, unsafe.Pointer)
	_nw_establishment_report_get_attempt_started_after_milliseconds func(Nw_establishment_report_t) uint64
	_nw_establishment_report_get_duration_milliseconds func(Nw_establishment_report_t) uint64
	_nw_establishment_report_get_previous_attempt_count func(Nw_establishment_report_t) uint32
	_nw_establishment_report_get_proxy_configured func(Nw_establishment_report_t) bool
	_nw_establishment_report_get_used_proxy func(Nw_establishment_report_t) bool
	_nw_ethernet_channel_cancel func(Nw_ethernet_channel_t)
	_nw_ethernet_channel_create func(uint16, Nw_interface_t) Nw_ethernet_channel_t
	_nw_ethernet_channel_create_with_parameters func(uint16, Nw_interface_t, Nw_parameters_t) Nw_ethernet_channel_t
	_nw_ethernet_channel_get_maximum_payload_size func(Nw_ethernet_channel_t) uint32
	_nw_ethernet_channel_send func(Nw_ethernet_channel_t, unsafe.Pointer, uint16, Nw_ethernet_address_t, unsafe.Pointer)
	_nw_ethernet_channel_set_queue func(Nw_ethernet_channel_t, unsafe.Pointer)
	_nw_ethernet_channel_set_receive_handler func(Nw_ethernet_channel_t, unsafe.Pointer)
	_nw_ethernet_channel_set_state_changed_handler func(Nw_ethernet_channel_t, unsafe.Pointer)
	_nw_ethernet_channel_start func(Nw_ethernet_channel_t)
	_nw_framer_async func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_copy_options func(Nw_framer_t) Nw_protocol_options_t
	_nw_framer_copy_parameters func(Nw_framer_t) Nw_parameters_t
	_nw_framer_copy_remote_endpoint func(Nw_framer_t) Nw_endpoint_t
	_nw_framer_create_definition func(unsafe.Pointer, uint32, unsafe.Pointer) Nw_protocol_definition_t
	_nw_framer_create_options func(Nw_protocol_definition_t) Nw_protocol_options_t
	_nw_framer_deliver_input func(Nw_framer_t, unsafe.Pointer, uintptr, Nw_framer_message_t, bool)
	_nw_framer_deliver_input_no_copy func(Nw_framer_t, uintptr, Nw_framer_message_t, bool) bool
	_nw_framer_mark_failed_with_error func(Nw_framer_t, int)
	_nw_framer_mark_ready func(Nw_framer_t)
	_nw_framer_message_access_value func(Nw_framer_message_t, unsafe.Pointer, bool) bool
	_nw_framer_message_copy_object_value func(Nw_framer_message_t, unsafe.Pointer) objc.ID
	_nw_framer_message_create func(Nw_framer_t) Nw_framer_message_t
	_nw_framer_message_set_object_value func(Nw_framer_message_t, unsafe.Pointer, objc.ID)
	_nw_framer_message_set_value func(Nw_framer_message_t, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_nw_framer_options_copy_object_value func(Nw_protocol_options_t, unsafe.Pointer) objc.ID
	_nw_framer_options_set_object_value func(Nw_protocol_options_t, unsafe.Pointer, objc.ID)
	_nw_framer_parse_input func(Nw_framer_t, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) bool
	_nw_framer_parse_output func(Nw_framer_t, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) bool
	_nw_framer_pass_through_input func(Nw_framer_t)
	_nw_framer_pass_through_output func(Nw_framer_t)
	_nw_framer_prepend_application_protocol func(Nw_framer_t, Nw_protocol_options_t) bool
	_nw_framer_protocol_create_message func(Nw_protocol_definition_t) Nw_framer_message_t
	_nw_framer_schedule_wakeup func(Nw_framer_t, uint64)
	_nw_framer_set_cleanup_handler func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_set_input_handler func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_set_output_handler func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_set_stop_handler func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_set_wakeup_handler func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_write_output func(Nw_framer_t, unsafe.Pointer, uintptr)
	_nw_framer_write_output_data func(Nw_framer_t, unsafe.Pointer)
	_nw_framer_write_output_no_copy func(Nw_framer_t, uintptr) bool
	_nw_group_descriptor_add_endpoint func(Nw_group_descriptor_t, Nw_endpoint_t) bool
	_nw_group_descriptor_create_multicast func(Nw_endpoint_t) Nw_group_descriptor_t
	_nw_group_descriptor_create_multiplex func(Nw_endpoint_t) Nw_group_descriptor_t
	_nw_group_descriptor_enumerate_endpoints func(Nw_group_descriptor_t, unsafe.Pointer)
	_nw_interface_get_index func(Nw_interface_t) uint32
	_nw_interface_get_name func(Nw_interface_t) unsafe.Pointer
	_nw_interface_get_type func(Nw_interface_t) unsafe.Pointer
	_nw_ip_create_metadata func() Nw_protocol_metadata_t
	_nw_ip_metadata_get_ecn_flag func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_ip_metadata_get_receive_time func(Nw_protocol_metadata_t) uint64
	_nw_ip_metadata_get_service_class func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_ip_metadata_set_ecn_flag func(Nw_protocol_metadata_t, unsafe.Pointer)
	_nw_ip_metadata_set_service_class func(Nw_protocol_metadata_t, unsafe.Pointer)
	_nw_ip_options_set_calculate_receive_time func(Nw_protocol_options_t, bool)
	_nw_ip_options_set_disable_fragmentation func(Nw_protocol_options_t, bool)
	_nw_ip_options_set_disable_multicast_loopback func(Nw_protocol_options_t, bool)
	_nw_ip_options_set_hop_limit func(Nw_protocol_options_t, uint8)
	_nw_ip_options_set_use_minimum_mtu func(Nw_protocol_options_t, bool)
	_nw_listener_cancel func(Nw_listener_t)
	_nw_listener_create func(Nw_parameters_t) Nw_listener_t
	_nw_listener_create_with_connection func(Nw_connection_t, Nw_parameters_t) Nw_listener_t
	_nw_listener_create_with_launchd_key func(Nw_parameters_t, unsafe.Pointer) Nw_listener_t
	_nw_listener_create_with_port func(unsafe.Pointer, Nw_parameters_t) Nw_listener_t
	_nw_listener_get_new_connection_limit func(Nw_listener_t) uint32
	_nw_listener_get_port func(Nw_listener_t) uint16
	_nw_listener_set_advertise_descriptor func(Nw_listener_t, Nw_advertise_descriptor_t)
	_nw_listener_set_advertised_endpoint_changed_handler func(Nw_listener_t, unsafe.Pointer)
	_nw_listener_set_new_connection_group_handler func(Nw_listener_t, unsafe.Pointer)
	_nw_listener_set_new_connection_handler func(Nw_listener_t, unsafe.Pointer)
	_nw_listener_set_new_connection_limit func(Nw_listener_t, uint32)
	_nw_listener_set_queue func(Nw_listener_t, unsafe.Pointer)
	_nw_listener_set_state_changed_handler func(Nw_listener_t, unsafe.Pointer)
	_nw_listener_start func(Nw_listener_t)
	_nw_multicast_group_descriptor_get_disable_unicast_traffic func(Nw_group_descriptor_t) bool
	_nw_multicast_group_descriptor_set_disable_unicast_traffic func(Nw_group_descriptor_t, bool)
	_nw_multicast_group_descriptor_set_specific_source func(Nw_group_descriptor_t, Nw_endpoint_t)
	_nw_parameters_clear_prohibited_interface_types func(Nw_parameters_t)
	_nw_parameters_clear_prohibited_interfaces func(Nw_parameters_t)
	_nw_parameters_copy func(Nw_parameters_t) Nw_parameters_t
	_nw_parameters_copy_default_protocol_stack func(Nw_parameters_t) Nw_protocol_stack_t
	_nw_parameters_copy_required_interface func(Nw_parameters_t) Nw_interface_t
	_nw_parameters_create func() Nw_parameters_t
	_nw_parameters_create_application_service func() Nw_parameters_t
	_nw_parameters_create_custom_ip func(uint8, unsafe.Pointer) Nw_parameters_t
	_nw_parameters_create_quic func(unsafe.Pointer) Nw_parameters_t
	_nw_parameters_create_secure_tcp func(unsafe.Pointer, unsafe.Pointer) Nw_parameters_t
	_nw_parameters_create_secure_udp func(unsafe.Pointer, unsafe.Pointer) Nw_parameters_t
	_nw_parameters_get_allow_ultra_constrained func(Nw_parameters_t) bool
	_nw_parameters_get_attribution func(Nw_parameters_t) nw_parameters_attribution_t
	_nw_parameters_get_expired_dns_behavior func(Nw_parameters_t) unsafe.Pointer
	_nw_parameters_get_fast_open_enabled func(Nw_parameters_t) bool
	_nw_parameters_get_include_peer_to_peer func(Nw_parameters_t) bool
	_nw_parameters_get_multipath_service func(Nw_parameters_t) unsafe.Pointer
	_nw_parameters_get_prefer_no_proxy func(Nw_parameters_t) bool
	_nw_parameters_get_prohibit_constrained func(Nw_parameters_t) bool
	_nw_parameters_get_prohibit_expensive func(Nw_parameters_t) bool
	_nw_parameters_get_required_interface_type func(Nw_parameters_t) unsafe.Pointer
	_nw_parameters_get_service_class func(Nw_parameters_t) unsafe.Pointer
	_nw_parameters_iterate_prohibited_interface_types func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_iterate_prohibited_interfaces func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_prohibit_interface func(Nw_parameters_t, Nw_interface_t)
	_nw_parameters_prohibit_interface_type func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_require_interface func(Nw_parameters_t, Nw_interface_t)
	_nw_parameters_requires_dnssec_validation func(Nw_parameters_t) bool
	_nw_parameters_set_allow_ultra_constrained func(Nw_parameters_t, bool)
	_nw_parameters_set_attribution func(Nw_parameters_t, nw_parameters_attribution_t)
	_nw_parameters_set_expired_dns_behavior func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_set_fast_open_enabled func(Nw_parameters_t, bool)
	_nw_parameters_set_include_peer_to_peer func(Nw_parameters_t, bool)
	_nw_parameters_set_multipath_service func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_set_prefer_no_proxy func(Nw_parameters_t, bool)
	_nw_parameters_set_privacy_context func(Nw_parameters_t, Nw_privacy_context_t)
	_nw_parameters_set_prohibit_constrained func(Nw_parameters_t, bool)
	_nw_parameters_set_prohibit_expensive func(Nw_parameters_t, bool)
	_nw_parameters_set_required_interface_type func(Nw_parameters_t, unsafe.Pointer)
	_nw_parameters_set_requires_dnssec_validation func(Nw_parameters_t, bool)
	_nw_parameters_set_service_class func(Nw_parameters_t, unsafe.Pointer)
	_nw_path_copy_effective_remote_endpoint func(Nw_path_t) Nw_endpoint_t
	_nw_path_enumerate_gateways func(Nw_path_t, unsafe.Pointer)
	_nw_path_enumerate_interfaces func(Nw_path_t, unsafe.Pointer)
	_nw_path_get_link_quality func(Nw_path_t) unsafe.Pointer
	_nw_path_get_status func(Nw_path_t) unsafe.Pointer
	_nw_path_get_unsatisfied_reason func(Nw_path_t) unsafe.Pointer
	_nw_path_has_dns func(Nw_path_t) bool
	_nw_path_has_ipv4 func(Nw_path_t) bool
	_nw_path_has_ipv6 func(Nw_path_t) bool
	_nw_path_is_constrained func(Nw_path_t) bool
	_nw_path_is_equal func(Nw_path_t, Nw_path_t) bool
	_nw_path_is_expensive func(Nw_path_t) bool
	_nw_path_is_ultra_constrained func(Nw_path_t) bool
	_nw_path_monitor_cancel func(Nw_path_monitor_t)
	_nw_path_monitor_create func() Nw_path_monitor_t
	_nw_path_monitor_create_for_ethernet_channel func() Nw_path_monitor_t
	_nw_path_monitor_create_with_type func(unsafe.Pointer) Nw_path_monitor_t
	_nw_path_monitor_prohibit_interface_type func(Nw_path_monitor_t, unsafe.Pointer)
	_nw_path_monitor_set_cancel_handler func(Nw_path_monitor_t, unsafe.Pointer)
	_nw_path_monitor_set_queue func(Nw_path_monitor_t, unsafe.Pointer)
	_nw_path_monitor_set_update_handler func(Nw_path_monitor_t, unsafe.Pointer)
	_nw_path_monitor_start func(Nw_path_monitor_t)
	_nw_path_uses_interface_type func(Nw_path_t, unsafe.Pointer) bool
	_nw_privacy_context_add_proxy func(Nw_privacy_context_t, Nw_proxy_config_t)
	_nw_privacy_context_clear_proxies func(Nw_privacy_context_t)
	_nw_privacy_context_create func(unsafe.Pointer) Nw_privacy_context_t
	_nw_privacy_context_disable_logging func(Nw_privacy_context_t)
	_nw_privacy_context_flush_cache func(Nw_privacy_context_t)
	_nw_privacy_context_require_encrypted_name_resolution func(Nw_privacy_context_t, bool, Nw_resolver_config_t)
	_nw_protocol_copy_ip_definition func() Nw_protocol_definition_t
	_nw_protocol_copy_quic_definition func() Nw_protocol_definition_t
	_nw_protocol_copy_tcp_definition func() Nw_protocol_definition_t
	_nw_protocol_copy_tls_definition func() Nw_protocol_definition_t
	_nw_protocol_copy_udp_definition func() Nw_protocol_definition_t
	_nw_protocol_copy_ws_definition func() Nw_protocol_definition_t
	_nw_protocol_definition_is_equal func(Nw_protocol_definition_t, Nw_protocol_definition_t) bool
	_nw_protocol_metadata_copy_definition func(Nw_protocol_metadata_t) Nw_protocol_definition_t
	_nw_protocol_metadata_is_framer_message func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_ip func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_quic func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_tcp func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_tls func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_udp func(Nw_protocol_metadata_t) bool
	_nw_protocol_metadata_is_ws func(Nw_protocol_metadata_t) bool
	_nw_protocol_options_copy_definition func(Nw_protocol_options_t) Nw_protocol_definition_t
	_nw_protocol_options_is_quic func(Nw_protocol_options_t) bool
	_nw_protocol_stack_clear_application_protocols func(Nw_protocol_stack_t)
	_nw_protocol_stack_copy_internet_protocol func(Nw_protocol_stack_t) Nw_protocol_options_t
	_nw_protocol_stack_copy_transport_protocol func(Nw_protocol_stack_t) Nw_protocol_options_t
	_nw_protocol_stack_iterate_application_protocols func(Nw_protocol_stack_t, unsafe.Pointer)
	_nw_protocol_stack_prepend_application_protocol func(Nw_protocol_stack_t, Nw_protocol_options_t)
	_nw_protocol_stack_set_transport_protocol func(Nw_protocol_stack_t, Nw_protocol_options_t)
	_nw_proxy_config_add_excluded_domain func(Nw_proxy_config_t, unsafe.Pointer)
	_nw_proxy_config_add_match_domain func(Nw_proxy_config_t, unsafe.Pointer)
	_nw_proxy_config_clear_excluded_domains func(Nw_proxy_config_t)
	_nw_proxy_config_clear_match_domains func(Nw_proxy_config_t)
	_nw_proxy_config_create_http_connect func(Nw_endpoint_t, Nw_protocol_options_t) Nw_proxy_config_t
	_nw_proxy_config_create_oblivious_http func(Nw_relay_hop_t, unsafe.Pointer, unsafe.Pointer, uintptr) Nw_proxy_config_t
	_nw_proxy_config_create_relay func(Nw_relay_hop_t, Nw_relay_hop_t) Nw_proxy_config_t
	_nw_proxy_config_create_socksv5 func(Nw_endpoint_t) Nw_proxy_config_t
	_nw_proxy_config_enumerate_excluded_domains func(Nw_proxy_config_t, unsafe.Pointer)
	_nw_proxy_config_enumerate_match_domains func(Nw_proxy_config_t, unsafe.Pointer)
	_nw_proxy_config_get_failover_allowed func(Nw_proxy_config_t) bool
	_nw_proxy_config_set_failover_allowed func(Nw_proxy_config_t, bool)
	_nw_proxy_config_set_username_and_password func(Nw_proxy_config_t, unsafe.Pointer, unsafe.Pointer)
	_nw_quic_add_tls_application_protocol func(Nw_protocol_options_t, unsafe.Pointer)
	_nw_quic_copy_sec_protocol_metadata func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_quic_copy_sec_protocol_options func(Nw_protocol_options_t) unsafe.Pointer
	_nw_quic_create_options func() Nw_protocol_options_t
	_nw_quic_get_application_error func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_application_error_reason func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_quic_get_idle_timeout func(Nw_protocol_options_t) uint32
	_nw_quic_get_initial_max_data func(Nw_protocol_options_t) uint64
	_nw_quic_get_initial_max_stream_data_bidirectional_remote func(Nw_protocol_options_t) uint64
	_nw_quic_get_initial_max_stream_data_unidirectional func(Nw_protocol_options_t) uint64
	_nw_quic_get_initial_max_streams_bidirectional func(Nw_protocol_options_t) uint64
	_nw_quic_get_initial_max_streams_unidirectional func(Nw_protocol_options_t) uint64
	_nw_quic_get_keepalive_interval func(Nw_protocol_metadata_t) uint16
	_nw_quic_get_max_datagram_frame_size func(Nw_protocol_options_t) uint16
	_nw_quic_get_max_udp_payload_size func(Nw_protocol_options_t) uint16
	_nw_quic_get_remote_idle_timeout func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_remote_max_streams_bidirectional func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_remote_max_streams_unidirectional func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_stream_application_error func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_stream_id func(Nw_protocol_metadata_t) uint64
	_nw_quic_get_stream_is_datagram func(Nw_protocol_options_t) bool
	_nw_quic_get_stream_is_unidirectional func(Nw_protocol_options_t) bool
	_nw_quic_get_stream_type func(Nw_protocol_metadata_t) uint8
	_nw_quic_get_stream_usable_datagram_frame_size func(Nw_protocol_metadata_t) uint16
	_nw_quic_set_application_error func(Nw_protocol_metadata_t, uint64, unsafe.Pointer)
	_nw_quic_set_idle_timeout func(Nw_protocol_options_t, uint32)
	_nw_quic_set_initial_max_data func(Nw_protocol_options_t, uint64)
	_nw_quic_set_initial_max_stream_data_bidirectional_remote func(Nw_protocol_options_t, uint64)
	_nw_quic_set_initial_max_stream_data_unidirectional func(Nw_protocol_options_t, uint64)
	_nw_quic_set_initial_max_streams_bidirectional func(Nw_protocol_options_t, uint64)
	_nw_quic_set_initial_max_streams_unidirectional func(Nw_protocol_options_t, uint64)
	_nw_quic_set_keepalive_interval func(Nw_protocol_metadata_t, uint16)
	_nw_quic_set_max_datagram_frame_size func(Nw_protocol_options_t, uint16)
	_nw_quic_set_max_udp_payload_size func(Nw_protocol_options_t, uint16)
	_nw_quic_set_stream_application_error func(Nw_protocol_metadata_t, uint64)
	_nw_quic_set_stream_is_datagram func(Nw_protocol_options_t, bool)
	_nw_quic_set_stream_is_unidirectional func(Nw_protocol_options_t, bool)
	_nw_relay_hop_add_additional_http_header_field func(Nw_relay_hop_t, unsafe.Pointer, unsafe.Pointer)
	_nw_relay_hop_create func(Nw_endpoint_t, Nw_endpoint_t, Nw_protocol_options_t) Nw_relay_hop_t
	_nw_release func(unsafe.Pointer)
	_nw_resolution_report_copy_preferred_endpoint func(Nw_resolution_report_t) Nw_endpoint_t
	_nw_resolution_report_copy_successful_endpoint func(Nw_resolution_report_t) Nw_endpoint_t
	_nw_resolution_report_get_endpoint_count func(Nw_resolution_report_t) uint32
	_nw_resolution_report_get_milliseconds func(Nw_resolution_report_t) uint64
	_nw_resolution_report_get_protocol func(Nw_resolution_report_t) unsafe.Pointer
	_nw_resolution_report_get_source func(Nw_resolution_report_t) unsafe.Pointer
	_nw_resolver_config_add_server_address func(Nw_resolver_config_t, Nw_endpoint_t)
	_nw_resolver_config_create_https func(Nw_endpoint_t) Nw_resolver_config_t
	_nw_resolver_config_create_tls func(Nw_endpoint_t) Nw_resolver_config_t
	_nw_retain func(unsafe.Pointer) unsafe.Pointer
	_nw_tcp_create_options func() Nw_protocol_options_t
	_nw_tcp_get_available_receive_buffer func(Nw_protocol_metadata_t) uint32
	_nw_tcp_get_available_send_buffer func(Nw_protocol_metadata_t) uint32
	_nw_tcp_options_set_connection_timeout func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_disable_ack_stretching func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_disable_ecn func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_enable_fast_open func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_enable_keepalive func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_keepalive_count func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_keepalive_idle_time func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_keepalive_interval func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_maximum_segment_size func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_no_delay func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_no_options func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_no_push func(Nw_protocol_options_t, bool)
	_nw_tcp_options_set_persist_timeout func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_retransmit_connection_drop_time func(Nw_protocol_options_t, uint32)
	_nw_tcp_options_set_retransmit_fin_drop func(Nw_protocol_options_t, bool)
	_nw_tls_copy_sec_protocol_metadata func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_tls_copy_sec_protocol_options func(Nw_protocol_options_t) unsafe.Pointer
	_nw_tls_create_options func() Nw_protocol_options_t
	_nw_txt_record_access_bytes func(Nw_txt_record_t, unsafe.Pointer) bool
	_nw_txt_record_access_key func(Nw_txt_record_t, unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_apply func(Nw_txt_record_t, unsafe.Pointer) bool
	_nw_txt_record_copy func(Nw_txt_record_t) Nw_txt_record_t
	_nw_txt_record_create_dictionary func() Nw_txt_record_t
	_nw_txt_record_create_with_bytes func(unsafe.Pointer, uintptr) Nw_txt_record_t
	_nw_txt_record_find_key func(Nw_txt_record_t, unsafe.Pointer) unsafe.Pointer
	_nw_txt_record_get_key_count func(Nw_txt_record_t) uintptr
	_nw_txt_record_is_dictionary func(Nw_txt_record_t) bool
	_nw_txt_record_is_equal func(Nw_txt_record_t, Nw_txt_record_t) bool
	_nw_txt_record_remove_key func(Nw_txt_record_t, unsafe.Pointer) bool
	_nw_txt_record_set_key func(Nw_txt_record_t, unsafe.Pointer, unsafe.Pointer, uintptr) bool
	_nw_udp_create_metadata func() Nw_protocol_metadata_t
	_nw_udp_create_options func() Nw_protocol_options_t
	_nw_udp_options_set_prefer_no_checksum func(Nw_protocol_options_t, bool)
	_nw_ws_create_metadata func(unsafe.Pointer) Nw_protocol_metadata_t
	_nw_ws_create_options func(unsafe.Pointer) Nw_protocol_options_t
	_nw_ws_metadata_copy_server_response func(Nw_protocol_metadata_t) Nw_ws_response_t
	_nw_ws_metadata_get_close_code func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_ws_metadata_get_opcode func(Nw_protocol_metadata_t) unsafe.Pointer
	_nw_ws_metadata_set_close_code func(Nw_protocol_metadata_t, unsafe.Pointer)
	_nw_ws_metadata_set_pong_handler func(Nw_protocol_metadata_t, unsafe.Pointer, unsafe.Pointer)
	_nw_ws_options_add_additional_header func(Nw_protocol_options_t, unsafe.Pointer, unsafe.Pointer)
	_nw_ws_options_add_subprotocol func(Nw_protocol_options_t, unsafe.Pointer)
	_nw_ws_options_set_auto_reply_ping func(Nw_protocol_options_t, bool)
	_nw_ws_options_set_client_request_handler func(Nw_protocol_options_t, unsafe.Pointer, unsafe.Pointer)
	_nw_ws_options_set_maximum_message_size func(Nw_protocol_options_t, uintptr)
	_nw_ws_options_set_skip_handshake func(Nw_protocol_options_t, bool)
	_nw_ws_request_enumerate_additional_headers func(Nw_ws_request_t, unsafe.Pointer) bool
	_nw_ws_request_enumerate_subprotocols func(Nw_ws_request_t, unsafe.Pointer) bool
	_nw_ws_response_add_additional_header func(Nw_ws_response_t, unsafe.Pointer, unsafe.Pointer)
	_nw_ws_response_create func(unsafe.Pointer, unsafe.Pointer) Nw_ws_response_t
	_nw_ws_response_enumerate_additional_headers func(Nw_ws_response_t, unsafe.Pointer) bool
	_nw_ws_response_get_selected_subprotocol func(Nw_ws_response_t) unsafe.Pointer
	_nw_ws_response_get_status func(Nw_ws_response_t) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_nw_advertise_descriptor_copy_txt_record_object, lib, "nw_advertise_descriptor_copy_txt_record_object")
	tryRegister(&_nw_advertise_descriptor_create_application_service, lib, "nw_advertise_descriptor_create_application_service")
	tryRegister(&_nw_advertise_descriptor_create_bonjour_service, lib, "nw_advertise_descriptor_create_bonjour_service")
	tryRegister(&_nw_advertise_descriptor_get_application_service_name, lib, "nw_advertise_descriptor_get_application_service_name")
	tryRegister(&_nw_advertise_descriptor_get_no_auto_rename, lib, "nw_advertise_descriptor_get_no_auto_rename")
	tryRegister(&_nw_advertise_descriptor_set_no_auto_rename, lib, "nw_advertise_descriptor_set_no_auto_rename")
	tryRegister(&_nw_advertise_descriptor_set_txt_record, lib, "nw_advertise_descriptor_set_txt_record")
	tryRegister(&_nw_advertise_descriptor_set_txt_record_object, lib, "nw_advertise_descriptor_set_txt_record_object")
	tryRegister(&_nw_browse_descriptor_create_application_service, lib, "nw_browse_descriptor_create_application_service")
	tryRegister(&_nw_browse_descriptor_create_bonjour_service, lib, "nw_browse_descriptor_create_bonjour_service")
	tryRegister(&_nw_browse_descriptor_get_application_service_name, lib, "nw_browse_descriptor_get_application_service_name")
	tryRegister(&_nw_browse_descriptor_get_bonjour_service_domain, lib, "nw_browse_descriptor_get_bonjour_service_domain")
	tryRegister(&_nw_browse_descriptor_get_bonjour_service_type, lib, "nw_browse_descriptor_get_bonjour_service_type")
	tryRegister(&_nw_browse_descriptor_get_include_txt_record, lib, "nw_browse_descriptor_get_include_txt_record")
	tryRegister(&_nw_browse_descriptor_set_include_txt_record, lib, "nw_browse_descriptor_set_include_txt_record")
	tryRegister(&_nw_browse_result_copy_endpoint, lib, "nw_browse_result_copy_endpoint")
	tryRegister(&_nw_browse_result_copy_txt_record_object, lib, "nw_browse_result_copy_txt_record_object")
	tryRegister(&_nw_browse_result_enumerate_interfaces, lib, "nw_browse_result_enumerate_interfaces")
	tryRegister(&_nw_browse_result_get_changes, lib, "nw_browse_result_get_changes")
	tryRegister(&_nw_browse_result_get_interfaces_count, lib, "nw_browse_result_get_interfaces_count")
	tryRegister(&_nw_browser_cancel, lib, "nw_browser_cancel")
	tryRegister(&_nw_browser_copy_browse_descriptor, lib, "nw_browser_copy_browse_descriptor")
	tryRegister(&_nw_browser_copy_parameters, lib, "nw_browser_copy_parameters")
	tryRegister(&_nw_browser_create, lib, "nw_browser_create")
	tryRegister(&_nw_browser_set_browse_results_changed_handler, lib, "nw_browser_set_browse_results_changed_handler")
	tryRegister(&_nw_browser_set_queue, lib, "nw_browser_set_queue")
	tryRegister(&_nw_browser_set_state_changed_handler, lib, "nw_browser_set_state_changed_handler")
	tryRegister(&_nw_browser_start, lib, "nw_browser_start")
	tryRegister(&_nw_connection_access_establishment_report, lib, "nw_connection_access_establishment_report")
	tryRegister(&_nw_connection_batch, lib, "nw_connection_batch")
	tryRegister(&_nw_connection_cancel, lib, "nw_connection_cancel")
	tryRegister(&_nw_connection_cancel_current_endpoint, lib, "nw_connection_cancel_current_endpoint")
	tryRegister(&_nw_connection_copy_current_path, lib, "nw_connection_copy_current_path")
	tryRegister(&_nw_connection_copy_description, lib, "nw_connection_copy_description")
	tryRegister(&_nw_connection_copy_endpoint, lib, "nw_connection_copy_endpoint")
	tryRegister(&_nw_connection_copy_parameters, lib, "nw_connection_copy_parameters")
	tryRegister(&_nw_connection_copy_protocol_metadata, lib, "nw_connection_copy_protocol_metadata")
	tryRegister(&_nw_connection_create, lib, "nw_connection_create")
	tryRegister(&_nw_connection_create_new_data_transfer_report, lib, "nw_connection_create_new_data_transfer_report")
	tryRegister(&_nw_connection_force_cancel, lib, "nw_connection_force_cancel")
	tryRegister(&_nw_connection_get_maximum_datagram_size, lib, "nw_connection_get_maximum_datagram_size")
	tryRegister(&_nw_connection_group_cancel, lib, "nw_connection_group_cancel")
	tryRegister(&_nw_connection_group_copy_descriptor, lib, "nw_connection_group_copy_descriptor")
	tryRegister(&_nw_connection_group_copy_parameters, lib, "nw_connection_group_copy_parameters")
	tryRegister(&_nw_connection_group_copy_path_for_message, lib, "nw_connection_group_copy_path_for_message")
	tryRegister(&_nw_connection_group_copy_protocol_metadata, lib, "nw_connection_group_copy_protocol_metadata")
	tryRegister(&_nw_connection_group_copy_protocol_metadata_for_message, lib, "nw_connection_group_copy_protocol_metadata_for_message")
	tryRegister(&_nw_connection_group_copy_remote_endpoint_for_message, lib, "nw_connection_group_copy_remote_endpoint_for_message")
	tryRegister(&_nw_connection_group_create, lib, "nw_connection_group_create")
	tryRegister(&_nw_connection_group_extract_connection, lib, "nw_connection_group_extract_connection")
	tryRegister(&_nw_connection_group_extract_connection_for_message, lib, "nw_connection_group_extract_connection_for_message")
	tryRegister(&_nw_connection_group_reinsert_extracted_connection, lib, "nw_connection_group_reinsert_extracted_connection")
	tryRegister(&_nw_connection_group_reply, lib, "nw_connection_group_reply")
	tryRegister(&_nw_connection_group_send_message, lib, "nw_connection_group_send_message")
	tryRegister(&_nw_connection_group_set_new_connection_handler, lib, "nw_connection_group_set_new_connection_handler")
	tryRegister(&_nw_connection_group_set_queue, lib, "nw_connection_group_set_queue")
	tryRegister(&_nw_connection_group_set_receive_handler, lib, "nw_connection_group_set_receive_handler")
	tryRegister(&_nw_connection_group_set_state_changed_handler, lib, "nw_connection_group_set_state_changed_handler")
	tryRegister(&_nw_connection_group_start, lib, "nw_connection_group_start")
	tryRegister(&_nw_connection_receive, lib, "nw_connection_receive")
	tryRegister(&_nw_connection_receive_message, lib, "nw_connection_receive_message")
	tryRegister(&_nw_connection_restart, lib, "nw_connection_restart")
	tryRegister(&_nw_connection_send, lib, "nw_connection_send")
	tryRegister(&_nw_connection_set_better_path_available_handler, lib, "nw_connection_set_better_path_available_handler")
	tryRegister(&_nw_connection_set_path_changed_handler, lib, "nw_connection_set_path_changed_handler")
	tryRegister(&_nw_connection_set_queue, lib, "nw_connection_set_queue")
	tryRegister(&_nw_connection_set_state_changed_handler, lib, "nw_connection_set_state_changed_handler")
	tryRegister(&_nw_connection_set_viability_changed_handler, lib, "nw_connection_set_viability_changed_handler")
	tryRegister(&_nw_connection_start, lib, "nw_connection_start")
	tryRegister(&_nw_content_context_copy_antecedent, lib, "nw_content_context_copy_antecedent")
	tryRegister(&_nw_content_context_copy_protocol_metadata, lib, "nw_content_context_copy_protocol_metadata")
	tryRegister(&_nw_content_context_create, lib, "nw_content_context_create")
	tryRegister(&_nw_content_context_foreach_protocol_metadata, lib, "nw_content_context_foreach_protocol_metadata")
	tryRegister(&_nw_content_context_get_expiration_milliseconds, lib, "nw_content_context_get_expiration_milliseconds")
	tryRegister(&_nw_content_context_get_identifier, lib, "nw_content_context_get_identifier")
	tryRegister(&_nw_content_context_get_is_final, lib, "nw_content_context_get_is_final")
	tryRegister(&_nw_content_context_get_relative_priority, lib, "nw_content_context_get_relative_priority")
	tryRegister(&_nw_content_context_set_antecedent, lib, "nw_content_context_set_antecedent")
	tryRegister(&_nw_content_context_set_expiration_milliseconds, lib, "nw_content_context_set_expiration_milliseconds")
	tryRegister(&_nw_content_context_set_is_final, lib, "nw_content_context_set_is_final")
	tryRegister(&_nw_content_context_set_metadata_for_protocol, lib, "nw_content_context_set_metadata_for_protocol")
	tryRegister(&_nw_content_context_set_relative_priority, lib, "nw_content_context_set_relative_priority")
	tryRegister(&_nw_data_transfer_report_collect, lib, "nw_data_transfer_report_collect")
	tryRegister(&_nw_data_transfer_report_copy_path_interface, lib, "nw_data_transfer_report_copy_path_interface")
	tryRegister(&_nw_data_transfer_report_get_duration_milliseconds, lib, "nw_data_transfer_report_get_duration_milliseconds")
	tryRegister(&_nw_data_transfer_report_get_path_count, lib, "nw_data_transfer_report_get_path_count")
	tryRegister(&_nw_data_transfer_report_get_path_radio_type, lib, "nw_data_transfer_report_get_path_radio_type")
	tryRegister(&_nw_data_transfer_report_get_received_application_byte_count, lib, "nw_data_transfer_report_get_received_application_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_ip_packet_count, lib, "nw_data_transfer_report_get_received_ip_packet_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_byte_count, lib, "nw_data_transfer_report_get_received_transport_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_duplicate_byte_count, lib, "nw_data_transfer_report_get_received_transport_duplicate_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_out_of_order_byte_count, lib, "nw_data_transfer_report_get_received_transport_out_of_order_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_application_byte_count, lib, "nw_data_transfer_report_get_sent_application_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_ip_packet_count, lib, "nw_data_transfer_report_get_sent_ip_packet_count")
	tryRegister(&_nw_data_transfer_report_get_sent_transport_byte_count, lib, "nw_data_transfer_report_get_sent_transport_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count, lib, "nw_data_transfer_report_get_sent_transport_retransmitted_byte_count")
	tryRegister(&_nw_data_transfer_report_get_state, lib, "nw_data_transfer_report_get_state")
	tryRegister(&_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds, lib, "nw_data_transfer_report_get_transport_minimum_rtt_milliseconds")
	tryRegister(&_nw_data_transfer_report_get_transport_rtt_variance, lib, "nw_data_transfer_report_get_transport_rtt_variance")
	tryRegister(&_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds, lib, "nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds")
	tryRegister(&_nw_endpoint_copy_address_string, lib, "nw_endpoint_copy_address_string")
	tryRegister(&_nw_endpoint_copy_port_string, lib, "nw_endpoint_copy_port_string")
	tryRegister(&_nw_endpoint_copy_txt_record, lib, "nw_endpoint_copy_txt_record")
	tryRegister(&_nw_endpoint_create_address, lib, "nw_endpoint_create_address")
	tryRegister(&_nw_endpoint_create_bonjour_service, lib, "nw_endpoint_create_bonjour_service")
	tryRegister(&_nw_endpoint_create_host, lib, "nw_endpoint_create_host")
	tryRegister(&_nw_endpoint_create_url, lib, "nw_endpoint_create_url")
	tryRegister(&_nw_endpoint_get_address, lib, "nw_endpoint_get_address")
	tryRegister(&_nw_endpoint_get_bonjour_service_domain, lib, "nw_endpoint_get_bonjour_service_domain")
	tryRegister(&_nw_endpoint_get_bonjour_service_name, lib, "nw_endpoint_get_bonjour_service_name")
	tryRegister(&_nw_endpoint_get_bonjour_service_type, lib, "nw_endpoint_get_bonjour_service_type")
	tryRegister(&_nw_endpoint_get_hostname, lib, "nw_endpoint_get_hostname")
	tryRegister(&_nw_endpoint_get_port, lib, "nw_endpoint_get_port")
	tryRegister(&_nw_endpoint_get_signature, lib, "nw_endpoint_get_signature")
	tryRegister(&_nw_endpoint_get_type, lib, "nw_endpoint_get_type")
	tryRegister(&_nw_endpoint_get_url, lib, "nw_endpoint_get_url")
	tryRegister(&_nw_error_copy_cf_error, lib, "nw_error_copy_cf_error")
	tryRegister(&_nw_error_get_error_code, lib, "nw_error_get_error_code")
	tryRegister(&_nw_error_get_error_domain, lib, "nw_error_get_error_domain")
	tryRegister(&_nw_establishment_report_copy_proxy_endpoint, lib, "nw_establishment_report_copy_proxy_endpoint")
	tryRegister(&_nw_establishment_report_enumerate_protocols, lib, "nw_establishment_report_enumerate_protocols")
	tryRegister(&_nw_establishment_report_enumerate_resolution_reports, lib, "nw_establishment_report_enumerate_resolution_reports")
	tryRegister(&_nw_establishment_report_enumerate_resolutions, lib, "nw_establishment_report_enumerate_resolutions")
	tryRegister(&_nw_establishment_report_get_attempt_started_after_milliseconds, lib, "nw_establishment_report_get_attempt_started_after_milliseconds")
	tryRegister(&_nw_establishment_report_get_duration_milliseconds, lib, "nw_establishment_report_get_duration_milliseconds")
	tryRegister(&_nw_establishment_report_get_previous_attempt_count, lib, "nw_establishment_report_get_previous_attempt_count")
	tryRegister(&_nw_establishment_report_get_proxy_configured, lib, "nw_establishment_report_get_proxy_configured")
	tryRegister(&_nw_establishment_report_get_used_proxy, lib, "nw_establishment_report_get_used_proxy")
	tryRegister(&_nw_ethernet_channel_cancel, lib, "nw_ethernet_channel_cancel")
	tryRegister(&_nw_ethernet_channel_create, lib, "nw_ethernet_channel_create")
	tryRegister(&_nw_ethernet_channel_create_with_parameters, lib, "nw_ethernet_channel_create_with_parameters")
	tryRegister(&_nw_ethernet_channel_get_maximum_payload_size, lib, "nw_ethernet_channel_get_maximum_payload_size")
	tryRegister(&_nw_ethernet_channel_send, lib, "nw_ethernet_channel_send")
	tryRegister(&_nw_ethernet_channel_set_queue, lib, "nw_ethernet_channel_set_queue")
	tryRegister(&_nw_ethernet_channel_set_receive_handler, lib, "nw_ethernet_channel_set_receive_handler")
	tryRegister(&_nw_ethernet_channel_set_state_changed_handler, lib, "nw_ethernet_channel_set_state_changed_handler")
	tryRegister(&_nw_ethernet_channel_start, lib, "nw_ethernet_channel_start")
	tryRegister(&_nw_framer_async, lib, "nw_framer_async")
	tryRegister(&_nw_framer_copy_options, lib, "nw_framer_copy_options")
	tryRegister(&_nw_framer_copy_parameters, lib, "nw_framer_copy_parameters")
	tryRegister(&_nw_framer_copy_remote_endpoint, lib, "nw_framer_copy_remote_endpoint")
	tryRegister(&_nw_framer_create_definition, lib, "nw_framer_create_definition")
	tryRegister(&_nw_framer_create_options, lib, "nw_framer_create_options")
	tryRegister(&_nw_framer_deliver_input, lib, "nw_framer_deliver_input")
	tryRegister(&_nw_framer_deliver_input_no_copy, lib, "nw_framer_deliver_input_no_copy")
	tryRegister(&_nw_framer_mark_failed_with_error, lib, "nw_framer_mark_failed_with_error")
	tryRegister(&_nw_framer_mark_ready, lib, "nw_framer_mark_ready")
	tryRegister(&_nw_framer_message_access_value, lib, "nw_framer_message_access_value")
	tryRegister(&_nw_framer_message_copy_object_value, lib, "nw_framer_message_copy_object_value")
	tryRegister(&_nw_framer_message_create, lib, "nw_framer_message_create")
	tryRegister(&_nw_framer_message_set_object_value, lib, "nw_framer_message_set_object_value")
	tryRegister(&_nw_framer_message_set_value, lib, "nw_framer_message_set_value")
	tryRegister(&_nw_framer_options_copy_object_value, lib, "nw_framer_options_copy_object_value")
	tryRegister(&_nw_framer_options_set_object_value, lib, "nw_framer_options_set_object_value")
	tryRegister(&_nw_framer_parse_input, lib, "nw_framer_parse_input")
	tryRegister(&_nw_framer_parse_output, lib, "nw_framer_parse_output")
	tryRegister(&_nw_framer_pass_through_input, lib, "nw_framer_pass_through_input")
	tryRegister(&_nw_framer_pass_through_output, lib, "nw_framer_pass_through_output")
	tryRegister(&_nw_framer_prepend_application_protocol, lib, "nw_framer_prepend_application_protocol")
	tryRegister(&_nw_framer_protocol_create_message, lib, "nw_framer_protocol_create_message")
	tryRegister(&_nw_framer_schedule_wakeup, lib, "nw_framer_schedule_wakeup")
	tryRegister(&_nw_framer_set_cleanup_handler, lib, "nw_framer_set_cleanup_handler")
	tryRegister(&_nw_framer_set_input_handler, lib, "nw_framer_set_input_handler")
	tryRegister(&_nw_framer_set_output_handler, lib, "nw_framer_set_output_handler")
	tryRegister(&_nw_framer_set_stop_handler, lib, "nw_framer_set_stop_handler")
	tryRegister(&_nw_framer_set_wakeup_handler, lib, "nw_framer_set_wakeup_handler")
	tryRegister(&_nw_framer_write_output, lib, "nw_framer_write_output")
	tryRegister(&_nw_framer_write_output_data, lib, "nw_framer_write_output_data")
	tryRegister(&_nw_framer_write_output_no_copy, lib, "nw_framer_write_output_no_copy")
	tryRegister(&_nw_group_descriptor_add_endpoint, lib, "nw_group_descriptor_add_endpoint")
	tryRegister(&_nw_group_descriptor_create_multicast, lib, "nw_group_descriptor_create_multicast")
	tryRegister(&_nw_group_descriptor_create_multiplex, lib, "nw_group_descriptor_create_multiplex")
	tryRegister(&_nw_group_descriptor_enumerate_endpoints, lib, "nw_group_descriptor_enumerate_endpoints")
	tryRegister(&_nw_interface_get_index, lib, "nw_interface_get_index")
	tryRegister(&_nw_interface_get_name, lib, "nw_interface_get_name")
	tryRegister(&_nw_interface_get_type, lib, "nw_interface_get_type")
	tryRegister(&_nw_ip_create_metadata, lib, "nw_ip_create_metadata")
	tryRegister(&_nw_ip_metadata_get_ecn_flag, lib, "nw_ip_metadata_get_ecn_flag")
	tryRegister(&_nw_ip_metadata_get_receive_time, lib, "nw_ip_metadata_get_receive_time")
	tryRegister(&_nw_ip_metadata_get_service_class, lib, "nw_ip_metadata_get_service_class")
	tryRegister(&_nw_ip_metadata_set_ecn_flag, lib, "nw_ip_metadata_set_ecn_flag")
	tryRegister(&_nw_ip_metadata_set_service_class, lib, "nw_ip_metadata_set_service_class")
	tryRegister(&_nw_ip_options_set_calculate_receive_time, lib, "nw_ip_options_set_calculate_receive_time")
	tryRegister(&_nw_ip_options_set_disable_fragmentation, lib, "nw_ip_options_set_disable_fragmentation")
	tryRegister(&_nw_ip_options_set_disable_multicast_loopback, lib, "nw_ip_options_set_disable_multicast_loopback")
	tryRegister(&_nw_ip_options_set_hop_limit, lib, "nw_ip_options_set_hop_limit")
	tryRegister(&_nw_ip_options_set_use_minimum_mtu, lib, "nw_ip_options_set_use_minimum_mtu")
	tryRegister(&_nw_listener_cancel, lib, "nw_listener_cancel")
	tryRegister(&_nw_listener_create, lib, "nw_listener_create")
	tryRegister(&_nw_listener_create_with_connection, lib, "nw_listener_create_with_connection")
	tryRegister(&_nw_listener_create_with_launchd_key, lib, "nw_listener_create_with_launchd_key")
	tryRegister(&_nw_listener_create_with_port, lib, "nw_listener_create_with_port")
	tryRegister(&_nw_listener_get_new_connection_limit, lib, "nw_listener_get_new_connection_limit")
	tryRegister(&_nw_listener_get_port, lib, "nw_listener_get_port")
	tryRegister(&_nw_listener_set_advertise_descriptor, lib, "nw_listener_set_advertise_descriptor")
	tryRegister(&_nw_listener_set_advertised_endpoint_changed_handler, lib, "nw_listener_set_advertised_endpoint_changed_handler")
	tryRegister(&_nw_listener_set_new_connection_group_handler, lib, "nw_listener_set_new_connection_group_handler")
	tryRegister(&_nw_listener_set_new_connection_handler, lib, "nw_listener_set_new_connection_handler")
	tryRegister(&_nw_listener_set_new_connection_limit, lib, "nw_listener_set_new_connection_limit")
	tryRegister(&_nw_listener_set_queue, lib, "nw_listener_set_queue")
	tryRegister(&_nw_listener_set_state_changed_handler, lib, "nw_listener_set_state_changed_handler")
	tryRegister(&_nw_listener_start, lib, "nw_listener_start")
	tryRegister(&_nw_multicast_group_descriptor_get_disable_unicast_traffic, lib, "nw_multicast_group_descriptor_get_disable_unicast_traffic")
	tryRegister(&_nw_multicast_group_descriptor_set_disable_unicast_traffic, lib, "nw_multicast_group_descriptor_set_disable_unicast_traffic")
	tryRegister(&_nw_multicast_group_descriptor_set_specific_source, lib, "nw_multicast_group_descriptor_set_specific_source")
	tryRegister(&_nw_parameters_clear_prohibited_interface_types, lib, "nw_parameters_clear_prohibited_interface_types")
	tryRegister(&_nw_parameters_clear_prohibited_interfaces, lib, "nw_parameters_clear_prohibited_interfaces")
	tryRegister(&_nw_parameters_copy, lib, "nw_parameters_copy")
	tryRegister(&_nw_parameters_copy_default_protocol_stack, lib, "nw_parameters_copy_default_protocol_stack")
	tryRegister(&_nw_parameters_copy_required_interface, lib, "nw_parameters_copy_required_interface")
	tryRegister(&_nw_parameters_create, lib, "nw_parameters_create")
	tryRegister(&_nw_parameters_create_application_service, lib, "nw_parameters_create_application_service")
	tryRegister(&_nw_parameters_create_custom_ip, lib, "nw_parameters_create_custom_ip")
	tryRegister(&_nw_parameters_create_quic, lib, "nw_parameters_create_quic")
	tryRegister(&_nw_parameters_create_secure_tcp, lib, "nw_parameters_create_secure_tcp")
	tryRegister(&_nw_parameters_create_secure_udp, lib, "nw_parameters_create_secure_udp")
	tryRegister(&_nw_parameters_get_allow_ultra_constrained, lib, "nw_parameters_get_allow_ultra_constrained")
	tryRegister(&_nw_parameters_get_attribution, lib, "nw_parameters_get_attribution")
	tryRegister(&_nw_parameters_get_expired_dns_behavior, lib, "nw_parameters_get_expired_dns_behavior")
	tryRegister(&_nw_parameters_get_fast_open_enabled, lib, "nw_parameters_get_fast_open_enabled")
	tryRegister(&_nw_parameters_get_include_peer_to_peer, lib, "nw_parameters_get_include_peer_to_peer")
	tryRegister(&_nw_parameters_get_multipath_service, lib, "nw_parameters_get_multipath_service")
	tryRegister(&_nw_parameters_get_prefer_no_proxy, lib, "nw_parameters_get_prefer_no_proxy")
	tryRegister(&_nw_parameters_get_prohibit_constrained, lib, "nw_parameters_get_prohibit_constrained")
	tryRegister(&_nw_parameters_get_prohibit_expensive, lib, "nw_parameters_get_prohibit_expensive")
	tryRegister(&_nw_parameters_get_required_interface_type, lib, "nw_parameters_get_required_interface_type")
	tryRegister(&_nw_parameters_get_service_class, lib, "nw_parameters_get_service_class")
	tryRegister(&_nw_parameters_iterate_prohibited_interface_types, lib, "nw_parameters_iterate_prohibited_interface_types")
	tryRegister(&_nw_parameters_iterate_prohibited_interfaces, lib, "nw_parameters_iterate_prohibited_interfaces")
	tryRegister(&_nw_parameters_prohibit_interface, lib, "nw_parameters_prohibit_interface")
	tryRegister(&_nw_parameters_prohibit_interface_type, lib, "nw_parameters_prohibit_interface_type")
	tryRegister(&_nw_parameters_require_interface, lib, "nw_parameters_require_interface")
	tryRegister(&_nw_parameters_requires_dnssec_validation, lib, "nw_parameters_requires_dnssec_validation")
	tryRegister(&_nw_parameters_set_allow_ultra_constrained, lib, "nw_parameters_set_allow_ultra_constrained")
	tryRegister(&_nw_parameters_set_attribution, lib, "nw_parameters_set_attribution")
	tryRegister(&_nw_parameters_set_expired_dns_behavior, lib, "nw_parameters_set_expired_dns_behavior")
	tryRegister(&_nw_parameters_set_fast_open_enabled, lib, "nw_parameters_set_fast_open_enabled")
	tryRegister(&_nw_parameters_set_include_peer_to_peer, lib, "nw_parameters_set_include_peer_to_peer")
	tryRegister(&_nw_parameters_set_multipath_service, lib, "nw_parameters_set_multipath_service")
	tryRegister(&_nw_parameters_set_prefer_no_proxy, lib, "nw_parameters_set_prefer_no_proxy")
	tryRegister(&_nw_parameters_set_privacy_context, lib, "nw_parameters_set_privacy_context")
	tryRegister(&_nw_parameters_set_prohibit_constrained, lib, "nw_parameters_set_prohibit_constrained")
	tryRegister(&_nw_parameters_set_prohibit_expensive, lib, "nw_parameters_set_prohibit_expensive")
	tryRegister(&_nw_parameters_set_required_interface_type, lib, "nw_parameters_set_required_interface_type")
	tryRegister(&_nw_parameters_set_requires_dnssec_validation, lib, "nw_parameters_set_requires_dnssec_validation")
	tryRegister(&_nw_parameters_set_service_class, lib, "nw_parameters_set_service_class")
	tryRegister(&_nw_path_copy_effective_remote_endpoint, lib, "nw_path_copy_effective_remote_endpoint")
	tryRegister(&_nw_path_enumerate_gateways, lib, "nw_path_enumerate_gateways")
	tryRegister(&_nw_path_enumerate_interfaces, lib, "nw_path_enumerate_interfaces")
	tryRegister(&_nw_path_get_link_quality, lib, "nw_path_get_link_quality")
	tryRegister(&_nw_path_get_status, lib, "nw_path_get_status")
	tryRegister(&_nw_path_get_unsatisfied_reason, lib, "nw_path_get_unsatisfied_reason")
	tryRegister(&_nw_path_has_dns, lib, "nw_path_has_dns")
	tryRegister(&_nw_path_has_ipv4, lib, "nw_path_has_ipv4")
	tryRegister(&_nw_path_has_ipv6, lib, "nw_path_has_ipv6")
	tryRegister(&_nw_path_is_constrained, lib, "nw_path_is_constrained")
	tryRegister(&_nw_path_is_equal, lib, "nw_path_is_equal")
	tryRegister(&_nw_path_is_expensive, lib, "nw_path_is_expensive")
	tryRegister(&_nw_path_is_ultra_constrained, lib, "nw_path_is_ultra_constrained")
	tryRegister(&_nw_path_monitor_cancel, lib, "nw_path_monitor_cancel")
	tryRegister(&_nw_path_monitor_create, lib, "nw_path_monitor_create")
	tryRegister(&_nw_path_monitor_create_for_ethernet_channel, lib, "nw_path_monitor_create_for_ethernet_channel")
	tryRegister(&_nw_path_monitor_create_with_type, lib, "nw_path_monitor_create_with_type")
	tryRegister(&_nw_path_monitor_prohibit_interface_type, lib, "nw_path_monitor_prohibit_interface_type")
	tryRegister(&_nw_path_monitor_set_cancel_handler, lib, "nw_path_monitor_set_cancel_handler")
	tryRegister(&_nw_path_monitor_set_queue, lib, "nw_path_monitor_set_queue")
	tryRegister(&_nw_path_monitor_set_update_handler, lib, "nw_path_monitor_set_update_handler")
	tryRegister(&_nw_path_monitor_start, lib, "nw_path_monitor_start")
	tryRegister(&_nw_path_uses_interface_type, lib, "nw_path_uses_interface_type")
	tryRegister(&_nw_privacy_context_add_proxy, lib, "nw_privacy_context_add_proxy")
	tryRegister(&_nw_privacy_context_clear_proxies, lib, "nw_privacy_context_clear_proxies")
	tryRegister(&_nw_privacy_context_create, lib, "nw_privacy_context_create")
	tryRegister(&_nw_privacy_context_disable_logging, lib, "nw_privacy_context_disable_logging")
	tryRegister(&_nw_privacy_context_flush_cache, lib, "nw_privacy_context_flush_cache")
	tryRegister(&_nw_privacy_context_require_encrypted_name_resolution, lib, "nw_privacy_context_require_encrypted_name_resolution")
	tryRegister(&_nw_protocol_copy_ip_definition, lib, "nw_protocol_copy_ip_definition")
	tryRegister(&_nw_protocol_copy_quic_definition, lib, "nw_protocol_copy_quic_definition")
	tryRegister(&_nw_protocol_copy_tcp_definition, lib, "nw_protocol_copy_tcp_definition")
	tryRegister(&_nw_protocol_copy_tls_definition, lib, "nw_protocol_copy_tls_definition")
	tryRegister(&_nw_protocol_copy_udp_definition, lib, "nw_protocol_copy_udp_definition")
	tryRegister(&_nw_protocol_copy_ws_definition, lib, "nw_protocol_copy_ws_definition")
	tryRegister(&_nw_protocol_definition_is_equal, lib, "nw_protocol_definition_is_equal")
	tryRegister(&_nw_protocol_metadata_copy_definition, lib, "nw_protocol_metadata_copy_definition")
	tryRegister(&_nw_protocol_metadata_is_framer_message, lib, "nw_protocol_metadata_is_framer_message")
	tryRegister(&_nw_protocol_metadata_is_ip, lib, "nw_protocol_metadata_is_ip")
	tryRegister(&_nw_protocol_metadata_is_quic, lib, "nw_protocol_metadata_is_quic")
	tryRegister(&_nw_protocol_metadata_is_tcp, lib, "nw_protocol_metadata_is_tcp")
	tryRegister(&_nw_protocol_metadata_is_tls, lib, "nw_protocol_metadata_is_tls")
	tryRegister(&_nw_protocol_metadata_is_udp, lib, "nw_protocol_metadata_is_udp")
	tryRegister(&_nw_protocol_metadata_is_ws, lib, "nw_protocol_metadata_is_ws")
	tryRegister(&_nw_protocol_options_copy_definition, lib, "nw_protocol_options_copy_definition")
	tryRegister(&_nw_protocol_options_is_quic, lib, "nw_protocol_options_is_quic")
	tryRegister(&_nw_protocol_stack_clear_application_protocols, lib, "nw_protocol_stack_clear_application_protocols")
	tryRegister(&_nw_protocol_stack_copy_internet_protocol, lib, "nw_protocol_stack_copy_internet_protocol")
	tryRegister(&_nw_protocol_stack_copy_transport_protocol, lib, "nw_protocol_stack_copy_transport_protocol")
	tryRegister(&_nw_protocol_stack_iterate_application_protocols, lib, "nw_protocol_stack_iterate_application_protocols")
	tryRegister(&_nw_protocol_stack_prepend_application_protocol, lib, "nw_protocol_stack_prepend_application_protocol")
	tryRegister(&_nw_protocol_stack_set_transport_protocol, lib, "nw_protocol_stack_set_transport_protocol")
	tryRegister(&_nw_proxy_config_add_excluded_domain, lib, "nw_proxy_config_add_excluded_domain")
	tryRegister(&_nw_proxy_config_add_match_domain, lib, "nw_proxy_config_add_match_domain")
	tryRegister(&_nw_proxy_config_clear_excluded_domains, lib, "nw_proxy_config_clear_excluded_domains")
	tryRegister(&_nw_proxy_config_clear_match_domains, lib, "nw_proxy_config_clear_match_domains")
	tryRegister(&_nw_proxy_config_create_http_connect, lib, "nw_proxy_config_create_http_connect")
	tryRegister(&_nw_proxy_config_create_oblivious_http, lib, "nw_proxy_config_create_oblivious_http")
	tryRegister(&_nw_proxy_config_create_relay, lib, "nw_proxy_config_create_relay")
	tryRegister(&_nw_proxy_config_create_socksv5, lib, "nw_proxy_config_create_socksv5")
	tryRegister(&_nw_proxy_config_enumerate_excluded_domains, lib, "nw_proxy_config_enumerate_excluded_domains")
	tryRegister(&_nw_proxy_config_enumerate_match_domains, lib, "nw_proxy_config_enumerate_match_domains")
	tryRegister(&_nw_proxy_config_get_failover_allowed, lib, "nw_proxy_config_get_failover_allowed")
	tryRegister(&_nw_proxy_config_set_failover_allowed, lib, "nw_proxy_config_set_failover_allowed")
	tryRegister(&_nw_proxy_config_set_username_and_password, lib, "nw_proxy_config_set_username_and_password")
	tryRegister(&_nw_quic_add_tls_application_protocol, lib, "nw_quic_add_tls_application_protocol")
	tryRegister(&_nw_quic_copy_sec_protocol_metadata, lib, "nw_quic_copy_sec_protocol_metadata")
	tryRegister(&_nw_quic_copy_sec_protocol_options, lib, "nw_quic_copy_sec_protocol_options")
	tryRegister(&_nw_quic_create_options, lib, "nw_quic_create_options")
	tryRegister(&_nw_quic_get_application_error, lib, "nw_quic_get_application_error")
	tryRegister(&_nw_quic_get_application_error_reason, lib, "nw_quic_get_application_error_reason")
	tryRegister(&_nw_quic_get_idle_timeout, lib, "nw_quic_get_idle_timeout")
	tryRegister(&_nw_quic_get_initial_max_data, lib, "nw_quic_get_initial_max_data")
	tryRegister(&_nw_quic_get_initial_max_stream_data_bidirectional_remote, lib, "nw_quic_get_initial_max_stream_data_bidirectional_remote")
	tryRegister(&_nw_quic_get_initial_max_stream_data_unidirectional, lib, "nw_quic_get_initial_max_stream_data_unidirectional")
	tryRegister(&_nw_quic_get_initial_max_streams_bidirectional, lib, "nw_quic_get_initial_max_streams_bidirectional")
	tryRegister(&_nw_quic_get_initial_max_streams_unidirectional, lib, "nw_quic_get_initial_max_streams_unidirectional")
	tryRegister(&_nw_quic_get_keepalive_interval, lib, "nw_quic_get_keepalive_interval")
	tryRegister(&_nw_quic_get_max_datagram_frame_size, lib, "nw_quic_get_max_datagram_frame_size")
	tryRegister(&_nw_quic_get_max_udp_payload_size, lib, "nw_quic_get_max_udp_payload_size")
	tryRegister(&_nw_quic_get_remote_idle_timeout, lib, "nw_quic_get_remote_idle_timeout")
	tryRegister(&_nw_quic_get_remote_max_streams_bidirectional, lib, "nw_quic_get_remote_max_streams_bidirectional")
	tryRegister(&_nw_quic_get_remote_max_streams_unidirectional, lib, "nw_quic_get_remote_max_streams_unidirectional")
	tryRegister(&_nw_quic_get_stream_application_error, lib, "nw_quic_get_stream_application_error")
	tryRegister(&_nw_quic_get_stream_id, lib, "nw_quic_get_stream_id")
	tryRegister(&_nw_quic_get_stream_is_datagram, lib, "nw_quic_get_stream_is_datagram")
	tryRegister(&_nw_quic_get_stream_is_unidirectional, lib, "nw_quic_get_stream_is_unidirectional")
	tryRegister(&_nw_quic_get_stream_type, lib, "nw_quic_get_stream_type")
	tryRegister(&_nw_quic_get_stream_usable_datagram_frame_size, lib, "nw_quic_get_stream_usable_datagram_frame_size")
	tryRegister(&_nw_quic_set_application_error, lib, "nw_quic_set_application_error")
	tryRegister(&_nw_quic_set_idle_timeout, lib, "nw_quic_set_idle_timeout")
	tryRegister(&_nw_quic_set_initial_max_data, lib, "nw_quic_set_initial_max_data")
	tryRegister(&_nw_quic_set_initial_max_stream_data_bidirectional_remote, lib, "nw_quic_set_initial_max_stream_data_bidirectional_remote")
	tryRegister(&_nw_quic_set_initial_max_stream_data_unidirectional, lib, "nw_quic_set_initial_max_stream_data_unidirectional")
	tryRegister(&_nw_quic_set_initial_max_streams_bidirectional, lib, "nw_quic_set_initial_max_streams_bidirectional")
	tryRegister(&_nw_quic_set_initial_max_streams_unidirectional, lib, "nw_quic_set_initial_max_streams_unidirectional")
	tryRegister(&_nw_quic_set_keepalive_interval, lib, "nw_quic_set_keepalive_interval")
	tryRegister(&_nw_quic_set_max_datagram_frame_size, lib, "nw_quic_set_max_datagram_frame_size")
	tryRegister(&_nw_quic_set_max_udp_payload_size, lib, "nw_quic_set_max_udp_payload_size")
	tryRegister(&_nw_quic_set_stream_application_error, lib, "nw_quic_set_stream_application_error")
	tryRegister(&_nw_quic_set_stream_is_datagram, lib, "nw_quic_set_stream_is_datagram")
	tryRegister(&_nw_quic_set_stream_is_unidirectional, lib, "nw_quic_set_stream_is_unidirectional")
	tryRegister(&_nw_relay_hop_add_additional_http_header_field, lib, "nw_relay_hop_add_additional_http_header_field")
	tryRegister(&_nw_relay_hop_create, lib, "nw_relay_hop_create")
	tryRegister(&_nw_release, lib, "nw_release")
	tryRegister(&_nw_resolution_report_copy_preferred_endpoint, lib, "nw_resolution_report_copy_preferred_endpoint")
	tryRegister(&_nw_resolution_report_copy_successful_endpoint, lib, "nw_resolution_report_copy_successful_endpoint")
	tryRegister(&_nw_resolution_report_get_endpoint_count, lib, "nw_resolution_report_get_endpoint_count")
	tryRegister(&_nw_resolution_report_get_milliseconds, lib, "nw_resolution_report_get_milliseconds")
	tryRegister(&_nw_resolution_report_get_protocol, lib, "nw_resolution_report_get_protocol")
	tryRegister(&_nw_resolution_report_get_source, lib, "nw_resolution_report_get_source")
	tryRegister(&_nw_resolver_config_add_server_address, lib, "nw_resolver_config_add_server_address")
	tryRegister(&_nw_resolver_config_create_https, lib, "nw_resolver_config_create_https")
	tryRegister(&_nw_resolver_config_create_tls, lib, "nw_resolver_config_create_tls")
	tryRegister(&_nw_retain, lib, "nw_retain")
	tryRegister(&_nw_tcp_create_options, lib, "nw_tcp_create_options")
	tryRegister(&_nw_tcp_get_available_receive_buffer, lib, "nw_tcp_get_available_receive_buffer")
	tryRegister(&_nw_tcp_get_available_send_buffer, lib, "nw_tcp_get_available_send_buffer")
	tryRegister(&_nw_tcp_options_set_connection_timeout, lib, "nw_tcp_options_set_connection_timeout")
	tryRegister(&_nw_tcp_options_set_disable_ack_stretching, lib, "nw_tcp_options_set_disable_ack_stretching")
	tryRegister(&_nw_tcp_options_set_disable_ecn, lib, "nw_tcp_options_set_disable_ecn")
	tryRegister(&_nw_tcp_options_set_enable_fast_open, lib, "nw_tcp_options_set_enable_fast_open")
	tryRegister(&_nw_tcp_options_set_enable_keepalive, lib, "nw_tcp_options_set_enable_keepalive")
	tryRegister(&_nw_tcp_options_set_keepalive_count, lib, "nw_tcp_options_set_keepalive_count")
	tryRegister(&_nw_tcp_options_set_keepalive_idle_time, lib, "nw_tcp_options_set_keepalive_idle_time")
	tryRegister(&_nw_tcp_options_set_keepalive_interval, lib, "nw_tcp_options_set_keepalive_interval")
	tryRegister(&_nw_tcp_options_set_maximum_segment_size, lib, "nw_tcp_options_set_maximum_segment_size")
	tryRegister(&_nw_tcp_options_set_no_delay, lib, "nw_tcp_options_set_no_delay")
	tryRegister(&_nw_tcp_options_set_no_options, lib, "nw_tcp_options_set_no_options")
	tryRegister(&_nw_tcp_options_set_no_push, lib, "nw_tcp_options_set_no_push")
	tryRegister(&_nw_tcp_options_set_persist_timeout, lib, "nw_tcp_options_set_persist_timeout")
	tryRegister(&_nw_tcp_options_set_retransmit_connection_drop_time, lib, "nw_tcp_options_set_retransmit_connection_drop_time")
	tryRegister(&_nw_tcp_options_set_retransmit_fin_drop, lib, "nw_tcp_options_set_retransmit_fin_drop")
	tryRegister(&_nw_tls_copy_sec_protocol_metadata, lib, "nw_tls_copy_sec_protocol_metadata")
	tryRegister(&_nw_tls_copy_sec_protocol_options, lib, "nw_tls_copy_sec_protocol_options")
	tryRegister(&_nw_tls_create_options, lib, "nw_tls_create_options")
	tryRegister(&_nw_txt_record_access_bytes, lib, "nw_txt_record_access_bytes")
	tryRegister(&_nw_txt_record_access_key, lib, "nw_txt_record_access_key")
	tryRegister(&_nw_txt_record_apply, lib, "nw_txt_record_apply")
	tryRegister(&_nw_txt_record_copy, lib, "nw_txt_record_copy")
	tryRegister(&_nw_txt_record_create_dictionary, lib, "nw_txt_record_create_dictionary")
	tryRegister(&_nw_txt_record_create_with_bytes, lib, "nw_txt_record_create_with_bytes")
	tryRegister(&_nw_txt_record_find_key, lib, "nw_txt_record_find_key")
	tryRegister(&_nw_txt_record_get_key_count, lib, "nw_txt_record_get_key_count")
	tryRegister(&_nw_txt_record_is_dictionary, lib, "nw_txt_record_is_dictionary")
	tryRegister(&_nw_txt_record_is_equal, lib, "nw_txt_record_is_equal")
	tryRegister(&_nw_txt_record_remove_key, lib, "nw_txt_record_remove_key")
	tryRegister(&_nw_txt_record_set_key, lib, "nw_txt_record_set_key")
	tryRegister(&_nw_udp_create_metadata, lib, "nw_udp_create_metadata")
	tryRegister(&_nw_udp_create_options, lib, "nw_udp_create_options")
	tryRegister(&_nw_udp_options_set_prefer_no_checksum, lib, "nw_udp_options_set_prefer_no_checksum")
	tryRegister(&_nw_ws_create_metadata, lib, "nw_ws_create_metadata")
	tryRegister(&_nw_ws_create_options, lib, "nw_ws_create_options")
	tryRegister(&_nw_ws_metadata_copy_server_response, lib, "nw_ws_metadata_copy_server_response")
	tryRegister(&_nw_ws_metadata_get_close_code, lib, "nw_ws_metadata_get_close_code")
	tryRegister(&_nw_ws_metadata_get_opcode, lib, "nw_ws_metadata_get_opcode")
	tryRegister(&_nw_ws_metadata_set_close_code, lib, "nw_ws_metadata_set_close_code")
	tryRegister(&_nw_ws_metadata_set_pong_handler, lib, "nw_ws_metadata_set_pong_handler")
	tryRegister(&_nw_ws_options_add_additional_header, lib, "nw_ws_options_add_additional_header")
	tryRegister(&_nw_ws_options_add_subprotocol, lib, "nw_ws_options_add_subprotocol")
	tryRegister(&_nw_ws_options_set_auto_reply_ping, lib, "nw_ws_options_set_auto_reply_ping")
	tryRegister(&_nw_ws_options_set_client_request_handler, lib, "nw_ws_options_set_client_request_handler")
	tryRegister(&_nw_ws_options_set_maximum_message_size, lib, "nw_ws_options_set_maximum_message_size")
	tryRegister(&_nw_ws_options_set_skip_handshake, lib, "nw_ws_options_set_skip_handshake")
	tryRegister(&_nw_ws_request_enumerate_additional_headers, lib, "nw_ws_request_enumerate_additional_headers")
	tryRegister(&_nw_ws_request_enumerate_subprotocols, lib, "nw_ws_request_enumerate_subprotocols")
	tryRegister(&_nw_ws_response_add_additional_header, lib, "nw_ws_response_add_additional_header")
	tryRegister(&_nw_ws_response_create, lib, "nw_ws_response_create")
	tryRegister(&_nw_ws_response_enumerate_additional_headers, lib, "nw_ws_response_enumerate_additional_headers")
	tryRegister(&_nw_ws_response_get_selected_subprotocol, lib, "nw_ws_response_get_selected_subprotocol")
	tryRegister(&_nw_ws_response_get_status, lib, "nw_ws_response_get_status")
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



// Accesses the TXT record to advertise with the service.
//
// Added in macOS 10.15.
// Accesses the TXT record to advertise with the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_copy_txt_record_object(_:)
func nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t) Nw_txt_record_t {
	return _nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor)
}

// nw_advertise_descriptor_create_application_service is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_application_service(_:)
func nw_advertise_descriptor_create_application_service(application_service_name unsafe.Pointer) Nw_advertise_descriptor_t {
	return _nw_advertise_descriptor_create_application_service(application_service_name)
}

// Initializes a Bonjour service to advertise.
//
// Added in macOS 10.14.
// Initializes a Bonjour service to advertise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_bonjour_service(_:_:_:)
func nw_advertise_descriptor_create_bonjour_service(name unsafe.Pointer, type_ unsafe.Pointer, domain unsafe.Pointer) Nw_advertise_descriptor_t {
	return _nw_advertise_descriptor_create_bonjour_service(name, type_, domain)
}

// nw_advertise_descriptor_get_application_service_name is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_application_service_name(_:)
func nw_advertise_descriptor_get_application_service_name(advertise_descriptor Nw_advertise_descriptor_t) unsafe.Pointer {
	return _nw_advertise_descriptor_get_application_service_name(advertise_descriptor)
}

// Checks whether the service prohibits automatic renaming in the event of a name conflict.
//
// Added in macOS 10.14.
// Checks whether the service prohibits automatic renaming in the event of a name conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_no_auto_rename(_:)
func nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t) bool {
	return _nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor)
}

// Sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict.
//
// Added in macOS 10.14.
// Sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_no_auto_rename(_:_:)
func nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor Nw_advertise_descriptor_t, no_auto_rename bool) {
	_nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor, no_auto_rename)
}

// Sets the TXT record as a raw buffer to advertise with the service.
//
// Added in macOS 10.14.
// Sets the TXT record as a raw buffer to advertise with the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record(_:_:_:)
func nw_advertise_descriptor_set_txt_record(advertise_descriptor Nw_advertise_descriptor_t, txt_record unsafe.Pointer, txt_length uintptr) {
	_nw_advertise_descriptor_set_txt_record(advertise_descriptor, txt_record, txt_length)
}

// Sets the TXT record to advertise with the service.
//
// Added in macOS 10.15.
// Sets the TXT record to advertise with the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record_object(_:_:)
func nw_advertise_descriptor_set_txt_record_object(advertise_descriptor Nw_advertise_descriptor_t, txt_record Nw_txt_record_t) {
	_nw_advertise_descriptor_set_txt_record_object(advertise_descriptor, txt_record)
}

// nw_browse_descriptor_create_application_service is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_application_service(_:)
func nw_browse_descriptor_create_application_service(application_service_name unsafe.Pointer) Nw_browse_descriptor_t {
	return _nw_browse_descriptor_create_application_service(application_service_name)
}

// Initializes a service descriptor used to discover a Bonjour service.
//
// Added in macOS 10.15.
// Initializes a service descriptor used to discover a Bonjour service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_bonjour_service(_:_:)
func nw_browse_descriptor_create_bonjour_service(type_ unsafe.Pointer, domain unsafe.Pointer) Nw_browse_descriptor_t {
	return _nw_browse_descriptor_create_bonjour_service(type_, domain)
}

// nw_browse_descriptor_get_application_service_name is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_application_service_name(_:)
func nw_browse_descriptor_get_application_service_name(descriptor Nw_browse_descriptor_t) unsafe.Pointer {
	return _nw_browse_descriptor_get_application_service_name(descriptor)
}

// Accesses the Bonjour service domain set on a browse descriptor.
//
// Added in macOS 10.15.
// Accesses the Bonjour service domain set on a browse descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_domain(_:)
func nw_browse_descriptor_get_bonjour_service_domain(descriptor Nw_browse_descriptor_t) unsafe.Pointer {
	return _nw_browse_descriptor_get_bonjour_service_domain(descriptor)
}

// Accesses the Bonjour service type set on a browse descriptor.
//
// Added in macOS 10.15.
// Accesses the Bonjour service type set on a browse descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_type(_:)
func nw_browse_descriptor_get_bonjour_service_type(descriptor Nw_browse_descriptor_t) unsafe.Pointer {
	return _nw_browse_descriptor_get_bonjour_service_type(descriptor)
}

// Checks if the browse descriptor requires including associated TXT records with all results.
//
// Added in macOS 10.15.
// Checks if the browse descriptor requires including associated TXT records with all results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_include_txt_record(_:)
func nw_browse_descriptor_get_include_txt_record(descriptor Nw_browse_descriptor_t) bool {
	return _nw_browse_descriptor_get_include_txt_record(descriptor)
}

// Requires including associated TXT records with all results generated for this service descriptor.
//
// Added in macOS 10.15.
// Requires including associated TXT records with all results generated for this service descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_set_include_txt_record(_:_:)
func nw_browse_descriptor_set_include_txt_record(descriptor Nw_browse_descriptor_t, include_txt_record bool) {
	_nw_browse_descriptor_set_include_txt_record(descriptor, include_txt_record)
}

// The discovered service endpoint.
//
// Added in macOS 10.15.
// The discovered service endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_copy_endpoint(_:)
func nw_browse_result_copy_endpoint(result Nw_browse_result_t) Nw_endpoint_t {
	return _nw_browse_result_copy_endpoint(result)
}

// Accesses the TXT record associated with a discovered service.
//
// Added in macOS 10.15.
// Accesses the TXT record associated with a discovered service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_copy_txt_record_object(_:)
func nw_browse_result_copy_txt_record_object(result Nw_browse_result_t) Nw_txt_record_t {
	return _nw_browse_result_copy_txt_record_object(result)
}

// Enumerates the list of interfaces on which the service was discovered.
//
// Added in macOS 10.15.
// Enumerates the list of interfaces on which the service was discovered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interfaces(_:_:)
func nw_browse_result_enumerate_interfaces(result Nw_browse_result_t, enumerator unsafe.Pointer) {
	_nw_browse_result_enumerate_interfaces(result, enumerator)
}

// Compares two discovered services and calculates changes between them.
//
// Added in macOS 10.15.
// Compares two discovered services and calculates changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_get_changes(_:_:)
func nw_browse_result_get_changes(old_result Nw_browse_result_t, new_result Nw_browse_result_t) Nw_browse_result_change_t {
	return _nw_browse_result_get_changes(old_result, new_result)
}

// Accesses the number of interfaces associated with a discovered service.
//
// Added in macOS 10.15.
// Accesses the number of interfaces associated with a discovered service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_get_interfaces_count(_:)
func nw_browse_result_get_interfaces_count(result Nw_browse_result_t) uintptr {
	return _nw_browse_result_get_interfaces_count(result)
}

// Stops browsing for services.
//
// Added in macOS 10.15.
// Stops browsing for services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_cancel(_:)
func nw_browser_cancel(browser Nw_browser_t) {
	_nw_browser_cancel(browser)
}

// Accesses the service descriptor with which the browser was created.
//
// Added in macOS 10.15.
// Accesses the service descriptor with which the browser was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_copy_browse_descriptor(_:)
func nw_browser_copy_browse_descriptor(browser Nw_browser_t) Nw_browse_descriptor_t {
	return _nw_browser_copy_browse_descriptor(browser)
}

// Accesses the parameters with which the browser was created.
//
// Added in macOS 10.15.
// Accesses the parameters with which the browser was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_copy_parameters(_:)
func nw_browser_copy_parameters(browser Nw_browser_t) Nw_parameters_t {
	return _nw_browser_copy_parameters(browser)
}

// Initializes a browser with a type of service to discover.
//
// Added in macOS 10.15.
// Initializes a browser with a type of service to discover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_create(_:_:)
func nw_browser_create(descriptor Nw_browse_descriptor_t, parameters Nw_parameters_t) Nw_browser_t {
	return _nw_browser_create(descriptor, parameters)
}

// Sets the handler to receive updates about discovered services.
//
// Added in macOS 10.15.
// Sets the handler to receive updates about discovered services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_browse_results_changed_handler(_:_:)
func nw_browser_set_browse_results_changed_handler(browser Nw_browser_t, handler unsafe.Pointer) {
	_nw_browser_set_browse_results_changed_handler(browser, handler)
}

// Sets the queue on which all browser events will be delivered.
//
// Added in macOS 10.15.
// Sets the queue on which all browser events will be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_queue(_:_:)
func nw_browser_set_queue(browser Nw_browser_t, queue unsafe.Pointer) {
	_nw_browser_set_queue(browser, queue)
}

// Sets a handler to receive browser state updates.
//
// Added in macOS 10.15.
// Sets a handler to receive browser state updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_state_changed_handler(_:_:)
func nw_browser_set_state_changed_handler(browser Nw_browser_t, state_changed_handler unsafe.Pointer) {
	_nw_browser_set_state_changed_handler(browser, state_changed_handler)
}

// Starts browsing for services.
//
// Added in macOS 10.15.
// Starts browsing for services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_start(_:)
func nw_browser_start(browser Nw_browser_t) {
	_nw_browser_start(browser)
}

// Requests a copy of the connection’s establishment report once the connection is in the ready state.
//
// Added in macOS 10.15.
// Requests a copy of the connection’s establishment report once the connection is in the ready state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_access_establishment_report(_:_:_:)
func nw_connection_access_establishment_report(connection Nw_connection_t, queue unsafe.Pointer, access_block unsafe.Pointer) {
	_nw_connection_access_establishment_report(connection, queue, access_block)
}

// Defines a block in which calls to send and receive are processed as a batch to improve performance.
//
// Added in macOS 10.14.
// Defines a block in which calls to send and receive are processed as a batch to improve performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_batch(_:_:)
func nw_connection_batch(connection Nw_connection_t, batch_block unsafe.Pointer) {
	_nw_connection_batch(connection, batch_block)
}

// Cancels the connection and gracefully disconnects any established network protocols.
//
// Added in macOS 10.14.
// Cancels the connection and gracefully disconnects any established network protocols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_cancel(_:)
func nw_connection_cancel(connection Nw_connection_t) {
	_nw_connection_cancel(connection)
}

// Causes the current endpoint to be rejected, allowing the connection to try another resolved address.
//
// Added in macOS 10.14.
// Causes the current endpoint to be rejected, allowing the connection to try another resolved address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_cancel_current_endpoint(_:)
func nw_connection_cancel_current_endpoint(connection Nw_connection_t) {
	_nw_connection_cancel_current_endpoint(connection)
}

// Accesses the network path the connection is using.
//
// Added in macOS 10.14.
// Accesses the network path the connection is using.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_current_path(_:)
func nw_connection_copy_current_path(connection Nw_connection_t) Nw_path_t {
	return _nw_connection_copy_current_path(connection)
}

// Copies the description of the connection as a string.
//
// Added in macOS 10.14.
// Copies the description of the connection as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_description(_:)
func nw_connection_copy_description(connection Nw_connection_t) unsafe.Pointer {
	return _nw_connection_copy_description(connection)
}

// Accesses the endpoint with which the connection was created.
//
// Added in macOS 10.14.
// Accesses the endpoint with which the connection was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_endpoint(_:)
func nw_connection_copy_endpoint(connection Nw_connection_t) Nw_endpoint_t {
	return _nw_connection_copy_endpoint(connection)
}

// Accesses the parameters with which the connection was created.
//
// Added in macOS 10.14.
// Accesses the parameters with which the connection was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_parameters(_:)
func nw_connection_copy_parameters(connection Nw_connection_t) Nw_parameters_t {
	return _nw_connection_copy_parameters(connection)
}

// Retrieves the connection-wide metadata for a specific protocol.
//
// Added in macOS 10.14.
// Retrieves the connection-wide metadata for a specific protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_protocol_metadata(_:_:)
func nw_connection_copy_protocol_metadata(connection Nw_connection_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	return _nw_connection_copy_protocol_metadata(connection, definition)
}

// Initializes a new connection to a remote endpoint.
//
// Added in macOS 10.14.
// Initializes a new connection to a remote endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_create(_:_:)
func nw_connection_create(endpoint Nw_endpoint_t, parameters Nw_parameters_t) Nw_connection_t {
	return _nw_connection_create(endpoint, parameters)
}

// Begins a new data transfer report, which can later be collected.
//
// Added in macOS 10.15.
// Begins a new data transfer report, which can later be collected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_create_new_data_transfer_report(_:)
func nw_connection_create_new_data_transfer_report(connection Nw_connection_t) Nw_data_transfer_report_t {
	return _nw_connection_create_new_data_transfer_report(connection)
}

// Cancels the connection and immediately disconnects any established network protocols.
//
// Added in macOS 10.14.
// Cancels the connection and immediately disconnects any established network protocols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_force_cancel(_:)
func nw_connection_force_cancel(connection Nw_connection_t) {
	_nw_connection_force_cancel(connection)
}

// Accesses the maximum size of a datagram message that can be sent on a connection.
//
// Added in macOS 10.14.
// Accesses the maximum size of a datagram message that can be sent on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_get_maximum_datagram_size(_:)
func nw_connection_get_maximum_datagram_size(connection Nw_connection_t) uint32 {
	return _nw_connection_get_maximum_datagram_size(connection)
}

// Cancels the connection group object and leaves the network group.
//
// Added in macOS 11.0.
// Cancels the connection group object and leaves the network group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_cancel(_:)
func nw_connection_group_cancel(group Nw_connection_group_t) {
	_nw_connection_group_cancel(group)
}

// Accesses the descriptor of the group you use to initialize the connection group.
//
// Added in macOS 11.0.
// Accesses the descriptor of the group you use to initialize the connection group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_descriptor(_:)
func nw_connection_group_copy_descriptor(group Nw_connection_group_t) Nw_group_descriptor_t {
	return _nw_connection_group_copy_descriptor(group)
}

// Accesses the parameters with which you initialize the connection group.
//
// Added in macOS 11.0.
// Accesses the parameters with which you initialize the connection group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_parameters(_:)
func nw_connection_group_copy_parameters(group Nw_connection_group_t) Nw_parameters_t {
	return _nw_connection_group_copy_parameters(group)
}

// Accesses the network path on which you receive the message.
//
// Added in macOS 11.0.
// Accesses the network path on which you receive the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_path_for_message(_:_:)
func nw_connection_group_copy_path_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_path_t {
	return _nw_connection_group_copy_path_for_message(group, context)
}

// nw_connection_group_copy_protocol_metadata is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata(_:_:)
func nw_connection_group_copy_protocol_metadata(group Nw_connection_group_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	return _nw_connection_group_copy_protocol_metadata(group, definition)
}

// nw_connection_group_copy_protocol_metadata_for_message is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata_for_message(_:_:_:)
func nw_connection_group_copy_protocol_metadata_for_message(group Nw_connection_group_t, context Nw_content_context_t, definition Nw_protocol_definition_t) Nw_protocol_metadata_t {
	return _nw_connection_group_copy_protocol_metadata_for_message(group, context, definition)
}

// Accesses the endpoint that originates the message you receive.
//
// Added in macOS 11.0.
// Accesses the endpoint that originates the message you receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_remote_endpoint_for_message(_:_:)
func nw_connection_group_copy_remote_endpoint_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_endpoint_t {
	return _nw_connection_group_copy_remote_endpoint_for_message(group, context)
}

// Initializes a new connection group with a group identifier.
//
// Added in macOS 11.0.
// Initializes a new connection group with a group identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_create(_:_:)
func nw_connection_group_create(group_descriptor Nw_group_descriptor_t, parameters Nw_parameters_t) Nw_connection_group_t {
	return _nw_connection_group_create(group_descriptor, parameters)
}

// nw_connection_group_extract_connection is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection(_:_:_:)
func nw_connection_group_extract_connection(group Nw_connection_group_t, endpoint Nw_endpoint_t, protocol_options Nw_protocol_options_t) Nw_connection_t {
	return _nw_connection_group_extract_connection(group, endpoint, protocol_options)
}

// Converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint.
//
// Added in macOS 11.0.
// Converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection_for_message(_:_:)
func nw_connection_group_extract_connection_for_message(group Nw_connection_group_t, context Nw_content_context_t) Nw_connection_t {
	return _nw_connection_group_extract_connection_for_message(group, context)
}

// nw_connection_group_reinsert_extracted_connection is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_reinsert_extracted_connection(_:_:)
func nw_connection_group_reinsert_extracted_connection(group Nw_connection_group_t, connection Nw_connection_t) bool {
	return _nw_connection_group_reinsert_extracted_connection(group, connection)
}

// Sends a reply to the specific endpoint that originates a group message you receive.
//
// Added in macOS 11.0.
// Sends a reply to the specific endpoint that originates a group message you receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_reply(_:_:_:_:)
func nw_connection_group_reply(group Nw_connection_group_t, inbound_message Nw_content_context_t, outbound_message Nw_content_context_t, content unsafe.Pointer) {
	_nw_connection_group_reply(group, inbound_message, outbound_message, content)
}

// Sends data to the entire group, or to a specific member of the group.
//
// Added in macOS 11.0.
// Sends data to the entire group, or to a specific member of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_send_message(_:_:_:_:_:)
func nw_connection_group_send_message(group Nw_connection_group_t, content unsafe.Pointer, endpoint Nw_endpoint_t, context Nw_content_context_t, completion unsafe.Pointer) {
	_nw_connection_group_send_message(group, content, endpoint, context, completion)
}

// nw_connection_group_set_new_connection_handler is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_new_connection_handler(_:_:)
func nw_connection_group_set_new_connection_handler(group Nw_connection_group_t, new_connection_handler unsafe.Pointer) {
	_nw_connection_group_set_new_connection_handler(group, new_connection_handler)
}

// Sets the queue on which you handle connection group events.
//
// Added in macOS 11.0.
// Sets the queue on which you handle connection group events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_queue(_:_:)
func nw_connection_group_set_queue(group Nw_connection_group_t, queue unsafe.Pointer) {
	_nw_connection_group_set_queue(group, queue)
}

// Sets a handler that receives inbound messages from members of the group.
//
// Added in macOS 11.0.
// Sets a handler that receives inbound messages from members of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_receive_handler(_:_:_:_:)
func nw_connection_group_set_receive_handler(group Nw_connection_group_t, maximum_message_size uint32, reject_oversized_messages bool, receive_handler unsafe.Pointer) {
	_nw_connection_group_set_receive_handler(group, maximum_message_size, reject_oversized_messages, receive_handler)
}

// Sets a handler that receives connection group state updates.
//
// Added in macOS 11.0.
// Sets a handler that receives connection group state updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_state_changed_handler(_:_:)
func nw_connection_group_set_state_changed_handler(group Nw_connection_group_t, state_changed_handler unsafe.Pointer) {
	_nw_connection_group_set_state_changed_handler(group, state_changed_handler)
}

// Joins the group and registers to receive messages.
//
// Added in macOS 11.0.
// Joins the group and registers to receive messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_start(_:)
func nw_connection_group_start(group Nw_connection_group_t) {
	_nw_connection_group_start(group)
}

// Schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time.
//
// Added in macOS 10.14.
// Schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_receive(_:_:_:_:)
func nw_connection_receive(connection Nw_connection_t, minimum_incomplete_length uint32, maximum_length uint32, completion unsafe.Pointer) {
	_nw_connection_receive(connection, minimum_incomplete_length, maximum_length, completion)
}

// Schedules a single receive completion handler for a complete message, as opposed to a range of bytes.
//
// Added in macOS 10.14.
// Schedules a single receive completion handler for a complete message, as opposed to a range of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_receive_message(_:_:)
func nw_connection_receive_message(connection Nw_connection_t, completion unsafe.Pointer) {
	_nw_connection_receive_message(connection, completion)
}

// Restarts a connection that is in the waiting state.
//
// Added in macOS 10.14.
// Restarts a connection that is in the waiting state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_restart(_:)
func nw_connection_restart(connection Nw_connection_t) {
	_nw_connection_restart(connection)
}

// Sends data on a connection.
//
// Added in macOS 10.14.
// Sends data on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_send(_:_:_:_:_:)
func nw_connection_send(connection Nw_connection_t, content unsafe.Pointer, context Nw_content_context_t, is_complete bool, completion unsafe.Pointer) {
	_nw_connection_send(connection, content, context, is_complete, completion)
}

// Sets a handler that receives updates when an alternative network path is preferred over the current path.
//
// Added in macOS 10.14.
// Sets a handler that receives updates when an alternative network path is preferred over the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_better_path_available_handler(_:_:)
func nw_connection_set_better_path_available_handler(connection Nw_connection_t, handler unsafe.Pointer) {
	_nw_connection_set_better_path_available_handler(connection, handler)
}

// Sets a handler that receives network path updates.
//
// Added in macOS 10.14.
// Sets a handler that receives network path updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_path_changed_handler(_:_:)
func nw_connection_set_path_changed_handler(connection Nw_connection_t, handler unsafe.Pointer) {
	_nw_connection_set_path_changed_handler(connection, handler)
}

// Sets the queue on which all connection events are delivered.
//
// Added in macOS 10.14.
// Sets the queue on which all connection events are delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_queue(_:_:)
func nw_connection_set_queue(connection Nw_connection_t, queue unsafe.Pointer) {
	_nw_connection_set_queue(connection, queue)
}

// Sets a handler to receive connection state updates.
//
// Added in macOS 10.14.
// Sets a handler to receive connection state updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_state_changed_handler(_:_:)
func nw_connection_set_state_changed_handler(connection Nw_connection_t, handler unsafe.Pointer) {
	_nw_connection_set_state_changed_handler(connection, handler)
}

// Sets a handler that receives updates when data can be sent and received.
//
// Added in macOS 10.14.
// Sets a handler that receives updates when data can be sent and received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_viability_changed_handler(_:_:)
func nw_connection_set_viability_changed_handler(connection Nw_connection_t, handler unsafe.Pointer) {
	_nw_connection_set_viability_changed_handler(connection, handler)
}

// Starts establishing a connection.
//
// Added in macOS 10.14.
// Starts establishing a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_start(_:)
func nw_connection_start(connection Nw_connection_t) {
	_nw_connection_start(connection)
}

// Accesses the optional message context that must be sent before the context you are sending.
//
// Added in macOS 10.14.
// Accesses the optional message context that must be sent before the context you are sending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_copy_antecedent(_:)
func nw_content_context_copy_antecedent(context Nw_content_context_t) Nw_content_context_t {
	return _nw_content_context_copy_antecedent(context)
}

// Retreives the metadata associated with a specific protocol.
//
// Added in macOS 10.14.
// Retreives the metadata associated with a specific protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_copy_protocol_metadata(_:_:)
func nw_content_context_copy_protocol_metadata(context Nw_content_context_t, protocol_ Nw_protocol_definition_t) Nw_protocol_metadata_t {
	return _nw_content_context_copy_protocol_metadata(context, protocol_)
}

// Initializes a custom message context.
//
// Added in macOS 10.14.
// Initializes a custom message context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_create(_:)
func nw_content_context_create(context_identifier unsafe.Pointer) Nw_content_context_t {
	return _nw_content_context_create(context_identifier)
}

// Iterates through all protocol metadata associated with the message context.
//
// Added in macOS 10.14.
// Iterates through all protocol metadata associated with the message context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_foreach_protocol_metadata(_:_:)
func nw_content_context_foreach_protocol_metadata(context Nw_content_context_t) {
	_nw_content_context_foreach_protocol_metadata(context)
}

// Accesses the expiration set for this message context.
//
// Added in macOS 10.14.
// Accesses the expiration set for this message context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_expiration_milliseconds(_:)
func nw_content_context_get_expiration_milliseconds(context Nw_content_context_t) uint64 {
	return _nw_content_context_get_expiration_milliseconds(context)
}

// Accesses the identifier used to create this message context.
//
// Added in macOS 10.14.
// Accesses the identifier used to create this message context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_identifier(_:)
func nw_content_context_get_identifier(context Nw_content_context_t) unsafe.Pointer {
	return _nw_content_context_get_identifier(context)
}

// Checks whether this context represents the final message being received.
//
// Added in macOS 10.14.
// Checks whether this context represents the final message being received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_is_final(_:)
func nw_content_context_get_is_final(context Nw_content_context_t) bool {
	return _nw_content_context_get_is_final(context)
}

// Accesses the relative value of priority used to reorder contexts when sending.
//
// Added in macOS 10.14.
// Accesses the relative value of priority used to reorder contexts when sending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_relative_priority(_:)
func nw_content_context_get_relative_priority(context Nw_content_context_t) float64 {
	return _nw_content_context_get_relative_priority(context)
}

// Set an optional message context that must be sent before the context you are sending.
//
// Added in macOS 10.14.
// Set an optional message context that must be sent before the context you are sending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_antecedent(_:_:)
func nw_content_context_set_antecedent(context Nw_content_context_t, antecedent_context Nw_content_context_t) {
	_nw_content_context_set_antecedent(context, antecedent_context)
}

// Sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded.
//
// Added in macOS 10.14.
// Sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_expiration_milliseconds(_:_:)
func nw_content_context_set_expiration_milliseconds(context Nw_content_context_t, expiration_milliseconds uint64) {
	_nw_content_context_set_expiration_milliseconds(context, expiration_milliseconds)
}

// Sets a Boolean indicating if this context represents the final message being sent.
//
// Added in macOS 10.14.
// Sets a Boolean indicating if this context represents the final message being sent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_is_final(_:_:)
func nw_content_context_set_is_final(context Nw_content_context_t, is_final bool) {
	_nw_content_context_set_is_final(context, is_final)
}

// Sets protocol metadata to configure per-message or per-packet properties.
//
// Added in macOS 10.14.
// Sets protocol metadata to configure per-message or per-packet properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_metadata_for_protocol(_:_:)
func nw_content_context_set_metadata_for_protocol(context Nw_content_context_t, protocol_metadata Nw_protocol_metadata_t) {
	_nw_content_context_set_metadata_for_protocol(context, protocol_metadata)
}

// Sets the relative value of priority used to reorder contexts when sending.
//
// Added in macOS 10.14.
// Sets the relative value of priority used to reorder contexts when sending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_relative_priority(_:_:)
func nw_content_context_set_relative_priority(context Nw_content_context_t, relative_priority float64) {
	_nw_content_context_set_relative_priority(context, relative_priority)
}

// Stops an outstanding data transfer report and calculates the results.
//
// Added in macOS 10.15.
// Stops an outstanding data transfer report and calculates the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect(_:_:_:)
func nw_data_transfer_report_collect(report Nw_data_transfer_report_t, queue unsafe.Pointer, collect_block unsafe.Pointer) {
	_nw_data_transfer_report_collect(report, queue, collect_block)
}

// Accesses the network interface the path used.
//
// Added in macOS 10.15.
// Accesses the network interface the path used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_copy_path_interface(_:_:)
func nw_data_transfer_report_copy_path_interface(report Nw_data_transfer_report_t, path_index uint32) Nw_interface_t {
	return _nw_data_transfer_report_copy_path_interface(report, path_index)
}

// Checks the duration of the data transfer report, from when it was started to when it was collected.
//
// Added in macOS 10.15.
// Checks the duration of the data transfer report, from when it was started to when it was collected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_duration_milliseconds(_:)
func nw_data_transfer_report_get_duration_milliseconds(report Nw_data_transfer_report_t) uint64 {
	return _nw_data_transfer_report_get_duration_milliseconds(report)
}

// Checks the number of valid paths in the report.
//
// Added in macOS 10.15.
// Checks the number of valid paths in the report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_count(_:)
func nw_data_transfer_report_get_path_count(report Nw_data_transfer_report_t) uint32 {
	return _nw_data_transfer_report_get_path_count(report)
}

// nw_data_transfer_report_get_path_radio_type is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_radio_type(_:_:)
func nw_data_transfer_report_get_path_radio_type(report Nw_data_transfer_report_t, path_index uint32) unsafe.Pointer {
	return _nw_data_transfer_report_get_path_radio_type(report, path_index)
}

// Accesses the number of bytes the connection delivered.
//
// Added in macOS 10.15.
// Accesses the number of bytes the connection delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_application_byte_count(_:_:)
func nw_data_transfer_report_get_received_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_received_application_byte_count(report, path_index)
}

// Accesses the number of IP packets the connection received.
//
// Added in macOS 10.15.
// Accesses the number of IP packets the connection received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_ip_packet_count(_:_:)
func nw_data_transfer_report_get_received_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_received_ip_packet_count(report, path_index)
}

// Accesses the number of bytes the transport protocol delivered.
//
// Added in macOS 10.15.
// Accesses the number of bytes the transport protocol delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_received_transport_byte_count(report, path_index)
}

// Accesses the number of duplicated bytes the transport protocol detected.
//
// Added in macOS 10.15.
// Accesses the number of duplicated bytes the transport protocol detected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_duplicate_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_duplicate_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_received_transport_duplicate_byte_count(report, path_index)
}

// Accesses the number of bytes the transport protocol received out of order.
//
// Added in macOS 10.15.
// Accesses the number of bytes the transport protocol received out of order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_out_of_order_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report, path_index)
}

// Accesses the number of bytes sent on the connection.
//
// Added in macOS 10.15.
// Accesses the number of bytes sent on the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_application_byte_count(_:_:)
func nw_data_transfer_report_get_sent_application_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_sent_application_byte_count(report, path_index)
}

// Accesses the number of IP packets the connection sent.
//
// Added in macOS 10.15.
// Accesses the number of IP packets the connection sent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_ip_packet_count(_:_:)
func nw_data_transfer_report_get_sent_ip_packet_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_sent_ip_packet_count(report, path_index)
}

// Accesses the number of bytes sent into the transport protocol.
//
// Added in macOS 10.15.
// Accesses the number of bytes sent into the transport protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_byte_count(_:_:)
func nw_data_transfer_report_get_sent_transport_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_sent_transport_byte_count(report, path_index)
}

// Accesses the number of bytes the transport protocol retransmitted.
//
// Added in macOS 10.15.
// Accesses the number of bytes the transport protocol retransmitted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(_:_:)
func nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report, path_index)
}

// Checks whether a data transfer report is collected.
//
// Added in macOS 10.15.
// Checks whether a data transfer report is collected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_state(_:)
func nw_data_transfer_report_get_state(report Nw_data_transfer_report_t) unsafe.Pointer {
	return _nw_data_transfer_report_get_state(report)
}

// Accesses the minimum round-trip time the transport protocol measured, in milliseconds.
//
// Added in macOS 10.15.
// Accesses the minimum round-trip time the transport protocol measured, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(_:_:)
func nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report, path_index)
}

// Accesses the variance of the round-trip time the transport protocol measured.
//
// Added in macOS 10.15.
// Accesses the variance of the round-trip time the transport protocol measured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_rtt_variance(_:_:)
func nw_data_transfer_report_get_transport_rtt_variance(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_transport_rtt_variance(report, path_index)
}

// Accesses the smoothed round-trip time the transport protocol measured, in milliseconds.
//
// Added in macOS 10.15.
// Accesses the smoothed round-trip time the transport protocol measured, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(_:_:)
func nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report Nw_data_transfer_report_t, path_index uint32) uint64 {
	return _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report, path_index)
}

// Copies the address of an endpoint as a string.
//
// Added in macOS 10.14.
// Copies the address of an endpoint as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_address_string(_:)
func nw_endpoint_copy_address_string(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_copy_address_string(endpoint)
}

// Copies the port of an endpoint as a string.
//
// Added in macOS 10.14.
// Copies the port of an endpoint as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_port_string(_:)
func nw_endpoint_copy_port_string(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_copy_port_string(endpoint)
}

// nw_endpoint_copy_txt_record is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_txt_record(_:)
func nw_endpoint_copy_txt_record(endpoint Nw_endpoint_t) Nw_txt_record_t {
	return _nw_endpoint_copy_txt_record(endpoint)
}

// Creates a network endpoint with an address structure.
//
// Added in macOS 10.14.
// Creates a network endpoint with an address structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_address(_:)
func nw_endpoint_create_address(address unsafe.Pointer) Nw_endpoint_t {
	return _nw_endpoint_create_address(address)
}

// Creates a network endpoint with a Bonjour service name, type, and domain.
//
// Added in macOS 10.14.
// Creates a network endpoint with a Bonjour service name, type, and domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_bonjour_service(_:_:_:)
func nw_endpoint_create_bonjour_service(name unsafe.Pointer, type_ unsafe.Pointer, domain unsafe.Pointer) Nw_endpoint_t {
	return _nw_endpoint_create_bonjour_service(name, type_, domain)
}

// Creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address.
//
// Added in macOS 10.14.
// Creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_host(_:_:)
func nw_endpoint_create_host(hostname unsafe.Pointer, port unsafe.Pointer) Nw_endpoint_t {
	return _nw_endpoint_create_host(hostname, port)
}

// Creates a network endpoint with a URL string.
//
// Added in macOS 10.15.
// Creates a network endpoint with a URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_url(_:)
func nw_endpoint_create_url(url unsafe.Pointer) Nw_endpoint_t {
	return _nw_endpoint_create_url(url)
}

// Accesses the address structure stored in an address endpoint.
//
// Added in macOS 10.14.
// Accesses the address structure stored in an address endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_address(_:)
func nw_endpoint_get_address(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_address(endpoint)
}

// Accesses the Bonjour service domain stored in an endpoint.
//
// Added in macOS 10.14.
// Accesses the Bonjour service domain stored in an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_domain(_:)
func nw_endpoint_get_bonjour_service_domain(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_domain(endpoint)
}

// Accesses the Bonjour service name stored in an endpoint.
//
// Added in macOS 10.14.
// Accesses the Bonjour service name stored in an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_name(_:)
func nw_endpoint_get_bonjour_service_name(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_name(endpoint)
}

// Accesses the Bonjour service type stored in an endpoint.
//
// Added in macOS 10.14.
// Accesses the Bonjour service type stored in an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_type(_:)
func nw_endpoint_get_bonjour_service_type(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_type(endpoint)
}

// Accesses the hostname stored in an endpoint.
//
// Added in macOS 10.14.
// Accesses the hostname stored in an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_hostname(_:)
func nw_endpoint_get_hostname(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_hostname(endpoint)
}

// Accesses the port stored in an endpoint, in host-byte order.
//
// Added in macOS 10.14.
// Accesses the port stored in an endpoint, in host-byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_port(_:)
func nw_endpoint_get_port(endpoint Nw_endpoint_t) uint16 {
	return _nw_endpoint_get_port(endpoint)
}

// nw_endpoint_get_signature is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_signature(_:_:)
func nw_endpoint_get_signature(endpoint Nw_endpoint_t, out_signature_length unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_signature(endpoint, out_signature_length)
}

// Accesses the type of a endpoint.
//
// Added in macOS 10.14.
// Accesses the type of a endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_type(_:)
func nw_endpoint_get_type(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_type(endpoint)
}

// Accesses the URL string stored in an endpoint.
//
// Added in macOS 10.15.
// Accesses the URL string stored in an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_url(_:)
func nw_endpoint_get_url(endpoint Nw_endpoint_t) unsafe.Pointer {
	return _nw_endpoint_get_url(endpoint)
}

// Returns a copy of a network error.
//
// Added in macOS 10.14.
// Returns a copy of a network error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_copy_cf_error(_:)
func nw_error_copy_cf_error(error_ Nw_error_t) ErrorRef {
	return _nw_error_copy_cf_error(error_)
}

// Accesses the specific code of the network error.
//
// Added in macOS 10.14.
// Accesses the specific code of the network error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_get_error_code(_:)
func nw_error_get_error_code(error_ Nw_error_t) int {
	return _nw_error_get_error_code(error_)
}

// Accesses the domain of the network error.
//
// Added in macOS 10.14.
// Accesses the domain of the network error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_get_error_domain(_:)
func nw_error_get_error_domain(error_ Nw_error_t) unsafe.Pointer {
	return _nw_error_get_error_domain(error_)
}

// Accesses the endpoint of the proxy the connection used.
//
// Added in macOS 10.15.
// Accesses the endpoint of the proxy the connection used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_copy_proxy_endpoint(_:)
func nw_establishment_report_copy_proxy_endpoint(report Nw_establishment_report_t) Nw_endpoint_t {
	return _nw_establishment_report_copy_proxy_endpoint(report)
}

// Iterates a list of protocol handshakes in order from first completed to last completed.
//
// Added in macOS 10.15.
// Iterates a list of protocol handshakes in order from first completed to last completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_protocols(_:_:)
func nw_establishment_report_enumerate_protocols(report Nw_establishment_report_t, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_protocols(report, enumerate_block)
}

// nw_establishment_report_enumerate_resolution_reports is a Network function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolution_reports(_:_:)
func nw_establishment_report_enumerate_resolution_reports(report Nw_establishment_report_t, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_resolution_reports(report, enumerate_block)
}

// Iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved.
//
// Added in macOS 10.15.
// Iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolutions(_:_:)
func nw_establishment_report_enumerate_resolutions(report Nw_establishment_report_t, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_resolutions(report, enumerate_block)
}

// Accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds.
//
// Added in macOS 10.15.
// Accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_attempt_started_after_milliseconds(_:)
func nw_establishment_report_get_attempt_started_after_milliseconds(report Nw_establishment_report_t) uint64 {
	return _nw_establishment_report_get_attempt_started_after_milliseconds(report)
}

// Checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state.
//
// Added in macOS 10.15.
// Checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_duration_milliseconds(_:)
func nw_establishment_report_get_duration_milliseconds(report Nw_establishment_report_t) uint64 {
	return _nw_establishment_report_get_duration_milliseconds(report)
}

// Checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state.
//
// Added in macOS 10.15.
// Checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_previous_attempt_count(_:)
func nw_establishment_report_get_previous_attempt_count(report Nw_establishment_report_t) uint32 {
	return _nw_establishment_report_get_previous_attempt_count(report)
}

// Checks whether a proxy was configured on the connection.
//
// Added in macOS 10.15.
// Checks whether a proxy was configured on the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_proxy_configured(_:)
func nw_establishment_report_get_proxy_configured(report Nw_establishment_report_t) bool {
	return _nw_establishment_report_get_proxy_configured(report)
}

// Checks whether the connection used a proxy.
//
// Added in macOS 10.15.
// Checks whether the connection used a proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_used_proxy(_:)
func nw_establishment_report_get_used_proxy(report Nw_establishment_report_t) bool {
	return _nw_establishment_report_get_used_proxy(report)
}

// Unregisters the channel from the interface.
//
// Added in macOS 10.15.
// Unregisters the channel from the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_cancel(_:)
func nw_ethernet_channel_cancel(ethernet_channel Nw_ethernet_channel_t) {
	_nw_ethernet_channel_cancel(ethernet_channel)
}

// Initializes an Ethernet channel on a specific interface with a custom Ethernet type.
//
// Added in macOS 10.15.
// Initializes an Ethernet channel on a specific interface with a custom Ethernet type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create(_:_:)
func nw_ethernet_channel_create(ether_type uint16, interface_ Nw_interface_t) Nw_ethernet_channel_t {
	return _nw_ethernet_channel_create(ether_type, interface_)
}

// nw_ethernet_channel_create_with_parameters is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create_with_parameters(_:_:_:)
func nw_ethernet_channel_create_with_parameters(ether_type uint16, interface_ Nw_interface_t, parameters Nw_parameters_t) Nw_ethernet_channel_t {
	return _nw_ethernet_channel_create_with_parameters(ether_type, interface_, parameters)
}

// nw_ethernet_channel_get_maximum_payload_size is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_get_maximum_payload_size(_:)
func nw_ethernet_channel_get_maximum_payload_size(ethernet_channel Nw_ethernet_channel_t) uint32 {
	return _nw_ethernet_channel_get_maximum_payload_size(ethernet_channel)
}

// Sends a single Ethernet frame over a channel to a specific Ethernet address.
//
// Added in macOS 10.15.
// Sends a single Ethernet frame over a channel to a specific Ethernet address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send(_:_:_:_:_:)
func nw_ethernet_channel_send(ethernet_channel Nw_ethernet_channel_t, content unsafe.Pointer, vlan_tag uint16, remote_address Nw_ethernet_address_t, completion unsafe.Pointer) {
	_nw_ethernet_channel_send(ethernet_channel, content, vlan_tag, remote_address, completion)
}

// Sets the queue on which all channel events are delivered.
//
// Added in macOS 10.15.
// Sets the queue on which all channel events are delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_queue(_:_:)
func nw_ethernet_channel_set_queue(ethernet_channel Nw_ethernet_channel_t, queue unsafe.Pointer) {
	_nw_ethernet_channel_set_queue(ethernet_channel, queue)
}

// Sets a handler to receive inbound Ethernet frames.
//
// Added in macOS 10.15.
// Sets a handler to receive inbound Ethernet frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_receive_handler(_:_:)
func nw_ethernet_channel_set_receive_handler(ethernet_channel Nw_ethernet_channel_t, handler unsafe.Pointer) {
	_nw_ethernet_channel_set_receive_handler(ethernet_channel, handler)
}

// Sets a handler to receive channel state updates.
//
// Added in macOS 10.15.
// Sets a handler to receive channel state updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_state_changed_handler(_:_:)
func nw_ethernet_channel_set_state_changed_handler(ethernet_channel Nw_ethernet_channel_t, handler unsafe.Pointer) {
	_nw_ethernet_channel_set_state_changed_handler(ethernet_channel, handler)
}

// Starts the process of registering the channel.
//
// Added in macOS 10.15.
// Starts the process of registering the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_start(_:)
func nw_ethernet_channel_start(ethernet_channel Nw_ethernet_channel_t) {
	_nw_ethernet_channel_start(ethernet_channel)
}

// Requests that a block be executed on the connection’s internal scheduling context.
//
// Added in macOS 10.15.
// Requests that a block be executed on the connection’s internal scheduling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_async(_:_:)
func nw_framer_async(framer Nw_framer_t, async_block unsafe.Pointer) {
	_nw_framer_async(framer, async_block)
}

// nw_framer_copy_options is a Network function.
//
// Added in macOS 12.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_options(_:)
func nw_framer_copy_options(framer Nw_framer_t) Nw_protocol_options_t {
	return _nw_framer_copy_options(framer)
}

// Accesses the parameters of the connection in which your protocol is running.
//
// Added in macOS 10.15.
// Accesses the parameters of the connection in which your protocol is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_parameters(_:)
func nw_framer_copy_parameters(framer Nw_framer_t) Nw_parameters_t {
	return _nw_framer_copy_parameters(framer)
}

// Accesses the remote endpoint of the connection in which your protocol is running.
//
// Added in macOS 10.15.
// Accesses the remote endpoint of the connection in which your protocol is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_remote_endpoint(_:)
func nw_framer_copy_remote_endpoint(framer Nw_framer_t) Nw_endpoint_t {
	return _nw_framer_copy_remote_endpoint(framer)
}

// Initializes a new protocol definition based on your protocol implementation.
//
// Added in macOS 10.15.
// Initializes a new protocol definition based on your protocol implementation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_create_definition(_:_:_:)
func nw_framer_create_definition(identifier unsafe.Pointer, flags uint32, start_handler unsafe.Pointer) Nw_protocol_definition_t {
	return _nw_framer_create_definition(identifier, flags, start_handler)
}

// Initializes a set of protocol options with a custom framer definition.
//
// Added in macOS 10.15.
// Initializes a set of protocol options with a custom framer definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_create_options(_:)
func nw_framer_create_options(framer_definition Nw_protocol_definition_t) Nw_protocol_options_t {
	return _nw_framer_create_options(framer_definition)
}

// Delivers an inbound message containing arbitrary data from your protocol to the application.
//
// Added in macOS 10.15.
// Delivers an inbound message containing arbitrary data from your protocol to the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_deliver_input(_:_:_:_:_:)
func nw_framer_deliver_input(framer Nw_framer_t, input_buffer unsafe.Pointer, input_length uintptr, message Nw_framer_message_t, is_complete bool) {
	_nw_framer_deliver_input(framer, input_buffer, input_length, message, is_complete)
}

// Delivers an inbound message containing a specific number of next received bytes.
//
// Added in macOS 10.15.
// Delivers an inbound message containing a specific number of next received bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_deliver_input_no_copy(_:_:_:_:)
func nw_framer_deliver_input_no_copy(framer Nw_framer_t, input_length uintptr, message Nw_framer_message_t, is_complete bool) bool {
	return _nw_framer_deliver_input_no_copy(framer, input_length, message, is_complete)
}

// Indicates to a connection that your protocol has encountered an error, or has gracefully closed.
//
// Added in macOS 10.15.
// Indicates to a connection that your protocol has encountered an error, or has gracefully closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_mark_failed_with_error(_:_:)
func nw_framer_mark_failed_with_error(framer Nw_framer_t, error_code int) {
	_nw_framer_mark_failed_with_error(framer, error_code)
}

// Indicates to a connection that your protocol’s handshake is complete.
//
// Added in macOS 10.15.
// Indicates to a connection that your protocol’s handshake is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_mark_ready(_:)
func nw_framer_mark_ready(framer Nw_framer_t) {
	_nw_framer_mark_ready(framer)
}

// Accesses a custom value stored in a framer message.
//
// Added in macOS 10.15.
// Accesses a custom value stored in a framer message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_access_value(_:_:_:)
func nw_framer_message_access_value(message Nw_framer_message_t, key unsafe.Pointer, access_value bool) bool {
	return _nw_framer_message_access_value(message, key, access_value)
}

// Accesses an NSObject value stored in a framer message.
//
// Added in macOS 10.15.
// Accesses an NSObject value stored in a framer message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_copy_object_value(_:_:)
func nw_framer_message_copy_object_value(message Nw_framer_message_t, key unsafe.Pointer) objc.ID {
	return _nw_framer_message_copy_object_value(message, key)
}

// Initializes an empty message from within a framer implementation.
//
// Added in macOS 10.15.
// Initializes an empty message from within a framer implementation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_create(_:)
func nw_framer_message_create(framer Nw_framer_t) Nw_framer_message_t {
	return _nw_framer_message_create(framer)
}

// Sets an NSObject value to be stored in a framer message.
//
// Added in macOS 10.15.
// Sets an NSObject value to be stored in a framer message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_set_object_value(_:_:_:)
func nw_framer_message_set_object_value(message Nw_framer_message_t, key unsafe.Pointer, value objc.ID) {
	_nw_framer_message_set_object_value(message, key, value)
}

// Sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released.
//
// Added in macOS 10.15.
// Sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_set_value(_:_:_:_:)
func nw_framer_message_set_value(message Nw_framer_message_t, key unsafe.Pointer, value unsafe.Pointer, dispose_value unsafe.Pointer) {
	_nw_framer_message_set_value(message, key, value, dispose_value)
}

// nw_framer_options_copy_object_value is a Network function.
//
// Added in macOS 12.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_options_copy_object_value(_:_:)
func nw_framer_options_copy_object_value(options Nw_protocol_options_t, key unsafe.Pointer) objc.ID {
	return _nw_framer_options_copy_object_value(options, key)
}

// nw_framer_options_set_object_value is a Network function.
//
// Added in macOS 12.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_options_set_object_value(_:_:_:)
func nw_framer_options_set_object_value(options Nw_protocol_options_t, key unsafe.Pointer, value objc.ID) {
	_nw_framer_options_set_object_value(options, key, value)
}

// Examines the content of input data while inside your input handler block.
//
// Added in macOS 10.15.
// Examines the content of input data while inside your input handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_parse_input(_:_:_:_:_:)
func nw_framer_parse_input(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer unsafe.Pointer, parse unsafe.Pointer) bool {
	return _nw_framer_parse_input(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
}

// Examines the content of output data while inside your output handler.
//
// Added in macOS 10.15.
// Examines the content of output data while inside your output handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_parse_output(_:_:_:_:_:)
func nw_framer_parse_output(framer Nw_framer_t, minimum_incomplete_length uintptr, maximum_length uintptr, temp_buffer unsafe.Pointer, parse unsafe.Pointer) bool {
	return _nw_framer_parse_output(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
}

// Indicates that your protocol no longer needs to handle input data.
//
// Added in macOS 10.15.
// Indicates that your protocol no longer needs to handle input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_pass_through_input(_:)
func nw_framer_pass_through_input(framer Nw_framer_t) {
	_nw_framer_pass_through_input(framer)
}

// Indicates that your protocol no longer needs to handle output data.
//
// Added in macOS 10.15.
// Indicates that your protocol no longer needs to handle output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_pass_through_output(_:)
func nw_framer_pass_through_output(framer Nw_framer_t) {
	_nw_framer_pass_through_output(framer)
}

// Dynamically adds another protocol that will run above your protocol after your protocol calls .
//
// Added in macOS 10.15.
// Dynamically adds another protocol that will run above your protocol after your protocol calls .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_prepend_application_protocol(_:_:)
func nw_framer_prepend_application_protocol(framer Nw_framer_t, protocol_options Nw_protocol_options_t) bool {
	return _nw_framer_prepend_application_protocol(framer, protocol_options)
}

// Initializes an empty message for a custom framer definition.
//
// Added in macOS 10.15.
// Initializes an empty message for a custom framer definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_protocol_create_message(_:)
func nw_framer_protocol_create_message(definition Nw_protocol_definition_t) Nw_framer_message_t {
	return _nw_framer_protocol_create_message(definition)
}

// Requests that the be called on your protocol at a specific time in the future.
//
// Added in macOS 10.15.
// Requests that the be called on your protocol at a specific time in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_schedule_wakeup(_:_:)
func nw_framer_schedule_wakeup(framer Nw_framer_t, milliseconds uint64) {
	_nw_framer_schedule_wakeup(framer, milliseconds)
}

// Sets a block to handle the final cleanup of allocations made by your protocol instance.
//
// Added in macOS 10.15.
// Sets a block to handle the final cleanup of allocations made by your protocol instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_cleanup_handler(_:_:)
func nw_framer_set_cleanup_handler(framer Nw_framer_t, cleanup_handler unsafe.Pointer) {
	_nw_framer_set_cleanup_handler(framer, cleanup_handler)
}

// Sets a block to handle new inbound data.
//
// Added in macOS 10.15.
// Sets a block to handle new inbound data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_input_handler(_:_:)
func nw_framer_set_input_handler(framer Nw_framer_t, input_handler unsafe.Pointer) {
	_nw_framer_set_input_handler(framer, input_handler)
}

// Sets a block to handle new outbound messages.
//
// Added in macOS 10.15.
// Sets a block to handle new outbound messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_output_handler(_:_:)
func nw_framer_set_output_handler(framer Nw_framer_t, output_handler unsafe.Pointer) {
	_nw_framer_set_output_handler(framer, output_handler)
}

// Sets a block to handle when the connection is being closed.
//
// Added in macOS 10.15.
// Sets a block to handle when the connection is being closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_stop_handler(_:_:)
func nw_framer_set_stop_handler(framer Nw_framer_t, stop_handler unsafe.Pointer) {
	_nw_framer_set_stop_handler(framer, stop_handler)
}

// Sets a handler to receive scheduled wakeup events.
//
// Added in macOS 10.15.
// Sets a handler to receive scheduled wakeup events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_wakeup_handler(_:_:)
func nw_framer_set_wakeup_handler(framer Nw_framer_t, wakeup_handler unsafe.Pointer) {
	_nw_framer_set_wakeup_handler(framer, wakeup_handler)
}

// Sends arbitrary output data in a buffer from your protocol to the next protocol.
//
// Added in macOS 10.15.
// Sends arbitrary output data in a buffer from your protocol to the next protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output(_:_:_:)
func nw_framer_write_output(framer Nw_framer_t, output_buffer unsafe.Pointer, output_length uintptr) {
	_nw_framer_write_output(framer, output_buffer, output_length)
}

// Sends arbitrary output data from your protocol to the next protocol.
//
// Added in macOS 10.15.
// Sends arbitrary output data from your protocol to the next protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output_data(_:_:)
func nw_framer_write_output_data(framer Nw_framer_t, output_data unsafe.Pointer) {
	_nw_framer_write_output_data(framer, output_data)
}

// Sends a specific number of bytes from a message while inside your output handler.
//
// Added in macOS 10.15.
// Sends a specific number of bytes from a message while inside your output handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output_no_copy(_:_:)
func nw_framer_write_output_no_copy(framer Nw_framer_t, output_length uintptr) bool {
	return _nw_framer_write_output_no_copy(framer, output_length)
}

// Adds a multicast address endpoint you specify to define an extra IP multicast group to join.
//
// Added in macOS 11.0.
// Adds a multicast address endpoint you specify to define an extra IP multicast group to join.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_add_endpoint(_:_:)
func nw_group_descriptor_add_endpoint(descriptor Nw_group_descriptor_t, endpoint Nw_endpoint_t) bool {
	return _nw_group_descriptor_add_endpoint(descriptor, endpoint)
}

// Creates group descriptor you use to join an IP multicast group on a local network.
//
// Added in macOS 11.0.
// Creates group descriptor you use to join an IP multicast group on a local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multicast(_:)
func nw_group_descriptor_create_multicast(multicast_group Nw_endpoint_t) Nw_group_descriptor_t {
	return _nw_group_descriptor_create_multicast(multicast_group)
}

// nw_group_descriptor_create_multiplex is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multiplex(_:)
func nw_group_descriptor_create_multiplex(remote_endpoint Nw_endpoint_t) Nw_group_descriptor_t {
	return _nw_group_descriptor_create_multiplex(remote_endpoint)
}

// Sets a handler to list all endpoints added to the group descriptor.
//
// Added in macOS 11.0.
// Sets a handler to list all endpoints added to the group descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints(_:_:)
func nw_group_descriptor_enumerate_endpoints(descriptor Nw_group_descriptor_t, enumerate_block unsafe.Pointer) {
	_nw_group_descriptor_enumerate_endpoints(descriptor, enumerate_block)
}

// Accesses the system interface index associated with the interface.
//
// Added in macOS 10.14.
// Accesses the system interface index associated with the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_index(_:)
func nw_interface_get_index(interface_ Nw_interface_t) uint32 {
	return _nw_interface_get_index(interface_)
}

// Accesses the name of the interface.
//
// Added in macOS 10.14.
// Accesses the name of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_name(_:)
func nw_interface_get_name(interface_ Nw_interface_t) unsafe.Pointer {
	return _nw_interface_get_name(interface_)
}

// Accesses the type of the interface, such as Wi-Fi or Loopback.
//
// Added in macOS 10.14.
// Accesses the type of the interface, such as Wi-Fi or Loopback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_type(_:)
func nw_interface_get_type(interface_ Nw_interface_t) unsafe.Pointer {
	return _nw_interface_get_type(interface_)
}

// Initializes an IP packet configuration with default settings.
//
// Added in macOS 10.14.
// Initializes an IP packet configuration with default settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_create_metadata()
func nw_ip_create_metadata() Nw_protocol_metadata_t {
	return _nw_ip_create_metadata()
}

// Checks the Explicit Congestion Notification flag value received on an IP packet.
//
// Added in macOS 10.14.
// Checks the Explicit Congestion Notification flag value received on an IP packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_ecn_flag(_:)
func nw_ip_metadata_get_ecn_flag(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_ip_metadata_get_ecn_flag(metadata)
}

// Access the time at which a packet was received, in nanoseconds, based on .
//
// Added in macOS 10.14.
// Access the time at which a packet was received, in nanoseconds, based on .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_receive_time(_:)
func nw_ip_metadata_get_receive_time(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_ip_metadata_get_receive_time(metadata)
}

// Accesses a specific service class to mark on an IP packet.
//
// Added in macOS 10.14.
// Accesses a specific service class to mark on an IP packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_service_class(_:)
func nw_ip_metadata_get_service_class(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_ip_metadata_get_service_class(metadata)
}

// Sets a specific Explicit Congestion Notification flag value to set on an IP packet.
//
// Added in macOS 10.14.
// Sets a specific Explicit Congestion Notification flag value to set on an IP packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_ecn_flag(_:_:)
func nw_ip_metadata_set_ecn_flag(metadata Nw_protocol_metadata_t, ecn_flag unsafe.Pointer) {
	_nw_ip_metadata_set_ecn_flag(metadata, ecn_flag)
}

// Sets a specific service class to mark on an IP packet.
//
// Added in macOS 10.14.
// Sets a specific service class to mark on an IP packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_service_class(_:_:)
func nw_ip_metadata_set_service_class(metadata Nw_protocol_metadata_t, service_class unsafe.Pointer) {
	_nw_ip_metadata_set_service_class(metadata, service_class)
}

// Configures a connection to deliver receive timestamps for IP packets.
//
// Added in macOS 10.14.
// Configures a connection to deliver receive timestamps for IP packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_calculate_receive_time(_:_:)
func nw_ip_options_set_calculate_receive_time(options Nw_protocol_options_t, calculate_receive_time bool) {
	_nw_ip_options_set_calculate_receive_time(options, calculate_receive_time)
}

// Configures a connection to disable fragmentation on outbound packets.
//
// Added in macOS 10.14.
// Configures a connection to disable fragmentation on outbound packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_fragmentation(_:_:)
func nw_ip_options_set_disable_fragmentation(options Nw_protocol_options_t, disable_fragmentation bool) {
	_nw_ip_options_set_disable_fragmentation(options, disable_fragmentation)
}

// nw_ip_options_set_disable_multicast_loopback is a Network function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_multicast_loopback(_:_:)
func nw_ip_options_set_disable_multicast_loopback(options Nw_protocol_options_t, disable_multicast_loopback bool) {
	_nw_ip_options_set_disable_multicast_loopback(options, disable_multicast_loopback)
}

// Configures the default hop limit for packets generated by a connection.
//
// Added in macOS 10.14.
// Configures the default hop limit for packets generated by a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_hop_limit(_:_:)
func nw_ip_options_set_hop_limit(options Nw_protocol_options_t, hop_limit uint8) {
	_nw_ip_options_set_hop_limit(options, hop_limit)
}

// Configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6.
//
// Added in macOS 10.14.
// Configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_use_minimum_mtu(_:_:)
func nw_ip_options_set_use_minimum_mtu(options Nw_protocol_options_t, use_minimum_mtu bool) {
	_nw_ip_options_set_use_minimum_mtu(options, use_minimum_mtu)
}

// Stops listening for inbound connections.
//
// Added in macOS 10.14.
// Stops listening for inbound connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_cancel(_:)
func nw_listener_cancel(listener Nw_listener_t) {
	_nw_listener_cancel(listener)
}

// Initializes a network listener, which will select a random port.
//
// Added in macOS 10.14.
// Initializes a network listener, which will select a random port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create(_:)
func nw_listener_create(parameters Nw_parameters_t) Nw_listener_t {
	return _nw_listener_create(parameters)
}

// Initializes a network listener to receive new streams on a multiplexed connection.
//
// Added in macOS 10.14.
// Initializes a network listener to receive new streams on a multiplexed connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_connection(_:_:)
func nw_listener_create_with_connection(connection Nw_connection_t, parameters Nw_parameters_t) Nw_listener_t {
	return _nw_listener_create_with_connection(connection, parameters)
}

// nw_listener_create_with_launchd_key is a Network function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_launchd_key(_:_:)
func nw_listener_create_with_launchd_key(parameters Nw_parameters_t, launchd_key unsafe.Pointer) Nw_listener_t {
	return _nw_listener_create_with_launchd_key(parameters, launchd_key)
}

// Initializes a network listener with a specified local port.
//
// Added in macOS 10.14.
// Initializes a network listener with a specified local port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_port(_:_:)
func nw_listener_create_with_port(port unsafe.Pointer, parameters Nw_parameters_t) Nw_listener_t {
	return _nw_listener_create_with_port(port, parameters)
}

// Checks the remaining number of inbound connections to deliver before rejecting connections.
//
// Added in macOS 10.15.
// Checks the remaining number of inbound connections to deliver before rejecting connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_get_new_connection_limit(_:)
func nw_listener_get_new_connection_limit(listener Nw_listener_t) uint32 {
	return _nw_listener_get_new_connection_limit(listener)
}

// The port on which the listener can accept connections.
//
// Added in macOS 10.14.
// The port on which the listener can accept connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_get_port(_:)
func nw_listener_get_port(listener Nw_listener_t) uint16 {
	return _nw_listener_get_port(listener)
}

// Sets a Bonjour service that advertises the listener on the local network.
//
// Added in macOS 10.14.
// Sets a Bonjour service that advertises the listener on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_advertise_descriptor(_:_:)
func nw_listener_set_advertise_descriptor(listener Nw_listener_t, advertise_descriptor Nw_advertise_descriptor_t) {
	_nw_listener_set_advertise_descriptor(listener, advertise_descriptor)
}

// Sets a handler that receives updates for the service endpoint being advertised.
//
// Added in macOS 10.14.
// Sets a handler that receives updates for the service endpoint being advertised.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_advertised_endpoint_changed_handler(_:_:)
func nw_listener_set_advertised_endpoint_changed_handler(listener Nw_listener_t, handler unsafe.Pointer) {
	_nw_listener_set_advertised_endpoint_changed_handler(listener, handler)
}

// nw_listener_set_new_connection_group_handler is a Network function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_group_handler(_:_:)
func nw_listener_set_new_connection_group_handler(listener Nw_listener_t, handler unsafe.Pointer) {
	_nw_listener_set_new_connection_group_handler(listener, handler)
}

// Sets a handler that receives inbound connections.
//
// Added in macOS 10.14.
// Sets a handler that receives inbound connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_handler(_:_:)
func nw_listener_set_new_connection_handler(listener Nw_listener_t, handler unsafe.Pointer) {
	_nw_listener_set_new_connection_handler(listener, handler)
}

// Resets the number of inbound connections to deliver before rejecting connections.
//
// Added in macOS 10.15.
// Resets the number of inbound connections to deliver before rejecting connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_limit(_:_:)
func nw_listener_set_new_connection_limit(listener Nw_listener_t, new_connection_limit uint32) {
	_nw_listener_set_new_connection_limit(listener, new_connection_limit)
}

// Sets the queue on which all listener events are delivered.
//
// Added in macOS 10.14.
// Sets the queue on which all listener events are delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_queue(_:_:)
func nw_listener_set_queue(listener Nw_listener_t, queue unsafe.Pointer) {
	_nw_listener_set_queue(listener, queue)
}

// Sets a handler to receive listener state updates.
//
// Added in macOS 10.14.
// Sets a handler to receive listener state updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_state_changed_handler(_:_:)
func nw_listener_set_state_changed_handler(listener Nw_listener_t, handler unsafe.Pointer) {
	_nw_listener_set_state_changed_handler(listener, handler)
}

// Registers for listening for inbound connections.
//
// Added in macOS 10.14.
// Registers for listening for inbound connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_start(_:)
func nw_listener_start(listener Nw_listener_t) {
	_nw_listener_start(listener)
}

// Checks a Boolean that indicates whether a connection group should reject unicast traffic.
//
// Added in macOS 11.0.
// Checks a Boolean that indicates whether a connection group should reject unicast traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_get_disable_unicast_traffic(_:)
func nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t) bool {
	return _nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor)
}

// Sets a Boolean that indicates whether a connection group should reject unicast traffic.
//
// Added in macOS 11.0.
// Sets a Boolean that indicates whether a connection group should reject unicast traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_disable_unicast_traffic(_:_:)
func nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor Nw_group_descriptor_t, disable_unicast_traffic bool) {
	_nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor, disable_unicast_traffic)
}

// Sets an optional address endpoint used to filter received multicast packets.
//
// Added in macOS 11.0.
// Sets an optional address endpoint used to filter received multicast packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_specific_source(_:_:)
func nw_multicast_group_descriptor_set_specific_source(multicast_descriptor Nw_group_descriptor_t, source Nw_endpoint_t) {
	_nw_multicast_group_descriptor_set_specific_source(multicast_descriptor, source)
}

// Removes all prohibited interface types.
//
// Added in macOS 10.14.
// Removes all prohibited interface types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interface_types(_:)
func nw_parameters_clear_prohibited_interface_types(parameters Nw_parameters_t) {
	_nw_parameters_clear_prohibited_interface_types(parameters)
}

// Removes all prohibited interface types.
//
// Added in macOS 10.14.
// Removes all prohibited interface types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interfaces(_:)
func nw_parameters_clear_prohibited_interfaces(parameters Nw_parameters_t) {
	_nw_parameters_clear_prohibited_interfaces(parameters)
}

// Peforms a deep copy of a parameters object.
//
// Added in macOS 10.14.
// Peforms a deep copy of a parameters object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy(_:)
func nw_parameters_copy(parameters Nw_parameters_t) Nw_parameters_t {
	return _nw_parameters_copy(parameters)
}

// Accesses the protocol stack used by connections and listeners.
//
// Added in macOS 10.14.
// Accesses the protocol stack used by connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy_default_protocol_stack(_:)
func nw_parameters_copy_default_protocol_stack(parameters Nw_parameters_t) Nw_protocol_stack_t {
	return _nw_parameters_copy_default_protocol_stack(parameters)
}

// Accesses the interface required on connections, listeners, and browsers.
//
// Added in macOS 10.14.
// Accesses the interface required on connections, listeners, and browsers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy_required_interface(_:)
func nw_parameters_copy_required_interface(parameters Nw_parameters_t) Nw_interface_t {
	return _nw_parameters_copy_required_interface(parameters)
}

// Initializes parameters for connections, listeners, and browsers with no protocols specified.
//
// Added in macOS 10.14.
// Initializes parameters for connections, listeners, and browsers with no protocols specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create()
func nw_parameters_create() Nw_parameters_t {
	return _nw_parameters_create()
}

// nw_parameters_create_application_service is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_application_service()
func nw_parameters_create_application_service() Nw_parameters_t {
	return _nw_parameters_create_application_service()
}

// Initializes parameters for connections and listeners using a custom IP protocol.
//
// Added in macOS 10.15.
// Initializes parameters for connections and listeners using a custom IP protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_custom_ip(_:_:)
func nw_parameters_create_custom_ip(custom_ip_protocol_number uint8, configure_ip unsafe.Pointer) Nw_parameters_t {
	return _nw_parameters_create_custom_ip(custom_ip_protocol_number, configure_ip)
}

// Initializes parameters for QUIC connections and listeners.
//
// Added in macOS 12.0.
// Initializes parameters for QUIC connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_quic(_:)
func nw_parameters_create_quic(configure_quic unsafe.Pointer) Nw_parameters_t {
	return _nw_parameters_create_quic(configure_quic)
}

// Initializes parameters for TLS or TCP connections and listeners.
//
// Added in macOS 10.14.
// Initializes parameters for TLS or TCP connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_tcp(_:_:)
func nw_parameters_create_secure_tcp(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) Nw_parameters_t {
	return _nw_parameters_create_secure_tcp(configure_tls, configure_tcp)
}

// Initializes parameters for DTLS or UDP connections and listeners.
//
// Added in macOS 10.14.
// Initializes parameters for DTLS or UDP connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_udp(_:_:)
func nw_parameters_create_secure_udp(configure_dtls unsafe.Pointer, configure_udp unsafe.Pointer) Nw_parameters_t {
	return _nw_parameters_create_secure_udp(configure_dtls, configure_udp)
}

// nw_parameters_get_allow_ultra_constrained is a Network function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_allow_ultra_constrained(_:)
func nw_parameters_get_allow_ultra_constrained(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_allow_ultra_constrained(parameters)
}

// Gets a flag that indicates whether the network request originates from the developer or the user.
//
// Added in macOS 12.0.
// Gets a flag that indicates whether the network request originates from the developer or the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_attribution(_:)
func nw_parameters_get_attribution(parameters Nw_parameters_t) nw_parameters_attribution_t {
	return _nw_parameters_get_attribution(parameters)
}

// Checks the behavior for how expired DNS answers should be used.
//
// Added in macOS 10.14.
// Checks the behavior for how expired DNS answers should be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_expired_dns_behavior(_:)
func nw_parameters_get_expired_dns_behavior(parameters Nw_parameters_t) unsafe.Pointer {
	return _nw_parameters_get_expired_dns_behavior(parameters)
}

// Checks if sending application data with protocol handshakes is enabled.
//
// Added in macOS 10.14.
// Checks if sending application data with protocol handshakes is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_fast_open_enabled(_:)
func nw_parameters_get_fast_open_enabled(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_fast_open_enabled(parameters)
}

// Checks whether a connection is allowed to use peer-to-peer link technologies.
//
// Added in macOS 10.14.
// Checks whether a connection is allowed to use peer-to-peer link technologies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_include_peer_to_peer(_:)
func nw_parameters_get_include_peer_to_peer(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_include_peer_to_peer(parameters)
}

// Checks if multipath is enabled on a connection.
//
// Added in macOS 10.14.
// Checks if multipath is enabled on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_multipath_service(_:)
func nw_parameters_get_multipath_service(parameters Nw_parameters_t) unsafe.Pointer {
	return _nw_parameters_get_multipath_service(parameters)
}

// Checks if proxies are ignored by default.
//
// Added in macOS 10.14.
// Checks if proxies are ignored by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prefer_no_proxy(_:)
func nw_parameters_get_prefer_no_proxy(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_prefer_no_proxy(parameters)
}

// Checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode.
//
// Added in macOS 10.15.
// Checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_constrained(_:)
func nw_parameters_get_prohibit_constrained(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_prohibit_constrained(parameters)
}

// Checks if connections, listeners, and browsers are prevented from using network paths marked as expensive.
//
// Added in macOS 10.14.
// Checks if connections, listeners, and browsers are prevented from using network paths marked as expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_expensive(_:)
func nw_parameters_get_prohibit_expensive(parameters Nw_parameters_t) bool {
	return _nw_parameters_get_prohibit_expensive(parameters)
}

// Accesses the interface type required on connections and listeners.
//
// Added in macOS 10.14.
// Accesses the interface type required on connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_required_interface_type(_:)
func nw_parameters_get_required_interface_type(parameters Nw_parameters_t) unsafe.Pointer {
	return _nw_parameters_get_required_interface_type(parameters)
}

// Checks the level of service quality used for connections.
//
// Added in macOS 10.14.
// Checks the level of service quality used for connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_service_class(_:)
func nw_parameters_get_service_class(parameters Nw_parameters_t) unsafe.Pointer {
	return _nw_parameters_get_service_class(parameters)
}

// Examines the list of prohibited interface types.
//
// Added in macOS 10.14.
// Examines the list of prohibited interface types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interface_types(_:_:)
func nw_parameters_iterate_prohibited_interface_types(parameters Nw_parameters_t, iterate_block unsafe.Pointer) {
	_nw_parameters_iterate_prohibited_interface_types(parameters, iterate_block)
}

// Examines the list of prohibited interfaces.
//
// Added in macOS 10.14.
// Examines the list of prohibited interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interfaces(_:_:)
func nw_parameters_iterate_prohibited_interfaces(parameters Nw_parameters_t, iterate_block unsafe.Pointer) {
	_nw_parameters_iterate_prohibited_interfaces(parameters, iterate_block)
}

// Prevents connections and listeners from using a specific interface.
//
// Added in macOS 10.14.
// Prevents connections and listeners from using a specific interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface(_:_:)
func nw_parameters_prohibit_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	_nw_parameters_prohibit_interface(parameters, interface_)
}

// Prevents connections, listeners, and browsers from using a specific interface type.
//
// Added in macOS 10.14.
// Prevents connections, listeners, and browsers from using a specific interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface_type(_:_:)
func nw_parameters_prohibit_interface_type(parameters Nw_parameters_t, interface_type unsafe.Pointer) {
	_nw_parameters_prohibit_interface_type(parameters, interface_type)
}

// Sets a specific interface to require on connections, listeners, and browsers.
//
// Added in macOS 10.14.
// Sets a specific interface to require on connections, listeners, and browsers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_require_interface(_:_:)
func nw_parameters_require_interface(parameters Nw_parameters_t, interface_ Nw_interface_t) {
	_nw_parameters_require_interface(parameters, interface_)
}

// Checks whether a connection requires DNSSEC validation when resolving endpoints.
//
// Added in macOS 13.0.
// Checks whether a connection requires DNSSEC validation when resolving endpoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_requires_dnssec_validation(_:)
func nw_parameters_requires_dnssec_validation(parameters Nw_parameters_t) bool {
	return _nw_parameters_requires_dnssec_validation(parameters)
}

// nw_parameters_set_allow_ultra_constrained is a Network function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_allow_ultra_constrained(_:_:)
func nw_parameters_set_allow_ultra_constrained(parameters Nw_parameters_t, allow_ultra_constrained bool) {
	_nw_parameters_set_allow_ultra_constrained(parameters, allow_ultra_constrained)
}

// Sets a flag that indicates whether the network request originates from the developer or the user.
//
// Added in macOS 12.0.
// Sets a flag that indicates whether the network request originates from the developer or the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_attribution(_:_:)
func nw_parameters_set_attribution(parameters Nw_parameters_t, attribution nw_parameters_attribution_t) {
	_nw_parameters_set_attribution(parameters, attribution)
}

// Sets the behavior for how expired DNS answers should be used.
//
// Added in macOS 10.14.
// Sets the behavior for how expired DNS answers should be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_expired_dns_behavior(_:_:)
func nw_parameters_set_expired_dns_behavior(parameters Nw_parameters_t, expired_dns_behavior unsafe.Pointer) {
	_nw_parameters_set_expired_dns_behavior(parameters, expired_dns_behavior)
}

// Enables sending application data with protocol handshakes.
//
// Added in macOS 10.14.
// Enables sending application data with protocol handshakes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_fast_open_enabled(_:_:)
func nw_parameters_set_fast_open_enabled(parameters Nw_parameters_t, fast_open_enabled bool) {
	_nw_parameters_set_fast_open_enabled(parameters, fast_open_enabled)
}

// Enables peer-to-peer link technologies for connections and listeners.
//
// Added in macOS 10.14.
// Enables peer-to-peer link technologies for connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_include_peer_to_peer(_:_:)
func nw_parameters_set_include_peer_to_peer(parameters Nw_parameters_t, include_peer_to_peer bool) {
	_nw_parameters_set_include_peer_to_peer(parameters, include_peer_to_peer)
}

// Enables multipath protocols to allow connections to use multiple interfaces.
//
// Added in macOS 10.14.
// Enables multipath protocols to allow connections to use multiple interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_multipath_service(_:_:)
func nw_parameters_set_multipath_service(parameters Nw_parameters_t, multipath_service unsafe.Pointer) {
	_nw_parameters_set_multipath_service(parameters, multipath_service)
}

// Sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system.
//
// Added in macOS 10.14.
// Sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prefer_no_proxy(_:_:)
func nw_parameters_set_prefer_no_proxy(parameters Nw_parameters_t, prefer_no_proxy bool) {
	_nw_parameters_set_prefer_no_proxy(parameters, prefer_no_proxy)
}

// Associates a privacy context with any connections or listeners that use the parameters.
//
// Added in macOS 11.0.
// Associates a privacy context with any connections or listeners that use the parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_privacy_context(_:_:)
func nw_parameters_set_privacy_context(parameters Nw_parameters_t, privacy_context Nw_privacy_context_t) {
	_nw_parameters_set_privacy_context(parameters, privacy_context)
}

// Prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode.
//
// Added in macOS 10.15.
// Prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_constrained(_:_:)
func nw_parameters_set_prohibit_constrained(parameters Nw_parameters_t, prohibit_constrained bool) {
	_nw_parameters_set_prohibit_constrained(parameters, prohibit_constrained)
}

// Prevents connections, listeners, and browsers from using network paths marked as expensive.
//
// Added in macOS 10.14.
// Prevents connections, listeners, and browsers from using network paths marked as expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_expensive(_:_:)
func nw_parameters_set_prohibit_expensive(parameters Nw_parameters_t, prohibit_expensive bool) {
	_nw_parameters_set_prohibit_expensive(parameters, prohibit_expensive)
}

// Sets an interface type to require on connections and listeners.
//
// Added in macOS 10.14.
// Sets an interface type to require on connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_required_interface_type(_:_:)
func nw_parameters_set_required_interface_type(parameters Nw_parameters_t, interface_type unsafe.Pointer) {
	_nw_parameters_set_required_interface_type(parameters, interface_type)
}

// Determines whether a connection requires DNSSEC validation when resolving endpoints.
//
// Added in macOS 13.0.
// Determines whether a connection requires DNSSEC validation when resolving endpoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_requires_dnssec_validation(_:_:)
func nw_parameters_set_requires_dnssec_validation(parameters Nw_parameters_t, requires_dnssec_validation bool) {
	_nw_parameters_set_requires_dnssec_validation(parameters, requires_dnssec_validation)
}

// Sets a level of service quality to use for connections.
//
// Added in macOS 10.14.
// Sets a level of service quality to use for connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_service_class(_:_:)
func nw_parameters_set_service_class(parameters Nw_parameters_t, service_class unsafe.Pointer) {
	_nw_parameters_set_service_class(parameters, service_class)
}

// Accesses the remote endpoint in use by a connection’s network path.
//
// Added in macOS 10.14.
// Accesses the remote endpoint in use by a connection’s network path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_copy_effective_remote_endpoint(_:)
func nw_path_copy_effective_remote_endpoint(path Nw_path_t) Nw_endpoint_t {
	return _nw_path_copy_effective_remote_endpoint(path)
}

// Enumerates the list of gateways configured on the interfaces available to a path.
//
// Added in macOS 10.15.
// Enumerates the list of gateways configured on the interfaces available to a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways(_:_:)
func nw_path_enumerate_gateways(path Nw_path_t, enumerate_block unsafe.Pointer) {
	_nw_path_enumerate_gateways(path, enumerate_block)
}

// Enumerates the list of all interfaces available to the path, in order of preference.
//
// Added in macOS 10.14.
// Enumerates the list of all interfaces available to the path, in order of preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces(_:_:)
func nw_path_enumerate_interfaces(path Nw_path_t, enumerate_block unsafe.Pointer) {
	_nw_path_enumerate_interfaces(path, enumerate_block)
}

// nw_path_get_link_quality is a Network function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_link_quality(_:)
func nw_path_get_link_quality(path Nw_path_t) unsafe.Pointer {
	return _nw_path_get_link_quality(path)
}

// Checks whether a path can be used by connections.
//
// Added in macOS 10.14.
// Checks whether a path can be used by connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_status(_:)
func nw_path_get_status(path Nw_path_t) unsafe.Pointer {
	return _nw_path_get_status(path)
}

// nw_path_get_unsatisfied_reason is a Network function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_unsatisfied_reason(_:)
func nw_path_get_unsatisfied_reason(path Nw_path_t) unsafe.Pointer {
	return _nw_path_get_unsatisfied_reason(path)
}

// Checks whether the path has a DNS server configured.
//
// Added in macOS 10.14.
// Checks whether the path has a DNS server configured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_dns(_:)
func nw_path_has_dns(path Nw_path_t) bool {
	return _nw_path_has_dns(path)
}

// Checks whether the path can route IPv4 traffic.
//
// Added in macOS 10.14.
// Checks whether the path can route IPv4 traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_ipv4(_:)
func nw_path_has_ipv4(path Nw_path_t) bool {
	return _nw_path_has_ipv4(path)
}

// Checks whether the path can route IPv6 traffic.
//
// Added in macOS 10.14.
// Checks whether the path can route IPv6 traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_ipv6(_:)
func nw_path_has_ipv6(path Nw_path_t) bool {
	return _nw_path_has_ipv6(path)
}

// Checks whether the path uses an interface in Low Data Mode.
//
// Added in macOS 10.15.
// Checks whether the path uses an interface in Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_constrained(_:)
func nw_path_is_constrained(path Nw_path_t) bool {
	return _nw_path_is_constrained(path)
}

// Compares if two paths are identical.
//
// Added in macOS 10.14.
// Compares if two paths are identical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_equal(_:_:)
func nw_path_is_equal(path Nw_path_t, other_path Nw_path_t) bool {
	return _nw_path_is_equal(path, other_path)
}

// Checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot.
//
// Added in macOS 10.14.
// Checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_expensive(_:)
func nw_path_is_expensive(path Nw_path_t) bool {
	return _nw_path_is_expensive(path)
}

// nw_path_is_ultra_constrained is a Network function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_ultra_constrained(_:)
func nw_path_is_ultra_constrained(path Nw_path_t) bool {
	return _nw_path_is_ultra_constrained(path)
}

// Stops receiving network path updates.
//
// Added in macOS 10.14.
// Stops receiving network path updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel(_:)
func nw_path_monitor_cancel(monitor Nw_path_monitor_t) {
	_nw_path_monitor_cancel(monitor)
}

// Initializes a path monitor to observe all available interface types.
//
// Added in macOS 10.14.
// Initializes a path monitor to observe all available interface types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create()
func nw_path_monitor_create() Nw_path_monitor_t {
	return _nw_path_monitor_create()
}

// nw_path_monitor_create_for_ethernet_channel is a Network function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create_for_ethernet_channel()
func nw_path_monitor_create_for_ethernet_channel() Nw_path_monitor_t {
	return _nw_path_monitor_create_for_ethernet_channel()
}

// Initializes a path monitor to observe a specific interface type.
//
// Added in macOS 10.14.
// Initializes a path monitor to observe a specific interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create_with_type(_:)
func nw_path_monitor_create_with_type(required_interface_type unsafe.Pointer) Nw_path_monitor_t {
	return _nw_path_monitor_create_with_type(required_interface_type)
}

// Prohibit a path monitor from using a specific interface type.
//
// Added in macOS 11.0.
// Prohibit a path monitor from using a specific interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_prohibit_interface_type(_:_:)
func nw_path_monitor_prohibit_interface_type(monitor Nw_path_monitor_t, interface_type unsafe.Pointer) {
	_nw_path_monitor_prohibit_interface_type(monitor, interface_type)
}

// Sets a handler to determine when a monitor is fully cancelled and will no longer deliver events.
//
// Added in macOS 10.14.
// Sets a handler to determine when a monitor is fully cancelled and will no longer deliver events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_cancel_handler(_:_:)
func nw_path_monitor_set_cancel_handler(monitor Nw_path_monitor_t, cancel_handler unsafe.Pointer) {
	_nw_path_monitor_set_cancel_handler(monitor, cancel_handler)
}

// Sets a queue on which to deliver path events.
//
// Added in macOS 10.14.
// Sets a queue on which to deliver path events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_queue(_:_:)
func nw_path_monitor_set_queue(monitor Nw_path_monitor_t, queue unsafe.Pointer) {
	_nw_path_monitor_set_queue(monitor, queue)
}

// Sets a handler to receive network path updates.
//
// Added in macOS 10.14.
// Sets a handler to receive network path updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_update_handler(_:_:)
func nw_path_monitor_set_update_handler(monitor Nw_path_monitor_t, update_handler unsafe.Pointer) {
	_nw_path_monitor_set_update_handler(monitor, update_handler)
}

// Starts monitoring path changes.
//
// Added in macOS 10.14.
// Starts monitoring path changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_start(_:)
func nw_path_monitor_start(monitor Nw_path_monitor_t) {
	_nw_path_monitor_start(monitor)
}

// Checks if connections using the path may send traffic over a specific interface type.
//
// Added in macOS 10.14.
// Checks if connections using the path may send traffic over a specific interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_uses_interface_type(_:_:)
func nw_path_uses_interface_type(path Nw_path_t, interface_type unsafe.Pointer) bool {
	return _nw_path_uses_interface_type(path, interface_type)
}

// Applies a proxy configuration to all connections associated with this context.
//
// Added in macOS 14.0.
// Applies a proxy configuration to all connections associated with this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_add_proxy(_:_:)
func nw_privacy_context_add_proxy(privacy_context Nw_privacy_context_t, proxy_config Nw_proxy_config_t) {
	_nw_privacy_context_add_proxy(privacy_context, proxy_config)
}

// Clears out any proxies added using
//
// Added in macOS 14.0.
// Clears out any proxies added using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_clear_proxies(_:)
func nw_privacy_context_clear_proxies(privacy_context Nw_privacy_context_t) {
	_nw_privacy_context_clear_proxies(privacy_context)
}

// Initializes a privacy context with a description string.
//
// Added in macOS 11.0.
// Initializes a privacy context with a description string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_create(_:)
func nw_privacy_context_create(description unsafe.Pointer) Nw_privacy_context_t {
	return _nw_privacy_context_create(description)
}

// Disables system logging of connection activity.
//
// Added in macOS 11.0.
// Disables system logging of connection activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_disable_logging(_:)
func nw_privacy_context_disable_logging(privacy_context Nw_privacy_context_t) {
	_nw_privacy_context_disable_logging(privacy_context)
}

// Flushes all cached data, such as TLS session state, created by connections associated with the privacy context.
//
// Added in macOS 11.0.
// Flushes all cached data, such as TLS session state, created by connections associated with the privacy context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_flush_cache(_:)
func nw_privacy_context_flush_cache(privacy_context Nw_privacy_context_t) {
	_nw_privacy_context_flush_cache(privacy_context)
}

// Requires that any DNS name resolution for connections associated with this context use encrypted transports, such as TLS or HTTPS.
//
// Added in macOS 11.0.
// Requires that any DNS name resolution for connections associated with this context use encrypted transports, such as TLS or HTTPS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_privacy_context_require_encrypted_name_resolution(_:_:_:)
func nw_privacy_context_require_encrypted_name_resolution(privacy_context Nw_privacy_context_t, require_encrypted_name_resolution bool, fallback_resolver_config Nw_resolver_config_t) {
	_nw_privacy_context_require_encrypted_name_resolution(privacy_context, require_encrypted_name_resolution, fallback_resolver_config)
}

// Accesses the system definition of the Internet Protocol.
//
// Added in macOS 10.14.
// Accesses the system definition of the Internet Protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_ip_definition()
func nw_protocol_copy_ip_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_ip_definition()
}

// Accesses the system definition of the QUIC transport protocol.
//
// Added in macOS 12.0.
// Accesses the system definition of the QUIC transport protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_quic_definition()
func nw_protocol_copy_quic_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_quic_definition()
}

// Accesses the system definition of the Transport Control Protocol.
//
// Added in macOS 10.14.
// Accesses the system definition of the Transport Control Protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_tcp_definition()
func nw_protocol_copy_tcp_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_tcp_definition()
}

// Accesses the system definition of the Transport Layer Security protocol.
//
// Added in macOS 10.14.
// Accesses the system definition of the Transport Layer Security protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_tls_definition()
func nw_protocol_copy_tls_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_tls_definition()
}

// Accesses the system definition of the User Datagram Protocol.
//
// Added in macOS 10.14.
// Accesses the system definition of the User Datagram Protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_udp_definition()
func nw_protocol_copy_udp_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_udp_definition()
}

// Accesses the system definition of the WebSocket protocol.
//
// Added in macOS 10.15.
// Accesses the system definition of the WebSocket protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_ws_definition()
func nw_protocol_copy_ws_definition() Nw_protocol_definition_t {
	return _nw_protocol_copy_ws_definition()
}

// Compares two protocol definitions, and returns true if they represent the same protocol implementation.
//
// Added in macOS 10.14.
// Compares two protocol definitions, and returns true if they represent the same protocol implementation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_definition_is_equal(_:_:)
func nw_protocol_definition_is_equal(definition1 Nw_protocol_definition_t, definition2 Nw_protocol_definition_t) bool {
	return _nw_protocol_definition_is_equal(definition1, definition2)
}

// Accesses the protocol definition associated with the metadata object.
//
// Added in macOS 10.14.
// Accesses the protocol definition associated with the metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_copy_definition(_:)
func nw_protocol_metadata_copy_definition(metadata Nw_protocol_metadata_t) Nw_protocol_definition_t {
	return _nw_protocol_metadata_copy_definition(metadata)
}

// Checks if a metadata object represents a custom framer protocol message.
//
// Added in macOS 10.15.
// Checks if a metadata object represents a custom framer protocol message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_framer_message(_:)
func nw_protocol_metadata_is_framer_message(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_framer_message(metadata)
}

// Checks whether a metadata object represents an IP packet.
//
// Added in macOS 10.14.
// Checks whether a metadata object represents an IP packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ip(_:)
func nw_protocol_metadata_is_ip(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_ip(metadata)
}

// Checks whether a metadata object contains QUIC connection state.
//
// Added in macOS 12.0.
// Checks whether a metadata object contains QUIC connection state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_quic(_:)
func nw_protocol_metadata_is_quic(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_quic(metadata)
}

// Checks whether a metadata object contains TCP connection state.
//
// Added in macOS 10.14.
// Checks whether a metadata object contains TCP connection state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tcp(_:)
func nw_protocol_metadata_is_tcp(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_tcp(metadata)
}

// Checks whether a metadata object contains TLS connection state.
//
// Added in macOS 10.14.
// Checks whether a metadata object contains TLS connection state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tls(_:)
func nw_protocol_metadata_is_tls(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_tls(metadata)
}

// Checks whether a metadata object represents a UDP datagram.
//
// Added in macOS 10.14.
// Checks whether a metadata object represents a UDP datagram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_udp(_:)
func nw_protocol_metadata_is_udp(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_udp(metadata)
}

// Checks whether a metadata object represents a WebSocket message.
//
// Added in macOS 10.15.
// Checks whether a metadata object represents a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ws(_:)
func nw_protocol_metadata_is_ws(metadata Nw_protocol_metadata_t) bool {
	return _nw_protocol_metadata_is_ws(metadata)
}

// Accesses the protocol definition associated with the options object.
//
// Added in macOS 10.14.
// Accesses the protocol definition associated with the options object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_options_copy_definition(_:)
func nw_protocol_options_copy_definition(options Nw_protocol_options_t) Nw_protocol_definition_t {
	return _nw_protocol_options_copy_definition(options)
}

// Checks whether an options object uses the QUIC protocol.
//
// Added in macOS 12.0.
// Checks whether an options object uses the QUIC protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_options_is_quic(_:)
func nw_protocol_options_is_quic(options Nw_protocol_options_t) bool {
	return _nw_protocol_options_is_quic(options)
}

// Removes all application protocols from the protocol stack.
//
// Added in macOS 10.14.
// Removes all application protocols from the protocol stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_clear_application_protocols(_:)
func nw_protocol_stack_clear_application_protocols(stack Nw_protocol_stack_t) {
	_nw_protocol_stack_clear_application_protocols(stack)
}

// Accesses the protocol stack’s Internet Protocol options.
//
// Added in macOS 10.14.
// Accesses the protocol stack’s Internet Protocol options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_internet_protocol(_:)
func nw_protocol_stack_copy_internet_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	return _nw_protocol_stack_copy_internet_protocol(stack)
}

// Accesses the options for the protocol stack’s transport protocol.
//
// Added in macOS 10.14.
// Accesses the options for the protocol stack’s transport protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_transport_protocol(_:)
func nw_protocol_stack_copy_transport_protocol(stack Nw_protocol_stack_t) Nw_protocol_options_t {
	return _nw_protocol_stack_copy_transport_protocol(stack)
}

// Iterates through the array of application protocol options that will be used by connections and listeners.
//
// Added in macOS 10.14.
// Iterates through the array of application protocol options that will be used by connections and listeners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_application_protocols(_:_:)
func nw_protocol_stack_iterate_application_protocols(stack Nw_protocol_stack_t, iterate_block unsafe.Pointer) {
	_nw_protocol_stack_iterate_application_protocols(stack, iterate_block)
}

// Adds a protocol onto the top of the protocol stack.
//
// Added in macOS 10.14.
// Adds a protocol onto the top of the protocol stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_prepend_application_protocol(_:_:)
func nw_protocol_stack_prepend_application_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	_nw_protocol_stack_prepend_application_protocol(stack, protocol_)
}

// Replaces the protocol stack’s transport protocol with a new set of options.
//
// Added in macOS 10.14.
// Replaces the protocol stack’s transport protocol with a new set of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_set_transport_protocol(_:_:)
func nw_protocol_stack_set_transport_protocol(stack Nw_protocol_stack_t, protocol_ Nw_protocol_options_t) {
	_nw_protocol_stack_set_transport_protocol(stack, protocol_)
}

// nw_proxy_config_add_excluded_domain is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_add_excluded_domain(_:_:)
func nw_proxy_config_add_excluded_domain(config Nw_proxy_config_t, excluded_domain unsafe.Pointer) {
	_nw_proxy_config_add_excluded_domain(config, excluded_domain)
}

// nw_proxy_config_add_match_domain is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_add_match_domain(_:_:)
func nw_proxy_config_add_match_domain(config Nw_proxy_config_t, match_domain unsafe.Pointer) {
	_nw_proxy_config_add_match_domain(config, match_domain)
}

// nw_proxy_config_clear_excluded_domains is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_excluded_domains(_:)
func nw_proxy_config_clear_excluded_domains(config Nw_proxy_config_t) {
	_nw_proxy_config_clear_excluded_domains(config)
}

// nw_proxy_config_clear_match_domains is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_match_domains(_:)
func nw_proxy_config_clear_match_domains(config Nw_proxy_config_t) {
	_nw_proxy_config_clear_match_domains(config)
}

// Initializes a legacy HTTP CONNECT configuration for a proxy server accessible using HTTP/1.1.
//
// Added in macOS 14.0.
// Initializes a legacy HTTP CONNECT configuration for a proxy server accessible using HTTP/1.1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_create_http_connect(_:_:)
func nw_proxy_config_create_http_connect(proxy_endpoint Nw_endpoint_t, proxy_tls_options Nw_protocol_options_t) Nw_proxy_config_t {
	return _nw_proxy_config_create_http_connect(proxy_endpoint, proxy_tls_options)
}

// Initializes an Oblivious HTTP proxy configuration using a relay and a gateway.
//
// Added in macOS 14.0.
// Initializes an Oblivious HTTP proxy configuration using a relay and a gateway.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_create_oblivious_http(_:_:_:_:)
func nw_proxy_config_create_oblivious_http(relay Nw_relay_hop_t, relay_resource_path unsafe.Pointer, gateway_key_config unsafe.Pointer, gateway_key_config_length uintptr) Nw_proxy_config_t {
	return _nw_proxy_config_create_oblivious_http(relay, relay_resource_path, gateway_key_config, gateway_key_config_length)
}

// Initializes a proxy configuration with one or two relay hops.
//
// Added in macOS 14.0.
// Initializes a proxy configuration with one or two relay hops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_create_relay(_:_:)
func nw_proxy_config_create_relay(first_hop Nw_relay_hop_t, second_hop Nw_relay_hop_t) Nw_proxy_config_t {
	return _nw_proxy_config_create_relay(first_hop, second_hop)
}

// Initializes a SOCKSv5 proxy configuration.
//
// Added in macOS 14.0.
// Initializes a SOCKSv5 proxy configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_create_socksv5(_:)
func nw_proxy_config_create_socksv5(proxy_endpoint Nw_endpoint_t) Nw_proxy_config_t {
	return _nw_proxy_config_create_socksv5(proxy_endpoint)
}

// nw_proxy_config_enumerate_excluded_domains is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_excluded_domains(_:_:)
func nw_proxy_config_enumerate_excluded_domains(config Nw_proxy_config_t, enumerator unsafe.Pointer) {
	_nw_proxy_config_enumerate_excluded_domains(config, enumerator)
}

// nw_proxy_config_enumerate_match_domains is a Network function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_match_domains(_:_:)
func nw_proxy_config_enumerate_match_domains(config Nw_proxy_config_t, enumerator unsafe.Pointer) {
	_nw_proxy_config_enumerate_match_domains(config, enumerator)
}

// Checks if a proxy configuration allows failover to non-proxied connections.
//
// Added in macOS 14.0.
// Checks if a proxy configuration allows failover to non-proxied connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_get_failover_allowed(_:)
func nw_proxy_config_get_failover_allowed(proxy_config Nw_proxy_config_t) bool {
	return _nw_proxy_config_get_failover_allowed(proxy_config)
}

// Configures whether or not a proxy configuration allows failover to non-proxied connections. Failover isn’t allowed by default.
//
// Added in macOS 14.0.
// Configures whether or not a proxy configuration allows failover to non-proxied connections. Failover isn’t allowed by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_set_failover_allowed(_:_:)
func nw_proxy_config_set_failover_allowed(proxy_config Nw_proxy_config_t, failover_allowed bool) {
	_nw_proxy_config_set_failover_allowed(proxy_config, failover_allowed)
}

// Sets a username and password to use as authentication for a proxy configuration.
//
// Added in macOS 14.0.
// Sets a username and password to use as authentication for a proxy configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_set_username_and_password(_:_:_:)
func nw_proxy_config_set_username_and_password(proxy_config Nw_proxy_config_t, username unsafe.Pointer, password unsafe.Pointer) {
	_nw_proxy_config_set_username_and_password(proxy_config, username, password)
}

// Adds a supported Application-Layer Protocol Negotiation value.
//
// Added in macOS 12.0.
// Adds a supported Application-Layer Protocol Negotiation value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_add_tls_application_protocol(_:_:)
func nw_quic_add_tls_application_protocol(options Nw_protocol_options_t, application_protocol unsafe.Pointer) {
	_nw_quic_add_tls_application_protocol(options, application_protocol)
}

// Accesses the result of the QUIC handshake.
//
// Added in macOS 12.0.
// Accesses the result of the QUIC handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_metadata(_:)
func nw_quic_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_quic_copy_sec_protocol_metadata(metadata)
}

// Accesses the handshake security options QUIC will use.
//
// Added in macOS 12.0.
// Accesses the handshake security options QUIC will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_options(_:)
func nw_quic_copy_sec_protocol_options(options Nw_protocol_options_t) unsafe.Pointer {
	return _nw_quic_copy_sec_protocol_options(options)
}

// Initializes a default set of QUIC connection options.
//
// Added in macOS 12.0.
// Initializes a default set of QUIC connection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_create_options()
func nw_quic_create_options() Nw_protocol_options_t {
	return _nw_quic_create_options()
}

// Accesses the QUIC application error code received from the peer.
//
// Added in macOS 12.0.
// Accesses the QUIC application error code received from the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_application_error(_:)
func nw_quic_get_application_error(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_application_error(metadata)
}

// Accesses the QUIC application error reason received from the peer.
//
// Added in macOS 12.0.
// Accesses the QUIC application error reason received from the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_application_error_reason(_:)
func nw_quic_get_application_error_reason(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_quic_get_application_error_reason(metadata)
}

// Accesses the idle timeout for the QUIC connection, in milliseconds.
//
// Added in macOS 12.0.
// Accesses the idle timeout for the QUIC connection, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_idle_timeout(_:)
func nw_quic_get_idle_timeout(options Nw_protocol_options_t) uint32 {
	return _nw_quic_get_idle_timeout(options)
}

// Accesses a QUIC connection’s initial maximum data transport parameter.
//
// Added in macOS 12.0.
// Accesses a QUIC connection’s initial maximum data transport parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_data(_:)
func nw_quic_get_initial_max_data(options Nw_protocol_options_t) uint64 {
	return _nw_quic_get_initial_max_data(options)
}

// Accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// Added in macOS 12.0.
// Accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_remote(_:)
func nw_quic_get_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t) uint64 {
	return _nw_quic_get_initial_max_stream_data_bidirectional_remote(options)
}

// Accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// Added in macOS 12.0.
// Accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_unidirectional(_:)
func nw_quic_get_initial_max_stream_data_unidirectional(options Nw_protocol_options_t) uint64 {
	return _nw_quic_get_initial_max_stream_data_unidirectional(options)
}

// Accesses a QUIC connection’s initial maximum number of bidirectional streams.
//
// Added in macOS 12.0.
// Accesses a QUIC connection’s initial maximum number of bidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_bidirectional(_:)
func nw_quic_get_initial_max_streams_bidirectional(options Nw_protocol_options_t) uint64 {
	return _nw_quic_get_initial_max_streams_bidirectional(options)
}

// Accesses a QUIC connection’s initial maximum number of unidirectional streams.
//
// Added in macOS 12.0.
// Accesses a QUIC connection’s initial maximum number of unidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_unidirectional(_:)
func nw_quic_get_initial_max_streams_unidirectional(options Nw_protocol_options_t) uint64 {
	return _nw_quic_get_initial_max_streams_unidirectional(options)
}

// Accesses the keepalive interval for the QUIC connection, in seconds.
//
// Added in macOS 12.0.
// Accesses the keepalive interval for the QUIC connection, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_keepalive_interval(_:)
func nw_quic_get_keepalive_interval(metadata Nw_protocol_metadata_t) uint16 {
	return _nw_quic_get_keepalive_interval(metadata)
}

// Accesses a QUIC connection’s maximum DATAGRAM frame size.
//
// Added in macOS 13.0.
// Accesses a QUIC connection’s maximum DATAGRAM frame size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_max_datagram_frame_size(_:)
func nw_quic_get_max_datagram_frame_size(options Nw_protocol_options_t) uint16 {
	return _nw_quic_get_max_datagram_frame_size(options)
}

// Accesses the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// Added in macOS 12.0.
// Accesses the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_max_udp_payload_size(_:)
func nw_quic_get_max_udp_payload_size(options Nw_protocol_options_t) uint16 {
	return _nw_quic_get_max_udp_payload_size(options)
}

// Accesses the idle timeout value from the peer’s transport parameters, in milliseconds.
//
// Added in macOS 12.0.
// Accesses the idle timeout value from the peer’s transport parameters, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_idle_timeout(_:)
func nw_quic_get_remote_idle_timeout(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_remote_idle_timeout(metadata)
}

// Accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create.
//
// Added in macOS 12.0.
// Accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_bidirectional(_:)
func nw_quic_get_remote_max_streams_bidirectional(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_remote_max_streams_bidirectional(metadata)
}

// Accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create.
//
// Added in macOS 12.0.
// Accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_unidirectional(_:)
func nw_quic_get_remote_max_streams_unidirectional(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_remote_max_streams_unidirectional(metadata)
}

// Accesses the QUIC application error code received from the peer for the stream.
//
// Added in macOS 12.0.
// Accesses the QUIC application error code received from the peer for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_application_error(_:)
func nw_quic_get_stream_application_error(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_stream_application_error(metadata)
}

// Accesses the QUIC stream identifier.
//
// Added in macOS 12.0.
// Accesses the QUIC stream identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_id(_:)
func nw_quic_get_stream_id(metadata Nw_protocol_metadata_t) uint64 {
	return _nw_quic_get_stream_id(metadata)
}

// Checks if a QUIC stream is a datagram flow, instead of a byte stream.
//
// Added in macOS 13.0.
// Checks if a QUIC stream is a datagram flow, instead of a byte stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_datagram(_:)
func nw_quic_get_stream_is_datagram(options Nw_protocol_options_t) bool {
	return _nw_quic_get_stream_is_datagram(options)
}

// Checks if a QUIC stream is unidirectional, instead of bidirectional.
//
// Added in macOS 12.0.
// Checks if a QUIC stream is unidirectional, instead of bidirectional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_unidirectional(_:)
func nw_quic_get_stream_is_unidirectional(options Nw_protocol_options_t) bool {
	return _nw_quic_get_stream_is_unidirectional(options)
}

// Accesses the stream type of the QUIC stream.
//
// Added in macOS 12.0.
// Accesses the stream type of the QUIC stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_type(_:)
func nw_quic_get_stream_type(stream_metadata Nw_protocol_metadata_t) uint8 {
	return _nw_quic_get_stream_type(stream_metadata)
}

// Accesses the maximum usable size of a datagram frame on a QUIC datagram flow.
//
// Added in macOS 13.0.
// Accesses the maximum usable size of a datagram frame on a QUIC datagram flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_usable_datagram_frame_size(_:)
func nw_quic_get_stream_usable_datagram_frame_size(metadata Nw_protocol_metadata_t) uint16 {
	return _nw_quic_get_stream_usable_datagram_frame_size(metadata)
}

// Sets the QUIC application error code to send for the connection.
//
// Added in macOS 12.0.
// Sets the QUIC application error code to send for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_application_error(_:_:_:)
func nw_quic_set_application_error(metadata Nw_protocol_metadata_t, application_error uint64, reason unsafe.Pointer) {
	_nw_quic_set_application_error(metadata, application_error, reason)
}

// Sets the idle timeout for the QUIC connection, in milliseconds.
//
// Added in macOS 12.0.
// Sets the idle timeout for the QUIC connection, in milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_idle_timeout(_:_:)
func nw_quic_set_idle_timeout(options Nw_protocol_options_t, idle_timeout uint32) {
	_nw_quic_set_idle_timeout(options, idle_timeout)
}

// Sets a QUIC connection’s initial maximum data transport parameter.
//
// Added in macOS 12.0.
// Sets a QUIC connection’s initial maximum data transport parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_data(_:_:)
func nw_quic_set_initial_max_data(options Nw_protocol_options_t, initial_max_data uint64) {
	_nw_quic_set_initial_max_data(options, initial_max_data)
}

// Sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// Added in macOS 12.0.
// Sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_remote(_:_:)
func nw_quic_set_initial_max_stream_data_bidirectional_remote(options Nw_protocol_options_t, initial_max_stream_data_bidirectional_remote uint64) {
	_nw_quic_set_initial_max_stream_data_bidirectional_remote(options, initial_max_stream_data_bidirectional_remote)
}

// Sets a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// Added in macOS 12.0.
// Sets a QUIC connection’s initial maximum stream data limit for unidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_unidirectional(_:_:)
func nw_quic_set_initial_max_stream_data_unidirectional(options Nw_protocol_options_t, initial_max_stream_data_unidirectional uint64) {
	_nw_quic_set_initial_max_stream_data_unidirectional(options, initial_max_stream_data_unidirectional)
}

// Sets a QUIC connection’s initial maximum number of bidirectional streams.
//
// Added in macOS 12.0.
// Sets a QUIC connection’s initial maximum number of bidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_bidirectional(_:_:)
func nw_quic_set_initial_max_streams_bidirectional(options Nw_protocol_options_t, initial_max_streams_bidirectional uint64) {
	_nw_quic_set_initial_max_streams_bidirectional(options, initial_max_streams_bidirectional)
}

// Sets a QUIC connection’s initial maximum number of unidirectional streams.
//
// Added in macOS 12.0.
// Sets a QUIC connection’s initial maximum number of unidirectional streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_unidirectional(_:_:)
func nw_quic_set_initial_max_streams_unidirectional(options Nw_protocol_options_t, initial_max_streams_unidirectional uint64) {
	_nw_quic_set_initial_max_streams_unidirectional(options, initial_max_streams_unidirectional)
}

// Sets the keepalive interval for the QUIC connection, in seconds.
//
// Added in macOS 12.0.
// Sets the keepalive interval for the QUIC connection, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_keepalive_interval(_:_:)
func nw_quic_set_keepalive_interval(metadata Nw_protocol_metadata_t, keepalive_interval uint16) {
	_nw_quic_set_keepalive_interval(metadata, keepalive_interval)
}

// Sets a QUIC connection’s maximum DATAGRAM frame size.
//
// Added in macOS 13.0.
// Sets a QUIC connection’s maximum DATAGRAM frame size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_max_datagram_frame_size(_:_:)
func nw_quic_set_max_datagram_frame_size(options Nw_protocol_options_t, max_datagram_frame_size uint16) {
	_nw_quic_set_max_datagram_frame_size(options, max_datagram_frame_size)
}

// Sets the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// Added in macOS 12.0.
// Sets the maximum length of a QUIC packet that can be received on a connection, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_max_udp_payload_size(_:_:)
func nw_quic_set_max_udp_payload_size(options Nw_protocol_options_t, max_udp_payload_size uint16) {
	_nw_quic_set_max_udp_payload_size(options, max_udp_payload_size)
}

// Sets the QUIC application error code to send for the stream.
//
// Added in macOS 12.0.
// Sets the QUIC application error code to send for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_application_error(_:_:)
func nw_quic_set_stream_application_error(metadata Nw_protocol_metadata_t, application_error uint64) {
	_nw_quic_set_stream_application_error(metadata, application_error)
}

// Configures a QUIC stream as a datagram flow, instead of a byte stream.
//
// Added in macOS 13.0.
// Configures a QUIC stream as a datagram flow, instead of a byte stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_datagram(_:_:)
func nw_quic_set_stream_is_datagram(options Nw_protocol_options_t, is_datagram bool) {
	_nw_quic_set_stream_is_datagram(options, is_datagram)
}

// Configures a QUIC stream as unidirectional, instead of bidirectional.
//
// Added in macOS 12.0.
// Configures a QUIC stream as unidirectional, instead of bidirectional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_unidirectional(_:_:)
func nw_quic_set_stream_is_unidirectional(options Nw_protocol_options_t, is_unidirectional bool) {
	_nw_quic_set_stream_is_unidirectional(options, is_unidirectional)
}

// Adds an HTTP header name and value pair to send as part of requests to the relay.
//
// Added in macOS 14.0.
// Adds an HTTP header name and value pair to send as part of requests to the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_relay_hop_add_additional_http_header_field(_:_:_:)
func nw_relay_hop_add_additional_http_header_field(relay_hop Nw_relay_hop_t, field_name unsafe.Pointer, field_value unsafe.Pointer) {
	_nw_relay_hop_add_additional_http_header_field(relay_hop, field_name, field_value)
}

// Creates a configuration for a secure relay accessible using HTTP/3, with an optional HTTP/2 fallback.
//
// Added in macOS 14.0.
// Creates a configuration for a secure relay accessible using HTTP/3, with an optional HTTP/2 fallback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_relay_hop_create(_:_:_:)
func nw_relay_hop_create(http3_relay_endpoint Nw_endpoint_t, http2_relay_endpoint Nw_endpoint_t, relay_tls_options Nw_protocol_options_t) Nw_relay_hop_t {
	return _nw_relay_hop_create(http3_relay_endpoint, http2_relay_endpoint, relay_tls_options)
}

// Releases a reference count on a Network.framework object.
//
// Added in macOS 10.14.
// Releases a reference count on a Network.framework object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_release
func nw_release(obj unsafe.Pointer) {
	_nw_release(obj)
}

// Accesses the resolved endpoint that the connection used for its first connection attempt.
//
// Added in macOS 11.0.
// Accesses the resolved endpoint that the connection used for its first connection attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_preferred_endpoint(_:)
func nw_resolution_report_copy_preferred_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	return _nw_resolution_report_copy_preferred_endpoint(resolution_report)
}

// Accesses the resolved endpoint that led to the established connection.
//
// Added in macOS 11.0.
// Accesses the resolved endpoint that led to the established connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_successful_endpoint(_:)
func nw_resolution_report_copy_successful_endpoint(resolution_report Nw_resolution_report_t) Nw_endpoint_t {
	return _nw_resolution_report_copy_successful_endpoint(resolution_report)
}

// Accesses the number of endpoints resolved in this step.
//
// Added in macOS 11.0.
// Accesses the number of endpoints resolved in this step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_endpoint_count(_:)
func nw_resolution_report_get_endpoint_count(resolution_report Nw_resolution_report_t) uint32 {
	return _nw_resolution_report_get_endpoint_count(resolution_report)
}

// Accesses the duration of this resolution step, from when the query was issued to when the response was complete.
//
// Added in macOS 11.0.
// Accesses the duration of this resolution step, from when the query was issued to when the response was complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_milliseconds(_:)
func nw_resolution_report_get_milliseconds(resolution_report Nw_resolution_report_t) uint64 {
	return _nw_resolution_report_get_milliseconds(resolution_report)
}

// Accesses the transport protocol your connection used for DNS resolution.
//
// Added in macOS 11.0.
// Accesses the transport protocol your connection used for DNS resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_protocol(_:)
func nw_resolution_report_get_protocol(resolution_report Nw_resolution_report_t) unsafe.Pointer {
	return _nw_resolution_report_get_protocol(resolution_report)
}

// Accesses the source of the DNS response.
//
// Added in macOS 11.0.
// Accesses the source of the DNS response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_source(_:)
func nw_resolution_report_get_source(resolution_report Nw_resolution_report_t) unsafe.Pointer {
	return _nw_resolution_report_get_source(resolution_report)
}

// Provides a well-known DNS server address to use instead of looking up the address dynamically.
//
// Added in macOS 11.0.
// Provides a well-known DNS server address to use instead of looking up the address dynamically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolver_config_add_server_address(_:_:)
func nw_resolver_config_add_server_address(config Nw_resolver_config_t, server_address Nw_endpoint_t) {
	_nw_resolver_config_add_server_address(config, server_address)
}

// Initializes a DNS-over-HTTPS resolver configuration.
//
// Added in macOS 11.0.
// Initializes a DNS-over-HTTPS resolver configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolver_config_create_https(_:)
func nw_resolver_config_create_https(url_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	return _nw_resolver_config_create_https(url_endpoint)
}

// Initializes a DNS-over-TLS resolver configuration.
//
// Added in macOS 11.0.
// Initializes a DNS-over-TLS resolver configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolver_config_create_tls(_:)
func nw_resolver_config_create_tls(server_endpoint Nw_endpoint_t) Nw_resolver_config_t {
	return _nw_resolver_config_create_tls(server_endpoint)
}

// Adds a reference count to a Network.framework object.
//
// Added in macOS 10.14.
// Adds a reference count to a Network.framework object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_retain
func nw_retain(obj unsafe.Pointer) unsafe.Pointer {
	return _nw_retain(obj)
}

// Initializes a default set of TCP connection options.
//
// Added in macOS 10.14.
// Initializes a default set of TCP connection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_create_options()
func nw_tcp_create_options() Nw_protocol_options_t {
	return _nw_tcp_create_options()
}

// Accesses the number of available bytes in the TCP receive buffer.
//
// Added in macOS 10.14.
// Accesses the number of available bytes in the TCP receive buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_get_available_receive_buffer(_:)
func nw_tcp_get_available_receive_buffer(metadata Nw_protocol_metadata_t) uint32 {
	return _nw_tcp_get_available_receive_buffer(metadata)
}

// Accesses the number of available bytes in the TCP send buffer.
//
// Added in macOS 10.14.
// Accesses the number of available bytes in the TCP send buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_get_available_send_buffer(_:)
func nw_tcp_get_available_send_buffer(metadata Nw_protocol_metadata_t) uint32 {
	return _nw_tcp_get_available_send_buffer(metadata)
}

// Sets the number of seconds that TCP waits before timing out its handshake.
//
// Added in macOS 10.14.
// Sets the number of seconds that TCP waits before timing out its handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_connection_timeout(_:_:)
func nw_tcp_options_set_connection_timeout(options Nw_protocol_options_t, connection_timeout uint32) {
	_nw_tcp_options_set_connection_timeout(options, connection_timeout)
}

// Disables TCP acknowledgment stretching.
//
// Added in macOS 10.14.
// Disables TCP acknowledgment stretching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ack_stretching(_:_:)
func nw_tcp_options_set_disable_ack_stretching(options Nw_protocol_options_t, disable_ack_stretching bool) {
	_nw_tcp_options_set_disable_ack_stretching(options, disable_ack_stretching)
}

// Disables negotiation of Explicit Congestion Notification markings.
//
// Added in macOS 10.14.
// Disables negotiation of Explicit Congestion Notification markings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ecn(_:_:)
func nw_tcp_options_set_disable_ecn(options Nw_protocol_options_t, disable_ecn bool) {
	_nw_tcp_options_set_disable_ecn(options, disable_ecn)
}

// Enables TCP Fast Open on a connection.
//
// Added in macOS 10.14.
// Enables TCP Fast Open on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_fast_open(_:_:)
func nw_tcp_options_set_enable_fast_open(options Nw_protocol_options_t, enable_fast_open bool) {
	_nw_tcp_options_set_enable_fast_open(options, enable_fast_open)
}

// Enables TCP keepalives.
//
// Added in macOS 10.14.
// Enables TCP keepalives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_keepalive(_:_:)
func nw_tcp_options_set_enable_keepalive(options Nw_protocol_options_t, enable_keepalive bool) {
	_nw_tcp_options_set_enable_keepalive(options, enable_keepalive)
}

// Sets the number of keepalive probes that TCP sends before terminating the connection.
//
// Added in macOS 10.14.
// Sets the number of keepalive probes that TCP sends before terminating the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_count(_:_:)
func nw_tcp_options_set_keepalive_count(options Nw_protocol_options_t, keepalive_count uint32) {
	_nw_tcp_options_set_keepalive_count(options, keepalive_count)
}

// Sets the number of seconds of idleness that TCP waits before sending keepalive probes.
//
// Added in macOS 10.14.
// Sets the number of seconds of idleness that TCP waits before sending keepalive probes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_idle_time(_:_:)
func nw_tcp_options_set_keepalive_idle_time(options Nw_protocol_options_t, keepalive_idle_time uint32) {
	_nw_tcp_options_set_keepalive_idle_time(options, keepalive_idle_time)
}

// Sets the number of seconds that TCP waits between sending keepalive probes.
//
// Added in macOS 10.14.
// Sets the number of seconds that TCP waits between sending keepalive probes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_interval(_:_:)
func nw_tcp_options_set_keepalive_interval(options Nw_protocol_options_t, keepalive_interval uint32) {
	_nw_tcp_options_set_keepalive_interval(options, keepalive_interval)
}

// Sets TCP’s maximum segment size in bytes.
//
// Added in macOS 10.14.
// Sets TCP’s maximum segment size in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_maximum_segment_size(_:_:)
func nw_tcp_options_set_maximum_segment_size(options Nw_protocol_options_t, maximum_segment_size uint32) {
	_nw_tcp_options_set_maximum_segment_size(options, maximum_segment_size)
}

// Disables Nagle’s algorithm for TCP.
//
// Added in macOS 10.14.
// Disables Nagle’s algorithm for TCP.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_delay(_:_:)
func nw_tcp_options_set_no_delay(options Nw_protocol_options_t, no_delay bool) {
	_nw_tcp_options_set_no_delay(options, no_delay)
}

// Sets TCP into no-options mode.
//
// Added in macOS 10.14.
// Sets TCP into no-options mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_options(_:_:)
func nw_tcp_options_set_no_options(options Nw_protocol_options_t, no_options bool) {
	_nw_tcp_options_set_no_options(options, no_options)
}

// Sets TCP into no-push mode.
//
// Added in macOS 10.14.
// Sets TCP into no-push mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_push(_:_:)
func nw_tcp_options_set_no_push(options Nw_protocol_options_t, no_push bool) {
	_nw_tcp_options_set_no_push(options, no_push)
}

// Sets the TCP persist timeout in seconds, as defined by RFC 6429.
//
// Added in macOS 10.14.
// Sets the TCP persist timeout in seconds, as defined by RFC 6429.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_persist_timeout(_:_:)
func nw_tcp_options_set_persist_timeout(options Nw_protocol_options_t, persist_timeout uint32) {
	_nw_tcp_options_set_persist_timeout(options, persist_timeout)
}

// Sets the number of seconds that TCP waits between retransmission attempts.
//
// Added in macOS 10.14.
// Sets the number of seconds that TCP waits between retransmission attempts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_connection_drop_time(_:_:)
func nw_tcp_options_set_retransmit_connection_drop_time(options Nw_protocol_options_t, retransmit_connection_drop_time uint32) {
	_nw_tcp_options_set_retransmit_connection_drop_time(options, retransmit_connection_drop_time)
}

// Causes TCP to drop its connection after not receiving an ACK after a FIN.
//
// Added in macOS 10.14.
// Causes TCP to drop its connection after not receiving an ACK after a FIN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_fin_drop(_:_:)
func nw_tcp_options_set_retransmit_fin_drop(options Nw_protocol_options_t, retransmit_fin_drop bool) {
	_nw_tcp_options_set_retransmit_fin_drop(options, retransmit_fin_drop)
}

// Accesses the result of the TLS handshake.
//
// Added in macOS 10.14.
// Accesses the result of the TLS handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_metadata(_:)
func nw_tls_copy_sec_protocol_metadata(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_tls_copy_sec_protocol_metadata(metadata)
}

// Accesses the handshake security options TLS will use.
//
// Added in macOS 10.14.
// Accesses the handshake security options TLS will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_options(_:)
func nw_tls_copy_sec_protocol_options(options Nw_protocol_options_t) unsafe.Pointer {
	return _nw_tls_copy_sec_protocol_options(options)
}

// Initializes a default set of TLS connection options.
//
// Added in macOS 10.14.
// Initializes a default set of TLS connection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_create_options()
func nw_tls_create_options() Nw_protocol_options_t {
	return _nw_tls_create_options()
}

// Accesses the raw bytes contained within a TXT record.
//
// Added in macOS 10.15.
// Accesses the raw bytes contained within a TXT record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes(_:_:)
func nw_txt_record_access_bytes(txt_record Nw_txt_record_t, access_bytes unsafe.Pointer) bool {
	return _nw_txt_record_access_bytes(txt_record, access_bytes)
}

// Accesses the value for a specific key in a TXT record dictionary.
//
// Added in macOS 10.15.
// Accesses the value for a specific key in a TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_access_key(_:_:_:)
func nw_txt_record_access_key(txt_record Nw_txt_record_t, key unsafe.Pointer, access_value unsafe.Pointer) bool {
	return _nw_txt_record_access_key(txt_record, key, access_value)
}

// Iterates through all keys in a TXT record dictionary.
//
// Added in macOS 10.15.
// Iterates through all keys in a TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_apply(_:_:)
func nw_txt_record_apply(txt_record Nw_txt_record_t, applier unsafe.Pointer) bool {
	return _nw_txt_record_apply(txt_record, applier)
}

// Performs a deep copy of a TXT record.
//
// Added in macOS 10.15.
// Performs a deep copy of a TXT record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_copy(_:)
func nw_txt_record_copy(txt_record Nw_txt_record_t) Nw_txt_record_t {
	return _nw_txt_record_copy(txt_record)
}

// Initializes a TXT record as a dictionary of strings.
//
// Added in macOS 10.15.
// Initializes a TXT record as a dictionary of strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_create_dictionary()
func nw_txt_record_create_dictionary() Nw_txt_record_t {
	return _nw_txt_record_create_dictionary()
}

// Initializes a TXT record with raw bytes.
//
// Added in macOS 10.15.
// Initializes a TXT record with raw bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_create_with_bytes(_:_:)
func nw_txt_record_create_with_bytes(txt_bytes unsafe.Pointer, txt_len uintptr) Nw_txt_record_t {
	return _nw_txt_record_create_with_bytes(txt_bytes, txt_len)
}

// Checks the status of value associated with a key in a TXT record dictionary.
//
// Added in macOS 10.15.
// Checks the status of value associated with a key in a TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_find_key(_:_:)
func nw_txt_record_find_key(txt_record Nw_txt_record_t, key unsafe.Pointer) unsafe.Pointer {
	return _nw_txt_record_find_key(txt_record, key)
}

// Accesses the number of keys stored in the TXT record dictionary.
//
// Added in macOS 10.15.
// Accesses the number of keys stored in the TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_get_key_count(_:)
func nw_txt_record_get_key_count(txt_record Nw_txt_record_t) uintptr {
	return _nw_txt_record_get_key_count(txt_record)
}

// Checks whether a TXT record conforms to a dictionary format.
//
// Added in macOS 10.15.
// Checks whether a TXT record conforms to a dictionary format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_is_dictionary(_:)
func nw_txt_record_is_dictionary(txt_record Nw_txt_record_t) bool {
	return _nw_txt_record_is_dictionary(txt_record)
}

// Checks whether two TXT records are equivalent.
//
// Added in macOS 10.15.
// Checks whether two TXT records are equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_is_equal(_:_:)
func nw_txt_record_is_equal(left Nw_txt_record_t, right Nw_txt_record_t) bool {
	return _nw_txt_record_is_equal(left, right)
}

// Removes a data value in a TXT record dictionary.
//
// Added in macOS 10.15.
// Removes a data value in a TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_remove_key(_:_:)
func nw_txt_record_remove_key(txt_record Nw_txt_record_t, key unsafe.Pointer) bool {
	return _nw_txt_record_remove_key(txt_record, key)
}

// Sets a data value in a TXT record dictionary.
//
// Added in macOS 10.15.
// Sets a data value in a TXT record dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_set_key(_:_:_:_:)
func nw_txt_record_set_key(txt_record Nw_txt_record_t, key unsafe.Pointer, value unsafe.Pointer, value_len uintptr) bool {
	return _nw_txt_record_set_key(txt_record, key, value, value_len)
}

// Initializes a default UDP message.
//
// Added in macOS 10.14.
// Initializes a default UDP message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_create_metadata()
func nw_udp_create_metadata() Nw_protocol_metadata_t {
	return _nw_udp_create_metadata()
}

// Initializes a default set of UDP connection options.
//
// Added in macOS 10.14.
// Initializes a default set of UDP connection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_create_options()
func nw_udp_create_options() Nw_protocol_options_t {
	return _nw_udp_create_options()
}

// Configures the connection to not send UDP checksums.
//
// Added in macOS 10.14.
// Configures the connection to not send UDP checksums.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_options_set_prefer_no_checksum(_:_:)
func nw_udp_options_set_prefer_no_checksum(options Nw_protocol_options_t, prefer_no_checksum bool) {
	_nw_udp_options_set_prefer_no_checksum(options, prefer_no_checksum)
}

// Initializes a WebSocket message with a specific type code.
//
// Added in macOS 10.15.
// Initializes a WebSocket message with a specific type code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_create_metadata(_:)
func nw_ws_create_metadata(opcode unsafe.Pointer) Nw_protocol_metadata_t {
	return _nw_ws_create_metadata(opcode)
}

// Initializes a default set of WebSocket connection options.
//
// Added in macOS 10.15.
// Initializes a default set of WebSocket connection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_create_options(_:)
func nw_ws_create_options(version unsafe.Pointer) Nw_protocol_options_t {
	return _nw_ws_create_options(version)
}

// Accesses the WebSocket server’s response sent during the handshake.
//
// Added in macOS 10.15.
// Accesses the WebSocket server’s response sent during the handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_copy_server_response(_:)
func nw_ws_metadata_copy_server_response(metadata Nw_protocol_metadata_t) Nw_ws_response_t {
	return _nw_ws_metadata_copy_server_response(metadata)
}

// Accesses the close code on a WebSocket message.
//
// Added in macOS 10.15.
// Accesses the close code on a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_close_code(_:)
func nw_ws_metadata_get_close_code(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_ws_metadata_get_close_code(metadata)
}

// Checks the type code on a WebSocket message.
//
// Added in macOS 10.15.
// Checks the type code on a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_opcode(_:)
func nw_ws_metadata_get_opcode(metadata Nw_protocol_metadata_t) unsafe.Pointer {
	return _nw_ws_metadata_get_opcode(metadata)
}

// Sets a close code on a WebSocket message.
//
// Added in macOS 10.15.
// Sets a close code on a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_close_code(_:_:)
func nw_ws_metadata_set_close_code(metadata Nw_protocol_metadata_t, close_code unsafe.Pointer) {
	_nw_ws_metadata_set_close_code(metadata, close_code)
}

// Sets a handler on a Ping message to be invoked when the corresponding Pong message is received.
//
// Added in macOS 10.15.
// Sets a handler on a Ping message to be invoked when the corresponding Pong message is received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_pong_handler(_:_:_:)
func nw_ws_metadata_set_pong_handler(metadata Nw_protocol_metadata_t, client_queue unsafe.Pointer, pong_handler unsafe.Pointer) {
	_nw_ws_metadata_set_pong_handler(metadata, client_queue, pong_handler)
}

// Adds additional HTTP header fields to be sent by the client during the WebSocket handshake.
//
// Added in macOS 10.15.
// Adds additional HTTP header fields to be sent by the client during the WebSocket handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_add_additional_header(_:_:_:)
func nw_ws_options_add_additional_header(options Nw_protocol_options_t, name unsafe.Pointer, value unsafe.Pointer) {
	_nw_ws_options_add_additional_header(options, name, value)
}

// Adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment.
//
// Added in macOS 10.15.
// Adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_add_subprotocol(_:_:)
func nw_ws_options_add_subprotocol(options Nw_protocol_options_t, subprotocol unsafe.Pointer) {
	_nw_ws_options_add_subprotocol(options, subprotocol)
}

// Configures the connection to automatically reply to Ping messages instead of delivering them to you.
//
// Added in macOS 10.15.
// Configures the connection to automatically reply to Ping messages instead of delivering them to you.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_auto_reply_ping(_:_:)
func nw_ws_options_set_auto_reply_ping(options Nw_protocol_options_t, auto_reply_ping bool) {
	_nw_ws_options_set_auto_reply_ping(options, auto_reply_ping)
}

// Sets a handler to react to as a server to inbound WebSocket client handshakes.
//
// Added in macOS 10.15.
// Sets a handler to react to as a server to inbound WebSocket client handshakes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_client_request_handler(_:_:_:)
func nw_ws_options_set_client_request_handler(options Nw_protocol_options_t, client_queue unsafe.Pointer, handler unsafe.Pointer) {
	_nw_ws_options_set_client_request_handler(options, client_queue, handler)
}

// Sets the maximum allowed message size, in bytes, to be received by the WebSocket connection.
//
// Added in macOS 10.15.
// Sets the maximum allowed message size, in bytes, to be received by the WebSocket connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_maximum_message_size(_:_:)
func nw_ws_options_set_maximum_message_size(options Nw_protocol_options_t, maximum_message_size uintptr) {
	_nw_ws_options_set_maximum_message_size(options, maximum_message_size)
}

// Specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established.
//
// Added in macOS 10.15.
// Specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_skip_handshake(_:_:)
func nw_ws_options_set_skip_handshake(options Nw_protocol_options_t, skip_handshake bool) {
	_nw_ws_options_set_skip_handshake(options, skip_handshake)
}

// Enumerates additional HTTP headers in a WebSocket message.
//
// Added in macOS 10.15.
// Enumerates additional HTTP headers in a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_additional_headers(_:_:)
func nw_ws_request_enumerate_additional_headers(request Nw_ws_request_t, enumerator unsafe.Pointer) bool {
	return _nw_ws_request_enumerate_additional_headers(request, enumerator)
}

// Enumerates the supported subprotocols in a WebSocket message.
//
// Added in macOS 10.15.
// Enumerates the supported subprotocols in a WebSocket message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_subprotocols(_:_:)
func nw_ws_request_enumerate_subprotocols(request Nw_ws_request_t, enumerator unsafe.Pointer) bool {
	return _nw_ws_request_enumerate_subprotocols(request, enumerator)
}

// Adds an additional HTTP header to a WebSocket server response.
//
// Added in macOS 10.15.
// Adds an additional HTTP header to a WebSocket server response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_add_additional_header(_:_:_:)
func nw_ws_response_add_additional_header(response Nw_ws_response_t, name unsafe.Pointer, value unsafe.Pointer) {
	_nw_ws_response_add_additional_header(response, name, value)
}

// Initializes a WebSocket server response with a status and selected subprotocol.
//
// Added in macOS 10.15.
// Initializes a WebSocket server response with a status and selected subprotocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_create(_:_:)
func nw_ws_response_create(status unsafe.Pointer, selected_subprotocol unsafe.Pointer) Nw_ws_response_t {
	return _nw_ws_response_create(status, selected_subprotocol)
}

// Enumerates the additional HTTP headers in a WebSocket server response.
//
// Added in macOS 10.15.
// Enumerates the additional HTTP headers in a WebSocket server response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_enumerate_additional_headers(_:_:)
func nw_ws_response_enumerate_additional_headers(response Nw_ws_response_t, enumerator unsafe.Pointer) bool {
	return _nw_ws_response_enumerate_additional_headers(response, enumerator)
}

// Accesses the selected subprotocol in a WebSocket server response.
//
// Added in macOS 10.15.
// Accesses the selected subprotocol in a WebSocket server response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_get_selected_subprotocol(_:)
func nw_ws_response_get_selected_subprotocol(response Nw_ws_response_t) unsafe.Pointer {
	return _nw_ws_response_get_selected_subprotocol(response)
}

// Accesses the status of a WebSocket server response.
//
// Added in macOS 10.15.
// Accesses the status of a WebSocket server response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_get_status(_:)
func nw_ws_response_get_status(response Nw_ws_response_t) unsafe.Pointer {
	return _nw_ws_response_get_status(response)
}




