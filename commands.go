package dic

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

const (
	col0 = "DEFAULT SECTION"
	col1 = "SECTION NAME"
	col2 = "SERVER"
	col3 = "QUERY TYPES"
)

// Flags variable
var Flags = []cli.Flag{
	flagName,
	flagReverse,
	flagFollowCNAME,
	flagCnameMax,
	flagConfig,
}

// Commands variable
var Commands = []cli.Command{
	commandList,
	commandEdit,
	commandSet,
}

var usages = map[string]string{
	"list": "| dic l",
	"edit": "| dic e",
	"set":  "<section name> | dic s <section name>",
}

var flagName = cli.StringFlag{
	Name:  "name, n",
	Usage: "Pass a `<section name>` for temporary use",
}

var flagReverse = cli.BoolFlag{
	Name:  "reverse, r",
	Usage: "Reverse lookup for IP arguments (PTR); domain args remain normal",
}

var flagFollowCNAME = cli.BoolFlag{
	Name:  "follow-cname, f",
	Usage: "Follow CNAMEs and query A/AAAA for the target name",
}

var flagCnameMax = cli.IntFlag{
	Name:  "cname-max, m",
	Usage: "Maximum CNAME follow depth (default 5)",
	Value: defaultCnameMax,
}

var flagConfig = cli.StringFlag{
	Name:  "config, c",
	Usage: "Path to config file (default: ./config.toml or GOPATH path)",
}

var commandList = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "Show default session and config description",
	Action:  doList,
}

var commandSet = cli.Command{
	Name:    "set",
	Aliases: []string{"s"},
	Usage:   "Set a section to use by default",
	Action:  doSet,
}

var commandEdit = cli.Command{
	Name:    "edit",
	Aliases: []string{"e"},
	Usage:   "Open and edit config file",
	Action:  doEdit,
}

func init() {
	usages := setUsages()
	cli.CommandHelpTemplate = `NAME:
    {{.Name}} - {{.Usage}}
USAGE:
    dic {{.Name}} ` + usages + `{{ "\n\n" }}`
}

func setUsages() string {
	s := "{{if false}}"
	for _, command := range append(Commands) {
		s = s + fmt.Sprintf("{{else if (eq .Name %q)}}%s", command.Name, usages[command.Name])
	}
	return s + "{{end}}"
}

func doEdit(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "edit")
		os.Exit(exitErr)
	}

	open.Run(getAppPath(configFile))
	return nil
}

func doList(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "list")
		os.Exit(exitErr)
	}

	conf := newConfig()

	ds := getDefaultSection()
	fmt.Printf("%s\n %s\n\n", col0, ds)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 10, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", col1, col2, col3)

	for _, s := range conf.Sec {
		for i, a := range s.Args {
			if i == 0 {
				if len(a.Server) == 0 {
					a.Server = unspecified
				}

				if strings.EqualFold(ds, s.Name) {
					fmt.Fprintf(w, "*%s\t%s\t%s\n", s.Name, a.Server, a.Qtypes)
				} else {
					fmt.Fprintf(w, " %s\t%s\t%s\n", s.Name, a.Server, a.Qtypes)
				}
			} else {
				fmt.Fprintf(w, "\t%s\t%s\n", a.Server, a.Qtypes)
			}
		}
		fmt.Fprintln(w, "\t\t")
	}
	w.Flush()

	return nil
}

func doSet(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "set")
		os.Exit(exitErr)
	}

	s := removeNewline(c.Args().First())

	f, err := os.OpenFile(getAppPath(sectionFile), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		err := fmt.Errorf("%s", err)
		setError(err)
	}
	defer f.Close()

	fmt.Fprintln(f, s)
	fmt.Printf("Changed the default section to %q\n", s)
	return nil
}
