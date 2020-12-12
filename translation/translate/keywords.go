package translate

import "sort"

var special_idents = []string{
	// "actions",
	// "append_to_list",
	// "assert",
	// "break",
	// "check_any",
	// "check_value",
	// "create_list",
	// "debug_text",
	// "do_any",
	"do_all",
	"do_else",
	"do_elseif",
	"do_for_each",
	"do_if",
	"do_while",
	// "handler",
	// "input_param",
	// "interrupt",
	// "interrupts",
	// "label",
	// "order",
	// "param",
	// "remove_value",
	// "resume",
	// "return",
	// "run_script",
	// "set_value",
	// "add_faction_relation",
	// "add_relation_boost",
	// "attention",
	// "conditions",
	// "create_order",
	// "init",
	// "interrupt_after_time",
	// "match_content",
	// "match_distance",
	// "move_strafe",
	// "on_abort",
	// "position",
	// "replace",
	// "rotation",
	// "save_retval",
	// "set_command",
	// "set_command_action",
	// "set_order_syncpoint_reached",
	// "show_notification",
	// "shuffle_list",
	// "signal_objects",
	// "skill",
	// "substitute_text",
	// "unknown",
	// "wait",
	// "write_to_logbook",
}

func init() {
	sort.Strings(special_idents)
}
