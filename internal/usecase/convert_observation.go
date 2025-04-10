// usecase/convert_observation.go
package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/trkd-knt/fhir-sandbox/internal/domain"
)

func ConvertObservationJSONToCSV(input []byte) ([]string, [][]string, error) {
	var bundle struct {
		ResourceType string `json:"resourceType"`
		Type         string `json:"type"`
		Entry        []struct {
			Resource domain.Observation `json:"resource"`
		} `json:"entry"`
	}
	if err := json.Unmarshal(input, &bundle); err != nil {
		return nil, nil, err
	}

	var flatObservations []map[string]string
	allKeys := map[string]bool{}

	for _, entry := range bundle.Entry {
		flat := flattenObservation(entry.Resource)
		flatObservations = append(flatObservations, flat)
		for k := range flat {
			allKeys[k] = true
		}
	}

	headers := make([]string, 0, len(allKeys))
	for k := range allKeys {
		headers = append(headers, k)
	}
	sort.Strings(headers)

	records := [][]string{}
	for _, flat := range flatObservations {
		row := make([]string, len(headers))
		for i, k := range headers {
			row[i] = flat[k]
		}
		records = append(records, row)
	}

	return headers, records, nil
}

func ConvertObservationCSVToJSON(headers []string, records [][]string) ([]byte, error) {
	var entries []map[string]interface{}

	for _, row := range records {
		data := map[string]interface{}{}
		for i, col := range headers {
			if val := strings.TrimSpace(row[i]); val != "" {
				assignNestedField(data, col, val)
			}
		}
		data["resourceType"] = "Observation"
		entries = append(entries, map[string]interface{}{
			"resource": data,
			"request": map[string]string{
				"method": "PUT",
				"url":    fmt.Sprintf("Observation/%v", data["id"]),
			},
		})
	}

	bundle := map[string]interface{}{
		"resourceType": "Bundle",
		"type":         "transaction",
		"entry":        entries,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(bundle); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func flattenObservation(obs domain.Observation) map[string]string {
	flat := map[string]string{}
	flat["resourceType"] = obs.ResourceType
	flat["id"] = obs.ID
	flat["status"] = obs.Status
	flat["effectiveDateTime"] = obs.EffectiveDateTime

	flat["code_text"] = obs.Code.Text
	if len(obs.Code.Coding) > 0 {
		flat["code_coding_0_code"] = obs.Code.Coding[0].Code
		flat["code_coding_0_display"] = obs.Code.Coding[0].Display
	}

	if obs.ValueQuantity != nil {
		flat["valueQuantity_value"] = fmt.Sprintf("%f", obs.ValueQuantity.Value)
		flat["valueQuantity_unit"] = obs.ValueQuantity.Unit
	}

	if obs.ValueString != "" {
		flat["valueString"] = obs.ValueString
	}

	if obs.ValueCodeableConcept != nil {
		flat["valueCodeableConcept_text"] = obs.ValueCodeableConcept.Text
	}

	if obs.Subject.Reference != "" {
		flat["subject_reference"] = obs.Subject.Reference
		flat["subject_display"] = obs.Subject.Display
	}

	return flat
}
