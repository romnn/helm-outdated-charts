package main

import (
	"github.com/spf13/cobra"
)

var rootCmdLongUsage = `
Helm plugin to manage outdated dependencies of a Helm chart.

Examples:
  $ helm outdated list <pathToChart> 										- Checks if there's a newer version of any dependency available in the specified repository.
  $ helm outdated list <pathToChart> --repositories repo1.corp,repo2.corp 	- Checks if there's a newer version of any dependency available only using the given repositories.

  $ helm outdated update <pathToChart> 							- Updates all outdated dependencies to the latest version found in the repository.
  $ helm outdated update <pathToChart> --increment-chart-version	- Updates all outdated dependencies to the latest version found in the repository and increments the version of the Helm chart.
`

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "outdated",
		Long:      rootCmdLongUsage,
		ValidArgs: []string{"chartPath"},
	}

	cmd.AddCommand(
		newListOutdatedDependenciesCmd(),
		newUpdateOutdatedDependenciesCmd(),
	)

	return cmd
}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().IntP("max-column-width", "w", 60, "Max column width to use for tables")
	cmd.Flags().
		StringSliceP("repositories", "r", []string{}, "Limit search to the given repository URLs. Can also just provide a part of the URL.")
	cmd.Flags().
		StringSliceP("dependencies", "", []string{}, "Only considers the given dependencies.")
	cmd.Flags().Bool("debug", false, "Enable debug")
}
