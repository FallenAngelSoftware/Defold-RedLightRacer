local DO = {}

function DO.DisplayOptions()
	local node = gui.get_node("SelectorLine")
	if (ArrowsSelectedByKeyboard == 0) then
		gui.set_position( node, vmath.vector3(180, 560, 1) )
	elseif (ArrowsSelectedByKeyboard == 1) then
		gui.set_position( node, vmath.vector3(180, 510, 1) )
	elseif (ArrowsSelectedByKeyboard == 2) then
		gui.set_position( node, vmath.vector3(180, 430, 1) )
	elseif (ArrowsSelectedByKeyboard == 3) then
		gui.set_position( node, vmath.vector3(180, 240, 1) )
	elseif (ArrowsSelectedByKeyboard == 4) then
		gui.set_position( node, vmath.vector3(180, 190, 1) )
	elseif (ArrowsSelectedByKeyboard == 5) then
		gui.set_position( node, vmath.vector3(180, 140, 1) )
	elseif (ArrowsSelectedByKeyboard == 6) then
		gui.set_position( node, vmath.vector3(180, 90, 1) )
	end

	if (MusicVolume == 1.0) then
		label.set_text("#OptionsMusicVolume", "100%")
	elseif (MusicVolume == 0.75) then
		label.set_text("#OptionsMusicVolume", "75%")
	elseif (MusicVolume == 0.5) then
		label.set_text("#OptionsMusicVolume", "50%")
	elseif (MusicVolume == 0.25) then
		label.set_text("#OptionsMusicVolume", "25%")
	elseif (MusicVolume == 0.0) then
		label.set_text("#OptionsMusicVolume", "0%")
	end	

	if (EffectsVolume == 1.0) then
		label.set_text("#OptionsEffectsVolume", "100%")
	elseif (EffectsVolume == 0.75) then
		label.set_text("#OptionsEffectsVolume", "75%")
	elseif (EffectsVolume == 0.5) then
		label.set_text("#OptionsEffectsVolume", "50%")
	elseif (EffectsVolume == 0.25) then
		label.set_text("#OptionsEffectsVolume", "25%")
	elseif (EffectsVolume == 0.0) then
		label.set_text("#OptionsEffectsVolume", "0%")
	end	

	if (GameMode == 2) then
		label.set_text("#OptionsGameMode", "Adult")
	elseif (GameMode == 1) then
		label.set_text("#OptionsGameMode", "Teen")
	elseif (GameMode == 0) then
		label.set_text("#OptionsGameMode", "Child")
	end

	label.set_text("#SecretCodeOneText", "Secret Code #1:")
	label.set_text("#SecretCodeTwoText", "Secret Code #2:")
	label.set_text("#SecretCodeThreeText", "Secret Code #3:")
	label.set_text("#SecretCodeFourText", "Secret Code #4:")

	label.set_text("#SecretCodeOne", SecretCode[0])
	label.set_text("#SecretCodeTwo", SecretCode[1])
	label.set_text("#SecretCodeThree", SecretCode[2])
	label.set_text("#SecretCodeFour", SecretCode[3])

	if (SecretCodesUnlocked < 500 and Platform == Platform_iOS) then
		label.set_text("#SecretCodeOneText", " ")
		label.set_text("#SecretCodeTwoText", " ")
		label.set_text("#SecretCodeThreeText", " ")
		label.set_text("#SecretCodeFourText", " ")

		label.set_text("#SecretCodeOne", " ")
		label.set_text("#SecretCodeTwo", " ")
		label.set_text("#SecretCodeThree", " ")
		label.set_text("#SecretCodeFour", " ")

		gui.set_position( gui.get_node("SelectorLeft3"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorLeft4"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorLeft5"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorLeft6"), vmath.vector3(-999, -999, 1) )

		gui.set_position( gui.get_node("SelectorRight3"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorRight4"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorRight5"), vmath.vector3(-999, -999, 1) )
		gui.set_position( gui.get_node("SelectorRight6"), vmath.vector3(-999, -999, 1) )							
	end	
end

return DO