package main

import (
	"fmt"
	"ray_tracer/internal/misc"
	p "ray_tracer/internal/primatives"
	r "ray_tracer/internal/render"
	"time"
)

const (
	nx = 1600
	ny = 800
	ns = 200
)

func createWorld() (*p.World, *p.Camera){
	// world items
	camera := p.NewCamera()
	world := p.World{}

	floor := p.Sphere{
		Center: p.Vector{0, -100.5, -1},
		Radius: 100,
		Material: p.Lambertian{C: p.Vector{0.8, 0.8, 0}}}

	sphere1 := p.Sphere{
		Center: p.Vector{-0.75, 0, -1.5},
		Radius: 0.5,
		Material: p.Lambertian{C: p.Vector{0.8, 0.3, 0}}}

	sphereRight := p.Sphere{
		Center: p.Vector{0.75, 0, -1.5},
		Radius: 0.5,
		Material: p.Metal{C: p.Vector{0.8, 0.6, 0.3}, Fuzz: 0.15}}

	glass := p.Sphere{
		Center: p.Vector{0, 0, -1},
		Radius: 0.5,
		Material: p.Dielectric{C:p.Vector{0.9, 0.9, 0.9}, RefractiveIndex:1.5}}

	world.Add(&sphere1)
	world.Add(&glass)
	world.Add(&sphereRight)
	world.Add(&floor)

	return &world, camera
}



func main() {
	start := time.Now()
	f := misc.OpenFile()
	defer f.Close()

	// create world
	world, camera := createWorld()

	// progress bar
	pgCh := make(chan int, ny)
	go misc.ProgressBar(pgCh, ny)

	// render image
	img := r.Render(world, camera, nx, ny, ns, pgCh)
	misc.WriteFile(f, img)

	fmt.Printf("\nDone.\nElapsed: %v\n", time.Since(start))
}
