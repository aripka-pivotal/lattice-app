package helpers

import(
	"os"
)

func FetchAppName() string {
        appName := os.Getenv("APP_NAME")
        if appName == "" {
                return "Lattice-app"
        }
        return appName
}

