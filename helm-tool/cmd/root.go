/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/iluoy/helm-tool/pkg/option"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "helm-tool",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	globalRootOptions = option.NewGlobalRootOptions()
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringSliceVar(&globalRootOptions.Namespaces, "namespace", []string{}, "all release in specific kubernetes namespaces")
	rootCmd.PersistentFlags().StringSliceVar(&globalRootOptions.Releases, "releases", []string{}, "specific the release name and namespace, i.e: release1.namespace1")
	rootCmd.PersistentFlags().BoolVarP(&globalRootOptions.AllNamespaces, "all-namespaces", "A", false, "all releases in all namespaces")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.helm-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
