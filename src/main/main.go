package main
import (
	"configs"
	"driver"
	"graph"
	"myswitch"
	"repairteam"
	"sync"
	"track"
	"train"
)

func main()

	var wg sync.WaitGroup
	/* Load configs from file */
	configs.Conf.Load()

	/* Create Global Arrays */
	train.Trains.NewTrains(configs.Conf.NumTrains())
	myswitch.Switches.NewSwitches(configs.Conf.NumSwitches())
	track.Tracks.NewTracks(configs.Conf.NumTracks())

	/* Load ALL */
	train.Trains.Load()
	myswitch.Switches.Load()
	track.Tracks.Load()
	graph.Load()

	repairteam.RepairNodeStation = graph.NewNode(graph.EDGE, configs.Conf.NumTracks())
	train.Trains.GetTrainByID(configs.Conf.NumTrains()).ChangePos(train.POS_STATION, repairteam.RepairNodeStation.ID())

	/* Start Thread if needed */
	if configs.Conf.Mode() == configs.SILENT {
		wg.Add(1)

		go client.Talk()
	}

	drivers := make([]*driver.Driver, configs.Conf.NumTrains()-1)
	for i := 0; i < len(drivers); i++ {
		drivers[i] = driver.New(train.Trains.GetTrainByID(i + 1))
		wg.Add(1)
		go drivers[i].Drive()
	}

	wg.Wait()
}
