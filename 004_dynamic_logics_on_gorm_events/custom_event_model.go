package main

import "github.com/jinzhu/gorm"

var diffScriptForPerClientID = []CustomEventModel {
	CustomEventModel{
		Model : gorm.Model{
			ID: 1,
		},
		ModelName: "users",
		ClientID:   "A",
		HookName: "AfterCreate",
		Content: `function main(user) {
		console.log("I Am From Client ID = "+Runtime.ClientID + ",And My Name Is " + user.Name)
	}`,
		Enabled:  true,
	},
	CustomEventModel{
		Model : gorm.Model{
			ID: 2,
		},
		ModelName: "users",
		ClientID:   "B",
		HookName: "AfterCreate",
		Content: `function main(user) {
		console.log("I Am From Client ID = "+Runtime.ClientID + ",And My Name Is " + user.Name)
	}`,
		Enabled:  true,
	},
}



type CustomEventModel struct {
	gorm.Model
	ModelName string
	HookName string
	ClientID string
	Content string
	Enabled bool
}


func GetCustomEvent(modelName string,hookName string,clientID string) *CustomEventModel {
	for _,script := range diffScriptForPerClientID {
		if script.ModelName == modelName && script.HookName == hookName && script.ClientID == clientID {
			return &script
		}
	}
	return nil
}










