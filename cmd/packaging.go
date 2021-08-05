package cmd

import (
	"github.com/spf13/cobra"

	"github.com/sudonims/hover/cmd/packaging"
)

func init() {
	initPackagingCmd.AddCommand(initLinuxSnapCmd)
	initPackagingCmd.AddCommand(initLinuxDebCmd)
	initPackagingCmd.AddCommand(initLinuxAppImageCmd)
	initPackagingCmd.AddCommand(initLinuxRpmCmd)
	initPackagingCmd.AddCommand(initLinuxPkgCmd)
	initPackagingCmd.AddCommand(initWindowsMsiCmd)
	initPackagingCmd.AddCommand(initDarwinBundleCmd)
	initPackagingCmd.AddCommand(initDarwinPkgCmd)
	initPackagingCmd.AddCommand(initDarwinDmgCmd)
	rootCmd.AddCommand(initPackagingCmd)
}

var initPackagingCmd = &cobra.Command{
	Use:   "init-packaging",
	Short: "Create configuration files for a packaging format",
}

var initLinuxSnapCmd = &cobra.Command{
	Use:   "linux-snap",
	Short: "Create configuration files for snap packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.LinuxSnapTask.Init()
	},
}

var initLinuxDebCmd = &cobra.Command{
	Use:   "linux-deb",
	Short: "Create configuration files for deb packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.LinuxDebTask.Init()
	},
}

var initLinuxAppImageCmd = &cobra.Command{
	Use:   "linux-appimage",
	Short: "Create configuration files for AppImage packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.LinuxAppImageTask.Init()
	},
}
var initLinuxRpmCmd = &cobra.Command{
	Use:   "linux-rpm",
	Short: "Create configuration files for rpm packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.LinuxRpmTask.Init()
	},
}
var initLinuxPkgCmd = &cobra.Command{
	Use:   "linux-pkg",
	Short: "Create configuration files for pacman pkg packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.LinuxPkgTask.Init()
	},
}
var initWindowsMsiCmd = &cobra.Command{
	Use:   "windows-msi",
	Short: "Create configuration files for msi packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.WindowsMsiTask.Init()
	},
}

var initDarwinBundleCmd = &cobra.Command{
	Use:   "darwin-bundle",
	Short: "Create configuration files for OSX bundle packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.DarwinBundleTask.Init()
	},
}

var initDarwinPkgCmd = &cobra.Command{
	Use:   "darwin-pkg",
	Short: "Create configuration files for OSX pkg installer packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.DarwinPkgTask.Init()
	},
}

var initDarwinDmgCmd = &cobra.Command{
	Use:   "darwin-dmg",
	Short: "Create configuration files for OSX dmg packaging",
	Run: func(cmd *cobra.Command, args []string) {
		assertHoverInitialized()

		packaging.DarwinDmgTask.Init()
	},
}
