

marshall-out.xml: marshall
	./marshall | xmllint --format - > marshall-out.xml

marshall: marshall.go
	go build -linkshared marshall.go
