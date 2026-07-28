package i18n

// catalogVI 越南语文案。key 集合必须与 catalogZH 完全一致（测试强制）。
var catalogVI = map[string]string{
	// command_help.go — bảng trợ giúp /help
	"help.title":          "Trợ giúp lệnh",
	"help.shortcuts":      "Phím tắt",
	"help.hint_search":    "Gõ / để tìm lệnh",
	"help.hint_select":    "↑↓ chọn lệnh",
	"help.hint_complete":  "Tab/Enter chấp nhận gợi ý",
	"help.hint_close":     "Esc đóng bảng lệnh hiện tại",
	"help.hint_copy_mode": "Ctrl+R bật/tắt chế độ chọn-sao chép (tắt mouse reporting để kéo chọn văn bản; nhấn lần nữa để khôi phục)",
	"help.modal_footer":   "  ↑↓ cuộn · Esc đóng",

	// commands.go — /lang đổi ngôn ngữ giao diện
	"lang.description": "Đổi ngôn ngữ giao diện (zh/en/vi)",
	"lang.current":     "Ngôn ngữ giao diện hiện tại: %s (hỗ trợ zh/en/vi, cách dùng: /lang <ngôn ngữ>)",
	"lang.switched":    "Đã đổi ngôn ngữ giao diện: %s",
	"lang.unsupported": "Ngôn ngữ %q không được hỗ trợ (hỗ trợ zh/en/vi)",
	"lang.save_failed": "Đã đổi ngôn ngữ nhưng lưu config thất bại: %v",
}
