UNAME := $(shell uname)

ifeq ($(UNAME),Linux)
	SED_CMD=sed -i -e
endif

ifeq ($(UNAME),Darwin)
	SED_CMD=sed -i '' -e
endif

.PHONY: test
test:
	go test -v ./...

.PHONY: generate-fakes
generate-fakes:
	go generate ./...
	counterfeiter -o fakes/fake_elasticacheiface_client.go \
		vendor/github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface/interface.go ElastiCacheAPI
	counterfeiter -o fakes/fake_cloudfront_client.go \
		vendor/github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface/interface.go CloudFrontAPI

	$(SED_CMD) "s|github.com/alphagov/paas-cf/tools/metrics/vendor/||" fakes/fake_*.go
