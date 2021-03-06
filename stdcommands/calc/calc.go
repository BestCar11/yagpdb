package calc

import (
	"fmt"
	"github.com/alfredxing/calc/compute"
	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/commands"
	"sync"
)

var (
	// calc/compute isnt threadsafe :'(
	computeLock sync.Mutex
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryTool,
	Name:         "Calc",
	Aliases:      []string{"c", "calculate"},
	Description:  "Calculator 2+2=5",
	RunInDM:      true,
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		&dcmd.ArgDef{Name: "Expression", Type: dcmd.String},
	},

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		computeLock.Lock()
		defer computeLock.Unlock()
		result, err := compute.Evaluate(data.Args[0].Str())
		if err != nil {
			return err, err
		}

		return fmt.Sprintf("Result: `%f`", result), nil
	},
}
