for dm in `cat domains.txt`; do
    heroku domains:add $dm
    heroku domains:add "*.$dm"
done
