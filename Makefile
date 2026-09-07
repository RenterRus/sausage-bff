GQL_VERSION := v0.17.95  

build:
	@go build cmd/main.go


init graphql:
	go get github.com/99designs/gqlgen/codegen/config@${GQL_VERSION} \
		github.com/99designs/gqlgen/internal/imports@${GQL_VERSION} \
		github.com/99designs/gqlgen/api@${GQL_VERSION} \
		github.com/99designs/gqlgen@${GQL_VERSION} \
	&& go run github.com/99designs/gqlgen@${GQL_VERSION} init 

gen graphql:
	go get github.com/99designs/gqlgen/codegen/config@${GQL_VERSION} \
		github.com/99designs/gqlgen/internal/imports@${GQL_VERSION} \
		github.com/99designs/gqlgen/api@${GQL_VERSION} \
		github.com/99designs/gqlgen@${GQL_VERSION} \
	&& go run github.com/99designs/gqlgen@${GQL_VERSION} generate 


update:
	@git pull && make build
