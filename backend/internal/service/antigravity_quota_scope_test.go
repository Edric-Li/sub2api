//go:build unit

package service

import "testing"

func TestResolveAntigravityQuotaScope(t *testing.T) {
	tests := []struct {
		name          string
		model         string
		expectedScope AntigravityQuotaScope
		expectedOK    bool
	}{
		// claude 系列
		{name: "claude-sonnet-4-5", model: "claude-sonnet-4-5", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude-opus-4-6", model: "claude-opus-4-6", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude-haiku-3-5", model: "claude-haiku-3-5", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude with thinking suffix", model: "claude-sonnet-4-5-thinking", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude uppercase", model: "CLAUDE-OPUS-4-6", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude with models/ prefix", model: "models/claude-sonnet-4-5", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},
		{name: "claude with spaces", model: "  claude-sonnet-4-5  ", expectedScope: AntigravityQuotaScopeClaude, expectedOK: true},

		// gemini 文本系列
		{name: "gemini-3-flash", model: "gemini-3-flash", expectedScope: AntigravityQuotaScopeGeminiText, expectedOK: true},
		{name: "gemini-3-pro-high", model: "gemini-3-pro-high", expectedScope: AntigravityQuotaScopeGeminiText, expectedOK: true},
		{name: "gemini-2.5-flash", model: "gemini-2.5-flash", expectedScope: AntigravityQuotaScopeGeminiText, expectedOK: true},
		{name: "gemini with models/ prefix", model: "models/gemini-3-flash", expectedScope: AntigravityQuotaScopeGeminiText, expectedOK: true},

		// gemini 图像系列
		{name: "gemini-3-pro-image", model: "gemini-3-pro-image", expectedScope: AntigravityQuotaScopeGeminiImage, expectedOK: true},
		{name: "gemini-3-pro-image-preview", model: "gemini-3-pro-image-preview", expectedScope: AntigravityQuotaScopeGeminiImage, expectedOK: true},
		{name: "gemini-2.5-flash-image", model: "gemini-2.5-flash-image", expectedScope: AntigravityQuotaScopeGeminiImage, expectedOK: true},
		{name: "gemini-2.5-flash-image-preview", model: "gemini-2.5-flash-image-preview", expectedScope: AntigravityQuotaScopeGeminiImage, expectedOK: true},
		{name: "gemini image with models/ prefix", model: "models/gemini-3-pro-image", expectedScope: AntigravityQuotaScopeGeminiImage, expectedOK: true},

		// 不支持的模型
		{name: "empty string", model: "", expectedScope: "", expectedOK: false},
		{name: "gpt-4", model: "gpt-4", expectedScope: "", expectedOK: false},
		{name: "llama-3", model: "llama-3", expectedScope: "", expectedOK: false},
		{name: "bare claude (no dash)", model: "claude", expectedScope: "", expectedOK: false},
		{name: "bare gemini (no dash)", model: "gemini", expectedScope: "", expectedOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scope, ok := resolveAntigravityQuotaScope(tt.model)
			if ok != tt.expectedOK {
				t.Errorf("resolveAntigravityQuotaScope(%q) ok = %v, want %v", tt.model, ok, tt.expectedOK)
			}
			if scope != tt.expectedScope {
				t.Errorf("resolveAntigravityQuotaScope(%q) scope = %q, want %q", tt.model, scope, tt.expectedScope)
			}
		})
	}
}
