package mp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PreferenceItem struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Quantity   int     `json:"quantity"`
	CurrencyID string  `json:"currency_id"`
	UnitPrice  float64 `json:"unit_price"`
}

func GetPayment(accessToken, paymentID string) (string, error) {
    url := fmt.Sprintf("https://api.mercadopago.com/v1/payments/%s", paymentID)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
    resp, err := http.DefaultClient.Do(req)
    if err != nil { return "", err }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 { return "", fmt.Errorf("mp get payment status %d", resp.StatusCode) }
    var body struct{ Status string `json:"status"` }
    if err := json.NewDecoder(resp.Body).Decode(&body); err != nil { return "", err }
    return body.Status, nil
}

type PreferenceRequest struct {
	Items          []PreferenceItem `json:"items"`
	NotificationURL string          `json:"notification_url"`
}

type PreferenceResponse struct {
	InitPoint string `json:"init_point"`
}

func CreatePreference(accessToken string, pref PreferenceRequest) (string, error) {
	b, _ := json.Marshal(pref)
	req, _ := http.NewRequest("POST", "https://api.mercadopago.com/checkout/preferences", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return "", err }
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 { return "", fmt.Errorf("mp create preference status %d", resp.StatusCode) }
	var pr struct{ Body PreferenceResponse `json:"body"` }
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&pr); err != nil { return "", err }
	if pr.Body.InitPoint == "" {
		// fallback: some responses return the fields at root
		resp.Body.Close()
		// decode root
		resp2, _ := http.DefaultClient.Do(req)
		defer resp2.Body.Close()
		var root PreferenceResponse
		if err := json.NewDecoder(resp2.Body).Decode(&root); err == nil && root.InitPoint != "" {
			return root.InitPoint, nil
		}
	}
	return pr.Body.InitPoint, nil
}
