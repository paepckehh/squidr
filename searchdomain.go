package squidr

import (
	"math/rand"
	"time"
)

func searchDomain(domain string) search {
	for _, item := range blocklist {
		if domain == item {
			return search{_blockDomain, true, false, false, true}
		}
	}
	for target, searchlist := range rewrite {
		for _, item := range searchlist {
			if domain == item {
				return search{target, false, true, false, true}
			}
		}
	}
	for target, searchlist := range redirect302 {
		for _, item := range searchlist {
			if domain == item {
				return search{target, false, false, false, true}
			}
		}
	}
	for target, searchlist := range redirect301 {
		for _, item := range searchlist {
			if domain == item {
				return search{target, false, false, true, true}
			}
		}
	}
	for target, searchlist := range redirect301randomMap {
		for _, item := range searchlist {
			if domain == item {
				targetMap, ok := randomMap[target]
				if !ok {
					return search{domain, false, false, false, false} // map does not exist,
				}
				l := len(targetMap.targets)
				if l < 1 {
					return search{domain, false, false, false, false} // not map targets
				}
				rand.Seed(time.Now().UnixNano())
				idx := rand.Intn(l - 1)
				if targetMap.doNotValidate {
					return search{targetMap.targets[idx], false, false, true, true}
				}
				for i := 0; i < l*3; i++ {
					target := targetMap.targets[idx]
					if isReachableViaProxy(targetMap, idx) {
						return search{target, false, false, true, true} // random picked target is reachable
					}
					idx = rand.Intn(l - 1) // next
				}
				for i := 0; i < l; i++ {
					if len(targetMap.targets[i]) > 0 {
						return search{domain, false, false, false, false} // one valid target
					}
				}
				delete(randomMap, target) // no valid target urls left, remove mapEntry
			}
		}
	}
	return search{domain, false, false, false, false}
}
