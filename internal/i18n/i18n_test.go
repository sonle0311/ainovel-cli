package i18n

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		in   string
		want Lang
		ok   bool
	}{
		{"", LangZH, true}, // 空值=默认中文，保持现有行为
		{"zh", LangZH, true},
		{"zh-CN", LangZH, true},
		{"ZH", LangZH, true},
		{"en", LangEN, true},
		{"en-US", LangEN, true},
		{"vi", LangVI, true},
		{"vi-VN", LangVI, true},
		{"fr", LangZH, false}, // 未支持语言回落中文并报告
		{"xx-YY", LangZH, false},
	}
	for _, c := range cases {
		got, ok := Normalize(c.in)
		if got != c.want || ok != c.ok {
			t.Errorf("Normalize(%q) = (%v, %v), want (%v, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestSetAndT(t *testing.T) {
	defer Set("")

	Set("")
	if got := T("help.title"); got != catalogZH["help.title"] {
		t.Errorf("默认语言应为中文，got %q", got)
	}

	Set("en")
	if got := T("help.title"); got != catalogEN["help.title"] {
		t.Errorf("en 语言应返回英文文案，got %q", got)
	}

	Set("vi")
	if got := T("help.title"); got != catalogVI["help.title"] {
		t.Errorf("vi 语言应返回越南语文案，got %q", got)
	}
}

func TestTFallback(t *testing.T) {
	defer Set("")

	// 未知 key 原样返回，方便排查遗漏
	Set("en")
	if got := T("no.such.key"); got != "no.such.key" {
		t.Errorf("未知 key 应原样返回，got %q", got)
	}
}

// TestCatalogParity 保证三份 catalog key 集合完全一致：
// 中文是源头，翻译缺失会在编译期测试中暴露而不是运行时静默回落。
func TestCatalogParity(t *testing.T) {
	for key := range catalogZH {
		if _, ok := catalogEN[key]; !ok {
			t.Errorf("catalogEN 缺少 key %q", key)
		}
		if _, ok := catalogVI[key]; !ok {
			t.Errorf("catalogVI 缺少 key %q", key)
		}
	}
	for key := range catalogEN {
		if _, ok := catalogZH[key]; !ok {
			t.Errorf("catalogEN 存在中文源没有的 key %q", key)
		}
	}
	for key := range catalogVI {
		if _, ok := catalogZH[key]; !ok {
			t.Errorf("catalogVI 存在中文源没有的 key %q", key)
		}
	}
}

// TestCatalogNoEmptyValues 空翻译等同缺失，禁止占位空串。
func TestCatalogNoEmptyValues(t *testing.T) {
	for name, catalog := range map[string]map[string]string{
		"zh": catalogZH, "en": catalogEN, "vi": catalogVI,
	} {
		for key, val := range catalog {
			if val == "" {
				t.Errorf("catalog %s key %q 值为空", name, key)
			}
		}
	}
}
