package helpers

import(
	"os"
	"encoding/json"
)

func FetchAppName() string {
        appName := os.Getenv("APP_NAME")
        if appName == "" {

		//get the app name from lattice name
		appName = os.Getenv("PROCESS_GUID")
		if appName == "" {
			vcapStr := os.Getenv("VCAP_APPLICATION")
			var vcapMap map[string]string

			json.Unmarshal([]byte(vcapStr), &vcapMap)

			appName = vcapMap["application_name"]
			if appName == "" {	
                		return "Lattice-app"
			}
		}
        }
        return appName
}

