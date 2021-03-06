package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	gcli "gopkg.in/urfave/cli.v1"
)

func init() {
	cmd := gcli.Command{
		Name:      "broadcastTransaction",
		Usage:     "Broadcast a raw transaction to the network.",
		ArgsUsage: "[transaction]",
		Action: func(c *gcli.Context) error {
			rawtx := c.Args().First()
			if rawtx == "" {
				return errors.New("raw transaction is empty")
			}

			v, err := broadcastTx(rawtx)
			if err != nil {
				return err
			}
			fmt.Println(v)
			return nil
		},
	}
	Commands = append(Commands, cmd)
}

func broadcastTx(rawtx string) (string, error) {
	var tx = struct {
		Rawtx string `json:"rawtx"`
	}{
		rawtx,
	}
	d, err := json.Marshal(tx)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/injectTransaction", nodeAddress)
	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return "", errConnectNodeFailed
	}
	defer rsp.Body.Close()
	v, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(v), "\""), nil
}
