// internal/domain/patient.go
package domain

// FHIR Patient Resource - 必要最小限の構造体
// 今後フィールド追加で拡張可能（map[string]interface{}で柔軟にしてもOK）

type Patient struct {
	ResourceType  string           `json:"resourceType"`
	ID            string           `json:"id,omitempty"`
	Active        *bool            `json:"active,omitempty"`
	Name          []HumanName      `json:"name,omitempty"`
	Telecom       []Contact        `json:"telecom,omitempty"`
	Gender        string           `json:"gender,omitempty"`
	BirthDate     string           `json:"birthDate,omitempty"`
	Address       []Address        `json:"address,omitempty"`
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	// その他拡張フィールドは随時追加
}

type HumanName struct {
	Use    string   `json:"use,omitempty"`
	Family string   `json:"family,omitempty"`
	Given  []string `json:"given,omitempty"`
}

type Contact struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
	Use    string `json:"use,omitempty"`
}

type Address struct {
	Use        string   `json:"use,omitempty"`
	Type       string   `json:"type,omitempty"`
	Text       string   `json:"text,omitempty"`
	Line       []string `json:"line,omitempty"`
	City       string   `json:"city,omitempty"`
	PostalCode string   `json:"postalCode,omitempty"`
	Country    string   `json:"country,omitempty"`
}

type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}

type Coding struct {
	System  string `json:"system,omitempty"`
	Code    string `json:"code,omitempty"`
	Display string `json:"display,omitempty"`
}
