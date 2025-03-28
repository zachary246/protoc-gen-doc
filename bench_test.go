package gendoc_test

import (
	"testing"

	"github.com/pseudomuto/protokit/utils"
	. "github.com/zachary246/protoc-gen-doc"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	set, _ := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	plugin := new(Plugin)

	for i := 0; i < b.N; i++ {
		plugin.Generate(req)
	}
}
