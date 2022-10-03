package main

import (
	"fmt"
	"log"
	"os"

	"github.com/i0Ek3/color"
	"github.com/i0Ek3/ognoc"
	"github.com/urfave/cli/v2"
)

var (
	newCmd          cli.Command
	nonAlphabetFlag cli.BoolFlag
	nonNumberFlag   cli.BoolFlag
	nonSpecharFlag  cli.BoolFlag
	lengthFlag      cli.IntFlag
)

func new(ctx *cli.Context) error {
	var options []ognoc.Option

	if !ctx.Bool("non-alphabet") {
		options = append(options, ognoc.WithAlphabet())
	}
	if !ctx.Bool("non-number") {
		options = append(options, ognoc.WithNumber())
	}
	if !ctx.Bool("non-spechar") {
		options = append(options, ognoc.WithSpechar())
	}

	fmt.Println(color.Cyan(ognoc.NewPassword(ctx.Int("length"), options...)))

	return nil
}

func init() {
	lengthFlag = cli.IntFlag{
		Name:    "length",
		Aliases: []string{"l"},
		Value:   ognoc.PwdLen,
		Usage:   "length of cipher",
	}

	nonAlphabetFlag = cli.BoolFlag{
		Name:    "non-alphabet",
		Aliases: []string{"no-alpa"},
		Value:   false,
		Usage:   "no alphabet contains in cipher",
	}

	nonNumberFlag = cli.BoolFlag{
		Name:    "non-number",
		Aliases: []string{"no-num"},
		Value:   false,
		Usage:   "no number contains in cipher",
	}

	nonSpecharFlag = cli.BoolFlag{
		Name:    "non-spechar",
		Aliases: []string{"no-spec"},
		Value:   false,
		Usage:   "no special char contains in cipher",
	}

	newCmd = cli.Command{
		Name:   "new",
		Usage:  "create a new cipher",
		Action: new,
		Flags: []cli.Flag{
			&lengthFlag,
			&nonAlphabetFlag,
			&nonNumberFlag,
			&nonSpecharFlag,
		},
	}
}

func main() {
	app := &cli.App{
		Name:     "ognoc",
		Usage:    "use ognoc to create a complicated cipher",
		Commands: []*cli.Command{&newCmd},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
