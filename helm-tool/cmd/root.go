/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/iluoy/helm-tool/pkg/option"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
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
	}

	pflags := rootCmd.PersistentFlags()
	globalRootOptions := option.GlobalRootOptions{}
	addFlags(pflags, globalRootOptions)
	pflags.Parse(args)
	return rootCmd

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewRootCmd(os.Args[1:]).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addFlags(pflag *pflag.FlagSet, globalRootOptions option.GlobalRootOptions) {
	pflag.StringSliceVar(&globalRootOptions.Namespaces, "namespace", []string{}, "all release in specific kubernetes namespaces")
	pflag.StringSliceVar(&globalRootOptions.Releases, "releases", []string{}, "specific the release name and namespace, i.e: release1.namespace1")
	pflag.BoolVarP(&globalRootOptions.AllNamespaces, "all-namespaces", "A", false, "all releases in all namespaces")
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.helm-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
