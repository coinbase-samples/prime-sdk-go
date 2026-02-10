.PHONY: fetch-spec

# Fetch the Prime API OpenAPI specification
fetch-spec:
	@mkdir -p apiSpec
	curl -o apiSpec/prime-public-api-spec.yaml https://api.prime.coinbase.com/v1/openapi.yaml
