package utils

import "time"

// time
const MIN_TO_SEC = uint64(60)
const HOUR_TO_SEC = 60 * MIN_TO_SEC
const DAY_TO_SEC = HOUR_TO_SEC * 24
const WEEK_TO_SEC = 7 * DAY_TO_SEC

// Date de référence (époque Unix)
var referenceTime = time.Unix(0, 0)

func IsDifferentDay(timestamp1, timestamp2 int64) bool {
	// Convertir les timestamps en time.Time
	time1 := time.Unix(timestamp1, 0)
	time2 := time.Unix(timestamp2, 0)

	// Assurez-vous que les deux temps sont dans la même zone horaire
	// pour une comparaison correcte. Modifier `time.UTC` selon les besoins.
	time1 = time1.In(time.UTC)
	time2 = time2.In(time.UTC)

	// Comparer les composantes de la date
	return time1.Year() != time2.Year() || time1.Month() != time2.Month() || time1.Day() != time2.Day()
}

func DaySinceEpoch(timestamp int64) int {
	targetTime := time.Unix(timestamp, 0)

	// Calculer la différence en jours
	duration := targetTime.Sub(referenceTime)
	days := int(duration.Hours() / 24)

	return days
}
