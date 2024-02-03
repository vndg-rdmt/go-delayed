.PHONY: test clean

test:
	@(go test ./ -v -coverprofile cover.out.tmp  && \
	cat cover.out.tmp | grep -v "_mock.go" > cover.out && \
	rm cover.out.tmp && \
	go tool cover -func cover.out) || true

clean:
	@rm -r ./*.out