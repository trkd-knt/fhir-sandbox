// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fhir-convert",
	Short: "FHIR <-> CSV Converter CLI",
	Long:  `A CLI to convert FHIR Bundle JSON to/from flat CSV for Patient, Observation, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// サブコマンド登録
	rootCmd.AddCommand(patientCmd)
	rootCmd.AddCommand(observationCmd)
}
