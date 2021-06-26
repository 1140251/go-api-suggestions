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

Confidence score is a value between 0 and 1. Only cities in GB are allowed to be returned.


## How to run
Install dependencies:
- `go mod download`

From the root of the app:
- `go run cmd/main.go`

or run the binary:
- `./build/main`

Run test:
- `go test ./...`

## Solution
Since the endpoint only returns suggestions of city within Great Britain, and because the file is loaded in memory,
I decided to write a small python script to skim cities which are not in GB.
This will make reading from file and filtering/sorting faster in our API code.

The first criteria for filtering is if there is a partial/total string match between city names.
If there is a confidence score is calculated and the slice of suggestions sorted in descending order.


### How the confidence score is scored
This requirement was arbitrary.
I decided to score the confidence level based on the distance in km between the coordinates input by the user and the actual coordinates of the city 
only if there is at least a partial match of strings.

I used a formula called `Haversine formula` which I found on Google and translated the JavaScript formula into Go code. Can't take all credit for that math :)

The farther the user coordinates the lower the score percentage will be.
