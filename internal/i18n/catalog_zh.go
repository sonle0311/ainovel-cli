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

	// commands.go — 命令描述
	"cmd.help.desc":      "查看命令列表",
	"cmd.model.desc":     "切换角色的模型与推理强度",
	"cmd.lang.desc":      "切换界面语言（zh/en/vi）",
	"cmd.config.desc":    "新增或编辑 Provider、模型与上下文窗口",
	"cmd.diag.desc":      "诊断小说创作健康度",
	"cmd.review.desc":    "切换逐章验收模式",
	"cmd.next.desc":      "验收后放行一个新章节",
	"cmd.import.desc":    "语义导入外部小说（无参数则恢复未完成导入；--guide 用自然语言调整切分）",
	"cmd.reopen.desc":    "重开已完结的书继续创作（方向先经裁定注入，再自动续跑）",
	"cmd.cocreate.desc":  "暂停创作，共创规划后续阶段走向",
	"cmd.simulate.desc":  "读取 ./simulate 生成或增量更新仿写画像",
	"cmd.importsim.desc": "导入已有仿写画像并按语料指纹合并",
	"cmd.export.desc":    "导出已完成章节为 TXT/EPUB",

	// command_palette.go — 命令弹窗
	"palette.title": "命令",
	"palette.hint":  "↑↓ 选择 · Tab/Enter 接受 · Esc 关闭",
	"palette.more":  " · 还有 %d 个命令",

	// commands.go — /lang 界面语言切换
	"lang.current":     "当前界面语言：%s（支持 zh/en/vi，用法：/lang <语言>）",
	"lang.switched":    "界面语言已切换：%s",
	"lang.unsupported": "不支持的语言 %q（支持 zh/en/vi）",
	"lang.save_failed": "语言已切换但写入配置失败：%v",

	// panels.go — 新建态首屏
	"welcome.feat_multi_model":      "多模型协作",
	"welcome.feat_multi_model_desc": "Architect 规划 / Writer 创作 / Editor 审阅",
	"welcome.feat_checkpoint":       "断点恢复",
	"welcome.feat_checkpoint_desc":  "崩溃或中断后从上次进度自动续写",
	"welcome.feat_steer":            "实时干预",
	"welcome.feat_steer_desc":       "创作过程中随时调整剧情走向",
	"welcome.feat_longform":         "分层长篇",
	"welcome.feat_longform_desc":    "支持卷-弧-章分层结构的长篇创作",
	"welcome.current_mode":          "当前模式：%s · %s",
	"welcome.prompt":                "在下方输入你的小说需求开始创作",
	"welcome.example_1":             "写一部 12 章都市悬疑小说，主角是一名女法医",
	"welcome.example_2":             "创作一部仙侠长篇，主角从凡人修炼至飞升",
	"welcome.example_3":             "写一个科幻短篇，讲述 AI 觉醒后的伦理困境",
	"welcome.import_hint":           "已有小说存稿想接着写？输入 /import <文件路径> 导入后续写",
	"welcome.tab_hint":              "Tab 切换模式 · 快速开始下 Enter 直接创作 · 共创规划下 Enter 进入对话",

	// cocreate.go — 启动模式
	"startup.mode_quick":             "快速开始",
	"startup.mode_cocreate":          "共创规划",
	"startup.mode_quick_subtitle":    "一句话直接开始写",
	"startup.mode_cocreate_subtitle": "先与 AI 对话澄清，再开始创作",
	"startup.mode_bar":               "启动模式",
	"startup.placeholder_quick":      "输入一句小说需求，Enter 直接开始创作",
	"startup.placeholder_cocreate":   "先输入你的核心想法，Enter 开始与 AI 共创",

	// model.go — 新建态底部提示
	"footer.new_quick":    "Tab 切换启动模式 · 输入 / 搜索命令 · Enter 直接开始创作 · Esc 清空输入",
	"footer.new_cocreate": "Tab 切换启动模式 · 输入 / 搜索命令 · Enter 开始共创对话 · Esc 清空输入",
}
