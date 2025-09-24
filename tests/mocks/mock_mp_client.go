package mocks

import (
    "github.com/arthu/shop-api-go/internal/mp"
)

type MockMPClient struct {
    GetPaymentFunc       func(accessToken, paymentID string) (string, error)
    CreatePreferenceFunc func(accessToken string, pref mp.PreferenceRequest) (string, error)
}

func (m MockMPClient) GetPayment(accessToken, paymentID string) (string, error) {
    if m.GetPaymentFunc != nil {
        return m.GetPaymentFunc(accessToken, paymentID)
    }
    return "", nil
}

func (m MockMPClient) CreatePreference(accessToken string, pref mp.PreferenceRequest) (string, error) {
    if m.CreatePreferenceFunc != nil {
        return m.CreatePreferenceFunc(accessToken, pref)
    }
    return "", nil
}
