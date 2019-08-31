/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

package model

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
