package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/labstack/gommon/color"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	originTuLingURL = "http://www.tuling123.com/openapi/api?key=%s&info=%s"
	tlkey           = "a5052a22b8232be1e387ff153e823975"
)

// command list which supported
var cmds = []cli.Command{
	{
		Name:  "listen",
		Usage: "Listen to you",
		Action: func(ctx *cli.Context) {
			if len(ctx.Args()) == 0 {
				fmt.Println("what did you say?")
				return
			}

			tuLingURL := fmt.Sprintf(originTuLingURL, tlkey, url.QueryEscape(ctx.Args()[0]))
			res, err := http.Get(tuLingURL)
			if err != nil {
				log.Println(err)
				return
			}

			defer res.Body.Close()

			reply := new(tlReply)
			decoder := json.NewDecoder(res.Body)
			decoder.Decode(reply)

			wl := []string{"<cd.url=互动百科@", "", "&prd=button_doc_jinru>", "", "<br>", "\n"}
			srp := strings.NewReplacer(wl...)
			ret := srp.Replace(reply.Text)

			log.Printf("%s: %s", color.Green("Alice say", color.B), color.Underline(ret))
		},
	},
}

type tlReply struct {
	code int
	Text string `json:"text"`
}

func main() {
	app := cli.NewApp()
	app.Name = "alice"
	app.Usage = "Command line tool to chat with alice"
	app.Version = "0.0.1"
	app.Author = "Nixon"
	app.Email = "xiaoyao.ning@gmail.com"
	app.Commands = cmds
	app.Run(os.Args)
}
