package bus

import "testing"

func TestGetBusArrivalDetail(t *testing.T) {
	resp, err := GetBusArrivalDetail("548路", 18, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Time)
}
