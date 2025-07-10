package analyzer

import (
    "time"
)

func GetArchetype(hourMap map[int]int, dayMap map[time.Weekday]int) string {
    nightCommits := 0
    workdayCommits := 0
    weekendCommits := 0

    for hour, count := range hourMap {
        if hour >= 20 || hour <= 3 {
            nightCommits += count
        }
        if hour >= 9 && hour <= 17 {
            workdayCommits += count
        }
    }

    for day, count := range dayMap {
        if day == time.Saturday || day == time.Sunday {
            weekendCommits += count
        }
    }

    // Apply simple rules
    if nightCommits > workdayCommits && nightCommits > weekendCommits {
        return "🌙 Night Owl"
    } else if weekendCommits > nightCommits && weekendCommits > workdayCommits {
        return "🧪 Weekend Hacker"
    } else if workdayCommits > nightCommits && workdayCommits > weekendCommits {
        return "🧑‍💼 9-to-5 Coder"
    } else {
        return "🌀 Chaos Coder"
    }
}
