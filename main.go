package main

import (
	"fmt"
	"math/rand"
)

type Bandit struct {
	arms  int
	rates []float64
}

func NewBandit(arms int) *Bandit {
	var rates []float64
	for i := 0; i < arms; i++ {
		rates = append(rates, rand.Float64())
	}
	return &Bandit{
		arms:  arms,
		rates: rates,
	}
}

func (b *Bandit) play(arm int) float64 {
	rate := b.rates[arm]
	if rand.Float64() < rate {
		return 1.0
	}
	return 0.0

}

type Agent struct {
	epsilon float64
	ns      []int
	qs      []float64
}

func NewAgent(epsilon float64, actionSize int) *Agent {
	ns := make([]int, 0)
	for i := 0; i < actionSize; i++ {
		ns = append(ns, 0)

	}
	qs := make([]float64, 0.0)
	for i := 0; i < actionSize; i++ {
		qs = append(qs, 0.0)
	}

	return &Agent{
		epsilon: epsilon,
		ns:      ns,
		qs:      qs,
	}
}

func (a *Agent) update(action int, reward float64) {
	a.ns[action] += 1
	a.qs[action] += (reward - a.qs[action]) / float64(len(a.ns))
}

func (a *Agent) getAction() int {
	if rand.Float64() < a.epsilon {
		return rand.Intn(len(a.ns))
	} else {
		presentMax := 0.0
		maxIndex := 0
		for i, q := range a.qs {
			if q >= presentMax {
				presentMax = q
				maxIndex = i
			}
		}
		return maxIndex
	}
}

func main() {
	bandit := NewBandit(10)
	agent := NewAgent(0.05, 10)
	steps := 1000
	totalReward := 0.0

	rates := make([]float64, 0.0)

	for i := 0; i < steps; i++ {
		action := agent.getAction()
		reward := bandit.play(action)
		agent.update(action, reward)

		totalReward += reward

		rates = append(rates, totalReward/float64(i+1))
	}

	fmt.Println("Win Rate")
	fmt.Println(rates)
	fmt.Println("Set rate for each arm")
	fmt.Println(bandit.rates)

}
