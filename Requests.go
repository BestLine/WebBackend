package main

import (
	"encoding/json"
	"fmt"
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

func PostAddVersion(c *fiber.Ctx) error {
	logrus.Debug("PostAddVersion")
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

func PostAddMethodic(c *fiber.Ctx) error {
	logrus.Debug("PostAddMethodic")
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
				"[{\"application\":\"StartTime 2023-10-02 15-05\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409474382\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-47\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409834495\"}," +
				"{\"application\":\"StartTime 2023-10-02 15-00 Stable Test 3 INAC 5rps\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409473677\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-39\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=408949351\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-39\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409833189\"}," +
				"{\"application\":\"StartTime\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409473120\"}," +
				"{\"application\":\"StartTime 2023-10-02 11-51 Stable test 18rps DOL\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409832239\"}," +
				"{\"application\":\"StartTime 2023-09-29 10-21Stability\",\"bucket\":\"jmeter_IntergraStand_BM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409831416\"}," +
				"{\"application\":\"StartTime 2023-10-01 11-57 COMM_LT_all_sms_email_voice_bell_250_65_15_16_tps\",\"bucket\":\"jmeter_COMM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=408945128\"}," +
				"{\"application\":\"StartTime  SmokeCrush\",\"bucket\":\"jmeter_Pretium\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409830287\"}]}")
		} else {
			return c.SendString("{\"count\":10,\"data\":" +
				"[{\"application\":\"StartTime 2023-10-02 15-05\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409474382\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-47\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409834495\"}," +
				"{\"application\":\"StartTime 2023-10-02 15-00 Stable Test 3 INAC 5rps\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409473677\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-39\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=408949351\"}," +
				"{\"application\":\"StartTime 2023-10-02 14-39\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409833189\"}," +
				"{\"application\":\"StartTime\",\"bucket\":\"jmeter_IDP\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409473120\"}," +
				"{\"application\":\"StartTime 2023-10-02 11-51 Stable test 18rps DOL\",\"bucket\":\"jmeter_OM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409832239\"}," +
				"{\"application\":\"StartTime 2023-09-29 10-21Stability\",\"bucket\":\"jmeter_IntergraStand_BM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409831416\"}," +
				"{\"application\":\"StartTime 2023-10-01 11-57 COMM_LT_all_sms_email_voice_bell_250_65_15_16_tps\",\"bucket\":\"jmeter_COMM\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=408945128\"}," +
				"{\"application\":\"StartTime  SmokeCrush\",\"bucket\":\"jmeter_Pretium\",\"cfurl\":\"https://confluence.veon.com/pages/viewpage.action?pageId=409830287\"}]}")
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

func GetHostList(c *fiber.Ctx) error {
	logrus.Debug("GetHostList")
	if viper.GetBool("server.debug") {
		var hosts []interface{}
		hosts = []interface{}{
			HostData{Host: "https://qa.load.com"},
			HostData{Host: "https://shit.box.ru"},
		}
		jsonData, err := json.Marshal(hosts)
		if err != nil {
			logrus.Error(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Отправка JSON-ответа
		c.Set("Content-Type", "application/json")
		fmt.Println(jsonData)
		//return c.Send(jsonData)
		return c.SendString("{\"host\":[\"http://qa-auto.vimpelcom.ru:8086\",\"http://ms-loadrtst026:8086\"]}")
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
