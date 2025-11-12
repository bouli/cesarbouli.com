package models

import (
	"boulihub/utils"
	"net/http"
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

func GetTemplateData(r *http.Request, forceWebSite string) templateData {

	if r.Host == "byeceebee.com" || forceWebSite == "byeceebee.com" {
		return templateData{
			Domain:            r.Host,
			Url_path:          r.URL.Path,
			Artist_name:       "ByeCeeBee",
			Meta_text:         "Listen to ByeCeeBee",
			Meta_title:        "Diana",
			Meta_url:          utils.GetHttpProtocolString(r) + "://" + r.Host,
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
	} else if r.Host == "salviasupernova.com.br" || forceWebSite == "salviasupernova.com.br" {
		return templateData{
			Domain:            r.Host,
			Url_path:          r.URL.Path,
			Artist_name:       "Salvia Supernova",
			Meta_text:         "Salvia Supernova",
			Meta_title:        "α",
			Meta_url:          utils.GetHttpProtocolString(r) + "://" + r.Host,
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
	} else if r.Host == "paralelossa.com.br" || forceWebSite == "paralelossa.com.br" {
		return templateData{
			Domain:            r.Host,
			Url_path:          r.URL.Path,
			Artist_name:       "Paralelos s/A",
			Meta_text:         "Escute Paralelos s/A.",
			Meta_title:        "Universo Surdo e Mudo",
			Meta_url:          utils.GetHttpProtocolString(r) + "://" + r.Host,
			Cta:               "Escute Paralelos s/A",
			Cta_url:           "universo-surdo-e-mudo",
			Cta_exists:        true,
			Soundcloud_link:   "",
			Soundcloud_exists: false,
			Facebook_link:     "",
			Facebook_exists:   false,
			Instagram_link:    "",
			Instagram_exists:  false,
			Youtube_link:      "https://www.youtube.com/paralelossa",
			Youtube_exists:    true,
			Spotify_link:      "https://open.spotify.com/artist/50R6Tiq0s02DaYQZcSr58Q",
			Spotify_exists:    true,
		}
	} else {
		return templateData{
			Domain:            "cesarbouli.com",
			Url_path:          r.URL.Path,
			Artist_name:       "Cesar Bouli",
			Meta_text:         "Cesar Bouli",
			Meta_title:        "Fios Naturais",
			Meta_url:          utils.GetHttpProtocolString(r) + "://" + r.Host,
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
