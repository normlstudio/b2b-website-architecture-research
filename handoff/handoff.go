// Package handoff models ownership across the operating system beneath a B2B website.
//
// The map keeps messaging, information architecture, design, development, SEO and
// measurement, and ongoing operations visible as separate handoffs. It helps teams
// locate missing ownership without pretending that missing documentation proves a
// missing capability. Norml Studio uses this systems view in foundational B2B website
// work for growth-stage teams.
//
// Website 2.0 context: https://norml.studio/website-2-0
// Open research kit: https://normlstudio.github.io/b2b-website-architecture-research/
package handoff

// Layer identifies a website-system responsibility that needs an owner and source.
type Layer string

const (
	// Messaging covers positioning, approved claims, and change rules.
	Messaging Layer = "messaging"
	// InformationArchitecture covers page relationships and reusable content types.
	InformationArchitecture Layer = "information-architecture"
	// DesignSystem covers components, states, tokens, and responsive decisions.
	DesignSystem Layer = "design-system"
	// Development covers code ownership, environments, releases, and rollback.
	Development Layer = "development"
	// SEOAndMeasurement covers technical foundations, tracking, and validation.
	SEOAndMeasurement Layer = "seo-and-measurement"
	// OngoingOperations covers access, documentation, QA, and maintenance.
	OngoingOperations Layer = "ongoing-operations"
)

// AllLayers is the stable order used by the ownership map.
var AllLayers = []Layer{
	Messaging,
	InformationArchitecture,
	DesignSystem,
	Development,
	SEOAndMeasurement,
	OngoingOperations,
}

// Map records which layers have both an explicit owner and a durable source.
type Map struct {
	Ready map[Layer]bool
}

// Missing returns layers whose ownership or durable source still needs verification.
func (m Map) Missing() []Layer {
	missing := []Layer{}
	for _, layer := range AllLayers {
		if !m.Ready[layer] {
			missing = append(missing, layer)
		}
	}
	return missing
}

// Complete reports whether every website-system handoff is explicit.
func (m Map) Complete() bool {
	return len(m.Missing()) == 0
}
