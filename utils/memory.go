package utils

import "fmt"

func FormatRam(ram uint64) string {
	const unit = 1000
	if ram < unit {
		return fmt.Sprintf("%d B", ram)
	}
	div, exp := int64(unit), 0
	for n := ram / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB",
		float64(ram)/float64(div), "kMGTPE"[exp])
}
