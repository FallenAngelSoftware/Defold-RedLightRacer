components {
  id: "Screen-4-HowToPlay"
  component: "/data/gui/Screen-4-HowToPlay-GUI.gui"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ScreenHeader"
  component: "/data/labels/ScreenHeader.label"
  position {
    x: 180.0
    y: 625.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "Screen4"
  component: "/Screen-4-HowToPlay/Screen4.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "MouseXY"
  component: "/data/labels/MouseXY.label"
  position {
    x: 5.0
    y: 630.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "BG"
  type: "sprite"
  data: "tile_set: \"/data/atlases/TitleFaded-BG.atlas\"\n"
  "default_animation: \"TitleFaded-BG\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 320.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "LineTop"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 607.0
    z: 0.25
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "LineBottom"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 55.0
    z: 0.25
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "HowToPlay"
  type: "label"
  data: "size {\n"
  "  x: 360.0\n"
  "  y: 320.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "scale {\n"
  "  x: 1.0\n"
  "  y: 0.9\n"
  "  z: 1.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "  z: 1.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: true\n"
  "text: \"Control your car by clicking/touching the screen.\\n"
  "\"\n"
  "  \"Clicking/touching the screen will steer the car left.\\n"
  "\"\n"
  "  \"Not clicking/touching the screen will steer the car right.\\n"
  "\"\n"
  "  \"\\n"
  "\"\n"
  "  \"Drive your car to the destination before the time runs out.\\n"
  "\"\n"
  "  \"\\n"
  "\"\n"
  "  \"While driving you must steer your car away from obstacles.\\n"
  "\"\n"
  "  \"If you hit an obstacle your car will slow down.\\n"
  "\"\n"
  "  \"\\n"
  "\"\n"
  "  \"Beat the clock to advance the level.\\n"
  "\"\n"
  "  \"Beat level 10 to win the game!\"\n"
  "font: \"/data/fonts/GameFont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 180.0
    y: 330.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
