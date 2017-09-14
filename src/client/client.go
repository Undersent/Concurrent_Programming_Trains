
package client

import (
    "fmt"
    "configs"
    "train"
    "myswitch"
    "track"
    "os"
    "os/exec"
)

func clearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}

func Talk() {
    var cmd int
    for {
        fmt.Println("funkcje:")
        fmt.Println("1    konfiguracja")
        fmt.Println("2    pociagi")
        fmt.Println("3    tory")
        fmt.Println("4    zwrotnice")
        fmt.Println("5    pozycje pociagow")
        fmt.Println("6    wyjscie")

        _, _ = fmt.Scanf("%d", &cmd)

        /* Clear screen */
        clearScreen()

        switch cmd {
        case 1:
            configs.Conf.Show()
        case 2:
            train.Trains.Show()
        case 3:
            track.Tracks.Show()
        case 4:
            myswitch.Switches.Show()
        case 5:
            train.Trains.ShowPos()
        case 6:
            os.Exit(0)
        }
    }
}
