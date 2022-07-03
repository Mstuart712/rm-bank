package cmd

import (
	"github.com/Mstuart712/rm-bank/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Execute(version string, exit func(int), args []string) {
	newRootCmd(version, exit).Execute(args)
}

type rootCmd struct {
	cmd     *cobra.Command
	project string
	exit    func(int)
}

func (c rootCmd) Execute(args []string) {
	c.cmd.SetArgs(args)
	if err := c.cmd.Execute(); err != nil {
		c.exit(1)
	}
}

func newRootCmd(version string, exit func(int)) *rootCmd {
	root := &rootCmd{
		exit: exit,
	}
	cmd := &cobra.Command{
		Use:          "rmb",
		Short:        "Rolemaster Bank for the Rolemaster Standard System",
		Version:      version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, f, err := setup(root.project)
			if err != nil {
				return err
			}
			defer db.Close()
			defer f.Close()

			p := tea.NewProgram(ui.Init(db, root.project))
			p.EnterAltScreen()
			defer p.ExitAltScreen()
			return p.Start()
		},
	}

	cmd.PersistentFlags().StringVarP(&root.project, "project", "p", "default", "Project name")

	cmd.AddCommand(
		newManCmd().cmd,
	)

	root.cmd = cmd
	return root
}
