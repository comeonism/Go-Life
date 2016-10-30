# Go-Life

[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in Go.  

## Implementation info: 

1. Since the grid is conceptually infinite - the current implementation maps the end of ledges to the beginnings of the opposite edges. So, for example - a pattern moving off the screen on the left will appear on the right. Just like in good old side-scrollers. Well, some of them. 
2. The 'grid.txt' file provides the initial set up for the grid. Zeroes represent dead cells, ones - live cells. Simple. 
3. Currently there is no fancy graphical representation of the grid. The state of the grid for each cycle gets printed to the console. Each cycle output is separated from the previous one by a line of '===='s.  