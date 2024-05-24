from flask import Flask, redirect, render_template

app = Flask(__name__)


@app.route("/")
def index():
    artist_name = "Cesar Bouli"
    meta_image = "http://cesarbouli.com/images/photo.jpg"
    meta_text = f"Ouça {artist_name}."
    meta_title = f"Ouça {artist_name}."
    cta = 'Ouça o Disco "Cesar Bouli"'
    cta_url = "cesarbouli"

    return render_template(
        "index.jinja",
        artist_name=artist_name,
        meta_image=meta_image,
        meta_text=meta_text,
        meta_title=meta_title,
        cta=cta,
        cta_url=cta_url,
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
