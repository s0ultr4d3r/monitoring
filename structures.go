package main

type MonAnsw struct {
	States []struct {
		Pk                     int           `json:"pk"`
		Alias                  string        `json:"alias"`
		Level                  string        `json:"level"`
		DependingCount         int           `json:"depending_count"`
		IsHidden               bool          `json:"is_hidden"`
		DirectlyResponsible    []interface{} `json:"directly_responsible"`
		DutyHightlightDisabled bool          `json:"duty_hightlight_disabled"`
		SubHostHostname        string        `json:"sub__host__hostname"`
		SubHostID              int           `json:"sub__host_id"`
		SubHostInferedZoneName interface{}   `json:"sub__host__infered_zone__name"`
		SubSubservice          string        `json:"sub__subservice"`
		SubID                  int           `json:"sub_id"`
		SubServiceService      string        `json:"sub__service__service"`
		SubServiceID           int           `json:"sub__service_id"`
		SubSubserviceStr       string        `json:"sub__subservice_str"`
		MonhostHostname        string        `json:"monhost__hostname"`
		ChecktimeHumanized     string        `json:"checktime_humanized"`
		Checktime              string        `json:"checktime"`
		Times                  []struct {
			Label             string `json:"label"`
			Timestamp         string `json:"timestamp"`
			DurationHumanized string `json:"duration_humanized"`
		} `json:"times"`
		TotalTime struct {
			Timestamp         string `json:"timestamp"`
			DurationHumanized string `json:"duration_humanized"`
		} `json:"total_time"`
		RackName                  interface{}   `json:"rack__name"`
		DatacenterAbbr            interface{}   `json:"datacenter__abbr"`
		Color                     string        `json:"color"`
		MsgDecoded                string        `json:"msg_decoded"`
		SmallMsg                  string        `json:"small_msg"`
		AlertCount                int           `json:"alert_count"`
		OkAlertCount              int           `json:"ok_alert_count"`
		State                     string        `json:"state"`
		Status                    int           `json:"status"`
		SubStatus                 int           `json:"sub__status"`
		SubAssignAuthorUsername   interface{}   `json:"sub__assign_author__username"`
		SubAssignedAt             string        `json:"sub__assigned_at"`
		SubTimeLeftHumanized      string        `json:"sub__time_left_humanized"`
		SubExecutorPk             interface{}   `json:"sub__executor__pk"`
		SubExecutorUsername       interface{}   `json:"sub__executor__username"`
		SubComment                string        `json:"sub__comment"`
		SubTasks                  interface{}   `json:"sub__tasks"`
		SubTasksCount             int           `json:"sub__tasks__count"`
		Responsibles              []interface{} `json:"responsibles"`
		Zones                     []interface{} `json:"zones"`
		HasChat                   bool          `json:"has_chat"`
		IncidentID                int           `json:"incident_id"`
		ChatNotifiedAt            string        `json:"chat_notified_at"`
		DependsOnID               interface{}   `json:"depends_on_id"`
		Maintenances              []interface{} `json:"maintenances"`
		MonitoringManagerNotified bool          `json:"monitoring_manager_notified"`
		IsCalled                  interface{}   `json:"is_called"`
		IsCalledLately            interface{}   `json:"is_called_lately"`
		CheckerID                 int           `json:"checker_id"`
		IsNotHardstateIn          bool          `json:"is_not_hardstate_in"`
		Group                     interface{}   `json:"group"`
		ExternalLink              string        `json:"external_link"`
	} `json:"states"`
}

type Prblm struct {
	Hostname  string
	Tag       string
	Problem   string
	MetaSrv   string
	StartTime string
	Id        uint64
}
