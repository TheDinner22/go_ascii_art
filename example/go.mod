module example

go 1.22.2

replace github.com/TheDinner22/image_reader => ../image_reader

replace github.com/TheDinner22/terminal_image_displayer => ../image_displayer

replace todo.com/iNeedToMakeARepoForThis => ../image_processor

require (
	github.com/TheDinner22/image_reader v0.0.0-20240701021525-9be740bdb7cb
	github.com/TheDinner22/terminal_image_displayer v0.0.0-20240701022300-125b9d3d4e79
    todo.com/iNeedToMakeARepoForThis v0
)
