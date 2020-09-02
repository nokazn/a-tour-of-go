go run ./images/images.go | sed -e 's/IMAGE:\(.*\)/<img src="data:image\/png;base64,\1">/g' > ./images/index.html
