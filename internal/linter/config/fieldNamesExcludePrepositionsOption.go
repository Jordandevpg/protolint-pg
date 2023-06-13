package config

// FieldNamesExcludePrepositionsOption represents the option for the FIELD_NAMES_EXCLUDE_PREPOSITIONS rule.
type FieldNamesExcludePrepositionsOption struct {
	CustomizableSeverityOption
	Prepositions []string `yaml:"prepositions"`
	Excludes     []string `yaml:"excludes"`
}
