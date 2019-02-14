-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local MTWO = {}

function MTWO.PlayMusic(music)

	if (music == 0) then msg.post("level:/Audio#Title-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 1) then msg.post("level:/Audio#Cutscenes-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 2) then msg.post("level:/Audio#Level12-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 3) then msg.post("level:/Audio#Level34-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 4) then msg.post("level:/Audio#Level56-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 5) then msg.post("level:/Audio#Level78-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 6) then msg.post("level:/Audio#Level910-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 7) then msg.post("level:/Audio#NewHighScore-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 8) then msg.post("level:/Audio#WinChild-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 9) then msg.post("level:/Audio#WinTeen-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 10) then msg.post("level:/Audio#WinAdult-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	elseif (music == 11) then msg.post("level:/Audio#MsPalser-BGM", "play_sound", {delay = 0, gain = MusicVolume})
	end

	CurrentMusicPlaying = music
	
end

return MTWO