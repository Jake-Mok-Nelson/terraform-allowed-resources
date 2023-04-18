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
	RulesFile         string `help:"Path to the file containing your rules." type:"existingfile" default:"rules.json" name:"rules" env:"TF_ALLOWED_RULES_FILE"`
	PlanFile          string `help:"Path to the Terraform plan file." type:"existingfile" default:"plan.tfplan" name:"plan" env:"TF_ALLOWED_PLAN_FILE"`
	DetailedExitCodes bool   `help:"Return detailed exit codes." default:"false" env:"TF_ALLOWED_DETAILED_EXIT_CODES"`
	Debug             bool   `help:"Enable debug mode." default:"false" env:"TF_ALLOWED_DEBUG"`
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
