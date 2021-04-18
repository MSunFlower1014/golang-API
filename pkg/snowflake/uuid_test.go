package snowflake

import "testing"

func TestGetUID(t *testing.T) {
	id := GetUID()
	t.Logf("id is %v", id)
}
