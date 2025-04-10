// cmd/patient.go
package cmd

import (
	"fmt"
	"os"

	"github.com/trkd-knt/fhir-sandbox/internal/infrastructure"
	"github.com/trkd-knt/fhir-sandbox/internal/usecase"

	"github.com/spf13/cobra"
)

var (
	patientInput  string
	patientOutput string
	toCSV         bool
	toJSON        bool
)

var patientCmd = &cobra.Command{
	Use:   "patient",
	Short: "Convert FHIR Patient resource (CSV ⇄ JSON)",
	Run: func(cmd *cobra.Command, args []string) {
		if patientInput == "" || patientOutput == "" {
			fmt.Println("[ERROR] --input と --output の指定が必要です")
			os.Exit(1)
		}
		if toCSV == toJSON {
			fmt.Println("[ERROR] --to-csv か --to-json のどちらか一方を指定してください")
			os.Exit(1)
		}

		if toCSV {
			fmt.Println("[INFO] JSON → CSV に変換中 (Patient)...")
			jsonBytes, err := infrastructure.ReadFile(patientInput)
			if err != nil {
				fmt.Println("[ERROR] JSON読み込み失敗:", err)
				os.Exit(1)
			}
			headers, records, err := usecase.ConvertPatientJSONToCSV(jsonBytes)
			if err != nil {
				fmt.Println("[ERROR] 変換失敗:", err)
				os.Exit(1)
			}
			if err := infrastructure.WriteCSVFile(patientOutput, headers, records); err != nil {
				fmt.Println("[ERROR] CSV書き込み失敗:", err)
				os.Exit(1)
			}
			fmt.Println("[SUCCESS] CSV出力完了:", patientOutput)
		} else if toJSON {
			fmt.Println("[INFO] CSV → JSON に変換中 (Patient)...")
			headers, records, err := infrastructure.ReadCSVFile(patientInput)
			if err != nil {
				fmt.Println("[ERROR] CSV読み込み失敗:", err)
				os.Exit(1)
			}
			jsonBytes, err := usecase.ConvertPatientCSVToJSON(headers, records)
			if err != nil {
				fmt.Println("[ERROR] 変換失敗:", err)
				os.Exit(1)
			}
			if err := infrastructure.WriteFile(patientOutput, jsonBytes); err != nil {
				fmt.Println("[ERROR] JSON書き込み失敗:", err)
				os.Exit(1)
			}
			fmt.Println("[SUCCESS] JSON出力完了:", patientOutput)
		}
	},
}

func init() {
	patientCmd.Flags().StringVar(&patientInput, "input", "", "入力ファイルのパス")
	patientCmd.Flags().StringVar(&patientOutput, "output", "", "出力ファイルのパス")
	patientCmd.Flags().BoolVar(&toCSV, "to-csv", false, "JSON → CSV に変換")
	patientCmd.Flags().BoolVar(&toJSON, "to-json", false, "CSV → JSON に変換")
}
