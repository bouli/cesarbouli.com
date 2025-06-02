package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type templateData struct {
	Domain            string
	Url_path          string
	Artist_name       string
	Meta_text         string
	Meta_title        string
	Meta_url          string
	Cta               string
	Cta_url           string
	Cta_exists        bool
	Soundcloud_link   string
	Soundcloud_exists bool
	Facebook_link     string
	Facebook_exists   bool
	Instagram_link    string
	Instagram_exists  bool
	Youtube_link      string
	Youtube_exists    bool
	Spotify_link      string
	Spotify_exists    bool
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		fmt.Println()
		main := template.Must(template.ParseFiles("templates/index.html"))
		main.Execute(w, getTemplateData(r))
	})
	http.HandleFunc("/como-se-fosse-a-ultima-vez", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQZYDKcQiEU", http.StatusFound)
	})

	http.HandleFunc("/diana", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://open.spotify.com/track/7GPSPivgaqYPQ6FstxXi9C", http.StatusFound)
	})

	http.HandleFunc("/happy-end", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://open.spotify.com/album/427He8WE8uuO4ECdchYhh8", http.StatusFound)
	})

	http.HandleFunc("/cesar-bouli", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://distrokid.com/hyperfollow/cesarbouli/cesar-bouli", http.StatusFound)
	})

	http.HandleFunc("/hey-irmao", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://distrokid.com/hyperfollow/cesarbouli/hey-irmo-2", http.StatusFound)
	})

	http.HandleFunc("/vale-a-pena", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://distrokid.com/hyperfollow/cesarbouli/vale-a-pena-4", http.StatusFound)
	})

	http.HandleFunc("/fios-naturais", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://distrokid.com/hyperfollow/cesarbouli/fios-naturais", http.StatusFound)
	})

	http.ListenAndServe(":80", nil)
}

func getTemplateData(r *http.Request) templateData {

	if r.Host == "byeceebee.com" {
		return templateData{
			Domain:            r.Host,
			Url_path:          r.URL.Path,
			Artist_name:       "ByeCeeBee",
			Meta_text:         "Listen to ByeCeeBee",
			Meta_title:        "Listen to ByeCeeBee",
			Meta_url:          getHttpString(r) + "://" + r.Host,
			Cta:               "Listen to ByeCeeBee",
			Cta_url:           "diana",
			Cta_exists:        true,
			Soundcloud_link:   "",
			Soundcloud_exists: false,
			Facebook_link:     "",
			Facebook_exists:   false,
			Instagram_link:    "https://www.instagram.com/byeceebee",
			Instagram_exists:  true,
			Youtube_link:      "https://www.youtube.com/@ByeCeeBee",
			Youtube_exists:    true,
			Spotify_link:      "https://open.spotify.com/artist/6sJjhXfbYaWZXeVL81BZiz",
			Spotify_exists:    true,
		}
	} else if r.Host == "salviasupernova.com.br" {
		return templateData{
			Domain:            r.Host,
			Url_path:          r.URL.Path,
			Artist_name:       "Salvia Supernova",
			Meta_text:         "Salvia Supernova",
			Meta_title:        "Salvia Supernova",
			Meta_url:          getHttpString(r) + "://" + r.Host,
			Cta:               "",
			Cta_url:           "",
			Cta_exists:        false,
			Soundcloud_link:   "",
			Soundcloud_exists: false,
			Facebook_link:     "",
			Facebook_exists:   false,
			Instagram_link:    "https://www.instagram.com/salviasupernova",
			Instagram_exists:  true,
			Youtube_link:      "https://www.youtube.com/@salviasupernova",
			Youtube_exists:    true,
			Spotify_link:      "https://open.spotify.com/artist/7i3TMpT0F2YqHDS1MEstLO",
			Spotify_exists:    true,
		}
	} else {
		return templateData{
			Domain:            "cesarbouli.com",
			Url_path:          r.URL.Path,
			Artist_name:       "Cesar Bouli",
			Meta_text:         "Cesar Bouli",
			Meta_title:        "Cesar Bouli",
			Meta_url:          getHttpString(r) + "://" + r.Host,
			Cta:               "Listen to Fios Naturais",
			Cta_url:           "fios-naturais",
			Cta_exists:        true,
			Soundcloud_link:   "https://soundcloud.com/cesar-bouli",
			Soundcloud_exists: true,
			Facebook_link:     "https://www.facebook.com/cesarbouli",
			Facebook_exists:   true,
			Instagram_link:    "https://www.instagram.com/cesarbouli",
			Instagram_exists:  true,
			Youtube_link:      "https://www.youtube.com/cesarbouli",
			Youtube_exists:    true,
			Spotify_link:      "https://open.spotify.com/artist/16K0CW2ObGqglDFtWjxbMb",
			Spotify_exists:    true,
		}
	}
}

func getHttpString(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	} else {
		return "http"
	}
}
