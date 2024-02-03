package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) issueTime(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issuePrice:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(3 * time.Second)
		sendTimeRequest(input)
	}()
}

func sendTimeRequest(request models.Request) {
	var timeEntry string
	if rand.Intn(10)%10 > 3 {
		timeEntry = "\nПриказ подписан Ректором " + (time.Now().Add(time.Hour * 24 * time.Duration(rand.Intn(10)))).Format(time.DateOnly)

		fmt.Println("calculated")
	} else {
		timeEntry = ""
	}
	answer := models.TimeRequest{
		AccessToken: 123,
		Signature:   timeEntry,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:2023/api/orders/%d/update_signature", request.OrderId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
