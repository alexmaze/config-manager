package main

// Rule rule
type Rule struct {
	Name  string     `mapstructure:"name"`
	Items []RuleItem `mapstructure:"item"`
}

// Type configuration type
type Type string

// Types
const (
	TypeFile   Type = "file"
	TypeFolder Type = "folder"
)

// RuleItem files || folders
type RuleItem struct {
	Type        Type   `mapstructure:"-"` // Auto detected by `Source`
	Source      string `mapstructure:"src"`
	Destination string `mapstructure:"dest"`
}
