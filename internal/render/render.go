package render

import (
	"image"
	"math"
	"math/rand"
	p "ray_tracer/internal/primatives"
	"runtime"
	"sync"
	"time"
)


var (
	white = p.Vector{1.0, 1.0, 1.0}
	blue = p.Vector{0.5, 0.7, 1.0}
)

func color(r p.Ray, world p.Hitable, depth int, rand *rand.Rand) p.Vector {
	hit, record := world.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		if depth < 50 {
			bounced, bouncedRay := record.Bounce(r, record, rand)
			if bounced{
				newColor := color(bouncedRay, world, depth+1, rand)
				return record.Material.Color().Multiply(newColor)
			}
		}
		// depth exceeded - return black
		return p.Vector{}
	}
	return background(r)
}

func background(r p.Ray) p.Vector {
	v := r.Direction.Normalise()
	// scale t to be between 0.0 and 1.0
	t := 0.5 * (v.Y + 1.0)
	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func sample(world *p.World, camera *p.Camera, i, j, nx, ny, ns int, rand *rand.Rand) p.Vector {
	rgb:= p.Vector{}
	for s := 0; s < ns; s++{
		u := (float64(i) + rand.Float64()) / float64(nx)
		v := (float64(j) + rand.Float64())/ float64(ny)

		ray := camera.RayAt(u,v)
		worldColor := color(ray, world, 0, rand)
		rgb = rgb.Add(worldColor)

	}
	return rgb.DivideScalar(float64(ns))
}


func Render(world *p.World, camera *p.Camera, nx, ny, ns int, pgCh chan<- int) *image.NRGBA{
	img := image.NewNRGBA(image.Rect(0, 0, nx, ny))

	// set up worker group
	var wg sync.WaitGroup
	cpus := 2* runtime.NumCPU()

	for core := 0; core < cpus; core++{
		wg.Add(1)
		go func(offset int){
			defer wg.Done()
			// create rng per thread
			rand := rand.New(rand.NewSource(time.Now().Unix()))

			for row := offset; row < ny; row +=cpus {
				for column := 0; column < nx; column++ {
					// draw pixel
					rgb := sample(world, camera, column, row, nx, ny, ns, rand)
					img.Set(column, ny-row-1, rgb)
				}
				// update progress bar
				pgCh <- 1
			}

		}(core)
	}
	wg.Wait()
	return img
}
