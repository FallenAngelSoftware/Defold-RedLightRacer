local UFPSAMXY = {}

function UFPSAMXY.UpdateFPSandMouseXY(dt)
	local frames = {}
	local frame_sum = 0
	local frame_index = 1
	local MAX_SAMPLES = 10
	UFPSAMXY.average_ms = 0
	UFPSAMXY.raw_average_time = 0
	UFPSAMXY.fps = 0
	
	if frames[frame_index] then frame_sum = frame_sum - frames[frame_index] end
	frame_sum = frame_sum + dt
	frames[frame_index] = dt
	frame_index = math.fmod(frame_index + 1, MAX_SAMPLES)
	UFPSAMXY.raw_average_time = frame_sum / MAX_SAMPLES
	UFPSAMXY.average_ms = math.floor(UFPSAMXY.raw_average_time * 10000) / 100
	UFPSAMXY.fps = math.floor(1.0 / UFPSAMXY.raw_average_time * 100) / 1000 + 1
		
	if (SecretCode[0] == 1 and SecretCode[1] == 0 and SecretCode[2] == 1 and SecretCode[3] == 0) then
		label.set_text(   "#MouseXY", "FPS=" .. tostring( math.floor(UFPSAMXY.fps) ) .. " [" .. tostring(  math.floor( mX * (360/WindowWidthTrue) )  ) .. "," .. tostring(  math.floor( mY * (640/WindowHeightTrue) )  ) .. "]"   )	
	end
end

return UFPSAMXY