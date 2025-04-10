// domain/observation.go
package domain

// Observation represents a simplified FHIR Observation resource
// 必要に応じてフィールドを拡張可能

type Observation struct {
	ResourceType         string            `json:"resourceType"`
	ID                   string            `json:"id,omitempty"`
	Status               string            `json:"status,omitempty"`
	Category             []CodeableConcept `json:"category,omitempty"`
	Code                 CodeableConcept   `json:"code,omitempty"`
	Subject              Reference         `json:"subject,omitempty"`
	EffectiveDateTime    string            `json:"effectiveDateTime,omitempty"`
	ValueQuantity        *Quantity         `json:"valueQuantity,omitempty"`
	ValueString          string            `json:"valueString,omitempty"`
	ValueCodeableConcept *CodeableConcept  `json:"valueCodeableConcept,omitempty"`
}

type Quantity struct {
	Value  float64 `json:"value,omitempty"`
	Unit   string  `json:"unit,omitempty"`
	System string  `json:"system,omitempty"`
	Code   string  `json:"code,omitempty"`
}

type Reference struct {
	Reference string `json:"reference,omitempty"`
	Display   string `json:"display,omitempty"`
}
