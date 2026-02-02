package dic

import (
	"fmt"

	"github.com/urfave/cli"
)

// RunCLI runs as cli
func RunCLI(ctx *cli.Context) error {
	opts := Options{
		Reverse:     ctx.Bool("reverse"),
		FollowCNAME: ctx.Bool("follow-cname"),
		CnameMax:    ctx.Int("cname-max"),
	}
	args, opts, configPath, showHelp, showVersion := applyFallbackArgs(ctx.Args(), opts, ctx.String("config"))
	if showHelp {
		cli.ShowAppHelp(ctx)
		return nil
	}
	if showVersion {
		fmt.Fprintln(ctx.App.Writer, ctx.App.Version)
		return nil
	}
	opts = normalizeOptions(opts)

	setResolv()
	setConfigPath(configPath)
	c := newConfig()
	c.setSection(ctx.String("name"))

	d := New(args, opts)
	l := newLines(d)
	l.output(d)
	return nil
}
