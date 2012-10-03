// Copyright 2012 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rog provides algorithms and data structures for creating roguelike games.

 package main

 import (
     "github.com/ajhager/rog"
     _ "github.com/ajhager/rog/wde"
 )

 func main() {
     rog.Open(20, 11, 2, "rog")
     for rog.IsOpen() {
         rog.Set(5, 5, nil, nil, "Hello, 世界!")
         if rog.Key() == rog.Escape {
             rog.Close()
         }
         rog.Flush()
     }
 }
*/
package rog

var (
	backend Backend
	console *Console
	timing  *stats
)

// SetBackend registers a backend and isn't normally used by the user.
func SetBackend(b Backend) {
	backend = b
}

// IsOpen returns whether the rog window is open or not.
func IsOpen() bool {
	return backend.IsOpen()
}

// Open creates a window and a root console with size width by height cells.
func Open(width, height, zoom int, title string) {
	console = NewConsole(width, height)

	backend.Open(width, height, zoom)
	backend.Name(title)

    timing = NewStats()
}

// Close shuts down the windowing system.
// No rog functions should be called after this.
func Close() {
	backend.Close()
}

// SetTitle changes the title of the window.
func SetTitle(title string) {
	backend.Name(title)
}

// Flush renders the root console to the window.
func Flush() {
	backend.Render(console)
	timing.Update()
}

// Mouse returns a struct representing the state of the mouse.
func Mouse() *MouseData {
	return backend.Mouse()
}

// Key returns the last key typed this frame.
func Key() int {
	return backend.Key()
}

// Dt returns length of the last frame in seconds.
func Dt() float64 {
	return timing.Dt
}

// Fps returns the number of rendered frames per second.
func Fps() int64 {
	return timing.Fps
}

func Blit(con *Console, x, y int) {
	console.Blit(con, x, y)
}

// Set draws on the root console.
func Set(x, y int, fg, bg Blender, data string, rest ...interface{}) {
	console.Set(x, y, fg, bg, data, rest...)
}

// Set draws on the root console with wrapping bounds of x, y, w, h.
func SetR(x, y, w, h int, fg, bg Blender, data string, rest ...interface{}) {
	console.SetR(x, y, w, h, fg, bg, data, rest...)
}

// Get returns the fg, bg colors and rune of the cell on the root console.
func Get(x, y int) (RGB, RGB, rune) {
	return console.Get(x, y)
}

// Fill draws a rect on the root console.
func Fill(x, y, w, h int, fg, bg Blender, ch rune) {
	console.Fill(x, y, w, h, fg, bg, ch)
}

// Clear draws a rect over the entire root console.
func Clear(fg, bg Blender, ch rune) {
	console.Clear(fg, bg, ch)
}

// Width returns the width of the root console in cells.
func Width() int {
	return console.Width()
}

// Height returns the height of the root console in cells.
func Height() int {
	return console.Height()
}
