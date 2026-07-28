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

	// panels.go — welcome screen
	"welcome.feat_multi_model":      "Multi-agent",
	"welcome.feat_multi_model_desc": "Architect plans / Writer writes / Editor reviews",
	"welcome.feat_checkpoint":       "Checkpoint",
	"welcome.feat_checkpoint_desc":  "Auto-resume from last progress after crash or interrupt",
	"welcome.feat_steer":            "Live steering",
	"welcome.feat_steer_desc":       "Adjust plot direction anytime during writing",
	"welcome.feat_longform":         "Long-form",
	"welcome.feat_longform_desc":    "Volume-arc-chapter layered structure for long novels",
	"welcome.current_mode":          "Current mode: %s · %s",
	"welcome.prompt":                "Type your novel request below to start",
	"welcome.example_1":             "Write a 12-chapter urban mystery starring a female forensic examiner",
	"welcome.example_2":             "Create a xianxia epic: a mortal cultivates all the way to ascension",
	"welcome.example_3":             "Write a sci-fi short story on the ethics of an awakened AI",
	"welcome.import_hint":           "Have an existing draft? Type /import <file path> to continue it",
	"welcome.tab_hint":              "Tab switch mode · Quick Start: Enter to create · Co-Create: Enter to chat",

	// cocreate.go — startup modes
	"startup.mode_quick":             "Quick Start",
	"startup.mode_cocreate":          "Co-Create",
	"startup.mode_quick_subtitle":    "start writing from one sentence",
	"startup.mode_cocreate_subtitle": "clarify with AI first, then write",
	"startup.mode_bar":               "Startup Mode",
	"startup.placeholder_quick":      "Type a one-line novel request, Enter to start",
	"startup.placeholder_cocreate":   "Type your core idea, Enter to co-create with AI",

	// model.go — footer hints on the new-book screen
	"footer.new_quick":    "Tab switch startup mode · / search commands · Enter start writing · Esc clear input",
	"footer.new_cocreate": "Tab switch startup mode · / search commands · Enter start co-creating · Esc clear input",
}
