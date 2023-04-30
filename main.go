package gif

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	Gif(os.Stdout)
}
