package primatives

type World struct {
	Elements []Hitable
}

func (w *World) Hit(r Ray, tMin, tMax float64) (bool, HitRecord) {
	hit_anything := false
	closest := tMax
	record := HitRecord{}

	for _, element := range w.Elements {
		hit, tempRecord := element.Hit(r, tMin, closest)
		if hit {
			hit_anything = true
			// this makes sure only the closest image is recorded
			closest = tempRecord.Time
			record = tempRecord
		}

	}
	return hit_anything, record

}

func (w *World) Add(hitable Hitable) {
	w.Elements = append(w.Elements, hitable)
}
