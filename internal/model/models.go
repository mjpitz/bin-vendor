package model

// GitHubRelease defines information needed for downloading an asset from GitHub.
type GitHubRelease struct {
	Repository   string            `json:"repo,omitempty"`
	Replacements map[string]string `json:"replacements,omitempty"`
	Format       string            `json:"format,omitempty"`
	Linux        string            `json:"linux,omitempty"`
	OSX          string            `json:"osx,omitempty"`
	Windows      string            `json:"windows,omitempty"`
}

// Tool defines the basic requirements for a tool.
type Tool struct {
	Name          string         `json:"name,omitempty"`
	Version       string         `json:"version,omitempty"`
	GitHubRelease *GitHubRelease `json:"github_release,omitempty"`
}

// Config defines the structure of the config file.
type Config struct {
	Name  string  `json:"name,omitempty"`
	Tools []*Tool `json:"tools,omitempty"`
}

// Render defines a structure used to render URLs with metadata about a tool.
type Render struct {
	Name    string
	Version string
	OS      string
	Arch    string
}