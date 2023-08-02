# Prime Go SDK README

## Overview

The *Prime Go SDK* is a sample libary that demonstrates the structure of a [Coinbase Prime](https://prime.coinbase.com/) driver for
the [REST APIs](https://docs.cloud.coinbase.com/prime/reference).

## License

The *Prime Go SDK* sample library is free and open source and released under the [Apache License, Version 2.0](LICENSE).

The application and code are only available for demonstration purposes.

## Usage

To use the *Prime Go SDK*, initialize the [Credentials](credentials.go) struct and create a new client. The Credentials struct is JSON
enabled. Ensure that Prime API credentials are stored in a secure manner.

```
credentials := &prime.Credentials{}
if err := json.Unmarshal([]byte(os.Getenv("PRIME_CREDENTIALS")), credentials); err != nil {
    return nil, fmt.Errorf("unable to deserialize prime credentials JSON: %w", err)
}

client := prime.NewClient(credentials, http.Client{})
```

Once the client is inialized, make the desired call. For example, to [list portfolios](https://github.com/coinbase-samples/prime-sdk-go/blob/main/list_portfolios.go), 
pass in the request object, check for an error, and if nil, process the response.


```
response, err := client.ListPortfolios(ctx, &prime.ListPortfoliosRequest{})
```

## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build *.go
```
