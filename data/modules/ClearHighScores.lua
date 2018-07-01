-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local M = {}

function M.ClearHighScores()
	HighScoresName = {}
	for i = 0, 2 do
		HighScoresName[i] = {}
		HighScoresName[i][0] = "The Fallen Angel"
		HighScoresName[i][1] = "Garry Kitchen"
		HighScoresName[i][2] = "Andre' LaMothe"
		HighScoresName[i][3] = "Defold Engine"
		HighScoresName[i][4] = "D.J. Fading Twilight"
		HighScoresName[i][5] = "Kubuntu Linux"
		HighScoresName[i][6] = "NeoPaint"
		HighScoresName[i][7] = "GoldWave"
		HighScoresName[i][8] = "You!"
		HighScoresName[i][9] = "& Them!"
		end

	HighScoresLevel = {}
	for i = 0, 2 do
		HighScoresLevel[i] = {}
		HighScoresLevel[i][0] = 10
		HighScoresLevel[i][1] = 9
		HighScoresLevel[i][2] = 8
		HighScoresLevel[i][3] = 7
		HighScoresLevel[i][4] = 6
		HighScoresLevel[i][5] = 5
		HighScoresLevel[i][6] = 4
		HighScoresLevel[i][7] = 3
		HighScoresLevel[i][8] = 2
		HighScoresLevel[i][9] = 1
	end

	HighScoresScore = {}
	for i = 0, 2 do
		HighScoresScore[i] = {}
		HighScoresScore[i][0] = 10000
		HighScoresScore[i][1] = 9000
		HighScoresScore[i][2] = 8000
		HighScoresScore[i][3] = 7000
		HighScoresScore[i][4] = 6000
		HighScoresScore[i][5] = 5000
		HighScoresScore[i][6] = 4000
		HighScoresScore[i][7] = 3000
		HighScoresScore[i][8] = 2000
		HighScoresScore[i][9] = 1000
	end
end

return M
