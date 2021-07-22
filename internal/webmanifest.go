package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"git.mills.io/yarnsocial/yarn"
)

type Icon struct {
	Src   string `json:"src"`
	Sizes string `json:"sizes"`
	Type  string `json:"type"`
}

type WebManifest struct {
	Name            string `json:"name"`
	ShortName       string `json:"short_name"`
	Icons           []Icon
	ThemeColor      string `json:"theme_color"`
	BackgroundColor string `json:"background_color"`
	Display         string `json:"display"`
}

func NewWebManifest(config *Config) WebManifest {
	var (
		prefix string
		icons  []Icon
	)

	if config.Debug {
		prefix = "/img"
	} else {
		prefix = fmt.Sprintf("/img/%s", yarn.Commit)
	}

	icons = []Icon{
		{
			Src:   prefix + "/android-chrome-192x192.png",
			Sizes: "192x192",
			Type:  "image/png",
		}, {
			Src:   prefix + "/android-chrome-512x512.png",
			Sizes: "512x512",
			Type:  "image/png",
		},
	}

	return WebManifest{
		Name:            config.Name,
		ShortName:       config.Name,
		Icons:           icons,
		ThemeColor:      "#ffffff",
		BackgroundColor: "#ffffff",
		Display:         "standalone",
	}
}

// WebManifestHandler ...
func (s *Server) WebManifestHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data, err := json.Marshal(NewWebManifest(s.config))
		if err != nil {
			log.WithError(err).Errorf("error serializing webmanifest")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

		if r.Method == http.MethodHead {
			return
		}

		_, _ = w.Write([]byte(data))
	}
}
