package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"boulihub/models"
	"boulihub/routes"
	"boulihub/utils"
)

func main() {

	var forceWebSite string
	flag.StringVar(&forceWebSite, "force-website", "", "force website")
	flag.Parse()

	fmt.Println(os.Args)
	//TODO: Make it handled
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fs_assets := http.FileServer(http.Dir("cesarcardoso.cc/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs_assets))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		destiny := routes.GetDestiny(r, forceWebSite)
		fmt.Println(destiny.DestinyType)
		if destiny.DestinyType == "redirect" {
			http.Redirect(w, r, destiny.DestinyUrl, http.StatusFound)
			return
		} else if destiny.DestinyType == "template" {
			main := template.Must(template.ParseFiles("templates/index.html"))
			main.Execute(w, models.GetTemplateData(r, destiny.DestinyUrl))
			return
		} else if destiny.DestinyType == "static" {
			fmt.Println("im static")
			http.ServeFile(w, r, destiny.DestinyUrl)
			return
		}

		http.Redirect(w, r, utils.GetHttpProtocolString(r)+"://"+r.Host, http.StatusNotFound)

	})

	http.ListenAndServe(":"+utils.GetServerPort(), nil)

}
