package benchmark

import "testing"

func Test_visitWebsite(t *testing.T) {
	wg.Add(1)

	go visitWebsite("https://www.google.com")
	wg.Wait()
	if totalDuration == 0 {
		t.Fail()
	}
}

func Test_Start(t *testing.T) {
	recordTotalDuration := totalDuration
	Start("https://www.facebook.com", 15)

	newTotalDuration := totalDuration
	if recordTotalDuration == newTotalDuration {
		t.Fail()
	}
}
