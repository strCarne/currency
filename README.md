
# Currency

Service 'currency' collects all conversion rates relative to BYN (Belorussian Rubles) from NBRB's (National Bank Of Republic of Belarus) API once a day.


## Run Locally

Clone the project
```bash
  git clone https://github.com/strCarne/currency
```

Go to the project directory
```bash
  cd currency
```

Run compose.yml
```bash
  docker compose up -d
```


## Documentation

To generate swagger:
```bash
go install github.com/swaggo/swag/cmd/swag@latest 
go get -u github.com/swaggo/echo-swagger

swag init --parseDependency --generalInfo cmd/currency/main.go --output api
```