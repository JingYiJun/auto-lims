package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func ListOpenableLab(c *fiber.Ctx) (err error) {
	var params struct {
		Apikey string `query:"apikey"`
		UID    string `query:"uid"`
	}
	err = c.QueryParser(&params)
	if err != nil {
		return err
	}

	username, err := checkApikey(params.Apikey)
	if err != nil {
		return err
	}

	var uid = config.DefaultUID
	if params.UID != "" {
		uid = params.UID
	}

	slog.LogAttrs(c.Context(), slog.LevelInfo, "list openable lab", slog.String("username", username), slog.String("uid", uid))

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":     "1",
			"pagesize": "20",
			"build":    "yfl",
			"search":   "",
			"appUser":  uid,
			"username": uid,
		}).
		SetHeader("Authorization", "Bearer "+config.Token).
		Get(listUrl)
	if err != nil {
		return err
	}

	var data []map[string]any
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return err
	}

	type sendStruct struct {
		DoorID int    `json:"doorID"`
		Name   string `json:"name"`
	}
	var respSend []sendStruct
	for _, d := range data {
		respSend = append(respSend, sendStruct{
			DoorID: int(d["agentid"].(float64)),
			Name:   d["lab_room_number"].(string),
		})
	}

	return c.JSON(respSend)
}

func OpenDoor(c *fiber.Ctx) (err error) {
	var params struct {
		Apikey string `query:"apikey"`
		DoorID string `query:"doorID"`
		UID    string `query:"uid"`
	}
	err = c.QueryParser(&params)
	if err != nil {
		return err
	}

	username, err := checkApikey(params.Apikey)
	if err != nil {
		return err
	}

	var uid = config.DefaultUID
	if params.UID != "" {
		uid = params.UID
	}

	slog.LogAttrs(c.Context(), slog.LevelInfo, "open door",
		slog.String("username", username),
		slog.String("uid", uid),
		slog.String("doorID", params.DoorID),
	)

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"agentId":  params.DoorID,
			"appUser":  uid,
			"username": uid,
		}).
		SetHeader("Authorization", "Bearer "+config.Token).
		Get(openUrl)
	if err != nil {
		return err
	}

	return c.Send(resp.Body())
}
