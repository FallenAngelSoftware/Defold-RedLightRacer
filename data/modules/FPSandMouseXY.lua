local UFPSAMXY = {}

function UFPSAMXY.UpdateFPSandMouseXY(dt)
	FPS_FrameCount = FPS_FrameCount + 1
	local currentTime = socket.gettime()
	if ( currentTime > (FPS_LastSecond+1) ) then
		FPS_LastSecond = socket.gettime()
		FPS_TenSeconds[FPS_CurrentSecond] = FPS_FrameCount
		FPS_FrameCount = 0
		if (FPS_CurrentSecond < 9) then
			FPS_CurrentSecond = FPS_CurrentSecond + 1
		else
			FPS_CurrentSecond = 0
		end

		FPS_Average = 0
		local index = 0
		for index = 0, 9 do
			FPS_Average = FPS_Average + FPS_TenSeconds[index]
		end
		FPS_Average = (FPS_Average / 10)
	end

	if (SecretCode[0] == 2 and SecretCode[1] == 7 and SecretCode[2] == 7 and SecretCode[3] == 7) then
		label.set_text(   "#MouseXY", "FPS=" .. tostring( math.floor(FPS_Average) ) .. " [" .. tostring(  math.floor( mX * (360/WindowWidthTrue) )  ) .. "," .. tostring(  math.floor( mY * (640/WindowHeightTrue) )  ) .. "]"   )	
	end
end

return UFPSAMXY