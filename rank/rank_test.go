package rank

import (
    "log"
    "math/rand"
    "testing"
)

func TestGetRank(t *testing.T) {
    rank := new(RankMgr)
    rank.Init()
    rank.PushRank(RankId(1), RankScore(rand.Uint32()))
    rank.PushRank(RankId(2), RankScore(rand.Uint32()))
    rank.PushRank(RankId(3), RankScore(rand.Uint32()))
    rank.Print()
    log.Println(rank.GetRankIndex(RankId(1)))
}
