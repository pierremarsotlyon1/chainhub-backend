package utils

func ForEachBlockRange(from uint64, to uint64, maxRange uint64, handler func(start uint64, end uint64) error) error {
	if maxRange == 0 {
		maxRange = 499
	}
	for start := from + 1; start <= to; start += maxRange + 1 {
		end := min(start+maxRange, to)
		if err := handler(start, end); err != nil {
			return err
		}
	}
	return nil
}
