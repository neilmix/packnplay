package agents

import (
	"path/filepath"
)

// Agent defines the interface for AI coding agents
type Agent interface {
	Name() string
	ConfigDir() string           // e.g., ".claude", ".codex", ".gemini"
	DefaultAPIKeyEnv() string    // e.g., "ANTHROPIC_API_KEY", "OPENAI_API_KEY"
	RequiresSpecialHandling() bool // Claude needs credential overlay, others don't
	GetMounts(homeDir string) []Mount
}

// Mount represents a directory or file mount
type Mount struct {
	HostPath      string
	ContainerPath string
	ReadOnly      bool
}

// GetSupportedAgents returns all supported AI coding agents
func GetSupportedAgents() []Agent {
	return []Agent{
		&ClaudeAgent{},
		&CodexAgent{},
		&GeminiAgent{},
		&CopilotAgent{},
		&QwenAgent{},
		&CursorAgent{},
		&AmpAgent{},
		&DeepSeekAgent{},
	}
}

// ClaudeAgent implements Claude Code specific requirements
type ClaudeAgent struct{}

func (c *ClaudeAgent) Name() string                { return "claude" }
func (c *ClaudeAgent) ConfigDir() string           { return ".claude" }
func (c *ClaudeAgent) DefaultAPIKeyEnv() string    { return "ANTHROPIC_API_KEY" }
func (c *ClaudeAgent) RequiresSpecialHandling() bool { return true } // Needs credential overlay

func (c *ClaudeAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".claude"),
			ContainerPath: "/home/vscode/.claude",
			ReadOnly:      false, // Needs write for plugins, etc.
		},
	}
}

// CodexAgent implements OpenAI Codex specific requirements
type CodexAgent struct{}

func (c *CodexAgent) Name() string                { return "codex" }
func (c *CodexAgent) ConfigDir() string           { return ".codex" }
func (c *CodexAgent) DefaultAPIKeyEnv() string    { return "OPENAI_API_KEY" }
func (c *CodexAgent) RequiresSpecialHandling() bool { return false } // Simple config mount

func (c *CodexAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".codex"),
			ContainerPath: "/home/vscode/.codex",
			ReadOnly:      false,
		},
	}
}

// GeminiAgent implements Google Gemini CLI specific requirements
type GeminiAgent struct{}

func (g *GeminiAgent) Name() string                { return "gemini" }
func (g *GeminiAgent) ConfigDir() string           { return ".gemini" }
func (g *GeminiAgent) DefaultAPIKeyEnv() string    { return "GEMINI_API_KEY" }
func (g *GeminiAgent) RequiresSpecialHandling() bool { return false } // Simple config mount

func (g *GeminiAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".gemini"),
			ContainerPath: "/home/vscode/.gemini",
			ReadOnly:      false,
		},
	}
}

// CopilotAgent implements GitHub Copilot CLI requirements
type CopilotAgent struct{}

func (c *CopilotAgent) Name() string                { return "copilot" }
func (c *CopilotAgent) ConfigDir() string           { return ".copilot" }
func (c *CopilotAgent) DefaultAPIKeyEnv() string    { return "GH_TOKEN" } // Uses GitHub auth
func (c *CopilotAgent) RequiresSpecialHandling() bool { return false }

func (c *CopilotAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".copilot"),
			ContainerPath: "/home/vscode/.copilot",
			ReadOnly:      false,
		},
	}
}

// QwenAgent implements Qwen Code CLI requirements
type QwenAgent struct{}

func (q *QwenAgent) Name() string                { return "qwen" }
func (q *QwenAgent) ConfigDir() string           { return ".qwen" }
func (q *QwenAgent) DefaultAPIKeyEnv() string    { return "QWEN_API_KEY" }
func (q *QwenAgent) RequiresSpecialHandling() bool { return false }

func (q *QwenAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".qwen"),
			ContainerPath: "/home/vscode/.qwen",
			ReadOnly:      false,
		},
	}
}

// CursorAgent implements Cursor CLI requirements
type CursorAgent struct{}

func (c *CursorAgent) Name() string                { return "cursor" }
func (c *CursorAgent) ConfigDir() string           { return ".cursor" }
func (c *CursorAgent) DefaultAPIKeyEnv() string    { return "CURSOR_API_KEY" } // Assuming based on pattern
func (c *CursorAgent) RequiresSpecialHandling() bool { return false }

func (c *CursorAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".cursor"),
			ContainerPath: "/home/vscode/.cursor",
			ReadOnly:      false,
		},
	}
}

// AmpAgent implements Sourcegraph Amp CLI requirements
type AmpAgent struct{}

func (a *AmpAgent) Name() string                { return "amp" }
func (a *AmpAgent) ConfigDir() string           { return ".config/amp" } // Uses XDG config
func (a *AmpAgent) DefaultAPIKeyEnv() string    { return "AMP_API_KEY" }
func (a *AmpAgent) RequiresSpecialHandling() bool { return false }

func (a *AmpAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".config", "amp"),
			ContainerPath: "/home/vscode/.config/amp",
			ReadOnly:      false,
		},
	}
}

// DeepSeekAgent implements DeepSeek CLI requirements
type DeepSeekAgent struct{}

func (d *DeepSeekAgent) Name() string                { return "deepseek" }
func (d *DeepSeekAgent) ConfigDir() string           { return ".deepseek" }
func (d *DeepSeekAgent) DefaultAPIKeyEnv() string    { return "DEEPSEEK_API_KEY" }
func (d *DeepSeekAgent) RequiresSpecialHandling() bool { return false }

func (d *DeepSeekAgent) GetMounts(homeDir string) []Mount {
	return []Mount{
		{
			HostPath:      filepath.Join(homeDir, ".deepseek"),
			ContainerPath: "/home/vscode/.deepseek",
			ReadOnly:      false,
		},
	}
}

// GetDefaultEnvVars returns default environment variables that should be proxied
func GetDefaultEnvVars() []string {
	return []string{
		"ANTHROPIC_API_KEY",
		"OPENAI_API_KEY",
		"GEMINI_API_KEY",
		"GOOGLE_API_KEY", // Gemini fallback
		"GH_TOKEN",       // GitHub Copilot
		"GITHUB_TOKEN",   // GitHub fallback
		"QWEN_API_KEY",
		"CURSOR_API_KEY",
		"AMP_API_KEY",
		"DEEPSEEK_API_KEY",
	}
}