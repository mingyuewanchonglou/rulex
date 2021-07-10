package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"rulex/plugin/http_server"
	"rulex/x"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/ngaut/log"
)

//
var engine *x.RuleEngine

//
func main() {
	gin.SetMode(gin.ReleaseMode)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT)
	engine = x.NewRuleEngine()
	////////
	hh := httpserver.NewHttpApiServer(2580, "plugin/http_server/templates/*")
	if e := engine.LoadPlugin(hh); e != nil {
		log.Fatal("rule load failed:", e)
	}
	for _, minEnd := range hh.AllMInEnd() {
		config := map[string]interface{}{}
		if err := json.Unmarshal([]byte(minEnd.Config), &config); err != nil {
			log.Fatal(err)
		}
		in1 := x.NewInEnd(minEnd.Type, minEnd.Name, minEnd.Description, &config)
		if err := engine.LoadInEnd(in1); err != nil {
			log.Fatal("InEnd load failed:", err)
		}
	}
	// in1 := x.NewInEnd("MQTT", "MQTT Stream", "MQTT Input Stream", &map[string]interface{}{
	// 	"server":   "127.0.0.1",
	// 	"port":     1883,
	// 	"username": "test",
	// 	"password": "test",
	// 	"clientId": "test",
	// })
	// if err0 := engine.LoadInEnd(in1); err0 != nil {
	// 	log.Fatal("InEnd load failed:", err0)
	// }
	// in2 := x.NewInEnd("COAP", "COAP Stream", "COAP Input Stream", &map[string]interface{}{
	// 	"server": "127.0.0.1",
	// 	"port":   1883,
	// })

	// if err := engine.LoadInEnd(in2); err != nil {
	// 	log.Fatal("InEnd load failed:", err)
	// }

	// 	out1 := x.NewOutEnd("mongo", "Data to mongodb", "Save data to mongodb",
	// 		&map[string]interface{}{
	// 			"mongourl": "mongodb+srv://rulex:rulex@cluster0.rsdmb.mongodb.net/test",
	// 		})
	// 	out1.Id = "MongoDB001"
	// 	if err1 := engine.LoadOutEnds(out1); err1 != nil {
	// 		log.Fatal("OutEnd load failed:", err1)
	// 	}
	// 	actions := `
	// local json = require("json")
	// Actions = {
	// 	function(data)
	// 	    dataToMongo("MongoDB001", data)
	// 	    print("[LUA Actions Callback]:dataToMongo Mqtt payload:", data)
	// 		return true, data
	// 	end
	// }
	// `
	// 	from := []string{"test"}
	// 	failed := `
	// function Failed(error)
	//   -- print("[LUA Callback] call failed from lua:", error)
	// end
	// `
	// 	success := `
	// function Success()
	//   -- print("[LUA Callback] call success from lua")
	// end
	// `
	// 	rule1 := x.NewRule(engine, "just_a_test_rule", "just_a_test_rule", from, success, actions, failed)
	// 	rule1.Id = "just_a_test_rule"

	// 	//
	// 	if e := engine.LoadRule(rule1); e != nil {
	// 		log.Fatal("rule load failed:", e)
	// 	}
	engine.Start()
	<-c
	os.Exit(0)
}
