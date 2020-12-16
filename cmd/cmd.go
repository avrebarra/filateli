package cmd

import (
	"path"

	"github.com/leaanthony/clir"
)

var cmd *clir.Cli

func customBanner(cli *clir.Cli) string {
	return ``
}

func Initialize() {
	cmd = clir.NewCli("filateli", "convert postman collection to markdown files", "v0")
	// Command.SetBannerFunction(customBanner)

	quiet := false
	cmd.BoolFlag("quiet", "perform quiet operation", &quiet)

	// default cli action on no params / flag
	cmd.Action(func() (err error) {
		cmd.PrintHelp()
		return
	})

	cmdBuild := cmd.NewSubCommand("build", "build postman collection file")
	{
		src := "."
		out := "."
		cmdBuild.StringFlag("src", "source folder (default: current dir)", &src)
		cmdBuild.StringFlag("out", "output folder (default: current dir)", &out)
		cmdBuild.Action(func() (err error) {
			srcclean := path.Clean(src)
			outclean := path.Clean(out)

			subcmd := NewCommandBuild(ConfigCommandBuild{
				Quiet:      quiet,
				SourcePath: srcclean,
				OutputPath: outclean,
			})

			return subcmd.Run()
		})
	}
}

func Run() error {
	return cmd.Run()
}
