go run ./slices/slices.go | sed -e 's/IMAGE:\(.*\)/<img src="data:image\/png;base64,\1">/g' > ./slices/index.html
