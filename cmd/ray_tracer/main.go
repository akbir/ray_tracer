package main

import (
	"fmt"
	"github.com/am-khan/ray_tracer/internal/misc"
	p "github.com/am-khan/ray_tracer/internal/primitives"
	r "github.com/am-khan/ray_tracer/internal/render"
	"time"
)

func createWorld(objectConfigs []misc.ObjectConfig) (*p.World, *p.Camera) {
	// world items
	camera := p.NewCamera()
	world := p.World{}

	for _, config := range objectConfigs {
		tmp_object := p.Sphere{Radius: config.Radius,
			Center: p.Vector{
				config.Center[0], config.Center[1], config.Center[2]}}

		material := config.Material
		switch material.Material {
		case "Lambertian":
			tmp_object.Material = p.Lambertian{
				C: p.Vector{material.Color[0], material.Color[1], material.Color[2]}}

		case "Metal":
			tmp_object.Material = p.Metal{
				C:    p.Vector{material.Color[0], material.Color[1], material.Color[2]},
				Fuzz: material.Property}

		case "Dielectric":
			tmp_object.Material = p.Dielectric{
				C:               p.Vector{material.Color[0], material.Color[1], material.Color[2]},
				RefractiveIndex: material.Property}
		}

		world.Add(&tmp_object)
	}

	return &world, camera
}

func main() {
	start := time.Now()
	f := misc.OpenFile()
	defer f.Close()

	config := misc.GetConfig()

	// create world
	world, camera := createWorld(config.World)

	// progress bar
	pgCh := make(chan int, config.Dimensions.Height)
	go misc.ProgressBar(pgCh, config.Dimensions.Height)

	// render image
	img := r.Render(world, camera, config.Dimensions.Width, config.Dimensions.Height, config.Dimensions.Aliasing, pgCh)
	misc.WriteFile(f, img)

	fmt.Printf("\nDone.\nElapsed: %v\n", time.Since(start))
}
