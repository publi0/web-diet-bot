package client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthWebDiet struct {
	N string `json:"id_nutricionista"`
	P string `json:"id_paciente"`
}

type API interface {
	UploadPhoto(photo, title, description, n, p string)
	GetUserAuth(phone, birthYear string) (*AuthWebDiet, error)
}

type WebDiet struct {
	Client *http.Client
}

func (w WebDiet) UploadPhoto(photo, title, description, n, p string) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().In(location)

	data := url.Values{}
	data.Set("n", n)
	data.Set("p", p)
	data.Set("tipo", "diario")
	data.Set("data", now.Format("02/01/2006"))
	data.Set("horario", now.Format("15:04"))
	data.Set("refeicao", title)
	data.Set("comentario", description)
	data.Set("arquivo", photo)

	req, _ := http.NewRequest("POST", "https://pt.webdiet.com.br/api/app/uploadArquivo_webdiet3.php", strings.NewReader(data.Encode()))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Host", "pt.webdiet.com.br")
	req.Header.Set("Origin", "null")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)")
	resp, err := w.Client.Do(req)
	if err != nil {
		log.Println("Error calling webdiet")
		log.Println(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != http.StatusOK || string(bodyText) != "ok" {
		log.Println(bodyText)
		log.Fatal(err)
	}
	log.Println("Success sending to webdiet")
}

func (w WebDiet) GetUserAuth(phone, birthYear string) (*AuthWebDiet, error) {
	data := url.Values{}
	data.Set("tefone", phone)
	data.Set("nascimento", birthYear)

	req, _ := http.NewRequest("POST", "https://webdiet.com.br/api/v3/login_v4.php", strings.NewReader(data.Encode()))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Host", "webdiet.com.br")
	req.Header.Set("Origin", "null")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)")

	resp, err := w.Client.Do(req)
	if err != nil {
		log.Println("Error calling webdiet")
		log.Println(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Error with user information")
		return nil, errors.New("invalid user information")
	}
	var auth AuthWebDiet

	json.Unmarshal(bodyText, &auth)

	log.Println("Success getting auth from webdiet")
	return &auth, nil
}
