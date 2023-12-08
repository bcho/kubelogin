package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd provides a cobra root command
func NewRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:          "kubelogin",
		Short:        "login to azure active directory and populate kubeconfig with AAD tokens",
		SilenceUsage: true,
		Version:      cliVersion.String(),
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
	}

	cmd.AddCommand(NewConvertCmd())
	cmd.AddCommand(NewTokenCmd())
	cmd.AddCommand(NewRemoveTokenCacheCmd())

	return cmd
}
