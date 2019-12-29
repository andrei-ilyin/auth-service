// create a package
package main

// import some pandora stuff
// and stuff you need for your scenario
// and protobuf contracts for your grpc service

import (
    "log"
    "context"
    // "strconv"
    "strings"
    "time"

    // "github.com/golang/protobuf/ptypes/timestamp"
    // "github.com/satori/go.uuid"

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

func (g *Gun) regular_scenario(client pb.AuthenticatorClient, ammo *Ammo) int {
    // Create session
    credentials := pb.Credentials{
        UserName: ammo.Param1,
        Password: ammo.Param2,
    }
    login_request := pb.LoginRequest{
        Credentials: &credentials,
    }
    clientDeadline := time.Now().Add(200 * time.Millisecond)
    ctx, _ := context.WithDeadline(context.TODO(), clientDeadline)
    login_response, err := client.Login(ctx, &login_request)
    if ctx.Err() == context.Canceled {
        return 503
    }
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
                case codes.InvalidArgument:
                    return 403
                case codes.DeadlineExceeded:
                    return 503
                default:
                    return 500
            }
        }
    }
    cookie := login_response.GetCookie()

    // Query resources
    for _, resource := range strings.Split(ammo.Param3, ",") {
        validation_request := pb.ValidationRequest{
            Cookie: cookie,
            Resource: resource,
        }
        clientDeadline = time.Now().Add(200 * time.Millisecond)
        ctx, _ = context.WithDeadline(context.TODO(), clientDeadline)
        validation_response, err := client.Validate(
            ctx, &validation_request)
        if ctx.Err() == context.Canceled {
            return 503
        }
        if err != nil {
            if e, ok := status.FromError(err); ok {
                switch e.Code() {
                    case codes.InvalidArgument:
                        return 403
                    case codes.DeadlineExceeded:
                        return 503
                    default:
                        return 500
                }
            }
        }
        if (validation_response == nil) {
            return 500
        }
    }

    // Remove session
    logout_request := pb.LogoutRequest{
        Cookie: cookie,
    }
    clientDeadline = time.Now().Add(200 * time.Millisecond)
    ctx, _ = context.WithDeadline(context.TODO(), clientDeadline)
    logout_response, err := client.Logout(ctx, &logout_request)
    if ctx.Err() == context.Canceled {
        return 503
    }
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
                case codes.InvalidArgument:
                    return 403
                case codes.DeadlineExceeded:
                    return 503
                default:
                    return 500
            }
        }
    }
    if (logout_response == nil) {
        return 500
    }

    return 200
}

/*
func (g *Gun) case2_method(client pb.AuthenticatorClient, ammo *Ammo) int {
        code := 0

        var itemIDs []int64
        for _, id := range strings.Split(ammo.Param1, ",") {
                if id == "" {
                        continue
                }
                itemID, err := strconv.ParseInt(id, 10, 64)
                if err != nil {
                        log.Printf("Ammo parse FATAL: %s", err)
                        code = 314
                }
                itemIDs = append(itemIDs, itemID)
        }

        // prepare item_id and warehouse_id
        item_id, err := strconv.ParseInt(ammo.Param1, 10, 0)
        if err != nil {
                log.Printf("Failed to parse ammo FATAL", err)
                code = 314
        }
        warehouse_id, err2 := strconv.ParseInt(ammo.Param2, 10, 0)
        if err2 != nil {
                log.Printf("Failed to parse ammo FATAL", err2)
                code = 314
        }

        items := []*pb.SomeItem{}
        items = append(items, &pb.SomeItem{
                item_id,
                warehouse_id,
                1,
                &timestamp.Timestamp{time.Now().Unix(), 111}
        })

        out2, err3 := client.GetSomeDataSecond(
                context.TODO(), &pb.SomeRequest{
                        uuid.Must(uuid.NewV4()).String(),
                        1,
                        items})
        if err3 != nil {
                log.Printf("FATAL", err3)
                code = 316
        }

        if out2 != nil {
                code = 200
        }

        return code
}
*/

func (g *Gun) shoot(ammo *Ammo) {
    code := 0
    sample := netsample.Acquire(ammo.Tag)

    conn := g.client
    client := pb.NewAuthenticatorClient(&conn)

    switch ammo.Tag {
        case "no_queries":
            code = g.regular_scenario(client, ammo)
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
