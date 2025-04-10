// usecase/convert_patient.go
package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/trkd-knt/fhir-sandbox/internal/domain"
)

// ConvertPatientJSONToCSV converts FHIR Bundle JSON (with Patient entries) to flat CSV
func ConvertPatientJSONToCSV(input []byte) ([]string, [][]string, error) {
	var bundle struct {
		ResourceType string `json:"resourceType"`
		Type         string `json:"type"`
		Entry        []struct {
			Resource domain.Patient `json:"resource"`
		} `json:"entry"`
	}
	if err := json.Unmarshal(input, &bundle); err != nil {
		return nil, nil, err
	}

	var flatPatients []map[string]string
	allKeys := map[string]bool{}

	for _, entry := range bundle.Entry {
		flat := flattenPatient(entry.Resource)
		flatPatients = append(flatPatients, flat)
		for k := range flat {
			allKeys[k] = true
		}
	}

	// ヘッダーの順序を固定（アルファベット順）
	headers := make([]string, 0, len(allKeys))
	for k := range allKeys {
		headers = append(headers, k)
	}
	sort.Strings(headers)

	// データ行作成
	records := [][]string{}
	for _, flat := range flatPatients {
		row := make([]string, len(headers))
		for i, k := range headers {
			row[i] = flat[k]
		}
		records = append(records, row)
	}

	return headers, records, nil
}

// flattenPatient converts a Patient into a flat map[string]string
func flattenPatient(p domain.Patient) map[string]string {
	flat := map[string]string{}
	flat["resourceType"] = p.ResourceType
	flat["id"] = p.ID
	flat["gender"] = p.Gender
	flat["birthDate"] = p.BirthDate
	if p.Active != nil {
		flat["active"] = strconv.FormatBool(*p.Active)
	}

	for i, name := range p.Name {
		flat[fmt.Sprintf("name_%d_family", i)] = name.Family
		flat[fmt.Sprintf("name_%d_use", i)] = name.Use
		for j, given := range name.Given {
			flat[fmt.Sprintf("name_%d_given_%d", i, j)] = given
		}
	}

	for i, tel := range p.Telecom {
		flat[fmt.Sprintf("telecom_%d_system", i)] = tel.System
		flat[fmt.Sprintf("telecom_%d_value", i)] = tel.Value
		flat[fmt.Sprintf("telecom_%d_use", i)] = tel.Use
	}

	for i, addr := range p.Address {
		flat[fmt.Sprintf("address_%d_city", i)] = addr.City
		flat[fmt.Sprintf("address_%d_country", i)] = addr.Country
		flat[fmt.Sprintf("address_%d_postalCode", i)] = addr.PostalCode
		for j, line := range addr.Line {
			flat[fmt.Sprintf("address_%d_line_%d", i, j)] = line
		}
	}

	return flat
}

// ConvertPatientCSVToJSON converts flat CSV (header + rows) to FHIR Bundle JSON with Patient entries
func ConvertPatientCSVToJSON(headers []string, records [][]string) ([]byte, error) {
	var entries []map[string]interface{}

	for _, row := range records {
		data := map[string]interface{}{}
		for i, col := range headers {
			if val := strings.TrimSpace(row[i]); val != "" {
				assignNestedField(data, col, val)
			}
		}
		entries = append(entries, map[string]interface{}{
			"resource": data,
			"request": map[string]string{
				"method": "PUT",
				"url":    fmt.Sprintf("Patient/%v", data["id"]),
			},
		})
		data["resourceType"] = "Patient"
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

// assignNestedField builds nested map structure from flat keys like name_0_given_0
func assignNestedField(data map[string]interface{}, flatKey string, value string) {
	parts := strings.Split(flatKey, "_")
	cur := data
	for i := 0; i < len(parts)-1; i++ {
		key := parts[i]
		// 配列っぽいか？
		idx, err := strconv.Atoi(parts[i+1])
		if err == nil {
			// parts[i]は配列キー、次がインデックス
			if _, ok := cur[key]; !ok {
				cur[key] = []interface{}{}
			}
			arr := cur[key].([]interface{})
			for len(arr) <= idx {
				arr = append(arr, map[string]interface{}{})
			}
			cur[key] = arr
			// 次のマップへ
			cur = arr[idx].(map[string]interface{})
			i++ // idxぶんスキップ
		} else {
			// 普通のキー
			if _, ok := cur[key]; !ok {
				cur[key] = map[string]interface{}{}
			}
			cur = cur[key].(map[string]interface{})
		}
	}
	cur[parts[len(parts)-1]] = value
}
