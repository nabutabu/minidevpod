/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	//"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "miniDevPod",
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

var createCmd = &cobra.Command{
	Use:   "create a dev pod",
	Short: "Short Description of create",
	Long:  "Long Description of create",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			return err
		}
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			return err
		}
		cpu, err := cmd.Flags().GetString("cpu")
		if err != nil {
			return err
		}
		memory, err := cmd.Flags().GetString("memory")
		if err != nil {
			return err
		}

		CreatePod(name, repo, branch, cpu, memory)

		return err
	},
}

var connectCmd = &cobra.Command{
	Use:     "Connect to an existing pod",
	Aliases: []string{"connect", "ssh"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Connect(args[0])
		return err
	},
}

var listCmd = &cobra.Command{
	Use:     "List all pods created by mini-devpod",
	Aliases: []string{"list", "ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return List()
	},
}

var deleteCmd = &cobra.Command{
	Use:     "Delete a devpod",
	Aliases: []string{"delete", "rm"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Delete(args[0])
	},
}

var forwardCmd = &cobra.Command{
	Use:  "Forward a port from host machine to devpod",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ports := strings.Split(args[1], ":")

		localPort, _ := strconv.Atoi(ports[0])
		remotePort, _ := strconv.Atoi(ports[1])

		return Forward(name, localPort, remotePort)
	},
}

var syncCmd = &cobra.Command{
	Use:  "Sync files on host machine to pod",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Sync(args[0], args[1])
	},
}

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.miniDevPod.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().String("name", "", "Name of new pod")
	createCmd.Flags().String("repo", "", "Repo you want in devpod")
	createCmd.Flags().String("branch", "main", "Branch to copy")
	createCmd.Flags().String("cpu", "500m", "CPU request")
	createCmd.Flags().String("memory", "1Gi", "Memory request")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(forwardCmd)
	rootCmd.AddCommand(syncCmd)
}
