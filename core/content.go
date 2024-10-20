package core

import (
	"os"

	"github.com/buger/jsonparser"
)

func (s *Server) update() (bool, error) {
	body, err := get(versionListURL)
	if err != nil {
		return false, err
	}

	newVersion, err := jsonparser.GetString(body, "latest", "release")
	if err != nil {
		return false, err
	}

	if newVersion == s.version {
		return false, nil
	}

	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		id, _ := jsonparser.GetString(value, "id")
		if id == newVersion {
			url, _ := jsonparser.GetString(value, "url")
			data, err := get(url)
			if err != nil {
				return
			}

			serverUri, err := jsonparser.GetString(data, "downloads", "server", "url")
			if err != nil {
				return
			}

			server, err := get(serverUri)
			if err != nil {
				return
			}

			if err = os.WriteFile(s.serverName+"/server.jar", server, 0775); err != nil {
				return
			}
		}
	}, "versions")

	return true, nil
}

func (s *Server) makeBackup() (err error) {
	return
}
