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
    Now:         {{.Main.Temp}} {{.Unit}}
    High:        {{.Main.TempMax}} {{.Unit}}
    Low:         {{.Main.TempMin}} {{.Unit}}
    Wind speed:  {{.Wind.Speed}}m/s
    Wind Dir:    {{.Wind.Deg}}°
`

func getCoords(w *owm.CurrentWeatherData) {
	lat, long, err := utils.ProcessCoords(coords)
	if err != nil {
		// error handled in function
		return
	}

	w.CurrentByCoordinates(&owm.Coordinates{
		Longitude: long,
		Latitude:  lat,
	})

	tmpl, err := template.New("current").Parse(currentTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from request")
		return
	}

	tmpl.Execute(os.Stdout, w)
}

func getCity(w *owm.CurrentWeatherData) {
	w.CurrentByName(city)

	tmpl, err := template.New("current").Parse(currentTemplate)
	if err != nil {
		fmt.Println("Bad JSON returned from request")
		return
	}

	tmpl.Execute(os.Stdout, w)
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

		// Set default values
		runCoords, runCity, err = utils.CheckFlags(coords, city)
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal(err)
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		w, err := owm.NewCurrent("c", "en", apiKey)
		if err != nil {
			fmt.Println("Bad API Key")
			log.Fatal(err)
			return
		}

		if runCoords {
			getCoords(w)
		} else if runCity {
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
