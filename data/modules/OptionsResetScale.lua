local RS = {}

function RS.ResetScale()
	if (DelayAllUserInput == 5) then
		gui.set_scale( gui.get_node("SelectorLeft"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft1"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft2"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft3"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft4"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft5"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorLeft6"), vmath.vector3(1, 1, 1) )
		
		gui.set_scale( gui.get_node("SelectorRight"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight1"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight2"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight3"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight4"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight5"), vmath.vector3(1, 1, 1) )
		gui.set_scale( gui.get_node("SelectorRight6"), vmath.vector3(1, 1, 1) )
	end
end

return RS