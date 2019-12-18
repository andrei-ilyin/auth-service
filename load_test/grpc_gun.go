// create a package
package main

// import some pandora stuff
// and stuff you need for your scenario
// and protobuf contracts for your grpc service

import (
        "log"
        "context"
        "time"

        "github.com/spf13/afero"
        "google.golang.org/grpc"

        pb "github.com/andrei_ilyin/auth-service"

        "github.com/yandex/pandora/cli"
        "github.com/yandex/pandora/components/phttp/import"
        "github.com/yandex/pandora/core"
        "github.com/yandex/pandora/core/aggregator/netsample"
        "github.com/yandex/pandora/core/import"
        "github.com/yandex/pandora/core/register"
)

type Ammo struct {
        Tag         string
        Param1      string
        Param2      string
        Param3      string
}

type Sample struct {
        URL              string
        ShootTimeSeconds float64
}

type GunConfig struct {
        Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
        // Configured on construction.
        client grpc.ClientConn
        conf   GunConfig
        // Configured on Bind, before shooting
        aggr core.Aggregator // May be your custom Aggregator.
        core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
        return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
        // create gRPC stub at gun initialization
        conn, err := grpc.Dial(
                g.conf.Target,
                grpc.WithInsecure(),
                grpc.WithTimeout(time.Second),
                grpc.WithUserAgent("load test, pandora custom shooter"))
        if err != nil {
                log.Fatalf("FATAL: %s", err)
        }
        g.client = *conn
        g.aggr = aggr
        g.GunDeps = deps
        return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
        customAmmo := ammo.(*Ammo)
        g.shoot(customAmmo)
}


func (g *Gun) send_login_call(client pb.AuthenticatorClient, ammo *Ammo) int {
        // Prepare request
        credentials := pb.Credentials{
            UserName: ammo.Param1,
            Password: ammo.Param2,
        }
        request := pb.LoginRequest{
            Credentials: &credentials,
        }

        // RPC call
        response, err := client.Login(context.TODO(), &request)

        // Process reply
        code := 0
        if err != nil {
                log.Printf("FATAL: %s", err)
                code = 500
        }
        if response != nil {
                code = 200
        }
        return code
}

func (g *Gun) shoot(ammo *Ammo) {
        code := 0
        sample := netsample.Acquire(ammo.Tag)

        conn := g.client
        client := pb.NewAuthenticatorClient(&conn)

        switch ammo.Tag {
        case "/login":
            code = g.send_login_call(client, ammo)
        default:
            code = 404
        }

        defer func() {
                sample.SetProtoCode(code)
                g.aggr.Report(sample)
        }()
}

func main() {
        //debug.SetGCPercent(-1)

        // Standard imports.
        fs := afero.NewOsFs()
        coreimport.Import(fs)

        // May not be imported, if you don't need http guns and etc.
        phttp.Import(fs)

        // Custom imports. Integrate your custom types into configuration system.
        coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

        register.Gun("grpc_gun", NewGun, func() GunConfig {
                return GunConfig{
                        Target: "default target",
                }
        })

        cli.Run()
}
