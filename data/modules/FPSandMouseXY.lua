-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local M = {}

function M.UpdateFPSandMouseXY()
	FPS_FrameCount = (FPS_FrameCount + 1)
	local currentTime = socket.gettime()
	if ( currentTime > (FPS_LastSecond + 1) ) then
		FPS_TenSeconds[FPS_CurrentSecond] = FPS_FrameCount
		if (FPS_CurrentSecond < 10) then
			FPS_CurrentSecond = (FPS_CurrentSecond + 1)
		else
			FPS_CurrentSecond = 0
		end

		for index = 0, 10 do
			FPS_Average = (FPS_Average + FPS_TenSeconds[index])
		end

		FPS_Average = (FPS_Average / 11)
		FPS_Average = math.floor(FPS_Average)

		FPS_LastSecond = socket.gettime()
		FPS_FrameCount = 0
	end

	label.set_text(   "#MouseXY", "FPS=" .. tostring(FPS_Average) .. " [" .. tostring(  math.floor( mX * (360/WindowWidthTrue) )  ) .. "," .. tostring(  math.floor( mY * (640/WindowHeightTrue) )  ) .. "]"   )	
end

return M
