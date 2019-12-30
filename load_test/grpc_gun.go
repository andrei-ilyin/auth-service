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
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

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

func (g *Gun) login(client pb.AuthenticatorClient, ammo *Ammo) int {
    // Prepare request
    credentials := pb.Credentials{
        UserName: ammo.Param1,
        Password: ammo.Param2,
    }
    cookie := pb.Cookie{
        SessionId: ammo.Param3,
    }
    request := pb.LoginRequest{
        Credentials: &credentials,
        Cookie: &cookie,
    }

    // Perform gRPC call
    deadline := time.Now().Add(1000 * time.Millisecond)
    ctx, _ := context.WithDeadline(context.TODO(), deadline)
    response, err := client.Login(ctx, &request)
    if ctx.Err() == context.Canceled {
        return 503
    }

    // Check return code
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
                case codes.DeadlineExceeded:
                    return 503
                default:
                    log.Printf("FATAL", err)
                    return 500
            }
        }
    }

    // Check response
    if response == nil {
        return 500
    }
    if response.GetStatus().GetCode() == pb.Status_INTERNAL_ERROR {
        return 500
    }

    return 200
}

func (g *Gun) logout(client pb.AuthenticatorClient, ammo *Ammo) int {
    // Prepare request
    cookie := pb.Cookie{
        SessionId: ammo.Param1,
    }
    request := pb.LogoutRequest{
        Cookie: &cookie,
    }

    // Perform gRPC call
    deadline := time.Now().Add(1000 * time.Millisecond)
    ctx, _ := context.WithDeadline(context.TODO(), deadline)
    response, err := client.Logout(ctx, &request)
    if ctx.Err() == context.Canceled {
        return 503
    }

    // Check return code
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
                case codes.DeadlineExceeded:
                    return 503
                default:
                    log.Printf("FATAL", err)
                    return 500
            }
        }
    }

    // Check response
    if response == nil {
        return 500
    }
    if response.GetStatus().GetCode() == pb.Status_INTERNAL_ERROR {
        return 500
    }

    return 200
}

func (g *Gun) validate(client pb.AuthenticatorClient, ammo *Ammo) int {
    // Prepare request
    cookie := pb.Cookie{
        SessionId: ammo.Param1,
    }
    request := pb.ValidationRequest{
        Cookie: &cookie,
        Resource: ammo.Param2,
    }

    // Perform gRPC call
    deadline := time.Now().Add(1000 * time.Millisecond)
    ctx, _ := context.WithDeadline(context.TODO(), deadline)
    response, err := client.Validate(ctx, &request)
    if ctx.Err() == context.Canceled {
        return 503
    }

    // Check return code
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
                case codes.DeadlineExceeded:
                    return 503
                default:
                    log.Printf("FATAL", err)
                    return 500
            }
        }
    }

    // Check response
    if response == nil {
        return 500
    }
    if response.GetStatus().GetCode() == pb.Status_INTERNAL_ERROR {
        return 500
    }

    return 200
}

func (g *Gun) shoot(ammo *Ammo) {
    code := 0
    sample := netsample.Acquire(ammo.Tag)

    conn := g.client
    client := pb.NewAuthenticatorClient(&conn)

    switch ammo.Tag {
        case "Login":
            code = g.login(client, ammo)
        case "Logout":
            code = g.logout(client, ammo)
        case "Validate":
            code = g.validate(client, ammo)
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
