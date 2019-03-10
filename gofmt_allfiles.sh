while true
do
	gofmt -w *.go;
	echo "All go files have been formated.";
	sleep 3;
done
