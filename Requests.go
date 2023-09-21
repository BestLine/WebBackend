package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func settings(c *fiber.Ctx) error {
	logrus.Debug("settings")
	return c.SendString("Target host: " + viper.GetString("proxy.target") + "\n" +
		"Debug mode: " + viper.GetString("server.debug"))
}

func PostCompareRelease(c *fiber.Ctx) error {
	logrus.Debug("PostCompareRelease")
	if viper.GetBool("server.debug") {
		return c.SendString("DEBUG MODE ENABLED!\nRESULT: OK")
	} else {
		err := SendPostRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func PostCreateBucket(c *fiber.Ctx) error {
	logrus.Debug("PostCreateBucket")
	if viper.GetBool("server.debug") {
		return c.SendString("DEBUG MODE ENABLED!\nRESULT: OK")
	} else {
		err := SendPostRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTableDataCurrentTests(c *fiber.Ctx) error {
	logrus.Debug("GetTableDataCurrentTests")
	if viper.GetBool("server.debug") {
		return c.SendString("[{\"SystemName\":\"IDP\",\"Bucket\":\"Jmeter_IDP\",\"Status\":\"В процессе\"}," +
			"{\"SystemName\":\"BackCRM\",\"Bucket\":\"Jmeter_BackCRM\",\"Status\":\"Присутствуют ошибки\"}]")
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTableDataReports(c *fiber.Ctx) error {

	bucket := c.Query("bucket")
	count := c.Query("count")
	logrus.Debug("GetTableDataReports: bucket = ", bucket, " count = ", count)
	//fmt.Println(count)
	if viper.GetBool("server.debug") {
		if bucket != "" {
			return c.SendString("{\"count\":10,\"data\":" +
				"[{\"application\":\"StartTime 2023-08-31 13-58 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=389557398\"}," +
				"{\"application\":\"StartTime 2023-08-23 01-38 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386705219\"}," +
				"{\"application\":\"StartTime 2023-08-23 00-07 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386705153\"}," +
				"{\"application\":\"StartTime 2023-08-22 16-02 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386704172\"}," +
				"{\"application\":\"StartTime 2023-08-21 09-15 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386698345\"}," +
				"{\"application\":\"StartTime 2023-08-18 09-55 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=383969194\"}," +
				"{\"application\":\"StartTime 2023-08-18 07-15 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384617383\"}," +
				"{\"application\":\"StartTime 2023-08-18 01-03 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384617004\"}," +
				"{\"application\":\"StartTime 2023-08-17 22-37 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384616672\"}," +
				"{\"application\":\"StartTime 2023-08-17 14-49 Max perf VLB2 0del\",\"bucket\":\"jmeter_MOCK\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=383967467\"}]}")
		} else {
			return c.SendString("{\"count\":10,\"data\":" +
				"[{\"application\":\"StartTime 2023-08-31 13-58 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=389557398\"}," +
				"{\"application\":\"StartTime 2023-08-23 01-38 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386705219\"}," +
				"{\"application\":\"StartTime 2023-08-23 00-07 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386705153\"}," +
				"{\"application\":\"StartTime 2023-08-22 16-02 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386704172\"}," +
				"{\"application\":\"StartTime 2023-08-21 09-15 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=386698345\"}," +
				"{\"application\":\"StartTime 2023-08-18 09-55 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=383969194\"}," +
				"{\"application\":\"StartTime 2023-08-18 07-15 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384617383\"}," +
				"{\"application\":\"StartTime 2023-08-18 01-03 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384617004\"}," +
				"{\"application\":\"StartTime 2023-08-17 22-37 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=384616672\"}," +
				"{\"application\":\"StartTime 2023-08-17 14-49 Max perf VLB2 0del\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=383967467\"}]}")
		}
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTableDataTests(c *fiber.Ctx) error {
	logrus.Debug("GetTableDataTests")
	if viper.GetBool("server.debug") {
		return c.SendString("[{\"IDP\", \"Jmeter_IDP\", \"5:56 28.07.2023\", \"В процессе проведения\", \"MaxPerf\"}," +
			"{\"BackCRM\", \"Jmeter_BackCRM\", \"3:20 28.07.2023\", \"Остановлен\", \"Stability\"}," +
			"{\"BackCRM\", \"Jmeter_BackCRM\", \"0:12 28.07.2023\", \"Завершён с ошибками\", \"MaxPerf\"}]")
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTableDataStatus(c *fiber.Ctx) error {
	logrus.Debug("GetTableDataStatus")
	if viper.GetBool("server.debug") {
		return c.SendString("[{\"Генераторы\", \"5 минут назад\", \"Штатный режим\"}," +
			"{\"Автоотчёт\", \"3 минуты назад\", \"Штатный режим\"}," +
			"{\"Jenkins\", \"243 минуты назад\", \"Не отвечает\"}]")
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetBucketList(c *fiber.Ctx) error {
	project := c.Query("project")
	bucket := c.Query("bucket")
	logrus.Debug("GetBucketList")
	logrus.Debug("project = ", project)
	logrus.Debug("bucket = ", bucket)
	if viper.GetBool("server.debug") {
		var buckets []interface{}
		if bucket != "" {
			//var buckets []ProjectData
			buckets = []interface{}{
				ProjectData{Project: "Project_1"},
				ProjectData{Project: "Project_2"},
				ProjectData{Project: "Project_3"},
				ProjectData{Project: "Project_4"},
				ProjectData{Project: "Project_5"},
			}
		} else {
			//var buckets []BucketData
			buckets = []interface{}{
				BucketData{Bucket: "Bucket_1"},
				BucketData{Bucket: "Bucket_2"},
				BucketData{Bucket: "Bucket_3"},
			}
		}
		// Преобразование данных в JSON
		jsonData, err := json.Marshal(buckets)
		if err != nil {
			logrus.Error(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Отправка JSON-ответа
		c.Set("Content-Type", "application/json")
		return c.Send(jsonData)
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetVersionsList(c *fiber.Ctx) error {
	project := c.Query("project")
	logrus.Debug("GetVersionsList")
	logrus.Debug("project = ", project)
	if viper.GetBool("server.debug") {
		var versions []interface{}
		versions = []interface{}{
			VersionData{Version: "1"},
			VersionData{Version: "2"},
			VersionData{Version: "3"},
		}
		jsonData, err := json.Marshal(versions)
		if err != nil {
			logrus.Error(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Отправка JSON-ответа
		c.Set("Content-Type", "application/json")
		return c.Send(jsonData)
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetProjectList(c *fiber.Ctx) error {
	if viper.GetBool("server.debug") {
		var buckets []interface{}
		buckets = []interface{}{
			ProjectData{Project: "Project_1"},
			ProjectData{Project: "Project_2"},
			ProjectData{Project: "Project_3"},
			ProjectData{Project: "Project_4"},
			ProjectData{Project: "Project_5"},
		}
		// Преобразование данных в JSON
		jsonData, err := json.Marshal(buckets)
		if err != nil {
			logrus.Error(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Отправка JSON-ответа
		c.Set("Content-Type", "application/json")
		return c.Send(jsonData)
	} else {
		err := SendGetRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}
