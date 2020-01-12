package main

import (
	"gopkg.in/yaml.v2"
)

func kubeweekly(Session Github)(){
	ConfigTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/kubeweekly/ContentList.yaml")
	
	var PushRepository = false
    var Config KubeweeklyContentList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.ContentLists {
		if s.Status.Delivered == false {

			YamlTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/kubeweekly/"+s.Content)

			var Content KubeweeklyContent
			yaml.Unmarshal(YamlTmpl, &Content)
			
			SendTelegram(Content)

			Config.ContentLists[i].Status.Delivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		Data, _ := yaml.Marshal(Config)
		Session.UpdateFile("zufardhiyaulhaq","announcer-system","./resources/kubeweekly/ContentList.yaml",Data)
	}
}
