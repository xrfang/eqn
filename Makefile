HASH=$(shell git log -n1 --pretty=format:%h)
REVS=$(shell git log --oneline|wc -l)
develop: --setver --geneh
	rm -f esp*
	go build -ldflags="-s -w" .
windows: export GOOS=windows
windows: export GOARCH=amd64
windows: develop
release:
	make
	make windows
	upx --best --lzma eqn*
clean:
	rm -fr eqn* version.go
--geneh: #generate error handler
	@for tpl in `find . -type f |grep errors.tpl`; do \
	    target=`echo $$tpl|sed 's/\.tpl/\.go/'`; \
	    pkg=`basename $$(dirname $$tpl)`; \
		sed "s/package main/package $$pkg/" src/errors.go > $$target; \
    done
--setver:
	cp verinfo.tpl version.go
	sed -i 's/{_G_HASH}/$(HASH)/' version.go
	sed -i 's/{_G_REVS}/$(REVS)/' version.go
