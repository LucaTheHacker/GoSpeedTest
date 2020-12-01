package GoSpeedTest

import (
	"encoding/json"
	"os/exec"
	"strconv"
)

// SpeedTest runs a speedtest using Ookla.
// SpeedTest from Ookla needs to be installed in order to run this command. (https://www.speedtest.net/apps/cli)
func SpeedTest(serverID int) (result *Result, err error) {
	if serverID == -1 { serverID = 4302 }
	cmd := exec.Command("speedtest", "--server-id=" + strconv.Itoa(serverID), "--format=json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		return
	}
	return
}

// GetServers returns a list of servers recommended (bu Ookla) to run a Speedtest.
func GetServers() (*[]Server, error) {
	cmd := exec.Command("speedtest", "-L", "--format=json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return &[]Server{}, err
	}

	var result Result
	err = json.Unmarshal(out, &result)
	if err != nil {
		return &[]Server{}, err
	}
	return result.Servers, nil
}
