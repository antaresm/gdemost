package handlers

import (
	"encoding/json"
	"gdemost/database"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var BridgeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var payload []byte
	id, err := strconv.Atoi(vars["id"])
	if err == nil {
		bridge := database.DataInstance.GetBridge(uint(id))
		payload, _ = json.Marshal(bridge)
	} else {
		bridges := database.DataInstance.GetBridges()
		payload, _ = json.Marshal(bridges)
	}
	println(id)

	w.Write([]byte(payload))

})

