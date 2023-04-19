package main

import (
	"fmt"
	manager "github.com/LiveAlone/GoLibSourceAnalyse/codes/manage"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var Root = &cobra.Command{
	Use: "patch",
}

var dir string

var patch = &cobra.Command{
	Use: "patch",
	Run: func(cmd *cobra.Command, args []string) {
		patcher := manager.NewPatcher(dir, -1)
		if err := patcher.Backup(); err != nil {
			log.Fatalf("backup error: %v", err)
		}

		if err := patcher.Inject(); err != nil {
			log.Fatalf("inject error: %v", err)
		}
	},
}

var cover = &cobra.Command{
	Use: "recover",
	Run: func(cmd *cobra.Command, args []string) {
		patcher := manager.NewPatcher(dir, -1)
		if err := patcher.Recover(); err != nil {
			log.Fatalf("recover error: %v", err)
		}
	},
}

func init() {
	patch.Flags().StringVarP(&dir, "dir", "d", "", "source code root dir")
	cover.Flags().StringVarP(&dir, "dir", "d", "", "source code root dir")
}

func main() {
	Root.AddCommand(patch)
	Root.AddCommand(cover)
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
