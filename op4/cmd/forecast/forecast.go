package forecast

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"weather/cmd/utils"

	"github.com/spf13/cobra"

	owm "github.com/briandowns/openweathermap"
)

var (
	coords    string
	city      string
	runCity   bool
	runCoords bool
)

var apiKey = os.Getenv("OWM_API_KEY")

const forecastTemplate = `Weather Forecast for {{.City.Name}}:
{{range .List}}Date & Time: {{.DtTxt}}
Conditions:  {{range .Weather}}{{.Description}}{{end}}
Temp:        {{.Main.Temp}}°C
High:        {{.Main.TempMax}}°C
Low:         {{.Main.TempMin}}°C
Wind speed:  {{.Wind.Speed}}m/s
Wind Dir:    {{.Wind.Deg}}°

{{end}}
`

func getCity(w *owm.ForecastWeatherData) {
	w.DailyByName(city, 40)

	forcast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	if forcast.City.Name == "" {
		fmt.Println("Invalid arguments, please double check them")
		fmt.Println("Have you entered coordinates to the city flag? Or mistyped the city name?")
		return
	}

	tmpl, err := template.New("forecast").Parse(forecastTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from forecast request")
		return
	}

	tmpl.Execute(os.Stdout, forcast)

}

func getCoords(w *owm.ForecastWeatherData) {
	lat, long, err := utils.ProcessCoords(coords)
	if err != nil {
		// error handled in function
		return
	}

	// Run query
	w.DailyByCoordinates(&owm.Coordinates{
		Longitude: lat,
		Latitude:  long,
	}, 40)

	forcast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	tmpl, err := template.New("forecast").Parse(forecastTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from forecast request")
		return
	}

	tmpl.Execute(os.Stdout, forcast)
}

// ForecastCmd represents the forecast command
var ForecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get a weather forcast for a city or location",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check API key is valid
		if apiKey == "" {
			fmt.Println("Cannot get API Key from env, check the OWM_API_KEY env variable is corretly set")
			return
		}

		err := owm.ValidAPIKey(apiKey)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Set default values
		runCoords, runCity, err = utils.CheckFlags(coords, city)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create weather forcast object
		w, err := owm.NewForecast("5", "C", "en", apiKey)
		if err != nil {
			fmt.Println("Bad API Key")
			log.Fatalln(err)
			return
		}

		if runCoords {
			// If coordinates were specified
			getCoords(w)
		} else if runCity {
			// If a city was specified
			getCity(w)
		}
	},
}

func init() {
	// Set flag options
	// Note if a flag option is set, the value in the first argument is filled out
	ForecastCmd.Flags().StringVarP(&coords, "coords", "l", "", "The coordinates in lat,long format to get the weather forcast for")
	ForecastCmd.Flags().StringVarP(&city, "city", "c", "", "The city to get the weather forcast for")
}
