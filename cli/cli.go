package cli

import (
	"flag"
	"fmt"
	"os"
)

func paramError(argName string) error {
	return fmt.Errorf("environment variable or argument <%s> is required", argName)
}

// GetParams gets parameters from the environment variables or cli flags and might throw an error
func GetParams() (map[string]string, error) {
	params := make(map[string]string)

	var address, pghost, pgport, pgdatabase, pguser, pgpass string

	// prefer environment variables to arguments
	var set bool

	if address, set = os.LookupEnv("address"); set == false {
		flag.StringVar(&address, "address", "localhost", "the address of the web server")
	}

	if pghost, set = os.LookupEnv("pghost"); set == false {
		flag.StringVar(&pghost, "pghost", "localhost", "the address of the database")
	}

	if pgport, set = os.LookupEnv("pgport"); set == false {
		flag.StringVar(&pgport, "pgport", "5432", "the TCP port of the database")
	}

	if pgdatabase, set = os.LookupEnv("pgdatabase"); set == false {
		flag.StringVar(&pgdatabase, "pgdatabase", "", "the database name (required)")
	}

	if pguser, set = os.LookupEnv("pguser"); set == false {
		flag.StringVar(&pguser, "pguser", "", "the database user (required)")
	}

	if pgpass, set = os.LookupEnv("pgpass"); set == false {
		flag.StringVar(&pgpass, "pgpass", "", "the database user credential (required)")
	}

	flag.Parse()

	if pgdatabase == "" {
		return nil, paramError("pgdatabase")
	} else if pguser == "" {
		return nil, paramError("pguser")
	} else if pgpass == "" {
		return nil, paramError("pgpass")
	}

	params = map[string]string{
		"address":    address,
		"pghost":     pghost,
		"pgport":     pgport,
		"pgdatabase": pgdatabase,
		"pguser":     pguser,
		"pgpass":     pgpass,
	}

	return params, nil
}
