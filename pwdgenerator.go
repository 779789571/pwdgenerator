package main

import (
	"./lib_core"
	"pwdgenerator/gologger"
)

const version = "0.1"
const banner = `

                 .___                                         __                
________  _  ____| _/ ____   ____   ____   ________________ _/  |_  ___________ 
\____ \ \/ \/ / __ | / ___\_/ __ \ /    \_/ __ \_  __ \__  \\   __\/  _ \_  __ \
|  |_> >     / /_/ |/ /_/  >  ___/|   |  \  ___/|  | \// __ \|  | (  <_> )  | \/
|   __/ \/\_/\____ |\___  / \___  >___|  /\___  >__|  (____  /__|  \____/|__|   
|__|              \/_____/      \/     \/     \/           \/                   

design by jinx0v0
`

func showBanner() {
	gologger.Printf(banner)
	gologger.Printf("current version: %s\n\r", version)
}

func main() {
	showBanner()
	options := lib_core.ParseOptions()
	lib_core.Start(options)
	//result :=lib_core.Uniq(options)

}
