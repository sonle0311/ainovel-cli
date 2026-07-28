// Package i18n 提供 TUI 界面文案的多语言查找。
//
// 设计约束：
//   - 中文（zh）是源语言与最终回落，未配置 ui_language 时行为与历史版本完全一致。
//   - 语言在启动时 Set 一次（读配置之后、UI/引擎 goroutine 启动之前），
//     之后只读；Set 不做并发保护。
//   - 只覆盖界面文案（菜单、帮助、状态栏、提示）。创作提示词、写作规则等
//     LLM 侧内容不走此包——那些由 voice/rules 覆盖层管理。
package i18n

import "strings"

// Lang 界面语言代号。
type Lang string

const (
	LangZH Lang = "zh"
	LangEN Lang = "en"
	LangVI Lang = "vi"
)

// current 当前生效语言；启动时 Set 一次，之后只读。
var current = LangZH

// Normalize 把用户配置值归一为受支持语言。空值=默认中文。
// 未支持的值回落中文并返回 false，由调用方决定是否告警。
func Normalize(raw string) (Lang, bool) {
	s := strings.ToLower(strings.TrimSpace(raw))
	// 容忍 BCP 47 区域后缀（zh-CN / en-US / vi-VN）
	if i := strings.IndexAny(s, "-_"); i >= 0 {
		s = s[:i]
	}
	switch s {
	case "", "zh":
		return LangZH, s == "" || s == "zh"
	case "en":
		return LangEN, true
	case "vi":
		return LangVI, true
	default:
		return LangZH, false
	}
}

// Set 设置当前界面语言。未支持的值静默回落中文（配置校验层负责告警）。
func Set(raw string) {
	lang, _ := Normalize(raw)
	current = lang
}

// Current 返回当前生效语言。
func Current() Lang {
	return current
}

// T 按当前语言查找文案。查找顺序：当前语言 → 中文 → key 原样返回。
// key 原样返回让遗漏在界面上直接可见，而不是渲染成空白。
func T(key string) string {
	if current != LangZH {
		if catalog := catalogFor(current); catalog != nil {
			if val, ok := catalog[key]; ok {
				return val
			}
		}
	}
	if val, ok := catalogZH[key]; ok {
		return val
	}
	return key
}

func catalogFor(lang Lang) map[string]string {
	switch lang {
	case LangEN:
		return catalogEN
	case LangVI:
		return catalogVI
	default:
		return nil
	}
}
