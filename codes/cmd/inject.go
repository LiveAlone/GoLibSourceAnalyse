package main

import (
	"fmt"
	brownfox "github.com/LiveAlone/GoLibSourceAnalyse/codes/manage"
	"github.com/spf13/cobra"
	"os"
)

var Root = &cobra.Command{
	Use: "inject",
}

var dir string

var jump = &cobra.Command{
	Use: "jump",
	Run: func(cmd *cobra.Command, args []string) {
		bf := brownfox.NewBrownFox(dir, -1)
		if err := bf.Backup(); err != nil {
			panic(err)
		}
		if err := bf.Inject(); err != nil {
			panic(err)
		}
	},
}

var over = &cobra.Command{
	Use: "over",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("over")
		//bf := brownfox.NewBrownFox(dir, -1)
		//if err := bf.Restore(); err != nil {
		//	panic(err)
		//}
	},
}

func init() {
	jump.Flags().StringVarP(&dir, "dir", "d", "", "source code root dir")
	over.Flags().StringVarP(&dir, "dir", "d", "", "source code root dir")
}
func main() {
	Root.AddCommand(jump)
	Root.AddCommand(over)
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
