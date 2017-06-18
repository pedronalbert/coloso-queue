package urid_test

import "testing"
import "coloso-queue/utils/urid"

func TestGenerate(t *testing.T) {
  var expectedURID = "LAN_12345"
  var URID = urid.Generate("lan", 12345)

  if URID != expectedURID {
    t.Fatalf("expected URID: %s got URID: %s", expectedURID, URID)
  }
}

func TestGetRegion(t *testing.T) {
  var expectedRegion = "LAN"
  var region = urid.GetRegion("LAN_123456")

  if expectedRegion != region {
    t.Fatalf("exptected region: %s got region: %s", expectedRegion, region)
  }
}

func TestGetId(t *testing.T) {
  var exptectedID = "123456"
  var ID = urid.GetID("LAN_123456")

  if exptectedID != ID {
    t.Fatalf("exptected ID: %s got ID: %s", exptectedID, ID)
  }
}
