/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	oTag    string
	oName   string
	oNotes  string
	oAssets []string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new release",
	Long: `Create a new release on the configured Git hosting platform.

This command creates a release with the specified tag, name, and optional notes and assets.
The tag must be a valid Git tag that exists in the repository.`,
	Run: Run,
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Required flags
	createCmd.Flags().StringVarP(&oTag, "tag", "t", "", "Tag referenced by the release (required)")
	createCmd.Flags().StringVarP(&oName, "name", "n", "", "Release name (required)")

	// Mark required flags
	createCmd.MarkFlagRequired("tag")
	createCmd.MarkFlagRequired("name")

	// Optional flags
	createCmd.Flags().StringVar(&oNotes, "notes", "", "Release notes")
	createCmd.Flags().StringSliceVar(&oAssets, "assets", []string{}, "Paths or directories to include as release assets")
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("create called")
}
