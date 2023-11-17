package handlers

import (
	"fmt"
	"go-sample-webserver/pkg/renders"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HomeHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"home.page.tmpl",
	)
	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func AboutHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"about.page.tmp2l",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func handleError(ctx *fiber.Ctx, err error) *fiber.Error {

	log.Print("Error processing request ", err)
	return fiber.NewError(
		fiber.StatusInternalServerError,
		fmt.Sprintf("Error processing request \n %v", err),
	)
}
