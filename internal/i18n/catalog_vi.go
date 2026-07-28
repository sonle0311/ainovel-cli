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

	// panels.go — màn hình chào
	"welcome.feat_multi_model":      "Đa agent",
	"welcome.feat_multi_model_desc": "Architect lên kế hoạch / Writer viết / Editor duyệt",
	"welcome.feat_checkpoint":       "Checkpoint",
	"welcome.feat_checkpoint_desc":  "Tự động viết tiếp từ tiến độ cũ sau khi crash hoặc gián đoạn",
	"welcome.feat_steer":            "Can thiệp trực tiếp",
	"welcome.feat_steer_desc":       "Điều chỉnh hướng cốt truyện bất cứ lúc nào khi đang viết",
	"welcome.feat_longform":         "Truyện dài",
	"welcome.feat_longform_desc":    "Cấu trúc phân tầng quyển-arc-chương cho truyện dài kỳ",
	"welcome.current_mode":          "Chế độ hiện tại: %s · %s",
	"welcome.prompt":                "Nhập yêu cầu tiểu thuyết của bạn bên dưới để bắt đầu",
	"welcome.example_1":             "Viết truyện trinh thám đô thị 12 chương, nhân vật chính là nữ pháp y",
	"welcome.example_2":             "Sáng tác trường thiên tiên hiệp: từ phàm nhân tu luyện đến phi thăng",
	"welcome.example_3":             "Viết truyện ngắn khoa học viễn tưởng về đạo đức của AI thức tỉnh",
	"welcome.import_hint":           "Đã có bản thảo? Gõ /import <đường dẫn file> để nhập và viết tiếp",
	"welcome.tab_hint":              "Tab đổi chế độ · Bắt đầu nhanh: Enter để viết · Đồng sáng tạo: Enter để trò chuyện",

	// cocreate.go — chế độ khởi động
	"startup.mode_quick":             "Bắt đầu nhanh",
	"startup.mode_cocreate":          "Đồng sáng tạo",
	"startup.mode_quick_subtitle":    "một câu là bắt đầu viết ngay",
	"startup.mode_cocreate_subtitle": "trao đổi với AI cho rõ trước, rồi mới viết",
	"startup.mode_bar":               "Chế độ khởi động",
	"startup.placeholder_quick":      "Nhập một câu yêu cầu tiểu thuyết, Enter để bắt đầu",
	"startup.placeholder_cocreate":   "Nhập ý tưởng cốt lõi, Enter để đồng sáng tạo với AI",

	// model.go — gợi ý thanh dưới màn hình tạo mới
	"footer.new_quick":    "Tab đổi chế độ khởi động · / tìm lệnh · Enter bắt đầu viết · Esc xoá input",
	"footer.new_cocreate": "Tab đổi chế độ khởi động · / tìm lệnh · Enter bắt đầu đồng sáng tạo · Esc xoá input",
}
