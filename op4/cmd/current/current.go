package current

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"weather/cmd/utils"

	"github.com/spf13/cobra"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")

var (
	coords    string
	city      string
	runCity   bool
	runCoords bool
)

const currentTemplate = `Current weather for {{.Name}}:
    Conditions: {{range .Weather}} {{.Description}} {{end}}
    Now:         {{.Main.Temp}}°C
    High:        {{.Main.TempMax}}°C
    Low:         {{.Main.TempMin}}°C
    Wind speed:  {{.Wind.Speed}}m/s
    Wind Dir:    {{.Wind.Deg}}°
`

// Calls the weathermap current weather api and prints to terminal for the given GPS coordinates
func getCoords(w *owm.CurrentWeatherData) {
	lat, long, err := utils.ProcessCoords(coords)
	if err != nil {
		// error handled in function
		return
	}

	// API call
	w.CurrentByCoordinates(&owm.Coordinates{
		Longitude: long,
		Latitude:  lat,
	})

	// Parse output template
	tmpl, err := template.New("current").Parse(currentTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from request")
		return
	}

	// Parse API response and print to stdout
	err = tmpl.Execute(os.Stdout, w)
	if err != nil {
		fmt.Println("BAD JSON response")
	}
}

// Calls the weathermap current weather api and prints to terminal for a city
func getCity(w *owm.CurrentWeatherData) {
	// API call
	w.CurrentByName(city)

	// Check valid repsonse
	// Can be pretty sure if all these are true, an invalid argument has been passed to it
	if (w.Name == "") && (w.Main.Temp == 0.0) && (w.Main.TempMin == 0.0) && (w.Wind.Deg == 0) {
		fmt.Println("Invalid arguments, please double check them")
		fmt.Println("Have you entered coordinates to the city flag? Or mistyped the city name?")
		return
	}

	// Parse output template
	tmpl, err := template.New("current").Parse(currentTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from request")
		return
	}

	// Parse API response and print to terminal
	err = tmpl.Execute(os.Stdout, w)
	if err != nil {
		fmt.Println("BAD JSON response")
	}
}

// CurrentCmd represents the current command
var CurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current weather for a city or location",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check API key is valid
		if apiKey == "" {
			fmt.Println("Cannot get API Key from env, check the OWM_API_KEY env variable is corretly set")
			return
		}
		err := owm.ValidAPIKey(apiKey)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			return
		}

		// Check user provided flags
		runCoords, runCity, err = utils.CheckFlags(coords, city)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new current weather API object
		w, err := owm.NewCurrent("c", "en", apiKey)
		if err != nil {
			fmt.Println("Bad API Key")
			log.Fatal(err)
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
	CurrentCmd.Flags().StringVarP(&coords, "coords", "l", "", "The coordinates in lat,long format to get the weather forcast for")
	CurrentCmd.Flags().StringVarP(&city, "city", "c", "", "The city to get the weather forcast for")
}
