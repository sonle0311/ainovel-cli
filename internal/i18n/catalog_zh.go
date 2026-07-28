package i18n

// catalogZH 中文文案（源语言）。key 命名：<模块>.<语义>，与 TUI 文件对应。
var catalogZH = map[string]string{
	// command_help.go — /help 帮助面板
	"help.title":          "命令帮助",
	"help.shortcuts":      "快捷键",
	"help.hint_search":    "输入 / 搜索命令",
	"help.hint_select":    "↑↓ 选择命令候选",
	"help.hint_complete":  "Tab/Enter 接受补全",
	"help.hint_close":     "Esc 关闭当前命令面板",
	"help.hint_copy_mode": "Ctrl+R 切换选中复制模式（关闭鼠标上报后可拖拽选中复制，再按一次恢复）",
	"help.modal_footer":   "  ↑↓ 滚动 · Esc 关闭",
}
