package jpushclient

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}

type AndroidNotice struct {
	Alert     string                 `json:"alert"`
	Title     string                 `json:"title,omitempty"`
	BuilderId int                    `json:"builder_id,omitempty"`
	Intent    map[string]string      `json:"intent,使用intent里的url指定点击通知栏后跳转的目标页面,可选"`
	Extras    map[string]interface{} `json:"extras,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	this.IOS = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}
