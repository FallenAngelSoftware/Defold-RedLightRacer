-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local M = {}

function M.StopMusic(music)
	if (music == 0) then msg.post("level:/Audio#Title-BGM", "stop_sound")
	elseif (music == 1) then msg.post("level:/Audio#Cutscenes-BGM", "stop_sound")
	elseif (music == 2) then msg.post("level:/Audio#Level12-BGM", "stop_sound")
	elseif (music == 3) then msg.post("level:/Audio#Level34-BGM", "stop_sound")
	elseif (music == 4) then msg.post("level:/Audio#Level56-BGM", "stop_sound")
	elseif (music == 5) then msg.post("level:/Audio#Level78-BGM", "stop_sound")
	elseif (music == 6) then msg.post("level:/Audio#Level910-BGM", "stop_sound")
	elseif (music == 7) then msg.post("level:/Audio#NewHighScore-BGM", "stop_sound")
	elseif (music == 8) then msg.post("level:/Audio#WinChild-BGM", "stop_sound")
	elseif (music == 9) then msg.post("level:/Audio#WinTeen-BGM", "stop_sound")
	elseif (music == 10) then msg.post("level:/Audio#WinAdult-BGM", "stop_sound")

	end
end

return M
