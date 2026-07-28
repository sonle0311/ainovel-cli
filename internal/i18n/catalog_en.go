package i18n

// catalogEN 英文文案。key 集合必须与 catalogZH 完全一致（测试强制）。
var catalogEN = map[string]string{
	// command_help.go — /help panel
	"help.title":          "Command Help",
	"help.shortcuts":      "Shortcuts",
	"help.hint_search":    "Type / to search commands",
	"help.hint_select":    "↑↓ select a command",
	"help.hint_complete":  "Tab/Enter accept completion",
	"help.hint_close":     "Esc close current panel",
	"help.hint_copy_mode": "Ctrl+R toggle copy mode (disables mouse reporting so you can drag-select; press again to restore)",
	"help.modal_footer":   "  ↑↓ scroll · Esc close",

	// commands.go — /lang UI language switch
	"lang.description": "Switch UI language (zh/en/vi)",
	"lang.current":     "Current UI language: %s (supported: zh/en/vi, usage: /lang <lang>)",
	"lang.switched":    "UI language switched: %s",
	"lang.unsupported": "Unsupported language %q (supported: zh/en/vi)",
	"lang.save_failed": "Language switched but saving config failed: %v",
}
