package test

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 奖项结构
type Prize struct {
	Name  string
	Stock int
}

// 抽奖算法
func drawLots(participants []string, prizes []Prize) (map[string]string, map[string]bool, error) {
	totalParticipants := len(participants)
	if totalParticipants == 0 {
		return nil, nil, fmt.Errorf("no participants")
	}

	winners := make(map[string]string)
	indices := make(map[int]struct{})

	for _, prize := range prizes {
		if prize.Stock <= 0 {
			continue
		}

		for i := 0; i < prize.Stock; i++ {
			if len(indices) >= totalParticipants {
				break
			}

			var index int64
			for {
				randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(totalParticipants)))
				if err != nil {
					return nil, nil, fmt.Errorf("error generating random number: %v", err)
				}

				index = randIndex.Int64()
				if _, exists := indices[int(index)]; !exists {
					break
				}
			}

			indices[int(index)] = struct{}{}
			winners[participants[index]] = prize.Name
		}
	}

	// 标记未中奖的用户
	nonWinners := make(map[string]bool)
	for _, participant := range participants {
		if _, won := winners[participant]; !won {
			nonWinners[participant] = true
		}
	}

	return winners, nonWinners, nil
}
