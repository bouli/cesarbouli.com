from flask import Flask, redirect, render_template, request

app = Flask(__name__)


@app.route("/")
def index():
    if request.headers["Host"] == "cesarbouli.com":
        domain = "cesarbouli.com"
        artist_name = "Cesar Bouli"
        meta_text = f"Ouça {artist_name}."
        meta_title = f"Ouça {artist_name}."
        meta_url = f"http://{domain}"
        cta = 'Ouça o Disco "Cesar Bouli"'
        cta_url = "cesarbouli"
        soundcloud_link = "https://soundcloud.com/cesar-bouli"
        facebook_link = "https://www.facebook.com/cesarbouli"
        instagram_link = "https://instagram.com/salviasupernova"
        youtube_link = "https://www.youtube.com/cesarbouli"
        spotify_link = "https://open.spotify.com/artist/16K0CW2ObGqglDFtWjxbMb"
    elif request.headers["Host"] == "salviasupernova.com.br":
        domain = "salviasupernova.com.br"
        artist_name = "Salvia Supernova"
        meta_text = f"Ouça {artist_name}."
        meta_title = f"Ouça {artist_name}."
        meta_url = f"http://{domain}"
        cta = 'Ouça o Disco "α"'
        cta_url = "alpha"
        soundcloud_link = False
        facebook_link = False
        instagram_link = "https://www.instagram.com/salviasupernova"
        youtube_link = "https://www.youtube.com/@salviasupernova"
        spotify_link = "https://open.spotify.com/artist/7i3TMpT0F2YqHDS1MEstLO"
    else:
        domain = "byeceebee.com"
        artist_name = "ByeCeeBee"
        meta_text = f"Listen to {artist_name}."
        meta_title = f"Listen to {artist_name}."
        meta_url = f"http://{domain}"
        cta = 'Listen to "Diana"'
        cta_url = "diana"
        soundcloud_link = False
        facebook_link = False
        instagram_link = "https://www.instagram.com/byeceebee"
        youtube_link = "https://www.youtube.com/@ByeCeeBee"
        spotify_link = "https://open.spotify.com/artist/6sJjhXfbYaWZXeVL81BZiz"

    return render_template(
        "index.jinja",
        domain=domain,
        artist_name=artist_name,
        meta_text=meta_text,
        meta_title=meta_title,
        meta_url=meta_url,
        cta=cta,
        cta_url=cta_url,
        soundcloud_link=soundcloud_link,
        facebook_link=facebook_link,
        instagram_link=instagram_link,
        youtube_link=youtube_link,
        spotify_link=spotify_link,
    )


@app.route("/alpha")
def alpha():
    return redirect("https://distrokid.com/hyperfollow/salviasupernova/3j29b2Stjw3")


@app.route("/diana")
def diana():
    return redirect("https://open.spotify.com/track/7GPSPivgaqYPQ6FstxXi9C")


@app.route("/cesar-bouli")
def cesarbouli():
    return redirect("https://distrokid.com/hyperfollow/cesarbouli/cesar-bouli")


@app.route("/hey-irmao")
def heyirmao():
    return redirect("https://distrokid.com/hyperfollow/cesarbouli/hey-irmo-2")


@app.route("/como-se-fosse-a-ultima-vez")
def comosefosseaultimavez():
    return redirect("http://cesarbouli.com/como-se-fosse-a-ultima-vez")


@app.route("/vale-a-pena")
def valeapena():
    return redirect("https://distrokid.com/hyperfollow/cesarbouli/vale-a-pena-4")
