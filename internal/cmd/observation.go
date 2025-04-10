// cmd/observation.go
package cmd

import (
	"fmt"
	"os"

	"github.com/trkd-knt/fhir-sandbox/internal/infrastructure"
	"github.com/trkd-knt/fhir-sandbox/internal/usecase"

	"github.com/spf13/cobra"
)

var (
	obsInput  string
	obsOutput string
	obsToCSV  bool
	obsToJSON bool
)

var observationCmd = &cobra.Command{
	Use:   "observation",
	Short: "Convert FHIR Observation resource (CSV ⇄ JSON)",
	Run: func(cmd *cobra.Command, args []string) {
		if obsInput == "" || obsOutput == "" {
			fmt.Println("[ERROR] --input と --output の指定が必要です")
			os.Exit(1)
		}
		if obsToCSV == obsToJSON {
			fmt.Println("[ERROR] --to-csv か --to-json のどちらか一方を指定してください")
			os.Exit(1)
		}

		if obsToCSV {
			fmt.Println("[INFO] JSON → CSV に変換中 (Observation)...")
			jsonBytes, err := infrastructure.ReadFile(obsInput)
			if err != nil {
				fmt.Println("[ERROR] JSON読み込み失敗:", err)
				os.Exit(1)
			}
			headers, records, err := usecase.ConvertObservationJSONToCSV(jsonBytes)
			if err != nil {
				fmt.Println("[ERROR] 変換失敗:", err)
				os.Exit(1)
			}
			if err := infrastructure.WriteCSVFile(obsOutput, headers, records); err != nil {
				fmt.Println("[ERROR] CSV書き込み失敗:", err)
				os.Exit(1)
			}
			fmt.Println("[SUCCESS] CSV出力完了:", obsOutput)
		} else if obsToJSON {
			fmt.Println("[INFO] CSV → JSON に変換中 (Observation)...")
			headers, records, err := infrastructure.ReadCSVFile(obsInput)
			if err != nil {
				fmt.Println("[ERROR] CSV読み込み失敗:", err)
				os.Exit(1)
			}
			jsonBytes, err := usecase.ConvertObservationCSVToJSON(headers, records)
			if err != nil {
				fmt.Println("[ERROR] 変換失敗:", err)
				os.Exit(1)
			}
			if err := infrastructure.WriteFile(obsOutput, jsonBytes); err != nil {
				fmt.Println("[ERROR] JSON書き込み失敗:", err)
				os.Exit(1)
			}
			fmt.Println("[SUCCESS] JSON出力完了:", obsOutput)
		}
	},
}

func init() {
	observationCmd.Flags().StringVar(&obsInput, "input", "", "入力ファイルのパス")
	observationCmd.Flags().StringVar(&obsOutput, "output", "", "出力ファイルのパス")
	observationCmd.Flags().BoolVar(&obsToCSV, "to-csv", false, "JSON → CSV に変換")
	observationCmd.Flags().BoolVar(&obsToJSON, "to-json", false, "CSV → JSON に変換")
}
