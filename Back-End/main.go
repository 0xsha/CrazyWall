package main

import (
    "back-end/config"
    "back-end/database"
    "back-end/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/helmet/v2"
    "log"
)

func main() {



    // initialize database
    database.InitDB()

    // initialize fiber
    app := fiber.New()


    // setup security headers XSSProtection , ContentTypeNosniff , XFrameOptions , etc
    app.Use(helmet.New())

    // CORS setup allows only domain defined in config
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "*",
    }))

    // Stay online after crashes
    app.Use(recover.New())

    // Enable logger for dev
    if config.GetKey("APP_DEBUG")  == "true"  {
        app.Use(logger.New())


    }


    // initialize routes
    routes.InitRoutes(app)


    //if config.GetKey("APP_DEBUG")  == "false"  {
    //
    //
    //
    //    // Certificate manager
    //    m := &autocert.Manager{
    //        Prompt: autocert.AcceptTOS,
    //        // Replace with your domain
    //        HostPolicy: autocert.HostWhitelist("crazywall.io" , "api.crazywall.io"),
    //        // Folder to store the certificates
    //        Cache: autocert.DirCache("./certs"),
    //    }
    //
    //    // TLS Config
    //    cfg := &tls.Config{
    //        // Get Certificate from Let's Encrypt
    //        GetCertificate: m.GetCertificate,
    //        // By default NextProtos contains the "h2"
    //        // This has to be removed since Fasthttp does not support HTTP/2
    //        // Or it will cause a flood of PRI method logs
    //        // http://webconcepts.info/concepts/http-method/PRI
    //        NextProtos: []string{
    //            "http/1.1", "acme-tls/1",
    //        },
    //    }
    //    ln, err := tls.Listen("tcp", ":443", cfg)
    //    if err != nil {
    //        panic(err)
    //    }
    //
    //    // Start server
    //    log.Fatal(app.Listener(ln))
    //
    //
    //
    //    return
    //}

        /* add this stage app is rest configured let's init database*/
    err := app.Listen("0.0.0.0:" + config.GetKey("APP_PORT"))
    if err != nil {
        log.Println(err)
        return 
    }
}