package cli

import (
	"encoding/json"
	"fmt"
	"reflect"

	gcli "gopkg.in/urfave/cli.v1"
)

func init() {
	cmd := gcli.Command{
		Name:      "version",
		Usage:     "List the current version of Skycoin components.",
		ArgsUsage: "[options]",
		Flags: []gcli.Flag{
			gcli.BoolFlag{
				Name:  "j,json",
				Usage: "Returns the results in JSON format.",
			},
		},
		Action: func(c *gcli.Context) error {
			var ver = struct {
				Skycoin string `json:"skycoin"`
				Cli     string `json:"cli"`
				RPC     string `json:"rpc"`
				Wallet  string `json:"wallet"`
			}{
				"0.1",
				"0.1",
				"0.1",
				"0.1",
			}

			jsonFmt := c.Bool("json")
			if jsonFmt {
				d, err := json.MarshalIndent(ver, "", "    ")
				if err != nil {
					return err
				}
				fmt.Println(string(d))
				return nil
			}

			v := reflect.ValueOf(ver)
			t := reflect.TypeOf(ver)
			for i := 0; i < v.NumField(); i++ {
				fmt.Printf("%s:%v\n", t.Field(i).Tag.Get("json"), v.Field(i).Interface())
			}

			return nil
		},
	}
	Commands = append(Commands, cmd)
}
