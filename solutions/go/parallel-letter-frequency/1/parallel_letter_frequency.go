package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
  c := make(chan FreqMap)
  m := FreqMap{}

  var wg sync.WaitGroup

  wg.Add(len(texts))

  go func(w *sync.WaitGroup, ch chan FreqMap) {
    defer close(c)

    wg.Wait()
  }(&wg, c)

  for _, text := range(texts) {
    go func (t string) {
      defer wg.Done()
      c <- Frequency(t)
      return
    }(text)
  }

  for fm := range(c) {
    for key, val := range(fm) {
      m[key] += val
    }
  }

  return m
}
