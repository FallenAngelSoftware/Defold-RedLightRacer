-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local M = {}

function M.UpdateFPSandMouseXY(dt)
	local frames = {}
	local frame_sum = 0
	local frame_index = 1
	local MAX_SAMPLES = 10
	M.average_ms = 0
	M.raw_average_time = 0
	M.fps = 0
	
	if frames[frame_index] then frame_sum = frame_sum - frames[frame_index] end
	frame_sum = frame_sum + dt
	frames[frame_index] = dt
	frame_index = math.fmod(frame_index + 1, MAX_SAMPLES)
	M.raw_average_time = frame_sum / MAX_SAMPLES
	M.average_ms = math.floor(M.raw_average_time * 10000) / 100
	M.fps = math.floor(1.0 / M.raw_average_time * 100) / 1000 + 1
		
	if (SecretCode[0] == 1 and SecretCode[1] == 0 and SecretCode[2] == 1 and SecretCode[3] == 0) then
		label.set_text(   "#MouseXY", "FPS=" .. tostring( math.floor(M.fps) ) .. " [" .. tostring(  math.floor( mX * (360/WindowWidthTrue) )  ) .. "," .. tostring(  math.floor( mY * (640/WindowHeightTrue) )  ) .. "]"   )	
	end
end

return M
