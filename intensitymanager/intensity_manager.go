package intensitymanager

type IntensityError string

const (
	InvalidRangeError IntensityError = IntensityError("Invalid range")
)

func (e IntensityError) Error() string {
	return string(e)
}

type Segment struct {
	point     int
	intensity int
}

type IntensityManager struct {
	Segments []Segment
}

// Add increase the intensity in the interval [from, to) by amount
func (pi *IntensityManager) Add(from, to, amount int) error {
	if from > to {
		return InvalidRangeError
	}

	updatedSegments := []Segment{}
	i := 0
	for i < len(pi.Segments) && pi.Segments[i].point < from {
		updatedSegments = append(updatedSegments, pi.Segments[i])
		i++
	}

	if i < len(pi.Segments) && pi.Segments[i].point == from {
		updatedSegments = append(updatedSegments, Segment{from, pi.Segments[i].intensity + amount})
		i++
	} else if len(updatedSegments) > 0 {
		updatedSegments = append(updatedSegments, Segment{from, updatedSegments[len(updatedSegments)-1].intensity + amount})
	} else {
		updatedSegments = append(updatedSegments, Segment{from, amount})
	}

	for i < len(pi.Segments) && pi.Segments[i].point < to {
		updatedSegments = append(updatedSegments, Segment{pi.Segments[i].point, pi.Segments[i].intensity + amount})
		i++
	}

	if i < len(pi.Segments) && pi.Segments[i].point == to {
		updatedSegments = append(updatedSegments, pi.Segments[i:]...)
	} else {
		updatedSegments = append(updatedSegments, Segment{to, updatedSegments[len(updatedSegments)-1].intensity - amount})
		updatedSegments = append(updatedSegments, pi.Segments[i:]...)
	}

	cleanUp(&updatedSegments)

	pi.Segments = updatedSegments
	return nil

}

// Set sets the intensity in the interval [from, to) to amount
func (pi *IntensityManager) Set(from, to, amount int) error {
	if from > to {
		return InvalidRangeError
	}

	updatedSegments := []Segment{}

	toIntensity := 0
	i := 0
	for i < len(pi.Segments) && pi.Segments[i].point < from {
		toIntensity = pi.Segments[i].intensity
		updatedSegments = append(updatedSegments, pi.Segments[i])
		i++
	}

	if len(updatedSegments) == 0 || i < len(pi.Segments) && pi.Segments[i].point != from {
		updatedSegments = append(updatedSegments, Segment{from, amount})
	}

	for i < len(pi.Segments) && pi.Segments[i].point >= from && pi.Segments[i].point < to {
		toIntensity = pi.Segments[i].intensity
		updatedSegments = append(updatedSegments, Segment{pi.Segments[i].point, amount})
		i++
	}

	if i < len(pi.Segments) && pi.Segments[i].point == to {
		updatedSegments = append(updatedSegments, pi.Segments[i:]...)
	} else {
		updatedSegments = append(updatedSegments, Segment{to, toIntensity})
		updatedSegments = append(updatedSegments, pi.Segments[i:]...)
	}

	cleanUp(&updatedSegments)

	pi.Segments = updatedSegments
	return nil

}

func cleanUp(updatedSegments *[]Segment) {
	i := 0
	n := len(*updatedSegments)
	cleanedUpSegments := []Segment{}
	for i < n && (*updatedSegments)[i].intensity == 0 {
		i++
	}

	for i < n {
		if len(cleanedUpSegments) == 0 || (*updatedSegments)[i].intensity != cleanedUpSegments[len(cleanedUpSegments)-1].intensity {
			cleanedUpSegments = append(cleanedUpSegments, (*updatedSegments)[i])
		}
		i++
	}

	*updatedSegments = cleanedUpSegments

}
