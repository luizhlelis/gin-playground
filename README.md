# Gin playground

## Running the project

In the [src](./src) folder, type this command to run in `debug` mode:

```bash
    go run main.go
```

below an example of `curl`:

```bash
curl --request POST \
  --url http://localhost:3000/v1/workedHours \
  --header 'Content-Type: application/json' \
  --header 'accept: */*' \
  --data '{
	"TemperatureCelsius": 10,
	"summary": "Worked Hours Test"
}'
```
