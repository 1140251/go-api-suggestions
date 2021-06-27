## Task

Create an API endpoint `GET /suggestions?q=Wok&latitude=43.70011&longitude=-79.4163`
which takes 3 parameters (city name, latitude and longitude) and returns a list of suggestions.

Each suggestions is: 
```
{
    name: <city name>,
    latitude: <latitude>,
    longitude: <longitude>,
    score: <confidence score>
}
```

Confidence score is a value between 0 and 1. Only cities in GB allowed to be returned.

## How to run
Install dependencies:
- `go mod download`

From the root of the app:
- `go run cmd/main.go`

or run the binary:
- `./build/main`

Run test:
- `go test -v ./...`

I have included a Postman collection as JSON to run the query after the service has started.
Or simply e.g. `curl http://localhost:8080/api/v1/suggestions?q=London&latitude=51.507&longitude=0.127`

## Solution
Since the endpoint only returns suggestions of city within Great Britain, and because the file is loaded in memory,
I decided to write a small python script to skim cities which are not in GB.
This will make reading from file and filtering/sorting faster in our API code.

What the API is returned is based on 3 criteria:
- The location must be within GB country code. The python script handles this already
- There has to be a partial or total string match.
- A confidence score calculated based on distance and the slice of response sorted in descending order.


### How the confidence score is calculated:
This requirement is arbitrary.
I decided to score the confidence level based on the distance in km between the coordinates input by the user and the actual coordinates of the city 
only if there is at least a partial match of strings.

I used a formula called `Haversine formula` which I found on Google and translated from JavaScript into Go code. Can't take all credit for that math :)

For example if I input "London" and London's coordinates, the string match returns:
London and Londonderry. However, since London's coordinates are much close to London than Londonderry,
the confidence will be much higher for the former, and the return will be sorted accordingly.