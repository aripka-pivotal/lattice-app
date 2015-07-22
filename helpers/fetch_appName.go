package helpers

import(
	"os"
)

func FetchAppName() string {
        appName := os.Getenv("APP_NAME")
        if appName == "" {

		//get the app name from lattice name
		appName = os.Getenv("PROCESS_GUID")
		if appName == "" {
                	return "Lattice-app"
		}
        }
        return appName
}

