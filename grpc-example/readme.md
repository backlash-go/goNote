```shell script

xixianbin@xixianbindeMac-mini helloworld % pwd
/Users/xixianbin/go/src/gobackend/goNote/grpc-example/helloworld

xixianbin@xixianbindeMac-mini helloworld % protoc -I. --go_out=plugins=grpc:. helloworld.proto
xixianbin@xixianbindeMac-mini helloworld % ls
helloworld.pb.go        helloworld.proto


```