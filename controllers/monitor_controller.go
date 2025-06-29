package controllers

import (
	"vps-monitor/services"

	"github.com/gofiber/fiber/v2"
)

type MetricsController struct {
	systemService *services.SystemService
}

func NewMetricsController(systemService *services.SystemService) *MetricsController {
	return &MetricsController{
		systemService: systemService,
	}
}

func (c *MetricsController) GetMetricsInfo(ctx *fiber.Ctx) error {
	data, err := c.systemService.GetMetricsInfo()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"data": data,
	})
}

func (c *MetricsController) GetStorageInfo(ctx *fiber.Ctx) error {
	storageInfo, err := c.systemService.GetStorageUsage()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(storageInfo)
}
