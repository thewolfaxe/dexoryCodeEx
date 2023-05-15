# Option 4 - Weather App
## About
This is a CLI weather app capable of getting a 5 day forecast in 3 hour increments and the current weather for any given GPS coordinates or city name. 

It uses a Golang wrapper for the [OpenWeatherMap API](https://openweathermap.org/api)

It also makes use of the [Cobra](https://cobra.dev/) library for building good looking CLI programs. 

## Building the code
Ensure that your Go environment is setup correctly. See official Go [documentation](https://go.dev/doc/install)

Build the code with:
`go build`

This should download required dependencies and build an executable called `weather`

## Usage
Ensure that the env variable is set in the terminal you are running the code in with:
`export OWM_API_KEY=**key**`
See the email for the API key

Run the CLI program with `./weather`

The main commands available are `forecast` and `current`.

Use `forecast` to get the 5 day estimate in 3 hour increments. 

Use `current` to get the current weather.

Both commands support `-c` or `--city` and `-l` or `--coords`  flags. 
Specify latitude and longitude in a `lat,long` format (note the lack of any spaces).

Only 1 flag can be used at a time

## Expected outputs
### Forecast
The `forecast` subcommand should print out blocks of text describing the temperature (avg, high and low for the area), the weather conditions in a brief description, the wind speed and direction for each location. The blocks will be split up in 3 hour increments over 5 days.
### Current
The `current` subcommand should print out the current weather conditions, temperature and wind speed and direction for the given location. 