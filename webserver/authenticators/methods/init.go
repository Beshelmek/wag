package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/NHAS/wag/config"
	"github.com/NHAS/wag/webserver/authenticators"
)

// from: https://github.com/duo-labs/webauthn.io/blob/3f03b482d21476f6b9fb82b2bf1458ff61a61d41/server/response.go#L15
func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func init() {
	authenticators.MFA[authenticators.TotpMFA] = new(Totp)
	authenticators.MFA[authenticators.WebauthnMFA] = new(Webauthn)
}

func resultMessage(err error) (string, int) {
	if err == nil {
		return "Success", http.StatusOK
	}

	msg := "Validation failed"
	if strings.Contains(err.Error(), "locked") {
		msg = "Account is locked contact: " + config.Values().HelpMail
	}
	return msg, http.StatusBadRequest
}
