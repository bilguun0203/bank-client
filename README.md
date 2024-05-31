# Bank Client

CLI client for Khan Bank.

## Usage

### Create new config file

```sh
go run . -config clients/config.json -create
```

Edit the resulting JSON file, or you can create it with all the necessary fields filled using the following command.

```sh
go run . \
    -config clients/config.json \
    -create \
    -username=<username> \
    -password=<password> \
    -deviceid=<uuid> \
    -useragent=<useragent>
```

### Login

Use the interactive login if your bank requires OTP verification.

```sh
go run . -login
```

### Download Transactions

Use the following command to get a list of transactions. To save the output as a JSON file, use the -save-path flag and specify the file name.

```sh
go run . -transactions -account=<account_number> -start-date=2024-04-01 -end-date=2024-05-30 -save-path=transactions.json
```
