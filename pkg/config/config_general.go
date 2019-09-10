package config

// General contains the general information about the novel
type General struct {
	Title       string        `yaml:"title"`
	AltTitle    string        `yaml:"alt-title"`
	Author      string        `yaml:"author"`
	Cover       string        `yaml:"cover"`
	Raw         string        `yaml:"raw"`
	Translators []*Translator `yaml:"translators"`
}

// Translator contains the name and website of the translators
type Translator struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}
