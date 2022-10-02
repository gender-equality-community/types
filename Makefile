default: README.md

README.md: *.go .doc/*
	cat .doc/* > $@
	goreadme --types --constants --functions --methods github.com/gender-equality-community/types >> $@

%.pb.go: proto/%.proto
	protoc -I proto/ $< --go_out=module=github.com/gender-equality-community/types:.
