get:
	rm -rf go.sum
	go get ./...


VERSION=0.1.7
tag:
	git tag -a v$(VERSION) -m "Release version $(VERSION)"
	git push origin v$(VERSION)

push:
	git push origin tag v$(VERSION)