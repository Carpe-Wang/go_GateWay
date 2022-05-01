package load_balance

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}

type WeightNode struct {
	addr            string
	Weight          int //初始化时对节点约定的权重
	currentWeight   int //节点临时权重，每轮都会变化
	effectiveWeight int //有效权重, 默认与weight相同 , totalWeight = sum(effectiveWeight)  //出现故障就-1
}

//1, currentWeight = currentWeight + effectiveWeight
//2, 选中最大的currentWeight节点为选中节点
//3, currentWeight = currentWeight - totalWeight

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{
		addr:   params[0],
		Weight: int(parInt),
	}
	node.effectiveWeight = node.Weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	var best *WeightNode
	total := 0
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		//1 计算所有有效权重
		total += w.effectiveWeight
		//2 修改当前节点临时权重
		w.currentWeight += w.effectiveWeight
		//3 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1，直到恢复到weight大小
		if w.effectiveWeight < w.Weight {
			w.effectiveWeight++
		}

		//4 选中最大临时权重节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}

	if best == nil {
		return ""
	}
	//5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get(string) (string, error) {
	return r.Next(), nil
}

func (r *WeightRoundRobinBalance) Update() {

}
