package main

import (
	"encoding/json"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	clientIp := r.RemoteAddr

	td, err := CreateToken(userId, clientIp)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	err = SaveRefreshToken(userId, td.RefreshToken)
	if err != nil {
		http.Error(w, "Could not save refresh token", http.StatusInternalServerError)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	json.NewEncoder(w).Encode(tokens)
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	json.NewDecoder(r.Body).Decode(&body)

	userId := body["user_id"]
	refreshToken := body["refresh_token"]
	clientIp := r.RemoteAddr

	isValid, err := ValidateRefreshToken(userId, refreshToken)
	if err != nil || !isValid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	td, err := CreateToken(userId, clientIp)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Check if IP has changed
	// Example: send warning email (mock)

	err = SaveRefreshToken(userId, td.RefreshToken)
	if err != nil {
		http.Error(w, "Could not save refresh token", http.StatusInternalServerError)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	json.NewEncoder(w).Encode(tokens)
}
