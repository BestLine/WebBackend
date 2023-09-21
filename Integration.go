package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

func SendPostRequest(c *fiber.Ctx) error {
	path := c.Params("path")
	logrus.Debug("ProxyPostRequest: ", path)
	logrus.Debug("ProxyPostRequest requestData: ", c.Body())
	targetURL := viper.GetString("proxy.target") + path
	response, err := http.Post(targetURL, "application/json", bytes.NewBuffer(c.Body()))
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(response.StatusCode).SendString(string(responseBody))
}

func SendGetRequest(c *fiber.Ctx) error {
	path := c.Params("path")
	logrus.Debug("ProxyGetRequest: ", path)
	targetURL := viper.GetString("proxy.target") + path
	response, err := http.Get(targetURL)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(response.StatusCode).SendString(string(responseBody))
}
