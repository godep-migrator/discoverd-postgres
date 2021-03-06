package main

import (
    "os"

    "github.com/codegangsta/cli"
    "github.com/dictybase/discoverd-postgres/centos"
)

func main() {
    pgver := "9.2"
    if len(os.Getenv("PGVERSION")) > 1 {
        pgver = os.Getenv("PGVERSION")
    }
    app := cli.NewApp()
    app.Name = "discoverd-postgres"
    app.Version = "1.0.0"
    app.Usage = "Starts an instance of postgresql server"
    app.Flags = []cli.Flag{
        cli.StringFlag{"data", "/var/lib/pgsql/" + pgver + "/data", "postgresql data directory"},
        cli.StringFlag{"service", "postgresql", "discoverd service name"},
        cli.StringFlag{"pgbin", "/usr/pgsql-" + pgver + "/bin", "postgres binary directory"},
        cli.StringFlag{"discoverd", "", "discoverd service ip, by default will be picked up from DISCOVERD env variable"},
        cli.StringFlag{"port", "5432", "postgresql port, default is 5432"},
        cli.StringFlag{"config", "/config", "folder from where postgresql config files will be copied during start up"},
    }
    app.Action = func(c *cli.Context) {
        centos.RunPg(c)
    }
    app.Run(os.Args)
}
