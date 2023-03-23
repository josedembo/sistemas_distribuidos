package main

import "github.com/gofiber/fiber/v2"

type Vars struct{
    Msg string
}

var vars = new(Vars)

func main() {
    app := fiber.New()

    app.Get("/api", func(c *fiber.Ctx) error {
        vars.Msg = "Este é o método GET"
        return c.Render("template.html", vars)
    })

    app.Post("/api", func (c *fiber.Ctx) error {
        vars.Msg = "Este é o metodo POST"
        return c.Render("template.html", vars)
        
    })

    app.Delete("/api", func (c *fiber.Ctx) error{
        vars.Msg = "Este é o metodo DELETE"
        return c.Render("template.html", vars)
    })

    app.Put("/api", func (c *fiber.Ctx) error{
        vars.Msg = "Este é o método PUT"
        return c.Render("template.html", vars)
    })

    

    app.Listen(":3000")
}