package main

// Import the raw raylib rendering hooks and standard string formatting
import (
	"fmt"
	rl "://github.com"
)

// Define your player coordinate data box by hand
type Player struct {
	X, Y, Z float32
}

func main() {
	// 1. Initialize a clean 800x600 desktop window context
	rl.InitWindow(800, 600, "Gorilla - Game Engine Loop")
	defer rl.CloseWindow() // Safeguards your memory cleanup when the engine stops

	// Set up your initial player state tracking
	player := Player{X: 0.0, Y: 0.5, Z: 0.0}

	// 2. Define your 3D Camera coordinates and matrix vectors by hand
	camera := rl.Camera3D{
		Position:   rl.NewVector3(0.0, 5.0, 10.0), // Camera location in space (High and back)
		Target:     rl.NewVector3(0.0, 0.5, 0.0),  // Exact coordinate point the screen is watching
		Up:         rl.NewVector3(0.0, 1.0, 0.0),  // Sets the Y-axis as skyward
		Fovy:       45.0,                          // View lens span angle
		Projection: rl.CameraPerspective,           // Math calculation flag enabling actual 3D depth
	}

	// Cap your hardware execution speed at 60 frames per second
	rl.SetTargetFPS(60)

	fmt.Println("game.go loop successfully initialized.")

	// 3. THE MASTER GRAPHICS & PROCESSING TICK
	for !rl.WindowShouldClose() { // Iterates continuously until the window is exited
		
		// --- 1. MOVEMENT INPUT LOGIC ---
		// Reading hardware keyboard presses by hand (Uses standard float32 steps)
		var speed float32 = 0.1
		
		if rl.IsKeyDown(rl.KeyW) { player.Z -= speed } // Forward
		if rl.IsKeyDown(rl.KeyS) { player.Z += speed } // Backward
		if rl.IsKeyDown(rl.KeyA) { player.X -= speed } // Left
		if rl.IsKeyDown(rl.KeyD) { player.X += speed } // Right

		// Smoothly anchor the camera focal point to follow your player's real-time position
		camera.Target = rl.NewVector3(player.X, player.Y, player.Z)
		camera.Position = rl.NewVector3(player.X, player.Y + 4.5, player.Z + 10.0)

		// --- 2. THE HARDWARE DRAWING PIPELINE ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(26, 26, 46, 255)) // Clears the previous frame canvas with a deep blue hex shade

		// Shift the graphic processor instructions from 2D coordinates into your custom 3D view matrix
		rl.BeginMode3D(camera)

		// Render a flat vector baseline reference grid (Grid Size 50, Step distance 1.0)
		rl.DrawGrid(50, 1.0)

		// Render your Player 3D cube mesh by hand into the spatial coordinates
		// Parameters: (Vector3[X,Y,Z], Width, Height, Depth, HexColor)
		rl.DrawCube(rl.NewVector3(player.X, player.Y, player.Z), 1.0, 1.0, 1.0, rl.Purple)
		rl.DrawCubeWires(rl.NewVector3(player.X, player.Y, player.Z), 1.0, 1.0, 1.0, rl.Black) // Edges look defined

		// Safely escape out of the 3D calculation loop
		rl.EndMode3D()

		// Output current coordinates onto the window frame so you can verify changes visually
		rl.DrawText(fmt.Sprintf("PLAYER POS -> X: %.2f, Z: %.2f", player.X, player.Z), 10, 10, 20, rl.White)

		// Push all processed drawings from memory directly onto the window screen display
		rl.EndDrawing()
	}
}
