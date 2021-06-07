// this is rank tools
package rank

import (
	"log"
	"sort"
)

var ()

const ()

type RankScore uint32
type RankId uint32

type RankValue struct {
	Id    RankId    // id
	Score RankScore // score
}

type RankMgr struct {
	Ranks   []*RankValue
	RankMap map[RankId]*RankValue
}

func (r *RankMgr) Init() {
	log.Println("init rank.")
	r.Ranks = make([]*RankValue, 0)
	r.RankMap = make(map[RankId]*RankValue, 0)
}

// GetRankIndex get rank by id
func (r *RankMgr) GetRankIndex(id RankId) int {
	res := -1
	if _, ok := r.RankMap[id]; ok {

		for i, rank := range r.Ranks {
			if rank.Id == id {
				res = i
				break
			}

		}
	}
	return res
}

// push new element in ranks
func (r *RankMgr) PushRank(id RankId, score RankScore) {
	if rank, ok := r.RankMap[id]; ok {
		rank.Score = score
	} else {

		r.RankMap[id] = &RankValue{
			Id:    id,
			Score: score,
		}
		r.Ranks = append(r.Ranks, r.RankMap[id])
	}
    r.sort()
	log.Println("push new element ", id, ":", score)
}

// sort ranks
func (r *RankMgr) sort() {
	sort.Slice(r.Ranks, func(i, j int) bool {
		return r.Ranks[i].Score > r.Ranks[j].Score
	})
}

// print ranks
func (r *RankMgr) Print() {
	log.Println(r.Ranks)
}
