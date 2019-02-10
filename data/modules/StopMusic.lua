-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local M = {}

function M.StopMusic(music)
	if (CurrentMusicPlaying == 0) then msg.post("level:/Audio#Title-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 1) then msg.post("level:/Audio#Cutscenes-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 2) then msg.post("level:/Audio#Level12-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 3) then msg.post("level:/Audio#Level34-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 4) then msg.post("level:/Audio#Level56-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 5) then msg.post("level:/Audio#Level78-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 6) then msg.post("level:/Audio#Level910-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 7) then msg.post("level:/Audio#NewHighScore-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 8) then msg.post("level:/Audio#WinChild-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 9) then msg.post("level:/Audio#WinTeen-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 10) then msg.post("level:/Audio#WinAdult-BGM", "stop_sound")
	elseif (CurrentMusicPlaying == 11) then msg.post("level:/Audio#MsPalser-BGM", "stop_sound")
	end
end

return M
