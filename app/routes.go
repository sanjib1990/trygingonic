package app

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"trygonic/app/config"
	"trygonic/app/utils/Logger"
	"trygonic/app/utils/response"

	"github.com/gin-gonic/gin"
)

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	Logger.Get().Info(frame.File + " " + strconv.Itoa(frame.Line) + " " + frame.Function)
}

func RegisterRoutes(engine *gin.Engine) {
	authEngine := engine.Group("/client/v4", func(c *gin.Context) {
		authEmail := c.GetHeader("X-Auth-Email")

		if config.Values.AuthKey != authEmail {
			response.Send(c, gin.H{}, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		authKey := c.GetHeader("X-Auth-Key")

		if config.Values.AuthSecret != authKey {
			response.Send(c, gin.H{}, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		c.Next()
	})

	authEngine.DELETE("/zones/:zoneId/custom_hostnames/:domainId", func(c *gin.Context) {
		Logger.Get().Info("ZoneID " + c.Param("zoneId") + " DOMAIN ID " + c.Param("domainId"))
		response.Send(c, gin.H{
			"id": "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	authEngine.POST("/zones/:zoneId/custom_hostnames", func(c *gin.Context) {
		Logger.Get().Info("ZoneID " + c.Param("zoneId"))
		response.Send(c, gin.H{
			"success":  true,
			"errors":   []string{},
			"messages": []string{},
			"result": map[string]interface{}{
				"id":       "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
				"hostname": "app.example.com",
				"ssl": map[string]interface{}{
					"id":     "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
					"status": "pending_validation",
					"method": "txt",
					"type":   "dv",
					"validation_records": []interface{}{
						map[string]interface{}{
							"txt_name":  "_acme-challenge.app.example.com",
							"txt_value": "810b7d5f01154524b961ba0cd578acc2",
							"http_url":  "http://app.example.com/.well-known/pki-validation/ca3-da12a1c25e7b48cf80408c6c1763b8a2.txt",
							"http_body": "ca3-574923932a82475cb8592200f1a2a23d",
							"emails": []string{
								"administrator@example.com",
								"webmaster@example.com",
							},
						},
					},
					"validation_errors": []interface{}{
						map[string]string{
							"message": "SERVFAIL looking up CAA for app.example.com",
						},
					},
					"hosts": []string{
						"app.example.com",
						"*.app.example.com",
					},
					"issuer":        "DigiCertInc",
					"serial_number": "6743787633689793699141714808227354901",
					"signature":     "SHA256WithRSA",
					"uploaded_on":   "2020-02-06T18:11:23.531995Z",
					"expires_on":    "2021-02-06T18:11:23.531995Z",
					"custom_csr_id": "7b163417-1d2b-4c84-a38a-2fb7a0cd7752",
					"settings": map[string]interface{}{
						"http2":           "on",
						"min_tls_version": "1.2",
						"tls_1_3":         "on",
						"ciphers": []string{
							"ECDHE-RSA-AES128-GCM-SHA256",
							"AES128-SHA",
						},
						"early_hints": "on",
					},
					"bundle_method":         "ubiquitous",
					"wildcard":              false,
					"certificate_authority": "digicert",
					"custom_certificate":    "-----BEGIN CERTIFICATE-----\\nMIIFJDCCBAygAwIBAgIQD0ifmj/Yi5NP/2gdUySbfzANBgkqhkiG9w0BAQsFADBN\\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMScwJQYDVQQDEx5E...SzSHfXp5lnu/3V08I72q1QNzOCgY1XeL4GKVcj4or6cT6tX6oJH7ePPmfrBfqI/O\\nOeH8gMJ+FuwtXYEPa4hBf38M5eU5xWG7\\n-----END CERTIFICATE-----\\n",
					"custom_key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAwQHoetcl9+5ikGzV6cMzWtWPJHqXT3wpbEkRU9Yz7lgvddmG\ndtcGbg/1CGZu0jJGkMoppoUo4c3dts3iwqRYmBikUP77wwY2QGmDZw2FvkJCJlKn\nabIRuGvBKwzESIXgKk2016aTP6/dAjEHyo6SeoK8lkIySUvK0fyOVlsiEsCmOpid\ntnKX/a+50GjB79CJH4ER2lLVZnhePFR/zUOyPxZQQ4naHf7yu/b5jhO0f8fwt+py\nFxIXjbEIdZliWRkRMtzrHOJIhrmJ2A1J7iOrirbbwillwjjNVUWPf3IJ3M12S9pE\newooaeO2izNTERcG9HzAacbVRn2Y2SWIyT/18QIDAQABAoIBACbhTYXBZYKmYPCb\nHBR1IBlCQA2nLGf0qRuJNJZg5iEzXows/6tc8YymZkQE7nolapWsQ+upk2y5Xdp/\naxiuprIs9JzkYK8Ox0r+dlwCG1kSW+UAbX0bQ/qUqlsTvU6muVuMP8vZYHxJ3wmb\n+ufRBKztPTQ/rYWaYQcgC0RWI20HTFBMxlTAyNxYNWzX7RKFkGVVyB9RsAtmcc8g\n+j4OdosbfNoJPS0HeIfNpAznDfHKdxDk2Yc1tV6RHBrC1ynyLE9+TaflIAdo2MVv\nKLMLq51GqYKtgJFIlBRPQqKoyXdz3fGvXrTkf/WY9QNq0J1Vk5ERePZ54mN8iZB7\n9lwy/AkCgYEA6FXzosxswaJ2wQLeoYc7ceaweX/SwTvxHgXzRyJIIT0eJWgx13Wo\n/WA3Iziimsjf6qE+SI/8laxPp2A86VMaIt3Z3mJN/CqSVGw8LK2AQst+OwdPyDMu\niacE8lj/IFGC8mwNUAb9CzGU3JpU4PxxGFjS/eMtGeRXCWkK4NE+G08CgYEA1Kp9\nN2JrVlqUz+gAX+LPmE9OEMAS9WQSQsfCHGogIFDGGcNf7+uwBM7GAaSJIP01zcoe\nVAgWdzXCv3FLhsaZoJ6RyLOLay5phbu1iaTr4UNYm5WtYTzMzqh8l1+MFFDl9xDB\nvULuCIIrglM5MeS/qnSg1uMoH2oVPj9TVst/ir8CgYEAxrI7Ws9Zc4Bt70N1As+U\nlySjaEVZCMkqvHJ6TCuVZFfQoE0r0whdLdRLU2PsLFP+q7qaeZQqgBaNSKeVcDYR\n9B+nY/jOmQoPewPVsp/vQTCnE/R81spu0mp0YI6cIheT1Z9zAy322svcc43JaWB7\nmEbeqyLOP4Z4qSOcmghZBSECgYACvR9Xs0DGn+wCsW4vze/2ei77MD4OQvepPIFX\ndFZtlBy5ADcgE9z0cuVB6CiL8DbdK5kwY9pGNr8HUCI03iHkW6Zs+0L0YmihfEVe\nPG19PSzK9CaDdhD9KFZSbLyVFmWfxOt50H7YRTTiPMgjyFpfi5j2q348yVT0tEQS\nfhRqaQKBgAcWPokmJ7EbYQGeMbS7HC8eWO/RyamlnSffdCdSc7ue3zdVJxpAkQ8W\nqu80pEIF6raIQfAf8MXiiZ7auFOSnHQTXUbhCpvDLKi0Mwq3G8Pl07l+2s6dQG6T\nlv6XTQaMyf6n1yjzL+fzDrH3qXMxHMO/b13EePXpDMpY7HQpoLDi\n-----END RSA PRIVATE KEY-----\n",
				},
				"custom_metadata": map[string]string{
					"key": "value",
				},
				"custom_origin_server": "origin2.example.com",
				"custom_origin_sni":    "sni.example.com",
				"status":               "pending",
				"verification_errors": []string{
					"None of the A or AAAA records are owned by this account and the pre-generated ownership verification token was not found.",
				},
				"ownership_verification": map[string]string{
					"type":  "txt",
					"name":  "_cf-custom-hostname.app.example.com",
					"value": "5cc07c04-ea62-4a5a-95f0-419334a875a4",
				},
				"ownership_verification_http": map[string]string{
					"http_url":  "http://custom.test.com/.well-known/cf-custom-hostname-challenge/0d89c70d-ad9f-4843-b99f-6cc0252067e9",
					"http_body": "5cc07c04-ea62-4a5a-95f0-419334a875a4",
				},
				"created_at": "2020-02-06T18:11:23.531995Z",
			},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.GET("/ping", func(c *gin.Context) {
		Logger.Get().Info("Ping")
		response.Send(c, gin.H{
			"all": "ok",
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.GET("/ping/:id", func(c *gin.Context) {
		Logger.Get().Info("GET Ping ID " + c.Param("id"))
		response.Send(c, gin.H{
			"ok": c.Param("id"),
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/ping/:id", func(c *gin.Context) {
		Logger.Get().Info("POST Ping ID " + c.Param("id"))
		response.Send(c, gin.H{
			"ok": c.Param("id"),
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/validate", func(c *gin.Context) {
		Logger.Get().Info("validate")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/suggestions/suggest", func(c *gin.Context) {
		Logger.Get().Info("Suggestions suggest")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/suggestions/fetch", func(c *gin.Context) {
		Logger.Get().Info("Suggestions fetch")
		response.Send(c, gin.H{
			"rules": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/validate/mlid", func(c *gin.Context) {
		Logger.Get().Info("validate mlid")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/sumo/post", func(c *gin.Context) {
		buf := make([]byte, c.Request.ContentLength)
		_, _ = c.Request.Body.Read(buf)
		_ = c.Request.Body.Close()
		Logger.Get().Info("SUMO request received " + string(buf))
		response.Send(c, gin.H{}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.GET("/render/view", func(c *gin.Context) {
		renderMap := map[string]interface{}{
			"render_profile": c.Request.URL.Query().Has("render_profile"),
			"render_name":    c.Request.URL.Query().Has("render_name"),
			"name":           c.Request.URL.Query().Get("name"),
			"alert":          c.Request.URL.Query().Has("alert"),
		}

		wd, _ := os.Getwd()
		t, err := template.ParseFiles(wd + "/resources/views/temp.html")

		if err != nil {
			Logger.Get().Info("Parse error " + err.Error())
			c.String(200, err.Error())
			return
		}

		err = t.Execute(c.Writer, renderMap)
		if err != nil {
			Logger.Get().Info("Template Err " + err.Error())
		}
	})

	engine.GET("/render/view/gin", func(c *gin.Context) {
		renderMap := map[string]interface{}{
			"render_profile": c.Request.URL.Query().Has("render_profile"),
			"render_name":    c.Request.URL.Query().Has("render_name"),
			"name":           c.Request.URL.Query().Get("name"),
		}

		//wd, _ := os.Getwd()
		c.HTML(200, "temp.html", renderMap)
	})

	engine.GET("/api/rzp/v1/internal/account/config", func(c *gin.Context) {
		strRes := `{"id":"10000000000000","name":"Test Account","brand_color":"#2371EC","handle":null,"transaction_report_email":["test@razorpay.com"],"logo_url":null,"invoice_label_field":null,"auto_capture_late_auth":false,"fee_credits_threshold":null,"amount_credits_threshold":null,"refund_credits_threshold":null,"balance_threshold":null,"display_name":null,"default_refund_speed":"normal","fee_bearer":"platform","billing_label":"Test Account","website":"http://abc.com","receipt_email_enabled":true,"parent_id":null,"international":false,"convert_currency":false,"category":"5691","logo_large_size_url":null,"rs_filter":true,"max_payment_amount":50000000,"max_international_payment_amount":50000000,"is_suspended":false,"is_live":true,"org_id":"100000razorpay","org_custom_code":"rzp","branding_logo":null,"login_logo":"","country_code":"IN","currency_code":"INR","org_name":"Razorpay Software Private Ltd","policy_url":null,"time_zone":"Asia/Kolkata","custom_org_branding":false,"methods":{"entity":"methods","upi":true},"mcc":"5691"}`
		res := map[string]interface{}{}
		json.Unmarshal([]byte(strRes), &res)
		c.JSON(http.StatusOK, res)
	})

	engine.POST("/api/rzp/v1/merchant/analytics", func(c *gin.Context) {
		strRes := `{"current":{"total":1,"last_updated_at":1707660539,"result":[{"value":1100}]},"previous":{"total":1,"last_updated_at":1707660539,"result":[{"value":1000}]},"current_histogram":{"total":7,"last_updated_at":1707660539,"result":[{"timestamp":1707157800,"value":100},{"timestamp":1705257000,"value":200},{"timestamp":1705602600,"value":100},{"timestamp":1706034600,"value":300},{"timestamp":1706725800,"value":100},{"timestamp":1706466600,"value":200},{"timestamp":1705516200,"value":100}]},"previous_histogram":{"last_updated_at":1707660539,"result":[{"timestamp":1707157800,"value":100},{"timestamp":1705257000,"value":200},{"timestamp":1705602600,"value":100},{"timestamp":1706034600,"value":200},{"timestamp":1706725800,"value":100},{"timestamp":1706466600,"value":200},{"timestamp":1705516200,"value":100}]}}`
		res := map[string]interface{}{}
		json.Unmarshal([]byte(strRes), &res)
		c.JSON(http.StatusOK, res)
	})

	engine.POST("/care/rzp/twirp/rzp.care.dashboard.home.v1.HomeService/GetMerchantKeyUpdate", func(c *gin.Context) {
		strRes := `{"data":[{"title":"2 Payment failure occured","message":"Review recent payment failures due to customer, bank, or business related issues","action":"payment_failed","params":{"from":"1706207400","to":"1706898599","status":"failed"},"type":"failed"},{"title":"1 Payment failure occured","message":"Review recent payment failures due to customer, bank, or business related issues","action":"payment_details","params":{"payment_id":"pay_Ixasdadadas"},"type":"failed"},{"title":"International payment enabled","message":"Accept payments via payment gateway, payment pages, payment links, and invoices","action":"international","params":{},"type":"info"},{"title":"Action required on your request to increase transaction limit","message":"Need more information","action":"international","params":{},"type":"need_clarification"}]}`
		res := map[string]interface{}{}
		json.Unmarshal([]byte(strRes), &res)
		c.JSON(http.StatusOK, res)
	})

	engine.POST("/care/rzp/twirp/rzp.care.dashboard.home.v1.HomeService/GetMerchantSettlementOverview", func(c *gin.Context) {
		strRes := `{"is_transacted":true,"is_settlement":true,"settlement_schedule":"FLBYKXvweeGWrV","settlement":{"current_balance":"60000","current_balance_currency":"INR","settlement_currency":"INR","today":{"total_amount":"1002","total_count":"9","delayed_total_amount":"800","delayed_count":"5","failed_total_amount":"0","failed_count":"0","created_total_amount":"0","created_count":"0","processed_total_amount":"202","processed_count":"4","title_key":"SETL_TODAY_MULTIPLE"},"previous":{"total_amount":"0","total_count":"1","created_at":"1659586909"},"upcoming_settlement":{"title_key":"UPCOMING_SETL_SKIPPED_NO_NEXT_SETTLEMENT","settlement_amount":"0","next_settlement_time":"0"}}}`
		res := map[string]interface{}{}
		json.Unmarshal([]byte(strRes), &res)
		c.JSON(http.StatusOK, res)
	})

	engine.GET("/dashboard/rzp/user/session", func(c *gin.Context) {
		strRes := `{"success":true,"data":{"token":"somerandomstringof36characterslength","email":"example@example.com","name":"John Doe","merchant_id":"10000000000000","role":"owner","merchant_name":"Merchant Name","logo":null,"user_id":"IY9DXiTMu25fDc"}}`
		res := map[string]interface{}{}
		json.Unmarshal([]byte(strRes), &res)
		c.JSON(http.StatusOK, res)
	})

	engine.NoRoute(func(c *gin.Context) {
		Logger.Get().Info("Route not found")
		response.Send(c, gin.H{}, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})
}
