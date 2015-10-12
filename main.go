package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/containerops/generator/cmd"
	"github.com/containerops/generator/modules"
	"github.com/containerops/wrench/setting"
)

func main() {

	if err := setting.SetConfig("conf/containerops.conf"); err != nil {
		fmt.Printf("Read config error: %v", err.Error())
	}

	if err := modules.LoadBuildList("/conf/pool.json"); err != nil {
		log.Fatal("load config file /conf/pool.json error")
	}

	app := cli.NewApp()

	app.Name = setting.AppName
	app.Usage = setting.Usage
	app.Version = setting.Version
	app.Author = setting.Author
	app.Email = setting.Email

	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
