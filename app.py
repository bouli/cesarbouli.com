from flask import Flask, redirect, render_template

app = Flask(__name__)


@app.route("/")
def index():
    domain = "cesarbouli.com"
    artist_name = "Cesar Bouli"
    meta_text = f"Ouça {artist_name}."
    meta_title = f"Ouça {artist_name}."
    meta_url = "http://cesarbouli.com"
    cta = 'Ouça o Disco "Cesar Bouli"'
    cta_url = "cesarbouli"
    soundcloud_link = "https://soundcloud.com/cesar-bouli"
    facebook_link = "https://www.facebook.com/cesarbouli"
    instagram_link = "https://www.instagram.com/cesarbouli"
    youtube_link = "https://www.youtube.com/cesarbouli"
    spotify_link = "https://open.spotify.com/artist/16K0CW2ObGqglDFtWjxbMb"


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
