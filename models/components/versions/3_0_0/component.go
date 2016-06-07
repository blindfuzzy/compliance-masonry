package component

import (
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/version"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
)

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    *common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications *common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies          `yaml:"satisfies" json:"satisfies"`
	base.Base
}

// VerifySchemaCompatibility will check that the current component schema version is
// compatible with the current masonry toolchain.
func (c *Component) VerifySchemaCompatibility(fileName string) error {
	if c != nil {
		requirements := version.NewRequirements(fileName, "component", c.SchemaVersion,
			constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion)
		return requirements.VerifyVersion()
	}
	return nil
}

func (c Component) GetVerifications() *common.VerificationReferences {
	return c.Verifications
}

func (c Component) GetReferences() *common.GeneralReferences {
	return c.References
}

func (c Component) GetName() string {
	return c.Name
}

func (c Component) GetKey() string {
	return c.Key
}

func (c *Component) SetKey(key string) {
	c.Key = key
}

func (c Component) GetAllSatisfies() []base.Satisfies {
	// Have to do manual conversion
	baseSatisfies := make([]base.Satisfies, len(c.Satisfies))
	for idx, value := range c.Satisfies {
		baseSatisfies[idx] = value
	}
	return baseSatisfies
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string             `yaml:"control_key" json:"control_key"`
	StandardKey string             `yaml:"standard_key" json:"standard_key"`
	Narrative   []NarrativeSection `yaml:"narrative" json:"narrative"`
	CoveredBy   common.CoveredByList      `yaml:"covered_by" json:"covered_by"`
}


func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

func (s Satisfies)GetControlKey() string {
	return s.ControlKey
}

func (s Satisfies)GetStandardKey() string {
	return s.StandardKey
}

func (s Satisfies) GetNarratives() []base.Narrative {
	// Have to do manual conversion
	baseNarrative := make([]base.Narrative, len(s.Narrative))
	for idx, value := range s.Narrative {
		baseNarrative[idx] = value
	}
	return baseNarrative
}

// NarrativeSection contains the key and text for a particular narrative section.
type NarrativeSection struct {
	Key  string `yaml:"key,omitempty" json:"key,omitempty"`
	Text string `yaml:"text" json:"text"`
}

func (ns NarrativeSection) GetKey() string {
	return ns.Key
}

func (ns NarrativeSection) GetText() string {
	return ns.Text
}
