package handlers

import (
	"fmt"
	"go-sample-webserver/src/renders"

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
		"about.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func MajorsRoomHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"majors.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func GeneralsRoomHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"generals.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func SearchAvailabilityHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"search-availability.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func MakeReservationsHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"make-reservation.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func ContactHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html; charset=utf-8")

	err := renders.RenderHtmlTemplate(
		ctx.Response().BodyWriter(),
		"contact.page.tmpl",
	)

	if err != nil {
		return handleError(ctx, err)
	}
	return nil
}

func handleError(ctx *fiber.Ctx, err error) *fiber.Error {

	logger := ctx.Context().Logger()
	logger.Printf("Error processing request2 %v", err)

	return fiber.NewError(
		fiber.StatusInternalServerError,
		fmt.Sprintf("Error processing request \n %v", err),
	)
}
