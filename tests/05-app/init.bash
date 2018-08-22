if [ -z "$appName" ]; then
    export name=$(date | md5sum | cut -b 1-19)
    export appName="test-app-$name"

    rm -rf app.json
    cp app.info app.json
    sed -i  "s/@APPNAME@/${appName}/" app.json
fi
