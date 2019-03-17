package dic

import "github.com/urfave/cli"

// RunCLI runs as cli
func RunCLI(cli *cli.Context) error {
	setResolv()
	c := newConfig()
	c.setSection(cli.String("name"))

	d := New(cli.Args())
	l := newLines(d)
	l.output(d)
	return nil
}
