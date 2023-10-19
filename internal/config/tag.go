package config

type Tag struct {
	Id          int64    `json:"id,omitempty" yaml:"id" xml:"id" toml:"id"`
	Name        string   `json:"name,omitempty" yaml:"name" xml:"name" toml:"name"`
	Label       string   `json:"label,omitempty" yaml:"label" xml:"label" toml:"label"`
	Keywords    []string `json:"keywords,omitempty" yaml:"keywords" xml:"keywords" toml:"keywords"`
	Description string   `json:"description,omitempty" yaml:"description" xml:"description" toml:"description"`
}
