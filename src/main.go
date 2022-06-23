package main

import (
	"fmt"

	"github.com/Jake-Mok-Nelson/terraform-allowed-resources/src/placeholder"
	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type EvalCmd struct {
	RulesFile string `help:"Path to the file containing your rules." arg type:"existingfile" default:"rules.json" name:"rules"`
	PlanFile  string `help:"Path to the Terraform plan file." arg type:"existingfile" default:"plan.tfplan" name:"plan"`
	Debug     bool   `help:"Enable debug mode." default:"false"`
}

func (r *EvalCmd) Run(ctx *Context) error {
	fmt.Println("eval", r.RulesFile, r.PlanFile)
	placeholder.PlaceHolder()
	return nil
}

var cli struct {
	Eval EvalCmd `cmd help:"Evaluate the Terraform Plan against provided rules."`
}

func main() {

	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{})
	ctx.FatalIfErrorf(err)
}
