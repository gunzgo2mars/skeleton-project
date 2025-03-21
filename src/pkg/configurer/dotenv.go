package configurer

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func LoadDotEnv(
	secrets interface{},
	secretPathName string,
	secretPrefix string,
	osenv string,
) error {
	if osenv == "" {
		osenv = "APPENV"
	}

	currentEnvironment, ok := os.LookupEnv(osenv)
	if ok {
		secretPathName = secretPathName + "." + currentEnvironment + ".env"
		log.Printf("LoadDotEnv env: %s, secretPathName: %s\n", currentEnvironment, secretPathName)
		if err := godotenv.Load(secretPathName); err != nil {
			panic(err)
		}
	}

	// env to struct
	err := envconfig.Process(secretPrefix, secrets)
	if err != nil {
		panic("Error unmarshalling env vars")
	}
	return nil
}

func HotfixNewLineCert(certString string) string {
	newLinedString := strings.ReplaceAll(certString, `\n`, "\n")
	return newLinedString
}
