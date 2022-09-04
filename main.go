package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	paymenttripay "gotripay/paymentTripay"
	"gotripay/tripay"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home go run main",
		})
	})
	r.GET("/payment", func(c *gin.Context) {

		tr := paymenttripay.New("DEV-WhvAPUhrvIiMTklIn1CTp3WIJs1vJLP99MHGTcJl", "SvkVQ-5kEgq-4tx1r-fD10X-0rZ4U", "T11858", paymenttripay.Development)

		req := paymenttripay.RequestTransaction{
			PaymentMethod: "BNIVA",
			MerchantRef:   "INV69",
			Amount:        20000,
			CustomerName:  "akbarfa",
			CustomerEmail: "fania@123.com",
			CustomerPhone: "081234567891",
			OrderItems:    []paymenttripay.Item{0: {Sku: "duar", Name: "duar", Price: 20000, Quantity: 1}},
		}
		b, err := tr.RequestClosedTransaction(req)
		if err != nil {
			log.Panic(err)
			return
		}
		v := tripay.ClosedTransactionResponse{}
		if err := json.Unmarshal(b, &v); err != nil {
			log.Panicln(err)
		}
		fmt.Println(v.Success)
		// fmt.Println(v.Data.Reference)
		c.JSON(http.StatusOK, gin.H{
			"message": v,
		})
		// os.WriteFile("dump/requestclosed.json", b, 0644)

	})
	r.POST("/callback", func(c *gin.Context) {
		var input tripay.Callback

		err := c.ShouldBindJSON(&input)
		if err != nil {
			errorMessage := gin.H{"errors": err}

			c.JSON(http.StatusUnprocessableEntity, errorMessage)
			return
		}

		fmt.Println(c.Request.Header.Get("X-Callback-Signature"))
		xcallbacksignature := c.Request.Header.Get("X-Callback-Signature")
		tr := paymenttripay.New("DEV-WhvAPUhrvIiMTklIn1CTp3WIJs1vJLP99MHGTcJl", "SvkVQ-5kEgq-4tx1r-fD10X-0rZ4U", "T11858", paymenttripay.Development)

		// fmt.Println(tr)
		h := hmac.New(sha256.New, tr.ApiKey)
		b, err := json.Marshal(&input)
		if err != nil {
			return
		}
		h.Write(b)
		signature := hex.EncodeToString(h.Sum(nil))
		c.JSON(http.StatusOK, gin.H{
			"server":   signature,
			"callback": xcallbacksignature,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//notifikasi FCM Method POST

// 	url := "https://fcm.googleapis.com/fcm/send"

// 	requestBody := strings.NewReader(`
// {
// 	"to":"coKAOJCzS0-mn8CNSKRBlF:APA91bGGhI4MvlTJ0zNliTAjmhDV9nC0n5yUp01LhlvdI0EQjkakPte5BW0srVfQ7UbETApv3z9bhQzeGWEF78dOlPdtgvCvixhEjRYC9sOnD56LO7mLU2VJbA3eiiXqwKGVvfLHf639",
// 	"priority": "high",
// 	"soundName": "default",
// 	"notification": {
// 	"title": "Halo Wandi Pratama",
// 	"body": "Ini adalah FCM melalui postman"
// 	}
// 	}
// `)

// 	req, err := http.NewRequest("POST", url, requestBody)
// 	req.Header.Add("Authorization", "key=AAAACoktzJI:APA91bEn9fap-WEvJRM_r3240PTuWz7GDPONKKAgQajHOmMZf3lRbRRCnh7Jr-6Hq_QoTzgQ7t2LxuINfFa-EXvvImCD1_KO82wurhxg1u-Yj8RkDbv03HJ7T7mVzdx06gtNuRJYNHhl")
// 	req.Header.Add("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error on response.\n[ERROR] -", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("Error while reading the response bytes:", err)
// 	}
// 	log.Println(string([]byte(body)))
