
package driver

import (
    "track"
    "train"
    "configs"
    "fmt"
    "time"
    "myswitch"
    "graph"
    "repairteam"
)

// Driver - train driver
type Driver struct {
    train   *train.Train
}

// New - create New Driver
func New(t *train.Train) *Driver {
    d := new(Driver)
    d.train = t

    return d
}

// Drive - main drover function
func (d *Driver) Drive() {
    for {
        var wTime float64
        var speed int
        var newSwitchID int
        var curSwitchID int

        curSwitchID = d.train.StartPoint()
        route := d.train.Route()

        /* we go to the station */
        d.train.ChangePos(train.POS_STATION, route[0])
        firstTime := true

        for i := 0; i < len(route) - 1 && route[i] != 0; i++ {
            /* get new Track */
            tr := track.Tracks.GetTrackByID(route[i])

            /* set track as busy */
            tr.Busy()

            /* if is a station wait for people */
            if tr.Type() == track.STATION {
                d.train.ChangePos(train.POS_STATION, route[i])

                /* print info */
                if configs.Conf.Mode() == configs.NOISY {
                    fmt.Printf("nr %d     jedzie na stacje: %d\n", d.train.ID(), tr.ID())
                }

                wTime = float64(tr.HTime()) * float64(configs.Conf.SPerH())
            } else { /* it's normal track, so let's go */
                d.train.ChangePos(train.POS_TRACK, route[i])

                /* print info */
                if configs.Conf.Mode() == configs.NOISY {
                    fmt.Printf("nr %d     jedzie na tor: %d\n", d.train.ID(), tr.ID())
                }

                /* our speed is min speed of train speed and track speed */
                if tr.Speed() < d.train.MaxSpeed() {
                    speed = tr.Speed()
                } else {
                    speed = d.train.MaxSpeed()
                }

                wTime = float64(tr.Len()) / float64(speed) * float64(configs.Conf.SPerH())
            }

            /* time for driving or waiting for people */
            time.Sleep(time.Duration(wTime *float64(time.Second)))

            /* Free Track */
            tr.Free()
            repairteam.UseItem(graph.NewNode(graph.EDGE, tr.ID()))

            /* Print info */
            if configs.Conf.Mode() == configs.NOISY {
                if tr.Type() == track.NORMAL {
                    fmt.Printf("nr %d     opuszcza tor %d\n", d.train.ID(), tr.ID())
                } else {
                    fmt.Printf("nr %d     opuszcza stacje %d\n", d.train.ID(), tr.ID())
                }
            }

            vers := tr.Vers()

            /* choose Switch (start or end) */
            if firstTime == false {
                for j := 0; j < len(vers); j++ {
                    if vers[j] != 0 {
                        newSwitchID = vers[j]

                        if newSwitchID != curSwitchID {
                            curSwitchID = newSwitchID
                            break
                        }
                    }
                }
            }

            firstTime = false

            /* enter switch */
            s := myswitch.Switches.GetSwitchByID(curSwitchID)

            s.Busy()

            d.train.ChangePos(train.POS_SWITCH, s.ID())

            /* print info */
            if configs.Conf.Mode() == configs.NOISY {
                fmt.Printf("nr %d     jedzie na zwrotnice: %d\n", d.train.ID(), s.ID())
            }

            /* wait */
            wTime = float64(s.StayTime()) * float64(configs.Conf.SPerH())
            time.Sleep(time.Duration(wTime *float64(time.Second)))

            /* go out */
            s.Free()

            repairteam.UseItem(graph.NewNode(graph.VERTEX, s.ID()))

            /* print info */
            if configs.Conf.Mode() == configs.NOISY {
                fmt.Printf("nr %d     opuszcza zwrotnice: %d\n", d.train.ID(), s.ID())
            }
        }
    }
}
