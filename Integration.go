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
	//path := c.Hostname() + c.OriginalURL()
	logrus.Debug("ProxyPostRequest url: ", c.BaseURL())
	logrus.Debug("ProxyPostRequest path: ", c.Path())
	logrus.Debug("ProxyPostRequest target url: ", viper.GetString("proxy.target")+c.Path())
	logrus.Debug("ProxyPostRequest requestData: ", string(c.Body()))
	targetURL := viper.GetString("proxy.target") + c.Path()
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
	logrus.Debug("ProxyPostRequest response: ", string(responseBody))
	return c.Status(response.StatusCode).SendString(string(responseBody))
}

func SendGetRequest(c *fiber.Ctx) error {
	//path := c.Params("path")
	logrus.Debug("ProxyGetRequest url: ", c.BaseURL())
	logrus.Debug("ProxyGetRequest path: ", c.Path())
	logrus.Debug("ProxyGetRequest target url: ", viper.GetString("proxy.target")+c.Path())
	targetURL := viper.GetString("proxy.target") + c.Path()
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
	logrus.Debug("ProxyPostRequest response: ", string(responseBody))
	return c.Status(response.StatusCode).SendString(string(responseBody))
}
